# Shared Alerts — 2026-06-13T05:55Z

## P1 (High) 🚨
- **Code Simplifier — 5-day failure streak** (#38793, OPEN, Day 5): HTTP 429 rate-limited Jun 13 (316.9 AIC). Yesterday's runaway (4,219 AIC) exhausted provider quota. Fix: max-turns:30, bash allowlist, max-ai-credits:1500. DO NOT RE-FILE.
- **upload_artifact malformed 400** (#38998, OPEN): Smoke Copilot ~95% failure 2+ days. `safe_outputs` job fails on non-fatal artifact upload. DO NOT RE-FILE.
- **Daily Safe Outputs Git Simulator — 5-day streak** (#aw_gitsim13, NEW Jun 13): memory/git-simulator branch missing. Issue filed this run. DO NOT RE-FILE.

## P2 (Watch) ⚠️
- **#38794** Duplicate of #38793 (Code Simplifier) — recommend closing.
- **Smoke Trigger & Smoke Multi Caller** (#38999, OPEN): 100% startup_failure for reusable-workflow callers. DO NOT RE-FILE.
- **Agentic workflows out of sync** (#38993, OPEN): Lock files need recompilation. DO NOT RE-FILE.
- **Copilot CLI Deep Research** (#39022, Jun 13): Exceeded tool denial limit. DO NOT RE-FILE.

## Resolved (Jun 12→13) ✅
- #38758 Failure cascade Jun 12 — CLOSED
- #38767 Failure Investigator blind spot — CLOSED
- #38379 Daily News Node.js chroot — CLOSED
- #38809 Code Simplifier runaway fix — CLOSED

## Systemic Notes
- **Health score trend:** 68→83→87→85→83→75→75→76 (slight recovery after cascade closure)
- **Code Simplifier Day 5**: Mode shift from runaway → rate-limit shows Jun 12 spike has downstream effects on provider quota. ROOT FIX STILL PENDING.
- **upload_artifact P1**: `safe_outputs` job design flaw — one failing optional safe-output should not fail the whole job.
- **Git Simulator temp-ID gap**: `#aw_gitsim10` was never created as a real issue in prior runs. Agents must verify temp-ID resolution.
- **Duplicate creation**: Ongoing #38793/#38794 pattern. Agents must search for existing issues before creating.

## Do Not Re-File (Active Issues)
- #38793: Code Simplifier 5-day tracker (P1)
- #38794: Code Simplifier duplicate (should be closed, not re-filed)
- #38998: upload_artifact malformed 400 (P1)
- #38999: Smoke Trigger startup_failure (P2)
- #38993: Agentic workflows out of sync
- #39022: Copilot CLI Deep Research tool denial
- #aw_gitsim13: Daily Safe Outputs Git Simulator (NEW, P1)
