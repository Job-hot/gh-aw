# Workflow Health — 2026-04-27T12:20Z

Score: 74/100 (↑+1 from 73 Apr 26). 204 workflows. Run: §24994541064

## KEY FINDINGS

### Compilation Status
- 204/204 lock files present ✅
- **0 missing lock files** ✅
- **0 stale lock files** ✅ (increased from 203 to 204 due to new workflow)

### Today's Failures (Apr 27)
- **Daily Fact About gh-aw**: `codex: command not found` — issue #28703 auto-created. Still part of P0 #28596 (aw-gpu-runner-T4 / codex binary missing).
- **Documentation Unbloat**: Claude engine terminated (OAuth token/auth issue) — issue #28659. Covered by failure investigation #28673.

### Scheduled Run Summary (Apr 27)
- 30 scheduled runs executed: 28 success, 1 failure, 1 action_required
- Effective success rate: ~93% (28/30 excl action_required)
- Good recovery from P0 GPU runner: Daily News and Daily Issues Report Generator now running on different runners ✅

### P0 Issues (Active)
- **codex: command not found** (#28596 OPEN): Daily Fact About gh-aw still failing with codex engine — `codex` binary not found on runners. Auto-issue #28703 created today. Root cause: codex CLI not installed on runners used by this workflow.

### P1 Issues (Active/Ongoing)
- **Documentation Unbloat claude auth failure** (#28659 OPEN, #28673 investigation): OAuth token issue with Claude engine in this workflow.
- **AWF binary CDN 502** (#28529 OPEN): Intermittent, not observed today
- **Daily Community Attribution model not supported** (#28025/#28235 OPEN): NOT observed today ✅
- **Safe outputs "session not found" at 37min** (#27755 OPEN)
- **Smoke Claude** (#27030 OPEN) + **Smoke Copilot** (#27028 OPEN): Not observed today
- **awf-api-proxy sidecar unhealthy** (#27888 OPEN)
- **GitHub Remote MCP Auth Test REGRESSION** (#27965 OPEN): Day 6+, still failing
- **GitHub App rate limit exhaustion** (#27251 OPEN)
- **CODEX_HOME variable collision** (#27512 OPEN)

### P2 Issues (Watch)
- **THREAT_DETECTION_RESULT parse failure**: Recurring
- **Safe Outputs SEC-004** (#27235 OPEN)
- **Daily Documentation Updater protected files** (#27801 OPEN)
- **Performance regressions** (#27280/#27279/#27278 OPEN)
- **MCP gateway long-running drops** (#23153 OPEN)

## Issues Created This Run
- None (auto-issues already created by failure-investigator)

## Issues Updated
- None (existing issues tracking active failures)

## Positive Trends
- Daily Community Attribution NOT failing today (was recurring P1)
- Daily Go Function Namer success (CDN 502 resolved)
- 28/30 scheduled runs successful
