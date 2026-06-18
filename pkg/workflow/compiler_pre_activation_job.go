package workflow

import (
	"encoding/json"
	"errors"
	"fmt"
	"maps"
	"strings"

	"github.com/github/gh-aw/pkg/constants"
	"github.com/github/gh-aw/pkg/logger"
	"github.com/github/gh-aw/pkg/sliceutil"
	"github.com/github/gh-aw/pkg/stringutil"
)

var compilerActivationJobsLog = logger.New("workflow:compiler_activation_jobs")

// buildPreActivationJob creates a unified pre-activation job that combines membership checks and stop-time validation.
// This job exposes a single "activated" output that indicates whether the workflow should proceed.
func (c *Compiler) buildPreActivationJob(data *WorkflowData, needsPermissionCheck bool) (*Job, error) {
	compilerActivationJobsLog.Printf("Building pre-activation job: needsPermissionCheck=%v, hasStopTime=%v", needsPermissionCheck, data.StopTime != "")

	// Extract custom steps and outputs from jobs.pre-activation if present
	customSteps, customOutputs, err := c.extractPreActivationCustomFields(data.Jobs)
	if err != nil {
		return nil, fmt.Errorf("failed to extract pre-activation custom fields: %w", err)
	}

	// Add setup action steps at the beginning of the job
	setupActionRef := c.resolveActionReference("./actions/setup", data)
	if setupActionRef == "" {
		return nil, errors.New("setup action reference is required but could not be resolved")
	}

	steps, needsContentsRead := c.buildPreActivationInitialSteps(data, setupActionRef, needsPermissionCheck)
	permissions := buildPreActivationPermissions(data, needsContentsRead)

	// Add optional check steps (stop-time, skip-if, skip-roles/bots, command position)
	steps = c.appendPreActivationOptionalCheckSteps(data, steps)

	// Append custom steps and on.steps, collecting step IDs for output wiring
	var onStepIDs []string
	steps, onStepIDs, err = appendPreActivationCustomAndOnSteps(data, steps, customSteps)
	if err != nil {
		return nil, err
	}

	// Build the activated output expression from all configured checks
	activatedExpression, err := buildPreActivationActivatedExpression(data, needsPermissionCheck)
	if err != nil {
		return nil, err
	}

	// Build outputs map and job-level if condition
	outputs := buildPreActivationOutputsMap(data, onStepIDs, customOutputs, activatedExpression)
	jobIfCondition := c.buildPreActivationJobIfCondition(data, needsPermissionCheck)

	// In script mode, explicitly add a cleanup step (mirrors post.js in dev/release/action mode).
	if c.actionMode.IsScript() {
		steps = append(steps, c.generateScriptModeCleanupStep())
	}

	return &Job{
		Name:        string(constants.PreActivationJobName),
		If:          jobIfCondition,
		RunsOn:      c.formatFrameworkJobRunsOn(data),
		Environment: c.indentYAMLLines(resolveSafeOutputsEnvironment(data), "    "),
		Permissions: permissions,
		Steps:       steps,
		Outputs:     outputs,
		Needs:       sliceutil.Deduplicate(data.OnNeeds),
	}, nil
}

// buildPreActivationPermissions constructs the permissions string for the pre-activation job
// based on whether contents:read, actions:read, or on.permissions are needed.
func buildPreActivationPermissions(data *WorkflowData, needsContentsRead bool) string {
	var perms *Permissions
	if needsContentsRead {
		perms = NewPermissionsContentsRead()
	}
	if data.RateLimit != nil {
		if perms == nil {
			perms = NewPermissions()
		}
		perms.Set(PermissionActions, PermissionRead)
	}
	if data.OnPermissions != nil {
		if perms == nil {
			perms = NewPermissions()
		}
		perms.Merge(data.OnPermissions)
	}
	if perms != nil {
		return perms.RenderToYAML()
	}
	return ""
}

// buildPreActivationInitialSteps initializes the steps slice with checkout, setup, and optional membership/rate-limit steps.
// It returns the initial steps and whether contents:read permission is needed.
func (c *Compiler) buildPreActivationInitialSteps(data *WorkflowData, setupActionRef string, needsPermissionCheck bool) (steps []string, needsContentsRead bool) {
steps = append(steps, c.generateCheckoutActionsFolder(data)...)
needsContentsRead = (c.actionMode.IsDev() || c.actionMode.IsScript()) && len(c.generateCheckoutActionsFolder(data)) > 0
steps = append(steps, c.generateSetupStep(data, setupActionRef, SetupActionDestination, false, "", "")...)
if needsPermissionCheck {
steps = c.generateMembershipCheck(data, steps)
}
if data.RateLimit != nil {
steps = c.generateRateLimitCheck(data, steps)
}
return steps, needsContentsRead
}

