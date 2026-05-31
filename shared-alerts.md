# Shared Alerts — 2026-05-31T13:05Z

## P0 (Critical) 🚨
_None active — both prior P0s resolved:_
- ~~**safe_outputs add_comment validation** (#35351)~~: **CLOSED** ✅ (PR #35901 fixed it)
- ~~**Copilot CLI engine systemic failure** (#35388)~~: **CLOSED** ✅

## P1 (High) 🚨
- **Smoke tests systemic**: All variants 100% failing on main. Issues: #35964, #35959, #36018, #36019, #35955, #35954. DO NOT re-file.
- **LintMonster backlog** (#35368 epic): 2218+ findings, batch issues #36050, #36051, #36052. DO NOT re-file.
- **Failure-reporters duplication**: #35984 (60% dup rate, dedup gate unimplemented). DO NOT re-file.
- **Step Name Alignment**: #36062 (80% failure rate). DO NOT re-file.
- **CGO CI**: Mixed 50% fail — PR #35883 pending review.
- **CJS shard 4 CI**: Filed 2026-05-29. DO NOT re-file.

## P2 (Watch) ⚠️
- **Daily Safe Output Tool Optimizer** (#35316): 115 turns / 14.9M tokens runaway
- **Ubuntu Actions Image Analyzer** (#35378): intermittent
- **Daily Firewall Logs Collector** (#36047): filed May 31 — assess severity
- **Design Decision Gate** (#36044): filed May 31 — assess severity
- **Test Quality Sentinel** (#36043): filed May 31 — assess severity
- **Daily Syntax Error Check** (#36089): filed May 31 — new failure
- **PR Sous Chef** (#35951): filed May 30
- **chaos-test**: 5 open PRs (#36120–#36124) with 0 merges — stall

## Resolved ✅
- **safe_outputs add_comment** (#35351): CLOSED
- **Copilot CLI engine** (#35388): CLOSED
- CGO build failures (#35028): stabilizing
- Daily Community Attribution Updater (#35105): resolved

## Do Not Re-File ✅
- Smoke tests: #35964, #35959, #36018, #36019, #35955, #35954
- LintMonster: #35368 epic
- Step Name Alignment: #36062
- Failure-reporters dedup: #35984
- CJS shard 4 CI: filed 2026-05-29
