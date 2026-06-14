# Agent Performance Analyzer — Latest Run

**Timestamp:** 2026-06-14T13:15:31Z
**Workflow:** Agent Performance Analyzer — Meta-Orchestrator
**Run ID:** 27499941088
**Run URL:** https://github.com/github/gh-aw/actions/runs/27499941088

## Executive Summary

- **Agents analyzed:** 25 active workflow types (~246 total workflows)
- **Quality score:** 57/100 (→ stable vs yesterday's 57)
- **Effectiveness score:** 55/100 (→ stable)
- **Ecosystem health:** 71/100 (↓2 from 73)
- **AIC today:** 19,985 (↓27% from Jun 12 spike, but still +48% above 4-day baseline of ~13,496)
- **AIC Budget Crisis:** Day 7 — Code Simplifier now hitting api-proxy invocation cap (maxRuns 50/50) in addition to HTTP 429

## Active P0/P1 Issues (Do Not Re-File)

- **Code Simplifier Day 7** (#39179 failure, #39199 P1 api-proxy cap, #39077 AIC root fix): DO NOT RE-FILE.
- **upload_artifact malformed 400** (#38998, P1): Smoke Copilot partial dark. DO NOT RE-FILE.
- **Daily Safe Outputs Git Simulator Day 7** (#39024, P1): Issue Monster queued Copilot fix. DO NOT RE-FILE.
- **Daily Model Inventory Checker Day 5+** (#39170, P1): session.idle timeout. DO NOT RE-FILE.
- **Smoke Trigger & Multi Caller** (#38999, P2). DO NOT RE-FILE.
- **Failure Investigator pre-fetch gap** (#39037). DO NOT RE-FILE.
- **Performance regression cluster** (#38870-38872). DO NOT RE-FILE.
- **Dictation Prompt Generator** (#39196 failure, #39200 P2 investigation): Filed today. DO NOT RE-FILE.
- **Smoke Gemini** (#39172, P1): Filed today. DO NOT RE-FILE.
- **GitHub Remote MCP Auth Test** (#39193): Filed today. DO NOT RE-FILE.

## New Issues Filed This Run
- None new — all actionable items already tracked by Failure Investigator and other agents.

## Agent Rankings

### Top Performers (Q/E scores)
1. Agentic Maintenance (Q:82, E:88) — 100% success
2. Bot Detection (Q:80, E:90) — 100% success, ~13s runtime
3. Auto-Triage Issues (Q:80, E:82) — 100% success
4. Issue Monster (Q:78, E:85) — 100% success, proactively queued Git Simulator fix
5. PR Sous Chef (Q:78, E:80) — 100% success (13 runs today, 4 minor errors in tool calls)
6. PR Triage Agent (Q:78, E:80) — 100% success
7. Claude Code User Documentation Review (Q:75, E:78)
8. Avenger (Q:72, E:75) — RECOVERED: 2 successful runs today (after failure yesterday)
9. Daily A/B Testing Advisor (Q:75, E:72)
10. Package Specification Enforcer (Q:72, E:75)

### Underperformers
- Code Simplifier (Q:10, E:5) — CRITICAL Day 7: now api-proxy cap (maxRuns 50/50) + HTTP 429 (#39199)
- Dictation Prompt Generator (Q:30, E:20) — NEW failure today: session.idle 570s timeout (#39196, #39200)
- Smoke Gemini (Q:25, E:15) — NEW: using preview TTS model → unknown model error (#39172)
- GitHub Remote MCP Auth Test (Q:30, E:20) — NEW: missing required tool (#39193)
- Daily Model Inventory Checker (Q:35, E:25) — Day 5 streak (#39170)
- jsweep (Q:42, E:35) — Tool denial cluster (#39185), high AIC (1316/run)
- Daily CLI Tools Exploratory Tester (Q:40, E:30) — AIC rate limits

## Pattern Detection
- **Session.idle timeout cluster EXPANDING**: Dictation Prompt Generator (570s) joins Model Inventory Checker (60s) — may be platform-level Copilot SDK regression
- **Code Simplifier escalation**: Added api-proxy cap exhaustion on top of existing HTTP 429 — hardening required
- **AIC normalization but still elevated**: 19,985 today (↓27% from Jun 12 spike) but +48% above 4-day baseline
- **Avenger recovery**: 2 successful runs after yesterday's failure — positively signals resilience
- **copilot-swe-agent PR momentum**: 77% merge rate (23/30 PRs merged in 7d) — strong performer

## Coordination Notes for Other Orchestrators
- **WH:** Dictation Prompt Generator joins session.idle timeout cluster — monitor for platform regression; 2 new P1s today (Smoke Gemini, GitHub Remote MCP Auth)
- **Campaign:** Avenger recovery positive; jsweep optimizer filed issue #39208 showing 1315 AIC/run
- **All:** AIC root fix (raise max-ai-credits to 2000) STILL unimplemented Day 7. ROOT CAUSE: Code Simplifier now both HTTP 429 AND api-proxy cap exhausted. ESCALATE IMMEDIATELY.
- **All:** Do NOT re-file issues listed above — check before creating trackers
