# Agent Performance — 2026-05-09
Run: §25601701383 | Q:74→74 E:71→71

## Ecosystem Overview (May 9)
- Overall quality: 74/100 (→ stable plateau, day 8), effectiveness: 71/100 (→ stable)
- 218 workflows (+1), health: 61/100 (→ stable, day 3)
- Engines: copilot (140), claude (60), codex (12), pi (2), opencode/gemini/crush (1 each)
- PR-review cluster: Scout/Archie/Q//cloclo — ~100+ action_required/day (ongoing)
- **P0 ongoing**: Smoke Gemini 100% failure (35+ days), fetch TypeError

## Top Performers (May 9)
1. **Agentic Maintenance** (Q:90 E:92) — Confirmed in-progress ✅
2. **Issue Monster** (Q:85 E:87) — Active, selected #30986, #30982, #30954 for Copilot ✅
3. **Auto-Close Parent Issues** (Q:82 E:85) — 1/1 success today ✅
4. **AI Moderator** (Q:78 E:76) — Recovering: 2/3 success today ↑
5. **Bot Detection** (Q:80 E:80) — Stable ✅
6. **PR Triage Agent** (Q:80 E:80) — Stable ✅
7. **Content Moderation** (Q:75 E:73) — 2/3 success today

## Key Patterns Detected (May 9)
- `over-creation`+`repetition`: PR-review cluster (Scout/Archie/Q//cloclo) — 0% success, 34 runs today
- `over-creation`: Plan Command — 5 [plan] issues in seconds (#31207-#31211)
- `engine-failure` P0: Smoke Gemini — fetch failed, 35+ days
- `missing-output`: Smoke Pi — noop-violation (no safe outputs called)
- `missing-tool`: Smoke Codex — web-fetch MCP tool absent
- `under-creation`: Resource Summarizer, Doc Build Deploy
- `zombie`: Deployment Incident Monitor — 4 runs, 0 conclusions
- `scope-creep` improving: AI Moderator (2/3 success vs prior 0%)

## New Failures Today (May 9)
- **Dev** (#31185): failed
- **Stale PR Cleanup** (#31183): failed
- **Weekly Editors Health Check** (#31181): failed
- **Plan Command**: over-creation (#31207-#31211, 5 issues in one batch)

## Active Issues (May 9)
- **P0 ongoing**: Smoke Gemini (35+ days) — #30175 fix ineffective
- **P0 ongoing**: Smoke CI CGO/EROFS — #29666
- **P0 ongoing**: APM unpack systemic — #30252
- **P0 ongoing**: Daily Model Inventory Checker — #30043
- **P0 ongoing**: config.models unsupported field — #30307
- **P1 ongoing**: Smoke macOS ARM64 — filed 2026-05-07
- **P1 ongoing**: CI TestStrictModePermissions
- **P1 ongoing**: MCP gateway session timeout — #23153
- **P1 ongoing**: Performance Regression — #30180

## 7-day Quality Trend
- Quality:      74→74→74→74→74→74→74→74 (→ stable plateau, day 8)
- Effectiveness: 71→71→71→71→71→71→71→71 (→ stable plateau, day 8)

## Actions This Run
- Discussion created: Agent Performance Report — Week of 2026-05-09
- No new P0/P1 issues filed (existing tracked items unchanged)
- Pattern analysis complete: Plan Command over-creation newly detected

Last updated: 2026-05-09T13:00Z by agent-performance-manager
