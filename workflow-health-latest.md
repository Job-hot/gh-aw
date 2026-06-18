# Workflow Health — 2026-06-18T06:10Z

Score: 67/100 (↓1 from 68)
Workflows: 250 | Lock files: 250/250 (100% ✅) | Run: §27740219260

## KEY FINDINGS

### Status (June 18)
- **Compilation:** 250/250 workflows have lock files (100% ✅)
- **Code Simplifier (Day 11, #39199/#39489):** Still 0/5 fail. api-proxy cap + HTTP 429. DO NOT RE-FILE.
- **Daily Model Inventory Checker (Day 9, #39471):** Confirmed failing Jun 18. session.idle 60s. DO NOT RE-FILE.
- **Daily Safe Outputs Git Simulator (Day 10+):** Still failing Jun 18. Branch missing. Issue Monster queued fix. DO NOT RE-FILE.
- **Daily Safe Output Integrator (Day 10, #39477):** Still 0/5 fail. Tool denial. DO NOT RE-FILE.
- **Daily BYOK Ollama Test (Day 10, #39476):** Still failing. transient_bad_request. DO NOT RE-FILE.
- **Tool Denial Cluster (systemic, filed Jun 16):** 7+ workflows. Daily MCP Tool Concurrency Analysis added Jun 17. DO NOT RE-FILE.
- **Smoke Copilot (~75% fail, #38998):** upload_artifact malformed 400 continues. + New auto-filed today: #39994/#39992/#39989/#39988/#39987/#39986. DO NOT RE-FILE.
- **AIC Budget Crisis (Day 11, #39077):** Root fix still pending. DO NOT RE-FILE.

### Recovering/Resolved ✅
- **Daily Documentation Updater (#39775):** RECOVERED — 3 consecutive successes Jun 18.
- **Daily Workflow Updater (#39753):** RECOVERED — 3 consecutive successes.
- **Instructions Janitor (#39757):** RECOVERED — successes Jun 17-18.
- **Glossary Maintainer (#39769):** RECOVERED — success Jun 18.
- **Avenger:** Healthy — 4+ consecutive successes today.
- **AI Moderator:** Running today.

### New Patterns (Jun 18)
- **Smoke cluster broadened:** 6 new auto-filed failure issues across Copilot, AOAI, Codex, Claude, Gemini. Same day = possible systemic issue.
- **Daily Compiler Quality Check (Day 2, #39724 closed):** model_not_supported_error (`gpt-5-mini`). Config fix needed.
- **LintMonster time-of-day pattern:** Success at 00:50, failure at 03:44. Possible race condition.

### Warnings (P2) ⚠️
- Daily Compiler Quality Check: Day 2 config error (gpt-5-mini model unavailable)
- Daily News: Day 5+ engine exit
- LintMonster: Alternating (success midnight, fail ~04:00)
- Daily Cache Strategy Analyzer (#39451): ~50% alternating

### Actions Taken This Run
- 1 comment added to #29109 (health dashboard Jun 18)

## Do Not Re-File
All from previous runs plus new Jun 18:
- #39994/#39992/#39989/#39988/#39987/#39986 — Smoke test auto-filed Jun 18
- All previous: see shared-alerts.md
