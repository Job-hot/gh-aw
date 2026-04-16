# Shared Alerts — 2026-04-16T12:10Z

## P0 (Critical)
- **Daily Issues Report Generator** (#26393, OPEN): 10+ consecutive day failure. node:command not found in Copilot runner. No fix yet.
- **Smoke Gemini** (#26351 OPEN, #26456 deep-report OPEN): 14+ day 100% failure. Gemini API proxy handler crash. Investigation underway.

## P1 (High)
- **Smoke Claude** (no issue): ~60% failure rate on schedule runs (3/5 recent failed). PR runs succeed. Environment context divergence. Monitor closely.
- **Smoke Cross-Repo PR Create/Update** (#25221, #25217): stale 8+ days. No fix applied.
- **Daily Firewall Logs** (#25456): safe_outputs process failure.
- **Schema Feature Coverage Checker** (#25992): Protected-files config blocks PR creation.
- **GitHub Remote MCP Auth Test**: 100% failure — #24829 closed not_planned. Test still failing.

## P2 (Watch)
- **MCP Rate Limit** (#26239 OPEN): Circuit breaker needed. Risk of cascading failures.
- **GitHub MCP get_me 403 errors** (#26458 OPEN): Authentication/permission failures.
- **~15 intermittent workflow failures** today (Apr 16): AI Moderator, Daily Go Function Namer, Instructions Janitor, etc. — likely Copilot engine flakiness.

## Copilot Version Status
- v1.0.27 AVAILABLE (PR #26367 OPEN for upgrade)
- v1.0.21 ACTIVE (current in production as of Apr 16)
- Claude Code 2.1.109 available
- Codex 0.120.0 available
- Gemini 0.38.0 available — may fix Smoke Gemini after CLI bump

## Recoveries (Apr 15-16)
- ✅ Auto-Triage Issues: recovered (failed 01:07Z Apr 16, success 07:04Z)
- ✅ Contribution Check: running successfully

## Ecosystem State
- 191 compiled workflows. Health: 71/100 (↓1 Apr 16 12:10Z)
- 16 stale lock files (down from 18, likely checkout artifact)
- P0 failures: 2 (unchanged since Apr 12)
- P1 failures: 4 ongoing

## Orchestrator Summaries (Apr 16)
- Agent Performance (last: Apr 15 04:37Z): 5 tools upgraded via CLI Version Checker (#26367 open)
- Workflow Health (Apr 16 12:10Z): Score 71/100 (↓1). 191 workflows, 2 P0 failures. Dashboard: #aw_whdapr16

Last updated: 2026-04-16T12:10Z by workflow-health-manager
