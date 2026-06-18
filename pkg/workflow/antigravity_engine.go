package workflow

import (
	"fmt"
	"maps"

	"github.com/github/gh-aw/pkg/constants"
	"github.com/github/gh-aw/pkg/logger"
	"github.com/github/gh-aw/pkg/workflow/compilerenv"
)

var antigravityLog = logger.New("workflow:antigravity_engine")

// AntigravityEngine represents the Google Antigravity CLI agentic engine
type AntigravityEngine struct {
	BaseEngine
}

func NewAntigravityEngine() *AntigravityEngine {
	return &AntigravityEngine{
		BaseEngine: BaseEngine{
			id:           "antigravity",
			displayName:  "Antigravity CLI",
			description:  "Antigravity CLI with headless mode and LLM gateway support",
			experimental: true,
			capabilities: EngineCapabilities{
				ToolsAllowlist:   true,
				MaxTurns:         true,
				MaxContinuations: false, // Antigravity CLI does not support --max-autopilot-continues-style continuation mode
				WebSearch:        false,
				NativeAgentFile:  false, // Antigravity does not support agent file natively; the compiler prepends the agent file content to prompt.txt
			},
			dedicatedLLMGatewayPort: constants.AntigravityLLMGatewayPort,
		},
	}
}

// GetModelEnvVarName returns the native environment variable name that the Antigravity CLI uses
// for model selection. Setting ANTIGRAVITY_MODEL is equivalent to passing --model to the CLI.
func (e *AntigravityEngine) GetModelEnvVarName() string {
	return constants.AntigravityCLIModelEnvVar
}

// GetRequiredSecretNames returns the list of secrets required by the Antigravity engine
// This includes ANTIGRAVITY_API_KEY and optionally MCP_GATEWAY_API_KEY, GITHUB_MCP_SERVER_TOKEN,
// HTTP MCP header secrets, and mcp-scripts secrets
func (e *AntigravityEngine) GetRequiredSecretNames(workflowData *WorkflowData) []string {
	antigravityLog.Print("Collecting required secrets for Antigravity engine")
	secrets := []string{"ANTIGRAVITY_API_KEY"}

	// Add common MCP secrets (MCP_GATEWAY_API_KEY if MCP servers present, mcp-scripts secrets)
	secrets = append(secrets, collectCommonMCPSecrets(workflowData)...)

	// Add GitHub token for GitHub MCP server if present
	if hasGitHubTool(workflowData.ParsedTools) {
		antigravityLog.Print("Adding GITHUB_MCP_SERVER_TOKEN secret")
		secrets = append(secrets, "GITHUB_MCP_SERVER_TOKEN")
	}

	// Add HTTP MCP header secret names
	headerSecrets := collectHTTPMCPHeaderSecrets(workflowData.Tools)
	for varName := range headerSecrets {
		secrets = append(secrets, varName)
	}
	if len(headerSecrets) > 0 {
		antigravityLog.Printf("Added %d HTTP MCP header secrets", len(headerSecrets))
	}

	return secrets
}

// GetSecretValidationStep returns the secret validation step for the Antigravity engine.
// Returns an empty step if custom command is specified.
func (e *AntigravityEngine) GetSecretValidationStep(workflowData *WorkflowData) GitHubActionStep {
	return BuildDefaultSecretValidationStep(
		workflowData,
		[]string{"ANTIGRAVITY_API_KEY"},
		"Antigravity CLI",
		"https://antigravity.google/docs/cli-overview",
	)
}

func (e *AntigravityEngine) GetInstallationSteps(workflowData *WorkflowData) []GitHubActionStep {
	antigravityLog.Printf("Generating installation steps for Antigravity engine: workflow=%s", workflowData.Name)

	// Skip installation if custom command is specified
	if workflowData.EngineConfig != nil && workflowData.EngineConfig.Command != "" {
		antigravityLog.Printf("Skipping installation steps: custom command specified (%s)", workflowData.EngineConfig.Command)
		return []GitHubActionStep{}
	}

	version := string(constants.DefaultAntigravityVersion)
	if workflowData.EngineConfig != nil && workflowData.EngineConfig.Version != "" {
		version = workflowData.EngineConfig.Version
	}
	installSteps := GenerateAntigravityInstallerSteps(version, "Install Antigravity CLI")
	return BuildNpmEngineInstallStepsWithAWF(installSteps, workflowData)
}

// GetDeclaredOutputFiles returns the output files that Antigravity may produce.
// Antigravity CLI writes structured error reports to /tmp/antigravity-client-error-*.json
// with a timestamp in the filename (e.g. antigravity-client-error-Turn.run-sendMessageStream-2026-02-21T20-45-59-824Z.json).
// These files provide detailed diagnostics when the Antigravity API call fails.
// GetPreBundleSteps moves these files into /tmp/gh-aw/ so all artifact paths share a common
// ancestor under /tmp/gh-aw/ and the actions/upload-artifact LCA calculation stays correct.
func (e *AntigravityEngine) GetDeclaredOutputFiles() []string {
	return []string{
		constants.TmpAntigravityClientErrorGlob,
	}
}

