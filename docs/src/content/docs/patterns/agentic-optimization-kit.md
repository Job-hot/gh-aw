---
title: Agentic Optimization Kit
description: Use the built-in Agentic Optimization Kit to consolidate token auditing, optimization targeting, and agentic observability into one weekly executive report with actionable prompt artifacts.
---

> [!WARNING]
> **Experimental:** The Agentic Optimization Kit is still experimental! Things may break, change, or be removed without deprecation at any time.

The Agentic Optimization Kit is a weekly unified analyst that consolidates token auditing, optimization targeting, and agentic observability into one executive report with actionable prompt artifacts. It runs on a weekly schedule, analyzes the last 7 days of Copilot workflow run data, selects a high-ROI optimization target, and produces five diagnostic charts alongside a structured discussion report.

This pattern is useful when a repository has enough agentic activity that maintainers need to track cost trends, identify expensive workflows, and surface concrete optimization opportunities — all in one place.

## What it does

The kit runs through three analysis phases each week:

1. **Baseline audit** — Aggregates token usage, cost, turns, action minutes, errors, and warnings across all Copilot workflow runs from the last 7 days. Flags dominant workflows (>30% of total tokens), expensive-per-run workflows (>100k avg tokens/run), and noisy workflows (>50% error/warning rate).

2. **Optimization target selection and analysis** — Selects the highest-token workflow not optimized in the last 14 days, audits its runs in depth, and produces ranked recommendations covering tool usage efficiency, token efficiency, reliability, and prompt structure.

3. **Episode and observability analysis** — Analyzes 30 days of episodic run data to compute workflow instability scores, value proxies, and portfolio positioning. Identifies stale workflows, overlap pairs, and consolidation candidates.

After analysis, it generates five diagnostic charts and publishes a discussion with the full executive report. If repeated escalation signals are detected, it opens one escalation issue.

## Visual diagnostics

The kit produces five charts at publication quality:

### 1. Token Usage by Workflow

A horizontal bar chart of the top 15 workflows by total tokens over the last 7 days. Bars are color-coded by risk flag: dominant (>30% token share), expensive-per-run (>100k avg), or noisy (>50% error rate).

This answers: which workflows are consuming the most tokens, and are those levels expected?

### 2. Historical Token Trend

A line chart of daily total tokens and total cost from the rolling 90-day summary. Captures week-over-week direction to surface whether spend is improving, flat, or growing.

### 3. Episode Risk–Cost Frontier

A scatter plot of execution episodes in cost-risk space. The x-axis is episode cost, the y-axis is a composite risk score (risky nodes, poor-control signals, MCP failures, blocked requests, escalation eligibility), and point size reflects run count. Pareto-frontier outliers are annotated.

This answers: which execution chains are both expensive and risky?

### 4. Workflow Stability Matrix

A heatmap of workflow instability signals. Rows are workflows sorted by instability score; columns are the six underlying signal rates (risky run rate, poor-control rate, resource-heavy rate, latest-success fallback rate, blocked-request rate, MCP failure rate).

This answers: which workflows are chronically unstable, and in which dimensions?

### 5. Repository Portfolio Map

A scatter plot positioning all workflows in value–cost space. The x-axis is recent cost, the y-axis is an evidence-based value proxy (successful usage, stability, repeat use, absence of overkill signals), and the quadrants map to actionable categories: `keep`, `optimize`, `simplify`, `review`.

This answers: which workflows justify their cost, and which warrant a maintainer decision?

## Report structure

The weekly discussion follows this structure:

1. **Executive Summary** — period, total runs/tokens/cost/action-minutes, heavy-hitters, optimization target selection.
2. **Visual Diagnostics** — all five charts, each with a one-sentence `Decision` and `Why it matters`.
3. **Escalation Targets** — workflows that crossed repeated-threshold conditions (omitted when none).
4. **Optimization Target** — the selected workflow, its run metrics, and a ranked recommendation table (up to five recommendations with estimated savings per run).
5. **Five Actionable Prompts** — ready-to-use prompt artifacts covering optimization, stability fix, consolidation, right-sizing, and escalation.
6. **Full Per-Workflow Breakdown** — complete baseline table in a collapsible section.
7. **Episode Detail** — top episodes by risk score in a collapsible section.
8. **Portfolio Opportunities** — stale workflows, overlap pairs, and overkill candidates in a collapsible section.
9. **Optimization Analysis Detail** — full four-area analysis for the selected target in a collapsible section.

