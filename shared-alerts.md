# Shared Alerts — 2026-04-14T04:37Z

## P2 (High)
- **Smoke Claude schedule failure** (ongoing, #25727): Still failing on daily schedule, passes on PR runs. Environment-specific — schedule vs PR context divergence.
- **Smoke Cross-Repo PR Create/Update** (#25221, #25217, Apr 8): STALE 6 days. No fix applied. Needs escalation.
- **Daily Semgrep Scan** (new fail Apr 13): 0/1 success — security scan degraded. P2.
- **Documentation Unbloat inconsistent** (ongoing): ~$55/week Claude, 50% success. Cost gate needed.
- **Daily Issues Report recurring failure** (#25265, #25503): Copilot agent crash pattern.
- **GitHub Remote MCP Auth Test** (persistent): 100% failure — #24829 closed not_planned. Test still failing.

## P3 (Watch)
- **Smoke Gemini** (#25216): 100% failure. Gemini 0.37.2 now available (see #26158) — may fix after CLI bump.
- **Daily Firewall Logs** (#25456): safe_outputs process failure.

## Copilot Version Status
- v1.0.25 TRACKED (new --remote/--no-remote flags; see #26158)
- v1.0.21 ACTIVE (current in production as of Apr 14)
- Claude Code 2.1.105 tracked in #26158
- Codex 0.120.0 tracked in #26158
- Gemini 0.37.2 tracked in #26158

## Recoveries (Apr 11-14)
- ✅ Smoke Copilot: RECOVERED
- ✅ Contribution Check: RECOVERED
- ✅ 20 PRs merged by Copilot bot (OTel, security, workflow fixes)

## Ecosystem State
- ~187 compiled workflows. Health: 74/100 (→ stable Apr 14 04:37Z)
- Engine split: ~124 copilot, ~41 claude, ~18 codex, ~4 others
- v1.0.21 currently active

## Orchestrator Summaries (Apr 14)
- Agent Performance (Apr 14 04:37Z): Q:74↑1 E:66↑1. CLI Version Checker standout (4 upgrades). Weekly discussion created.
- Workflow Health (Apr 13 12:12Z): Score 74/100 (stable). Semgrep new failure.
- Campaign Manager (last known: Mar 16 17:41Z): Status unknown — no recent update

Last updated: 2026-04-14T04:37Z by agent-performance-analyzer
