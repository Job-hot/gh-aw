package workflow

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/github/gh-aw/pkg/console"
	"github.com/github/gh-aw/pkg/constants"
)

// generateMainJobSteps generates the complete sequence of steps for the main agent execution job
// This is the heart of the workflow, orchestrating all steps from checkout through AI execution to artifact upload
func (c *Compiler) generateMainJobSteps(yaml *strings.Builder, data *WorkflowData) error {
	compilerYamlLog.Printf("Generating main job steps for workflow: %s", data.Name)

	// Phase 1: Initial setup, checkout, and repository imports
	checkoutMgr, needsCheckout, err := c.generateInitialAndCheckoutSteps(yaml, data)
	if err != nil {
		return err
	}

	// Phase 2: Runtime detection, custom steps, and workspace setup
	customStepsContainCheckout := c.generateRuntimeAndWorkspaceSetupSteps(yaml, data, needsCheckout)
	needsGitConfig := needsCheckout || customStepsContainCheckout

	// Phase 3: Engine installation, MCP setup, and pre-agent preparation
	engine, err := c.generateEngineInstallAndPreAgentSteps(yaml, data, needsGitConfig)
	if err != nil {
		return err
	}

	// Pre-warm the allowed domains cache so that engine execution steps (GetExecutionSteps)
	// can reuse the pre-computed value instead of re-running the expensive map+sort
	// operation inside each engine's domain helper.  The result is stored on WorkflowData
	// and is also used later by generateOutputCollectionStep.
	_, _ = c.computeAllowedDomainsForSanitization(data)

	// Phase 4: Agent execution and immediate post-agent steps
	artifactPaths, logFileFull, err := c.generateAgentRunSteps(yaml, data, engine, needsGitConfig)
	if err != nil {
		return err
	}

	// Phase 5: Artifact collection, log parsing, upload, and cleanup
	return c.generatePostAgentCollectionAndUpload(yaml, data, engine, artifactPaths, logFileFull, checkoutMgr)
}

// generateInitialAndCheckoutSteps emits the OTLP mask step, pre-steps, all checkout steps
// (default workspace checkout, dev-mode CLI build, additional checkouts), repository import
// checkouts, legacy agent import checkout, and the merge-.github-folder step.
// It returns the CheckoutManager (needed later for token invalidation and dev-mode restore)
// and a flag indicating whether the default workspace checkout was emitted.
func (c *Compiler) generateInitialAndCheckoutSteps(yaml *strings.Builder, data *WorkflowData) (*CheckoutManager, bool, error) {
	// Mask OTLP telemetry headers early so authentication tokens cannot leak in runner
	// debug logs. The workflow-level OTEL_EXPORTER_OTLP_HEADERS env var is available
	// from the very first step, so masking can happen before any other work.
	if isOTLPHeadersPresent(data) {
		yaml.WriteString(generateOTLPHeadersMaskStep())
	}
	// Mask custom OTLP attribute values so user-supplied values cannot leak into runner logs.
	if isOTLPAttributesPresent(data) {
		yaml.WriteString(generateOTLPAttributesMaskStep())
	}

	// Add pre-steps before checkout and the subsequent built-in steps in this agent job.
	// This allows users to mint short-lived tokens (via custom actions) in the same
	// job as checkout, so the tokens are never dropped by the GitHub Actions runner's
	// add-mask behaviour that silently suppresses masked values across job boundaries.
	// Step outputs are available as ${{ steps.<id>.outputs.<name> }} and can be
	// referenced directly in checkout.token. Some compiler-injected setup steps may
	// still be emitted earlier than these pre-steps.
	c.generatePreSteps(yaml, data)

	// Determine if we need to add a checkout step
	needsCheckout := c.shouldAddCheckoutStep(data)
	compilerYamlLog.Printf("Checkout step needed: %t", needsCheckout)

	// Build a CheckoutManager with any user-configured checkouts
	checkoutMgr := NewCheckoutManager(data.CheckoutConfigs)

	// Propagate the platform (host) repo resolved by the activation job so that
	// checkout steps in this job and in safe_outputs can use the correct repository
	// for .github/.agents sparse checkouts when called cross-repo.
	// The activation job exposes this as needs.activation.outputs.target_repo.
	if hasWorkflowCallTrigger(data.On) && !data.InlinedImports {
		checkoutMgr.SetCrossRepoTargetRepo("${{ needs.activation.outputs.target_repo }}")
	}

	// Mint checkout app tokens directly in the agent job before checkout steps are executed.
	c.emitCheckoutAppTokenSteps(yaml, data, checkoutMgr)

	// Emit default checkout, dev-mode build, additional checkouts, and manifest step.
	c.emitDefaultCheckoutAndAdditional(yaml, data, checkoutMgr, needsCheckout)

	// Emit checkout steps for repository imports and legacy agent import.
	c.emitImportCheckoutSteps(yaml, data)

	// Add merge remote .github folder step for repository imports or agent imports.
	if err := c.emitMergeGitHubFolderStep(yaml, data); err != nil {
		return nil, false, err
	}

	return checkoutMgr, needsCheckout, nil
}

// emitCheckoutAppTokenSteps mints checkout app tokens directly in the agent job
// before checkout steps run, making them available as step outputs within the same job.
func (c *Compiler) emitCheckoutAppTokenSteps(yaml *strings.Builder, data *WorkflowData, checkoutMgr *CheckoutManager) {
	if !checkoutMgr.HasAppAuth() {
		return
	}
	compilerYamlLog.Print("Generating checkout app token minting steps in agent job")
	var checkoutPermissions *Permissions
	if data.CachedPermissions != nil {
		checkoutPermissions = data.CachedPermissions
	} else if data.Permissions != "" {
		checkoutPermissions = NewPermissionsParser(data.Permissions).ToPermissions()
	} else {
		checkoutPermissions = NewPermissions()
	}
	for _, step := range checkoutMgr.GenerateCheckoutAppTokenSteps(c, checkoutPermissions) {
		yaml.WriteString(step)
	}
}

