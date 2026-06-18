package workflow

import (
	"fmt"
	"maps"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"

	"github.com/github/gh-aw/pkg/console"
	"github.com/github/gh-aw/pkg/constants"
	"github.com/github/gh-aw/pkg/logger"
)

var compilerMainJobLog = logger.New("workflow:compiler_main_job")

func isBuiltinJobName(jobName string) bool {
	_, isBuiltIn := constants.KnownBuiltInJobNames[jobName]
	return isBuiltIn
}

// buildMainJob creates the main agent job that runs the AI agent with the configured engine and tools.
// This job depends on the activation job if it exists, and handles the main workflow logic.
func (c *Compiler) buildMainJob(data *WorkflowData, activationJobCreated bool) (*Job, error) {
workflowLog.Printf("Building main job for workflow: %s", data.Name)

// Build initial steps (setup, checkout, safe-outputs runtime paths)
steps := c.buildMainJobInitialSteps(data)

// Find custom jobs that depend on pre_activation - these are handled by the activation job
customJobsBeforeActivation := c.getCustomJobsDependingOnPreActivation(data.Jobs)

// Compute job-level if: condition
jobCondition := c.buildMainJobJobCondition(data, activationJobCreated, customJobsBeforeActivation)

// Generate main agent job steps
var stepBuilder strings.Builder
if err := c.generateMainJobSteps(&stepBuilder, data); err != nil {
return nil, fmt.Errorf("failed to generate main job steps: %w", err)
}
if stepsContent := stepBuilder.String(); stepsContent != "" {
steps = append(steps, stepsContent)
}

// Build dependencies and engine.env content (used for output wiring and warnings)
depends, engineEnvContent := c.buildMainJobDependencies(data, activationJobCreated)

// Warn when engine.env references built-in job outputs that are not direct dependencies
c.warnBuiltinJobsInEngineEnv(depends, engineEnvContent)

// Build outputs, environment variables, and permissions
outputs, err := c.buildMainJobOutputs(data)
if err != nil {
return nil, err
}
env := c.buildMainJobEnv(data)
permissions, err := c.buildMainJobPermissions(data)
if err != nil {
return nil, err
}

agentConcurrency := GenerateJobConcurrencyConfig(data)

if c.actionMode.IsScript() {
steps = append(steps, c.generateScriptModeCleanupStep())
}

return &Job{
Name:        string(constants.AgentJobName),
If:          jobCondition,
RunsOn:      c.indentYAMLLines(data.RunsOn, "    "),
Environment: c.indentYAMLLines(data.Environment, "    "),
Container:   c.indentYAMLLines(data.Container, "    "),
Services:    c.indentYAMLLines(data.Services, "    "),
Permissions: c.indentYAMLLines(permissions, "    "),
Concurrency: c.indentYAMLLines(agentConcurrency, "    "),
Env:         env,
Steps:       steps,
Needs:       depends,
Outputs:     outputs,
}, nil
}

// buildMainJobInitialSteps creates the setup, checkout, and runtime-path steps for the main job.
func (c *Compiler) buildMainJobInitialSteps(data *WorkflowData) []string {
var steps []string
setupActionRef := c.resolveActionReference("./actions/setup", data)
if setupActionRef != "" || c.actionMode.IsScript() {
steps = append(steps, c.generateCheckoutActionsFolder(data)...)
agentTraceID := fmt.Sprintf("${{ needs.%s.outputs.setup-trace-id }}", constants.ActivationJobName)
agentParentSpanID := setupParentSpanNeedsExpr(constants.ActivationJobName)
steps = append(steps, c.generateSetupStep(data, setupActionRef, SetupActionDestination, false, agentTraceID, agentParentSpanID)...)
}
if data.SafeOutputs != nil {
steps = append(steps, c.generateSetRuntimePathsStep()...)
}
return steps
}

// buildMainJobJobCondition computes the job-level if: condition for the main job.
// When the activation job is present, the condition is delegated to the activation job unless
// it references custom jobs that don't depend on pre_activation.
func (c *Compiler) buildMainJobJobCondition(data *WorkflowData, activationJobCreated bool, customJobsBeforeActivation []string) string {
jobCondition := data.If
if activationJobCreated {
if c.referencesCustomJobOutputs(data.If, data.Jobs) && len(customJobsBeforeActivation) > 0 {
jobCondition = "" // Activation job handles this condition
} else if !c.referencesCustomJobOutputs(data.If, data.Jobs) {
jobCondition = "" // Main job depends on activation job, so no need for inline condition
}
// Note: If data.If references custom jobs that DON'T depend on pre_activation,
// we keep the condition on the agent job
}
if activationJobCreated && hasMaxDailyAICGuardrail(data) {
guard := &ExpressionNode{Expression: fmt.Sprintf("needs.%s.outputs.daily_ai_credits_exceeded != 'true'", constants.ActivationJobName)}
if jobCondition == "" {
jobCondition = RenderCondition(guard)
} else {
jobCondition = RenderCondition(BuildAnd(&ExpressionNode{Expression: stripExpressionWrapper(jobCondition)}, guard))
}
}
return jobCondition
}

