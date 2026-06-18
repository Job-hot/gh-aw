// This file provides helper functions for AWF (Agentic Workflow Firewall) integration.
//
// AWF is the network firewall/sandbox used by gh-aw to control network egress for
// AI agent execution. This file consolidates common AWF logic that was previously
// duplicated across multiple engine implementations (Copilot, Claude, Codex).
//
// # Key Functions
//
// AWF Command Building:
//   - BuildAWFCommand() - Builds complete AWF command with all arguments
//   - BuildAWFArgs() - Constructs common AWF arguments from configuration
//   - GetAWFCommandPrefix() - Determines AWF command (custom vs standard)
//   - WrapCommandInShell() - Wraps engine command in shell for AWF execution
//
// AWF Configuration:
//   - GetAWFDomains() - Combines allowed/blocked domains from various sources
//   - GetSSLBumpArgs() - Returns SSL bump configuration arguments
//   - GetAWFImageTag() - Returns pinned AWF image tag
//
// These functions extract shared AWF patterns from engine implementations,
// providing a consistent and maintainable approach to AWF integration.

package workflow

import (
	"encoding/json"
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/github/gh-aw/pkg/constants"
	"github.com/github/gh-aw/pkg/logger"
	"github.com/github/gh-aw/pkg/workflow/compilerenv"
)

var awfHelpersLog = logger.New("workflow:awf_helpers")

const (
	awfArcDindPrefixArgsVarName = "GH_AW_DOCKER_HOST_PATH_PREFIX_ARGS"
	awfDockerHostVarName        = "GH_AW_DOCKER_HOST"
	awfToolCacheMountVarName    = "GH_AW_TOOL_CACHE_MOUNT"
	awfMaxAICreditsVarName      = "GH_AW_MAX_AI_CREDITS"
	awfConfigRuntimePathExpr    = "${RUNNER_TEMP}/gh-aw/awf-config.json"
	awfModelsJSONPathExpr       = "/tmp/gh-aw/models.json"
	// Bash regex used in [[ ... =~ ... ]] to detect TCP Docker hosts (ARC/DinD).
	// Any tcp:// DOCKER_HOST indicates the Docker daemon runs on a separate filesystem,
	// requiring --docker-host-path-prefix so AWF bind-mounts resolve against the daemon.
	// This covers localhost, pod IPs, K8s service names (e.g., tcp://dind:2375), and
	// any other TCP Docker daemon configuration.
	awfArcDindDockerHostRegex    = `^tcp://`
	awfArcDindHostPathPrefixFlag = "--docker-host-path-prefix /tmp/gh-aw"

	// awfArcDindChrootBinariesSourcePath is the runner-side directory that AWF overlays
	// at /usr/local/bin inside chroot mode for ARC/DinD split-filesystem runners.
	// This is the gh-aw staging directory that holds pre-downloaded binaries (e.g., copilot).
	awfArcDindChrootBinariesSourcePath = "/tmp/gh-aw"

	// awfArcDindChrootIdentityHome is the home directory path exported inside chroot mode
	// for ARC/DinD runners. A dedicated directory under /tmp/gh-aw is used so that the
	// runner user has a consistent home that exists on the daemon-visible filesystem.
	awfArcDindChrootIdentityHome = "/tmp/gh-aw/home"
)

// AWFCommandConfig contains configuration for building AWF commands.
// This struct centralizes all the parameters needed to construct an AWF-wrapped command.
type AWFCommandConfig struct {
	// EngineName is the engine ID (e.g., "copilot", "claude", "codex")
	EngineName string

	// EngineCommand is the command to execute inside AWF
	EngineCommand string

	// LogFile is the path to the log file
	LogFile string

	// WorkflowData contains all workflow configuration
	WorkflowData *WorkflowData

	// UsesTTY indicates if the engine requires a TTY (e.g., Claude)
	UsesTTY bool

	// AllowedDomains is the comma-separated list of allowed domains
	AllowedDomains string

	// PathSetup is optional shell commands to run before the engine command
	// (e.g., npm PATH setup)
	PathSetup string

	// ExcludeEnvVarNames is the list of environment variable names to exclude from
	// the agent container's visible environment via --exclude-env. These are the env
	// var keys whose step-env values contain secret references (${{ secrets.* }}).
	// Computed from the engine's GetRequiredSecretNames() so that every secret-bearing
	// variable is excluded — the agent can never read raw token values via `env`/`printenv`.
	// Requires AWF v0.25.3+ for --exclude-env support.
	ExcludeEnvVarNames []string

	// ResolveMaxAICreditsFromEnv switches maxAiCredits runtime resolution from an inline
	// GitHub Actions expression in run: to the GH_AW_MAX_AI_CREDITS step env variable.
	// When true and max-ai-credits is unset, BuildAWFCommand emits:
	//   GH_AW_MAX_AI_CREDITS="${GH_AW_MAX_AI_CREDITS:-<default>}"
	// instead of embedding ${{ vars.* }} directly in run:.
	ResolveMaxAICreditsFromEnv bool
}

