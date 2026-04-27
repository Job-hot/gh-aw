# Shared Alerts — 2026-04-27T12:20Z

## P0 (Critical)
- **codex: command not found** (#28596 OPEN, Apr 27 still active): Daily Fact About gh-aw failing with `codex: command not found`. Root cause: codex CLI binary not installed on workflow runners. Fix: install codex binary on runner or switch to different engine. Auto-issue #28703 created today.

## P1 (High)
- **Documentation Unbloat claude auth failure** (#28659 OPEN, #28673 investigation): Claude engine auth termination — OAuth token issue.
- **GitHub Remote MCP Authentication Test** (#27965 OPEN): Day 6+ — model not supported (gpt-5.1-codex-mini). Persistent regression.
- **AWF binary CDN 502** (#28529 OPEN Apr 26): Not observed today — may be resolved.
- **Safe outputs "session not found" at 37min** (#27755 OPEN): Long-running workflows at risk.
- **awf-api-proxy sidecar unhealthy** (#27888 OPEN): Docker compose failures.
- **GitHub App rate limit exhaustion** (#27251 OPEN).
- **CODEX_HOME variable collision** (#27512 OPEN).
- **Smoke Copilot** (#27028 OPEN) + **Smoke Claude** (#27030 OPEN): Not observed today.

## P2 (Watch)
- **THREAT_DETECTION_RESULT parse failure**: Recurring.
- **Safe Outputs SEC-004** (#27235 OPEN).
- **Daily Documentation Updater protected files** (#27801 OPEN).
- **Performance regressions** (#27280/#27279/#27278 OPEN).
- **MCP gateway long-running drops** (#23153 OPEN).

## Trends (Apr 27)
- 204 workflows, 0 missing lock files
- Scheduled success rate ~93% (28/30 runs)
- Daily Community Attribution NOT failing today (improvement vs Apr 26)
- P0 GPU runner issue partially resolved: Daily News + Daily Issues Report Generator working; Daily Fact About gh-aw still failing with codex binary
- Documentation Unbloat new failure continues (claude auth)

## Resolved (Recent)
- Stale lock files ✅ RESOLVED Apr 23
- CLI Version Checker Docker failure ✅ RESOLVED Apr 24
- Daily Community Attribution recurring failure: NOT observed Apr 27 (may be improving)
