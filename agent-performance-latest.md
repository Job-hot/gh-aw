# Agent Performance — 2026-04-23
Run: §24817058307 | Q:68↓3 E:62↓5

## Ecosystem Overview (Apr 23)
- Overall quality: 68/100 (↓-3 from 71), effectiveness: 62/100 (↓-5 from 67)
- 21 runs observed (4 in_progress at analysis time), 17 completed
- Success rate: ~47% (8/17) — significant drop from 90% yesterday
- Root causes: #27888 awf-api-proxy sidecar unhealthy + model not supported + ongoing P1s

## Top Performers
1. **Test Quality Sentinel** (Q:90 E:92) — 2/2 success, ~5min, consistent
2. **Design Decision Gate** (Q:85 E:82) — 2/2 success today (IMPROVED from 50% yesterday)
3. **Smoke OpenCode** (Q:83 E:80) — success ✅, new engine stable
4. **Changeset Generator** (Q:80 E:78) — success, Codex engine working for this workflow
5. **Agent Container Smoke Test** (Q:80 E:82) — success

## Regressed This Run 📉
- **GitHub Remote MCP Authentication Test** (#27965 auto-created today): REGRESSED (was RESOLVED Apr 22). Root cause: `gpt-5.1-codex-mini` model not supported on this subscription tier. Previously worked, now failing again — model availability change suspected.
- **CLI Version Checker** (#27966 auto-created today): docker compose startup failure. Related to #27888 awf-api-proxy sidecar unhealthy.

## Improved This Run 📈
- **Design Decision Gate**: 2/2 success today (was 50% success Apr 22 with push bundle failure)

## Watch / Needs Improvement
- **awf-api-proxy sidecar** (Q:20 E:10) — #27888 blocking Docker-compose-based workflows
- **GitHub Remote MCP Auth Test** (Q:15 E:0) — REGRESSED, model not supported
- **CLI Version Checker** (Q:20 E:5) — container startup failure #27966
- **Smoke Copilot/Claude** (Q:10 E:5) — ongoing P1 (#27028, #27030)
- **Smoke Codex** (Q:10 E:5) — stale lock files #27724 #27731
- **Smoke Crush/Gemini** (Q:10 E:5) — new engines not yet stable

## P0/P1 Active (Apr 23)
- P1: awf-api-proxy sidecar unhealthy (#27888) — blocking multiple workflows
- P1: GitHub Remote MCP Auth REGRESSION — model not supported (gpt-5.1-codex-mini, #27965)
- P1: Stale lock files (#27724 + #27731) — needs `make recompile`
- P1: Safe outputs session not found at 37min (#27755) — infrastructure
- P1: Design Decision Gate max_turns=5 (#27470)
- P1: Smoke Copilot (#27028), Smoke Claude (#27030)

## Issues/Actions This Run
- Discussion created (performance report)
- No new improvement issues created (existing issues #27965, #27966 cover today's failures)

Last updated: 2026-04-23T04:40Z by agent-performance-manager
