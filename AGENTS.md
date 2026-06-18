# gh-aw (GitHub Agentic Workflows)

`gh aw` CLI extension — compiles markdown workflows into GitHub Actions. **copilot** is a separate runtime engine, not this tool.

## Ambient Context

Load only at first invocation; load everything else lazily via skills.

| File | Purpose |
|---|---|
| `AGENTS.md` | Global rules and routing |
| `SKILL.md` | Repository capability summary |

## Critical Rules

1. Changed files → run `make agent-report-progress` (must pass), then `report_progress`.
2. Go changes → run `make fmt`.
3. Workflow `.md` changes under `.github/workflows/` → run `make recompile`.
4. Never add `.lock.yml` to `.gitignore`.

## Upstream-Managed Workflows (read-only)

Workflows with `source:` frontmatter (e.g. `source: githubnext/agentic-ops@<ref>`) are read-only. Do not edit their source files or `.lock.yml` files directly. To update: run `gh aw update`, bump `source: ...@...`, or change upstream first.

## Skill Routing

Load a skill only when the task matches its intent.

| Intent | Skill |
|---|---|
| Workflow create/update/debug/upgrade | `.github/skills/agentic-workflows/SKILL.md` |
| Engineering conventions, validation, command playbooks | `.github/skills/developer/SKILL.md` |
| Error handling design/patterns | `.github/skills/error-recovery-patterns/SKILL.md` |
| GitHub MCP usage | `.github/skills/github-mcp-server/SKILL.md` |
| Issue/PR/workflow/discussion/label queries | `.github/skills/github-*-query/SKILL.md` |
| Documentation writing | `.github/skills/documentation/SKILL.md` |
| git/gh/checkout credential review | `.github/skills/checkout-credential-review/SKILL.md` |
