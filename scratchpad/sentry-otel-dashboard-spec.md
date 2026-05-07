# Sentry OTEL Dashboard Specification

> Added in May 2026 to define the first OTEL-first Sentry dashboard and alert baseline for `gh-aw` runtime telemetry.

## Goal

Provide a Sentry dashboard and alert model that uses live OTEL conclusion-span attributes emitted during workflow execution.
This is intended to replace logs-first operational reporting for the common optimization and observability questions raised in Discussion #28810.

## Source of Truth

Use only `gh-aw.job.conclusion` spans emitted by `actions/setup/js/send_otlp_span.cjs`.

Why this span:
- It exists for every completed job.
- It already carries workflow, engine, run, rate-limit, output, observability, and optimization attributes.
- It does not require GitHub API log downloads or post-run artifact reconstruction.

## Required Filters

Every dashboard query should start with these filters unless a panel explicitly needs a broader scope:

- `span.op: gh-aw.job.conclusion`
- `service.name: gh-aw`

Recommended slice dimensions:
- `gh-aw.workflow.name`
- `gh-aw.engine.id`
- `gh-aw.run.actor`
- `gh-aw.staged`
- `gh-aw.observability.runtime_status`
- `gh-aw.observability.posture`

## Query Building Rules

Use these conventions consistently so panels from different dashboards can be compared without reinterpreting the query shape.

- Base filter: `span.op:gh-aw.job.conclusion service.name:gh-aw`
- Default time window for fleet views: 7 days
- Default time window for alerts and incident triage views: 1 hour to 24 hours depending on signal volatility
- Default group-by for workflow-level tuning: `gh-aw.workflow.name`
- Default secondary group-by for engine comparisons: `gh-aw.engine.id`
- Prefer `avg(...)` for stable score panels, `p95(...)` for regression alerts, and `sum(...)` for fleet cost and token volume
- When a numeric attribute is optional, expect nulls and treat the panel as partial coverage rather than backfilling from logs

## Sentry Query Recipes

This section translates the dashboard set into concrete panel recipes.
The exact Sentry UI labels may vary, but the filter, aggregate, and group-by should remain the same.

### Fleet Overview Queries

1. Run count by workflow
Filter:
```text
span.op:gh-aw.job.conclusion service.name:gh-aw
```
Aggregate:
```text
count()
```
Group by:
```text
gh-aw.workflow.name
```

2. Runtime status distribution
Filter:
```text
span.op:gh-aw.job.conclusion service.name:gh-aw
```
Aggregate:
```text
count()
```
Group by:
```text
gh-aw.observability.runtime_status
```

3. Average runtime risk by workflow
Filter:
```text
span.op:gh-aw.job.conclusion service.name:gh-aw
```
Aggregate:
```text
avg(gh-aw.optimization.runtime_risk_score)
```
Group by:
```text
gh-aw.workflow.name
```

4. Average optimization score by workflow
Filter:
```text
span.op:gh-aw.job.conclusion service.name:gh-aw
```
Aggregate:
```text
avg(gh-aw.optimization.score)
```
Group by:
```text
gh-aw.workflow.name
```

5. Average estimated cost by workflow
Filter:
```text
span.op:gh-aw.job.conclusion service.name:gh-aw has:gh-aw.optimization.estimated_cost_usd
```
Aggregate:
```text
avg(gh-aw.optimization.estimated_cost_usd)
```
Group by:
```text
gh-aw.workflow.name
```

### Cost and Token Queries

1. Total token volume by workflow
Filter:
```text
span.op:gh-aw.job.conclusion service.name:gh-aw has:gh-aw.optimization.total_tokens
```
Aggregate:
```text
sum(gh-aw.optimization.total_tokens)
```
Group by:
```text
gh-aw.workflow.name
```

2. Average tokens by engine
Filter:
```text
span.op:gh-aw.job.conclusion service.name:gh-aw has:gh-aw.optimization.total_tokens
```
Aggregate:
```text
avg(gh-aw.optimization.total_tokens)
```
Group by:
```text
gh-aw.engine.id
```

3. Token intensity distribution
Filter:
```text
span.op:gh-aw.job.conclusion service.name:gh-aw
```
Aggregate:
```text
count()
```
Group by:
```text
gh-aw.optimization.intensity
```

4. Cache efficiency by workflow
Filter:
```text
span.op:gh-aw.job.conclusion service.name:gh-aw has:gh-aw.optimization.cache_efficiency
```
Aggregate:
```text
avg(gh-aw.optimization.cache_efficiency)
```
Group by:
```text
gh-aw.workflow.name
```

