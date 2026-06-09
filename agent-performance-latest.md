# Agent Performance Analyzer — Latest Run

**Timestamp:** 2026-06-09T13:45:00Z  
**Workflow:** Agent Performance Analyzer — Meta-Orchestrator  
**Run ID:** 27209785615  
**Run URL:** https://github.com/github/gh-aw/actions/runs/27209785615

## Executive Summary

- **Agents analyzed:** 22 active workflow types (80 runs in window)
- **Quality score:** 67/100 (→ -1 from Jun 8's 66 — essentially stable)
- **Effectiveness score:** 63/100 (↑3 from 60 — recovery continuing)
- **Ecosystem health:** 83/100 (maintained; cascade recovery holding)

## Active P0/P1 Issues (Do Not Re-File)
- **Daily Compiler Quality Check** (#38021 OPEN, 4th day): tool denials (5/5) — DO NOT RE-FILE
- **Tool Denial Cluster** (#aw_tdcluster9 filed Jun 9): systemic shell() pattern — DO NOT RE-FILE
- **Safe Output Health Monitor** (#38039 OPEN): AI credits exceeded — DO NOT RE-FILE
- **Code Simplifier** (#38026 OPEN): Missing tools — DO NOT RE-FILE
- **AI Credits Cluster (3 known)**: #38025, #38024, #38039 — individual issues OPEN — DO NOT RE-FILE individual ones

## New Issue Filed This Run
- **AI Credits Cluster Expansion** (#aw_aic_exp9): 8 workflows hitting max-ai-credits limit (up from 3 this morning); includes Workflow Health Manager, Impact Efficiency Report, Daily AgentRx, Smoke Gemini — systemic budget review needed

## Agent Rankings

### Top Performers
1. copilot-swe-agent (Q:88, E:85) — 11/20 PRs merged (55%) Jun 8-9; strong throughput
2. Auto-Close Parent Issues (Q:82, E:85) — 100% success
3. Smoke CI (Q:80, E:78) — 75% success (3/4)
4. Bot Detection (Q:78, E:78) — 100% success
5. Avenger (Q:75, E:75) — 100% success
6. Daily File Diet (Q:75, E:75) — 100% success
7. Agentic Maintenance (Q:74, E:72) — mixed (1/2)
8. Running Copilot Code Review (Q:74, E:74) — 100% success

### Needing Improvement
- Daily Compiler Quality Check (Q:20, E:10) — 0% (tool denial, Day 4, #38021)
- AI Credits Cluster x8: Daily AgentRx, CLI Tools Tester, WH Manager, Impact Efficiency, Safe Output Health, Test Quality, Matt Pocock, Smoke Gemini (Q:35-45, E:20-30)
- CJS (Q:40, E:30) — 0/3 failure (infrastructure issues)
- Auth/Token failures: 3 workflows failing on authentication

## Pattern Detection
- **Productive:** copilot-swe-agent continues strong output (55% merge, 12 substantive PRs)
- **AI credits expansion (P1):** 8 workflows hitting max-ai-credits — up from 3 this morning; Workflow Health Manager now included (blind spot risk)
- **Tool denial cluster (Day 4):** Compiler Quality + Deep Research + jsweep using shell() pattern (blocked)
- **action_required pattern (normal):** Q (13 runs) + AI Moderator (6 runs) — this is EXPECTED behavior, not failure
- **Recovery holding:** Health score 83/100; cascade resolution stable

## Coverage Notes
- Strong: maintenance, compilation, PR review, contribution checking, spec enforcement
- Weak: Daily News (failing), Daily Workflow Updater (failing), Workflow Health Manager (AI credits)
- Risk: Workflow Health Manager failing = blind spot in health monitoring coverage

## Coordination Notes for Other Orchestrators
- **WH:** AI credits cluster has expanded to 8 workflows including WH Manager itself — health monitoring blind spot
- **Campaign:** copilot-swe-agent throughput healthy, 55% merge rate; complex campaigns viable
- **All:** AI credits cluster expansion is the new P1 — systemic budget config review needed
- **All:** Tool denial cluster (Day 4) still unresolved — prompt engineering fix needed for shell() usage
