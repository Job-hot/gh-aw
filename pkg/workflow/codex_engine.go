package workflow

import (
	"fmt"
	"maps"
	"regexp"
	"sort"
	"strings"

	"github.com/github/gh-aw/pkg/constants"
	"github.com/github/gh-aw/pkg/logger"
	"github.com/github/gh-aw/pkg/workflow/compilerenv"
)

var codexEngineLog = logger.New("workflow:codex_engine")

// detectionResponseSchema is the JSON Schema for Codex detection runs.
// It constrains the model output to exactly the threat detection result fields.
// The schema is written to detectionSchemaFilePath before Codex runs and passed
// via --output-schema; the structured result is written to detectionResultFilePath
// via --output-last-message for direct parsing without log scraping.
const detectionResponseSchema = `{"type":"object","properties":{"prompt_injection":{"type":"boolean"},"secret_leak":{"type":"boolean"},"malicious_patch":{"type":"boolean"},"reasons":{"type":"array","items":{"type":"string"}}},"required":["prompt_injection","secret_leak","malicious_patch","reasons"],"additionalProperties":false}`

// detectionSchemaFilePath is the path where the detection JSON schema is written
// before Codex runs. It is referenced by --output-schema.
const detectionSchemaFilePath = "/tmp/gh-aw/threat-detection/detection_schema.json"

// detectionResultFilePath is the path where Codex writes the final structured
// verdict via --output-last-message. The parser reads this file directly instead
// of scraping the log stream, eliminating false parse_error warnings from noisy
// SSE/tracing output.
const detectionResultFilePath = "/tmp/gh-aw/threat-detection/detection_result.json"

// Pre-compiled regexes for Codex log parsing (performance optimization)
var (
	codexToolCallOldFormat    = regexp.MustCompile(`\] tool ([^(]+)\(`)
	codexToolCallNewFormat    = regexp.MustCompile(`^tool ([^(]+)\(`)
	codexExecCommandOldFormat = regexp.MustCompile(`\] exec (.+?) in`)
	codexExecCommandNewFormat = regexp.MustCompile(`^exec (.+?) in`)
	codexDurationPattern      = regexp.MustCompile(`in\s+(\d+(?:\.\d+)?)\s*s`)
	codexTokenUsagePattern    = regexp.MustCompile(`(?i)tokens\s+used[:\s]+(\d+)`)
	codexTotalTokensPattern   = regexp.MustCompile(`total_tokens:\s*(\d+)`)
)

// CodexEngine represents the Codex agentic engine
type CodexEngine struct {
	BaseEngine
}

func NewCodexEngine() *CodexEngine {
	return &CodexEngine{
		BaseEngine: BaseEngine{
			id:           "codex",
			displayName:  "Codex",
			description:  "Uses OpenAI Codex CLI with MCP server support",
			experimental: false,
			capabilities: EngineCapabilities{
				ToolsAllowlist:   true,
				MaxTurns:         true,  // AWF max-turns is supported for Codex runs
				MaxContinuations: false, // Codex does not support --max-autopilot-continues-style continuation mode
				WebSearch:        true,  // Codex has built-in web-search support
				NativeAgentFile:  false, // Codex does not support agent file natively; the compiler prepends the agent file content to prompt.txt
			},
			dedicatedLLMGatewayPort: constants.CodexLLMGatewayPort,
		},
	}
}

// GetModelEnvVarName returns an empty string because the Codex CLI does not support
// selecting the model via a native environment variable. Model selection for Codex
// is done via the --model flag in the shell command.
func (e *CodexEngine) GetModelEnvVarName() string {
	return ""
}

// GetRequiredSecretNames returns the list of secrets required by the Codex engine
// This includes CODEX_API_KEY, OPENAI_API_KEY, and optionally MCP_GATEWAY_API_KEY and mcp-scripts secrets
func (e *CodexEngine) GetRequiredSecretNames(workflowData *WorkflowData) []string {
	return append([]string{"CODEX_API_KEY", "OPENAI_API_KEY"}, collectCommonMCPSecrets(workflowData)...)
}

// GetSecretValidationStep returns the secret validation step for the Codex engine.
// Returns an empty step if custom command is specified.
func (e *CodexEngine) GetSecretValidationStep(workflowData *WorkflowData) GitHubActionStep {
	return BuildDefaultSecretValidationStep(
		workflowData,
		[]string{"CODEX_API_KEY", "OPENAI_API_KEY"},
		"Codex",
		"https://github.github.com/gh-aw/reference/engines/#openai-codex",
	)
}

