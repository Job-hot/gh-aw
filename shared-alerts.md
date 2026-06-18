# Shared Alerts — 2026-06-17T14:00Z (updated by Agent Performance Analyzer)

## P1 (High) 🚨
- **Tool Denial Cluster — SYSTEMIC**: NOW 7+ workflows. Daily MCP Tool Concurrency Analysis (#39763) is new member Jun 17. Systemic issue filed Jun 16. DO NOT RE-FILE per-workflow.
- **Daily Model Inventory Checker — Day 8** (#39471): session.idle timeout 60000ms. DO NOT RE-FILE.
- **AIC Budget Crisis — Day 10, Code Simplifier** (#39077, #39199, #39489, #39729): api-proxy cap (maxRuns 50/50) + HTTP 429. Root fix STILL PENDING DAY 10. DO NOT RE-FILE.
- **upload_artifact malformed 400** (#38998): Smoke Copilot ~75-95% failure. DO NOT RE-FILE.
- **Daily Safe Outputs Git Simulator — 9-day streak**: memory/git-simulator branch missing. Issue Monster queued fix. DO NOT RE-FILE.
- **Smoke Gemini** (#39172): preview TTS model → unknown model error. DO NOT RE-FILE.
- **Daily BYOK Ollama Test — Day 9** (#39476): transient_bad_request. DO NOT RE-FILE.
- **Daily Safe Output Integrator — Day 9** (#39477): tool denial 5/5 + ECONNREFUSED in fallback. DO NOT RE-FILE.
- **Daily Cache Strategy Analyzer** (#39451): Codex model 404. ~50% alternating. DO NOT RE-FILE.
- **Incomplete Result Cluster — NEW Jun 17**: Package Spec Enforcer (#39787), Package Spec Extractor (#39762), Daily Docs Updater (#39775), Daily Workflow Updater (#39753). Systemic issue filed #aw_inc_result. DO NOT RE-FILE individually.

## P2 (Watch) ⚠️
- **LintMonster** (#39511 closed Jun 16): Still failing Jun 16 + Jun 17 despite issue closed. Commented on #39511.
- **Smoke Codex/Antigravity** (#39560-39562): no safe outputs, tracked batch. Watch.
- **Smoke Trigger & Smoke Multi Caller** (#38999): 100% startup_failure. DO NOT RE-FILE.
- **Agentic workflows out of sync** (#38993). DO NOT RE-FILE.
- **Failure Investigator pre-fetch gap** (#39037). DO NOT RE-FILE.
- **Performance regression spike** (#38870, #38871, #38872). DO NOT RE-FILE.
- **Dictation Prompt Generator** (#39196, #39200): session.idle 570s timeout. DO NOT RE-FILE.
- **GitHub Remote MCP Auth Test** (#39193, #39505): missing required tool. DO NOT RE-FILE.
- **Daily Compiler Threat Spec Optimizer** (#39343): tool denial cluster (weekly pattern). Watch.
- **pr-code-quality-reviewer** (#39507): compile failure. Watch.
- **New individual failures Jun 17** (watch, may be transient): Test Quality Sentinel (#39782), Matt Pocock Skills Reviewer (#39781), Glossary Maintainer (#39769), Daily News (#39758), Daily Compiler Quality Check (#39724), Metrics Collector Infrastructure Agent (#39727), Design Decision Gate (#39779/#39776, 2 issues for same workflow), Instructions Janitor (#39757).
- **copilot-swe-agent merge rate declining**: 67% (30/45 PRs) ↓ from 77% — watch trend.
- **Metrics Collector failing**: #39727 — affects shared metrics infrastructure. Watch.

## Resolved ✅
- **PR Sous Chef**: FULLY RECOVERED (9+ consecutive successes Jun 15-17). Close issue recommended.
- **Avenger**: Healthy.
- **AI Moderator**: Recovering (consecutive successes Jun 17).

## Systemic Notes
- **Tool denial cluster**: NOW 7+ workflows — systemic issue filed Jun 16. Check allow-list for shell git ops.
- **AIC cluster Day 10**: Code Simplifier still blocked. Root fix #39479 still pending.
- **Incomplete Result Cluster**: 4 workflows Jun 17 — new pattern. May indicate token budget exhaustion or platform issue. Track daily.
- **Compilation**: 249/249 (100%) — all lock files present.
- **Ecosystem health**: 66/100 (↓2 from 68). Quality: 55/100 (↓2 from 57). Effectiveness: 53/100 (↓2 from 55).

## Do Not Re-File (Full List)
Tool denial cluster, PR Sous Chef (RESOLVED), Daily Model Inventory #39471, Code Simplifier #39199/#39489/#39729, upload_artifact #38998, Smoke Trigger #38999, Agentic workflows out of sync #38993, Git Simulator #39024, Failure Investigator #39037, Perf regression #38870/#38871/#38872, AIC Budget Crisis #39077, Smoke Gemini #39172, Dictation #39196/#39200, Remote MCP Auth #39193/#39505, BYOK Ollama #39476, Safe Output Integrator #39477, AI Moderator #39452 (recovering), Cache Strategy #39451, Compiler Threat #39343, Smoke Pi/Codex/Antigravity #39560-#39562, pr-code-quality-reviewer #39507, LintMonster #39511 (commented), Incomplete Result cluster #39787/#39762/#39775/#39753, Test Quality Sentinel #39782, Matt Pocock #39781, Design Decision Gate #39779/#39776, Glossary Maintainer #39769, Daily News #39758, Daily Compiler Quality #39724, Metrics Collector #39727, Instructions Janitor #39757

---
## Updated 2026-06-18T06:10Z (Workflow Health Manager)

### P1 Changes (Jun 18)
- Code Simplifier: Day 11 (unchanged)
- Daily Safe Outputs Git Simulator: Day 10+ confirmed Jun 18 failure
- Daily Model Inventory Checker: Day 9 confirmed Jun 18 failure
- Daily Safe Output Integrator: Day 10 (unchanged)
- Daily BYOK Ollama Test: Day 10 (unchanged)
- Tool Denial Cluster: 7+ workflows (unchanged)

### New (Jun 18)
- **Smoke test cluster broadened:** 6 new auto-filed issues: #39994/#39992/#39989/#39988/#39987/#39986 across Copilot, AOAI, Codex, Claude, Gemini — possible systemic smoke test regression.
- **Daily Compiler Quality Check (Day 2):** model_not_supported_error. Config fix: update gpt-5-mini model name.

### Recoveries (Jun 18)
- Daily Docs Updater (#39775): RECOVERED ✅
- Daily Workflow Updater (#39753): RECOVERED ✅
- Instructions Janitor (#39757): RECOVERED ✅
- Glossary Maintainer (#39769): RECOVERED ✅

### Do Not Re-File Additions (Jun 18)
Smoke test auto-filed: #39994, #39992, #39989, #39988, #39987, #39986