// emitDefaultCheckoutAndAdditional emits the default workspace checkout (with optional dev-mode
// CLI build), additional user-configured checkouts, and the checkout manifest step.
func (c *Compiler) emitDefaultCheckoutAndAdditional(yaml *strings.Builder, data *WorkflowData, checkoutMgr *CheckoutManager, needsCheckout bool) {
	if needsCheckout {
		// Emit the default workspace checkout, applying any user-supplied overrides
		for _, line := range checkoutMgr.GenerateDefaultCheckoutStep(c.trialMode, c.trialLogicalRepoSlug, c.getActionPin) {
			yaml.WriteString(line)
		}
		// Add CLI build steps in dev mode (after automatic checkout, before other steps)
		if c.actionMode.IsDev() {
			if _, hasAgenticWorkflows := data.Tools["agentic-workflows"]; hasAgenticWorkflows {
				compilerYamlLog.Printf("Generating CLI build steps for dev mode (agentic-workflows tool enabled)")
				c.generateDevModeCLIBuildSteps(yaml)
			} else {
				compilerYamlLog.Printf("Skipping CLI build steps in dev mode (agentic-workflows tool not enabled)")
			}
		}
	}
	// Emit additional (non-default) user-configured checkouts
	for _, line := range checkoutMgr.GenerateAdditionalCheckoutSteps(c.getActionPin) {
		yaml.WriteString(line)
	}
	// Emit a manifest step that records the path and resolved default branch for each
	// non-default cross-repo checkout.
	for _, line := range checkoutMgr.GenerateCheckoutManifestStep(c.getActionPin) {
		yaml.WriteString(line)
	}
}

// emitImportCheckoutSteps emits checkout steps for repository imports and the legacy agent import.
func (c *Compiler) emitImportCheckoutSteps(yaml *strings.Builder, data *WorkflowData) {
	if len(data.RepositoryImports) > 0 {
		compilerYamlLog.Printf("Adding checkout steps for %d repository imports", len(data.RepositoryImports))
		c.generateRepositoryImportCheckouts(yaml, data.RepositoryImports)
	}
	if data.AgentFile != "" && data.AgentImportSpec != "" {
		compilerYamlLog.Printf("Adding checkout step for legacy agent import: %s", data.AgentImportSpec)
		c.generateLegacyAgentImportCheckout(yaml, data.AgentImportSpec)
	}
}

// emitMergeGitHubFolderStep emits the merge-remote-.github-folder step when repository imports
// or legacy agent imports are present.
func (c *Compiler) emitMergeGitHubFolderStep(yaml *strings.Builder, data *WorkflowData) error {
	needsGithubMerge := (len(data.RepositoryImports) > 0) || (data.AgentFile != "" && data.AgentImportSpec != "")
	if !needsGithubMerge {
		return nil
	}
	compilerYamlLog.Printf("Adding merge remote .github folder step")
	yaml.WriteString("      - name: Merge remote .github folder\n")
	fmt.Fprintf(yaml, "        uses: %s\n", getCachedActionPin("actions/github-script", data))
	yaml.WriteString("        env:\n")
	if len(data.RepositoryImports) > 0 {
		repoImportsJSON, err := json.Marshal(data.RepositoryImports)
		if err != nil {
			return fmt.Errorf("failed to marshal repository imports for merge step: %w", err)
		}
		writeYAMLEnv(yaml, "          ", "GH_AW_REPOSITORY_IMPORTS", string(repoImportsJSON))
	}
	if data.AgentFile != "" && data.AgentImportSpec != "" {
		writeYAMLEnv(yaml, "          ", "GH_AW_AGENT_FILE", data.AgentFile)
		writeYAMLEnv(yaml, "          ", "GH_AW_AGENT_IMPORT_SPEC", data.AgentImportSpec)
	}
	yaml.WriteString("        with:\n")
	yaml.WriteString("          script: |\n")
	yaml.WriteString("            const { setupGlobals } = require('${{ runner.temp }}/gh-aw/actions/setup_globals.cjs');\n")
	yaml.WriteString("            setupGlobals(core, github, context, exec, io, getOctokit);\n")
	yaml.WriteString("            const { main } = require('${{ runner.temp }}/gh-aw/actions/merge_remote_agent_github_folder.cjs');\n")
	yaml.WriteString("            await main();\n")
	return nil
}

// generateRuntimeAndWorkspaceSetupSteps emits runtime setup steps, the gh-aw temp directory
// creation step, GitHub Enterprise CLI configuration, DIFC proxy start, custom steps, cache
// steps, cache-memory steps, and repo-memory steps.
// It mutates data.CustomSteps (via deduplication) and returns whether the custom steps
// themselves contain a checkout action (used by the caller to compute needsGitConfig).
func (c *Compiler) generateRuntimeAndWorkspaceSetupSteps(yaml *strings.Builder, data *WorkflowData, needsCheckout bool) bool {
	runtimeSetupSteps := resolveRuntimeSetupSteps(data)

	customStepsContainCheckout := data.CustomSteps != "" && ContainsCheckout(data.CustomSteps)
	compilerYamlLog.Printf("Custom steps contain checkout: %t (len(customSteps)=%d)", customStepsContainCheckout, len(data.CustomSteps))

	emitRuntimeStepsIfBeforeCustom(yaml, needsCheckout, customStepsContainCheckout, runtimeSetupSteps)

	// Create /tmp/gh-aw/ base directory for all temporary files
	// This must be created before custom steps so they can use the temp directory
	yaml.WriteString("      - name: Create gh-aw temp directory\n")
	yaml.WriteString("        run: bash \"${RUNNER_TEMP}/gh-aw/actions/create_gh_aw_tmp_dir.sh\"\n")

	// Configure gh CLI for GitHub Enterprise hosts (*.ghe.com / GHES).
	// This step runs configure_gh_for_ghe.sh which:
	//   1. Detects the GitHub host from GITHUB_SERVER_URL
	//   2. For github.com: exits immediately (no-op)
	//   3. For GHE/GHES: authenticates gh CLI with the enterprise host and sets
	//      GH_HOST=<host> in GITHUB_ENV so every subsequent step in this job
	//      picks up the correct host without manual per-step configuration.
	// Must run after the setup action (so the script is available at ${RUNNER_TEMP}/gh-aw/actions/)
	// and before any custom steps that invoke gh CLI commands.
	yaml.WriteString("      - name: Configure gh CLI for GitHub Enterprise\n")
	yaml.WriteString("        run: bash \"${RUNNER_TEMP}/gh-aw/actions/configure_gh_for_ghe.sh\"\n")
	yaml.WriteString("        env:\n")
	yaml.WriteString("          GH_TOKEN: ${{ github.token }}\n")

	// Start DIFC proxy for pre-agent gh CLI calls (only when guard policies are configured
	// and pre-agent steps with GH_TOKEN are present). The proxy routes gh CLI calls through
	// integrity filtering before the agent runs. Must start before custom steps.
	c.generateStartDIFCProxyStep(yaml, data)

	c.emitCustomStepsBlock(yaml, data, customStepsContainCheckout, runtimeSetupSteps)

	// Add cache steps if cache configuration is present
	compilerYamlLog.Printf("Generating cache steps for workflow")
	generateCacheSteps(yaml, data, c.verbose)

	// Add cache-memory steps if cache-memory configuration is present
	compilerYamlLog.Printf("Generating cache-memory steps for workflow")
	generateCacheMemorySteps(yaml, data)

	// Add repo-memory clone steps if repo-memory configuration is present
	compilerYamlLog.Printf("Generating repo-memory steps for workflow")
	generateRepoMemorySteps(yaml, data)

	return customStepsContainCheckout
}