// appendPreActivationOptionalCheckSteps appends all optional filter check steps to the steps slice.
// This includes stop-time, skip-if-match, skip-if-no-match, skip-if-check-failing,
// skip-roles, skip-bots, and command-position checks.
func (c *Compiler) appendPreActivationOptionalCheckSteps(data *WorkflowData, steps []string) []string {
	if data.StopTime != "" {
		compilerActivationJobsLog.Printf("Adding stop-time check step: stop_time=%s", data.StopTime)
		steps = append(steps, generatePreActivationStopTimeStep(data)...)
	}
	hasSkipIfCheck := data.SkipIfMatch != nil || data.SkipIfNoMatch != nil
	if hasSkipIfCheck && data.ActivationGitHubApp != nil {
		steps = append(steps, c.buildPreActivationAppTokenMintStep(data.ActivationGitHubApp)...)
	}
	skipIfToken := c.resolvePreActivationSkipIfToken(data)
	if data.SkipIfMatch != nil {
		compilerActivationJobsLog.Printf("Adding skip-if-match check step: query=%s, max=%d", data.SkipIfMatch.Query, data.SkipIfMatch.Max)
		steps = append(steps, generatePreActivationSkipIfMatchStep(data, skipIfToken)...)
	}
	if data.SkipIfNoMatch != nil {
		compilerActivationJobsLog.Printf("Adding skip-if-no-match check step: query=%s, min=%d", data.SkipIfNoMatch.Query, data.SkipIfNoMatch.Min)
		steps = append(steps, generatePreActivationSkipIfNoMatchStep(data, skipIfToken)...)
	}
	if data.SkipIfCheckFailing != nil {
		compilerActivationJobsLog.Printf("Adding skip-if-check-failing check step: include=%v, exclude=%v", data.SkipIfCheckFailing.Include, data.SkipIfCheckFailing.Exclude)
		steps = append(steps, generatePreActivationSkipIfCheckFailingStep(data)...)
	}
	if len(data.SkipRoles) > 0 {
		steps = append(steps, generatePreActivationSkipRolesStep(data)...)
	}
	if len(data.SkipBots) > 0 {
		steps = append(steps, generatePreActivationSkipBotsStep(data)...)
	}
	if len(data.Command) > 0 {
		steps = append(steps, generatePreActivationCommandPositionStep(data)...)
	}
	return steps
}

// generatePreActivationStopTimeStep returns the step YAML lines for the stop-time check.
func generatePreActivationStopTimeStep(data *WorkflowData) []string {
	cleanStopTime := stringutil.StripANSI(data.StopTime)
	return []string{
		"      - name: Check stop-time limit\n",
		fmt.Sprintf("        id: %s\n", constants.CheckStopTimeStepID),
		fmt.Sprintf("        uses: %s\n", getCachedActionPin("actions/github-script", data)),
		"        env:\n",
		fmt.Sprintf("          GH_AW_STOP_TIME: %q\n", cleanStopTime),
		fmt.Sprintf("          GH_AW_WORKFLOW_NAME: %q\n", data.Name),
		"        with:\n",
		"          script: |\n",
		generateGitHubScriptWithRequire("check_stop_time.cjs"),
	}
}

// generatePreActivationSkipIfMatchStep returns the step YAML lines for the skip-if-match check.
func generatePreActivationSkipIfMatchStep(data *WorkflowData, skipIfToken string) []string {
	steps := []string{
		"      - name: Check skip-if-match query\n",
		fmt.Sprintf("        id: %s\n", constants.CheckSkipIfMatchStepID),
		fmt.Sprintf("        uses: %s\n", getCachedActionPin("actions/github-script", data)),
		"        env:\n",
		fmt.Sprintf("          GH_AW_SKIP_QUERY: %q\n", data.SkipIfMatch.Query),
		fmt.Sprintf("          GH_AW_WORKFLOW_NAME: %q\n", data.Name),
		fmt.Sprintf("          GH_AW_SKIP_MAX_MATCHES: \"%d\"\n", data.SkipIfMatch.Max),
	}
	if data.SkipIfMatch.Scope != "" {
		steps = append(steps, fmt.Sprintf("          GH_AW_SKIP_SCOPE: %q\n", data.SkipIfMatch.Scope))
	}
	steps = append(steps, "        with:\n")
	if skipIfToken != "" {
		steps = append(steps, fmt.Sprintf("          github-token: %s\n", skipIfToken))
	}
	return append(steps, "          script: |\n", generateGitHubScriptWithRequire("check_skip_if_match.cjs"))
}