func shouldUseWorkflowCallNetworkAllowedInput(data *WorkflowData) bool {
	return data != nil &&
		data.NetworkPermissions != nil &&
		data.NetworkPermissions.AllowedInput &&
		hasWorkflowCallTrigger(data.On)
}

func buildModelsJSONPathExportScript() string {
	return fmt.Sprintf(`export GH_AW_MODELS_JSON_PATH="%s"`, awfModelsJSONPathExpr)
}

// applyDefaultMaxAICreditsEnvToMap adds the runtime max-ai-credits GitHub Actions expression
// to env when no compile-time max-ai-credits is configured.
//
// This keeps the organization/repository variable override behavior while allowing AWF run:
// scripts to read GH_AW_MAX_AI_CREDITS from step env instead of embedding ${{ vars.* }}
// directly in run blocks.
func applyDefaultMaxAICreditsEnvToMap(env map[string]string, workflowData *WorkflowData) {
	if env == nil {
		return
	}
	if workflowData != nil && workflowData.EngineConfig != nil && workflowData.EngineConfig.MaxAICredits != 0 {
		return
	}
	if workflowData != nil && workflowData.IsDetectionRun {
		env[awfMaxAICreditsVarName] = compilerenv.BuildDefaultDetectionMaxAICreditsExpression(strconv.FormatInt(constants.DefaultDetectionMaxAICredits, 10))
		return
	}
	env[awfMaxAICreditsVarName] = compilerenv.BuildDefaultMaxAICreditsExpression(strconv.FormatInt(constants.DefaultMaxAICredits, 10))
}

// injectMaxAICreditsExpression inserts "maxAiCredits":expr into the apiProxy
// JSON object of awfConfigJSON directly after the "maxRuns" field value.
//
// expr is a shell variable reference such as "${GH_AW_MAX_AI_CREDITS}". The
// caller emits a local export line before the printf command that assigns the
// GitHub Actions runtime expression to that variable, so the ${{ }} expression
// lives on one clean, dedicated line rather than being embedded inside the JSON.
//
// shellEscapeArgWithVarPreserved is then used to double-quote the JSON arg while
// preserving the ${varName} reference for bash expansion and escaping bare $ signs
// (e.g. "$schema" → "\$schema").
func injectMaxAICreditsExpression(awfConfigJSON string, expr string) string {
	const maxRunsKey = `"maxRuns":`
	idx := strings.Index(awfConfigJSON, maxRunsKey)
	if idx == -1 {
		awfHelpersLog.Print("Warning: could not find maxRuns in AWF config JSON; maxAiCredits expression not injected")
		return awfConfigJSON
	}
	// Scan past the integer value of maxRuns.
	valueEnd := idx + len(maxRunsKey)
	for valueEnd < len(awfConfigJSON) && awfConfigJSON[valueEnd] >= '0' && awfConfigJSON[valueEnd] <= '9' {
		valueEnd++
	}
	return awfConfigJSON[:valueEnd] + `,"maxAiCredits":` + expr + awfConfigJSON[valueEnd:]
}

func buildWorkflowCallNetworkAllowedUpdateScript() (string, error) {
	ecosystemMap := make(map[string][]string, safeAllocationCapacity(len(ecosystemDomains), len(compoundEcosystems)))
	for ecosystem := range ecosystemDomains {
		ecosystemMap[ecosystem] = getEcosystemDomains(ecosystem)
	}
	for ecosystem := range compoundEcosystems {
		ecosystemMap[ecosystem] = getEcosystemDomains(ecosystem)
	}

	ecosystemJSON, err := json.Marshal(ecosystemMap)
	if err != nil {
		return "", fmt.Errorf("marshal network allowed ecosystem map: %w", err)
	}

	return fmt.Sprintf(`python3 - <<'PY'
import json
import os
from pathlib import Path

runner_temp = os.environ.get("RUNNER_TEMP")
if not runner_temp:
    raise SystemExit("RUNNER_TEMP is not set")

config_path = Path(runner_temp) / "gh-aw" / "awf-config.json"
try:
    config = json.loads(config_path.read_text())
except FileNotFoundError as exc:
    raise SystemExit(f"Missing AWF config file at {config_path}") from exc
except json.JSONDecodeError as exc:
    raise SystemExit(f"Invalid AWF config JSON at {config_path}: {exc}") from exc
except OSError as exc:
    raise SystemExit(f"Failed to read AWF config file at {config_path}: {exc}") from exc

network_allowed = os.environ.get(%q, "")
tokens = [token.strip() for token in network_allowed.split(",") if token.strip()]

if tokens:
    ecosystem_map = json.loads(r'''%s''')
    allow_domains = config.setdefault("network", {}).setdefault("allowDomains", [])
    seen = set(allow_domains)
    for token in tokens:
        for domain in ecosystem_map.get(token, [token]):
            if domain not in seen:
                allow_domains.append(domain)
                seen.add(domain)

try:
    config_path.write_text(json.dumps(config, separators=(",", ":"), ensure_ascii=False) + "\n")
except OSError as exc:
    raise SystemExit(f"Failed to write AWF config file at {config_path}: {exc}") from exc
PY`, string(WorkflowCallNetworkAllowedEnvVar), string(ecosystemJSON)), nil
}