func (e *CodexEngine) GetInstallationSteps(workflowData *WorkflowData) []GitHubActionStep {
	codexEngineLog.Printf("Generating installation steps for Codex engine: workflow=%s", workflowData.Name)

	// Skip installation if custom command is specified
	if workflowData.EngineConfig != nil && workflowData.EngineConfig.Command != "" {
		codexEngineLog.Printf("Skipping installation steps: custom command specified (%s)", workflowData.EngineConfig.Command)
		return []GitHubActionStep{}
	}

	steps := BuildStandardNpmEngineInstallStepsNoCooldown(
		"@openai/codex",
		string(constants.DefaultCodexVersion),
		"Install Codex CLI",
		"codex",
		workflowData,
	)

	// Add AWF installation step if firewall is enabled
	if isFirewallEnabled(workflowData) {
		firewallConfig := getFirewallConfig(workflowData)
		agentConfig := getAgentConfig(workflowData)
		var awfVersion string
		if firewallConfig != nil {
			awfVersion = firewallConfig.Version
		}

		// Install AWF binary (or skip if custom command is specified)
		awfInstall := generateAWFInstallationStep(awfVersion, agentConfig)
		if len(awfInstall) > 0 {
			steps = append(steps, awfInstall)
		}
	}

	return steps
}

// GetDeclaredOutputFiles returns the output files that Codex may produce.
// Use /tmp/gh-aw for Codex runtime logs because ${RUNNER_TEMP}/gh-aw is
// mounted read-only inside the AWF chroot sandbox.
func (e *CodexEngine) GetDeclaredOutputFiles() []string {
	// Return the Codex log directory for artifact collection.
	return []string{
		constants.TmpMcpConfigLogsDir,
	}
}

// GetAgentManifestFiles returns Codex-specific instruction files that should be
// treated as security-sensitive manifests.  AGENTS.md is the primary OpenAI
// Codex agent-instruction file; modifying it can redirect agent behaviour.
// CLAUDE.md and GEMINI.md are also listed because repositories often use multiple
// engines and Codex runs alongside them.
func (e *CodexEngine) GetAgentManifestFiles() []string {
	return []string{"AGENTS.md", "CLAUDE.md", "GEMINI.md"}
}

// GetAgentManifestPathPrefixes returns Codex-specific config directory prefixes.
// The .codex/ directory can contain agent configuration and task-specific settings.
func (e *CodexEngine) GetAgentManifestPathPrefixes() []string {
	return []string{".codex/"}
}

// GetHarnessScriptName returns the filename of the JavaScript harness script that wraps
// Codex CLI execution with retry logic for transient OpenAI API errors.
func (e *CodexEngine) GetHarnessScriptName() string {
	return "codex_harness.cjs"
}

// buildCodexModelConfig determines the model parameter string and environment variable name
// to use for model selection in the Codex CLI command.
func buildCodexModelConfig(workflowData *WorkflowData) (modelParam, modelEnvVar string, isDetectionJob bool) {
	isDetectionJob = workflowData.SafeOutputs == nil
	if isDetectionJob {
		modelEnvVar = constants.EnvVarModelDetectionCodex
	} else {
		modelEnvVar = constants.EnvVarModelAgentCodex
	}
	modelParam = fmt.Sprintf(`${%s:+ --model "$%s"}`, modelEnvVar, modelEnvVar)
	return
}

// buildCodexExecutionFlags builds the web-search, web-fetch and execution-policy CLI flags.
func buildCodexExecutionFlags(workflowData *WorkflowData, firewallEnabled bool) (webSearchParam, webFetchParam, executionPolicyParam string) {
	webSearchParam = ` -c web_search="disabled"`
	if workflowData.ParsedTools != nil && workflowData.ParsedTools.WebSearch != nil {
		webSearchParam = ""
	}
	webFetchParam = ` -c fetch="disabled"`
	if workflowData.ParsedTools != nil && workflowData.ParsedTools.WebFetch != nil {
		webFetchParam = ""
	}
	executionPolicyParam = ` --sandbox workspace-write --skip-git-repo-check -c approval_policy="never" `
	if firewallEnabled {
		executionPolicyParam = " --dangerously-bypass-approvals-and-sandbox --skip-git-repo-check "
	}
	return
}

