# Workflow Health — 2026-04-24T12:10Z

Score: 72/100 (↑+4 from 68 Apr 23). 201 workflows. Run: §24888666710

## KEY FINDINGS

### Compilation Status
- 201/201 lock files present ✅
- **0 stale lock files** ✅
- **1 known compilation error**: dependabot-go-checker.md — `vulnerability-alerts` permission not allowed at job level (P1, issue created Apr 23 — still open)

### Today's Run Stats (30 scheduled runs, Apr 24)
- Success: 28 (93%)
- Failures: 2 (7%)
  - **Daily Fact About gh-aw**: Start MCP Gateway step failed → #28245 auto-created
  - **Daily Community Attribution**: model not supported (400) → #28235 auto-created (recurring #28025)

### P0 Issues
- None today

### P1 Issues (Active)
- **dependabot-go-checker compilation failure** (#aw_deplck, Apr 23 issue created): `vulnerability-alerts: read` not allowed at job-level
- **Daily Community Attribution model not supported** (#28025/#28235 OPEN): Recurring 400 error — model unavailable
- **Daily Fact About gh-aw MCP Gateway failure** (#28245 NEW Apr 24): `Start MCP Gateway` step failing — different failure from yesterday
- **Safe outputs "session not found" at 37min** (#27755 OPEN): Long-running workflows at risk
- **Design Decision Gate push bundle failure** (#27756 OPEN)
- **Smoke Claude** (#27030 OPEN): Ongoing
- **Smoke Copilot** (#27028 OPEN): Ongoing since Apr 14
- **awf-api-proxy sidecar unhealthy** (#27888 OPEN): Docker compose failures
- **GitHub Remote MCP Auth Test REGRESSION** (#27965 OPEN): model not supported
- **node not found on GPU runner** (#27534 OPEN): Recurring
- **GitHub App rate limit exhaustion** (#27251 OPEN)
- **CODEX_HOME variable collision** (#27512 OPEN)
- **Design Decision Gate max_turns=5** (#27470 OPEN)

### P2 Issues
- **THREAT_DETECTION_RESULT parse failure** (#aw_thrdet, Apr 23 issue created): Recurring
- **Daily Fact About gh-aw** (#28035 Apr 23): Separate auto-issue from yesterday
- **Safe Outputs SEC-004** (#27235 OPEN)
- **Daily Documentation Updater protected files** (#27801 OPEN)
- **Performance regressions** (#27280/#27279/#27278 OPEN)
- **MCP gateway long-running drops** (#23153 OPEN)

## Open Issues (workflow-health related)
- #28245 Daily Fact About gh-aw MCP Gateway failure (NEW Apr 24)
- #28235 Daily Community Attribution model not supported (NEW Apr 24 — recurring)
- #28025 Daily Community Attribution model not supported (Apr 23, prior instance)
- #28035 Daily Fact About gh-aw (auto, Apr 23)
- P1/P0 list above + prior issues