// generatePreActivationSkipIfNoMatchStep returns the step YAML lines for the skip-if-no-match check.
func generatePreActivationSkipIfNoMatchStep(data *WorkflowData, skipIfToken string) []string {
	steps := []string{
		"      - name: Check skip-if-no-match query\n",
		fmt.Sprintf("        id: %s\n", constants.CheckSkipIfNoMatchStepID),
		fmt.Sprintf("        uses: %s\n", getCachedActionPin("actions/github-script", data)),
		"        env:\n",
		fmt.Sprintf("          GH_AW_SKIP_QUERY: %q\n", data.SkipIfNoMatch.Query),
		fmt.Sprintf("          GH_AW_WORKFLOW_NAME: %q\n", data.Name),
		fmt.Sprintf("          GH_AW_SKIP_MIN_MATCHES: \"%d\"\n", data.SkipIfNoMatch.Min),
	}
	if data.SkipIfNoMatch.Scope != "" {
		steps = append(steps, fmt.Sprintf("          GH_AW_SKIP_SCOPE: %q\n", data.SkipIfNoMatch.Scope))
	}
	steps = append(steps, "        with:\n")
	if skipIfToken != "" {
		steps = append(steps, fmt.Sprintf("          github-token: %s\n", skipIfToken))
	}
	return append(steps, "          script: |\n", generateGitHubScriptWithRequire("check_skip_if_no_match.cjs"))
}

// generatePreActivationSkipIfCheckFailingStep returns the step YAML lines for the skip-if-check-failing check.
func generatePreActivationSkipIfCheckFailingStep(data *WorkflowData) []string {
	cfg := data.SkipIfCheckFailing
	steps := []string{
		"      - name: Check skip-if-check-failing\n",
		fmt.Sprintf("        id: %s\n", constants.CheckSkipIfCheckFailingStepID),
		fmt.Sprintf("        uses: %s\n", getCachedActionPin("actions/github-script", data)),
	}
	if len(cfg.Include) > 0 || len(cfg.Exclude) > 0 || cfg.Branch != "" || cfg.AllowPending {
		steps = append(steps, "        env:\n")
		if len(cfg.Include) > 0 {
			includeJSON, _ := json.Marshal(cfg.Include) //nolint:jsonmarshalignoredeerror // marshaling a string slice cannot fail
			steps = append(steps, fmt.Sprintf("          GH_AW_SKIP_CHECK_INCLUDE: %q\n", string(includeJSON)))
		}
		if len(cfg.Exclude) > 0 {
			excludeJSON, _ := json.Marshal(cfg.Exclude) //nolint:jsonmarshalignoredeerror // marshaling a string slice cannot fail
			steps = append(steps, fmt.Sprintf("          GH_AW_SKIP_CHECK_EXCLUDE: %q\n", string(excludeJSON)))
		}
		if cfg.Branch != "" {
			steps = append(steps, fmt.Sprintf("          GH_AW_SKIP_BRANCH: %q\n", cfg.Branch))
		}
		if cfg.AllowPending {
			steps = append(steps, "          GH_AW_SKIP_CHECK_ALLOW_PENDING: \"true\"\n")
		}
	}
	return append(steps, "        with:\n", "          script: |\n", generateGitHubScriptWithRequire("check_skip_if_check_failing.cjs"))
}

// generatePreActivationSkipRolesStep returns the step YAML lines for the skip-roles check.
func generatePreActivationSkipRolesStep(data *WorkflowData) []string {
	return []string{
		"      - name: Check skip-roles\n",
		fmt.Sprintf("        id: %s\n", constants.CheckSkipRolesStepID),
		fmt.Sprintf("        uses: %s\n", getCachedActionPin("actions/github-script", data)),
		"        env:\n",
		fmt.Sprintf("          GH_AW_SKIP_ROLES: %q\n", strings.Join(data.SkipRoles, ",")),
		fmt.Sprintf("          GH_AW_WORKFLOW_NAME: %q\n", data.Name),
		"        with:\n",
		"          github-token: ${{ secrets.GITHUB_TOKEN }}\n",
		"          script: |\n",
		generateGitHubScriptWithRequire("check_skip_roles.cjs"),
	}
}

// generatePreActivationSkipBotsStep returns the step YAML lines for the skip-bots check.
func generatePreActivationSkipBotsStep(data *WorkflowData) []string {
	steps := []string{
		"      - name: Check skip-bots\n",
		fmt.Sprintf("        id: %s\n", constants.CheckSkipBotsStepID),
		fmt.Sprintf("        uses: %s\n", getCachedActionPin("actions/github-script", data)),
		"        env:\n",
		fmt.Sprintf("          GH_AW_SKIP_BOTS: %q\n", strings.Join(data.SkipBots, ",")),
		fmt.Sprintf("          GH_AW_WORKFLOW_NAME: %q\n", data.Name),
	}
	if data.AllowBotAuthoredTriggerComment {
		steps = append(steps, "          GH_AW_ALLOW_BOT_AUTHORED_TRIGGER_COMMENT: \"true\"\n")
	}
	return append(steps, "        with:\n", "          script: |\n", generateGitHubScriptWithRequire("check_skip_bots.cjs"))
}

