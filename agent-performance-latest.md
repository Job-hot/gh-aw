# Agent Performance — 2026-04-30
Run: §25148001674 | Q:74→74 E:71→71

## Ecosystem Overview (Apr 30)
- Overall quality: 74/100 (→ stable), effectiveness: 71/100 (→ stable)
- 17 completed runs observed today: 15 success, 2 failure
- 3 in-progress (Agent Performance Analyzer, Schema Consistency Checker, jsweep)
- Session cost: $7.67 | 15.1M tokens | 140 action-minutes | 229 turns
- Engines: copilot:12, claude:8, codex:1

## Top Performers
1. **Test Quality Sentinel** (Q:90 E:92) — 3/3 success ✅
2. **Smoke CI** (Q:88 E:87) — 3/3 success ✅
3. **Daily Caveman Optimizer** (Q:85 E:85) — 2/2 success ✅
4. **CLI Version Checker** (Q:83 E:82) — 1/1 success ✅
5. **Documentation Unbloat** (Q:80 E:79) — 1/1 success ✅ (recovered from auth failure)

## Failures (Apr 30)
- **GitHub Remote MCP Authentication Test** (Q:10 E:0) — Day 9+ (#27965) — 0 tokens, infra failure
- **Design Decision Gate 🏗️** (Q:62 E:60) — 2/3 (67%) — 1 failure on PR copilot/add-github-ref-constraint-support

## 7-day Trends
- Quality: 72→73→74→74→74→74→74 (→ stable)
- Effectiveness: 68→69→70→71→71→71→71 (→ stable)
- Success rate: 93%→94%→95%→93%→57%→73%→85% (recovering)
- P1 open: 13→13→13→13→13→13→13 (→ stagnant)

## Issues/Actions This Run
- Discussion created (performance report, Apr 30)
- No new improvement issues (existing issues cover active failures)
- P1 backlog unchanged at 13 open items

Last updated: 2026-04-30T05:00Z by agent-performance-manager