// buildMainJobEngineEnvContent concatenates all engine.env values into a single string
// for dependency scanning and built-in job reference warnings.
func buildMainJobEngineEnvContent(data *WorkflowData) string {
if data.EngineConfig == nil || len(data.EngineConfig.Env) == 0 {
return ""
}
var b strings.Builder
for _, envValue := range data.EngineConfig.Env {
b.WriteByte('\n')
b.WriteString(envValue)
}
compilerMainJobLog.Printf("Including %d engine.env values in agent job dependency scan", len(data.EngineConfig.Env))
return b.String()
}

// buildMainJobReferencedDependencies adds any custom jobs referenced in workflow content or
// engine.env as direct dependencies if they are not already in depends.
func (c *Compiler) buildMainJobReferencedDependencies(data *WorkflowData, depends []string, engineEnvContent string) []string {
var contentBuilder strings.Builder
contentBuilder.WriteString(data.MarkdownContent)
if data.CustomSteps != "" {
contentBuilder.WriteByte('\n')
contentBuilder.WriteString(data.CustomSteps)
}
contentBuilder.WriteString(engineEnvContent)
for _, jobName := range c.getReferencedCustomJobs(contentBuilder.String(), data.Jobs) {
if isBuiltinJobName(jobName) {
continue
}
if !slices.Contains(depends, jobName) {
depends = append(depends, jobName)
compilerMainJobLog.Printf("Added direct dependency on custom job '%s' because it's referenced in workflow content or engine.env", jobName)
}
}
return depends
}

// buildMainJobDependencies assembles the needs list for the main job and returns the
// engine.env content string (used by warnBuiltinJobsInEngineEnv).
func (c *Compiler) buildMainJobDependencies(data *WorkflowData, activationJobCreated bool) (depends []string, engineEnvContent string) {
if activationJobCreated {
depends = []string{string(constants.ActivationJobName)}
}
// Add custom jobs that don't depend on pre_activation or agent
if data.Jobs != nil {
for _, jobName := range slices.Sorted(maps.Keys(data.Jobs)) {
if isBuiltinJobName(jobName) {
continue
}
if configMap, ok := data.Jobs[jobName].(map[string]any); ok {
if !jobDependsOnPreActivation(configMap) && !jobDependsOnAgent(configMap) {
depends = append(depends, jobName)
}
}
}
}
engineEnvContent = buildMainJobEngineEnvContent(data)
depends = c.buildMainJobReferencedDependencies(data, depends, engineEnvContent)
return depends, engineEnvContent
}

// warnBuiltinJobsInEngineEnv emits a warning when engine.env values reference built-in job
// outputs via needs expressions that cannot be satisfied at runtime.
func (c *Compiler) warnBuiltinJobsInEngineEnv(depends []string, engineEnvContent string) {
if engineEnvContent == "" {
return
}
builtinNames := make([]string, 0, len(constants.KnownBuiltInJobNames))
for name := range constants.KnownBuiltInJobNames {
builtinNames = append(builtinNames, name)
}
sort.Strings(builtinNames)
builtinsWarned := make(map[string]struct{})
for _, builtinJobName := range builtinNames {
if slices.Contains(depends, builtinJobName) {
continue
}
_, warned := builtinsWarned[builtinJobName]
if !warned && strings.Contains(engineEnvContent, fmt.Sprintf("needs.%s.", builtinJobName)) {
builtinsWarned[builtinJobName] = struct{}{}
warningMsg := fmt.Sprintf(
"engine.env references built-in job '%s' in a needs expression. "+
"Built-in jobs are managed by the compiler and cannot be added as direct agent dependencies; "+
"this expression will silently evaluate to an empty string at runtime.",
builtinJobName,
)
fmt.Fprintln(os.Stderr, console.FormatWarningMessage(warningMsg))
c.IncrementWarningCount()
}
}
}

