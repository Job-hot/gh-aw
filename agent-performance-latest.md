# Agent Performance Analyzer — Latest Run

**Timestamp:** 2026-06-06T13:02:11Z  
**Workflow:** Agent Performance Analyzer  
**Run ID:** 27062936303  
**Run URL:** https://github.com/github/gh-aw/actions/runs/27062936303

## Executive Summary

- **Agents analyzed:** 26 active workflow types (100 runs)
- **Quality score:** 70/100 (↓2 from 72)
- **Effectiveness:** 64/100 (↓2 from 66)
- **Ecosystem health:** 74/100 (↓7 from 81 — 4-day declining trend)

## Key Events This Period (June 2–6)

### Resolved ✅
- **Safe Outputs MCP P0** (June 6): `create_pull_request` path bug fixed PR #37299. RESOLVED.
- **CJS typecheck** (#36410): Was 0% P1, now **85% success (11/13 runs)** — LIKELY RESOLVED. WH: confirm closure.

### Improving ⬆️
- **CGO unit tests** (#35028): Was 0%, now 67% success (8/12 runs). Still 2 failures.

### New Issues (Do Not Re-File)
- **max-ai-credits migration** (#37282, #37284): 6 workflows using deprecated `max-effective-tokens`. Migration tracked.
- All other critical issues unchanged from shared-alerts.md.

## Agent Rankings

### Top Performers
1. copilot-swe-agent (Q:90, E:92) — 26/30 PR merges, 8 WIP PRs, major platform fixes
2. Agentic Maintenance (Q:85, E:90) — 100% success, proactive migrations
3. License Compliance Check (Q:83, E:88) — 100% success (3/3)
4. Auto-Close Parent Issues (Q:80, E:85) — 100% success (2/2)
5. Smoke CI (Q:78, E:82) — 93% success (14/15)

### Needing Improvement
- AI Moderator (Q:35, E:20) — 0% success, no fix path (persistent)
- Safe Output Health Monitor (Q:30, E:25) — token budget, #37264 OPEN
- Daily BYOK Ollama Test (E:15) — auth failures, #37211 OPEN
- Doc Build - Deploy (E:62) — 62% success (3/8 failures), watch list

## Pattern Detection

- **Productive:** copilot-swe-agent fast iteration, Agentic Maintenance proactive hygiene
- **Token exhaustion cluster (static):** Safe Output Health Monitor + jsweep + Daily Compiler + Firewall Logs — 4 days no improvement
- **Auth config cluster:** Doc Healer + Model Inventory + BYOK Ollama — #37039 fix insufficient; #37271 filed
- **Failure reporter duplication** (#35984): 60% duplicate rate — unresolved

## No New Issues Created
All critical issues already tracked in shared-alerts.md.

## Coordination Notes
- For WH: CJS at 85% — recommend closing/updating #36410
- For Campaign Manager: copilot-swe-agent throughput high (26 merges) — consider more complex tasks
- Token guard #37145 still not implemented — token exhaustion cluster growing risk