## Escalation conditions

The kit creates one escalation issue when any of the following holds across the last 14 days:

- Any episode is marked `escalation_eligible`.
- Two or more runs for the same workflow are classified as `risky`.
- Two or more runs have `new_mcp_failure` or `blocked_requests_increase` in their regression reason codes.
- Two or more runs for the same workflow have a medium or high severity `resource_heavy_for_domain` or `poor_agentic_control` assessment.

If no threshold is crossed, no issue is created.

## Relationship to the Agentic Observability Kit

The Agentic Optimization Kit and the [Agentic Observability Kit](/gh-aw/patterns/agentic-observability-kit/) are complementary:

- The **Observability Kit** focuses on operational regression detection, episode-level risk, and portfolio cleanup at repository scope.
- The **Optimization Kit** adds a weekly cost-focused audit layer with optimization targeting, historical trend tracking, token efficiency analysis, and the five actionable prompts.

Repositories with active Copilot usage benefit from running both. The Observability Kit surfaces operational regressions; the Optimization Kit surfaces spend efficiency and actionable reduction opportunities.

## Metrics and scoring

The kit uses several derived metrics to make the charts decision-oriented.

`episode_risk_score`
A composite risk score for a single execution episode. Combines risky nodes, poor-control nodes, MCP failures, blocked requests, trend signals, and escalation eligibility. Escalation-eligible episodes are weighted at 2× because they already represent multiple threshold crossings.

`workflow_instability_score`
A workflow-level score derived from repeated risky runs, poor-control assessments, resource-heavy assessments, latest-success fallback usage, blocked requests, and MCP failures. Used to separate chronic instability from one-off incidents.

`workflow_value_proxy`
A repository-local proxy for workflow value combining recent successful usage (0.35), stability (0.25), repeat invocations (0.20), and the absence of overkill signals (0.20). Used to position workflows in the portfolio map quadrants.

`workflow_overlap_score`
An approximate similarity score between two workflows. Blends task domain, schedule family, behavior cluster, naming similarity, and assessment similarity. Values ≥ 0.55 are flagged as strong consolidation candidates.

## Flags and thresholds

| Flag | Condition | Implication |
|------|-----------|-------------|
| `is_dominant` | Workflow accounts for > 30% of total tokens | Single workflow concentrating systemic risk |
| `is_expensive_per_run` | Avg tokens/run > 100,000 | Prompt or tool configuration is the primary cost lever |
| `is_noisy` | Error + warning count > 50% of run count | Reliability waste consuming tokens without useful output |

Optimization cooldown prevents the same workflow from being selected again within 14 days.

## Cost caveats

The kit's cost signals are estimates, not billing records.

- `action_minutes` is derived from workflow duration rounded to billable minutes. Use it for relative comparison and trend detection, not as a GitHub invoice figure.
- `estimated_cost` reflects engine log fields. It is appropriate for portfolio prioritization, but should be treated as a run-level estimate rather than finance-grade accounting.
- Effective Tokens (when used instead of raw cost) normalize token classes and apply a per-model multiplier. They make cross-run and cross-model comparisons more useful, but they are not a billing unit.

## Source workflow

The built-in workflow lives at [`/.github/workflows/agentic-optimization-kit.md`](https://github.com/github/gh-aw/blob/main/.github/workflows/agentic-optimization-kit.md).

## Related documentation

- [Agentic Observability Kit](/gh-aw/patterns/agentic-observability-kit/) — operational regression detection and portfolio cleanup
- [Cost Management](/gh-aw/reference/cost-management/) — Actions minutes, inference spend, and optimization levers
- [`gh aw logs`](/gh-aw/setup/cli/#logs) — cross-run log analysis used as the primary data source
- [`gh aw audit`](/gh-aw/setup/cli/#audit) — single-run detailed reports