// buildCodexStructuredOutputConfig returns the --output-schema and -o flags for detection runs,
// and the shell command that writes the schema file before Codex starts.
func buildCodexStructuredOutputConfig(workflowData *WorkflowData) (structuredOutputParam, detectionSchemaWriteCmd string) {
	if !workflowData.IsDetectionRun {
		return
	}
	structuredOutputParam = fmt.Sprintf(` --output-schema %s -o %s`, detectionSchemaFilePath, detectionResultFilePath)
	detectionSchemaWriteCmd = fmt.Sprintf("mkdir -p /tmp/gh-aw/threat-detection && printf '%%s' '%s' > %s", detectionResponseSchema, detectionSchemaFilePath)
	codexEngineLog.Printf("Enabling structured outputs for Codex detection run")
	return
}

// buildCodexCustomArgsParam joins any user-specified extra CLI args into a single string.
func buildCodexCustomArgsParam(workflowData *WorkflowData) string {
	if workflowData.EngineConfig == nil || len(workflowData.EngineConfig.Args) == 0 {
		return ""
	}
	var sb strings.Builder
	for _, arg := range workflowData.EngineConfig.Args {
		sb.WriteString(arg + " ")
	}
	return sb.String()
}

// buildCodexBaseCommand returns the command name and harness script name to use for Codex execution.
func (e *CodexEngine) buildCodexBaseCommand(workflowData *WorkflowData) (commandName, harnessScriptName string) {
	commandName = "codex"
	if workflowData.EngineConfig != nil && workflowData.EngineConfig.Command != "" {
		commandName = workflowData.EngineConfig.Command
		codexEngineLog.Printf("Using custom command: %s", commandName)
	}
	harnessScriptName = e.GetHarnessScriptName()
	if workflowData.EngineConfig != nil && workflowData.EngineConfig.HarnessScript != "" {
		harnessScriptName = workflowData.EngineConfig.HarnessScript
		codexEngineLog.Printf("Using custom harness script: %s", harnessScriptName)
	}
	return
}

// assembleCodexCLICommand combines the base command with all parameter strings into the final CLI command.
func assembleCodexCLICommand(commandName, harnessScriptName, modelParam, webSearchParam, webFetchParam, executionPolicyParam, structuredOutputParam, customArgsParam string) string {
	if harnessScriptName != "" {
		execPrefix := fmt.Sprintf(`%s %s/%s %s`, nodeRuntimeResolutionCommand, SetupActionDestinationShell, harnessScriptName, commandName)
		return fmt.Sprintf("%s exec%s%s%s%s%s%s --prompt-file /tmp/gh-aw/aw-prompts/prompt.txt",
			execPrefix, modelParam, webSearchParam, webFetchParam, executionPolicyParam, structuredOutputParam, customArgsParam)
	}
	return fmt.Sprintf("%s exec%s%s%s%s%s%s \"$INSTRUCTION\"",
		commandName, modelParam, webSearchParam, webFetchParam, executionPolicyParam, structuredOutputParam, customArgsParam)
}

// buildCodexCommandWithPathSetup prepends the npm PATH setup (and optional MCP CLI path) to the codex command.
func buildCodexCommandWithPathSetup(workflowData *WorkflowData, codexCommand, harnessScriptName string) string {
	npmPathSetup := GetNpmBinPathSetup()
	var codexCommandWithSetup string
	if harnessScriptName != "" {
		codexCommandWithSetup = fmt.Sprintf(`%s && %s`, npmPathSetup, codexCommand)
	} else {
		codexCommandWithSetup = fmt.Sprintf(`%s && INSTRUCTION="$(cat /tmp/gh-aw/aw-prompts/prompt.txt)" && %s`, npmPathSetup, codexCommand)
	}
	if mcpCLIPath := GetMCPCLIPathSetup(workflowData); mcpCLIPath != "" {
		codexCommandWithSetup = fmt.Sprintf("%s && %s", mcpCLIPath, codexCommandWithSetup)
	}
	return codexCommandWithSetup
}

