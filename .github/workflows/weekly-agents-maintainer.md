---
emoji: "🛠️"
name: AGENTS.md Maintainer
description: Weekly automation to keep AGENTS.md accurate by reviewing merged pull requests and updated source files, then opening a PR when updates are needed.
on:
  schedule:
    - cron: "weekly"
  workflow_dispatch: null
permissions:
  contents: read
  pull-requests: read
strict: true
tracker-id: weekly-agents-maintainer
engine: copilot
timeout-minutes: 25
network:
  allowed:
    - defaults
    - github
tools:
  bash: true
  github:
    mode: gh-proxy
    toolsets: [repos, pull_requests]
  edit:
safe-outputs:
  create-pull-request:
    title-prefix: "[agents] "
    labels: [automation, documentation]
    allowed-files: [AGENTS.md]
    expires: 2d
  noop: null
---

# AGENTS.md Maintainer

Keep `AGENTS.md` accurate and current by reviewing merged pull requests and updated source files on a weekly cadence.

## Task

1. Identify merged pull requests and repository source file changes since the last weekly run.
2. Review `AGENTS.md` for any missing or stale guidance based on:
   - merged PRs affecting workflow conventions, agent behavior, automation guidance, or repo policies
   - changed or added source files that affect the information captured in `AGENTS.md`
   - renamed or moved files referenced by `AGENTS.md`
3. If `AGENTS.md` needs updates, edit only that file and create a pull request using the configured `create-pull-request` safe output.
4. If no updates are needed, call `noop` with a concise explanation.

## Guidance

- Use GitHub tools and repository history for analysis.
- Prefer `gh` and GitHub MCP reads for merged PR metadata and file change details.
- Use `git` or `gh` to determine what changed in the last 7 days.
- Do not modify files outside `AGENTS.md`.

## Pull Request Requirements

- Keep the PR focused on maintaining `AGENTS.md`.
- If a PR is created, ensure the title starts with `[agents] ` and the description explains what was updated and why.
- Do not open a PR if `AGENTS.md` is already accurate.