// buildAWFArcDindProbes returns shell probe/reference variable strings for ARC DinD topology.
// When AWF supports --docker-host-path-prefix, it also returns the prefix probe and ref.
func buildAWFArcDindProbes(firewallConfig *FirewallConfig) (dockerHostProbe, dockerHostRef, prefixProbe, prefixArgsRef string) {
dockerHostProbe = fmt.Sprintf(`%s=""
if [[ "${DOCKER_HOST:-}" =~ %s ]]; then
  %s="${DOCKER_HOST}"
fi`,
awfDockerHostVarName,
awfArcDindDockerHostRegex,
awfDockerHostVarName,
)
dockerHostRef = fmt.Sprintf("${%s:+--docker-host \"$%s\"}", awfDockerHostVarName, awfDockerHostVarName)
if awfSupportsDockerHostPathPrefix(firewallConfig) {
chrootPatchBody := ""
if awfSupportsChrootConfig(firewallConfig) {
chrootPatchBody = "\n" + buildArcDindChrootConfigPatchBody()
}
prefixProbe = fmt.Sprintf(`%s=""
if [[ "${DOCKER_HOST:-}" =~ %s ]]; then
  %s="%s"%s
fi`,
awfArcDindPrefixArgsVarName,
awfArcDindDockerHostRegex,
awfArcDindPrefixArgsVarName,
awfArcDindHostPathPrefixFlag,
chrootPatchBody)
prefixArgsRef = fmt.Sprintf("${%s}", awfArcDindPrefixArgsVarName)
}
return dockerHostProbe, dockerHostRef, prefixProbe, prefixArgsRef
}

// buildAWFToolCacheMountProbe returns a shell probe and reference for conditionally mounting
// the runner tool cache into the AWF container.
func buildAWFToolCacheMountProbe() (probe, ref string) {
probe = fmt.Sprintf(`%s=""
GH_AW_TOOL_CACHE="${RUNNER_TOOL_CACHE:-/opt/hostedtoolcache}"
if [ -d "$GH_AW_TOOL_CACHE" ]; then
  if [[ "$GH_AW_TOOL_CACHE" != /opt/* ]]; then
    %s="$GH_AW_TOOL_CACHE:$GH_AW_TOOL_CACHE:ro"
  fi
elif [ -d "/home/runner/work/_tool" ]; then
  %s="/home/runner/work/_tool:/home/runner/work/_tool:ro"
fi`,
awfToolCacheMountVarName,
awfToolCacheMountVarName,
awfToolCacheMountVarName,
)
ref = fmt.Sprintf("${%s:+--mount \"$%s\"}", awfToolCacheMountVarName, awfToolCacheMountVarName)
return probe, ref
}

// buildAWFMaxAICreditsLines injects a runtime max-AI-credits shell variable into the AWF
// config JSON when no compile-time budget value is set. Returns the updated JSON and the
// shell export line (empty when MaxAICredits is already set at compile time).
func buildAWFMaxAICreditsLines(config AWFCommandConfig, awfConfigJSON string) (updatedJSON, exportLine string) {
if config.WorkflowData != nil && config.WorkflowData.EngineConfig != nil && config.WorkflowData.EngineConfig.MaxAICredits != 0 {
return awfConfigJSON, ""
}
defaultMaxAICredits := strconv.FormatInt(constants.DefaultMaxAICredits, 10)
if config.WorkflowData != nil && config.WorkflowData.IsDetectionRun {
defaultMaxAICredits = strconv.FormatInt(constants.DefaultDetectionMaxAICredits, 10)
}
awfConfigJSON = injectMaxAICreditsExpression(awfConfigJSON, fmt.Sprintf("${%s}", awfMaxAICreditsVarName))
if config.ResolveMaxAICreditsFromEnv {
exportLine = fmt.Sprintf(`%s="${%s:-%s}"`, awfMaxAICreditsVarName, awfMaxAICreditsVarName, defaultMaxAICredits)
} else {
expr := compilerenv.BuildDefaultMaxAICreditsExpression(defaultMaxAICredits)
if config.WorkflowData != nil && config.WorkflowData.IsDetectionRun {
expr = compilerenv.BuildDefaultDetectionMaxAICreditsExpression(defaultMaxAICredits)
}
exportLine = fmt.Sprintf(`%s="%s"`, awfMaxAICreditsVarName, expr)
}
awfHelpersLog.Printf("Injected maxAiCredits local var reference into AWF config JSON")
return awfConfigJSON, exportLine
}

