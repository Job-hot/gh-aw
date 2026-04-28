# Shared Alerts — 2026-04-28T12:20Z

## P0 (Critical)
- **Daily Fact About gh-aw codex crash** (ongoing, daily auto-issues): codex engine agent failure. Recurring every day.

## P1 (High)
- **THREAT_DETECTION_RESULT parse failure** — ESCALATED TODAY: ≥3 workflows failing with "No THREAT_DETECTION_RESULT found in detection log". Was 1-2 workflows on Apr 27, now ≥3 on Apr 28. Tracked in #28866. May indicate detection model degradation.
- **Documentation Unbloat claude auth failure** (#28659 OPEN): Claude OAuth token issue. Recurring.
- **GitHub Remote MCP Authentication Test** (#27965 OPEN): Day 7+ of model-not-supported error.
- **Safe outputs session not found** (#23153 OPEN): Long-running workflows at risk.
- **awf-api-proxy sidecar unhealthy** (#27888 OPEN): Docker compose failures.
- **GitHub App rate limit exhaustion** (#27251 OPEN).
- **CODEX_HOME variable collision** (#27512 OPEN).

## P2 (Watch)
- **Safe Outputs SEC-004** (#27235 OPEN).
- **Daily Documentation Updater protected files** (#27801 OPEN).
- **Performance regressions** (#27280/#27279/#27278 OPEN).
- **MCP gateway long-running drops** (#23153 OPEN).

## Trends (Apr 28)
- 204 workflows, 0 missing lock files
- Scheduled success rate: 57% (17/30) — REGRESSION from 93% yesterday
- THREAT_DETECTION_RESULT failures spreading to new workflows (systemic P1+)
- CI integration tests failing (4 jobs)

## Resolved (Recent)
- #28596 (aw-gpu-runner-T4 node not found for Daily News / Daily Issues Report) — CLOSED Apr 27
  - Note: Daily Fact About gh-aw still failing separately with codex engine
