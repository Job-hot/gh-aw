# Shared Alerts — 2026-06-16T06:17Z (updated by Workflow Health Manager)

## P1 (High) 🚨
- **Tool Denial Cluster — SYSTEMIC (NEW)**: 6+ workflows hitting 5/5 `tool_denials_exceeded`: Daily Safe Output Integrator, Daily Agent of the Day Blog Writer, Breaking Change Checker, Delight, jsweep, Daily Compiler Threat Spec Optimizer. Pattern: workflows try `git checkout -b` / shell git ops blocked by allow-list. Systemic issue filed today. DO NOT RE-FILE per-workflow.
- **Daily Model Inventory Checker — Day 7** (#39471, auto-filed today): session.idle timeout 60000ms. DO NOT RE-FILE.
- **AIC Budget Crisis — Day 9, Code Simplifier** (#39077, #39199, #39489): api-proxy cap (maxRuns 50/50) + HTTP 429. Root fix STILL PENDING DAY 9. #39479 filed Jun 16 for non-retryable 429 handling. DO NOT RE-FILE.
- **upload_artifact malformed 400** (#38998): Smoke Copilot ~95% failure. DO NOT RE-FILE.
- **Daily Safe Outputs Git Simulator — 8-day streak** (#39024 expired): memory/git-simulator branch missing. Issue Monster queued Copilot fix. DO NOT RE-FILE.
- **Smoke Gemini** (#39172): preview TTS model → unknown model error. DO NOT RE-FILE.
- **Daily BYOK Ollama Test — Day 8** (#39476): transient_bad_request every inference call. DO NOT RE-FILE.
- **Daily Safe Output Integrator — Day 8** (#39477): tool denial 5/5 + ECONNREFUSED in fallback. DO NOT RE-FILE.
- **AI Moderator** (#39452, Jun 15): intermittent pre_activation pagination timeout (page 205+). Mostly OK Jun 16 (recovering?). Watch.
- **Daily Cache Strategy Analyzer** (#39451, Jun 15): Codex gpt-5-codex-alpha-2025-11-07 returns 404. ~50% alternating success rate. DO NOT RE-FILE.

## P2 (Watch) ⚠️
- **Smoke Trigger & Smoke Multi Caller** (#38999): 100% startup_failure. DO NOT RE-FILE.
- **Agentic workflows out of sync** (#38993). DO NOT RE-FILE.
- **Failure Investigator pre-fetch gap** (#39037). DO NOT RE-FILE.
- **Performance regression spike** (#38870, #38871, #38872): Compile ops 165-269% slower. DO NOT RE-FILE.
- **Dictation Prompt Generator** (#39196, #39200): session.idle 570s timeout. DO NOT RE-FILE.
- **GitHub Remote MCP Auth Test** (#39193): missing required tool. DO NOT RE-FILE.
- **Daily Compiler Threat Spec Optimizer** (#39343, P2): tool denial cluster (weekly pattern). Watch.
- **PR Sous Chef** (#39330): RESOLVED — 9+ consecutive successes Jun 15–16. FULLY RECOVERED. Close issue recommended.

## Resolved (Jun 15→16) ✅
- **PR Sous Chef**: 9+ consecutive successes — FULLY RECOVERED. Commented on #39330.
- **Avenger**: Still healthy.

## Systemic Notes
- **Tool denial cluster EXPANDING**: 6+ workflows — all individually filing issues but root cause is allow-list blocking `git checkout -b` and other shell git ops. Systemic issue filed Jun 16.
- **session.idle timeout cluster EXPANDING**: Model Inventory (60s) + Dictation Prompt Generator (570s) — platform-level Copilot SDK regression likely.
- **AIC cluster Day 9**: Code Simplifier still blocked. #39479 addresses non-retryable 429 handling.
- **Codex model 404**: Daily Cache Strategy Analyzer uses deprecated model path.
- **Compilation**: 249/249 (100%) — all lock files present.

## Do Not Re-File (All Active Issues)
- Tool denial cluster: systemic issue filed Jun 16
- PR Sous Chef recovery: commented on #39330 (RESOLVED)
- Daily Model Inventory Checker Day 7: auto-filed today #39471
- #39199/#39489: Code Simplifier Day 9 (P1)
- #38998: upload_artifact malformed 400 (P1)
- #38999: Smoke Trigger startup_failure (P2)
- #38993: Agentic workflows out of sync
- #39024: Daily Safe Outputs Git Simulator (expired, DO NOT RE-FILE — Issue Monster queued fix)
- #39037: Failure Investigator pre-fetch blind spot
- #38870/#38871/#38872: Performance regression cluster
- #39077: AIC Budget Crisis Day 9 (escalating)
- #39172: Smoke Gemini preview model (P1)
- #39196/#39200: Dictation Prompt Generator (P2)
- #39193: GitHub Remote MCP Auth missing tool
- #39476: Daily BYOK Ollama Test Day 8 (P1)
- #39477: Daily Safe Output Integrator Day 8 (P1)
- #39452: AI Moderator pagination timeout (P1, intermittent)
- #39451: Daily Cache Strategy Analyzer Codex 404 (P1)
- #39343: Daily Compiler Threat Spec Optimizer tool denial (P2)

---
# Agent Performance Update — 2026-06-16T14:30Z

## New Observations (for Campaign Manager + Workflow Health Manager)
- **Smoke no-safe-outputs cluster NEW**: Pi (#39560), Codex (#39561), Antigravity (#39562) all produced no safe outputs today on copilot/add-custom-validation-safe-outputs branch. May indicate PR-branch safe-output precondition issue. Investigate before merging.
- **AIC crisis spreading**: Impact Efficiency Report hit rate limit (#39497). Root fix #39077 Day 9 — escalate urgency.
- **copilot-swe-agent merge rate watch**: 62% in current sample vs 77% historical. Monitor next 5 PRs.
- **Campaign Manager offline 2+ days**: campaign-manager-latest.md absent — coordination risk for campaign-aware agents.
- **pr-code-quality-reviewer compile failure** (#39507): Missing issues:read permission. Quick fix needed.