// resolveRuntimeSetupSteps detects runtime requirements from the workflow data, deduplicates
// them against any user-supplied custom steps, and returns the final set of setup steps.
// It may mutate data.CustomSteps to remove user-customized setup actions.
func resolveRuntimeSetupSteps(data *WorkflowData) []GitHubActionStep {
	requirements := DetectRuntimeRequirements(data)
	if len(requirements) > 0 && data.CustomSteps != "" {
		deduplicated, filtered, err := DeduplicateRuntimeSetupStepsFromCustomSteps(data.CustomSteps, requirements)
		if err != nil {
			compilerYamlLog.Printf("Warning: failed to deduplicate runtime setup steps: %v", err)
		} else {
			data.CustomSteps = deduplicated
			requirements = filtered
		}
	}
	steps := GenerateRuntimeSetupSteps(requirements, data)
	compilerYamlLog.Printf("Detected runtime requirements: %d runtimes, %d setup steps", len(requirements), len(steps))
	return steps
}

// emitRuntimeStepsIfBeforeCustom emits runtime setup steps before custom steps when either the
// workspace was already checked out (needsCheckout) or custom steps contain no checkout of their own.
func emitRuntimeStepsIfBeforeCustom(yaml *strings.Builder, needsCheckout, customStepsContainCheckout bool, runtimeSetupSteps []GitHubActionStep) {
	if needsCheckout || !customStepsContainCheckout {
		compilerYamlLog.Printf("Adding %d runtime steps before custom steps (needsCheckout=%t, !customStepsContainCheckout=%t)", len(runtimeSetupSteps), needsCheckout, !customStepsContainCheckout)
		for _, step := range runtimeSetupSteps {
			for _, line := range step {
				yaml.WriteString(line)
				yaml.WriteByte('\n')
			}
		}
	}
}

// emitCustomStepsBlock emits the custom steps, optionally injecting proxy env vars and
// inserting runtime setup steps after the first checkout when needed.
func (c *Compiler) emitCustomStepsBlock(yaml *strings.Builder, data *WorkflowData, customStepsContainCheckout bool, runtimeSetupSteps []GitHubActionStep) {
	if data.CustomSteps == "" {
		return
	}
	customStepsToEmit := data.CustomSteps
	if hasDIFCProxyNeeded(data) {
		customStepsToEmit = injectProxyEnvIntoCustomSteps(customStepsToEmit)
	}
	if customStepsContainCheckout && len(runtimeSetupSteps) > 0 {
		compilerYamlLog.Printf("Calling addCustomStepsWithRuntimeInsertion: %d runtime steps to insert after checkout", len(runtimeSetupSteps))
		c.addCustomStepsWithRuntimeInsertion(yaml, customStepsToEmit, runtimeSetupSteps, data.ParsedTools)
	} else {
		compilerYamlLog.Printf("Calling addCustomStepsAsIs (customStepsContainCheckout=%t, runtimeStepsCount=%d)", customStepsContainCheckout, len(runtimeSetupSteps))
		c.addCustomStepsAsIs(yaml, customStepsToEmit)
	}
}

// generateEngineInstallAndPreAgentSteps emits git credential configuration, the PR-ready-for-review
// checkout, engine installation steps, GitHub MCP app token minting, MCP lockdown detection, guard
// variable parsing, DIFC proxy stop, activation artifact download, comment-memory file preparation,
// base-.github-folder restore, pre-agent steps, MCP gateway setup, and MCP CLI mount.
// It returns the resolved CodingAgentEngine for use in subsequent phases.
func (c *Compiler) generateEngineInstallAndPreAgentSteps(yaml *strings.Builder, data *WorkflowData, needsGitConfig bool) (CodingAgentEngine, error) {
	// Configure git credentials for agentic workflows.
	// Git credential configuration requires a .git directory in the workspace, which is only
	// present when the repository was checked out. Skip these steps when checkout is disabled
	// and no custom steps perform a checkout, since git remote set-url origin would fail
	// with "fatal: not a git repository" otherwise.
	compilerYamlLog.Printf("Git credential configuration needed: %t", needsGitConfig)
	if needsGitConfig {
		for _, line := range c.generateGitConfigurationSteps() {
			yaml.WriteString(line)
		}
	}

	// Add step to checkout PR branch if the event is pull_request
	c.generatePRReadyForReviewCheckout(yaml, data)

	// Resolve the agentic engine and ensure MCP gateway defaults are set
	engine, err := c.getAgenticEngine(data.AI)
	if err != nil {
		return nil, fmt.Errorf("failed to resolve agentic engine from AI configuration: %w", err)
	}
	if HasMCPServers(data) {
		ensureDefaultMCPGatewayConfig(data)
	}

	// Emit engine installation steps (Node.js, Playwright CLI, etc.)
	c.emitEngineAndPlaywrightInstallSteps(yaml, data, engine)

	// Mint the GitHub MCP App token, emit lockdown detection, guard vars, and stop DIFC proxy.
	c.emitGitHubMCPTokenAndLockdownSteps(yaml, data)

	// Download activation artifact and prepare comment-memory files.
	c.emitActivationArtifactAndCommentMemorySteps(yaml, data)

	// Restore base .github folders and inline sub-agents from the activation artifact.
	c.emitBaseGitHubRestoreAndInlineSubAgentsSteps(yaml, data)

	// Add pre-agent steps, MCP setup, and MCP CLI mount.
	return engine, c.emitPreAgentAndMCPSetup(yaml, data, engine)
}

// emitEngineAndPlaywrightInstallSteps emits engine-specific installation steps followed by
// Playwright CLI install steps when playwright is configured in CLI mode.
func (c *Compiler) emitEngineAndPlaywrightInstallSteps(yaml *strings.Builder, data *WorkflowData, engine CodingAgentEngine) {
	installSteps := engine.GetInstallationSteps(data)
	compilerYamlLog.Printf("Adding %d engine installation steps for %s", len(installSteps), engine.GetID())
	for _, step := range installSteps {
		for _, line := range step {
			yaml.WriteString(line)
			yaml.WriteByte('\n')
		}
	}
	// Add Playwright CLI install steps when playwright is configured in CLI mode.
	// These run after Node.js is available (set up by the engine install steps above).
	for _, step := range generatePlaywrightCLIInstallSteps(data) {
		for _, line := range step {
			yaml.WriteString(line)
			yaml.WriteByte('\n')
		}
	}
}

