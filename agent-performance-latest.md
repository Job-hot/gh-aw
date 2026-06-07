# Agent Performance Analyzer — Latest Run

**Timestamp:** 2026-06-07T13:30:00Z  
**Workflow:** Agent Performance Analyzer  
**Run ID:** 27093462163  
**Run URL:** https://github.com/github/gh-aw/actions/runs/27093462163

## Executive Summary

- **Agents analyzed:** 25 active workflow types (100 runs)
- **Quality score:** 68/100 (↓2 from 70)
- **Effectiveness:** 62/100 (↓2 from 64)
- **Ecosystem health:** 71/100 (↓3 — 5-day declining trend: 82→81→78→74→71)

## Active P1 Issues (Do Not Re-File)
- **CJS re-regression** (#37503, OPEN Jun 7): Closed prematurely Jun 3 (#36410), re-regressed Jun 7
- **CGO unit tests** (#35028, OPEN): 17% success rate in recent window
- **Token guard escalation** (`#aw_tgescalate`, filed this run): Day 5 without #37145 implementation
- **Safe Output Health Monitor**: #37501 CLOSED Jun 7 — monitor for recurrence (Day 5 token exhaustion)

## Resolved Since Last Run ✅
- Code Simplifier: Fixed via #37489
- Daily Compiler Quality Check (partial): #37485 unblocked
- Documentation Healer model pinning: #37505
- safe-output-health AIC: #37506

## Agent Rankings

### Top Performers
1. copilot-swe-agent (Q:91, E:93) — 22 PRs merged June 7, exceptional throughput
2. Agentic Maintenance (Q:85, E:88) — 100% success
3. License Compliance Check (Q:83, E:86) — 100% success
4. Auto-Triage Issues (Q:79, E:82) — 4/4 success
5. CI (Q:78, E:80) — 3/3 success

### Needing Improvement
- AI Moderator (Q:35, E:20) — 0% success, no fix path (5+ days)
- Safe Output Health Monitor (Q:30, E:25) — token budget Day 5
- Workflow Forecast Report (Q:45, E:35) — 3 duplicate failures today
- Daily Documentation Healer (Q:48, E:35) — 4th consecutive failure (fix in progress #37505)

## Pattern Detection
- **Productive:** copilot-swe-agent fast-lane (22 merges/day)
- **Token exhaustion cluster (Day 5):** Safe Output Health Monitor + Daily Compiler → escalated `#aw_tgescalate`
- **CI blockage cluster:** CGO (17%) + CJS re-regression (#37503)
- **Auth/config failure cluster (Day 4):** Documentation Healer (fix in progress via #37505)
- **Forecast duplication:** 3 identical failure issues today (#37459, #37486, #37499)

## Issues Created This Run
- `#aw_tgescalate`: Token guard escalation (Day 5, blocking monitoring)

## Coordination Notes
- **WH:** CJS #37503 filed today (re-regression). CGO #35028 still active — fresh P1 attention needed.
- **WH:** Token guard #37145 escalated — Day 5. Assign owner immediately.
- **Campaign:** copilot-swe-agent throughput very high (22 merges/day) — scaling complex tasks is viable.
- **All:** Ecosystem health 5-day declining trend (82→71) — systemic attention needed.
