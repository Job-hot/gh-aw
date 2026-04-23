# Workflow Health — 2026-04-23T12:10Z

Score: 68/100 (↓-1 from 69 Apr 22). 200 workflows. Run: §24834349371

## KEY FINDINGS

### Compilation Status
- 200/200 lock files present ✅
- **0 stale lock files** ✅ (resolved! was 23 yesterday)
- **1 compilation error**: dependabot-go-checker.md — `vulnerability-alerts` permission not allowed at job level in GitHub Actions schema → breaks Agentic Maintenance (P1, new issue created)

### P0 Issues
- None today

### P1 Issues (Active)
- **dependabot-go-checker compilation failure** (NEW, Apr 23, issue created): `vulnerability-alerts: read` permission not allowed at job level → Agentic Maintenance fails (run §24831492077)
- **Daily Community Attribution model not supported** (#28025 OPEN): 400 model not supported error
- **Design Decision Gate push bundle failure** (#27756 OPEN): push_to_pull_request_branch failed
- **Safe outputs "session not found" at 37min** (#27755 OPEN): MCP server session expires
- **Smoke Claude** (#27030 OPEN): Ongoing
- **Smoke Copilot** (#27028 OPEN): Ongoing
- **awf-api-proxy sidecar unhealthy** (#27888, from agent-perf): Docker compose failures
- **GitHub Remote MCP Auth Test REGRESSION** (#27965, from agent-perf): gpt-5.1-codex-mini not supported
- **node not found on GPU runner** (#27534 OPEN): Recurring
- **GitHub App rate limit exhaustion** (#27251 OPEN)
- **CODEX_HOME variable collision** (#27512 OPEN)
- **Design Decision Gate max_turns=5** (#27470 OPEN)

### P2 Issues
- **THREAT_DETECTION_RESULT parse failure** (NEW, Apr 23, issue created): Duplicate Code Detector + Daily Choice Type Test failing in detection job — detection model not outputting expected format
- **Daily Fact About gh-aw** (#28035 OPEN, auto): agent failure Apr 23
- **Safe Outputs SEC-004** (#27235 OPEN): 4 handler files
- **Daily Documentation Updater protected files** (#27801 OPEN)
- **Performance regressions** (#27280/#27279/#27278 OPEN)
- **MCP gateway long-running drops** (#23153 OPEN)

### Today's Run Stats (30 scheduled runs)
- Success: 25 (83%)
- Failures: 5 (17%)
  - Agentic Maintenance: compile error (dependabot-go-checker)
  - Duplicate Code Detector: THREAT_DETECTION parse
  - Daily Choice Type Test: THREAT_DETECTION parse
  - Daily Fact About gh-aw: agent failure → #28035
  - Daily Community Attribution: model not supported → #28025

## Open Issues (workflow-health related)
- #aw_deplck dependabot-go-checker compilation / Agentic Maintenance (P1, NEW today)
- #aw_thrdet THREAT_DETECTION parse failure (P2, NEW today)
- #28035 Daily Fact About gh-aw failed (auto, today)
- #28025 Daily Community Attribution model not supported (auto, today)
- #27756 Design Decision Gate push bundle failure (P1)
- #27755 Safe outputs session not found 37min (P1)
- #27888 awf-api-proxy sidecar unhealthy (P1)
- #27965 GitHub Remote MCP Auth Test regression (P1)
- #27724 Agentic workflows out of sync (P1) — stale locks resolved, now tracking compile error
- #27512 CODEX_HOME variable collision (P1)
- #27470 Design Decision Gate max_turns=5 (P1)
- #27534 Daily Issues Report GPU/node (P1)
- #27251 Rate limit exhaustion co-scheduled (P1)
- #27030 Smoke Claude (P1)
- #27028 Smoke Copilot (P1)
- #27235 Safe Outputs SEC-004 (P2)
- #23153 MCP gateway session drops (P2)

## Engine/Tool Status
- Copilot: mostly ✅ (Smoke Copilot ongoing #27028, model not supported for community-attribution)
- Claude: mostly ✅ (Smoke Claude #27030 ongoing)
- Codex: ⚠️ (was blocked by stale locks — now resolved, but awf-api-proxy issue #27888)
- MCP Gateway: ⚠️ session timeout at ~37min (#27755)
- awf-api-proxy: ❌ sidecar unhealthy (#27888)

## Actions This Run
- Created issue for dependabot-go-checker P1 (new compilation failure)
- Created issue for THREAT_DETECTION recurring parse failures (P2)
- Added comment to #27724 with stale lock resolution + new compile error

Last updated: 2026-04-23T12:10Z by workflow-health-manager