// generatePreActivationCommandPositionStep returns the step YAML lines for the command position check.
func generatePreActivationCommandPositionStep(data *WorkflowData) []string {
	commandsJSON, _ := json.Marshal(data.Command) //nolint:jsonmarshalignoredeerror // marshaling a string slice cannot fail
	steps := []string{
		"      - name: Check command position\n",
		fmt.Sprintf("        id: %s\n", constants.CheckCommandPositionStepID),
		fmt.Sprintf("        uses: %s\n", getCachedActionPin("actions/github-script", data)),
		"        env:\n",
		fmt.Sprintf("          GH_AW_COMMANDS: %q\n", string(commandsJSON)),
	}
	if data.CommandPlaceholder != "" {
		steps = append(steps, fmt.Sprintf("          GH_AW_COMMAND_PLACEHOLDER: %q\n", data.CommandPlaceholder))
	}
	return append(steps, "        with:\n", "          script: |\n", generateGitHubScriptWithRequire("check_command_position.cjs"))
}

// appendPreActivationCustomAndOnSteps appends custom steps from jobs.pre-activation and on.steps to the
// steps slice. It returns the updated steps and a list of on.steps step IDs for output wiring.
func appendPreActivationCustomAndOnSteps(data *WorkflowData, steps []string, customSteps []string) ([]string, []string, error) {
	if len(customSteps) > 0 {
		compilerActivationJobsLog.Printf("Adding %d custom steps to pre-activation job", len(customSteps))
		steps = append(steps, customSteps...)
	}
	var onStepIDs []string
	if len(data.OnSteps) > 0 {
		compilerActivationJobsLog.Printf("Adding %d on.steps to pre-activation job", len(data.OnSteps))
		for i, stepMap := range data.OnSteps {
			stepYAML, err := ConvertStepToYAML(stepMap)
			if err != nil {
				return nil, nil, fmt.Errorf("failed to convert on.steps[%d] to YAML: %w", i, err)
			}
			steps = append(steps, stepYAML)
			if id, ok := stepMap["id"].(string); ok && id != "" {
				onStepIDs = append(onStepIDs, id)
			}
		}
	}
	return steps, onStepIDs, nil
}

// buildPreActivationCheckConditions builds the list of ConditionNode values for all configured checks.
func buildPreActivationCheckConditions(data *WorkflowData, needsPermissionCheck bool) []ConditionNode {
	var conditions []ConditionNode
	if needsPermissionCheck {
		conditions = append(conditions, BuildComparison(
			BuildPropertyAccess(fmt.Sprintf("steps.%s.outputs.%s", constants.CheckMembershipStepID, constants.IsTeamMemberOutput)),
			"==", BuildStringLiteral("true"),
		))
	}
	if data.StopTime != "" {
		conditions = append(conditions, BuildComparison(
			BuildPropertyAccess(fmt.Sprintf("steps.%s.outputs.%s", constants.CheckStopTimeStepID, constants.StopTimeOkOutput)),
			"==", BuildStringLiteral("true"),
		))
	}
	if data.SkipIfMatch != nil {
		conditions = append(conditions, BuildComparison(
			BuildPropertyAccess(fmt.Sprintf("steps.%s.outputs.%s", constants.CheckSkipIfMatchStepID, constants.SkipCheckOkOutput)),
			"==", BuildStringLiteral("true"),
		))
	}
	if data.SkipIfNoMatch != nil {
		conditions = append(conditions, BuildComparison(
			BuildPropertyAccess(fmt.Sprintf("steps.%s.outputs.%s", constants.CheckSkipIfNoMatchStepID, constants.SkipNoMatchCheckOkOutput)),
			"==", BuildStringLiteral("true"),
		))
	}
	if data.SkipIfCheckFailing != nil {
		conditions = append(conditions, BuildComparison(
			BuildPropertyAccess(fmt.Sprintf("steps.%s.outputs.%s", constants.CheckSkipIfCheckFailingStepID, constants.SkipIfCheckFailingOkOutput)),
			"==", BuildStringLiteral("true"),
		))
	}
	if len(data.SkipRoles) > 0 {
		conditions = append(conditions, BuildComparison(
			BuildPropertyAccess(fmt.Sprintf("steps.%s.outputs.%s", constants.CheckSkipRolesStepID, constants.SkipRolesOkOutput)),
			"==", BuildStringLiteral("true"),
		))
	}
	if len(data.SkipBots) > 0 {
		conditions = append(conditions, BuildComparison(
			BuildPropertyAccess(fmt.Sprintf("steps.%s.outputs.%s", constants.CheckSkipBotsStepID, constants.SkipBotsOkOutput)),
			"==", BuildStringLiteral("true"),
		))
	}
	if data.RateLimit != nil {
		conditions = append(conditions, BuildComparison(
			BuildPropertyAccess(fmt.Sprintf("steps.%s.outputs.%s", constants.CheckRateLimitStepID, constants.RateLimitOkOutput)),
			"==", BuildStringLiteral("true"),
		))
	}
	if len(data.Command) > 0 {
		conditions = append(conditions, BuildComparison(
			BuildPropertyAccess(fmt.Sprintf("steps.%s.outputs.%s", constants.CheckCommandPositionStepID, constants.CommandPositionOkOutput)),
			"==", BuildStringLiteral("true"),
		))
	}
	return conditions
}

