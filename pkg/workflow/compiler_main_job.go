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
	steps, err := c.buildMainJobStepList(data)
	if err != nil {
		return nil, err
	}
	jobCondition := c.buildMainJobCondition(data, activationJobCreated)
	depends, engineEnvContent := c.buildMainJobDependencies(data, activationJobCreated)
	c.warnOnBuiltInEngineEnvNeeds(engineEnvContent, depends)
	outputs := c.buildMainJobOutputs(data)
	env := c.buildMainJobEnv(data)
	permissions, err := c.buildMainJobPermissions(data)
	if err != nil {
		return nil, err
	}

	// In script mode, explicitly add a cleanup step (mirrors post.js in dev/release/action mode).
	if c.actionMode.IsScript() {
		steps = append(steps, c.generateScriptModeCleanupStep())
	}

	agentConcurrency := GenerateJobConcurrencyConfig(data)
	job := &Job{
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
	}

	return job, nil
}

func (c *Compiler) buildMainJobStepList(data *WorkflowData) ([]string, error) {
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

	var stepBuilder strings.Builder
	if err := c.generateMainJobSteps(&stepBuilder, data); err != nil {
		return nil, fmt.Errorf("failed to generate main job steps: %w", err)
	}
	if stepsContent := stepBuilder.String(); stepsContent != "" {
		steps = append(steps, stepsContent)
	}
	return steps, nil
}

func (c *Compiler) buildMainJobCondition(data *WorkflowData, activationJobCreated bool) string {
	jobCondition := data.If
	customJobsBeforeActivation := c.getCustomJobsDependingOnPreActivation(data.Jobs)
	if activationJobCreated {
		if c.referencesCustomJobOutputs(data.If, data.Jobs) && len(customJobsBeforeActivation) > 0 {
			jobCondition = ""
		} else if !c.referencesCustomJobOutputs(data.If, data.Jobs) {
			jobCondition = ""
		}
	}
	if activationJobCreated && hasMaxDailyEffectiveTokensGuardrail(data) {
		guard := &ExpressionNode{Expression: fmt.Sprintf("needs.%s.outputs.daily_effective_workflow_exceeded != 'true'", constants.ActivationJobName)}
		if jobCondition == "" {
			return RenderCondition(guard)
		}
		return RenderCondition(BuildAnd(&ExpressionNode{Expression: stripExpressionWrapper(jobCondition)}, guard))
	}
	return jobCondition
}

func (c *Compiler) buildMainJobDependencies(data *WorkflowData, activationJobCreated bool) ([]string, string) {
	depends := c.getMainJobBaseDependencies(data, activationJobCreated)

	contentBuilder := strings.Builder{}
	contentBuilder.WriteString(data.MarkdownContent)
	if data.CustomSteps != "" {
		contentBuilder.WriteByte('\n')
		contentBuilder.WriteString(data.CustomSteps)
	}

	engineEnvContent := c.buildEngineEnvDependencyContent(data)
	contentBuilder.WriteString(engineEnvContent)

	referencedJobs := c.getReferencedCustomJobs(contentBuilder.String(), data.Jobs)
	for _, jobName := range referencedJobs {
		if isBuiltinJobName(jobName) || slices.Contains(depends, jobName) {
			continue
		}
		depends = append(depends, jobName)
		compilerMainJobLog.Printf("Added direct dependency on custom job '%s' because it's referenced in workflow content or engine.env", jobName)
	}
	return depends, engineEnvContent
}

func (c *Compiler) getMainJobBaseDependencies(data *WorkflowData, activationJobCreated bool) []string {
	var depends []string
	if activationJobCreated {
		depends = []string{string(constants.ActivationJobName)}
	}
	if data.Jobs == nil {
		return depends
	}
	for _, jobName := range slices.Sorted(maps.Keys(data.Jobs)) {
		if isBuiltinJobName(jobName) {
			continue
		}
		configMap, ok := data.Jobs[jobName].(map[string]any)
		if ok && !jobDependsOnPreActivation(configMap) && !jobDependsOnAgent(configMap) {
			depends = append(depends, jobName)
		}
	}
	return depends
}

func (c *Compiler) buildEngineEnvDependencyContent(data *WorkflowData) string {
	if data.EngineConfig == nil || len(data.EngineConfig.Env) == 0 {
		return ""
	}
	var engineEnvBuilder strings.Builder
	for _, envValue := range data.EngineConfig.Env {
		engineEnvBuilder.WriteByte('\n')
		engineEnvBuilder.WriteString(envValue)
	}
	compilerMainJobLog.Printf("Including %d engine.env values in agent job dependency scan", len(data.EngineConfig.Env))
	return engineEnvBuilder.String()
}

