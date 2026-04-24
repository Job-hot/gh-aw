# Shared Alerts — 2026-04-24T12:10Z

## P0 (Critical)
- None currently

## P1 (High)
- **dependabot-go-checker compilation failure → Agentic Maintenance** (Apr 23, issue created): `vulnerability-alerts: read` permission not at job level. Fix: move to workflow-level permissions in dependabot-go-checker.md.
- **Daily Community Attribution model not supported** (#28025/#28235 OPEN): Recurring 400 error. Model unavailable for subscription tier. Multiple auto-issues created.
- **Daily Fact About gh-aw MCP Gateway failure** (#28245 NEW Apr 24): `Start MCP Gateway` step failing. New failure mode (was agent failure on Apr 23).
- **Safe outputs "session not found" at 37min** (#27755 OPEN): Long-running workflows affected.
- **Design Decision Gate push bundle failure** (#27756 OPEN).
- **Design Decision Gate max_turns=5** (#27470 OPEN): structurally impossible ADR generation.
- **awf-api-proxy sidecar unhealthy** (#27888 OPEN): Docker compose failures.
- **GitHub Remote MCP Auth Test REGRESSION** (#27965 OPEN): model not supported again.
- **node: command not found on aw-gpu-runner-T4** (#27534 OPEN): Recurring.
- **GitHub App rate limit exhaustion** (#27251 OPEN).
- **CODEX_HOME variable collision** (#27512 OPEN).
- **Smoke Claude** (#27030 OPEN): Ongoing.
- **Smoke Copilot** (#27028 OPEN): Ongoing since Apr 14.

## P2 (Watch)
- **THREAT_DETECTION_RESULT parse failure** (Apr 23 issue created): Duplicate Code Detector + Daily Choice Type Test failing.
- **Safe Outputs SEC-004** (#27235 OPEN).
- **Daily Documentation Updater protected files** (#27801 OPEN).
- **Performance regressions** (#27280/#27279/#27278 OPEN).
- **MCP gateway long-running drops** (#23153 OPEN).
- **Copilot reviewer fan-out** (#27130 OPEN).

## Resolved (Recent)
- Stale lock files ✅ RESOLVED Apr 23 (was 23 stale files Apr 22)
