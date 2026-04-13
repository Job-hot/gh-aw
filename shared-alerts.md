# Shared Alerts — 2026-04-13T12:12Z

## P2 (High)
- **Smoke Claude schedule failure** (ongoing, #25727): Still failing on daily schedule, but PASSES on PR-triggered runs. Environment-specific issue — monitor.
- **Smoke Cross-Repo PR Create/Update** (#25221, #25217, Apr 8): Both still failing. Persistent — no fix yet.
- **Documentation Unbloat inconsistent** (ongoing): Claude workflow ~$55/week, inconsistent 50% success.
- **Daily Semgrep Scan** (new fail Apr 13, #aw_semgrep1): 0/1 success Apr 13 — issue created this WH run.
- **GitHub Remote MCP Auth Test** (persistent): 100% failure — #24829 closed not_planned, test still failing.
- **Daily Issues Report recurring failure** (#25265, #25503): Copilot agent crash pattern.

## P3 (Watch)
- **Smoke Gemini** (#25216): 100% failure (Gemini CLI 0.37.0 compat). Open issue.
- **Daily Firewall Logs** (#25456): safe_outputs process failure.

## Copilot Version Status
- v1.0.21 ACTIVE (current in production)
- Issue #25978: CLI bump tracking (Copilot 1.0.24, Claude Code 2.1.104, Codex 0.120.0, MCP Gateway v0.2.18) — open, not yet PRed

## Recoveries / Fixes (Apr 11-13)
- ✅ Smoke Copilot: RECOVERED — passing scheduled runs
- ✅ Contribution Check: Now passing (was report_incomplete)
- ✅ 20 PRs merged by Copilot bot (OTel, security, workflow fixes)
- ✅ Fixed detection squid crash, multiple agent assignments, push_repo_memory gating

## Note on Stale Shared Alert (resolved)
- "#25548 DDG (Design Decision Gate)" reference was INCORRECT
- Issue #25548 is actually "feat: collect Docker operational logs on failure for AWF diagnostics" (enhancement)
- Design Decision Gate tracking may be separate — not a current P2 issue

## Ecosystem State
- ~187 compiled workflows. Health: ~74/100 (→ stable Apr 13 12:12Z)
- Engine split: ~124 copilot, ~41 claude, ~18 codex, ~4 others
- v1.0.21 currently active

## Orchestrator Summaries (Apr 13)
- Agent Performance (Apr 13 04:47Z): Q:73↑3 E:65↑5. Smoke Copilot recovered. Smoke Claude still failing on schedule.
- Workflow Health (Apr 13 12:12Z): Score 74/100 (stable). Semgrep new failure. API rate-limited during run.
- Campaign Manager (last known: Mar 16 17:41Z): P0/#20315 resolved likely

Last updated: 2026-04-13T12:12Z by workflow-health-manager
