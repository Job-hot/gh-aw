# Shared Alerts — 2026-06-11T06:07Z

## P1 (High) 🚨
- **AI Credits Cluster June 11** (#aw_aic11, filed Jun 11 OPEN): Code Simplifier, Ubuntu Actions Image Analyzer, Workflow Skill Extractor (all 1K max hit), Matt Pocock (5K 24h limit). Fix: raise max-ai-credits to 2000. DO NOT RE-FILE individuals already tracked: #38499, #38500, #38501, #38497.
- **Daily News Node.js chroot** (#38379, Jun 10 OPEN): Node.js unreachable in AWF chroot (exit 127). Assigned @zarenner. DO NOT RE-FILE.
- **Daily Safe Outputs Git Simulator** (#aw_gitsim10, Jun 10 OPEN): memory/git-simulator branch missing signed-commit seed. DO NOT RE-FILE.

## P2 (Watch) ⚠️
- **jsweep tool denial** (#38505, Jun 11): Agent attempted `rm -rf /tmp/x`; guardrail fired (2/2). One-off, monitor.
- **agentic workflows out of sync** (#38274, Jun 10, CLOSED): maintenance resolved.

## Resolved (Jun 9→11) ✅
- #38119 AI credits cluster tracker — CLOSED (old tracker)
- #38389 Failure cascade — CLOSED
- Jun 10 bulk individual failures — most CLOSED on credits reset
- Prior P0/P1 issues from Jun 9-10 all resolved

## Systemic Notes
- **Health score trend:** 68→83→87→85 (slight dip, AI credits cluster persists)
- **AI credits persistent pattern:** Analysis-heavy workflows at 1K limit routinely exceed. Fix needed: raise to 2K. Third day of same pattern.
- **memory/* bootstrap:** New memory branches require signed-commit for first push. Manual intervention required.
- **Daily News AWF chroot:** Node.js PATH resolution failing; fix: bind-mount node runtime. @zarenner assigned.
- **Smoke tests:** Some smoke test failures (Gemini #38515, Pi #38513, Antigravity #38512) — likely from PR testing, not systemic.

## Do Not Re-File (Active Issues)
- #aw_aic11: AI Credits Cluster June 11 (4 workflows)
- #aw_gitsim10: Daily Safe Outputs Git Simulator (memory branch needs seed)
- #38379: Daily News chroot (assigned @zarenner)
- #38499, #38500, #38501, #38497: AI credits individual failures — covered by #aw_aic11
- #38505: jsweep tool denial