// buildAWFConfigFileSetup generates shell commands to write the AWF config JSON to disk.
// Returns the shell setup script and whether the config file will be written.
func buildAWFConfigFileSetup(config AWFCommandConfig) (setup string, written bool) {
awfConfigJSON, err := BuildAWFConfigJSON(config)
if err != nil {
awfHelpersLog.Printf("Warning: failed to build AWF config JSON: %v", err)
return "", false
}
awfConfigJSON, maxAICreditsExportLine := buildAWFMaxAICreditsLines(config, awfConfigJSON)
var printfArg string
if maxAICreditsExportLine != "" {
printfArg = shellEscapeArgWithVarPreserved(awfConfigJSON, awfMaxAICreditsVarName)
} else {
printfArg = shellEscapeArg(awfConfigJSON)
}
setup = fmt.Sprintf("printf '%%s\\n' %s > %q", printfArg, awfConfigRuntimePathExpr)
if maxAICreditsExportLine != "" {
setup = maxAICreditsExportLine + "\n" + setup
}
if shouldUseWorkflowCallNetworkAllowedInput(config.WorkflowData) {
if updateScript, updateErr := buildWorkflowCallNetworkAllowedUpdateScript(); updateErr == nil {
setup += "\n" + updateScript
} else {
awfHelpersLog.Printf("Warning: failed to build workflow_call network_allowed updater: %v", updateErr)
}
}
setup += fmt.Sprintf("\ncp %q %s", awfConfigRuntimePathExpr, constants.AWFConfigFilePath)
awfHelpersLog.Print("Using AWF config file (--config flag)")
return setup, true
}

// buildAWFExpandableArgs constructs the AWF expandable args string (container workdir, mounts,
// optional --config prefix, upload artifact mount, and service port expressions).
func buildAWFExpandableArgs(config AWFCommandConfig, configWasWritten bool) string {
ghAwDir := constants.GhAwRootDirShell
expandableArgs := fmt.Sprintf(
`--container-workdir "${GITHUB_WORKSPACE}" --mount "%s:%s:ro" --mount "%s:/host%s:ro"`,
ghAwDir, ghAwDir, ghAwDir, ghAwDir,
)
if configWasWritten {
expandableArgs = fmt.Sprintf("--config %q ", awfConfigRuntimePathExpr) + expandableArgs
}
if config.WorkflowData != nil && config.WorkflowData.SafeOutputs != nil && config.WorkflowData.SafeOutputs.UploadArtifact != nil {
stagingDir := SafeOutputsUploadArtifactsDir
expandableArgs += fmt.Sprintf(` --mount "%s:%s:rw"`, stagingDir, stagingDir)
awfHelpersLog.Print("Added read-write mount for upload_artifact staging directory")
}
if config.WorkflowData != nil && config.WorkflowData.ServicePortExpressions != "" {
expandableArgs += fmt.Sprintf(` --allow-host-service-ports "%s"`, config.WorkflowData.ServicePortExpressions)
awfHelpersLog.Printf("Added --allow-host-service-ports with %s", config.WorkflowData.ServicePortExpressions)
}
return expandableArgs
}

// assembleAWFShellCommand assembles the final AWF shell script from its component parts.
// dynamicRefs is the pre-combined string of tool-cache mount ref, ARC DinD docker-host ref,
// and ARC DinD prefix args ref (all shell-expansion variables that may be empty at runtime).
func assembleAWFShellCommand(preamble []string, awfCmd, expandableArgs, dynamicRefs, awfArgsStr, shellWrappedCmd, logFile string) string {
return fmt.Sprintf(`set -o pipefail
%s
# shellcheck disable=SC1003
%s %s %s %s \
  -- %s 2>&1 | tee -a %s`,
strings.Join(preamble, "\n"),
awfCmd, expandableArgs, dynamicRefs, awfArgsStr,
shellWrappedCmd, shellEscapeArg(logFile))
}

// BuildAWFCommand builds the complete AWF shell command string for a given engine configuration.
// It assembles ARC DinD probes, tool cache mounts, config file setup, and expandable args,
// then combines them into a single pipefail-guarded shell script.
//
// The returned command contains only args that are safe to pass through shellJoinArgs.
// Expandable-var args (--container-workdir "${GITHUB_WORKSPACE}" and --mount "${RUNNER_TEMP}/...")
// are appended raw so that shell variable expansion is not suppressed by single-quoting.
func BuildAWFCommand(config AWFCommandConfig) string {
awfHelpersLog.Printf("Building AWF command for engine: %s", config.EngineName)
awfCommand := GetAWFCommandPrefix(config.WorkflowData)
awfArgs := BuildAWFArgs(config)
firewallConfig := getFirewallConfig(config.WorkflowData)

arcDockerHostProbe, arcDockerHostRef, arcPrefixProbe, arcPrefixArgsRef := buildAWFArcDindProbes(firewallConfig)
toolCacheProbe, toolCacheRef := buildAWFToolCacheMountProbe()
configFileSetup, configWritten := buildAWFConfigFileSetup(config)
expandableArgs := buildAWFExpandableArgs(config, configWritten)
modelsJSONPathExport := buildModelsJSONPathExportScript()

writeAgentCLIStartMs := "printf '%s' \"$(date +%s%3N)\" > " + shellEscapeArg(AgentCLIStartMsPath)
preCreateLog := fmt.Sprintf("(umask 177 && touch %s)", shellEscapeArg(config.LogFile))

var preamble []string
preamble = append(preamble, writeAgentCLIStartMs)
if config.PathSetup != "" {
preamble = append(preamble, config.PathSetup)
}
preamble = append(preamble, preCreateLog)
if configFileSetup != "" {
preamble = append(preamble, configFileSetup)
}
preamble = append(preamble, modelsJSONPathExport, arcDockerHostProbe, arcPrefixProbe, toolCacheProbe)

awfHelpersLog.Print("Successfully built AWF command")
dynamicRefs := toolCacheRef + " " + arcDockerHostRef + " " + arcPrefixArgsRef
return assembleAWFShellCommand(
preamble, awfCommand, expandableArgs, dynamicRefs,
shellJoinArgs(awfArgs), WrapCommandInShell(config.EngineCommand), config.LogFile)
}

