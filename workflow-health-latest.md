# Workflow Health — 2026-06-01T06:08Z

Score: 82/100 (↑1 from 81)
Workflows: 237 | Lock files: 237/237 (100% ✅) | Run: §26738226272

## KEY FINDINGS

### Status (June 1)
- **Compilation:** 237/237 workflows have lock files (100% ✅) — unchanged
- **New P0/P1 issues filed this run: 0** (no new systemic issues; tracked below)
- **IMPROVEMENT:** Most smoke test issues CLOSED (#35959, #36018, #36019, #35955, #35954)
- **Token Budget Exhaustion pattern** newly observed affecting 2+ workflows

### Critical Issues (P1) 🚨
- **Step Name Alignment** (#36062 OPEN): Still failing, new run failure #36187 (Claude engine terminating). 100% fail.
- **Failure-reporters duplication**: #35984 still OPEN
- **LintMonster backlog**: #36050 still open; #36175, #36173 new batch today

### P2 Issues ⚠️
- **jsweep JavaScript Unbloater** (#36183): Token budget exhausted  
- **Daily Compiler Quality Check** (#36172): Token budget exhausted
- **Daily Firewall Logs Collector** (#36171): New instance (was #36047 closed, recurring)
- **Daily SPDD Spec Planner** (#36138): Failed May 31
- **Token budget exhaustion pattern**: jsweep + daily-compiler both hit limit — may be systemic

### Actions Taken This Run
- No new issues filed (failures already tracked or P2 via auto-filing)
- Updated shared memory with June 1 assessment

### Trends
- Score: 82/100 (↑1 from 81)
- Smoke tests: mostly CLOSED — major improvement
- Token budget exhaustion: new recurring pattern affecting analytics/quality workflows
- CGO: 33% failure rate, 1 recent failure in last 5 runs

Last updated: 2026-06-01T06:08:00Z by workflow-health-manager
