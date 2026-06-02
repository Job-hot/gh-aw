# Shared Alerts — 2026-06-02T06:03Z

## P0 (Critical) 🚨
_None active_

## P1 (High) 🚨
- **CJS typecheck broken** (NEW - issue filed 2026-06-02): 100% fail on main, triggered by PR #36358. DO NOT re-file.
- **CGO unit tests** (#35028 OPEN): Escalated to 100% failing today (20 runs). Auto-notifier already tracking. DO NOT re-file.
- **Step Name Alignment** (#36062 OPEN): Passed today — possibly resolved. Monitor.
- **LintMonster backlog** (#36050 open): Intermittent failures. DO NOT re-file.
- **Failure-reporters duplication**: #35984 (dedup gate unimplemented). DO NOT re-file.

## P2 (Watch) ⚠️
- **Token budget exhaustion (systemic)**: jsweep (#36183) + daily-compiler (#36172) — watch for escalation
- **chaos-test PR stall**: 10+ open PRs (#36120–#36124, #36251–#36256), 0 merges — worsening
- **Daily Safe Output Tool Optimizer** (#35316): runaway token usage
- **Daily Firewall Logs Collector** (#36171): Recurring failure
- **Daily SPDD Spec Planner** (#36138): Failed May 31
- **Safe Outputs Conformance SEC-005** (#36079): update_activation_comment.cjs allowlist gap
- **CLI Tools Test** (#36076): compile --workflow-name flag undocumented

## Resolved ✅ (since May 30)
- **Smoke tests**: #35959, #36018, #36019, #35955, #35954 all CLOSED
- **Daily Syntax Error Check** (#36089): CLOSED
- **Lint batches** #36051, #36052: CLOSED
- **Daily Firewall Logs Collector** #36047: CLOSED (recurred as #36171)
- **Design Decision Gate** #36044: CLOSED
- **Test Quality Sentinel** #36043: CLOSED
- **safe_outputs add_comment** (#35351): CLOSED
- **Copilot CLI engine** (#35388): CLOSED

## Do Not Re-File
- CJS typecheck: filed 2026-06-02
- CGO unit tests: #35028
- Smoke tests: #35964, #35959, #36018, #36019, #35955, #35954
- LintMonster: #36050, #36051, #36052, #36175, #36173
- Step Name Alignment: #36062, #36187
- Failure-reporters dedup: #35984
- Token budget exhaustion (jsweep, daily-compiler): #36183, #36172
