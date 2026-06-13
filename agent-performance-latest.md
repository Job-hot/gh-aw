# Agent Performance Analyzer — Latest Run

**Timestamp:** 2026-06-13T13:15:00Z
**Workflow:** Agent Performance Analyzer — Meta-Orchestrator
**Run ID:** 27467672212
**Run URL:** https://github.com/github/gh-aw/actions/runs/27467672212

## Executive Summary

- **Agents analyzed:** 22 active workflow types (~246 total workflows)
- **Quality score:** 57/100 (↓5 from 62)
- **Effectiveness score:** 55/100 (↓3 from 58)
- **Ecosystem health:** 72/100 (↓3 from 75)
- **AIC cluster:** 6 agents blocked today (up from 5 yesterday) — root fix STILL pending Day 5+

## Active P0/P1 Issues (Do Not Re-File)

- **Code Simplifier Day 5** (#39013, OPEN, HTTP 429 rate-limited): Root fix pending. DO NOT RE-FILE.
- **upload_artifact malformed 400** (#38998, OPEN, P1): Smoke Copilot ~95% failure. DO NOT RE-FILE.
- **Daily Safe Outputs Git Simulator** (#39024, OPEN, P1, 5-day streak): memory/git-simulator branch missing. DO NOT RE-FILE.
- **Smoke Trigger/Multi Caller startup_failure** (#38999, OPEN, P2). DO NOT RE-FILE.
- **Agentic workflows out of sync** (#38993, OPEN). DO NOT RE-FILE.
- **Failure Investigator pre-fetch blind spot** (#39037, OPEN): empty failed_run_ids. DO NOT RE-FILE.
- **AIC cluster systemic tracker** (#aw_aic_day5, NEW this run): 6-agent cluster escalation. DO NOT RE-FILE.

## New Issues Filed This Run
- **1 systemic issue**: AIC Budget Crisis Day 5 — 6-agent cluster escalation (#aw_aic_day5)

## Agent Rankings

### Top Performers (Q/E scores)
1. Agentic Maintenance (Q:82, E:88) — 100% success, 2m18s
2. Bot Detection (Q:80, E:90) — 100% success, 13s
3. Auto-Triage Issues (Q:80, E:82) — 100% success, 5m24s
4. Issue Monster (Q:78, E:85) — 100% success, 30s
5. PR Sous Chef (Q:78, E:80) — 100% success, 9m59s
6. PR Triage Agent (Q:78, E:80) — 100% success, 9m14s
7. Claude Code User Documentation Review (Q:75, E:78) — 100% success, 14m58s
8. Lint Monster (Q:75, E:78) — Actionable refactoring issues (#39003, #39004)
9. Daily A/B Testing Advisor (Q:75, E:72) — Active experiments (#39062, #39063)
10. Duplicate Code Detector (Q:72, E:75) — 3 issues filed (#39026-#39028)

### Underperformers
- Code Simplifier (Q:10, E:5) — CRITICAL: Day 5 HTTP 429 (#39013)
- Dev (Q:30, E:20) — No safe outputs (#39046)
- Test Quality Sentinel (Q:35, E:25) — AIC exceeded (#39059)
- Matt Pocock Skills Reviewer (Q:35, E:25) — AIC exceeded (#39050)
- Daily CLI Tools Exploratory Tester (Q:40, E:30) — AIC rate limit (#39031)
- jsweep (Q:42, E:35) — Tool denial (#39020)
- Copilot CLI Deep Research (Q:42, E:32) — Tool denial (#39022)
- Failure Investigator (6h) (Q:42, E:35) — Pre-fetch blind spot (#39037)
- Avenger (Q:55, E:45) — Failed today (#39073)

## Pattern Detection
- **AIC cluster EXPANDING**: 6 agents at/near limit today (up from 5 yesterday). Root fix Day 5+.
- **Performance regression spike (NEW)**: 3 compile ops 165-269% slower (#38870-#38872)
- **Code quality momentum**: Lint Monster + Duplicate Detector + Static Analysis all productive
- **Tool denial cluster**: jsweep + Copilot CLI Deep Research hitting tool denial limits

## Coordination Notes for Other Orchestrators
- **WH:** AIC cluster now 6 agents (escalating); performance regression spike (#38870-#38872) is new today
- **Campaign:** Code quality momentum good (Lint Monster + Duplicate Detector + Static Analysis)
- **All:** AIC root fix (raise max-ai-credits to 2000) STILL unimplemented Day 5. ESCALATE NOW.
- **All:** Do NOT re-file issues listed in shared-alerts.md — check before creating trackers
