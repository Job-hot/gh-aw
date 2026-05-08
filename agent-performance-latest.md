# Agent Performance — 2026-05-08
Run: §25556968317 | Q:74→74 E:71→71

## Ecosystem Overview (May 8)
- Overall quality: 74/100 (→ stable plateau, day 7), effectiveness: 71/100 (→ stable)
- 217 workflows (→ unchanged), health: 61/100 (→ stable)
- Engines: copilot (140), claude (60), codex (12), pi (2), opencode/gemini/crush (1 each)
- PR-review cluster: Scout/Archie/Q//cloclo — ~100+ action_required/day (ongoing)
- **P0 ongoing**: Smoke Gemini 100% failure (35+ days), fetch failed TypeError

## Top Performers (May 8)
1. **Issue Monster** (Q:85 E:87) — Active, selected #30986, #30982, #30954 for Copilot ✅
2. **Agentic Maintenance** (Q:90 E:92) — Stable ✅
3. **Bot Detection** (Q:80 E:80) — Success ✅
4. **PR Triage Agent** (Q:80 E:80) — Success ✅
5. **Claude Code User Documentation Review** (Q:82 E:80) — Success ✅

## Key Patterns Detected (May 8)
- `over-creation`+`repetition`: PR-review cluster (Scout/Archie/Q//cloclo) — 0% success
- `engine-failure` P0: Smoke Gemini — fetch failed, 35+ days
- `missing-output`: Smoke Pi — no safe outputs called (noop violation)
- `missing-tool`: Smoke Codex — web-fetch MCP tool absent
- `under-creation` recurring: Resource Summarizer, Doc Build Deploy
- `scope-creep`: AI Moderator, Content Moderation on PR diff events

## New Failures Today (May 8)
- **AI Moderator** (#31016): failed
- **Auto-Triage Issues** (#30996): failed
- **Copilot PR Conversation NLP Analysis** (#31006): failed
- **Daily News** (#30967): failed
- **Dev** (#30970): failed
- **Smoke Codex** (#30979): cache miss + missing web-fetch
- **Smoke Gemini** (#30976): 100% failure (35+ days)
- **Smoke Pi** (#30975): no safe outputs

## Active Issues (May 8)
- **P0 ongoing**: Smoke Gemini (35+ days) — #30175 closed ineffective, #30976 new
- **P0 ongoing**: Smoke CI CGO/EROFS — #29666
- **P0 ongoing**: APM unpack systemic — #30252
- **P0 ongoing**: Daily Model Inventory Checker — #30043
- **P0 ongoing**: config.models unsupported field — #30307
- **P1 ongoing**: Smoke macOS ARM64 — filed 2026-05-07
- **P1 ongoing**: CI TestStrictModePermissions
- **P1 ongoing**: MCP gateway session timeout — #23153
- **P1 ongoing**: Performance Regression — #30180

## 7-day Quality Trend
- Quality:      73→74→74→74→74→74→74 (→ stable plateau, day 7)
- Effectiveness: 69→70→71→71→71→71→71 (→ stable plateau, day 7)

## Actions This Run
- Discussion created: Agent Performance Report — Week of 2026-05-08
- No new issues filed (existing P0/P1 tracked, Smoke Pi noop-violation noted)
- Pattern analysis complete: 8 workflows with failures today

Last updated: 2026-05-08T13:03Z by agent-performance-manager
