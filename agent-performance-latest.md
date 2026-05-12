# Agent Performance — 2026-05-12
Run: §25736968122 | Q:74→74 E:71→71

## Ecosystem Overview (May 12)
- Overall quality: 74/100 (→ stable plateau, day 11), effectiveness: 71/100 (→ stable)
- 219 workflows (+1 new), health: 64/100 (↑ +2)
- Engines: copilot (140), claude (60), codex (12), pi (2), opencode/gemini/crush (1 each)
- PR-review cluster: Q/Scout/Archie/cloclo/Grumpy/Security Review/PR Nitpick/PR Code Quality — ~272 wasted run-attempts/day (0% success)
- **P0**: None ✅ (APM #30252 CLOSED)
- **Pattern**: Under-creation dominant (8/19 profiled = 42%), consistent with prior run

## Top Performers (May 12)
1. **Agentic Maintenance** (Q:90 E:92) — Stable top performer ✅
2. **Issue Monster** (Q:85 E:87) — Active and effective ✅
3. **Auto-Close Parent Issues** (Q:82 E:85) — 100% success rate ✅
4. **Bot Detection** (Q:80 E:80) — Stable ✅
5. **PR Triage Agent** (Q:80 E:80) — Stable ✅

## Key Patterns Detected (May 12)
- `under-creation` (8 agents, 42%): PR-review cluster (8), Daily Fact, Resource Summarizer, Plan Command, Deployment Incident Monitor, Design Decision Gate, Go Logger Enhancement, Step Name Alignment, jsweep
- `inconsistency` (6 agents): PR-review cluster, Content Moderation, AI Moderator, Daily Fact, Resource Summarizer, Plan Command
- `over-creation` (2 agents): PR-review cluster (run-attempts), Plan Command (output bursts)
- `scope-creep` (1 agent): AI Moderator (recovering)

## New This Run (2026-05-12)
- **4 same-day failures**: Design Decision Gate #31626, Go Logger Enhancement #31628, Step Name Alignment #31636, jsweep #31637 — possible shared root cause (PR #31418 side-effect or engine issue)
- **Daily Fact still failing** post-PR#31411 merge; issues #31432, #31524 open
- **Smoke failures**: Gemini (#31575), Pi (#31563), Codex (#31567) ongoing

## Active Issues (May 12)
- **P1**: Daily Fact parse failures (#31432, #31524), Smoke Gemini (#31575), MCP gateway timeout (#23153), Performance Regression (#30180)
- **P2**: PR-review cluster waste (~272/day), Firewall reporting (#31607, #31620), 4 new failures (watch), Deployment Incident Monitor (zombie), Node.js 20 deprecation (Sep 2026)

## 11-day Quality Trend
- Quality:      74→74→74→74→74→74→74→74→74→74→74→74 (→ stable plateau, day 11)
- Effectiveness: 71→71→71→71→71→71→71→71→71→71→71→71 (→ stable plateau, day 11)

## Actions This Run
- Discussion created: Agent Performance Report — Week of 2026-05-12
- Pattern analysis: pattern-detector classified 19 agents
- No new P0/P1 issues filed (existing issues cover active P1s)
- Flagged 4 same-day failures as potential shared-cause cluster

Last updated: 2026-05-12T13:19Z by agent-performance-manager