// GetAgentManifestFiles returns Antigravity-specific instruction files that should be
// treated as security-sensitive manifests.  A fork PR that modifies these files
// can redirect the agent's behaviour or expand which files it treats as instructions.
// ANTIGRAVITY.md is the primary per-project context file; AGENTS.md is the cross-engine
// convention that Antigravity CLI also reads.
func (e *AntigravityEngine) GetAgentManifestFiles() []string {
	return []string{"ANTIGRAVITY.md", "AGENTS.md"}
}

// GetAgentManifestPathPrefixes returns Antigravity-specific config directory prefixes.
// The .antigravity/ directory contains settings.json and other configuration that could
// expand which files are treated as instructions or alter agent behaviour.
// Protecting this directory prevents fork PRs from injecting malicious configuration.
func (e *AntigravityEngine) GetAgentManifestPathPrefixes() []string {
	return []string{".antigravity/"}
}

// GetPreBundleSteps returns a step that moves Antigravity CLI error reports from /tmp/ into
// /tmp/gh-aw/ before the unified artifact upload. This keeps all artifact paths under
// /tmp/gh-aw/ so that actions/upload-artifact computes the correct least-common-ancestor
// path and downstream jobs find files at the expected locations.
func (e *AntigravityEngine) GetPreBundleSteps(workflowData *WorkflowData) []GitHubActionStep {
	return []GitHubActionStep{
		{
			"      - name: Move Antigravity error files to artifact directory",
			"        if: always()",
			"        run: mv /tmp/antigravity-client-error-*.json /tmp/gh-aw/ 2>/dev/null || true",
		},
	}
}

// buildAntigravityAgyCommand builds the core agy CLI command string and returns it along
// with a flag indicating whether the model was explicitly configured.
func buildAntigravityAgyCommand(workflowData *WorkflowData) (string, bool) {
	modelConfigured := workflowData.EngineConfig != nil && workflowData.EngineConfig.Model != ""

	// Auto-approve all tool executions so non-interactive CI runs don't block on permission prompts.
	var agyArgs []string
	agyArgs = append(agyArgs, "--dangerously-skip-permissions")

	commandName := "agy"
	if workflowData.EngineConfig != nil && workflowData.EngineConfig.Command != "" {
		commandName = workflowData.EngineConfig.Command
	}

	// Append the prompt arg raw (not through shellJoinArgs) to preserve shell expansion
	agyCommand := fmt.Sprintf(`%s %s --prompt "$(cat /tmp/gh-aw/aw-prompts/prompt.txt)"`, commandName, shellJoinArgs(agyArgs))
	return agyCommand, modelConfigured
}

// buildAntigravityCommandString builds the full shell command string to run the Antigravity CLI,
// wrapping it with AWF if the firewall is enabled.
func buildAntigravityCommandString(workflowData *WorkflowData, agyCommand, logFile string, firewallEnabled bool) string {
	if !firewallEnabled {
		return fmt.Sprintf(`set -o pipefail
printf '%%s' "$(date +%%s%%3N)" > %s
touch %s
(umask 177 && touch %s)
%s 2>&1 | tee -a %s`, AgentCLIStartMsPath, AgentStepSummaryPath, logFile, agyCommand, logFile)
	}

	// Get allowed domains: prefer the pre-warmed cache on WorkflowData to avoid
	// re-running the expensive map+sort operation.
	var allowedDomains string
	if workflowData.CachedAllowedDomainsComputed {
		allowedDomains = workflowData.CachedAllowedDomainsStr
	} else {
		allowedDomains = GetAllowedDomainsForEngine(constants.AntigravityEngine,
			workflowData.NetworkPermissions,
			workflowData.Tools,
			workflowData.Runtimes,
		)
	}
	if workflowData.EngineConfig != nil && workflowData.EngineConfig.APITarget != "" {
		allowedDomains = mergeAPITargetDomains(allowedDomains, workflowData.EngineConfig.APITarget)
	}

	npmPathSetup := GetNpmBinPathSetup()
	agyCommandWithPath := fmt.Sprintf("%s && %s", npmPathSetup, agyCommand)
	if mcpCLIPath := GetMCPCLIPathSetup(workflowData); mcpCLIPath != "" {
		agyCommandWithPath = fmt.Sprintf("%s && %s", mcpCLIPath, agyCommandWithPath)
	}

	return BuildAWFCommand(AWFCommandConfig{
		EngineName:         "antigravity",
		EngineCommand:      agyCommandWithPath,
		LogFile:            logFile,
		WorkflowData:       workflowData,
		UsesTTY:            false,
		AllowedDomains:     allowedDomains,
		PathSetup:          "touch " + AgentStepSummaryPath,
		ExcludeEnvVarNames: ComputeAWFExcludeEnvVarNames(workflowData, []string{"ANTIGRAVITY_API_KEY", "GEMINI_API_KEY"}),
	})
}

