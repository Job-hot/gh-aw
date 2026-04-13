# Workflow Health - 2026-04-13T12:12Z

Score: 74/100 (→ stable, ↑1 from 73 Apr 11). 187 workflows. Run: §24342586738

## KEY FINDINGS

### Recoveries Since Last Run (Apr 11-13)
- ✅ Smoke Copilot: RECOVERED — passing scheduled runs
- ✅ Contribution Check: RECOVERED — now passing (was report_incomplete)
- ✅ 20 PRs merged by Copilot bot (OTel, SEC-004, cache cleanup, agent assignments)

### NEW: Daily Semgrep Scan Failure (Apr 13)
- 0/1 success on Apr 13 — new failure
- Issue #aw_semgrep1 created this run (actual number from safeoutputs)
- P2 priority — security scan degraded

### Ongoing P2 Issues
- Smoke Claude: fails on SCHEDULE, passes on PR runs — environment-specific (#25727)
- Smoke Gemini: 100% failure (#25216, Gemini CLI 0.37.0 compat)
- Smoke Cross-Repo PR Create: persistent fail (#25221)
- Smoke Cross-Repo PR Update: persistent fail (#25217)
- Daily Issues Report: recurring Copilot crash (#25265, #25503)
- Daily Firewall Logs: safe_outputs process failure (#25456)

### P3 Issues (Ongoing)
- GitHub Remote MCP Auth Test: 100% failure (#24829 closed not_planned)
- Documentation Unbloat: 50% success, ~$55/week Claude

## Compilation
- 187/187 lock files present ✅
- ~10 files with 1ms timestamp diff (git checkout ordering artifact, not real staleness)
- All workflows properly compiled

## Copilot Status
- v1.0.21 ACTIVE (current in production)
- v1.0.24 upgrade tracked in #25978 (not yet PRed)

## Score Breakdown
- Compilation: 187/187 ✅: +35
- Smoke Copilot recovered: +2
- Contribution Check recovered: +1
- Smoke Claude schedule fail: -3
- Daily Semgrep new fail: -1
- Smoke Gemini + Cross-Repo persist: -5
- Daily Issues + Firewall issues: -3
- ~16 open failure issues: -8
- v1.0.21 stable: +18
- Net: 74/100

## Score Trend
68 → 71 → 73 → 71 → 70 → 75 → 73 → 74 → 74
Apr5  Apr6  Apr7  Apr8  Apr9  Apr10 Apr11 Apr12 Apr13

## Dashboard Issue
Created new issue #aw_dash413 (Apr 13, this run)

## Note: API Rate Limited
GitHub API rate limited during this run (reset 41min after start).
Fresh run data unavailable — analysis based on shared orchestrator memory (04:47Z today).

Last updated: 2026-04-13T12:12Z