// buildCodexAWFCommand builds the AWF-wrapped command for Codex execution when the firewall is enabled.
func buildCodexAWFCommand(workflowData *WorkflowData, codexCommand, logFile string, harnessScriptName, detectionSchemaWriteCmd string) string {
	var allowedDomains string
	if workflowData.CachedAllowedDomainsComputed {
		allowedDomains = workflowData.CachedAllowedDomainsStr
	} else {
		allowedDomains = GetAllowedDomainsForEngine(constants.CodexEngine, workflowData.NetworkPermissions, workflowData.Tools, workflowData.Runtimes)
	}
	if workflowData.EngineConfig != nil && workflowData.EngineConfig.APITarget != "" {
		allowedDomains = mergeAPITargetDomains(allowedDomains, workflowData.EngineConfig.APITarget)
	}

	pathSetup := "mkdir -p \"$CODEX_HOME/logs\" && touch " + AgentStepSummaryPath
	if workflowData.IsDetectionRun {
		pathSetup = pathSetup + " && " + detectionSchemaWriteCmd
	}

	return BuildAWFCommand(AWFCommandConfig{
		EngineName:         "codex",
		EngineCommand:      buildCodexCommandWithPathSetup(workflowData, codexCommand, harnessScriptName),
		LogFile:            logFile,
		WorkflowData:       workflowData,
		UsesTTY:            false,
		AllowedDomains:     allowedDomains,
		PathSetup:          pathSetup,
		ExcludeEnvVarNames: ComputeAWFExcludeEnvVarNames(workflowData, []string{"CODEX_API_KEY", "OPENAI_API_KEY"}),
	})
}

// buildCodexPlainCommand builds the non-AWF shell command for Codex execution.
func buildCodexPlainCommand(codexCommand, logFile string, harnessScriptName, detectionSchemaWriteCmd string) string {
	schemaWritePrefix := ""
	if detectionSchemaWriteCmd != "" {
		schemaWritePrefix = detectionSchemaWriteCmd + " && "
	}
	if harnessScriptName != "" {
		return fmt.Sprintf(`set -o pipefail
printf '%%s' "$(date +%%s%%3N)" > %s
touch %s
(umask 177 && touch %s)
mkdir -p "$CODEX_HOME/logs"
%s%s 2>&1 | tee %s`, AgentCLIStartMsPath, AgentStepSummaryPath, logFile, schemaWritePrefix, codexCommand, logFile)
	}
	return fmt.Sprintf(`set -o pipefail
printf '%%s' "$(date +%%s%%3N)" > %s
touch %s
(umask 177 && touch %s)
INSTRUCTION="$(cat "$GH_AW_PROMPT")"
mkdir -p "$CODEX_HOME/logs"
%s%s 2>&1 | tee %s`, AgentCLIStartMsPath, AgentStepSummaryPath, logFile, schemaWritePrefix, codexCommand, logFile)
}

// buildCodexBaseEnv builds the base environment variable map for a Codex execution step.
func buildCodexBaseEnv(workflowData *WorkflowData, firewallEnabled bool) map[string]string {
	effectiveGitHubToken := getEffectiveGitHubToken("")
	env := map[string]string{
		"CODEX_API_KEY":       "${{ secrets.CODEX_API_KEY || secrets.OPENAI_API_KEY }}",
		"GITHUB_STEP_SUMMARY": AgentStepSummaryPath,
		"GH_AW_PROMPT":        constants.AwPromptsFile,
		// Tag the step as a GitHub AW agentic execution for discoverability by agents
		"GITHUB_AW":        "true",
		"RUNNER_TEMP":      "${{ runner.temp }}",
		"GH_AW_MCP_CONFIG": constants.CodexMcpConfigTomlPath,
		// Keep Codex runtime state in /tmp/gh-aw because ${RUNNER_TEMP}/gh-aw is
		// mounted read-only inside the AWF chroot sandbox.
		"CODEX_HOME": constants.TmpMcpConfigDir,
		// Enable verbose RUST_LOG only in debug mode.
		"RUST_LOG":                     "${{ runner.debug == 1 && 'trace,hyper_util=info,mio=info,reqwest=info,os_info=info,codex_otel=warn,codex_core=debug,ocodex_exec=debug' || 'warn' }}",
		"GH_AW_GITHUB_TOKEN":           effectiveGitHubToken,
		"GITHUB_PERSONAL_ACCESS_TOKEN": effectiveGitHubToken,
		"OPENAI_API_KEY":               "${{ secrets.CODEX_API_KEY || secrets.OPENAI_API_KEY }}",
	}
	injectWorkflowCallNetworkAllowedEnv(env, workflowData)
	if workflowData.IsDetectionRun {
		env["GH_AW_PHASE"] = "detection"
	} else {
		env["GH_AW_PHASE"] = "agent"
	}
	if IsRelease() {
		env["GH_AW_VERSION"] = GetVersion()
	} else {
		env["GH_AW_VERSION"] = "dev"
	}
	// Add GH_AW_SAFE_OUTPUTS if output is needed
	applySafeOutputEnvToMap(env, workflowData)
	// Propagate W3C trace context so engine spans nest under the gh-aw.agent.setup span.
	applyTraceContextEnvToMap(env)
	// In sandbox (AWF) mode, set git identity environment variables so the first git commit succeeds.
	if firewallEnabled {
		maps.Copy(env, getGitIdentityEnvVars())
	}
	return env
}

