# Shared Alerts — 2026-06-06T05:56Z

## P0 (Critical) 🚨
_None active_

## P1 (High) 🚨
- **CJS typecheck** (#36410 OPEN): 100% failure since Jun 1. DO NOT re-file.
- **CGO unit tests** (#35028 OPEN): 100% failure since May 27. DO NOT re-file.
- **Daily Firewall Logs Collector**: 8th+ consecutive failure, token budget exhaustion. Filed Jun 4. DO NOT re-file.
- **Daily BYOK Ollama Test** (#37211 OPEN): Auth failures continue. Filed Jun 5. DO NOT re-file.
- **Daily Documentation Healer** (re-opened Jun 6): #37039 closed without full fix — still failing. NEW consolidated issue filed today. DO NOT re-file.
- **Daily Model Inventory Checker** (re-opened Jun 6): BYOK auth still missing. Covered by today's issue. DO NOT re-file.

## P2 (Watch) ⚠️
- **Daily Sentrux Report** (Jun 6): 3rd+ consecutive failure, #37019 CLOSED → covered by today's consolidated issue. DO NOT re-file.
- **Safe Output Health Monitor** (#37264 OPEN): Token budget exhaustion — monitoring coverage degraded. DO NOT re-file.
- **AI Moderator**: Still blocked (0% success rate, ongoing). Old issue exists. Monitor.
- **PR Sous Chef** (#37216 OPEN): 2+ failures Jun 5-6. DO NOT re-file.
- **Code Simplifier**: Recurring failure. #36829 old. Monitor for escalation.
- **jsweep + Daily Compiler Quality Check**: Token budget. Monitor.
- **Smoke tests on various PR branches**: CJS/CGO cascade. Covered by PR-level issues.

## Resolved ✅ (since Jun 5)
_None resolved today_

## Systemic Notes
- **Model/auth config cluster**: Documentation Healer (effort param) + Model Inventory (BYOK auth) + BYOK Ollama (auth) + Firewall Logs (token budget) = 4 workflows with bootstrap/config failures. Previous fix in #37039 insufficient.
- **Token budget exhaustion cluster**: Safe Output Health Monitor + jsweep + Daily Compiler + Firewall Logs = 4 workflows hitting token limits. Systemic pressure increasing.
- **CI blockage**: CJS + CGO both 100% failing — affects ALL PR validation.
- **Health score trending down**: 82 → 81 → 78 → 74 (4-day drop). No resolutions, new failures accumulating.
- **Monitoring gap**: Safe Output Health Monitor itself failing → observability degraded.