// buildAWFContainerEnvArgs returns AWF args for container environment: --tty, --env-all, --exclude-env.
func buildAWFContainerEnvArgs(config AWFCommandConfig, firewallConfig *FirewallConfig) []string {
var args []string
if config.UsesTTY {
args = append(args, "--tty")
}
args = append(args, "--env-all")
if awfSupportsExcludeEnv(firewallConfig) {
sortedExclude := make([]string, len(config.ExcludeEnvVarNames))
copy(sortedExclude, config.ExcludeEnvVarNames)
sort.Strings(sortedExclude)
for _, excludedVar := range sortedExclude {
args = append(args, "--exclude-env", excludedVar)
}
} else {
awfHelpersLog.Printf("Skipping --exclude-env: AWF version %q is older than minimum %s", getAWFImageTag(firewallConfig), constants.AWFExcludeEnvMinVersion)
}
return args
}

// buildAWFContainerMountAndLoggingArgs returns AWF args for custom mounts, log level,
// proxy logs directory, audit directory, and optional diagnostic logging.
func buildAWFContainerMountAndLoggingArgs(config AWFCommandConfig, firewallConfig *FirewallConfig, agentConfig *AgentSandboxConfig) []string {
var args []string
if agentConfig != nil && len(agentConfig.Mounts) > 0 {
sortedMounts := make([]string, len(agentConfig.Mounts))
copy(sortedMounts, agentConfig.Mounts)
sort.Strings(sortedMounts)
for _, mount := range sortedMounts {
args = append(args, "--mount", mount)
}
awfHelpersLog.Printf("Added %d custom mounts from agent config", len(sortedMounts))
}
awfLogLevel := string(constants.AWFDefaultLogLevel)
if firewallConfig != nil && firewallConfig.LogLevel != "" {
awfLogLevel = firewallConfig.LogLevel
}
args = append(args, "--log-level", awfLogLevel)
args = append(args, "--proxy-logs-dir", string(constants.AWFProxyLogsDir))
args = append(args, "--audit-dir", string(constants.AWFAuditDir))
if isFeatureEnabled(constants.AwfDiagnosticLogsFeatureFlag, config.WorkflowData) {
args = append(args, "--diagnostic-logs")
awfHelpersLog.Print("Added --diagnostic-logs because awf-diagnostic-logs feature flag is enabled")
}
return args
}

// buildAWFContainerNetworkArgs returns AWF args for host access, allow-host-ports, skip-pull, and CLI proxy.
func buildAWFContainerNetworkArgs(config AWFCommandConfig, firewallConfig *FirewallConfig) []string {
var args []string
args = append(args, "--enable-host-access")
awfHelpersLog.Print("Added --enable-host-access for API proxy and MCP gateway")
if awfSupportsAllowHostPorts(firewallConfig) {
mcpGatewayPort := int(DefaultMCPGatewayPort)
if config.WorkflowData != nil && config.WorkflowData.SandboxConfig != nil &&
config.WorkflowData.SandboxConfig.MCP != nil && config.WorkflowData.SandboxConfig.MCP.Port > 0 {
mcpGatewayPort = config.WorkflowData.SandboxConfig.MCP.Port
}
hostPorts := fmt.Sprintf("80,443,%d", mcpGatewayPort)
args = append(args, "--allow-host-ports", hostPorts)
awfHelpersLog.Printf("Added --allow-host-ports %s for MCP gateway access", hostPorts)
} else {
awfHelpersLog.Printf("Skipping --allow-host-ports: AWF version %q requires at least %s", getAWFImageTag(firewallConfig), constants.AWFAllowHostPortsMinVersion)
}
args = append(args, "--skip-pull")
awfHelpersLog.Print("Using --skip-pull since images are pre-downloaded")
if isGitHubCLIModeEnabled(config.WorkflowData) {
if awfSupportsCliProxy(firewallConfig) {
args = append(args, "--difc-proxy-host", "host.docker.internal:18443")
args = append(args, "--difc-proxy-ca-cert", constants.TmpDIFCProxyTLSCACert)
awfHelpersLog.Print("Added --difc-proxy-host and --difc-proxy-ca-cert for CLI proxy sidecar")
} else {
awfHelpersLog.Printf("Skipping CLI proxy flags: AWF version %q is older than minimum %s", getAWFImageTag(firewallConfig), constants.AWFCliProxyMinVersion)
}
}
return args
}

