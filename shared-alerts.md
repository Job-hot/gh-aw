# Shared Alerts — 2026-04-17T04:39Z

## P0 (Critical)
- **Daily Issues Report Generator** (#26393, OPEN): 10+ consecutive day failure. node:command not found in Copilot runner. No fix yet.
- **Smoke Gemini** (#26351 OPEN): 14+ day 100% failure. Gemini API proxy handler crash. Gemini 0.38.0 available - may fix.

## P1 (High)
- **Smoke Claude** (#26777, #26790 OPEN): ~60% failure rate on schedule runs. Environment context divergence.
- **Smoke Codex** (#26767 OPEN): Recent failure.
- **Smoke OpenCode** (#26765 OPEN): Recent failure.
- **Super Linter Report**: EACCES permission denied on super-linter.log upload. Structural bug.
- **Smoke Cross-Repo PR Create/Update** (#25221, #25217): stale 8+ days. No fix applied.
- **Daily Firewall Logs** (#25456): safe_outputs process failure.
- **Daily Observability Report for AWF Firewall and MCP Gateway** (#26761 OPEN): Recent failure.

## P2 (Watch)
- **Contribution Check**: 50% success rate, improving but inconsistent
- **PR Triage Agent** (#26778 OPEN): 67% success rate
- **Auto-Triage Issues** (#26364 OPEN): 67% success rate, intermittent
- **MCP Rate Limit** (#26239 OPEN): Circuit breaker needed.
- **GitHub MCP get_me 403 errors** (#26458 OPEN): Auth/permission failures.
- **AI Moderator** (#26793 OPEN): Recent failure.
- **[aw] Failure Investigator (6h)**: 50% success rate (2 runs).

## Copilot Version Status
- v1.0.27 AVAILABLE (#26803 OPEN for upgrade — was #26367)
- v1.0.21 ACTIVE (current in production as of Apr 17)
- Claude Code 2.1.109 available
- Codex 0.120.0 available
- Gemini 0.38.0 available — may fix Smoke Gemini

## Recoveries (Apr 16-17)
- ✅ GitHub Remote MCP Auth Test: recovered (1 run 100% success Apr 17)
- ✅ Documentation Unbloat: 100% success Apr 17
- ✅ CLI Version Checker: 100% success Apr 17

## Ecosystem State
- 82 unique scheduled workflows active. Success rate: 83.3% (↑ from ~76% Apr 16)
- P0 failures: 2 (unchanged since Apr 12)
- P1 failures: 5 ongoing
- Overall quality trend: ↑ improving

## Orchestrator Summaries
- Agent Performance (Apr 17 04:39Z): Q:77 E:75. 83.3% ecosystem success rate. Top: Agentic Maintenance (100%), Issue Monster (95%). P0 unchanged.
- Workflow Health (Apr 16 12:10Z): Score 71/100. 191 workflows. 16 stale lock files.

Last updated: 2026-04-17T04:39Z by agent-performance-manager
