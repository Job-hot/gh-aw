# Agent Performance — 2026-04-24
Run: §24872587802 | Q:72↑4 E:68↑6

## Ecosystem Overview (Apr 24)
- Overall quality: 72/100 (↑+4 from 68), effectiveness: 68/100 (↑+6 from 62)
- 30 runs observed: 13 success, 1 failure, 12 cancelled (bulk CI push), 4 in_progress
- Effective success rate: 93% (13/14 non-cancelled completed) — major improvement from 47% yesterday
- Cancellations were bulk (same-timestamp CI push, not individual failures)

## Top Performers
1. **Test Quality Sentinel** (Q:90 E:92) — 2/2 success, consistent
2. **Design Decision Gate** (Q:88 E:85) — 2/2 success, continuing improvement streak
3. **CLI Version Checker** (Q:80 E:82) — RECOVERED ✅ (was failing Apr 23 due to docker)
4. **Auto-Triage Issues** (Q:80 E:80) — success
5. **Agent Persona Explorer** (Q:78 E:75) — success
6. **Documentation Unbloat** (Q:78 E:75) — success
7. **Daily AW Cross-Repo Compile Check** (Q:75 E:75) — success

## Regressed / Still Failing 📉
- **GitHub Remote MCP Authentication Test** (Q:10 E:0) — STILL failing (#27965 P1). Persistent model unavailability (gpt-5.1-codex-mini not supported).

## Improved This Run 📈
- **CLI Version Checker**: RECOVERED (success today vs failure Apr 23 #27966)
- **Overall success rate**: 93% vs 47% yesterday — major improvement
- **Bulk cancellations**: Not actual failures; caused by CI push batching

## P1 Active (Apr 24)
- **GitHub Remote MCP Auth Test** (#27965): Persistent model not supported
- **awf-api-proxy sidecar** (#27888): Was blocking, today's docker-based runs appear cancelled (not tested)
- **Smoke Copilot** (#27028), **Smoke Claude** (#27030): Not run today (cancelled)
- **Daily Community Attribution** (#28025): Model unavailable
- **Safe outputs session not found 37min** (#27755)
- **Design Decision Gate max_turns=5** (#27470)
- **dependabot-go-checker compilation** (#aw_deplck)

## Issues/Actions This Run
- Discussion created (performance report)
- No new improvement issues (existing issues cover open P1s)

Last updated: 2026-04-24T04:50Z by agent-performance-manager
