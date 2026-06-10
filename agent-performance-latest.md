# Agent Performance Analyzer — Latest Run

**Timestamp:** 2026-06-10T13:53:00Z  
**Workflow:** Agent Performance Analyzer — Meta-Orchestrator  
**Run ID:** 27280947818  
**Run URL:** https://github.com/github/gh-aw/actions/runs/27280947818

## Executive Summary

- **Agents analyzed:** 22 active workflow types (~80 runs in window)
- **Quality score:** 68/100 (↑1 from Jun 9's 67 — stable/improving)
- **Effectiveness score:** 64/100 (↑1 from 63 — slow recovery continuing)
- **Ecosystem health:** 87/100 (↑4 from 83 — P0/P1 cascade fully resolved Jun 9→10)

## Active P0/P1 Issues (Do Not Re-File)
- **AI Credits Cluster (8 workflows)** (#aw_aic_exp9 filed Jun 9 — OPEN, Day 2): Daily AgentRx, CLI Tools Tester, WH Manager, Impact Efficiency, Safe Output Health, Test Quality, Matt Pocock, Smoke Gemini — DO NOT RE-FILE; systemic budget config fix needed
  - NEW Jun 10 failure issues: #38288, #38259, #38329, #38260, #38300, #38302, #38278, #38296 — all covered by #aw_aic_exp9 cluster
- **Daily Safe Outputs Git Simulator** (#aw_gitsim10 filed Jun 10 — OPEN): memory/git-simulator branch missing signed-commit seed — DO NOT RE-FILE
- **Auto-Triage composite** (#38309 filed Jun 10): Sub-Issue Closer + Auto-Triage Issues transient failure — DO NOT RE-FILE

## New Issues Filed This Run
- **None** — all active failures already tracked; do not create duplicates

## Agent Rankings

### Top Performers
1. copilot-swe-agent (Q:90, E:88) — 8 active PRs Jun 10; addressing real user bugs (safe_outputs timeout, SHA-pin, context propagation, OTLP wiring)
2. Agentic Maintenance (Q:82, E:85) — 2/2 successful runs; CLI version updates (#38298)
3. Bot Detection / Avenger (Q:80, E:82) — 100% success; focused scope
4. Daily File Diet (Q:80, E:80) — 100% success; consistent repo hygiene
5. Content Moderation (Q:78, E:75) — 66% success (4/6 runs)
6. AI Moderator (Q:75, E:72) — action_required is EXPECTED behavior (human review requests)
7. Daily AIC Consumption Report (Q:78, E:78) — 100% success; good observability output
8. Issue Monster (Q:72, E:65) — 1 success, 1 failure today

### Needing Improvement
- AI Credits Cluster x8: Q:35-45, E:20-30 — budget exhaustion, Day 2 (#aw_aic_exp9)
- Auto-Triage Issues (Q:40, E:25) — failed today, #38308
- Sub-Issue Closer (Q:40, E:25) — failed today, #38305
- Code Simplifier (Q:35, E:20) — failed today, #38278
- Daily News (Q:45, E:35) — failed today, #38318

## Pattern Detection
- **Productive:** copilot-swe-agent highest engagement rate; 8 PRs addressing user-filed bugs same-day
- **AI Credits persistence (Day 2):** 8 workflows still failing; issues being re-created daily → root budget config fix not applied
- **Failure Investigator working:** #38309 correctly aggregates Auto-Triage + Sub-Issue Closer failures into composite issue
- **action_required pattern (normal):** AI Moderator, Q — EXPECTED behavior; human review requesting
- **Recovery holding:** Ecosystem health 87/100; all Jun 9 P0/P1 closed

## Coverage Notes
- Strong: PR code review, dep management, security/moderation, repo hygiene, observability
- Weak: Daily News (failing), Daily Ambient Context (failing), memory/* bootstrap lifecycle
- Risk: AI credits cluster unresolved Day 2 — affects health monitoring, test quality, optimization workflows

## Coordination Notes for Other Orchestrators
- **WH:** AI credits cluster Day 2 — health monitoring (WH Manager) still at risk; budget fix needed this sprint
- **Campaign:** copilot-swe-agent throughput healthy, 8 active PRs; complex campaigns viable
- **All:** AI credits budget config must be escalated; daily re-creation of same failure issues is noise
- **All:** memory/* branch lifecycle — new runbook needed for signed-commit seed on first push
