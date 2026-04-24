# Shared Alerts — 2026-04-23T12:10Z

## P0 (Critical)
- None currently

## P1 (High)
- **dependabot-go-checker compilation failure → Agentic Maintenance broken** (NEW Apr 23, #aw_deplck): `vulnerability-alerts: read` permission rejected by GitHub Actions schema validator → `agentic-maintenance` workflow fails to compile workflows step (199/200, 1 error). Fix: remove or move `vulnerability-alerts: read` from job-level permissions in `dependabot-go-checker.md`.
- **Daily Community Attribution model not supported** (#28025 NEW Apr 23): 400 error, gpt model unavailable for subscription tier. Recurring model availability issue.
- **Safe outputs "session not found" at 37min** (#27755 OPEN, Apr 22): MCP server returning session not found at 37min. All long-running workflows at risk.
- **Design Decision Gate push bundle failure** (#27756 OPEN, Apr 22): push_to_pull_request_branch failed.
- **Design Decision Gate max_turns=5** (#27470 OPEN): structurally impossible ADR generation.
- **awf-api-proxy sidecar unhealthy** (#27888 OPEN, Apr 23): Docker compose failures, CLI Version Checker and potentially other workflows blocked.
- **GitHub Remote MCP Auth Test REGRESSION** (#27965, Apr 23): Was resolved Apr 22, failing again. gpt-5.1-codex-mini model not supported.
- **node: command not found on aw-gpu-runner-T4** (#27534 OPEN): Recurring.
- **GitHub App rate limit exhaustion** (#27251 OPEN): Co-scheduled at 23:44 UTC.
- **CODEX_HOME variable collision** (#27512 OPEN): cp same-file error.
- **Smoke Claude** (#27030 OPEN): Ongoing.
- **Smoke Copilot** (#27028 OPEN): Ongoing since Apr 14.

## P2 (Watch)
- **THREAT_DETECTION_RESULT parse failure** (NEW Apr 23, #aw_thrdet): Duplicate Code Detector + Daily Choice Type Test — detection model not outputting expected JSON format. Recurring across at least 2 days.
- **Daily Fact About gh-aw** (#28035 OPEN, Apr 23): agent failure, auto-issue.
- **Safe Outputs SEC-004** (#27235 OPEN): 4 handler files need sanitization.
- **Daily Documentation Updater protected files** (#27801 OPEN): Tried to modify .github/aw/ files.
- **Performance regressions** (#27280/#27279/#27278 OPEN).
- **MCP gateway long-running drops** (#23153 OPEN): Session not found after 30-45min.
- **Copilot reviewer fan-out** (#27130 OPEN): 6 review workflows per push.

## Resolved (Recent)
- Stale lock files ✅ RESOLVED Apr 23 (was 23 stale files Apr 22 — all in sync now)
- Design Decision Gate push failure ✅ IMPROVED Apr 23 (2/2 success, was 50% failure Apr 22)
- Smoke OpenCode ✅ NEW engine working

## Ecosystem State
- 200 workflows
- Stale lock files: 0 ✅ (resolved from 23)
- Compilation errors: 1 (dependabot-go-checker)
- Schedule success rate: ~83% today (25/30 runs)
- P0 failures: 0
- P1 failures: Agentic Maintenance, Smoke Copilot/Claude, awf-api-proxy, session timeout
- Overall quality trend: Q:68 (↓-1 from 69)

## Orchestrator Summaries
- Workflow Health (Apr 23 12:10Z): Score 68/100. 200 workflows. 0 stale locks ✅. 83% success (25/30). NEW P1: dependabot-go-checker compile error. NEW P2: THREAT_DETECTION parse failure.
- Agent Performance (Apr 24 04:50Z): Q:72 E:68. 93% success (13/14). CLI Version Checker RECOVERED. GitHub Remote MCP Auth still failing (#27965).
- Workflow Health (Apr 22 12:11Z): Score 69/100. 197 workflows. 23 stale locks. 90% success (27/30). Protected files P2. Codex 401 P1.
- Agent Performance (Apr 22 04:37Z): Q:71 E:67. 18 workflows, 29 runs. Safe outputs 37min P1. DDG push failure P1.

Last updated: 2026-04-23T12:10Z by workflow-health-manager
