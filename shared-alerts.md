# Shared Alerts — 2026-06-17T06:10Z (updated by Workflow Health Manager)

## P1 (High) 🚨
- **Tool Denial Cluster — SYSTEMIC**: 6+ workflows hitting 5/5 `tool_denials_exceeded`: Daily Safe Output Integrator, Daily Agent of the Day Blog Writer, Breaking Change Checker, Delight, jsweep, Daily Compiler Threat Spec Optimizer. Systemic issue filed Jun 16. DO NOT RE-FILE per-workflow.
- **Daily Model Inventory Checker — Day 8** (#39471): session.idle timeout 60000ms. DO NOT RE-FILE.
- **AIC Budget Crisis — Day 10, Code Simplifier** (#39077, #39199, #39489): api-proxy cap (maxRuns 50/50) + HTTP 429. Root fix STILL PENDING DAY 10. #39479 filed. DO NOT RE-FILE.
- **upload_artifact malformed 400** (#38998): Smoke Copilot ~75-95% failure. DO NOT RE-FILE.
- **Daily Safe Outputs Git Simulator — 9-day streak**: memory/git-simulator branch missing. Issue Monster queued Copilot fix. DO NOT RE-FILE.
- **Smoke Gemini** (#39172): preview TTS model → unknown model error. DO NOT RE-FILE.
- **Daily BYOK Ollama Test — Day 9** (#39476): transient_bad_request every inference call. DO NOT RE-FILE.
- **Daily Safe Output Integrator — Day 9** (#39477): tool denial 5/5 + ECONNREFUSED in fallback. DO NOT RE-FILE.
- **Daily Cache Strategy Analyzer** (#39451): Codex gpt-5-codex-alpha-2025-11-07 returns 404. ~50% alternating success. DO NOT RE-FILE.

## P2 (Watch) ⚠️
- **LintMonster** (#39511 closed Jun 16): Still failing Jun 16 + Jun 17 despite issue closed. Regression noted. Commented on #39511.
- **Smoke Codex/Antigravity** (#39560-39562): no safe outputs, tracked batch. Watch.
- **Smoke Trigger & Smoke Multi Caller** (#38999): 100% startup_failure. DO NOT RE-FILE.
- **Agentic workflows out of sync** (#38993). DO NOT RE-FILE.
- **Failure Investigator pre-fetch gap** (#39037). DO NOT RE-FILE.
- **Performance regression spike** (#38870, #38871, #38872). DO NOT RE-FILE.
- **Dictation Prompt Generator** (#39196, #39200): session.idle 570s timeout. DO NOT RE-FILE.
- **GitHub Remote MCP Auth Test** (#39193, #39505): missing required tool. DO NOT RE-FILE.
- **Daily Compiler Threat Spec Optimizer** (#39343): tool denial cluster (weekly pattern). Watch.
- **pr-code-quality-reviewer** (#39507): compile failure. Watch.

## Resolved ✅
- **PR Sous Chef**: FULLY RECOVERED (9+ consecutive successes Jun 15-17). Close issue recommended.
- **Avenger**: Healthy.
- **AI Moderator**: Recovering (consecutive successes Jun 17).

## Transient (1 failure, watch)
- **Documentation Noob Tester**: 1 failure Jun 17, 4 previous successes. Likely transient. DO NOT FILE.

## Systemic Notes
- **Tool denial cluster**: 6+ workflows — all individually filing issues but root cause is allow-list blocking `git checkout -b` and other shell git ops. Systemic issue filed Jun 16.
- **session.idle timeout cluster**: Model Inventory (60s) + Dictation Prompt Generator (570s) — platform-level Copilot SDK regression likely.
- **AIC cluster Day 10**: Code Simplifier still blocked. #39479 addresses non-retryable 429 handling.
- **Codex model 404**: Daily Cache Strategy Analyzer uses deprecated model path.
- **Compilation**: 249/249 (100%) — all lock files present.

## Do Not Re-File (All Active Issues)
- Tool denial cluster: systemic issue filed Jun 16
- PR Sous Chef: RESOLVED
- Daily Model Inventory Checker Day 8: #39471
- #39199/#39489: Code Simplifier Day 10 (P1)
- #38998: upload_artifact malformed 400 (P1)
- #38999: Smoke Trigger startup_failure (P2)
- #38993: Agentic workflows out of sync
- #39024: Daily Safe Outputs Git Simulator (expired, Issue Monster fix queued)
- #39037: Failure Investigator pre-fetch blind spot
- #38870/#38871/#38872: Performance regression cluster
- #39077: AIC Budget Crisis Day 10 (escalating)
- #39172: Smoke Gemini preview model (P1)
- #39196/#39200: Dictation Prompt Generator (P2)
- #39193/#39505: GitHub Remote MCP Auth missing tool
- #39476: Daily BYOK Ollama Test Day 9 (P1)
- #39477: Daily Safe Output Integrator Day 9 (P1)
- #39452: AI Moderator pagination timeout (recovering)
- #39451: Daily Cache Strategy Analyzer Codex 404 (P1)
- #39343: Daily Compiler Threat Spec Optimizer tool denial (P2)
- #39560-39562: Smoke Pi/Codex/Antigravity no safe outputs
- #39507: pr-code-quality-reviewer compile failure
- #39511: LintMonster (closed, regression commented today)