// applyCodexModelAndCustomEnv sets the max-turns, model, custom engine/agent env vars and
// mcp-scripts secrets on the given env map for a Codex execution step.
func applyCodexModelAndCustomEnv(env map[string]string, workflowData *WorkflowData, firewallEnabled bool, modelConfigured bool, modelEnvVar string) {
	if workflowData.ToolsStartupTimeout != "" {
		env["GH_AW_STARTUP_TIMEOUT"] = workflowData.ToolsStartupTimeout
	}
	if workflowData.ToolsTimeout != "" {
		env["GH_AW_TOOL_TIMEOUT"] = workflowData.ToolsTimeout
	}
	if workflowData.EngineConfig != nil && workflowData.EngineConfig.MaxTurns != "" {
		env["GH_AW_MAX_TURNS"] = workflowData.EngineConfig.MaxTurns
	} else {
		env["GH_AW_MAX_TURNS"] = compilerenv.BuildDefaultMaxTurnsExpression()
	}
	if modelConfigured {
		codexEngineLog.Printf("Setting %s env var for model: %s", modelEnvVar, workflowData.EngineConfig.Model)
		env[modelEnvVar] = workflowData.EngineConfig.Model
	} else {
		env[modelEnvVar] = compilerenv.BuildModelOverrideExpression(modelEnvVar, compilerenv.DefaultModelCodex, constants.CodexDefaultModel)
	}
	if workflowData.EngineConfig != nil && len(workflowData.EngineConfig.Env) > 0 {
		maps.Copy(env, workflowData.EngineConfig.Env)
	}
	agentConfig := getAgentConfig(workflowData)
	if agentConfig != nil && len(agentConfig.Env) > 0 {
		maps.Copy(env, agentConfig.Env)
		codexEngineLog.Printf("Added %d custom env vars from agent config", len(agentConfig.Env))
	}
	if IsMCPScriptsEnabled(workflowData.MCPScripts) {
		mcpScriptsSecrets := collectMCPScriptsSecrets(workflowData.MCPScripts)
		for varName, secretExpr := range mcpScriptsSecrets {
			if _, exists := env[varName]; !exists {
				env[varName] = secretExpr
			}
		}
	}
}

// GetExecutionSteps returns the GitHub Actions steps for executing Codex
func (e *CodexEngine) GetExecutionSteps(workflowData *WorkflowData, logFile string) []GitHubActionStep {
	modelConfigured := workflowData.EngineConfig != nil && workflowData.EngineConfig.Model != ""
	firewallEnabled := isFirewallEnabled(workflowData)
	codexEngineLog.Printf("Building Codex execution steps: workflow=%s, modelConfigured=%v, firewall=%v",
		workflowData.Name, modelConfigured, firewallEnabled)

	var steps []GitHubActionStep

	modelParam, modelEnvVar, _ := buildCodexModelConfig(workflowData)
	webSearchParam, webFetchParam, executionPolicyParam := buildCodexExecutionFlags(workflowData, firewallEnabled)
	structuredOutputParam, detectionSchemaWriteCmd := buildCodexStructuredOutputConfig(workflowData)
	customArgsParam := buildCodexCustomArgsParam(workflowData)
	commandName, harnessScriptName := e.buildCodexBaseCommand(workflowData)
	codexCommand := assembleCodexCLICommand(commandName, harnessScriptName, modelParam, webSearchParam, webFetchParam, executionPolicyParam, structuredOutputParam, customArgsParam)

	var command string
	if firewallEnabled {
		command = buildCodexAWFCommand(workflowData, codexCommand, logFile, harnessScriptName, detectionSchemaWriteCmd)
	} else {
		command = buildCodexPlainCommand(codexCommand, logFile, harnessScriptName, detectionSchemaWriteCmd)
	}

	env := buildCodexBaseEnv(workflowData, firewallEnabled)
	applyCodexModelAndCustomEnv(env, workflowData, firewallEnabled, modelConfigured, modelEnvVar)

	stepLines := []string{
		"      - name: Execute Codex CLI",
		"        id: agentic_execution",
	}
	allowedSecrets := e.GetRequiredSecretNames(workflowData)
	filteredEnv := FilterEnvForSecrets(env, allowedSecrets)
	addCliProxyGHTokenToEnv(filteredEnv, workflowData)
	stepLines = FormatStepWithCommandAndEnv(stepLines, command, filteredEnv)
	steps = append(steps, GitHubActionStep(stepLines))
	return steps
}