// emitGitHubMCPTokenAndLockdownSteps mints the GitHub MCP App token in the agent job, then
// emits MCP lockdown detection, guard variable parsing, and the DIFC proxy stop step.
func (c *Compiler) emitGitHubMCPTokenAndLockdownSteps(yaml *strings.Builder, data *WorkflowData) {
	// Mint the GitHub MCP App token directly in the agent job.
	// The token cannot be passed via job outputs from the activation job because
	// actions/create-github-app-token calls ::add-mask:: on the token, and the
	// GitHub Actions runner silently drops masked values in job outputs (runner v2.308+).
	for _, step := range c.generateGitHubMCPAppTokenMintingSteps(data) {
		yaml.WriteString(step)
	}
	c.generateGitHubMCPLockdownDetectionStep(yaml, data)
	c.generateParseGuardVarsStep(yaml, data)
	// Stop DIFC proxy before starting the MCP gateway to avoid double-filtering.
	c.generateStopDIFCProxyStep(yaml, data)
}

// emitActivationArtifactAndCommentMemorySteps downloads the activation artifact from the
// activation job and, when comment-memory is configured, prepares the comment-memory files.
func (c *Compiler) emitActivationArtifactAndCommentMemorySteps(yaml *strings.Builder, data *WorkflowData) {
	// Download activation artifact (contains aw_info.json and prompt.txt).
	// Must happen BEFORE pre-agent-steps so the base-branch snapshot is available.
	compilerYamlLog.Print("Adding activation artifact download step")
	activationArtifactName := artifactPrefixExprForDownstreamJob(data) + constants.ActivationArtifactName
	yaml.WriteString("      - name: Download activation artifact\n")
	fmt.Fprintf(yaml, "        uses: %s\n", c.getActionPin("actions/download-artifact"))
	yaml.WriteString("        with:\n")
	fmt.Fprintf(yaml, "          name: %s\n", activationArtifactName)
	yaml.WriteString("          path: /tmp/gh-aw\n")
	if data.SafeOutputs != nil && data.SafeOutputs.CommentMemory != nil {
		yaml.WriteString("      - name: Prepare comment memory files\n")
		fmt.Fprintf(yaml, "        uses: %s\n", getCachedActionPin("actions/github-script", data))
		yaml.WriteString("        with:\n")
		fmt.Fprintf(yaml, "          github-token: %s\n", getEffectiveSafeOutputGitHubToken(data.SafeOutputs.CommentMemory.GitHubToken))
		yaml.WriteString("          script: |\n")
		yaml.WriteString("            const { setupGlobals } = require('${{ runner.temp }}/gh-aw/actions/setup_globals.cjs');\n")
		yaml.WriteString("            setupGlobals(core, github, context, exec, io, getOctokit);\n")
		yaml.WriteString("            const { main } = require('${{ runner.temp }}/gh-aw/actions/setup_comment_memory_files.cjs');\n")
		yaml.WriteString("            await main();\n")
	}
}

// emitBaseGitHubRestoreAndInlineSubAgentsSteps restores trusted base-.github folders from the
// activation artifact (overwriting any PR-branch-injected files) and restores inline sub-agents.
func (c *Compiler) emitBaseGitHubRestoreAndInlineSubAgentsSteps(yaml *strings.Builder, data *WorkflowData) {
	// Restore agent config folders from the base branch snapshot in the activation artifact.
	// The activation job saved these before the PR checkout ran, so this step overwrites any
	// PR-branch-injected files (e.g. forked skill/instruction files) with trusted base content.
	// IMPORTANT: This must run BEFORE pre-agent-steps so that APM-restored skills
	// placed in .github/skills/ by pre-agent-steps are not clobbered by this restore.
	if ShouldGeneratePRCheckoutStep(data) {
		registry := GetGlobalEngineRegistry()
		generateRestoreBaseGitHubFoldersStep(yaml,
			registry.GetAllAgentManifestFolders(),
			registry.GetAllAgentManifestFiles(),
		)
	}
	// Restore inline sub-agents written during the activation job.
	// This step runs AFTER the base-branch restore so the engine-specific agent directory
	// is not clobbered. Inline sub-agents are enabled by default.
	if isFeatureEnabled(constants.FeatureFlag("inline-agents"), data) {
		generateRestoreInlineSubAgentsStep(yaml, data)
		generateRestoreInlineSkillsStep(yaml, data)
	}
}

// emitPreAgentAndMCPSetup emits pre-agent steps, MCP gateway setup, and MCP CLI mount.
func (c *Compiler) emitPreAgentAndMCPSetup(yaml *strings.Builder, data *WorkflowData, engine CodingAgentEngine) error {
	// Add pre-agent-steps (if any) after base-branch restore but before MCP setup.
	// Running after base restore ensures APM-restored skills (.github/skills/) are not
	// overwritten by the restore step above in PR context.
	// Running before MCP setup ensures pre-agent-steps can install/configure MCP
	// dependencies that the gateway may reference when it starts.
	c.generatePreAgentSteps(yaml, data)
	if err := c.generateMCPSetup(yaml, data.Tools, engine, data); err != nil {
		return fmt.Errorf("failed to generate MCP setup: %w", err)
	}
	// Mount MCP servers as CLI tools (runs after gateway is started)
	c.generateMCPCLIMountStep(yaml, data)
	return nil
}

