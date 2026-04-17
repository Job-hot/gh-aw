# Shared Alerts — 2026-04-17T12:10Z

## P0 (Critical)
- **Daily Fact About gh-aw** (#26852, OPEN TODAY): 10/10 consecutive failures (Codex engine). Auto-issue created today.

## P1 (High)
- **Daily Community Attribution** (#26848, OPEN TODAY): 50% failure rate. Copilot engine crash during README edit.
- **Smoke Claude** (#26777, #26790 OPEN): Last schedule run Apr 14 = failure. ~60% failure rate. Monitor.
- **Super Linter Report**: EACCES permission denied (structural bug).
- **Smoke Cross-Repo PR Create/Update** (#25221, #25217): stale 9+ days.
- **Daily Firewall Logs** (#25456): safe_outputs process failure.
- **Daily Observability Report for AWF Firewall and MCP Gateway** (#26761 OPEN): Recent failure.

## P2 (Watch)
- **PR Triage Agent** (#26778 OPEN): 67% success rate.
- **Auto-Triage Issues** (#26364 OPEN): 67% success rate, intermittent.
- **MCP Rate Limit** (#26239 OPEN): Circuit breaker needed.
- **GitHub MCP get_me 403 errors** (#26458 OPEN): Auth/permission failures.

## Copilot Version Status
- v1.0.27 AVAILABLE (#26803 OPEN for upgrade)
- v1.0.21 ACTIVE (production as of Apr 17)
- Claude Code 2.1.109 available
- Codex 0.120.0 available
- Gemini: API key invalid (issues closed not_planned Apr 17 by pelikhan)

## Recoveries (Apr 17)
- ✅ Daily Issues Report Generator (#26393): Closed not_planned by pelikhan
- ✅ Smoke Gemini (#26351): Closed not_planned by pelikhan
- ✅ Smoke Codex: 1/1 success Apr 14 (recovered)
- ✅ Compilation: 0 stale lock files (was 16, now fully clean)

## Ecosystem State
- 194 workflows total (+3 from Apr 16: 191 → 194)
- 194/194 lock files present, 0 stale
- Recent schedule success rate: ~93% (27/29 in sample)
- P0 failures: 1 (daily-fact, down from 2)
- P1 failures: 2 ongoing (smoke-claude, daily-community-attribution)
- Overall quality trend: ↑ improving

## Orchestrator Summaries
- Workflow Health (Apr 17 12:10Z): Score 73/100. 194 workflows. 0 stale lock files. daily-fact P0.
- Agent Performance (Apr 17 04:39Z): Q:77 E:75. 83.3% ecosystem success rate.

Last updated: 2026-04-17T12:10Z by workflow-health-manager
