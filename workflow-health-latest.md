# Workflow Health — 2026-06-16T06:17Z

Score: 70/100 (↓2 from 72)
Workflows: 249 | Lock files: 249/249 (100% ✅) | Run: §27598171193

## KEY FINDINGS

### Status (June 16)
- **Compilation:** 249/249 workflows have lock files (100% ✅)
- **PR Sous Chef (P1 → RESOLVED):** 9+ consecutive successes since ~19:58 Jun 15. Commented on #39330. FULLY RECOVERED.
- **Tool Denial Cluster (NEW SYSTEMIC P1):** 6+ workflows hitting 5/5 tool denial threshold. Systemic issue filed today.
- **Daily Safe Outputs Git Simulator (Day 8+):** branch missing. In progress today. DO NOT RE-FILE.
- **Code Simplifier (Day 9+, #39199/#39489):** api-proxy cap + HTTP 429. Ongoing.
- **Daily BYOK Ollama Test (Day 8, #39476):** transient_bad_request. Filed today. DO NOT RE-FILE.
- **Daily Safe Output Integrator (Day 8, #39477/#39449):** tool denial 5/5. Filed today. DO NOT RE-FILE.
- **Daily Model Inventory Checker (Day 7, #39471):** session.idle 60s. Auto-filed today. DO NOT RE-FILE.
- **AI Moderator (#39452):** intermittent pre_activation pagination timeout. Mostly OK Jun 16.
- **Daily Cache Strategy Analyzer (#39451):** Codex gpt-5-codex 404. Alternating ~50% success.

### Resolved Since Jun 15 ✅
- **PR Sous Chef:** 9+ consecutive successes Jun 15–16 — FULLY RECOVERED. Commented on #39330.

### Critical Issues (P1) 🚨
- **Tool Denial Cluster** (SYSTEMIC, NEW): 6+ workflows — Daily Safe Output Integrator, Daily Agent of the Day Blog Writer, Breaking Change Checker, Delight, jsweep, Daily Compiler Threat Spec Optimizer. Systemic issue filed today. DO NOT RE-FILE per-workflow.
- **Daily Safe Outputs Git Simulator** (Day 8+): branch missing. DO NOT RE-FILE.
- **Code Simplifier** (#39199, Day 9+): api-proxy cap. DO NOT RE-FILE.
- **Daily BYOK Ollama Test** (#39476, Day 8): transient_bad_request. DO NOT RE-FILE.
- **Daily Safe Output Integrator** (#39477, Day 8): tool denial. DO NOT RE-FILE.
- **Daily Model Inventory Checker** (#39471, Day 7): session.idle 60s. DO NOT RE-FILE.
- **upload_artifact malformed 400** (#38998): Smoke Copilot. DO NOT RE-FILE.
- **Smoke Gemini** (#39172): preview model error. DO NOT RE-FILE.
- **AI Moderator** (#39452): intermittent pagination timeout. Watch but DO NOT RE-FILE.
- **Daily Cache Strategy Analyzer** (#39451): Codex 404. DO NOT RE-FILE.

### Warnings (P2) ⚠️
- **Daily Compiler Threat Spec Optimizer** (#39343): tool denial (weekly pattern). Watch.
- **AIC Budget Crisis** (#39077): Day 9 — root fix still unimplemented.
- **Smoke Trigger & Multi Caller** (#38999): 100% startup_failure. DO NOT RE-FILE.
- **Dictation Prompt Generator** (#39196, #39200): session.idle 570s. DO NOT RE-FILE.

### Systemic Patterns
- **Tool denial cluster EXPANDING**: 6+ workflows hitting 5/5 `tool_denials_exceeded` — likely workflows using `git checkout -b` or shell commands that are blocked by the allow-list. Systemic issue filed today.
- **session.idle timeout cluster**: Model Inventory (60s) + Dictation Prompt Generator (570s) — platform-level Copilot SDK regression likely.
- **AIC cluster**: Code Simplifier Day 9 — root fix still blocking. #39479 filed Jun 16 for non-retryable 429 handling.
- **Codex model 404**: Daily Cache Strategy Analyzer uses gpt-5-codex-alpha-2025-11-07 (returns 404).

### Actions Taken This Run
- 1 comment added to #29109 (health dashboard Jun 16)
- 1 comment added to #39330 (PR Sous Chef recovery confirmation)
- 1 systemic issue created (tool denial cluster)

## Do Not Re-File
- PR Sous Chef: RESOLVED (commented on #39330)
- Tool denial cluster: systemic issue filed today
- #39477: Daily Safe Output Integrator Day 8 (P1)
- #39476: Daily BYOK Ollama Test Day 8 (P1)
- #39471: Daily Model Inventory Checker Day 7
- #39199/#39489: Code Simplifier Day 9+
- #38998: upload_artifact malformed 400
- #38999: Smoke Trigger startup_failure
- #38993: Agentic workflows out of sync
- #39024: Daily Safe Outputs Git Simulator (expired, DO NOT RE-FILE)
- #39037: Failure Investigator pre-fetch blind spot
- #38870/#38871/#38872: Performance regression cluster
- #39077: AIC Budget Crisis
- #39172: Smoke Gemini preview model
- #39196/#39200: Dictation Prompt Generator
- #39193: GitHub Remote MCP Auth missing tool
- #39452: AI Moderator pagination timeout
- #39451: Daily Cache Strategy Analyzer Codex 404
- #39343: Daily Compiler Threat Spec Optimizer tool denial (P2)
