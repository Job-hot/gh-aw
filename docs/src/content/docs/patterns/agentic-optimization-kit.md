---
title: Agentic Optimization Kit
description: Use the built-in Agentic Optimization Kit to consolidate token auditing, optimization targeting, and agentic observability into one weekly report with prompt artifacts and cost-reduction recommendations.
---

> [!WARNING]
> **Experimental:** The Agentic Optimization Kit is still experimental. Things may break, change, or be removed without deprecation at any time.

The Agentic Optimization Kit is a weekly consolidated workflow that combines token auditing, optimization targeting, and agentic observability into one actionable executive report. It analyzes the last 7 days of Copilot workflow runs, selects a high-ROI optimization target, produces ranked recommendations with prompt artifacts, and publishes a chart-backed discussion.

It runs as a single workflow on a weekly schedule, replacing the need to run a token audit, an optimization analysis, and an observability report separately.

## What it does

The kit runs three ordered phases in one session:

1. **Baseline Audit** — downloads the last 7 days of Copilot workflow run data and aggregates token usage, cost, action minutes, and reliability signals per workflow. Flags dominant workflows (>30% of total tokens), expensive-per-run workflows (>100k tokens/run), and noisy workflows (>50% error/warning rate). Saves a daily snapshot to repo memory for trend tracking.

2. **Optimization Target Analysis** — selects the highest-cost workflow not optimized in the last 14 days, audits its individual runs, evaluates tool usage and prompt efficiency, and produces a ranked list of ≤5 recommendations with estimated token savings and concrete actions.

3. **Episode and Observability Analysis** — uses the `agentic-workflows` MCP to retrieve episode-level data for the full repository over the last 30 days. Builds episode risk scores, workflow instability scores, and value proxies. Generates five charts and publishes a discussion with executive summary, portfolio map, and recommended actions.

## Deployment

Install the built-in workflow for repository-scoped optimization reports:

```aw wrap
---
on:
  schedule: weekly on monday
  workflow_dispatch:
permissions:
  contents: read
  actions: read
  issues: read
  pull-requests: read
  discussions: read
engine: copilot
strict: true
tools:
  cli-proxy: true
  agentic-workflows:
  github:
    toolsets: [default, discussions]
  bash:
    - "*"
safe-outputs:
  create-issue:
    title-prefix: "[agentic-optimization escalation] "
    labels: [agentics, warning, observability]
    close-older-issues: true
    max: 1
timeout-minutes: 35
---

# Agentic Optimization Kit

Run the weekly consolidated audit, optimization targeting, and observability analysis. Publish a chart-backed discussion with findings and recommended actions.
```

The workflow requires `actions: read` (for log access), `discussions: read` (for observability context), `issues: read` (for escalation deduplication), and `pull-requests: read`. The `cli-proxy: true` setting is required for the agent to read workflow source files via the `gh` CLI.

## Charts generated

Each weekly report includes five charts:

| Chart | What it shows |
|-------|---------------|
| **Token Usage by Workflow** | Top 15 workflows by total tokens, colored by dominant/expensive/noisy flags |
| **Historical Token Trend** | Daily total tokens and cost over the rolling 90-day window (from repo memory) |
| **Episode Risk–Cost Frontier** | Episodes plotted by cost and risk score; size = run count |
| **Workflow Stability Matrix** | Per-workflow heatmap of risky run rate, poor-control rate, resource-heavy rate, and failure signals |
| **Repository Portfolio Map** | Workflows in cost–value space; quadrants labeled keep / optimize / simplify / review |

## Memory and trend tracking

The kit stores daily audit snapshots and a 90-day rolling summary in repo memory (`memory/token-audit` branch). It also maintains an optimization log that records which workflows have been targeted, preventing the same workflow from being selected again within a 14-day cooldown window.

When no prior snapshots exist, the kit starts fresh with the current run's data. The trend chart is skipped or simplified until at least two data points are available.

## Escalation behavior

When the kit detects patterns that warrant owner action (repeated high-risk episodes, chronically unstable workflows, or unresolved optimization targets), it opens one escalation issue with a `[agentic-optimization escalation]` prefix. The `close-older-issues: true` setting ensures stale escalations from prior runs are closed when a new one is filed, keeping the issue tracker clean.

If no action is needed, the kit calls `noop` without creating an issue.

## Relationship to other kits

The Agentic Optimization Kit consolidates signals from two standalone workflows:

- [Agentic Observability Kit](/gh-aw/patterns/agentic-observability-kit/) — focuses on episode-level behavior, regression detection, and portfolio analysis. The optimization kit includes this layer as Phase 3, running it after the baseline audit and target analysis.
- Token Audit — provides the raw per-workflow cost and usage breakdown. The optimization kit runs an equivalent baseline audit as Phase 1.

If you already run the Agentic Observability Kit separately, consider whether consolidating into this single weekly workflow reduces session overhead while retaining the same coverage. The two workflows can coexist — the optimization kit runs deeper analysis on a single selected target, while the observability kit covers the full portfolio in more detail.

## When to use it

This pattern fits when:

- A repository runs multiple Copilot workflows and maintainers want a single weekly report covering cost, reliability, and optimization opportunities.
- Per-run inspection is too noisy and a consolidated portfolio view is needed.
- The team wants evidence-based prioritization for model downgrades, prompt tightening, or deterministic automation replacements.
- Trend tracking across weeks is more useful than individual run analysis.