// generateAgentRunSteps emits the git credentials cleaner, engine config steps, CLI proxy start,
// AI execution, CLI proxy stop, Copilot error detection, agent-execution-complete marker,
// post-agent git credential regeneration, firewall log collection, engine pre-bundle steps,
// MCP gateway stop, secret redaction, agent step summary append, and output collection.
// It returns the initial set of artifact paths (to be extended by the caller) and the
// agent stdio log path constant.
func (c *Compiler) generateAgentRunSteps(yaml *strings.Builder, data *WorkflowData, engine CodingAgentEngine, needsGitConfig bool) ([]string, string, error) {
	var artifactPaths []string
	artifactPaths = append(artifactPaths, "/tmp/gh-aw/aw-prompts/prompt.txt")
	logFileFull := "/tmp/gh-aw/agent-stdio.log"

	// Clean credentials before executing the agentic engine.
	for _, line := range c.generateCredentialsCleanerStep(data.KnownActionCredentialEnvVars) {
		yaml.WriteString(line)
	}
	// Emit an audit step after credentials have been cleaned but before the agent begins execution.
	c.generatePreAgentAuditStep(yaml)

	// Emit engine config steps (from RenderConfig) before the AI execution step.
	if err := c.emitEngineConfigSteps(yaml, data, engine); err != nil {
		return nil, "", err
	}

	// Start CLI proxy, run the engine, stop CLI proxy, detect errors, and mark completion.
	c.generateStartCliProxyStep(yaml, data)
	compilerYamlLog.Printf("Generating engine execution steps for %s", engine.GetID())
	c.generateEngineExecutionSteps(yaml, data, engine, logFileFull)
	c.generateStopCliProxyStep(yaml, data)
	c.generateDetectAgentErrorsStep(yaml, data, engine)
	compilerYamlLog.Print("Marking agent execution as complete for step order tracking")
	c.stepOrderTracker.MarkAgentExecutionComplete()

	// Regenerate git credentials and collect firewall logs + pre-bundle steps.
	c.emitPostAgentGitAndFirewallSteps(yaml, data, engine, needsGitConfig)

	// Stop MCP gateway, redact secrets, append step summary, and collect safe outputs.
	if err := c.emitRedactionAndSummarySteps(yaml, data); err != nil {
		return nil, "", err
	}

	return artifactPaths, logFileFull, nil
}

// emitEngineConfigSteps emits the engine config steps (from RenderConfig) that write runtime
// config files to disk before the AI execution step.
func (c *Compiler) emitEngineConfigSteps(yaml *strings.Builder, data *WorkflowData, engine CodingAgentEngine) error {
	if len(data.EngineConfigSteps) == 0 {
		return nil
	}
	compilerYamlLog.Printf("Adding %d engine config steps for %s", len(data.EngineConfigSteps), engine.GetID())
	for _, step := range data.EngineConfigSteps {
		stepYAML, err := ConvertStepToYAML(step)
		if err != nil {
			return fmt.Errorf("failed to render engine config step: %w", err)
		}
		yaml.WriteString(stepYAML)
	}
	return nil
}

// emitPostAgentGitAndFirewallSteps regenerates git credentials after agent execution and
// collects firewall logs and engine pre-bundle steps before secret redaction.
func (c *Compiler) emitPostAgentGitAndFirewallSteps(yaml *strings.Builder, data *WorkflowData, engine CodingAgentEngine, needsGitConfig bool) {
	// Regenerate git credentials after agent execution.
	// Only emit these steps when a checkout was performed (requires a .git directory).
	if needsGitConfig {
		for _, line := range c.generateGitConfigurationSteps() {
			yaml.WriteString(line)
		}
	}
	// Collect firewall logs BEFORE secret redaction so secrets in logs can be redacted.
	for _, step := range engine.GetFirewallLogsCollectionStep(data) {
		for _, line := range step {
			yaml.WriteString(line)
			yaml.WriteByte('\n')
		}
	}
	// Run engine pre-bundle steps to relocate files before secret redaction.
	for _, step := range engine.GetPreBundleSteps(data) {
		for _, line := range step {
			yaml.WriteString(line)
			yaml.WriteByte('\n')
		}
	}
}

// emitRedactionAndSummarySteps stops the MCP gateway, runs secret redaction, appends the
// agent step summary, and runs the safe-output collection step.
func (c *Compiler) emitRedactionAndSummarySteps(yaml *strings.Builder, data *WorkflowData) error {
	// Stop MCP gateway after agent execution and before secret redaction.
	c.generateStopMCPGateway(yaml, data)
	// Add secret redaction step BEFORE any artifact uploads.
	c.generateSecretRedactionStep(yaml, yaml.String(), data)
	// Append the agent step summary to the real $GITHUB_STEP_SUMMARY after secrets are redacted.
	c.generateAgentStepSummaryAppend(yaml)
	// Add output collection step only if safe-outputs feature is used.
	if data.SafeOutputs != nil {
		if err := c.generateOutputCollectionStep(yaml, data); err != nil {
			return err
		}
	}
	return nil
}

// collectArtifactPaths gathers all paths for the unified artifact upload.
// It starts from the initial paths already accumulated by generateAgentRunSteps and appends
// engine-declared output paths, log directories, observability files, safe-outputs files,
// patch/bundle paths, and firewall audit paths.
func (c *Compiler) collectArtifactPaths(data *WorkflowData, engine CodingAgentEngine, logFileFull string, initialPaths []string) []string {
	paths := initialPaths

	// Merge engine-declared output files into the unified artifact.
	paths = append(paths, getEngineArtifactPaths(engine)...)

	// Collect MCP logs.
	paths = append(paths, "/tmp/gh-aw/mcp-logs/")

	// Collect DIFC proxy logs (proxy-tls certs + container stderr) when proxy was injected
	paths = append(paths, difcProxyLogPaths(data)...)

	// Collect MCPScripts logs path if mcp-scripts is enabled
	if IsMCPScriptsEnabled(data.MCPScripts) {
		paths = append(paths, "/tmp/gh-aw/mcp-scripts/logs/")
	}

	// Include the aggregated agent_usage.json (requires AWF v0.25.8+)
	if isFirewallEnabled(data) {
		paths = append(paths, "/tmp/gh-aw/"+constants.TokenUsageFilename)
	}

	// Collect agent stdio logs path for unified upload
	paths = append(paths, logFileFull)

	// Include the pre-agent audit file
	paths = append(paths, constants.PreAgentAuditFilePath)

	// Collect agent-generated files path for unified upload
	paths = append(paths, "/tmp/gh-aw/agent/")

	return appendObservabilityAndOutputPaths(paths, data)
}

// appendObservabilityAndOutputPaths appends rate-limit logs, OTLP spans, safe-output files,
// patch/bundle globs, and firewall audit paths to the artifact paths slice.
func appendObservabilityAndOutputPaths(paths []string, data *WorkflowData) []string {
	// Collect GitHub API rate-limit log for observability.
	paths = append(paths, "/tmp/gh-aw/"+constants.GithubRateLimitsFilename)

	// Collect OTLP span mirror — enables post-hoc trace debugging without a live collector.
	if isOTLPEnabled(data) {
		paths = append(paths, "/tmp/gh-aw/"+constants.OtelJsonlFilename)
		paths = append(paths, "/tmp/gh-aw/"+constants.OtlpExportErrorsFilename)
	}

	// Collect safe outputs and agent output paths for the unified artifact.
	if data.SafeOutputs != nil {
		paths = append(paths, "/tmp/gh-aw/"+constants.SafeOutputsFilename)
		paths = append(paths, "/tmp/gh-aw/"+constants.AgentOutputFilename)
		if data.SafeOutputs.CommentMemory != nil {
			paths = append(paths, "/tmp/gh-aw/comment-memory/")
		}
	}

	// Collect git patch/bundle paths when safe-outputs PR operations or threat detection are configured.
	threatDetectionNeedsPatches := IsDetectionJobEnabled(data.SafeOutputs)
	if usesPatchesAndCheckouts(data.SafeOutputs) || threatDetectionNeedsPatches {
		paths = append(paths, "/tmp/gh-aw/aw-*.patch")
		paths = append(paths, "/tmp/gh-aw/aw-*.bundle")
	}

	// Include firewall audit/observability logs in the unified agent artifact (AWF v0.25.0+).
	if isFirewallEnabled(data) {
		paths = append(paths, constants.AWFConfigFilePath)
		paths = append(paths, constants.AWFProxyLogsDir+"/")
		paths = append(paths, constants.AWFAuditDir+"/")
		paths = append(paths, constants.AWFReflectFilePath)
	}

	return paths
}

