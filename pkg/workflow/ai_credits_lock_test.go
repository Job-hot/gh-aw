//go:build !integration

package workflow

import (
	"os"
	"path/filepath"
	"regexp"
	"testing"

	"github.com/github/gh-aw/pkg/stringutil"
	"github.com/github/gh-aw/pkg/testutil"
)

func compileWorkflowToLockContent(t *testing.T, workflow string) string {
	t.Helper()

	testDir := testutil.TempDir(t, "ai-credits-lock-*")
	workflowFile := filepath.Join(testDir, "workflow.md")
	if err := os.WriteFile(workflowFile, []byte(workflow), 0o644); err != nil {
		t.Fatalf("failed to write test workflow: %v", err)
	}

	compiler := NewCompiler()
	if err := compiler.CompileWorkflow(workflowFile); err != nil {
		t.Fatalf("failed to compile workflow: %v", err)
	}

	lockFile := stringutil.MarkdownToLockFile(workflowFile)
	lockContent, err := os.ReadFile(lockFile)
	if err != nil {
		t.Fatalf("failed to read lock file: %v", err)
	}

	return string(lockContent)
}

func TestCompiledLockUsesGitHubVarFallbacksForAICreditDefaults(t *testing.T) {
	lockContent := compileWorkflowToLockContent(t, `---
on:
  workflow_dispatch:
permissions:
  contents: read
engine: claude
---

# AI credit defaults`)

	if !regexp.MustCompile(`GH_AW_MAX_AI_CREDITS:\s+\$\{\{\s*vars\.GH_AW_DEFAULT_MAX_AI_CREDITS \|\| '1000'\s*\}\}`).MatchString(lockContent) {
		t.Fatalf("expected lock file to include GH_AW_MAX_AI_CREDITS vars fallback, got:\n%s", lockContent)
	}
	if !regexp.MustCompile(`GH_AW_MAX_DAILY_AI_CREDITS:\s+\$\{\{\s*vars\.GH_AW_DEFAULT_MAX_DAILY_AI_CREDITS \|\| '5000'\s*\}\}`).MatchString(lockContent) {
		t.Fatalf("expected lock file to include GH_AW_MAX_DAILY_AI_CREDITS vars fallback, got:\n%s", lockContent)
	}
}

func TestCompiledLockPrefersExplicitAICreditFrontmatterValues(t *testing.T) {
	lockContent := compileWorkflowToLockContent(t, `---
on:
  workflow_dispatch:
permissions:
  contents: read
engine: claude
max-ai-credits: 1234
max-daily-ai-credits: 5678
---

# Explicit AI credit defaults`)

	if !regexp.MustCompile(`GH_AW_MAX_AI_CREDITS:\s+"?1234"?`).MatchString(lockContent) {
		t.Fatalf("expected lock file to include explicit GH_AW_MAX_AI_CREDITS value, got:\n%s", lockContent)
	}
	if !regexp.MustCompile(`GH_AW_MAX_DAILY_AI_CREDITS:\s+"?5678"?`).MatchString(lockContent) {
		t.Fatalf("expected lock file to include explicit GH_AW_MAX_DAILY_AI_CREDITS value, got:\n%s", lockContent)
	}
}