// buildAntigravityBaseEnv builds the base environment variable map for an Antigravity execution step.
// It includes standard keys, phase/version markers, optional MCP config, and firewall-specific vars.
func buildAntigravityBaseEnv(workflowData *WorkflowData, firewallEnabled bool) map[string]string {
	env := map[string]string{
		"ANTIGRAVITY_API_KEY": "${{ secrets.ANTIGRAVITY_API_KEY }}",
		"GH_AW_PROMPT":        constants.AwPromptsFile,
		// Tag the step as a GitHub AW agentic execution for discoverability by agents
		"GITHUB_AW":        "true",
		"GITHUB_WORKSPACE": "${{ github.workspace }}",
		"RUNNER_TEMP":      "${{ runner.temp }}",
		// Override GITHUB_STEP_SUMMARY with a path that exists inside the sandbox.
		"GITHUB_STEP_SUMMARY": AgentStepSummaryPath,
		// Enable verbose debug logging from Antigravity CLI for better diagnostics.
		"DEBUG": "antigravity-cli:*",
		// Trust the workspace to prevent Antigravity CLI v1.x from overriding --yolo.
		"ANTIGRAVITY_CLI_TRUST_WORKSPACE": "true",
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
	if HasMCPServers(workflowData) {
		env["GH_AW_MCP_CONFIG"] = "${{ github.workspace }}/.antigravity/settings.json"
	}
	if firewallEnabled {
		env["ANTIGRAVITY_API_BASE_URL"] = fmt.Sprintf("http://host.docker.internal:%d", constants.AntigravityLLMGatewayPort)
		maps.Copy(env, getGitIdentityEnvVars())
	}
	// Add safe outputs env
	applySafeOutputEnvToMap(env, workflowData)
	// Propagate W3C trace context so engine spans nest under the gh-aw.agent.setup span.
	applyTraceContextEnvToMap(env)
	return env
}

// applyAntigravityModelAndCustomEnv sets the max-turns, model, custom engine/agent env vars, and
// the GEMINI_API_KEY sync on the given env map for an Antigravity execution step.
func applyAntigravityModelAndCustomEnv(env map[string]string, workflowData *WorkflowData, modelConfigured bool) {
	if workflowData.EngineConfig != nil && workflowData.EngineConfig.MaxTurns != "" {
		env["GH_AW_MAX_TURNS"] = workflowData.EngineConfig.MaxTurns
	} else {
		env["GH_AW_MAX_TURNS"] = compilerenv.BuildDefaultMaxTurnsExpression()
	}

	// Set the model environment variable only when explicitly configured.
	if modelConfigured {
		antigravityLog.Printf("Setting %s env var for model: %s", constants.AntigravityCLIModelEnvVar, workflowData.EngineConfig.Model)
		env[constants.AntigravityCLIModelEnvVar] = workflowData.EngineConfig.Model
	}

	// Add custom environment variables from engine config.
	if workflowData.EngineConfig != nil && len(workflowData.EngineConfig.Env) > 0 {
		maps.Copy(env, workflowData.EngineConfig.Env)
	}

	// Add custom environment variables from agent config
	agentConfig := getAgentConfig(workflowData)
	if agentConfig != nil && len(agentConfig.Env) > 0 {
		maps.Copy(env, agentConfig.Env)
		antigravityLog.Printf("Added %d custom env vars from agent config", len(agentConfig.Env))
	}

	// Keep GEMINI_API_KEY aligned with the effective ANTIGRAVITY_API_KEY by default.
	if _, hasGeminiKey := env["GEMINI_API_KEY"]; !hasGeminiKey {
		env["GEMINI_API_KEY"] = env["ANTIGRAVITY_API_KEY"]
	}
}

// GetExecutionSteps returns the GitHub Actions steps for executing Antigravity
func (e *AntigravityEngine) GetExecutionSteps(workflowData *WorkflowData, logFile string) []GitHubActionStep {
	antigravityLog.Printf("Generating execution steps for Antigravity engine: workflow=%s, firewall=%v", workflowData.Name, isFirewallEnabled(workflowData))

	var steps []GitHubActionStep

	// Write .antigravity/settings.json with context.includeDirectories and tools.core.
	settingsStep := e.generateAntigravitySettingsStep(workflowData)
	steps = append(steps, settingsStep)

	agyCommand, modelConfigured := buildAntigravityAgyCommand(workflowData)
	firewallEnabled := isFirewallEnabled(workflowData)
	command := buildAntigravityCommandString(workflowData, agyCommand, logFile, firewallEnabled)

	env := buildAntigravityBaseEnv(workflowData, firewallEnabled)
	applyAntigravityModelAndCustomEnv(env, workflowData, modelConfigured)

	// Generate the execution step
	stepLines := []string{
		"      - name: Execute Antigravity CLI",
		"        id: agentic_execution",
	}

	// Filter environment variables for security
	allowedSecrets := append([]string{"GEMINI_API_KEY"}, e.GetRequiredSecretNames(workflowData)...)
	filteredEnv := FilterEnvForSecrets(env, allowedSecrets)

	// Inject GH_TOKEN for CLI proxy
	addCliProxyGHTokenToEnv(filteredEnv, workflowData)

	// Format step with command and env
	stepLines = FormatStepWithCommandAndEnv(stepLines, command, filteredEnv)

	steps = append(steps, GitHubActionStep(stepLines))
	return steps
}