// buildMainJobBaseOutputsMap constructs the initial outputs map with standard telemetry,
// trace, and optional safe-outputs and artifact prefix keys.
func buildMainJobBaseOutputsMap(data *WorkflowData) map[string]string {
outputs := map[string]string{
"model": "${{ needs.activation.outputs.model }}",
// effective_tokens is the total ET for the run, captured by the MCP gateway log parser step.
"effective_tokens": fmt.Sprintf("${{ steps.%s.outputs.effective_tokens }}", constants.ParseMCPGatewayStepID),
// aic is the total AI Credits cost for the run (1 AIC == 0.01 USD).
"aic": fmt.Sprintf("${{ steps.%s.outputs.aic }}", constants.ParseMCPGatewayStepID),
// ambient_context is the first-request context size metric.
"ambient_context": fmt.Sprintf("${{ steps.%s.outputs.ambient_context }}", constants.ParseMCPGatewayStepID),
// ai_credits_rate_limit_error is true when MCP gateway logs indicate AI credits exhaustion.
"ai_credits_rate_limit_error": fmt.Sprintf("${{ steps.%s.outputs.ai_credits_rate_limit_error || 'false' }}", constants.ParseMCPGatewayStepID),
// unknown_model_ai_credits is true when the AWF API proxy rejects an unknown model.
"unknown_model_ai_credits": fmt.Sprintf("${{ steps.%s.outputs.unknown_model_ai_credits || 'false' }}", constants.ParseMCPGatewayStepID),
"setup-trace-id":    "${{ steps.setup.outputs.trace-id }}",
"setup-span-id":     "${{ steps.setup.outputs.span-id }}",
"setup-parent-span-id": "${{ steps.setup.outputs.parent-span-id || steps.setup.outputs.span-id }}",
}
if hasWorkflowCallTrigger(data.On) {
outputs[constants.ArtifactPrefixOutputName] = "${{ needs.activation.outputs.artifact_prefix }}"
compilerMainJobLog.Print("Added artifact_prefix output to agent job (workflow_call context)")
}
if data.SafeOutputs != nil {
outputs["output"] = "${{ steps.collect_output.outputs.output }}"
outputs["output_types"] = "${{ steps.collect_output.outputs.output_types }}"
outputs["has_patch"] = "${{ steps.collect_output.outputs.has_patch }}"
}
return outputs
}

// buildMainJobOutputs assembles the full job outputs map, adding checkout, cache-memory,
// and engine error detection outputs on top of the base outputs.
func (c *Compiler) buildMainJobOutputs(data *WorkflowData) (map[string]string, error) {
outputs := buildMainJobBaseOutputsMap(data)

if ShouldGeneratePRCheckoutStep(data) {
outputs["checkout_pr_success"] = "${{ steps.checkout-pr.outputs.checkout_pr_success || 'true' }}"
compilerMainJobLog.Print("Added checkout_pr_success output (workflow has contents read access)")
} else {
compilerMainJobLog.Print("Skipped checkout_pr_success output (workflow lacks contents read access)")
}

if data.CacheMemoryConfig != nil && len(data.CacheMemoryConfig.Caches) > 0 {
for i := range data.CacheMemoryConfig.Caches {
stepID := fmt.Sprintf("restore_cache_memory_%d", i)
outputs[fmt.Sprintf("cache_memory_restore_%d_matched_key", i)] = fmt.Sprintf("${{ steps.%s.outputs.cache-matched-key || '' }}", stepID)
outputs[fmt.Sprintf("cache_memory_restore_%d_cache_hit", i)] = fmt.Sprintf("${{ steps.%s.outputs.cache-hit || 'false' }}", stepID)
}
}

engine, engineErr := c.getAgenticEngine(data.AI)
if engineErr == nil && engine.GetErrorDetectionScriptId() != "" {
stepRef := fmt.Sprintf("steps.%s.outputs", constants.DetectAgentErrorsStepID)
outputs["inference_access_error"] = fmt.Sprintf("${{ %s.inference_access_error || 'false' }}", stepRef)
compilerMainJobLog.Printf("Added inference_access_error output (engine=%s, step=%s)", engine.GetID(), constants.DetectAgentErrorsStepID)
outputs["mcp_policy_error"] = fmt.Sprintf("${{ %s.mcp_policy_error || 'false' }}", stepRef)
compilerMainJobLog.Printf("Added mcp_policy_error output (engine=%s, step=%s)", engine.GetID(), constants.DetectAgentErrorsStepID)
outputs["agentic_engine_timeout"] = fmt.Sprintf("${{ %s.agentic_engine_timeout || 'false' }}", stepRef)
compilerMainJobLog.Printf("Added agentic_engine_timeout output (engine=%s, step=%s)", engine.GetID(), constants.DetectAgentErrorsStepID)
outputs["model_not_supported_error"] = fmt.Sprintf("${{ %s.model_not_supported_error || 'false' }}", stepRef)
compilerMainJobLog.Printf("Added model_not_supported_error output (engine=%s, step=%s)", engine.GetID(), constants.DetectAgentErrorsStepID)
}

return outputs, nil
}

