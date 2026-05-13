# Agent Performance — 2026-05-13
Run: §25801923272 | Q:74→74 E:71→71

## Ecosystem Overview (May 13)
- Overall quality: 74/100 (→ stable plateau, Day 12), effectiveness: 71/100 (→ stable)
- 223 workflows (+4 new since May 12), health: 63/100 (↓ -1)
- Engines: copilot (140), claude (60), codex (12), pi (2), opencode/gemini/crush (1 each)
- PR-review cluster: Q/Scout/Archie/cloclo/Grumpy/Security Review/PR Nitpick/PR Code Quality — ~272 wasted run-attempts/day (0% success)
- **P0**: Q, Scout (0% success, structural trigger failure)
- **Pattern**: Under-creation dominant (9/18 profiled = 50%), CI cluster compound friction

## Top Performers (May 13)
1. **Agentic Maintenance** (Q:90 E:92) — 100% success ✅
2. **Issue Monster** (Q:85 E:87) — Active and effective ✅
3. **Auto-Close Parent Issues** (Q:82 E:85) — 100% success ✅
4. **Bot Detection** (Q:80 E:80) — Stable ✅
5. **PR Triage Agent** (Q:80 E:80) — Stable ✅

## Pattern Classification (May 13 — 18 agents profiled)
- P0 (2): Q, Scout
- P1 (6): Agentic Commands, CGO, CJS, PR Sous Chef, Daily Fact, Daily Security Red Team
- P2 (5): Content Moderation, AI Moderator, Smoke CI, Semantic Function Refactoring, Daily Cache Strategy
- OK (5): Agentic Maintenance, Issue Monster, Bot Detection, PR Triage, Auto-Triage

## New This Run (2026-05-13)
- **5 new failures**: PR Sous Chef #31931, Draft PR Cleanup #31929, Semantic Function Refactoring #31827, Daily Security Red Team #31817, Daily Cache Strategy #31773
- **Daily agents batch failure**: Daily Fact/Security/Cache all failing → suspected shared root cause
- **CI cluster**: CGO (22%) + CJS (25%) + Smoke CI (50%) = compounding PR friction
- **Moderation twin 57%**: Content Moderation + AI Moderator identical split → shared upstream instability

## Active Issues (May 13)
- **P0/PR-cluster**: Q+Scout 0% success, #31724 (watch)
- **P1**: Daily Fact (#31432, #31524), Daily Security Red Team (#31817), CGO/CJS (#31860), PR Sous Chef (#31931), MCP gateway timeout (#23153), Performance Regression (#30180)
- **P2**: PR-review cluster waste (#31724), Security findings (#31708, #31704), CI integration (#31860), Node.js 20 deprecation (Sep 2026)

## 12-day Quality Trend
- Quality:      74→74→74→74→74→74→74→74→74→74→74→74→74 (→ plateau, Day 12)
- Effectiveness: 71→71→71→71→71→71→71→71→71→71→71→71→71 (→ plateau, Day 12)

## Actions This Run
- Discussion created: Agent Performance Report — Week of 2026-05-13
- Pattern analysis: pattern-detector classified 18 agents (P0:2, P1:6, P2:5, OK:5)
- Identified daily agents batch failure as shared root cause hypothesis
- No new P0/P1 issues filed (existing issues cover active items)

Last updated: 2026-05-13T13:26Z by agent-performance-manager
