# Workflow Health — 2026-05-27T05:54Z

Score: 90/100 (stable). ~236 workflows. Run: §26493510814

## KEY FINDINGS

### Status (May 27)
- **Compilation:** 236/236 workflows have lock files (100% ✅)
- **Scheduled runs analyzed:** 500 runs (149 unique workflows active)
- **Failing (>=50% rate, 2+ runs):** 31 workflows
- **Active failures tracked:** 7 P1 (100% fail), 6 P2 moderate

### Critical Issues (P0/P1) 🚨
- **CGO build failures** (#35028): 37% fail rate on push — ongoing
- **PR Sous Chef** (#35073, #35067, #35052): 21% fail — safe outputs permission + pre_activation setup
- **Daily News / Daily Issues Report Generator**: 100% fail (2/2) — Copilot CLI execution failure
- **Daily Fact About gh-aw**: 100% fail (2/2) — Codex CLI execution failure
- **Super Linter Report**: 100% fail (2/2) — super-linter infrastructure
- **Daily Agentic Workflow Token Usage Audit**: 100% fail (2/2) — setup job failure
- **Daily AW Cross-Repo Compile Check**: 100% fail (2/2) — cache-memory git failure
- **Daily Community Attribution Updater** (#35105): repo-memory patch size exceeded

### Moderate Issues (P2)
- Step Name Alignment: 67% fail (#35135 cache_memory miss)
- Agentic Maintenance: 26% fail (6/23) — no open issue
- Contribution Check: 25% fail (3/12) — safe outputs failure
- Avenger: 12% fail (4/32) — intermittent

### Actions Taken This Run
- Created health dashboard issue (Workflow Health Dashboard — 2026-05-27)
- No new P0/P1 issues created — all problems tracked in existing issues
- All 100% failing daily workflows root causes identified but no issues filed

### Trends
- Score: 90/100 (stable, improved estimate methodology)
- CGO: ongoing (P0 #35028)
- Safe outputs permission blocking: systemic pattern
- 87/236 workflows inactive (no recent runs)

Last updated: 2026-05-27T05:54:00Z by workflow-health-manager