// buildAWFExtraArgs returns remaining AWF args: API base paths, SSL bump, custom firewall/agent args, and memory.
func buildAWFExtraArgs(config AWFCommandConfig, firewallConfig *FirewallConfig, agentConfig *AgentSandboxConfig) []string {
var args []string
if openaiBasePath := extractAPIBasePath(config.WorkflowData, "OPENAI_BASE_URL"); openaiBasePath != "" {
args = append(args, "--openai-api-base-path", openaiBasePath)
awfHelpersLog.Printf("Added --openai-api-base-path=%s", openaiBasePath)
}
if anthropicBasePath := extractAPIBasePath(config.WorkflowData, "ANTHROPIC_BASE_URL"); anthropicBasePath != "" {
args = append(args, "--anthropic-api-base-path", anthropicBasePath)
awfHelpersLog.Printf("Added --anthropic-api-base-path=%s", anthropicBasePath)
}
if geminiBasePath := extractAPIBasePath(config.WorkflowData, "GEMINI_API_BASE_URL"); geminiBasePath != "" {
args = append(args, "--gemini-api-base-path", geminiBasePath)
awfHelpersLog.Printf("Added --gemini-api-base-path=%s", geminiBasePath)
}
args = append(args, getSSLBumpArgs(firewallConfig)...)
if firewallConfig != nil && len(firewallConfig.Args) > 0 {
args = append(args, firewallConfig.Args...)
}
if agentConfig != nil && len(agentConfig.Args) > 0 {
args = append(args, agentConfig.Args...)
awfHelpersLog.Printf("Added %d custom args from agent config", len(agentConfig.Args))
}
if agentConfig != nil && agentConfig.Memory != "" {
args = append(args, "--memory-limit", agentConfig.Memory)
awfHelpersLog.Printf("Set AWF memory limit to %s", agentConfig.Memory)
}
return args
}

// BuildAWFArgs constructs common AWF arguments from configuration.
// This extracts the shared AWF argument building logic from engine implementations.
//
// The following flags are expressed in the generated JSON config file written by
// BuildAWFCommand and are therefore not emitted here:
//   - --allow-domains / --block-domains   → network.allowDomains / network.blockDomains
//   - --enable-api-proxy                  → apiProxy.enabled
//   - --image-tag                         → container.imageTag
//   - --openai-api-target                 → apiProxy.targets.openai.host
//   - --anthropic-api-target              → apiProxy.targets.anthropic.host
//   - --copilot-api-target                → apiProxy.targets.copilot.host
//   - --gemini-api-target                 → apiProxy.targets.gemini.host
//
// Parameters:
//   - config: AWF command configuration
//
// Returns:
//   - []string: List of AWF arguments (safe args only; expandable-var args like
//     --container-workdir and --mount are handled by BuildAWFCommand)
func BuildAWFArgs(config AWFCommandConfig) []string {
awfHelpersLog.Printf("Building AWF args for engine: %s", config.EngineName)
firewallConfig := getFirewallConfig(config.WorkflowData)
agentConfig := getAgentConfig(config.WorkflowData)
var awfArgs []string
awfArgs = append(awfArgs, buildAWFContainerEnvArgs(config, firewallConfig)...)
awfArgs = append(awfArgs, buildAWFContainerMountAndLoggingArgs(config, firewallConfig, agentConfig)...)
awfArgs = append(awfArgs, buildAWFContainerNetworkArgs(config, firewallConfig)...)
awfArgs = append(awfArgs, buildAWFExtraArgs(config, firewallConfig, agentConfig)...)
awfHelpersLog.Printf("Built %d AWF arguments", len(awfArgs))
return awfArgs
}

// GetAWFCommandPrefix determines the AWF command to use (custom or standard).
// This extracts the common pattern for determining AWF command from agent config.
//
// Parameters:
//   - workflowData: The workflow data containing agent configuration
//
// Returns:
//   - string: The AWF command to use (e.g., "sudo -E awf" or custom command)
func GetAWFCommandPrefix(workflowData *WorkflowData) string {
	agentConfig := getAgentConfig(workflowData)
	if agentConfig != nil && agentConfig.Command != "" {
		awfHelpersLog.Printf("Using custom AWF command: %s", agentConfig.Command)
		return agentConfig.Command
	}

	awfHelpersLog.Print("Using standard AWF command")
	return string(constants.AWFDefaultCommand)
}

// buildAWFImageTagWithDigests returns an image tag value for AWF's --image-tag flag.
// When known firewall container digests are available, it appends AWF's digest
// metadata format:
//
//	<tag>,squid=sha256:...,agent=sha256:...,api-proxy=sha256:...,cli-proxy=sha256:...
//
// This keeps AWF sidecar configuration aligned with digest-pinned pre-download images.
func buildAWFImageTagWithDigests(imageTag string, workflowData *WorkflowData) string {
	if imageTag == "" {
		return imageTag
	}

	type digestSpec struct {
		name  string
		image string
	}
	specs := []digestSpec{
		{name: "squid", image: constants.DefaultFirewallRegistry + "/squid:" + imageTag},
		{name: "agent", image: constants.DefaultFirewallRegistry + "/agent:" + imageTag},
		{name: "agent-act", image: constants.DefaultFirewallRegistry + "/agent-act:" + imageTag},
		{name: "api-proxy", image: constants.DefaultFirewallRegistry + "/api-proxy:" + imageTag},
		{name: "cli-proxy", image: constants.DefaultFirewallRegistry + "/cli-proxy:" + imageTag},
	}

	parts := []string{imageTag}
	for _, spec := range specs {
		digest := lookupContainerDigest(spec.image, workflowData)
		if digest == "" {
			continue
		}
		parts = append(parts, spec.name+"="+digest)
	}

	if len(parts) == 1 {
		return imageTag
	}
	return strings.Join(parts, ",")
}

