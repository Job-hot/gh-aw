# Agentic Workflow Audit — Repo Memory

Persistent state seeded by the Agentic Workflow Audit Agent.

## Layout

- `audits/<YYYY-MM-DD>.json` — full daily audit snapshot
- `audits/index.json` — chronological index of all audits
- `metrics/daily.jsonl` — one-line-per-day metrics for fast trend computation
- `patterns/errors.json` — recurring error signatures
- `patterns/missing-tools.json` — recurring missing-tool events
- `patterns/mcp-failures.json` — recurring MCP server failures
- `trending/workflow_health/history.jsonl` — daily health snapshots for charts
- `trending/token_cost/history.jsonl` — daily token/cost snapshots for charts

## Conventions

- Dates in `YYYY-MM-DD` (UTC).
- `agent_runs` counts runs with a non-null `engine_id` (`copilot`, `claude`, `codex`, etc.).
- `filtered_skips` counts runs that the workflow's pre_activation gates rejected (no engine_id, conclusion=cancelled).
- Append-only files: never rewrite history; add a new row each audit.
