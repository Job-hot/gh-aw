# Agent Performance Analyzer — Latest Run

**Timestamp:** 2026-06-11T14:15:00Z  
**Workflow:** Agent Performance Analyzer — Meta-Orchestrator  
**Run ID:** 27352042326  
**Run URL:** https://github.com/github/gh-aw/actions/runs/27352042326

## Executive Summary

- **Agents analyzed:** 22 active workflow types (~55 active workflows, 96 runs in 24h window)
- **Quality score:** 67/100 (→ stable)
- **Effectiveness score:** 63/100 (→ stable)
- **Ecosystem health:** 83/100 (↓2 from 85)
- **Total AIC (24h):** 10,907.71 (−18.5% from yesterday)
- **Total Actions minutes (24h):** 1,235 min

## Active P0/P1 Issues (Do Not Re-File)

- **AI Credits Cluster — now 6+ workflows** (#38520 CLOSED, no open tracker): Code Simplifier, Ubuntu Image Analyzer, Workflow Skill Extractor, Matt Pocock Skills Reviewer, Package Specification Extractor (1,023.9 AIC — largest single consumer), Failure Investigator (6h) (over by 25.1 AIC). Root fix: raise max-ai-credits to 2000 for analysis-heavy workflows.
  - Individual failure issues: #38499, #38500, #38501, #38497, #38576, #38559 — DO NOT RE-FILE
- **Daily News chroot** (#38379, assigned @zarenner, Day 3+) — DO NOT RE-FILE
- **Daily Safe Outputs Git Simulator** (#aw_gitsim10, ongoing) — DO NOT RE-FILE

## New Issues Filed This Run
- **1 improvement issue created**: raise max-ai-credits for Failure Investigator (6h) — critical meta-monitor blind spot
- **Performance discussion created**

## Agent Rankings

### Top Performers
1. copilot-swe-agent (Q:90, E:88) — 5 active PRs (#38615, #38616, #38611, #38610, #38573); systematic improvements to failure classification, AIC detection, metric renames
2. Agentic Maintenance (Q:85, E:90) — 100% success, multiple runs today
3. Auto-Triage Issues (Q:82, E:88) — 5+ successful runs today (100%)
4. Daily AIC Usage Audit (Q:82, E:80) — comprehensive 96-run telemetry (#38609)
5. AB Testing Advisor (Q:82, E:78) — high-quality hypothesis-driven issues (#38591, #38590)
6. Daily File Diet (Q:80, E:85) — 5-day success streak
7. PR Triage Agent (Q:80, E:75) — accurate PR categorization (#38619)
8. Content Moderation (Q:78, E:82) — 4/4 success today
9. Bot Detection / Avenger (Q:78, E:80) — 100% success
10. Daily AIC Consumption Report (Q:75, E:72) — OTLP type mismatch discovery (#38612)

### Needing Improvement
- Package Specification Extractor (Q:35, E:25) — highest single AIC consumer (1,023.9); hitting 1K limit
- Matt Pocock Skills Reviewer (Q:35, E:25) — AI credits Day 3+ (#38497, #38617)
- Failure Investigator (6h) (Q:35, E:20) — meta-failure: AI credits exhaustion (#38559)
- Daily News (Q:35, E:20) — Node.js chroot failure Day 3+ (#38379)
- Daily AW Cross-Repo Compile Check (Q:40, E:30) — failed (#38570), high AIC risk (653.58)
- GitHub MCP Structural Analysis (Q:40, E:30) — failed (#38614)
- Daily MCP Tool Concurrency Analysis (Q:45, E:30) — failed (#38578)
- PR Sous Chef (Q:45, E:35) — failed today (#38618)
- Daily Safe Outputs Conformance Checker (Q:45, E:30) — failed (#38540)

## Pattern Detection
- **AI credits cluster EXPANDING**: Now 6+ workflows at 1K limit (up from 4 in #38520). Package Specification Extractor became largest single-run consumer (1,023.9 AIC). Failure Investigator is now also a victim (#38559) — meta-monitor blind spot.
- **copilot-swe-agent healthy throughput**: 5 active PRs on systematic improvements (failure titles, AIC detection, metric renames, log enhancement)
- **Smoke test flakiness**: Multiple engines (Copilot, Codex, Pi, Antigravity, AOAI) failing — likely transient / test-env issues
- **action_required (normal)**: AI Moderator, Q, PR Description Updater, Label Closed PRs — EXPECTED behavior

## Coverage Notes
- Strong: PR code review, dependency management, security/moderation, repo hygiene, token observability, AB testing
- Weak: Daily News (failing), Windows integration (timeout), memory branch bootstrap
- Risk: AI credits cluster unresolved Day 3 — now affects meta-monitor (Failure Investigator)

## Coordination Notes for Other Orchestrators
- **WH:** AI credits cluster grew to 6+ workflows; Failure Investigator now in cluster — update monitoring scope
- **Campaign:** copilot-swe-agent healthy, 5 active PRs; systematic AIC / failure-title improvements in flight
- **All:** AI credits root fix (raise max-ai-credits to 2000) MUST be applied to Failure Investigator first — meta-monitor reliability critical
- **All:** Package Specification Extractor is now the largest single-run AIC consumer (1,023.9) — needs separate budget tuning