// buildPreActivationActivatedExpression builds the "${{ ... }}" expression string for the
// "activated" job output by combining all configured check conditions with AND.
func buildPreActivationActivatedExpression(data *WorkflowData, needsPermissionCheck bool) (string, error) {
	conditions := buildPreActivationCheckConditions(data, needsPermissionCheck)
	var activatedNode ConditionNode
	switch len(conditions) {
	case 0:
		if len(data.OnSteps) > 0 || len(data.OnNeeds) > 0 || len(data.SkipAuthorAssociations) > 0 {
			compilerActivationJobsLog.Printf(
				"Pre-activation created with no output checks (on.steps=%d, on.needs=%d, skip-author-associations=%d); activated output is unconditionally true",
				len(data.OnSteps), len(data.OnNeeds), len(data.SkipAuthorAssociations),
			)
			activatedNode = BuildStringLiteral("true")
		} else {
			return "", errors.New("developer error: pre-activation job created without permission check or stop-time configuration")
		}
	case 1:
		activatedNode = conditions[0]
	default:
		activatedNode = conditions[0]
		for i := 1; i < len(conditions); i++ {
			activatedNode = BuildAnd(activatedNode, conditions[i])
		}
	}
	return fmt.Sprintf("${{ %s }}", activatedNode.Render()), nil
}

// buildPreActivationOutputsMap builds the outputs map for the pre-activation job, including
// the activated expression, trace IDs, on.steps outcomes, and custom outputs.
func buildPreActivationOutputsMap(data *WorkflowData, onStepIDs []string, customOutputs map[string]string, activatedExpression string) map[string]string {
	outputs := map[string]string{
		"activated":            activatedExpression,
		"setup-trace-id":       "${{ steps.setup.outputs.trace-id }}",
		"setup-span-id":        "${{ steps.setup.outputs.span-id }}",
		"setup-parent-span-id": "${{ steps.setup.outputs.parent-span-id || steps.setup.outputs.span-id }}",
	}
	if len(data.Command) > 0 {
		outputs[constants.MatchedCommandOutput] = fmt.Sprintf("${{ steps.%s.outputs.%s }}", constants.CheckCommandPositionStepID, constants.MatchedCommandOutput)
	} else {
		outputs[constants.MatchedCommandOutput] = "''"
	}
	if len(onStepIDs) > 0 {
		compilerActivationJobsLog.Printf("Wiring %d on.steps step outcomes as pre-activation outputs", len(onStepIDs))
		for _, id := range onStepIDs {
			outputs[id+"_result"] = fmt.Sprintf("${{ steps.%s.outcome }}", id)
		}
	}
	if len(customOutputs) > 0 {
		compilerActivationJobsLog.Printf("Adding %d custom outputs to pre-activation job", len(customOutputs))
		maps.Copy(outputs, customOutputs)
	}
	return outputs
}

// buildPreActivationJobIfCondition constructs the job-level if: condition for the pre-activation job.
// It combines the base condition (data.If), label guards, comment-author guards, and
// skip-author-association guards as appropriate.
func (c *Compiler) buildPreActivationJobIfCondition(data *WorkflowData, needsPermissionCheck bool) string {
	// Base condition: pass through data.If unless it references custom job or pre-activation outputs
	var jobIfCondition string
	if !c.referencesCustomJobOutputs(data.If, data.Jobs) && !referencesPreActivationOutputs(data.If) {
		jobIfCondition = data.If
	}
	// When labels is specified, guard the job with a label-name condition
	if len(data.LabelNames) > 0 {
		labelIfCondition := buildLabelNamesCondition(data.LabelNames)
		if jobIfCondition != "" {
			jobIfCondition = RenderCondition(BuildAnd(
				&ExpressionNode{Expression: labelIfCondition},
				&ExpressionNode{Expression: jobIfCondition},
			))
		} else {
			jobIfCondition = labelIfCondition
		}
	}
	// For comment-triggered workflows requiring permission checks, add a static author-association guard
	if needsPermissionCheck && hasCommentEventInOn(data.On) && !botsContainExpression(data.Bots) && !strings.Contains(data.On, "${{") {
		commentAuthCondition := RenderCondition(buildCommentAuthorAssociationCondition(data.Bots))
		if jobIfCondition != "" {
			jobIfCondition = RenderCondition(BuildAnd(
				&ExpressionNode{Expression: commentAuthCondition},
				&ExpressionNode{Expression: jobIfCondition},
			))
		} else {
			jobIfCondition = commentAuthCondition
		}
	}
	// Add skip-author-associations guard to exit early for matching event/association combinations
	if len(data.SkipAuthorAssociations) > 0 {
		skipAuthorAssocCondition := RenderCondition(buildSkipAuthorAssociationsCondition(data.SkipAuthorAssociations))
		if jobIfCondition != "" {
			jobIfCondition = RenderCondition(BuildAnd(
				&ExpressionNode{Expression: skipAuthorAssocCondition},
				&ExpressionNode{Expression: jobIfCondition},
			))
		} else {
			jobIfCondition = skipAuthorAssocCondition
		}
	}
	return jobIfCondition
}


