# Agent Performance Analyzer — Latest Run

**Timestamp:** 2026-06-08T14:02:14Z  
**Workflow:** Agent Performance Analyzer — Meta-Orchestrator  
**Run ID:** 27142793593  
**Run URL:** https://github.com/github/gh-aw/actions/runs/27142793593

## Executive Summary

- **Agents analyzed:** 22 active workflow types (100 runs in window)
- **Quality score:** 66/100 (↓2 from 68)
- **Effectiveness:** 60/100 (↓2 from 62)
- **Ecosystem health:** 68/100 (6-day declining trend: 82→81→78→74→71→68)

## Active P0/P1 Issues (Do Not Re-File)
- **Failure Cascade** (#37721 OPEN): awf-cli-proxy exit + LLM max_runs_exceeded — DO NOT RE-FILE
- **CJS re-regression** (#aw_cjs8 filed Jun 8): Re-regression — DO NOT RE-FILE
- **CGO unit tests** (#35028 OPEN): 100% failing — DO NOT RE-FILE
- **Daily Compiler Quality Check** (#37730 OPEN, 3rd day): Tool denials escalated P1 — DO NOT RE-FILE
- **Safe Output Health Monitor** (#37759 OPEN): Token exhaustion — DO NOT RE-FILE
- **Code Simplifier** (#37733 OPEN): Re-failed Jun 8 — DO NOT RE-FILE

## Improvement Issue Filed This Run
- **Issue Lifecycle Gap** (#aw_isg_jun8): Premature P1 closure pattern (CJS 2nd occurrence) — systemic process fix needed

## Agent Rankings

### Top Performers
1. copilot-swe-agent (Q:91, E:93) — 7 WIP PRs + 7 merges Jun 8; exceptional throughput
2. Agentic Maintenance (Q:85, E:88) — 100% success
3. License Compliance Check (Q:83, E:86) — 100% success
4. Auto-Close Parent Issues (Q:80, E:85) — 100% success
5. Auto-Triage Issues (Q:79, E:82) — 100% success
6. Issue Monster (Q:75, E:78) — 100% success
7. Agentic Commands (Q:70, E:65) — 36% (cascade-impacted)

### Needing Improvement
- AI Moderator (Q:35, E:20) — 0% (cascade)
- CJS (Q:30, E:25) — 0% (infrastructure, P1)
- CGO (Q:30, E:25) — 0% (infrastructure, P1)
- Claude Code User Documentation Review (Q:40, E:30) — 0% (fresh failure)

## Pattern Detection
- **Productive:** copilot-swe-agent fast-lane (7 merges/day Jun 8)
- **Cascade cluster (P0):** awf-cli-proxy → Smoke CI + AI Moderator + Q + Content Moderation + Agentic Commands
- **Tool denial cluster (Day 3):** Compiler Quality + Safe Output Health Monitor
- **Issue lifecycle gap:** CJS premature closure (2nd occurrence) → systemic process issue filed
- **Action_required flood:** ~25 action_required conclusions in window from cascade

## Coverage Notes
- copilot-swe-agent active on: bug fixes, doc fixes, linter improvements, A/B experiments, dependency reviews
- Weak coverage: Daily News, Dev, Layout Spec Maintainer, Weekly Workflow Analysis (all failing)
- Strong coverage: maintenance, compilation, spec enforcement, dependency management

## Coordination Notes for Other Orchestrators
- **WH:** Cascade #37721 is the root cause. awf-cli-proxy container exit must be patched before cascade clears.
- **Campaign:** copilot-swe-agent throughput healthy — complex campaigns viable; 7 merges in <3h Jun 8.
- **All:** Issue lifecycle gap issue filed (#aw_isg_jun8). See recommendations for P1 closure process.
- **All:** 6-day health decline continues (82→68). Cascade fix is the critical path.
