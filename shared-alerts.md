# Shared Alerts — 2026-06-14T06:05Z

## P1 (High) 🚨
- **AIC Budget Crisis — Day 6, 6-agent cluster** (#39077, OPEN): Code Simplifier + 5 other agents blocked. Root fix (raise max-ai-credits to 2000, add max-turns:30) STILL PENDING DAY 6. ESCALATE TO TEAM. DO NOT RE-FILE.
- **Code Simplifier — Day 6 failure** (#39179, OPEN, Avenger-filed Jun 14): HTTP 429 rate-limited. DO NOT RE-FILE.
- **upload_artifact malformed 400** (#38998, OPEN): Smoke Copilot ~95% failure 2+ days. DO NOT RE-FILE.
- **Daily Safe Outputs Git Simulator — 6-day streak** (#39024, OPEN, P1): memory/git-simulator branch missing. Issue Monster queued Copilot fix. DO NOT RE-FILE.
- **Daily Model Inventory Checker — 5-day streak** (#39170, OPEN, P1): copilot-sdk session.idle timeout 60000ms. DO NOT RE-FILE.

## P2 (Watch) ⚠️
- **Smoke Trigger & Smoke Multi Caller** (#38999, OPEN): 100% startup_failure reusable-workflow callers. DO NOT RE-FILE.
- **Agentic workflows out of sync** (#38993, OPEN): Lock files need recompilation. DO NOT RE-FILE.
- **Failure Investigator pre-fetch gap** (#39037, OPEN): Returns empty failed_run_ids despite in-window failures. DO NOT RE-FILE.
- **Performance regression spike** (#38870, #38871, #38872, OPEN): Compile ops 165-269% slower. DO NOT RE-FILE.

## Resolved (Jun 13→14) ✅
- None confirmed resolved today.

## Systemic Notes
- **AIC cluster Day 6**: Now Day 6. Code Simplifier alone drove 4.2K AIC Jun 12 + daily 429 since. Guardrail fix MUST ship.
- **upload_artifact P1**: 2+ day blocker. `safe_outputs` job needs non-fatal upload handling.
- **Git Simulator Day 6**: Issue Monster queued for Copilot fix. If no PR in 24h, escalate for manual branch creation.
- **Model Inventory Checker Day 5+**: Session timeout pattern — may indicate Copilot SDK regression.

## Do Not Re-File (All Active Issues)
- #39179: Code Simplifier Day 6 (P1)
- #38998: upload_artifact malformed 400 (P1)
- #38999: Smoke Trigger startup_failure (P2)
- #38993: Agentic workflows out of sync
- #39024: Daily Safe Outputs Git Simulator Day 6 (P1)
- #39170: Daily Model Inventory Checker 5-day (P1)
- #39037: Failure Investigator pre-fetch blind spot
- #38870/#38871/#38872: Performance regression cluster
- #39077: AIC Budget Crisis Day 6 (escalating)