// buildLabelNamesCondition constructs the GitHub Actions if: expression for labels filtering.
// The generated condition passes when:
//   - the event has no label object (github.event.label == null), which covers
//     workflow_dispatch, push, schedule, and any other non-labeled events, OR
//   - the triggering label name matches any of the specified names.
//
// Using github.event.label == null (rather than checking the name) is semantically
// clearer and handles cases where GitHub Actions evaluates missing nested properties
// as null before coercing to empty string.
func buildLabelNamesCondition(labelNames []string) string {
	// Pass through events without a label payload.
	// github.event.label is null for workflow_dispatch, push, schedule, etc.
	noLabelEvent := ConditionNode(BuildEquals(
		BuildPropertyAccess("github.event.label"),
		BuildNullLiteral(),
	))

	result := noLabelEvent
	for _, name := range labelNames {
		result = BuildOr(result, BuildEquals(
			BuildPropertyAccess("github.event.label.name"),
			BuildStringLiteral(name),
		))
	}

	return result.Render()
}

// hasCommentEventInOn reports whether the rendered on: section includes issue_comment or
// pull_request_review_comment events. These are the events flagged by RGS-004 because
// any GitHub user (including unaffiliated outsiders) can post a comment and trigger the workflow.
// data.On is compiled YAML generated by the compiler, so checking for the event name followed by a
// colon (':') reliably identifies a trigger key without false-positives from embedded strings.
func hasCommentEventInOn(on string) bool {
	return strings.Contains(on, "issue_comment:") || strings.Contains(on, "pull_request_review_comment:")
}

// botsContainExpression reports whether any entry in bots is a GitHub Actions expression
// (i.e. contains "${{"). When true, the static author_association guard must be disabled so
// that check_membership always runs and evaluates the bot list at runtime.
func botsContainExpression(bots []string) bool {
	for _, bot := range bots {
		if strings.Contains(bot, "${{") {
			return true
		}
	}
	return false
}

// generateReportSkipStep generates the "Report skip reason" step for the pre-activation job.
// The step runs with if: always() and writes skip reasons to the GitHub Actions job summary
// extractPreActivationCustomFields extracts custom steps and outputs from jobs.pre-activation field in frontmatter.
// It validates that only steps and outputs fields are present, and errors on any other fields.
// If both jobs.pre-activation and jobs.pre_activation are defined, imports from both.
// Returns (customSteps, customOutputs, error).
func (c *Compiler) extractPreActivationCustomFields(jobs map[string]any) ([]string, map[string]string, error) {
	if jobs == nil {
		return nil, nil, nil
	}
	var customSteps []string
	var customOutputs map[string]string
	// Check both jobs.pre-activation and jobs.pre_activation (users might define both by mistake)
	jobVariants := []string{"pre-activation", string(constants.PreActivationJobName)}
	for _, jobName := range jobVariants {
		preActivationJob, exists := jobs[jobName]
		if !exists {
			continue
		}
		configMap, ok := preActivationJob.(map[string]any)
		if !ok {
			return nil, nil, fmt.Errorf("jobs.%s must be an object, got %T", jobName, preActivationJob)
		}
		variantSteps, variantOutputs, err := extractPreActivationVariantFields(jobName, configMap)
		if err != nil {
			return nil, nil, err
		}
		customSteps = append(customSteps, variantSteps...)
		if len(variantOutputs) > 0 {
			if customOutputs == nil {
				customOutputs = make(map[string]string)
			}
			maps.Copy(customOutputs, variantOutputs)
		}
	}
	return customSteps, customOutputs, nil
}

