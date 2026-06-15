# Workflow Health — 2026-06-15T06:13Z

Score: 72/100 (↓1 from 73)
Workflows: 246 | Lock files: 246/246 (100% ✅) | Run: §27527526423

## KEY FINDINGS

### Status (June 15)
- **Compilation:** 246/246 workflows have lock files (100% ✅)
- **PR Sous Chef (P1, NEW):** "Fetch open non-draft PR queue" step failing since 22:31 UTC Jun 14 — 8+ consecutive failures (every 15 min). Filed today → #aw_sous_chef_p1
- **Daily Model Inventory Checker (P1, Day 6):** session.idle timeout 60000ms. #39170 expired/closed. Re-filed → #aw_model_inv_p1
- **Code Simplifier (P1, Day 8, #39199):** api-proxy cap (maxRuns 50/50) + HTTP 429. Root fix still pending.
- **upload_artifact malformed 400 (#38998):** Smoke Copilot partial failures continuing.
- **Daily Safe Outputs Git Simulator (Day 7, #39024 expired):** memory/git-simulator branch still missing. No Copilot PR yet.
- **Smoke Gemini (#39172):** preview model unknown error.

### Resolved Since Jun 14 ✅
- **Avenger:** 3 successful runs Jun 15 — fully recovered (no more monitoring needed).

### Critical Issues (P1) 🚨
- **PR Sous Chef** (NEW, P1): "Fetch open non-draft PR queue" failing 8+ runs. Filed today. DO NOT RE-FILE.
- **Daily Model Inventory Checker** (Day 6, re-filed): session.idle 60s timeout. Filed today. DO NOT RE-FILE.
- **Code Simplifier** (#39199, P1, Day 8): api-proxy cap exhaustion. DO NOT RE-FILE.
- **upload_artifact malformed** (#38998, P1): Smoke Copilot partial dark. DO NOT RE-FILE.
- **Daily Safe Outputs Git Simulator** (#39024 expired, Day 7): branch missing. DO NOT RE-FILE (Issue Monster queued fix).
- **Smoke Gemini** (#39172, P1): preview model error. DO NOT RE-FILE.

### Warnings (P2) ⚠️
- **Daily Compiler Threat Spec Optimizer:** Intermittent failures every ~1 week (May 24, Jun 8, Jun 15). Watch.
- **Sergo - Serena Go Expert:** 1 failure Jun 15 after 6-day success streak. Likely transient.
- **AIC Budget Crisis** (#39077): Day 8 — root fix still unimplemented.
- **Smoke Trigger & Multi Caller** (#38999): 100% startup_failure. DO NOT RE-FILE.
- **Failure Investigator blind spot** (#39037). DO NOT RE-FILE.
- **Dictation Prompt Generator** (#39196, #39200): session.idle 570s. DO NOT RE-FILE.

### Systemic Patterns
- session.idle timeout cluster: Model Inventory (60s) + Dictation Prompt Generator (570s) — may be platform-level Copilot SDK regression.
- PR Sous Chef queue fetch failure = possible GitHub API or script change at ~22:31 Jun 14.
- AIC cluster Day 8 — root fix still blocking Code Simplifier.

### Actions Taken This Run
- 1 issue created: PR Sous Chef P1 regression
- 1 issue created: Daily Model Inventory Checker Day 6 tracker
- 1 comment added to #39199 (Code Simplifier Day 8)
- 1 comment added to #29109 (health dashboard Jun 15)

## Do Not Re-File
- PR Sous Chef: filed today (see #aw_sous_chef_p1)
- Daily Model Inventory Checker Day 6: filed today (see #aw_model_inv_p1)
- #39199: Code Simplifier Day 8 (P1)
- #38998: upload_artifact malformed 400 (P1)
- #38999: Smoke Trigger startup_failure (P2)
- #38993: Agentic workflows out of sync
- #39024: Daily Safe Outputs Git Simulator expired (DO NOT RE-FILE — Issue Monster queued fix)
- #39037: Failure Investigator pre-fetch blind spot
- #38870/#38871/#38872: Performance regression cluster
- #39077: AIC Budget Crisis Day 8 (escalating)
- #39172: Smoke Gemini preview model (P1)
- #39196/#39200: Dictation Prompt Generator (P2)
- #39193: GitHub Remote MCP Auth missing tool