func (c *Compiler) warnOnBuiltInEngineEnvNeeds(engineEnvContent string, depends []string) {
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
		if _, warned := builtinsWarned[builtinJobName]; warned {
			continue
		}
		if !strings.Contains(engineEnvContent, fmt.Sprintf("needs.%s.", builtinJobName)) {
			continue
		}
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

func (c *Compiler) buildMainJobOutputs(data *WorkflowData) map[string]string {
	outputs := map[string]string{
		"model":            "${{ needs.activation.outputs.model }}",
		"effective_tokens": fmt.Sprintf("${{ steps.%s.outputs.effective_tokens }}", constants.ParseMCPGatewayStepID),
		"aic":              fmt.Sprintf("${{ steps.%s.outputs.aic }}", constants.ParseMCPGatewayStepID),
		"ambient_context":  fmt.Sprintf("${{ steps.%s.outputs.ambient_context }}", constants.ParseMCPGatewayStepID),
		"effective_tokens_rate_limit_error": fmt.Sprintf(
			"${{ steps.%s.outputs.effective_tokens_rate_limit_error || 'false' }}",
			constants.ParseMCPGatewayStepID,
		),
		"setup-trace-id":       "${{ steps.setup.outputs.trace-id }}",
		"setup-span-id":        "${{ steps.setup.outputs.span-id }}",
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
	if ShouldGeneratePRCheckoutStep(data) {
		outputs["checkout_pr_success"] = "${{ steps.checkout-pr.outputs.checkout_pr_success || 'true' }}"
		compilerMainJobLog.Print("Added checkout_pr_success output (workflow has contents read access)")
	} else {
		compilerMainJobLog.Print("Skipped checkout_pr_success output (workflow lacks contents read access)")
	}

	engine, err := c.getAgenticEngine(data.AI)
	if err != nil || engine.GetErrorDetectionScriptId() == "" {
		return outputs
	}
	stepRef := fmt.Sprintf("steps.%s.outputs", constants.DetectAgentErrorsStepID)
	outputs["inference_access_error"] = fmt.Sprintf("${{ %s.inference_access_error || 'false' }}", stepRef)
	outputs["mcp_policy_error"] = fmt.Sprintf("${{ %s.mcp_policy_error || 'false' }}", stepRef)
	outputs["agentic_engine_timeout"] = fmt.Sprintf("${{ %s.agentic_engine_timeout || 'false' }}", stepRef)
	outputs["model_not_supported_error"] = fmt.Sprintf("${{ %s.model_not_supported_error || 'false' }}", stepRef)
	compilerMainJobLog.Printf("Added engine error outputs (engine=%s, step=%s)", engine.GetID(), constants.DetectAgentErrorsStepID)
	return outputs
}

func (c *Compiler) buildMainJobEnv(data *WorkflowData) map[string]string {
	var env map[string]string
	if data.SafeOutputs != nil {
		env = map[string]string{
			"GH_AW_MCP_LOG_DIR": "/tmp/gh-aw/mcp-logs/safeoutputs",
			"DEFAULT_BRANCH":    "${{ github.event.repository.default_branch }}",
		}
		if data.SafeOutputs.UploadAssets != nil {
			env["GH_AW_ASSETS_BRANCH"] = fmt.Sprintf("%q", data.SafeOutputs.UploadAssets.BranchName)
			env["GH_AW_ASSETS_MAX_SIZE_KB"] = strconv.Itoa(data.SafeOutputs.UploadAssets.MaxSizeKB)
			env["GH_AW_ASSETS_ALLOWED_EXTS"] = fmt.Sprintf("%q", strings.Join(data.SafeOutputs.UploadAssets.AllowedExts, ","))
		} else {
			env["GH_AW_ASSETS_BRANCH"] = `""`
			env["GH_AW_ASSETS_MAX_SIZE_KB"] = "0"
			env["GH_AW_ASSETS_ALLOWED_EXTS"] = `""`
		}
	}
	if data.WorkflowID != "" {
		if env == nil {
			env = make(map[string]string)
		}
		env["GH_AW_WORKFLOW_ID_SANITIZED"] = SanitizeWorkflowIDForCacheKey(data.WorkflowID)
	}
	return env
}

func (c *Compiler) buildMainJobPermissions(data *WorkflowData) (string, error) {
	permissions := filterJobLevelPermissions(data.Permissions, data.CachedPermissions)
	needsContentsRead := (c.actionMode.IsDev() || c.actionMode.IsScript()) && len(c.generateCheckoutActionsFolder(data)) > 0
	permissions = ensureMainJobContentsRead(permissions, needsContentsRead)

	agentAllScripts := extractMainJobScripts(data)
	if len(agentAllScripts) == 0 {
		return permissions, nil
	}
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
	return permissions, nil
}

func ensureMainJobContentsRead(permissions string, needsContentsRead bool) string {
	if !needsContentsRead {
		return permissions
	}
	if permissions == "" {
		return NewPermissionsContentsRead().RenderToYAML()
	}
	parser := NewPermissionsParser(permissions)
	perms := parser.ToPermissions()
	if level, exists := perms.Get(PermissionContents); !exists || level == PermissionNone {
		perms.Set(PermissionContents, PermissionRead)
		return perms.RenderToYAML()
	}
	return permissions
}

func extractMainJobScripts(data *WorkflowData) []string {
	agentAllScripts := extractRunScriptsFromSectionYAML(data.PreSteps, "pre-steps")
	agentAllScripts = append(agentAllScripts, extractRunScriptsFromSectionYAML(data.CustomSteps, "steps")...)
	agentAllScripts = append(agentAllScripts, extractRunScriptsFromSectionYAML(data.PreAgentSteps, "pre-agent-steps")...)
	agentAllScripts = append(agentAllScripts, extractRunScriptsFromSectionYAML(data.PostSteps, "post-steps")...)
	if data.Jobs == nil {
		return agentAllScripts
	}
	agentJobName := string(constants.AgentJobName)
	agentAllScripts = append(agentAllScripts, extractRunScriptsFromJobSection(data.Jobs, agentJobName, "setup-steps")...)
	agentAllScripts = append(agentAllScripts, extractRunScriptsFromJobSection(data.Jobs, agentJobName, "pre-steps")...)
	return agentAllScripts
}