// GetSquidLogsSteps returns the steps for uploading and parsing Squid logs (after secret redaction)
func (e *CodexEngine) GetSquidLogsSteps(workflowData *WorkflowData) []GitHubActionStep {
	return defaultGetSquidLogsSteps(workflowData, codexEngineLog)
}

// applyPlaywrightToolToCodexConfig updates the result ToolsConfig with the playwright tool
// configuration derived from the source toolsConfig.
func applyPlaywrightToolToCodexConfig(toolsConfig, result *ToolsConfig) {
	if toolsConfig.Playwright == nil {
		return
	}
	playwrightConfig := &PlaywrightToolConfig{
		Version: toolsConfig.Playwright.Version,
		Args:    toolsConfig.Playwright.Args,
		Mode:    toolsConfig.Playwright.Mode,
	}
	result.Playwright = playwrightConfig
	// In CLI mode, playwright is not an MCP server — remove from raw map and skip MCP config entry.
	if playwrightConfig.IsCLIMode() {
		delete(result.raw, "playwright")
		return
	}
	// Also update the Custom map entry for playwright with allowed tools list
	playwrightMCP := map[string]any{
		"allowed": GetPlaywrightTools(),
	}
	if playwrightConfig.Version != "" {
		playwrightMCP["version"] = playwrightConfig.Version
	}
	if len(playwrightConfig.Args) > 0 {
		playwrightMCP["args"] = playwrightConfig.Args
	}
	result.raw["playwright"] = playwrightMCP
}

// expandNeutralToolsToCodexTools converts neutral tools to Codex-specific tools format
// This ensures that playwright tools get the same allowlist as the copilot agent
// Updated to use ToolsConfig instead of map[string]any
func (e *CodexEngine) expandNeutralToolsToCodexTools(toolsConfig *ToolsConfig) *ToolsConfig {
	if toolsConfig == nil {
		return &ToolsConfig{
			Custom: make(map[string]MCPServerConfig),
			raw:    make(map[string]any),
		}
	}

	// Create a copy of the tools config
	result := &ToolsConfig{
		GitHub:           toolsConfig.GitHub,
		Bash:             toolsConfig.Bash,
		WebFetch:         toolsConfig.WebFetch,
		WebSearch:        toolsConfig.WebSearch,
		Edit:             toolsConfig.Edit,
		Playwright:       toolsConfig.Playwright,
		AgenticWorkflows: toolsConfig.AgenticWorkflows,
		CacheMemory:      toolsConfig.CacheMemory,
		Timeout:          toolsConfig.Timeout,
		StartupTimeout:   toolsConfig.StartupTimeout,
		Custom:           make(map[string]MCPServerConfig),
		raw:              make(map[string]any),
	}

	// Copy custom tools and raw map
	maps.Copy(result.Custom, toolsConfig.Custom)
	maps.Copy(result.raw, toolsConfig.raw)

	applyPlaywrightToolToCodexConfig(toolsConfig, result)
	return result
}

// expandNeutralToolsToCodexToolsFromMap is a backward compatibility wrapper
// that accepts map[string]any instead of *ToolsConfig
func (e *CodexEngine) expandNeutralToolsToCodexToolsFromMap(tools map[string]any) map[string]any {
	toolsConfig, _ := ParseToolsConfig(tools)
	result := e.expandNeutralToolsToCodexTools(toolsConfig)
	return result.ToMap()
}

