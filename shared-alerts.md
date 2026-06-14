# Shared Alerts — 2026-06-14T13:15Z (updated by Agent Performance Analyzer)

## P1 (High) 🚨
- **AIC Budget Crisis — Day 7, Code Simplifier** (#39077, #39179, #39199): Code Simplifier now hitting BOTH HTTP 429 AND api-proxy invocation cap (maxRuns 50/50). Root fix (raise max-ai-credits to 2000, max-turns:30) STILL PENDING DAY 7. ESCALATE. DO NOT RE-FILE.
- **upload_artifact malformed 400** (#38998): Smoke Copilot ~95% failure. DO NOT RE-FILE.
- **Daily Safe Outputs Git Simulator — 7-day streak** (#39024): memory/git-simulator branch missing. Issue Monster queued Copilot fix. DO NOT RE-FILE.
- **Daily Model Inventory Checker — 5-day streak** (#39170): copilot-sdk session.idle timeout 60000ms. DO NOT RE-FILE.
- **Smoke Gemini** (#39172): preview TTS model → unknown model error. Filed today. DO NOT RE-FILE.

## P2 (Watch) ⚠️
- **Smoke Trigger & Smoke Multi Caller** (#38999): 100% startup_failure. DO NOT RE-FILE.
- **Agentic workflows out of sync** (#38993). DO NOT RE-FILE.
- **Failure Investigator pre-fetch gap** (#39037). DO NOT RE-FILE.
- **Performance regression spike** (#38870, #38871, #38872): Compile ops 165-269% slower. DO NOT RE-FILE.
- **Dictation Prompt Generator** (#39196, #39200): session.idle 570s timeout. Filed today. DO NOT RE-FILE.
- **GitHub Remote MCP Auth Test** (#39193): missing required tool. Filed today. DO NOT RE-FILE.

## Resolved (Jun 13→14) ✅
- **Avenger**: Recovered from yesterday's failure — 2 successful runs Jun 14. Monitor.

## Systemic Notes
- **Session.idle timeout cluster**: Dictation Prompt Generator (570s) + Daily Model Inventory Checker (60s). May be platform-level Copilot SDK regression. Watch for more agents joining.
- **AIC today**: 19,985 (↓27% from Jun 12 spike, but +48% above 4-day baseline ~13,496). Elevated but not spiking.
- **copilot-swe-agent PR merge rate**: 77% (23/30 merged in 7d) — strong performance.
- **app/github-actions PR merge rate**: 40% (4/10 merged in 7d) — below target.

## Do Not Re-File (All Active Issues)
- #39199: Code Simplifier api-proxy cap exhaustion (P1, filed today)
- #39179: Code Simplifier failure notice (P1)
- #38998: upload_artifact malformed 400 (P1)
- #38999: Smoke Trigger startup_failure (P2)
- #38993: Agentic workflows out of sync
- #39024: Daily Safe Outputs Git Simulator Day 7 (P1)
- #39170: Daily Model Inventory Checker 5-day tracker (P1)
- #39037: Failure Investigator pre-fetch blind spot
- #38870/#38871/#38872: Performance regression cluster
- #39077: AIC Budget Crisis Day 7 (escalating)
- #39172: Smoke Gemini preview model (P1)
- #39196: Dictation Prompt Generator failure notice (P2)
- #39200: Dictation Prompt Generator investigation (P2)
- #39193: GitHub Remote MCP Auth missing tool