// generateSummarySteps emits all GITHUB_STEP_SUMMARY log-parsing steps for the agent job.
// It covers agent log parsing, MCP scripts, MCP gateway, firewall logs, token usage,
// AWF reflect summary, and observability summary.
func (c *Compiler) generateSummarySteps(yaml *strings.Builder, data *WorkflowData, engine CodingAgentEngine) {
	// Parse agent logs for GITHUB_STEP_SUMMARY
	c.generateLogParsing(yaml, data, engine)

	// Parse mcp-scripts logs for GITHUB_STEP_SUMMARY (if mcp-scripts is enabled)
	if IsMCPScriptsEnabled(data.MCPScripts) {
		c.generateMCPScriptsLogParsing(yaml, data)
	}

	// Parse MCP gateway logs for GITHUB_STEP_SUMMARY.
	// The MCP gateway is always enabled, even when agent sandbox is disabled.
	c.generateMCPGatewayLogParsing(yaml, data)

	// Add firewall log parsing for all firewall-enabled engines.
	// This replaces the previous per-engine blocks (Copilot, Codex, Claude) and extends
	// support to all engines (including Gemini) so every agentic workflow uploads audit logs.
	if isFirewallEnabled(data) {
		firewallLogParsing := generateFirewallLogParsingStep(data.Name)
		for _, line := range firewallLogParsing {
			yaml.WriteString(line)
			yaml.WriteByte('\n')
		}
	}

	// Parse token-usage.jsonl and append to step summary (requires AWF v0.25.8+)
	if isFirewallEnabled(data) {
		c.generateTokenUsageSummary(yaml, data)
	}

	// Append AWF API proxy reflection data (available endpoints and models) to step summary.
	// This data is fetched from the /reflect endpoint by copilot_harness.cjs before the
	// agent exits and persisted to /tmp/gh-aw/awf-reflect.json.
	if isFirewallEnabled(data) {
		c.generateAWFReflectSummary(yaml, data)
	}

	// Synthesize a compact observability section from runtime artifacts when OTLP is enabled.
	c.generateObservabilitySummary(yaml, data)
}

// generatePostAgentCollectionAndUpload orchestrates the post-agent phase:
// engine output cleanup, access log collection, artifact path accumulation via collectArtifactPaths,
// step-summary generation via generateSummarySteps, safe-outputs/memory/staging artifact uploads,
// post-steps, the unified artifact upload, token invalidation, dev-mode actions restore,
// and step-order validation.
func (c *Compiler) generatePostAgentCollectionAndUpload(yaml *strings.Builder, data *WorkflowData, engine CodingAgentEngine, artifactPaths []string, logFileFull string, checkoutMgr *CheckoutManager) error {
	// Generate engine output cleanup step so workspace files are removed after collection.
	if len(getEngineArtifactPaths(engine)) > 0 {
		c.generateEngineOutputCleanup(yaml, engine)
	}

	// Extract and upload squid access logs (if any proxy tools were used)
	c.generateExtractAccessLogs(yaml, data.Tools)
	c.generateUploadAccessLogs(yaml, data.Tools)

	// Collect all artifact paths for the unified upload.
	artifactPaths = c.collectArtifactPaths(data, engine, logFileFull, artifactPaths)

	// Emit all GITHUB_STEP_SUMMARY log-parsing steps.
	c.generateSummarySteps(yaml, data, engine)

	// Write a minimal agent_output.json placeholder when the engine fails before producing safe outputs.
	if data.SafeOutputs != nil {
		c.generateAgentOutputPlaceholderStep(yaml)
	}

	c.emitCopilotCleanupMemoryAndUnifiedUpload(yaml, data, engine, artifactPaths, artifactPrefixExprForDownstreamJob(data))

	// In dev mode restore the actions/setup directory if an external root checkout replaced it.
	// We add a restore checkout step (if: always()) as the final step so the post-step
	// can always find action.yml and complete its /tmp/gh-aw cleanup.
	if c.actionMode.IsDev() && checkoutMgr.HasExternalRootCheckout() {
		yaml.WriteString(c.generateRestoreActionsSetupStep())
		compilerYamlLog.Print("Added restore actions folder step to agent job (dev mode with external root checkout)")
	}

	// Validate step ordering - this is a compiler check to ensure security
	if err := c.stepOrderTracker.ValidateStepOrdering(); err != nil {
		return fmt.Errorf("step ordering validation failed: %w", err)
	}
	return nil
}

// emitCopilotCleanupMemoryAndUnifiedUpload emits the Copilot engine cleanup step, all
// memory/staging artifact uploads, post-steps, and the unified artifact upload.
func (c *Compiler) emitCopilotCleanupMemoryAndUnifiedUpload(yaml *strings.Builder, data *WorkflowData, engine CodingAgentEngine, artifactPaths []string, agentArtifactPrefix string) {
	// Add post-execution cleanup step for Copilot engine
	if copilotEngine, ok := engine.(*CopilotEngine); ok {
		for _, line := range copilotEngine.GetCleanupStep(data) {
			yaml.WriteString(line)
			yaml.WriteByte('\n')
		}
	}

	generateRepoMemoryArtifactUpload(yaml, data, c.getActionPin)
	generateCacheMemoryGitCommitSteps(yaml, data)
	generateCacheMemoryValidation(yaml, data)
	generateCacheMemoryArtifactUpload(yaml, data, c.getActionPin)
	generateSafeOutputsAssetsArtifactUpload(yaml, data, c.getActionPin)
	generateSafeOutputsArtifactStagingUpload(yaml, data, c.getActionPin)

	c.generatePostSteps(yaml, data)
	c.generateUnifiedArtifactUpload(yaml, artifactPaths, agentArtifactPrefix)
}

