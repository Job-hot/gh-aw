# Agent Performance — 2026-05-16
Run: §25962522996 | Q:74→74 E:71→71 H:64/100

## Ecosystem Overview (May 16)
- Overall quality: 74/100 (→ plateau, Day 15), effectiveness: 71/100 (→ plateau, Day 15)
- 229 workflows (stable), health: 64/100 (→ stable, day 2 post-recovery)
- Engines: copilot (140+), claude (60+), codex (12), others
- **May 14 mass failure RECOVERY**: 36 open → 19 open [aw] failures in 48h post PR #32070 ✅
- **AWF Firewall v0.25.47 CONTAINED**: Smoke gate caught it before main ✅
- PR-review cluster: still ~272 wasted runs/day (0% success) — #31724
- CGO/CJS: still failing every push to main (P1, #29669)

## Top Performers (May 16)
1. **Agentic Maintenance** (Q:90 E:92) — 100% success ✅
2. **Issue Monster** (Q:85 E:87) — effective, ~6m39s runtime ✅
3. **Auto-Triage Issues** (Q:82 E:85) — 100% success ✅
4. **Bot Detection** (Q:82 E:83) — 100% success, 9s runtime ✅
5. **License Compliance Check** (Q:80 E:82) — ~98% success ✅
6. **PR Sous Chef** (Q:80 E:82) — 100% success (4/4) ✅
7. **Copilot SWE Agent** (Q:78 E:85) — 28/50 PRs merged (56%) ✅

## Pattern Classification (May 16)
- P0 (1): PR-review cluster (Q, Label Closed PRs, Doc Build-Deploy) — 0% success, ~272 wasted runs/day
- P1 (2): CGO/CJS regression (#29669), Daily Fact parse failures (#31432, #31524)
- P2 (3): Deployment Monitor zombie, Failure Investigator self-failure gap, Smoke Codex/Pi failures
- OK: Agentic Maintenance, Issue Monster, Auto-Close, Bot Detection, License, PR Triage, Auto-Triage, PR Sous Chef

## Active Issues (May 16)
- **PR-review cluster**: #31724, 0% success — highest-ROI fix
- **CGO/CJS**: #29669 open, failing every push
- **Daily Fact**: #31432, #31524 open
- **MCP gateway timeout**: #23153 open
- **Performance Regression**: #30180 open
- **Smoke Codex**: #32561 open
- **Smoke Pi**: #32553 open
- **AWF Firewall v0.25.47**: #32522 open (CONTAINED — PR#32503 closed)
- 19 open [aw] failure issues (recovering from 36 peak May 14)

## 15-day Quality Trend
- Quality:      74 (plateau, Day 15 — unchanged)
- Effectiveness: 71 (plateau, Day 15 — unchanged)
- Primary driver of plateau: PR-review cluster waste (~272 runs/day at 0%)

## Actions This Run
- Discussion created: Agent Performance Report — Week of 2026-05-16
- No new issues filed (existing #31724, #29669, #31432, #31524 cover active items)
- Updated shared memory + shared-alerts

Last updated: 2026-05-16T13:00Z by agent-performance-manager