// extractPreActivationVariantFields validates and extracts steps and outputs from a single
// pre-activation job variant config map (e.g. jobs.pre-activation or jobs.pre_activation).
func extractPreActivationVariantFields(jobName string, configMap map[string]any) ([]string, map[string]string, error) {
	allowedFields := map[string]struct{}{
		"steps":     {},
		"outputs":   {},
		"pre-steps": {}, // handled by generic built-in pre-steps insertion in compiler_jobs.go
	}
	for field := range configMap {
		if field == "setup-steps" {
			return nil, nil, fmt.Errorf(
				"jobs.%s.setup-steps is not allowed: setup-steps are refused for activation/pre-activation jobs because they can short-circuit protections",
				jobName,
			)
		}
		if _, ok := allowedFields[field]; !ok {
			return nil, nil, fmt.Errorf("jobs.%s: unsupported field '%s' - only 'steps', 'outputs', and 'pre-steps' are allowed", jobName, field)
		}
	}
	variantSteps, err := extractPreActivationVariantSteps(jobName, configMap)
	if err != nil {
		return nil, nil, err
	}
	variantOutputs, err := extractPreActivationVariantOutputs(jobName, configMap)
	if err != nil {
		return nil, nil, err
	}
	return variantSteps, variantOutputs, nil
}

// extractPreActivationVariantSteps returns the compiled YAML step strings from jobs.<variant>.steps.
func extractPreActivationVariantSteps(jobName string, configMap map[string]any) ([]string, error) {
	stepsValue, hasSteps := configMap["steps"]
	if !hasSteps {
		return nil, nil
	}
	stepsList, ok := stepsValue.([]any)
	if !ok {
		return nil, fmt.Errorf("jobs.%s.steps must be an array, got %T", jobName, stepsValue)
	}
	var steps []string
	for i, step := range stepsList {
		stepMap, ok := step.(map[string]any)
		if !ok {
			return nil, fmt.Errorf("jobs.%s.steps[%d] must be an object, got %T", jobName, i, step)
		}
		stepYAML, err := ConvertStepToYAML(stepMap)
		if err != nil {
			return nil, fmt.Errorf("failed to convert jobs.%s.steps[%d] to YAML: %w", jobName, i, err)
		}
		steps = append(steps, stepYAML)
	}
	compilerActivationJobsLog.Printf("Extracted %d custom steps from jobs.%s", len(stepsList), jobName)
	return steps, nil
}

// extractPreActivationVariantOutputs returns the string output map from jobs.<variant>.outputs.
func extractPreActivationVariantOutputs(jobName string, configMap map[string]any) (map[string]string, error) {
	outputsValue, hasOutputs := configMap["outputs"]
	if !hasOutputs {
		return nil, nil
	}
	outputsMap, ok := outputsValue.(map[string]any)
	if !ok {
		return nil, fmt.Errorf("jobs.%s.outputs must be an object, got %T", jobName, outputsValue)
	}
	result := make(map[string]string, len(outputsMap))
	for key, val := range outputsMap {
		valStr, ok := val.(string)
		if !ok {
			return nil, fmt.Errorf("jobs.%s.outputs.%s must be a string, got %T", jobName, key, val)
		}
		result[key] = valStr
	}
	compilerActivationJobsLog.Printf("Extracted %d custom outputs from jobs.%s", len(outputsMap), jobName)
	return result, nil
}

// buildPreActivationAppTokenMintStep generates a single GitHub App token mint step for use
// by all skip-if checks in the pre-activation job. The step ID is "pre-activation-app-token".
// Auth configuration comes from the top-level on.github-app field.
func (c *Compiler) buildPreActivationAppTokenMintStep(app *GitHubAppConfig) []string {
	var steps []string
	tokenStepID := constants.PreActivationAppTokenStepID

	steps = append(steps, "      - name: Generate GitHub App token for skip-if checks\n")
	steps = append(steps, fmt.Sprintf("        id: %s\n", tokenStepID))
	if app.shouldIgnoreMissingKey() {
		steps = append(steps, fmt.Sprintf("        if: %s\n", buildIgnoreIfMissingCondition(app)))
	}
	steps = append(steps, fmt.Sprintf("        uses: %s\n", getActionPin("actions/create-github-app-token")))
	steps = append(steps, "        with:\n")
	steps = append(steps, fmt.Sprintf("          client-id: %s\n", app.AppID))
	steps = append(steps, fmt.Sprintf("          private-key: %s\n", app.PrivateKey))

	owner := app.Owner
	if owner == "" {
		owner = "${{ github.repository_owner }}"
	}
	steps = append(steps, fmt.Sprintf("          owner: %s\n", owner))

	if len(app.Repositories) == 1 && app.Repositories[0] == "*" {
		// Org-wide access: omit repositories field entirely
	} else if len(app.Repositories) == 1 {
		steps = append(steps, fmt.Sprintf("          repositories: %s\n", app.Repositories[0]))
	} else if len(app.Repositories) > 1 {
		steps = append(steps, "          repositories: |-\n")
		for _, repo := range app.Repositories {
			steps = append(steps, fmt.Sprintf("            %s\n", repo))
		}
	} else {
		steps = append(steps, "          repositories: ${{ github.event.repository.name }}\n")
	}

	steps = append(steps, "          github-api-url: ${{ github.api_url }}\n")

	return steps
}