// addCustomStepsAsIs adds custom steps after sanitizing any GitHub Actions expressions
// found directly in run: fields.  Any ${{ ... }} expression in a run: script is
// extracted into an env: variable to prevent shell injection attacks; a compiler
// warning is emitted for every such extraction.
func (c *Compiler) addCustomStepsAsIs(yaml *strings.Builder, customSteps string) {
	customSteps = c.sanitizeAndWarnCustomSteps(customSteps)
	// Remove "steps:" line and adjust indentation
	lines := strings.Split(customSteps, "\n")
	if len(lines) > 1 {
		for _, line := range lines[1:] {
			// Skip empty lines
			if strings.TrimSpace(line) == "" {
				yaml.WriteString("\n")
				continue
			}

			// Simply add 6 spaces for job context indentation
			yaml.WriteString("      " + line + "\n")
		}
	}
}

// addCustomStepsWithRuntimeInsertion adds custom steps and inserts runtime steps after the first checkout.
// Like addCustomStepsAsIs it sanitizes any ${{ ... }} expressions found in run: fields before writing.
func (c *Compiler) addCustomStepsWithRuntimeInsertion(yaml *strings.Builder, customSteps string, runtimeSetupSteps []GitHubActionStep, tools *ToolsConfig) {
	customSteps = c.sanitizeAndWarnCustomSteps(customSteps)
	// Remove "steps:" line and adjust indentation
	lines := strings.Split(customSteps, "\n")
	if len(lines) <= 1 {
		return
	}

	insertedRuntime := false
	i := 1 // Start from index 1 to skip "steps:" line

	for i < len(lines) {
		line := lines[i]

		// Skip empty lines
		if strings.TrimSpace(line) == "" {
			yaml.WriteString("\n")
			i++
			continue
		}

		// Add the line with proper indentation
		yaml.WriteString("      " + line + "\n")

		// If runtime hasn't been inserted yet, try to detect and handle a checkout step.
		if !insertedRuntime {
			newI, inserted := processCheckoutStepForRuntimeInsertion(yaml, lines, i, runtimeSetupSteps)
			if inserted {
				insertedRuntime = true
				i = newI
				continue
			}
		}

		i++
	}
}

// processCheckoutStepForRuntimeInsertion checks whether the line at index i is a step header
// whose body contains a "uses: …checkout…" directive. If it is, the function copies the
// remaining step lines into yaml, injects runtimeSetupSteps immediately afterwards, and
// returns (newI, true) where newI is the index of the first line of the *next* step. When the
// current line is not a checkout step the function returns (i, false) and writes nothing.
func processCheckoutStepForRuntimeInsertion(yaml *strings.Builder, lines []string, i int, runtimeSetupSteps []GitHubActionStep) (newI int, inserted bool) {
	trimmed := strings.TrimSpace(lines[i])
	isStepStart := strings.HasPrefix(trimmed, "- name:") || strings.HasPrefix(trimmed, "- uses:")
	if !isStepStart {
		return i, false
	}

	// Look ahead to find "uses:" line with "checkout"
	isCheckoutStep := false
	for j := i + 1; j < len(lines); j++ {
		nextTrimmed := strings.TrimSpace(lines[j])
		// Stop if we hit the next step
		if strings.HasPrefix(nextTrimmed, "- name:") || strings.HasPrefix(nextTrimmed, "- uses:") {
			break
		}
		if strings.Contains(nextTrimmed, "uses:") && strings.Contains(nextTrimmed, "checkout") {
			isCheckoutStep = true
			break
		}
	}

	if !isCheckoutStep {
		return i, false
	}

	// Copy all step lines until the next step header.
	i++
	for i < len(lines) {
		nextLine := lines[i]
		nextTrimmed := strings.TrimSpace(nextLine)
		if strings.HasPrefix(nextTrimmed, "- name:") || strings.HasPrefix(nextTrimmed, "- uses:") {
			break
		}
		if nextTrimmed == "" {
			yaml.WriteString("\n")
		} else {
			yaml.WriteString("      " + nextLine + "\n")
		}
		i++
	}

	// Insert runtime steps after the checkout step.
	compilerYamlLog.Printf("Inserting %d runtime setup steps after checkout in custom steps", len(runtimeSetupSteps))
	for _, step := range runtimeSetupSteps {
		for _, stepLine := range step {
			yaml.WriteString(stepLine + "\n")
		}
	}

	return i, true
}

// generateRepositoryImportCheckouts generates checkout steps for repository imports
// Each repository is checked out into a temporary folder at .github/aw/imports/<owner>-<repo>-<sanitized-ref>
// relative to GITHUB_WORKSPACE. This allows the merge script to copy files from pre-checked-out folders instead of doing git operations
func (c *Compiler) generateRepositoryImportCheckouts(yaml *strings.Builder, repositoryImports []string) {
	for _, repoImport := range repositoryImports {
		compilerYamlLog.Printf("Generating checkout step for repository import: %s", repoImport)

		// Parse the import spec to extract owner, repo, and ref
		// Format: owner/repo@ref or owner/repo
		owner, repo, ref := parseRepositoryImportSpec(repoImport)
		if owner == "" || repo == "" {
			compilerYamlLog.Printf("Warning: failed to parse repository import: %s", repoImport)
			continue
		}

		// Generate a sanitized directory name for the checkout
		// Use a consistent format: owner-repo-ref
		// NOTE: Path must be relative to GITHUB_WORKSPACE for actions/checkout@v6
		sanitizedRef := sanitizeRefForPath(ref)
		checkoutPath := fmt.Sprintf(".github/aw/imports/%s-%s-%s", owner, repo, sanitizedRef)

		// Generate the checkout step
		fmt.Fprintf(yaml, "      - name: Checkout repository import %s/%s@%s\n", owner, repo, ref)
		fmt.Fprintf(yaml, "        uses: %s\n", getActionPin("actions/checkout"))
		yaml.WriteString("        with:\n")
		fmt.Fprintf(yaml, "          repository: %s/%s\n", owner, repo)
		fmt.Fprintf(yaml, "          ref: %s\n", ref)
		fmt.Fprintf(yaml, "          path: %s\n", checkoutPath)
		yaml.WriteString("          sparse-checkout: |\n")
		yaml.WriteString("            .github/\n")
		yaml.WriteString("          persist-credentials: false\n")

		compilerYamlLog.Printf("Added checkout step: %s/%s@%s -> %s", owner, repo, ref, checkoutPath)
	}
}

