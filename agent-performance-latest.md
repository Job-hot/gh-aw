# Agent Performance — 2026-05-15
Run: §25919679386 | Q:74→74 E:71→71 H:64/100

## Ecosystem Overview (May 15)
- Overall quality: 74/100 (→ plateau, Day 14), effectiveness: 71/100 (→ plateau, Day 14)
- 229 workflows (+4 new since May 14), health: 64/100 (↑ +2 — partial recovery)
- Engines: copilot (140+), claude (60+), codex (12), others
- **Mass failure RECOVERY**: PR #32070 merged May 14 restored ~8 workflows by May 15
- PR-review cluster: still ~272 wasted runs/day (0% success) — issue #31724
- CGO/CJS: still failing every push to main (P1, #29669)

## Top Performers (May 15)
1. **Agentic Maintenance** (Q:90 E:92) — 100% success ✅
2. **Issue Monster** (Q:85 E:87) — effective, 6m39s runtime ✅
3. **Auto-Close Parent Issues** (Q:82 E:85) — 100% success ✅
4. **Daily File Diet** (Q:80 E:80) — 16s, extremely efficient ✅
5. **License Compliance Check** (Q:80 E:82) — 98% success ✅

## Pattern Classification (May 15)
- P0 (1): PR-review cluster (Q, Label Closed PRs, Doc Build-Deploy) — 0% success, ~272 wasted runs/day
- P1 (2): CGO/CJS regression (#29669), Daily Fact parse failures (#31432, #31524)
- P2 (2): Failure Investigator self-failure during incidents, Deployment Monitor zombie
- OK (9): Agentic Maintenance, Issue Monster, Auto-Close, Daily File Diet, License, Bot Detection, PR Triage, Workflow Normalizer, Safe Output Health Monitor

## Active Issues (May 15)
- **PR-review cluster**: #31724, 0% success, structural — highest-ROI fix
- **CGO/CJS**: #29669 open, failing every push
- **Daily Fact**: #31432, #31524 open (post-PR#31411 still failing)
- **MCP gateway timeout**: #23153 open
- **Performance Regression**: #30180 open
- ~30+ open [aw] failure issues (recovering from May 14 peak of 36)

## 14-day Quality Trend
- Quality:      74 (plateau, Day 14 — unchanged)
- Effectiveness: 71 (plateau, Day 14 — unchanged)

## Actions This Run
- Discussion created: Agent Performance Report — Week of 2026-05-15
- Recovery event noted: PR #32070 effective, health score +2 to 64/100
- No new issues filed (existing #31724, #29669, #31432, #31524 cover active items)
- Updated shared memory + shared-alerts

Last updated: 2026-05-15T13:14Z by agent-performance-manager