5. Cost per created item
Filter:
```text
span.op:gh-aw.job.conclusion service.name:gh-aw has:gh-aw.optimization.estimated_cost_usd has:gh-aw.observability.created_item_count
```
Formula:
```text
avg(gh-aw.optimization.estimated_cost_usd) / clamp_min(avg(gh-aw.observability.created_item_count), 1)
```
Group by:
```text
gh-aw.workflow.name
```
If the formula editor does not support `clamp_min`, use two adjacent panels instead:
- `avg(gh-aw.optimization.estimated_cost_usd)` by workflow
- `avg(gh-aw.observability.created_item_count)` by workflow

### Reliability and Risk Queries

1. Erroring runs by workflow
Filter:
```text
span.op:gh-aw.job.conclusion service.name:gh-aw gh-aw.observability.runtime_status:error
```
Aggregate:
```text
count()
```
Group by:
```text
gh-aw.workflow.name
```

2. Degraded runs by workflow
Filter:
```text
span.op:gh-aw.job.conclusion service.name:gh-aw gh-aw.observability.runtime_status:degraded
```
Aggregate:
```text
count()
```
Group by:
```text
gh-aw.workflow.name
```

3. Average warning count by workflow
Filter:
```text
span.op:gh-aw.job.conclusion service.name:gh-aw
```
Aggregate:
```text
avg(gh-aw.observability.warning_count)
```
Group by:
```text
gh-aw.workflow.name
```

4. Average output error count by workflow
Filter:
```text
span.op:gh-aw.job.conclusion service.name:gh-aw
```
Aggregate:
```text
avg(gh-aw.observability.output_error_count)
```
Group by:
```text
gh-aw.workflow.name
```

5. Average turn count by workflow
Filter:
```text
span.op:gh-aw.job.conclusion service.name:gh-aw
```
Aggregate:
```text
avg(gh-aw.observability.turn_count)
```
Group by:
```text
gh-aw.workflow.name
```

### Firewall and Rate-Limit Queries

1. Firewall-enabled run count
Filter:
```text
span.op:gh-aw.job.conclusion service.name:gh-aw gh-aw.observability.firewall_enabled:true
```
Aggregate:
```text
count()
```
Group by:
```text
gh-aw.workflow.name
```

2. Blocked requests by workflow
Filter:
```text
span.op:gh-aw.job.conclusion service.name:gh-aw
```
Aggregate:
```text
sum(gh-aw.observability.blocked_requests)
```
Group by:
```text
gh-aw.workflow.name
```

3. Average rate-limit remaining by workflow
Filter:
```text
span.op:gh-aw.job.conclusion service.name:gh-aw has:gh-aw.github.rate_limit.remaining
```
Aggregate:
```text
avg(gh-aw.github.rate_limit.remaining)
```
Group by:
```text
gh-aw.workflow.name
```

4. Minimum rate-limit remaining by engine
Filter:
```text
span.op:gh-aw.job.conclusion service.name:gh-aw has:gh-aw.github.rate_limit.remaining
```
Aggregate:
```text
min(gh-aw.github.rate_limit.remaining)
```
Group by:
```text
gh-aw.engine.id
```

### Write Activity Queries

1. Write-capable run count
Filter:
```text
span.op:gh-aw.job.conclusion service.name:gh-aw gh-aw.observability.posture:write-capable
```
Aggregate:
```text
count()
```
Group by:
```text
gh-aw.workflow.name
```

2. Average created item count by workflow
Filter:
```text
span.op:gh-aw.job.conclusion service.name:gh-aw
```
Aggregate:
```text
avg(gh-aw.observability.created_item_count)
```
Group by:
```text
gh-aw.workflow.name
```

3. Write-capable run count by engine
Filter:
```text
span.op:gh-aw.job.conclusion service.name:gh-aw gh-aw.observability.posture:write-capable
```
Aggregate:
```text
count()
```
Group by:
```text
gh-aw.engine.id
```

## Dashboard Build Checklist

Use this checklist when creating the dashboards in Sentry:

1. Create a saved base filter for `span.op:gh-aw.job.conclusion service.name:gh-aw`.
2. Build fleet panels first, because they validate field coverage and cardinality.
3. Add `has:` guards to numeric fields that may be absent, especially cost, tokens, cache, and rate-limit data.
4. Prefer workflow-level grouping before actor-level grouping to keep panel cardinality bounded.
5. Use the same time window across companion panels in a dashboard.
6. Add one table panel per dashboard for top offenders by workflow.
7. Add one time-series panel per dashboard for trend direction.
8. Keep alert thresholds tied to p95 or min aggregates, not visual-only averages.

## Saved Searches

Create these saved searches first so dashboard builders can reuse the same base filters and avoid query drift.

