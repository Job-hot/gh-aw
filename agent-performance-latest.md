# Agent Performance — 2026-05-07
Run: §25498006520 | Q:74→74 E:71→71

## Ecosystem Overview (May 7)
- Overall quality: 74/100 (→ stable plateau, day 6), effectiveness: 71/100 (→ stable)
- 217 workflows (+3 from May 6), health: 61/100 (→ stable)
- Engines active: copilot, claude, codex; gemini 100% broken ongoing (33+ days)
- PR-review cluster: Scout/Archie/Q//cloclo — 4 agents, ~100+ action_required/day
- **NEW P1**: CJS workflow 100% failure since May 7 ~11:00 UTC (determine_automatic_lockdown.test.cjs:70)

## Top Performers (May 7)
1. **Agentic Maintenance** (Q:90 E:92) — 3/3 success ✅
2. **Issue Monster** (Q:85 E:85) — 3/3 success ✅
3. **License Compliance Check** (Q:89 E:88) — success ✅
4. **Auto-Triage Issues** (Q:83 E:82) — success ✅
5. **Bot Detection / PR Triage Agent / Daily File Diet** (Q:~80) — success ✅

## Key Patterns Detected (May 7)
- `over-creation`+`repetition`: PR-review cluster (Scout/Archie/Q//cloclo) — 0% success rate
- `scope-creep`: Content Moderation + AI Moderator firing on PR diff events outside domain
- `under-creation` NEW: CJS 100% failure (determine_automatic_lockdown.test.cjs mock not called)
- `under-creation` chronic: Resource Summarizer (2 skips), Doc Build - Deploy (3 action_required)
- Fleet health: 26% clean (5/19 agents no issues)

## Active Issues (May 7)
- **P0 ongoing**: Smoke Gemini (33+ days) — #30175, #29852
- **P0 ongoing**: Smoke CI CGO/EROFS — #29666
- **P0 ongoing**: APM unpack systemic — #30252
- **P0 ongoing**: Daily Model Inventory Checker — #30043
- **P1 NEW**: CJS 100% failure — determine_automatic_lockdown.test.cjs:70 (vi.fn() mock 0 calls)
- **P1 ongoing**: Smoke macOS ARM64 — ISSUE FILED May 7 ✅
- **P1 ongoing**: CI regression TestStrictModePermissions
- **P1 ongoing**: config.models unsupported field — #30307
- **P1 ongoing**: MCP gateway session timeout — #23153
- **P1 ongoing**: Performance Regression in Validation — #30180

## 7-day Quality Trend
- Quality:      73→74→74→74→74→74→74 (→ stable plateau, day 6)
- Effectiveness: 69→70→71→71→71→71→71 (→ stable plateau, day 6)

## Actions This Run
- Discussion created: Agent Performance Report — Week of 2026-05-07
- No new issues filed (CJS tracked in discussion; existing P0s tracked)
- Pattern-detected 19 agents: 5 clean (26%), 14 with patterns (74%)
- Top recommendation: Investigate CJS regression immediately + consolidate PR-review cluster

Last updated: 2026-05-07T13:15Z by agent-performance-manager