// buildMainJobEnv constructs the job-level environment variable map for the main job.
// This includes safe-outputs MCP logging, asset configuration, workflow ID, and UTC offset.
func (c *Compiler) buildMainJobEnv(data *WorkflowData) map[string]string {
var env map[string]string
if data.SafeOutputs != nil {
env = make(map[string]string)
env["GH_AW_MCP_LOG_DIR"] = constants.TmpMcpLogsSafeOutputsDir
// Add asset-related environment variables
if data.SafeOutputs.UploadAssets != nil {
env["GH_AW_ASSETS_BRANCH"] = fmt.Sprintf("%q", data.SafeOutputs.UploadAssets.BranchName)
env["GH_AW_ASSETS_MAX_SIZE_KB"] = strconv.Itoa(data.SafeOutputs.UploadAssets.MaxSizeKB)
env["GH_AW_ASSETS_ALLOWED_EXTS"] = fmt.Sprintf("%q", strings.Join(data.SafeOutputs.UploadAssets.AllowedExts, ","))
} else {
env["GH_AW_ASSETS_BRANCH"] = `""`
env["GH_AW_ASSETS_MAX_SIZE_KB"] = "0"
env["GH_AW_ASSETS_ALLOWED_EXTS"] = `""`
}
env["DEFAULT_BRANCH"] = "${{ github.event.repository.default_branch }}"
}
if data.WorkflowID != "" {
if env == nil {
env = make(map[string]string)
}
env["GH_AW_WORKFLOW_ID_SANITIZED"] = SanitizeWorkflowIDForCacheKey(data.WorkflowID)
}
if utcOffset := c.getCompiledProjectUTCOffset(); utcOffset != "" {
if env == nil {
env = make(map[string]string)
}
env["GH_AW_PROJECT_UTC"] = fmt.Sprintf("%q", utcOffset)
}
return env
}

// gatherAgentJobScripts collects all run: script strings from agent job step sections
// for permission inference and write-command detection.
func gatherAgentJobScripts(data *WorkflowData) []string {
agentJobName := string(constants.AgentJobName)
scripts := extractRunScriptsFromSectionYAML(data.PreSteps, "pre-steps")
scripts = append(scripts, extractRunScriptsFromSectionYAML(data.CustomSteps, "steps")...)
scripts = append(scripts, extractRunScriptsFromSectionYAML(data.PreAgentSteps, "pre-agent-steps")...)
scripts = append(scripts, extractRunScriptsFromSectionYAML(data.PostSteps, "post-steps")...)
if data.Jobs != nil {
scripts = append(scripts, extractRunScriptsFromJobSection(data.Jobs, agentJobName, "setup-steps")...)
scripts = append(scripts, extractRunScriptsFromJobSection(data.Jobs, agentJobName, "pre-steps")...)
}
return scripts
}

// buildMainJobPermissions computes the effective permissions string for the main job,
// including automatic contents:read augmentation and permission inference from run: scripts.
func (c *Compiler) buildMainJobPermissions(data *WorkflowData) (string, error) {
permissions := filterJobLevelPermissions(data.Permissions, data.CachedPermissions)
needsContentsRead := (c.actionMode.IsDev() || c.actionMode.IsScript()) && len(c.generateCheckoutActionsFolder(data)) > 0
if needsContentsRead {
if permissions == "" {
permissions = NewPermissionsContentsRead().RenderToYAML()
} else {
parser := NewPermissionsParser(permissions)
perms := parser.ToPermissions()
if level, exists := perms.Get(PermissionContents); !exists || level == PermissionNone {
perms.Set(PermissionContents, PermissionRead)
permissions = perms.RenderToYAML()
}
}
}
agentAllScripts := gatherAgentJobScripts(data)
if len(agentAllScripts) > 0 {
if writeCmds := detectWriteCommandsInShellScripts(agentAllScripts); len(writeCmds) > 0 {
return "", fmt.Errorf(
"agent job uses write gh command(s) [%s]; write operations are not permitted in agent job steps because the agent job runs with read-only permissions. Use safe-outputs for write operations. See: https://github.github.com/gh-aw/reference/safe-outputs/",
strings.Join(writeCmds, ", "),
)
}
if data.Permissions != "permissions: {}" && permissions != "" {
if inferred := inferPermissionsFromShellScripts(agentAllScripts); len(inferred) > 0 {
permissions = mergeInferredIntoPermissionsYAML(permissions, inferred)
}
}
}
return permissions, nil
}