// resolvePreActivationSkipIfToken returns the GitHub token expression to use for skip-if check
// steps in the pre-activation job. Priority: App token > custom github-token > empty (default).
// When non-empty, callers should emit `with.github-token: <value>` in the step.
func (c *Compiler) resolvePreActivationSkipIfToken(data *WorkflowData) string {
	if data.ActivationGitHubApp != nil {
		if data.ActivationGitHubApp.shouldIgnoreMissingKey() {
			return combineTokenExpressions(
				fmt.Sprintf("${{ steps.%s.outputs.token }}", constants.PreActivationAppTokenStepID),
				"${{ secrets.GITHUB_TOKEN }}",
			)
		}
		return fmt.Sprintf("${{ steps.%s.outputs.token }}", constants.PreActivationAppTokenStepID)
	}
	if data.ActivationGitHubToken != "" {
		return data.ActivationGitHubToken
	}
	return ""
}

// extractOnSteps extracts the 'steps' field from the 'on:' section of frontmatter.
// These steps are injected into the pre-activation job and their step outcome is wired
// as pre-activation outputs so users can reference them with:
//
//	needs.pre_activation.outputs.<id>_result   (contains outcome: success/failure/cancelled/skipped)
//
// Returns nil if on.steps is not configured.
// Returns an error if on.steps is not an array or contains non-object items.
func extractOnSteps(frontmatter map[string]any) ([]map[string]any, error) {
	onValue, exists := frontmatter["on"]
	if !exists || onValue == nil {
		return nil, nil
	}

	onMap, ok := onValue.(map[string]any)
	if !ok {
		return nil, nil
	}

	stepsValue, exists := onMap["steps"]
	if !exists || stepsValue == nil {
		return nil, nil
	}

	stepsList, ok := stepsValue.([]any)
	if !ok {
		return nil, fmt.Errorf("on.steps must be an array, got %T", stepsValue)
	}

	result := make([]map[string]any, 0, len(stepsList))
	for i, step := range stepsList {
		stepMap, ok := step.(map[string]any)
		if !ok {
			return nil, fmt.Errorf("on.steps[%d] must be an object, got %T", i, step)
		}
		result = append(result, stepMap)
	}

	return result, nil
}

// extractOnPermissions extracts the 'permissions' field from the 'on:' section of frontmatter.
// These permissions are merged into the pre-activation job permissions, allowing users to declare
// extra scopes required by their on.steps (e.g., issues: read for GitHub API calls).
//
// Returns nil if on.permissions is not configured.
func extractOnPermissions(frontmatter map[string]any) *Permissions {
	onValue, exists := frontmatter["on"]
	if !exists || onValue == nil {
		return nil
	}

	onMap, ok := onValue.(map[string]any)
	if !ok {
		return nil
	}

	permsValue, exists := onMap["permissions"]
	if !exists || permsValue == nil {
		return nil
	}

	parser := NewPermissionsParserFromValue(permsValue)
	return parser.ToPermissions()
}

// extractOnNeeds extracts the 'needs' field from the 'on:' section of frontmatter.
// These dependencies are added to both pre_activation and activation jobs.
//
// Returns nil if on.needs is not configured.
func extractOnNeeds(frontmatter map[string]any) ([]string, error) {
	onValue, exists := frontmatter["on"]
	if !exists || onValue == nil {
		return nil, nil
	}

	onMap, ok := onValue.(map[string]any)
	if !ok {
		return nil, nil
	}

	return parseOnNeedsValues(onMap)
}

func parseOnNeedsValues(onMap map[string]any) ([]string, error) {
	if onMap == nil {
		return nil, nil
	}

	needsValue, exists := onMap["needs"]
	if !exists || needsValue == nil {
		return nil, nil
	}

	needsList, ok := needsValue.([]any)
	if !ok {
		return nil, fmt.Errorf("on.needs must be an array, got %T", needsValue)
	}

	result := make([]string, 0, len(needsList))
	for i, need := range needsList {
		needStr, ok := need.(string)
		if !ok {
			return nil, fmt.Errorf("on.needs[%d] must be a string, got %T", i, need)
		}
		result = append(result, needStr)
	}

	return sliceutil.Deduplicate(result), nil
}

// referencesPreActivationOutputs returns true if the condition references the pre_activation job's
// own outputs (e.g., "needs.pre_activation.outputs.foo"). Such conditions cannot be applied to the
// pre_activation job itself (a job cannot reference its own outputs), so they are deferred to
// downstream jobs (activation, agent).
func referencesPreActivationOutputs(condition string) bool {
	if condition == "" {
		return false
	}
	return strings.Contains(condition, "needs."+string(constants.PreActivationJobName)+".outputs.")
}