// lookupContainerDigest resolves a container image digest from cache first, then
// falls back to embedded container pins.
func lookupContainerDigest(image string, workflowData *WorkflowData) string {
	var cache *ActionCache
	if workflowData != nil {
		cache = workflowData.ActionCache
	}
	if pin, ok := lookupContainerPin(image, cache); ok && pin.Digest != "" {
		return pin.Digest
	}
	return ""
}

// WrapCommandInShell wraps an engine command in a shell invocation for AWF execution.
// This is needed because AWF requires commands to be wrapped in shell for proper execution.
//
// set +o histexpand disables bash history expansion so that agent-authored strings
// containing '!' characters (e.g. "!**") cannot be silently misinterpreted or dropped.
// History expansion is meaningless for non-interactive execution and has no other effect.
//
// Parameters:
//   - command: The engine command to wrap (may include PATH setup and other initialization)
//
// Returns:
//   - string: Shell-wrapped command suitable for AWF execution
func WrapCommandInShell(command string) string {
	awfHelpersLog.Print("Wrapping command in shell for AWF execution")

	// Escape single quotes in the command by replacing ' with '\''
	escapedCommand := strings.ReplaceAll(command, "'", "'\\''")

	// Wrap in shell invocation.
	// set +o histexpand is first to prevent bash from expanding !-patterns in any
	// double-quoted strings that appear in the engine command or its arguments.
	return fmt.Sprintf("/bin/bash -c 'set +o histexpand; %s'", escapedCommand)
}

// addSecretEnvVarsFromMap calls add() for each key in envMap whose value contains a ${{ secrets.* }} reference.
func addSecretEnvVarsFromMap(add func(string), envMap map[string]string) {
for varName, varValue := range envMap {
if strings.Contains(varValue, "${{ secrets.") {
add(varName)
}
}
}

// ComputeAWFExcludeEnvVarNames returns the list of environment variable names that must be
// excluded from the agent container's visible environment via AWF's --exclude-env flag.
//
// Only env var names whose step-env value WILL contain a ${{ secrets.* }} reference are
// included, so non-secret vars (e.g. GH_DEBUG: "1" in mcp-scripts) are never excluded.
//
// Parameters:
//   - workflowData: the workflow being compiled
//   - coreSecretVarNames: engine-specific fixed secret env var names (e.g. ["COPILOT_GITHUB_TOKEN"])
//
// The function augments coreSecretVarNames with:
//   - MCP_GATEWAY_API_KEY when MCP servers are present
//   - GITHUB_MCP_SERVER_TOKEN when the GitHub tool is present
//   - HTTP MCP header secret var names (values always contain ${{ secrets.* }})
//   - mcp-scripts env var names whose values contain ${{ secrets.* }}
//   - engine.env var names whose values contain ${{ secrets.* }}
//   - agent.env var names whose values contain ${{ secrets.* }}
func ComputeAWFExcludeEnvVarNames(workflowData *WorkflowData, coreSecretVarNames []string) []string {
seen := make(map[string]struct{})
var names []string
add := func(name string) {
if _, ok := seen[name]; !ok {
seen[name] = struct{}{}
names = append(names, name)
}
}
for _, name := range coreSecretVarNames {
add(name)
}
if HasMCPServers(workflowData) {
add("MCP_GATEWAY_API_KEY")
}
if hasGitHubTool(workflowData.ParsedTools) {
add("GITHUB_MCP_SERVER_TOKEN")
}
for varName := range collectHTTPMCPHeaderSecrets(workflowData.Tools) {
add(varName)
}
if workflowData.MCPScripts != nil {
for _, toolConfig := range workflowData.MCPScripts.Tools {
addSecretEnvVarsFromMap(add, toolConfig.Env)
}
}
if workflowData.EngineConfig != nil {
addSecretEnvVarsFromMap(add, workflowData.EngineConfig.Env)
}
if agentConfig := getAgentConfig(workflowData); agentConfig != nil {
addSecretEnvVarsFromMap(add, agentConfig.Env)
}
if isGitHubCLIModeEnabled(workflowData) {
add("GH_TOKEN")
}
awfHelpersLog.Printf("Computed %d AWF env vars to exclude", len(names))
return names
}