func (e *CodexEngine) getShellEnvironmentPolicyVars(tools map[string]any, mcpTools []string) []string {
	// Collect all environment variables needed by MCP servers
	envVarNames := make(map[string]struct{})

	// Always include core environment variables
	envVarNames["PATH"] = struct{}{}
	envVarNames["HOME"] = struct{}{}

	// Add CODEX_API_KEY for authentication
	envVarNames["CODEX_API_KEY"] = struct{}{}
	envVarNames["OPENAI_API_KEY"] = struct{}{} // Fallback for CODEX_API_KEY

	// Check each MCP tool for required environment variables
	for _, toolName := range mcpTools {
		switch toolName {
		case "github":
			// GitHub MCP server needs GITHUB_PERSONAL_ACCESS_TOKEN
			envVarNames["GITHUB_PERSONAL_ACCESS_TOKEN"] = struct{}{}
		case "agentic-workflows":
			// Agentic workflows MCP server needs GITHUB_TOKEN
			envVarNames["GITHUB_TOKEN"] = struct{}{}
		case "safe-outputs":
			// Safe outputs MCP server needs several environment variables
			envVarNames["GH_AW_SAFE_OUTPUTS"] = struct{}{}
			envVarNames["GH_AW_ASSETS_BRANCH"] = struct{}{}
			envVarNames["GH_AW_ASSETS_MAX_SIZE_KB"] = struct{}{}
			envVarNames["GH_AW_ASSETS_ALLOWED_EXTS"] = struct{}{}
			envVarNames["GITHUB_REPOSITORY"] = struct{}{}
			envVarNames["GITHUB_SERVER_URL"] = struct{}{}
		default:
			// For custom MCP tools, check if they have env configuration
			if toolValue, ok := tools[toolName]; ok {
				if toolConfig, ok := toolValue.(map[string]any); ok {
					// Extract environment variable names from env configuration
					if env, hasEnv := toolConfig["env"].(map[string]any); hasEnv {
						for envKey := range env {
							envVarNames[envKey] = struct{}{}
						}
					}
				}
			}
		}
	}

	var sortedEnvVars []string
	for envVar := range envVarNames {
		sortedEnvVars = append(sortedEnvVars, envVar)
	}
	sort.Strings(sortedEnvVars)

	// Codex expects regex patterns for shell_environment_policy.include_only, not literal names.
	// Anchor each variable name to avoid accidental substring matches (for example "PATH" matching "PATH_SUFFIX").
	var includeOnlyPatterns []string
	for _, envVar := range sortedEnvVars {
		includeOnlyPatterns = append(includeOnlyPatterns, "^"+regexp.QuoteMeta(envVar)+"$")
	}
	return includeOnlyPatterns
}

// renderShellEnvironmentPolicy generates the [shell_environment_policy] section for config.toml
// This controls which environment variables are passed through to MCP servers for security
func (e *CodexEngine) renderShellEnvironmentPolicy(yaml *strings.Builder, tools map[string]any, mcpTools []string) {
	sortedEnvVars := e.getShellEnvironmentPolicyVars(tools, mcpTools)

	// Render [shell_environment_policy] section
	yaml.WriteString("          \n")
	yaml.WriteString("          [shell_environment_policy]\n")
	yaml.WriteString("          inherit = \"core\"\n")
	yaml.WriteString("          include_only = [")
	for i, envVar := range sortedEnvVars {
		if i > 0 {
			yaml.WriteString(", ")
		}
		yaml.WriteString("\"" + envVar + "\"")
	}
	yaml.WriteString("]\n")
}

func (e *CodexEngine) renderShellEnvironmentPolicyToml(yaml *strings.Builder, tools map[string]any, mcpTools []string, indent string) {
	sortedEnvVars := e.getShellEnvironmentPolicyVars(tools, mcpTools)

	yaml.WriteString(indent + "[shell_environment_policy]\n")
	yaml.WriteString(indent + "inherit = \"core\"\n")
	yaml.WriteString(indent + "include_only = [")
	for i, envVar := range sortedEnvVars {
		if i > 0 {
			yaml.WriteString(", ")
		}
		yaml.WriteString("\"" + envVar + "\"")
	}
	yaml.WriteString("]\n")
}

// RenderMCPConfig is implemented in codex_mcp.go

// renderCodexMCPConfig is implemented in codex_mcp.go

// ParseLogMetrics is implemented in codex_logs.go

// parseCodexToolCallsWithSequence is implemented in codex_logs.go

// updateMostRecentToolWithDuration is implemented in codex_logs.go

// extractCodexTokenUsage is implemented in codex_logs.go

// GetLogParserScriptId is implemented in codex_logs.go
