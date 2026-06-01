# Shared Alerts — 2026-06-01T06:08Z

## P0 (Critical) 🚨
_None active_

## P1 (High) 🚨
- **Step Name Alignment** (#36062 OPEN, #36187 new run): 100% fail, Claude engine terminating. DO NOT re-file.
- **LintMonster backlog** (#36050 open, #36175, #36173 new today): Ongoing. DO NOT re-file.
- **Failure-reporters duplication**: #35984 (dedup gate unimplemented). DO NOT re-file.
- **CGO CI**: Mixed ~33% fail. PR #35883 pending review. DO NOT re-file.

## P2 (Watch) ⚠️
- **Daily Safe Output Tool Optimizer** (#35316): runaway token usage
- **jsweep JavaScript Unbloater** (#36183): Token budget exhausted June 1
- **Daily Compiler Quality Check** (#36172): Token budget exhausted June 1
- **Daily Firewall Logs Collector** (#36171): Recurring failure (June 1 instance)
- **Daily SPDD Spec Planner** (#36138): Failed May 31
- **Safe Outputs Conformance SEC-005** (#36079): update_activation_comment.cjs allowlist gap
- **CLI Tools Test** (#36076): compile --workflow-name flag undocumented
- **chaos-test**: 5 open PRs (#36120–#36124) with 0 merges — stall

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
- Smoke tests: #35964, #35959, #36018, #36019, #35955, #35954
- LintMonster: #36050, #36051, #36052, #36175, #36173
- Step Name Alignment: #36062, #36187
- Failure-reporters dedup: #35984
- Token budget exhaustion (jsweep, daily-compiler): #36183, #36172
