# Agent Performance — 2026-05-06
Run: §25437379610 | Q:74→74 E:71→71

## Ecosystem Overview (May 6)
- Overall quality: 74/100 (→ stable plateau, day 5), effectiveness: 71/100 (→ stable)
- 214 workflows (+1 from May 5), health: 61/100 (↓-2 from CI regression on main)
- Engines active: copilot, claude, codex; gemini 100% broken ongoing (31+ days)
- PR-review cluster: 9 agents, ~65 action_required in sampled window, expected behavior
- New CI regression: TestStrictModePermissions failing on main since May 6 ~01:16 UTC

## Top Performers (May 6)
1. **Agentic Maintenance** (Q:90 E:92) — success ✅
2. **License Compliance Check** (Q:89 E:88) — success ✅
3. **Auto-Close Parent Issues** (Q:85 E:83) — success ✅
4. **Auto-Triage Issues** (Q:83 E:82) — success ✅ (recovered from #30205)
5. **Content Moderation** (Q:80 E:78) — partial success

## Key Patterns Detected (May 6)
- `over-triggering`: PR-review cluster (9 agents) all on same PR events — 100+ action_required/day
- `under-creation`: Resource Summarizer Agent — no safe-output completion signals
- `stagnant P0`: Smoke Gemini 31+ days, macOS ARM64 75+ days (no issue!)
- `recovery`: 6 P1 issues closed this cycle ✅ (recovery pipeline working)
- `new CI regression`: TestStrictModePermissions/strict mode integration test (Playwright MCP deprecation)

## Active Issues (May 6)
- **P0 ongoing**: Smoke Gemini 100% fail (31+ days) — #30175, #29852
- **P0 ongoing**: Smoke CI CGO/EROFS — #29666
- **P0 ongoing**: APM unpack systemic — #30252
- **P0 ongoing**: Daily Model Inventory Checker — #30043
- **P1 ongoing**: Smoke macOS ARM64 (NO ISSUE — 75+ days)
- **P1 new**: CI regression TestStrictModePermissions — filed by Workflow Health
- **P1 ongoing**: config.models unsupported field — #30307
- **P1 ongoing**: MCP gateway session timeout — #23153
- **P1 ongoing**: Performance Regression in Validation — #30180

## 7-day Quality Trend
- Quality:      73→74→74→74→74→74→74 (→ stable plateau, day 5)
- Effectiveness: 69→70→71→71→71→71→71 (→ stable plateau, day 5)

## Actions This Run
- Discussion created: Agent Performance Report — Week of 2026-05-06
- No new issues filed (existing issues tracked)
- Top recommendation: Consolidate PR-review cluster (9 agents, 100+ action_required/day)

Last updated: 2026-05-06T13:14Z by agent-performance-manager

## Pattern Detector Results (May 6)
- Fleet health: 33% clean (7/21 agents no issues)
- Critical: Smoke CI, CGO, Documentation Unbloat, Resource Summarizer, Label Closed PRs
- High: Scout, /cloclo, Q, Archie, AI Moderator, Content Moderation (over_creation + repetition)
- AI Moderator + Content Moderation: scope_creep (firing on PR diff events outside their domain)
- Medium: Grumpy, Security Review, PR Nitpick (repetition + inconsistency at cluster level)
- Healthy: Meta-orchestrators + maintenance workers (low severity)
