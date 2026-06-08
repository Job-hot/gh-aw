//go:build !integration

package workflow

import (
	"errors"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestExpandErrorMessages_UnwrapsJoinedErrors(t *testing.T) {
	err := errors.Join(
		NewValidationError("engine", "", "cannot be empty", "Add engine"),
		NewValidationError("permissions", "", "invalid scope", "Fix permissions"),
	)

	messages := ExpandErrorMessages(err)
	require.Len(t, messages, 2, "Joined errors should expand into separate messages")
	assert.Contains(t, messages[0], "field 'engine'", "Engine validation error should be preserved")
	assert.Contains(t, messages[1], "field 'permissions'", "Permissions validation error should be preserved")
}

func TestBuildPrioritizedErrorReportFromMessages_DefaultLimit(t *testing.T) {
	messages := []string{
		"workflow.md:4:1: error: deprecated field",
		"workflow.md:3:1: error: network.allowed requires strict mode",
		"workflow.md:2:1: error: invalid engine value 'copiliot'",
		"workflow.md:6:1: error: runtime version conflict",
		"workflow.md:5:1: error: event filter is invalid",
		"workflow.md:7:1: error: tools.github config invalid",
	}

	report := BuildPrioritizedErrorReportFromMessages(messages, false)

	require.Equal(t, 6, report.TotalCount, "All non-suppressed errors should be counted")
	require.Len(t, report.DisplayedErrors, 5, "Default report should limit output to five errors")
	assert.Equal(t, SeverityCritical, report.DisplayedErrors[0].Severity, "Critical errors should be first")
	assert.Contains(t, report.DisplayedErrors[0].Message, "invalid engine", "Highest-priority error should be the invalid engine")
	assert.Equal(t, SeverityHigh, report.DisplayedErrors[1].Severity, "High-priority errors should immediately follow critical errors")
	assert.Contains(t, report.DisplayedErrors[1].Message, "network.allowed", "The next prioritized error should be the high-priority network error")
	assert.Equal(t, 1, report.HiddenCount, "One error should be hidden when limiting output")
	require.NotNil(t, report.RecoveryPlan, "Multi-error reports should include a recovery plan")
	assert.NotEmpty(t, report.RecoveryPlan.Steps, "Recovery plan should contain steps")
}

func TestExpandErrorMessages_SplitsBundledSchemaFailures(t *testing.T) {
	messages := ExpandErrorMessages(errors.New(`/tmp/workflow.md:9:5: error: Multiple schema validation failures:
- 'tools/github' (line 9, col 5): Unknown property: foo
- 'on' (line 11, col 3): Unknown property: pull-request
 9 | foo: bar`))

	require.Len(t, messages, 2, "Bundled schema failures should be split into separate display messages")
	assert.Contains(t, messages[0], "tools/github", "The tool schema failure should be preserved")
	assert.Contains(t, messages[1], "pull-request", "The event schema failure should be preserved")
}

func TestBuildPrioritizedErrorReportFromMessages_ShowAll(t *testing.T) {
	messages := []string{
		"workflow.md:1:1: error: invalid engine value 'copiliot'",
		"workflow.md:2:1: error: runtime version conflict",
		"workflow.md:3:1: error: deprecated field",
	}

	report := BuildPrioritizedErrorReportFromMessages(messages, true)

	require.Len(t, report.DisplayedErrors, 3, "Show-all reports should include every prioritized error")
	assert.Zero(t, report.HiddenCount, "No errors should be hidden in show-all mode")
}

func TestBuildPrioritizedErrorReportFromMessages_SuppressesCascadingSyntaxErrors(t *testing.T) {
	messages := []string{
		"workflow.md:2:1: error: failed to parse YAML frontmatter: mapping values are not allowed in this context",
		"[2026-01-01T00:00:00Z] Validation failed for field 'engine'\nReason: cannot be empty",
	}

	report := BuildPrioritizedErrorReportFromMessages(messages, true)

	require.Len(t, report.DisplayedErrors, 1, "The syntax root cause should suppress cascading configuration errors")
	assert.Equal(t, SeverityCritical, report.DisplayedErrors[0].Severity, "The remaining error should be critical")
	assert.Equal(t, 1, report.SuppressedCount, "One cascading error should be suppressed")
}

func TestBuildPrioritizedErrorReportFromMessages_YAMLMissingColonGetsSyntaxSuggestion(t *testing.T) {
	messages := []string{
		`/tmp/workflow.md:2:1: error: missing ':' after key — YAML mapping entries require 'key: value' format
2 | on
3 |   workflow_dispatch:
  | ^`,
	}

	report := BuildPrioritizedErrorReportFromMessages(messages, true)
	require.Len(t, report.DisplayedErrors, 1, "Expected one prioritized error")
	assert.Equal(t, SeverityCritical, report.DisplayedErrors[0].Severity, "Missing colon should be classified as critical syntax")
	assert.Equal(t, "syntax", report.DisplayedErrors[0].Category, "Missing colon should be a syntax category")
	assert.Contains(t, report.DisplayedErrors[0].Suggestion, "Add \":\" after the key", "Suggestion should be YAML-specific")
}

func TestExpandErrorMessages_SynthesizesInvalidEngineTypoRootCause(t *testing.T) {
	raw := `/tmp/workflow.md:5:1: error: Multiple schema validation failures:
- 'engine' (line 5, col 1): must satisfy oneOf
- 'engine' (line 5, col 1): Unknown property: max-ai-credits
4 | on: push
5 | engine: CLAUD
6 | max-ai-credits: 1200`

	messages := ExpandErrorMessages(errors.New(raw))
	require.NotEmpty(t, messages, "Expected synthesized engine error plus original schema failures")
	assert.Contains(t, messages[0], `/tmp/workflow.md:5:1: error:`, "Synthesized error should preserve file path and line")
	assert.Contains(t, messages[0], `unknown engine "CLAUD"`, "Expected synthesized root-cause engine typo error to be first")
	assert.Contains(t, messages[0], `Did you mean "claude"?`, "Expected synthesized error to include the nearest suggestion")

	foundOriginalSchemaError := false
	for _, message := range messages[1:] {
		if strings.Contains(message, "must satisfy oneOf") {
			foundOriginalSchemaError = true
			break
		}
	}
	assert.True(t, foundOriginalSchemaError, "Original schema errors should remain in the expanded output")
}

func TestExpandErrorMessages_UnknownEngineWithoutSuggestionStillSurfacesRootCause(t *testing.T) {
	raw := `/tmp/workflow.md:5:1: error: Multiple schema validation failures:
- 'engine' (line 5, col 1): must satisfy oneOf
- 'engine' (line 5, col 1): Unknown property: max-ai-credits
4 | on: push
5 | engine: FOOBARXYZ
6 | max-ai-credits: 1200`

	messages := ExpandErrorMessages(errors.New(raw))
	require.NotEmpty(t, messages, "Expected synthesized engine error plus original schema failures")
	assert.Contains(t, messages[0], `unknown engine "FOOBARXYZ"`, "Unknown engine should surface even without a near-match suggestion")
	assert.NotContains(t, messages[0], `Did you mean`, "No suggestion should be added when there is no near match")
}

func TestBuildPrioritizedErrorReportFromMessages_ClassifiesFromHeadlineOnly(t *testing.T) {
	messages := []string{
		`/tmp/workflow.md:5:1: error: unknown engine "claud". Valid engines are: antigravity, claude, codex, copilot, crush, gemini, opencode, pi. Did you mean "claude"?
4 | on:
5 |   pull_request:
6 | engine: claud`,
	}

	report := BuildPrioritizedErrorReportFromMessages(messages, true)
	require.Len(t, report.DisplayedErrors, 1, "Expected one prioritized error")
	assert.Equal(t, "configuration", report.DisplayedErrors[0].Category, "Context lines should not change the error category")
	assert.Contains(t, report.DisplayedErrors[0].Suggestion, "Check the engine name", "Engine-specific suggestion should win over event/filter hinting")
}

func TestNewValidationError_ClassifiesSeverity(t *testing.T) {
	err := NewValidationError("network.allowed", "example.com", "requires strict mode", "Enable strict mode")

	require.NotNil(t, err, "Validation error should be created")
	assert.Equal(t, SeverityHigh, err.Severity, "Network strict-mode errors should be high priority")
	assert.Equal(t, "permissions", err.Category, "Network strict-mode errors should be categorized as permissions")
}
