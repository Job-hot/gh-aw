# Agent Performance — 2026-05-19
Run: §26101076456 | Q:74→74 E:71→71 H:63/100 (→0)

## Ecosystem Overview (May 19)
- Overall quality: 74/100 (plateau day 18 — breakout expected to 76-78 once Agentic Maintenance restored)
- Effectiveness: 71/100 (plateau day 18 — expect 73-75 once critical blockers resolved)
- 231 workflows (stable), health: 63/100 (stable but degraded from Agentic Maintenance)
- Engines: copilot (140), claude (60), codex (12 blocked), others (~19)
- **P1 CRITICAL**: Agentic Maintenance compile failure — orchestrator DOWN Day 2
- **P1 CRITICAL**: CGO/CJS failing every push (90+ days, #29669)
- **P1 CRITICAL**: Codex OPENAI_API_KEY excluded (#32446, 12 workflows blocked)

## Top Performers (May 19)
1. **Issue Monster** (Q:85 E:87) — effective, ~6m39s runtime ✅
2. **Auto-Triage Issues** (Q:82 E:85) — 100% success ✅
3. **Bot Detection** (Q:82 E:83) — 100% success, 9s runtime ✅
4. **License Compliance Check** (Q:80 E:82) — ~98% success ✅
5. **PR Sous Chef** (Q:80 E:82) — 100% success ✅
6. **Copilot SWE Agent** (Q:78 E:85) — 56% PR merge rate ✅

## Pattern Classification (May 19)
- **Healthy (no patterns)**: 5 agents (Issue Monster, Auto-Triage, Bot Detection, License, PR Sous Chef)
- **P1 (5)**: Agentic Maintenance (NEW Day 2), CGO/CJS (#29669), Codex OPENAI_API_KEY (#32446), MCP gateway timeout (#23153), Performance Regression (#30180)
- **P2 (5)**: UK AI Operational Resilience (NEW), ET budget exhaustion (#32717), engine-fail-after-completion (#32736), Step Name recurring (#32955), [aw-compat] warnings (#32528)
- **Pattern-detector**: under-creation (5), resource-waste (4), over-creation (1), inconsistency (1), stale-outputs (1)

## Active Issues (May 19)
- **Agentic Maintenance compile**: NEW P1 — orchestrator down Day 2 (needs issue)
- **CGO/CJS**: #29669 open, 90+ days — P1 CRITICAL (needs escalation)
- **Codex OPENAI_API_KEY**: #32446 open — P1 (12 workflows blocked)
- **MCP gateway timeout**: #23153 — P1
- **Performance Regression**: #30180 — P1
- **Daily Observability ET exhaustion**: #32717 — P2
- **Engine-fail-after-completion**: #32736 — P2
- **Step Name Alignment recurring**: #32955 — P2
- ~22 open [aw] failure issues (stable)

## 18-day Quality Trend
- Quality:       74 (plateau day 18 — expect 76-78 once Agentic Maintenance restored)
- Effectiveness: 71 (plateau day 18 — expect 73-75)
- Health:        63 (→0 stable but degraded)
- Primary blocker: Agentic Maintenance compile failure + CGO/CJS unresolved 90+ days

## Actions This Run
- Discussion created: Agent Performance Report — Week of May 19, 2026
- No new issues filed (Agentic Maintenance compile needs issue, all others tracked)
- Updated shared memory + shared-alerts
- Comprehensive analysis of 231 workflows with pattern detection

Last updated: 2026-05-19T13:44Z by agent-performance-analyzer
