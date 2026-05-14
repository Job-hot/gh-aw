# Agent Performance — 2026-05-14
Run: §25861985037 | Q:74→74 E:71→71 H:62/100

## Ecosystem Overview (May 14)
- Overall quality: 74/100 (→ plateau, Day 13), effectiveness: 71/100 (→ plateau, Day 13)
- 225 workflows (+2 new since May 13), health: 62/100 (↓ -1)
- Engines: copilot (140+), claude (60+), codex (12), others
- **P0 EVENT**: 33 [aw] failure issues created TODAY — highest single-day count observed
- PR-review cluster: still ~272 wasted runs/day (0% success) — issue #31724
- CGO: still failing every push to main (P1, #29669)

## Top Performers (May 14)
1. **Agentic Maintenance** (Q:90 E:92) — 100% success ✅
2. **Issue Monster** (Q:85 E:87) — 6m39s run, effective ✅
3. **Auto-Close Parent Issues** (Q:82 E:85) — 100% success ✅
4. **Daily File Diet** (Q:80 E:80) — 16s, efficient ✅
5. **License Compliance Check** (Q:80 E:82) — 98% success ✅

## Pattern Classification (May 14)
- P0 (1): Mass failure event (33 new issues today)
- P1 (3): CGO/CJS regression, Q+PR-cluster structural failure, Daily agents batch
- P2 (4): Moderation twin flapping, Safe Output Health Monitor first fail, Daily Fact, Failure Investigator down
- OK (7): Agentic Maintenance, Issue Monster, Auto-Close, Daily File Diet, License, Bot Detection, PR Triage

## Mass Failure Event (2026-05-14)
- 33 failure issues created today (#32045–#32119)
- Affected: daily agents, code quality agents, smoke tests, moderators, Failure Investigator itself
- Possible cause: safe-output infra disruption or engine availability window
- PR #32070 merged today (safe output bundle fix) — may resolve some
- Total open [aw] failures: 36

## Active Issues (May 14)
- **Mass failure**: 36 open [aw] failure issues (up from 30 morning)
- **CGO/CJS**: #29669 open, failing every push
- **Q/PR-cluster**: #31724, 0% success, structural
- **Daily Fact**: #31432, #31524 open
- **MCP gateway timeout**: #23153 open
- **Performance Regression**: #30180 open

## 13-day Quality Trend
- Quality:       74→74→74→74→74→74→74→74→74→74→74→74→74→74 (plateau, Day 13)
- Effectiveness:  71→71→71→71→71→71→71→71→71→71→71→71→71→71 (plateau, Day 13)

## Actions This Run
- Discussion created: Agent Performance Report — Week of 2026-05-14
- Identified mass failure event (33 issues) as top priority
- Updated shared memory + shared-alerts
- No new issues filed (existing coverage sufficient; mass event needs human triage)

Last updated: 2026-05-14T13:26Z by agent-performance-manager

## Pattern Detector Results (May 14)
- P0 (3): Q, Label Closed PRs, Doc Build-Deploy (0% success, ~272 wasted runs/day)
- P1 (5): CGO, CJS, Daily Fact, Daily Security Red Team, Daily Cache Strategy
- P2 (5): AI Moderator, Content Moderation, Agentic Commands, Smoke CI, PR Sous Chef
- OK (9): Agentic Maintenance, Issue Monster, Auto-Close, Daily File Diet, License, Bot Detection, PR Triage, Workflow Normalizer, Safe Output Health Monitor (watch)

Dominant pattern: `inconsistency` in 12/22 agents (55%) — systemic trigger/environment instability
Key recommendation: gate Q/Label Closed PRs/Doc Build-Deploy behind workflow_dispatch immediately
Systemic: 13-day quality plateau warrants freeze-and-fix sprint
