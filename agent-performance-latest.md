# Agent Performance Analyzer — Latest Run

**Timestamp:** 2026-06-02T14:03:32Z  
**Workflow:** Agent Performance Analyzer  
**Run ID:** 26824629412  
**Run URL:** https://github.com/github/gh-aw/actions/runs/26824629412

## Executive Summary

- **Agents analyzed:** 23 active workflow groups (~237 total workflows)
- **Quality score:** 72/100 (↓1 from 73 — new failures offset improvements)
- **Effectiveness:** 66/100 (↓1 — Agentic Commands degraded further)
- **Ecosystem health:** 81/100 (unchanged — CJS P1 still open)

## Key Changes Since Yesterday (June 1)

### Improvements ✅
- **Step Name Alignment** (#36062): PASSED today — likely resolved, monitor
- **Agentic Maintenance**: 75% ok (4 runs) — healthy
- **copilot-swe-agent**: 22 merged PRs recently, 5 open WIP PRs — high throughput

### New Issues / Regressions 🚨
- **CJS typecheck P1** (#36410, filed June 2): Still 100% failing (4 failed + more cancelled)
- **CGO** (2 runs, 0% success, 2 failures): continuing 100% failure
- **Typist - Go Type Analysis failed** (#36443): new failure report
- **Failure Investigator (6h) failed** (#36424): recurring failure reporter
- **Q workflow**: 0% ok in 18 runs — fully blocked (pattern unchanged)
- **Agentic Commands**: 22% ok in 18 runs — degraded from prior period

## Critical Findings (No Change — Do Not Re-File)

### P1 - High Priority
1. **CJS typecheck broken** (#36410, filed June 2): 100% fail since June 1 ~23:32Z
2. **CGO unit tests** (#35028): Escalated to 100% failure (20+ runs)
3. **Q workflow**: 0% success rate (18 runs, all cancelled) — fully blocked
4. **LintMonster backlog** (#36050): intermittent failures
5. **Failure-reporters duplication** (#35984): 60% duplicate rate

### P2 - Watch
- **Token budget exhaustion**: jsweep (#36183) + daily-compiler (#36172)
- **chaos-test PR stall**: 10+ open PRs, 0 merges
- **Agentic Commands degradation**: 22% ok (18 runs) — was higher previously
- **Typist workflow failing**: #36443 (new)

## Top Performers (unchanged)

1. spec-enforcer (85/100 quality, 88/100 effectiveness)
2. copilot-swe-agent (84/100 quality, 82/100 effectiveness, 22 merged PRs vs 8 open)
3. License Compliance Check (100% success rate)
4. Auto-Close Parent Issues (100% success rate)

## Pattern Detection

- **blocked**: Q (0%, 18 runs), CGO (0%, 5 runs), CJS (0%, 6 runs), AI Moderator (0%)
- **degraded**: Agentic Commands (22%, 18 runs), Smoke CI (9%, 11 runs)
- **over-creation/stall**: chaos-test flooding (10+ PRs, 0 merges)
- **token-exhaustion**: jsweep + daily-compiler (systemic)
- **healthy**: copilot-swe-agent, License Compliance, Auto-Close, Agentic Maintenance

## Issues Created This Run

None — all critical issues already tracked per Do Not Re-File list.

## Coordination Notes

### For Campaign Manager
- chaos-test stall persists (10+ open PRs, 0 merges) — should pause or escalate
- copilot-swe-agent throughput high (22 merged), consider assigning more complex tasks
- Typist workflow (#36443) needs investigation — new failure

### For Workflow Health Manager
- CJS typecheck P1 (#36410) filed June 2 — track resolution
- CGO escalating — auto-notifier #35028 open
- Step Name Alignment possibly resolved — confirm and close #36062/#36187 if confirmed

Last updated: 2026-06-02T14:03:32Z by agent-performance-analyzer
