// This file provides Copilot engine installation logic.
//
// This file contains functions for generating GitHub Actions steps to install
// the GitHub Copilot CLI and related sandbox infrastructure (AWF or SRT).
//
// Installation order:
//  1. Secret validation (COPILOT_GITHUB_TOKEN) — runs in the activation job
//  2. Node.js setup
//  3. Sandbox installation (SRT or AWF, if needed)
//  4. Copilot CLI installation
//
// The installation strategy differs based on sandbox mode:
//   - Standard mode: Global installation using official installer script
//   - SRT mode: Local npm installation for offline compatibility
//   - AWF mode: Global installation + AWF binary

package workflow

import (
	"strings"

	"github.com/github/gh-aw/pkg/constants"
	"github.com/github/gh-aw/pkg/logger"
)

var copilotInstallLog = logger.New("workflow:copilot_engine_installation")

// GetSecretValidationStep returns the secret validation step for the Copilot engine.
// Returns an empty step if:
//   - copilot-requests feature is enabled (uses GitHub Actions token instead), or
//   - COPILOT_PROVIDER_API_KEY or COPILOT_PROVIDER_BEARER_TOKEN is set in engine.env
//     (BYOK mode — the external provider handles authentication, so COPILOT_GITHUB_TOKEN
//     is not required for model routing).
func (e *CopilotEngine) GetSecretValidationStep(workflowData *WorkflowData) GitHubActionStep {
	if isFeatureEnabled(constants.CopilotRequestsFeatureFlag, workflowData) {
		copilotInstallLog.Print("Skipping secret validation step: copilot-requests feature enabled, using GitHub Actions token")
		return GitHubActionStep{}
	}
	if engineEnvHasKey(workflowData, constants.CopilotProviderAPIKey) ||
		engineEnvHasKey(workflowData, constants.CopilotProviderBearerToken) {
		copilotInstallLog.Print("Skipping COPILOT_GITHUB_TOKEN validation: BYOK provider credentials are configured")
		return GitHubActionStep{}
	}
	return BuildDefaultSecretValidationStep(
		workflowData,
		[]string{"COPILOT_GITHUB_TOKEN"},
		"GitHub Copilot CLI",
		"https://github.github.com/gh-aw/reference/engines/#github-copilot-default",
	)
}

// GetInstallationSteps generates the complete installation workflow for Copilot CLI.
// This includes Node.js setup, sandbox installation (SRT or AWF), and Copilot CLI installation.
// Secret validation is handled separately in the activation job via GetSecretValidationStep.
// The installation order is:
// 1. Node.js setup
// 2. Sandbox installation (AWF, if needed)
// 3. Copilot CLI installation
//
// If a custom command is specified in the engine configuration, this function skips
// standard Copilot CLI installation. When firewall is enabled, it still returns AWF
// runtime installation steps required for harness execution.
func (e *CopilotEngine) GetInstallationSteps(workflowData *WorkflowData) []GitHubActionStep {
	copilotInstallLog.Printf("Generating installation steps for Copilot engine: workflow=%s", workflowData.Name)

	// Skip standard Copilot CLI installation if custom command is specified.
	if workflowData.EngineConfig != nil && workflowData.EngineConfig.Command != "" {
		// Keep firewall runtime installation when firewall is enabled, since the
		// custom engine command still runs inside the AWF harness.
		if isFirewallEnabled(workflowData) {
			copilotInstallLog.Printf("Skipping Copilot CLI installation: custom command specified (%s); keeping AWF runtime installation because firewall is enabled", workflowData.EngineConfig.Command)
			return BuildNpmEngineInstallStepsWithAWF([]GitHubActionStep{}, workflowData)
		}
		copilotInstallLog.Printf("Skipping installation steps: custom command specified (%s)", workflowData.EngineConfig.Command)
		return []GitHubActionStep{}
	}

	// Copilot CLI is pinned to the default version constant.
	copilotVersion := string(constants.DefaultCopilotVersion)
	if workflowData.EngineConfig != nil {
		if workflowData.EngineConfig.Version != "" {
			copilotInstallLog.Printf("Ignoring pinned engine.version (%s): Copilot CLI install version is pinned to %s", workflowData.EngineConfig.Version, copilotVersion)
		}
		// Normalize engine config version to effective installed version so
		// downstream checks that consult EngineConfig.Version stay consistent.
		// This applies even when the original version was empty (unset), so all
		// downstream consumers observe the effective installed value.
		// This mutates workflowData by design because subsequent generation steps
		// in the same compile flow should observe the effective installed version.
		// Callers that reuse the same WorkflowData instance should expect this
		// field to be rewritten after installation-step generation.
		workflowData.EngineConfig.Version = copilotVersion
	}

	// Use the installer script for global installation
	copilotInstallLog.Print("Using new installer script for Copilot installation")
	npmSteps := GenerateCopilotInstallerSteps(copilotVersion, "Install GitHub Copilot CLI")

	// When the setup-copilot-resolver feature is enabled, gate the bash installer
	// step on the setup action's `copilot-cached` output. On cache hit, the setup
	// action already added the cached CLI to PATH and this step is skipped.
	// On cache miss, the installer runs as before.
	if shouldUseCopilotResolver(workflowData) {
		copilotInstallLog.Print("setup-copilot-resolver enabled: gating installer step on steps.setup.outputs.copilot-cached")
		npmSteps = gateStepsOnCopilotCached(npmSteps)
	}

	return BuildNpmEngineInstallStepsWithAWF(npmSteps, workflowData)
}