| Saved Search Name | Query |
|---|---|
| `gh-aw conclusion spans` | `span.op:gh-aw.job.conclusion service.name:gh-aw` |
| `gh-aw degraded runs` | `span.op:gh-aw.job.conclusion service.name:gh-aw gh-aw.observability.runtime_status:degraded` |
| `gh-aw error runs` | `span.op:gh-aw.job.conclusion service.name:gh-aw gh-aw.observability.runtime_status:error` |
| `gh-aw write-capable runs` | `span.op:gh-aw.job.conclusion service.name:gh-aw gh-aw.observability.posture:write-capable` |
| `gh-aw firewall-enabled runs` | `span.op:gh-aw.job.conclusion service.name:gh-aw gh-aw.observability.firewall_enabled:true` |
| `gh-aw token-bearing runs` | `span.op:gh-aw.job.conclusion service.name:gh-aw has:gh-aw.optimization.total_tokens` |

## Dashboard Blueprint

Use these dashboard names and panel types so the initial Sentry workspace is consistent across teams.

### 1. gh-aw Fleet Overview

Default window: 7 days

| Panel Title | Visualization | Backing Query |
|---|---|---|
| Runs by Workflow | table | `count()` grouped by `gh-aw.workflow.name` on `gh-aw conclusion spans` |
| Runs by Engine | bar chart | `count()` grouped by `gh-aw.engine.id` on `gh-aw conclusion spans` |
| Runtime Status Distribution | pie or stacked bar | `count()` grouped by `gh-aw.observability.runtime_status` on `gh-aw conclusion spans` |
| Average Runtime Risk | table | `avg(gh-aw.optimization.runtime_risk_score)` grouped by `gh-aw.workflow.name` on `gh-aw conclusion spans` |
| Average Optimization Score | line chart | `avg(gh-aw.optimization.score)` grouped by `gh-aw.workflow.name` on `gh-aw conclusion spans` |
| Average Estimated Cost | table | `avg(gh-aw.optimization.estimated_cost_usd)` grouped by `gh-aw.workflow.name` on `gh-aw conclusion spans` with `has:gh-aw.optimization.estimated_cost_usd` |

### 2. gh-aw Cost and Token Dashboard

Default window: 7 days

| Panel Title | Visualization | Backing Query |
|---|---|---|
| Total Tokens by Workflow | table | `sum(gh-aw.optimization.total_tokens)` grouped by `gh-aw.workflow.name` on `gh-aw token-bearing runs` |
| Average Tokens by Engine | bar chart | `avg(gh-aw.optimization.total_tokens)` grouped by `gh-aw.engine.id` on `gh-aw token-bearing runs` |
| Token Intensity Distribution | pie or bar | `count()` grouped by `gh-aw.optimization.intensity` on `gh-aw conclusion spans` |
| Cache Efficiency by Workflow | table | `avg(gh-aw.optimization.cache_efficiency)` grouped by `gh-aw.workflow.name` on `gh-aw conclusion spans` with `has:gh-aw.optimization.cache_efficiency` |
| Estimated Cost by Workflow | line chart | `avg(gh-aw.optimization.estimated_cost_usd)` grouped by `gh-aw.workflow.name` on `gh-aw conclusion spans` with `has:gh-aw.optimization.estimated_cost_usd` |
| Cost per Created Item | formula or paired table | `avg(cost_usd) / clamp_min(avg(created_item_count), 1)` or adjacent cost and item-count panels |

### 3. gh-aw Reliability and Risk Dashboard

Default window: 24 hours

| Panel Title | Visualization | Backing Query |
|---|---|---|
| Erroring Runs | table | `count()` grouped by `gh-aw.workflow.name` on `gh-aw error runs` |
| Degraded Runs | table | `count()` grouped by `gh-aw.workflow.name` on `gh-aw degraded runs` |
| Average Warning Count | line chart | `avg(gh-aw.observability.warning_count)` grouped by `gh-aw.workflow.name` on `gh-aw conclusion spans` |
| Average Output Error Count | line chart | `avg(gh-aw.observability.output_error_count)` grouped by `gh-aw.workflow.name` on `gh-aw conclusion spans` |
| Average Turn Count | table | `avg(gh-aw.observability.turn_count)` grouped by `gh-aw.workflow.name` on `gh-aw conclusion spans` |
| Top Runtime Risk Workflows | table | `p95(gh-aw.optimization.runtime_risk_score)` grouped by `gh-aw.workflow.name` on `gh-aw conclusion spans` |

### 4. gh-aw Firewall and Rate-Limit Dashboard

Default window: 24 hours

