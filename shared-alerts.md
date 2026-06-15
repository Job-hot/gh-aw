# Shared Alerts — 2026-06-15T06:13Z (updated by Workflow Health Manager)

## P1 (High) 🚨
- **PR Sous Chef — NEW regression** (P1): "Fetch open non-draft PR queue" step failing since ~22:31 UTC Jun 14, 8+ consecutive failures. Filed today. DO NOT RE-FILE.
- **Daily Model Inventory Checker — Day 6** (P1): session.idle timeout 60000ms. #39170 expired; re-filed today. DO NOT RE-FILE.
- **AIC Budget Crisis — Day 8, Code Simplifier** (#39077, #39199): api-proxy cap (maxRuns 50/50) + HTTP 429. Root fix STILL PENDING DAY 8. ESCALATE. DO NOT RE-FILE.
- **upload_artifact malformed 400** (#38998): Smoke Copilot ~95% failure. DO NOT RE-FILE.
- **Daily Safe Outputs Git Simulator — 7-day streak** (#39024 expired): memory/git-simulator branch missing. Issue Monster queued Copilot fix. DO NOT RE-FILE.
- **Smoke Gemini** (#39172): preview TTS model → unknown model error. DO NOT RE-FILE.

## P2 (Watch) ⚠️
- **Smoke Trigger & Smoke Multi Caller** (#38999): 100% startup_failure. DO NOT RE-FILE.
- **Agentic workflows out of sync** (#38993). DO NOT RE-FILE.
- **Failure Investigator pre-fetch gap** (#39037). DO NOT RE-FILE.
- **Performance regression spike** (#38870, #38871, #38872): Compile ops 165-269% slower. DO NOT RE-FILE.
- **Dictation Prompt Generator** (#39196, #39200): session.idle 570s timeout. DO NOT RE-FILE.
- **GitHub Remote MCP Auth Test** (#39193): missing required tool. DO NOT RE-FILE.
- **Daily Compiler Threat Spec Optimizer**: Intermittent failures (May 24, Jun 8, Jun 15 — escalating to weekly). Watch but no issue yet.
- **Sergo - Serena Go Expert**: 1 failure Jun 15 (likely transient after 6-day success streak). Monitor.

## Resolved (Jun 14→15) ✅
- **Avenger**: 3 successful runs Jun 15 — FULLY RECOVERED. Remove from monitoring.

## Systemic Notes
- **session.idle timeout cluster EXPANDING**: Model Inventory (60s) + Dictation Prompt Generator (570s) — platform-level Copilot SDK regression likely. New P1 filed today for Model Inventory.
- **PR Sous Chef regression**: High-frequency workflow (every 15 min) failing at queue-fetch step. Started ~22:31 Jun 14. Possible GitHub API issue or script change.
- **AIC today**: Still elevated. Day 8 of Code Simplifier crisis.
- **Compilation**: 246/246 (100%) — all lock files present.
- **Avenger**: Fully recovered 3/3 runs Jun 15 — no longer a watch item.

## Do Not Re-File (All Active Issues)
- PR Sous Chef regression: filed today
- Daily Model Inventory Checker Day 6: filed today
- #39199: Code Simplifier api-proxy cap exhaustion (P1)
- #38998: upload_artifact malformed 400 (P1)
- #38999: Smoke Trigger startup_failure (P2)
- #38993: Agentic workflows out of sync
- #39024: Daily Safe Outputs Git Simulator (expired, DO NOT RE-FILE — Issue Monster queued fix)
- #39170: Daily Model Inventory Checker (closed, superseded by today's re-file)
- #39037: Failure Investigator pre-fetch blind spot
- #38870/#38871/#38872: Performance regression cluster
- #39077: AIC Budget Crisis Day 8 (escalating)
- #39172: Smoke Gemini preview model (P1)
- #39196/#39200: Dictation Prompt Generator (P2)
- #39193: GitHub Remote MCP Auth missing tool