// addCliProxyGHTokenToEnv adds GH_TOKEN to the AWF step environment when GitHub
// mode is gh-proxy. The token is NOT used by AWF or its cli-proxy
// sidecar directly — the host difc-proxy (started by start_cli_proxy.sh) already
// has it. However, --env-all passes all step env vars into the agent container,
// so we explicitly set GH_TOKEN here to ensure --exclude-env GH_TOKEN can
// reliably strip it regardless of how the token enters the environment.
// The token is excluded from the agent container via --exclude-env GH_TOKEN, so only
// inject it when the effective AWF version supports both cli-proxy flags and
// --exclude-env.
//
// #nosec G101 -- This is NOT a hardcoded credential. It is a GitHub Actions expression
// template that is resolved at runtime by the GitHub Actions runner.
func addCliProxyGHTokenToEnv(env map[string]string, workflowData *WorkflowData) {
	firewallConfig := getFirewallConfig(workflowData)
	if isGitHubCLIModeEnabled(workflowData) &&
		isFirewallEnabled(workflowData) &&
		awfSupportsCliProxy(firewallConfig) &&
		awfSupportsExcludeEnv(firewallConfig) {
		env["GH_TOKEN"] = "${{ secrets.GH_AW_GITHUB_TOKEN || github.token }}"
		awfHelpersLog.Print("Added GH_TOKEN to env for CLI proxy (excluded from agent container)")
	}
}

// awfSupportsExcludeEnv returns true when the effective AWF version supports --exclude-env
// (introduced in AWF v0.25.3).
func awfSupportsExcludeEnv(firewallConfig *FirewallConfig) bool {
	return awfVersionAtLeast(firewallConfig, constants.AWFExcludeEnvMinVersion)
}

// awfVersionAtLeast returns true when the effective AWF version is at or above minVersion.
//
// If firewallConfig has no version set, DefaultFirewallVersion is used. "latest" always
// returns true. Non-semver strings (e.g. branch names) return false (conservative).
func awfVersionAtLeast(firewallConfig *FirewallConfig, minVersion constants.Version) bool {
	var versionStr string
	if firewallConfig != nil && firewallConfig.Version != "" {
		versionStr = firewallConfig.Version
	}
	return versionAtLeast(versionStr, string(constants.DefaultFirewallVersion), string(minVersion))
}

// awfSupportsCliProxy returns true when the effective AWF version supports --difc-proxy-host
// and --difc-proxy-ca-cert (introduced in AWF v0.26.0).
func awfSupportsCliProxy(firewallConfig *FirewallConfig) bool {
	return awfVersionAtLeast(firewallConfig, constants.AWFCliProxyMinVersion)
}

// awfSupportsAllowHostPorts returns true when the effective AWF version supports
// --allow-host-ports.
func awfSupportsAllowHostPorts(firewallConfig *FirewallConfig) bool {
	return awfVersionAtLeast(firewallConfig, constants.AWFAllowHostPortsMinVersion)
}

// awfSupportsDockerHostPathPrefix returns true when the effective AWF version supports
// --docker-host-path-prefix.
func awfSupportsDockerHostPathPrefix(firewallConfig *FirewallConfig) bool {
	return awfVersionAtLeast(firewallConfig, constants.AWFDockerHostPathPrefixMinVersion)
}

// awfSupportsTokenSteering returns true when the effective AWF version supports
// apiProxy.enableTokenSteering.
func awfSupportsTokenSteering(firewallConfig *FirewallConfig) bool {
	return awfVersionAtLeast(firewallConfig, constants.AWFTokenSteeringMinVersion)
}

// awfSupportsChrootConfig returns true when the effective AWF version supports
// chroot.binariesSourcePath and chroot.identity.* in the config file (AWF v0.27.1+).
func awfSupportsChrootConfig(firewallConfig *FirewallConfig) bool {
	return awfVersionAtLeast(firewallConfig, constants.AWFChrootConfigMinVersion)
}

// buildArcDindChrootConfigPatchBody returns the Python heredoc that patches the AWF
// config file with chroot.binariesSourcePath and chroot.identity.*. It is designed to be
// embedded inside a bash if-block that already guards on DOCKER_HOST=tcp://...
//
// The Python is intentionally kept compact to minimise script size and stay within
// GitHub Actions' 21 KB per-step expression limit.
// Both config paths are updated: ${RUNNER_TEMP}/gh-aw/awf-config.json (read by AWF) and
// /tmp/gh-aw/awf-config.json (used by the unified agent artifact upload).
func buildArcDindChrootConfigPatchBody() string {
	return fmt.Sprintf(`  python3 - <<'PY'
import json,os,subprocess as sp
from pathlib import Path
try:
 p=Path(os.environ["RUNNER_TEMP"])/"gh-aw"/"awf-config.json"
 c=json.loads(p.read_text())
 c["chroot"]={"binariesSourcePath":"%s","identity":{"user":sp.check_output(["id","-un"],text=True).strip(),"uid":int(sp.check_output(["id","-u"],text=True)),"gid":int(sp.check_output(["id","-g"],text=True)),"home":"%s"}}
 out=json.dumps(c,separators=(",",":"),ensure_ascii=False)+"\n"
 p.write_text(out)
 Path("%s/awf-config.json").write_text(out)
except Exception as e:
 raise SystemExit(f"chroot config patch failed: {e}") from e
PY`, awfArcDindChrootBinariesSourcePath, awfArcDindChrootIdentityHome, awfArcDindChrootBinariesSourcePath)
}
