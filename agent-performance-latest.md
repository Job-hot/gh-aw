# Agent Performance — Latest Run

**Timestamp:** 2026-06-15T14:49:18Z | **Run:** [§27554250019](https://github.com/github/gh-aw/actions/runs/27554250019)

## Summary: 56/100 Quality | 54/100 Effectiveness | 70/100 Health | 17,404 AIC (↓12.9%)

## Top Performers
1. Bot Detection (Q:82, E:91) — 100% success, ~13s
2. Agentic Maintenance (Q:82, E:88) — continued success
3. Auto-Triage Issues (Q:81, E:84) — 2/2 runs
4. Avenger (Q:74, E:82) — 🆙 FULLY RECOVERED (3/3 Jun 15)
5. Issue Monster (Q:78, E:85), copilot-swe-agent (Q:78, E:82, 77% merge rate)

## Underperformers
- Code Simplifier (Q:10, E:5) P1 Day 8: #39199 api-proxy cap + HTTP 429. Root fix #39077 pending.
- PR Sous Chef (Q:50, E:35) P1: queue-fetch failing 8+ runs (#39352)
- Daily Model Inventory (Q:35, E:25) Day 6 session.idle (#refiled today)
- Dictation Prompt Generator (Q:30, E:20) session.idle 570s (#39196, #39200)
- Smoke Gemini (Q:25, E:15) preview model error (#39172)
- GitHub Remote MCP Auth (Q:30, E:20) missing tool (#39193)

## New Failures Today (single-day; monitor)
#39385 Matt Pocock, #39384 Test Quality Sentinel, #39383 Design Decision Gate, #39353 Daily News, #39347 Layout Spec Maintainer (tool denial)

## Key Patterns
- **Tool denial cluster EXPANDING**: 3 workflows today (Layout Spec, MCP Concurrency, Compiler Threat Spec)
- **Session.idle cluster**: 2 workflows (platform SDK regression; watchdog #39282 merged today)
- **AIC declining**: 17,404 (↓12.9%); still ~29% above 4d baseline
- **Campaign Manager offline**: `campaign-manager-latest.md` absent — coordination gap
- **AIC observability gap**: gh-aw.aic Sentry field unqueryable (#39380)

## Do Not Re-File
#39199 Code Simplifier, #39352 PR Sous Chef, #38998 upload_artifact, #39024 Git Simulator, Model Inventory (re-filed today), #39172 Smoke Gemini, #39193 Remote MCP Auth, #39196/#39200 Dictation, #38999 Smoke Trigger, #39037 Failure Investigator, #38870-38872 Perf regression, #39077 AIC Crisis, #39345 Sergo, #39343 Compiler Threat Spec, #39347 Layout Spec, #39366 MCP Concurrency, #39383-39385 today's new failures, #39353 Daily News

## New Issues Filed: None (all tracked)