// parseRepositoryImportSpec parses a repository import specification
// Format: owner/repo@ref or owner/repo (defaults to "main" if no ref)
// Returns: owner, repo, ref
func parseRepositoryImportSpec(importSpec string) (owner, repo, ref string) {
	// Remove section reference if present (file.md#Section)
	cleanSpec := importSpec
	if before, _, ok := strings.Cut(importSpec, "#"); ok {
		cleanSpec = before
	}

	// Split on @ to get path and ref
	parts := strings.Split(cleanSpec, "@")
	pathPart := parts[0]
	ref = "main" // default ref
	if len(parts) > 1 {
		ref = parts[1]
	}

	// Parse path: owner/repo
	slashParts := strings.Split(pathPart, "/")
	if len(slashParts) != 2 {
		return "", "", ""
	}

	owner = slashParts[0]
	repo = slashParts[1]

	return owner, repo, ref
}

// generateLegacyAgentImportCheckout generates a checkout step for legacy agent imports
// Legacy format: owner/repo/path/to/file.md@ref
// This checks out the entire repository (not just .github folder) since the file could be anywhere
func (c *Compiler) generateLegacyAgentImportCheckout(yaml *strings.Builder, agentImportSpec string) {
	compilerYamlLog.Printf("Generating checkout step for legacy agent import: %s", agentImportSpec)

	// Parse the import spec to extract owner, repo, and ref
	owner, repo, ref := parseRepositoryImportSpec(agentImportSpec)
	if owner == "" || repo == "" {
		compilerYamlLog.Printf("Warning: failed to parse legacy agent import spec: %s", agentImportSpec)
		return
	}

	// Generate a sanitized directory name for the checkout
	sanitizedRef := sanitizeRefForPath(ref)
	checkoutPath := fmt.Sprintf("/tmp/gh-aw/repo-imports/%s-%s-%s", owner, repo, sanitizedRef)

	// Generate the checkout step
	fmt.Fprintf(yaml, "      - name: Checkout agent import %s/%s@%s\n", owner, repo, ref)
	fmt.Fprintf(yaml, "        uses: %s\n", getActionPin("actions/checkout"))
	yaml.WriteString("        with:\n")
	fmt.Fprintf(yaml, "          repository: %s/%s\n", owner, repo)
	fmt.Fprintf(yaml, "          ref: %s\n", ref)
	fmt.Fprintf(yaml, "          path: %s\n", checkoutPath)
	yaml.WriteString("          sparse-checkout: |\n")
	yaml.WriteString("            .github/\n")
	yaml.WriteString("          persist-credentials: false\n")

	compilerYamlLog.Printf("Added legacy agent checkout step: %s/%s@%s -> %s", owner, repo, ref, checkoutPath)
}

// generateDevModeCLIBuildSteps generates the steps needed to build the gh-aw CLI and Docker image in dev mode
// These steps are injected after checkout in dev mode to create a locally built Docker image that includes
// the gh-aw binary and all dependencies. The agentic-workflows MCP server uses this image instead of alpine:latest.
//
// The build process:
// 1. Setup Go using go.mod version
// 2. Build the gh-aw CLI binary for linux/amd64 (since it runs in a Linux container)
// 3. Setup Docker Buildx for advanced build features
// 4. Build Docker image and tag it as localhost/gh-aw:dev
//
// The built image is used by the agentic-workflows MCP server configuration (see mcp_config_builtin.go)
func (c *Compiler) generateDevModeCLIBuildSteps(yaml *strings.Builder) {
	compilerYamlLog.Print("Generating dev mode CLI build steps")

	// Step 1: Setup Go for building the CLI
	yaml.WriteString("      - name: Setup Go for CLI build\n")
	fmt.Fprintf(yaml, "        uses: %s\n", getActionPin("actions/setup-go"))
	yaml.WriteString("        with:\n")
	yaml.WriteString("          go-version-file: go.mod\n")
	yaml.WriteString("          cache: true\n")

	// Step 2: Build CLI binary for linux/amd64
	// Use the standard build command from CI/Makefile (not release build)
	// CGO_ENABLED=0 for static linking (required for Alpine containers)
	yaml.WriteString("      - name: Build gh-aw CLI\n")
	yaml.WriteString("        run: |\n")
	yaml.WriteString("          echo \"Building gh-aw CLI for linux/amd64...\"\n")
	yaml.WriteString("          mkdir -p dist\n")
	yaml.WriteString("          VERSION=$(git describe --tags --always --dirty)\n")
	yaml.WriteString("          CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \\\n")
	yaml.WriteString("            -ldflags \"-s -w -X main.version=${VERSION}\" \\\n")
	yaml.WriteString("            -o dist/gh-aw-linux-amd64 \\\n")
	yaml.WriteString("            ./cmd/gh-aw\n")
	yaml.WriteString("          # Copy binary to root for direct execution in user-defined steps\n")
	yaml.WriteString("          cp dist/gh-aw-linux-amd64 ./gh-aw\n")
	yaml.WriteString("          chmod +x ./gh-aw\n")
	yaml.WriteString("          echo \"✓ Built gh-aw CLI successfully\"\n")

	// Step 3: Setup Docker Buildx
	yaml.WriteString("      - name: Setup Docker Buildx\n")
	fmt.Fprintf(yaml, "        uses: %s\n", getActionPin("docker/setup-buildx-action"))

	// Step 4: Build Docker image
	// Use the Dockerfile at the repository root which expects BINARY build arg
	yaml.WriteString("      - name: Build gh-aw Docker image\n")
	fmt.Fprintf(yaml, "        uses: %s\n", getActionPin("docker/build-push-action"))
	yaml.WriteString("        with:\n")
	yaml.WriteString("          context: .\n")
	yaml.WriteString("          platforms: linux/amd64\n")
	yaml.WriteString("          push: false\n")
	yaml.WriteString("          load: true\n")
	yaml.WriteString("          tags: localhost/gh-aw:dev\n")
	yaml.WriteString("          build-args: |\n")
	yaml.WriteString("            BINARY=dist/gh-aw-linux-amd64\n")
}

// sanitizeAndWarnCustomSteps applies sanitizeCustomStepsYAML to the custom steps string,
// emits a compiler warning for every expression that was extracted, and returns the
// sanitized string.  If sanitization fails or produces no changes the original is returned.
func (c *Compiler) sanitizeAndWarnCustomSteps(customSteps string) string {
	sanitized, warnings, err := sanitizeCustomStepsYAML(customSteps)
	if err != nil {
		compilerYamlLog.Printf("Failed to sanitize custom steps YAML: %v", err)
		return customSteps
	}
	for _, w := range warnings {
		fmt.Fprintln(os.Stderr, console.FormatWarningMessage(w))
		c.IncrementWarningCount()
	}
	return sanitized
}