// shouldUseCopilotResolver reports whether the setup action's Copilot CLI
// resolver should be activated for this workflow. It requires both:
//   - the workflow's engine to be Copilot (only Copilot uses the toolcache bake), and
//   - the SetupCopilotResolverFeatureFlag to be enabled (default off until validated).
//
// Used by:
//   - GetInstallationSteps (this file): gate the bash installer step
//   - compiler_main_job.go: pass installCopilot=true to generateSetupStep
//   - threat_detection.go: pass installCopilot=true to generateSetupStep
func shouldUseCopilotResolver(workflowData *WorkflowData) bool {
	if workflowData == nil {
		return false
	}
	engineID := ""
	if workflowData.EngineConfig != nil && workflowData.EngineConfig.ID != "" {
		engineID = workflowData.EngineConfig.ID
	} else if workflowData.AI != "" {
		engineID = workflowData.AI
	}
	if engineID != "copilot" {
		return false
	}
	return isFeatureEnabled(constants.SetupCopilotResolverFeatureFlag, workflowData)
}

// gateStepsOnCopilotCached injects `if: steps.setup.outputs.copilot-cached != 'true'`
// into each step's YAML so the bash installer is skipped when the resolver hit
// the toolcache. The `if:` line is inserted directly after the `- name:` line
// (which is conventionally the first line of each step emitted by
// GenerateCopilotInstallerSteps).
//
// Steps without a recognisable `      - name:` opener are returned unmodified;
// any future refactor of the installer step shape should re-verify this here.
func gateStepsOnCopilotCached(steps []GitHubActionStep) []GitHubActionStep {
	const condition = "        if: steps.setup.outputs.copilot-cached != 'true'"
	out := make([]GitHubActionStep, 0, len(steps))
	for _, step := range steps {
		if len(step) == 0 {
			out = append(out, step)
			continue
		}
		if !strings.HasPrefix(step[0], "      - name:") {
			out = append(out, step)
			continue
		}
		gated := make([]string, 0, len(step)+1)
		gated = append(gated, step[0], condition)
		gated = append(gated, step[1:]...)
		out = append(out, GitHubActionStep(gated))
	}
	return out
}

// generateAWFInstallationStep creates a GitHub Actions step to install the AWF binary
// with SHA256 checksum verification to protect against supply chain attacks.
//
// The installation logic is implemented in a separate shell script (install_awf_binary.sh)
// which downloads the binary directly from GitHub releases, verifies its checksum against
// the official checksums.txt file, and installs it. This approach:
// - Eliminates trust in the installer script itself
// - Provides full transparency of the installation process
// - Protects against tampered or compromised installer scripts
// - Verifies the binary integrity before execution
//
// If a custom command is specified in the agent config, the installation is skipped
// as the custom command replaces the AWF binary.
func generateAWFInstallationStep(version string, agentConfig *AgentSandboxConfig) GitHubActionStep {
	// If custom command is specified, skip installation (command replaces binary)
	if agentConfig != nil && agentConfig.Command != "" {
		copilotInstallLog.Print("Skipping AWF binary installation (custom command specified)")
		// Return empty step - custom command will be used in execution
		return GitHubActionStep([]string{})
	}

	// Use default version for logging when not specified
	if version == "" {
		version = string(constants.DefaultFirewallVersion)
	}

	stepLines := []string{
		"      - name: Install AWF binary",
		"        run: bash \"${RUNNER_TEMP}/gh-aw/actions/install_awf_binary.sh\" " + version,
	}

	return GitHubActionStep(stepLines)
}
