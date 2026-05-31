# Shared Alerts — 2026-05-31T05:54Z

## P0 (Critical) 🚨
_None active — both prior P0s resolved:_
- ~~**safe_outputs add_comment validation** (#35351)~~: **CLOSED** ✅ (PR #35901 fixed it)
- ~~**Copilot CLI engine systemic failure** (#35388)~~: **CLOSED** ✅

## P1 (High) 🚨
- **Smoke tests systemic**: All variants 100% failing on main. Issues: #35964, #35959, #36018, #36019, #35955, #35954. Worsening — 6+ issues open. DO NOT re-file.
- **LintMonster backlog** (#35368 epic): 2218+ findings, 3 new batch issues today (#36050, #36051, #36052). Resource/timeout failures ongoing.
- **Failure-reporters duplication**: #35984 (latest: `add_comment` with no issue_number). ~60% duplicate rate. Dedup gate unimplemented. DO NOT re-file.
- **Step Name Alignment**: #36062 (new May 31). 80% failure rate. DO NOT re-file.
- **CGO CI**: Mixed (50% fail) — action_required runs suggest approval gate / dependency issue. PR #35883 pending.
- **CJS shard 4 CI**: Filed 2026-05-29. Ongoing. DO NOT re-file.

## P2 (Watch) ⚠️
- **Daily Safe Output Tool Optimizer** (#35316): 115 turns / 14.9M tokens runaway loop
- **Ubuntu Actions Image Analyzer** (#35378): intermittent failures
- **Daily Firewall Logs Collector** (#36047): filed May 31
- **Design Decision Gate** (#36044): filed May 31
- **Test Quality Sentinel** (#36043): filed May 31
- **PR Sous Chef** (#35951): filed May 30
- **chaos-test**: 5 open PRs with 0 merges — batch-creation stall

## Resolved ✅
- **safe_outputs add_comment** (#35351): CLOSED
- **Copilot CLI engine** (#35388): CLOSED
- CGO build failures (#35028): stabilizing
- Daily Community Attribution Updater (#35105): resolved

## Do Not Re-File ✅
- Smoke test failures: tracked in #35964, #35959, #36018, #36019, #35955, #35954
- LintMonster: tracked in #35368 epic
- Step Name Alignment: #36062 (just filed)
- Failure-reporters dedup: #35984
- CJS shard 4 CI: filed 2026-05-29