| Panel Title | Visualization | Backing Query |
|---|---|---|
| Firewall-Enabled Runs | table | `count()` grouped by `gh-aw.workflow.name` on `gh-aw firewall-enabled runs` |
| Blocked Requests by Workflow | table | `sum(gh-aw.observability.blocked_requests)` grouped by `gh-aw.workflow.name` on `gh-aw conclusion spans` |
| Average Rate-Limit Remaining | line chart | `avg(gh-aw.github.rate_limit.remaining)` grouped by `gh-aw.workflow.name` on `gh-aw conclusion spans` with `has:gh-aw.github.rate_limit.remaining` |
| Minimum Rate-Limit Remaining by Engine | table | `min(gh-aw.github.rate_limit.remaining)` grouped by `gh-aw.engine.id` on `gh-aw conclusion spans` with `has:gh-aw.github.rate_limit.remaining` |
| Blocked Requests by Posture | bar chart | `sum(gh-aw.observability.blocked_requests)` grouped by `gh-aw.observability.posture` on `gh-aw conclusion spans` |

### 5. gh-aw Write Activity Dashboard

Default window: 7 days

| Panel Title | Visualization | Backing Query |
|---|---|---|
| Write-Capable Runs by Workflow | table | `count()` grouped by `gh-aw.workflow.name` on `gh-aw write-capable runs` |
| Average Created Item Count | line chart | `avg(gh-aw.observability.created_item_count)` grouped by `gh-aw.workflow.name` on `gh-aw conclusion spans` |
| Write-Capable Runs by Engine | bar chart | `count()` grouped by `gh-aw.engine.id` on `gh-aw write-capable runs` |
| Created Item Type Coverage | table | `count()` grouped by `gh-aw.observability.created_item_types` on `gh-aw write-capable runs` |

## Implementation Notes

- There is no dashboard JSON or provisioning artifact format in this repo today.
- The repo does include a shared Sentry MCP workflow file at `.github/workflows/shared/mcp/sentry.md`, but that is for Sentry API access in workflows, not dashboard provisioning.
- Until a provisioning format is introduced, this document is the copy-paste source for building the first dashboard set.

## Alert Specification

Start with a small set of alerts tied to runtime facts that already exist.
Do not alert on every panel; alert on conditions that indicate operational regression.

### Alert 1: High Runtime Risk

- Condition: `p95(gh-aw.optimization.runtime_risk_score) >= 60` over 1 hour
- Scope: grouped by `gh-aw.workflow.name`
- Severity: warning
- Purpose: catch workflows trending toward unstable runs before they hard-fail

### Alert 2: Erroring Conclusion Runs

- Condition: count of `gh-aw.observability.runtime_status:error` >= 3 over 30 minutes
- Scope: grouped by `gh-aw.workflow.name`
- Severity: critical
- Purpose: detect repeated bad runs without opening action logs

### Alert 3: Degraded Network Behavior

- Condition: average `gh-aw.observability.blocked_requests > 0` over 1 hour
- Scope: grouped by `gh-aw.workflow.name`
- Severity: warning
- Purpose: highlight workflows that routinely collide with firewall policy or filtered MCP requests

### Alert 4: High Token Spend

- Condition: `p95(gh-aw.optimization.total_tokens) >= 250000` over 24 hours
- Scope: grouped by `gh-aw.workflow.name`
- Severity: warning
- Purpose: surface token-heavy runs suitable for prompt or model optimization

### Alert 5: Low Cache Efficiency

- Condition: average `gh-aw.optimization.cache_efficiency < 0.35` over 24 hours
- Scope: grouped by `gh-aw.workflow.name`
- Severity: warning
- Purpose: identify runs that are paying repeated input-token cost without meaningful cache reuse

### Alert 6: Rate-Limit Headroom Compression

- Condition: minimum `gh-aw.github.rate_limit.remaining < 500` over 1 hour
- Scope: grouped by `gh-aw.workflow.name`
- Severity: warning
- Purpose: detect workflows approaching the GitHub API scaling boundary

## What This Replaces

This dashboard baseline replaces the need to answer these common operational questions by downloading logs:

- Which workflows are the most expensive?
- Which workflows are consuming the most tokens?
- Which workflows are noisy, degraded, or warning-heavy?
- Which workflows are repeatedly hitting network or rate-limit constraints?
- Which workflows are mostly read-only versus write-capable?

## Known Gaps

The following higher-level assessment labels are not part of the runtime OTEL baseline yet:

- `resource_heavy_for_domain`
- `poor_agentic_control`
- `partially_reducible`
- `model_downgrade_available`

Those still come from the audit and logs analysis path rather than the live runtime producer.
They should be added only after a dedicated runtime producer exists for them.

## Recommended Rollout Order

1. Create the five dashboards above using only the current runtime span fields.
2. Run them long enough to identify which decisions still require higher-level assessments.
3. Add runtime-produced assessment flags only for those missing decisions.

This keeps the operational baseline OTEL-first and prevents a return to logs-first observability.