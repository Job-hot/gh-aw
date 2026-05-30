# Shared Alerts — 2026-05-30T05:41Z

## P0 (Critical) 🚨
- **safe_outputs add_comment validation** (#35351): `Target is "*" but no item_number specified` — affects PR Sous Chef, Contribution Check, Sub-Issue Closer, and others. Ongoing.
  - Root cause: agent omits required target identifier when `target: "*"` is configured
  - **ACTION NEEDED:** Fix safe_outputs target resolution in affected workflows
- **Copilot CLI engine systemic failure** (#35388): All workflows using `copilot` engine fail on "Execute GitHub Copilot CLI" step. Affects jsweep, Documentation Noob Tester, Copilot CLI Deep Research Agent, and many others. Ongoing since May 28.

## P1 (High) 🚨
- **Smoke tests systemic**: smoke-trigger, smoke-water, smoke-multi-caller all 100% failing. Issues: #35829, #35832, #35856, #35864, #35866, and more. Systemic smoke infrastructure failure.
- **CGO CI** (2026-05-30): Unit tests + custom linters failing on CGO CI workflow (run 26675666376). Build may be broken.
- **CJS shard 4 CI failure** (filed 2026-05-29): `push_to_pull_request_branch.test.cjs` lines 1027+1056 fail on EVERY push to main — CI permanently broken.
- **Step Name Alignment** (filed 2026-05-29): 80% failure rate (8/10 recent runs), chronic since May 20.
- **failure-reporters duplication**: 20 issues/day, 60% duplicate rate — dedup gate unimplemented (ongoing).
- **LintMonster** (#35370, epic #35368): 2218+ finding backlog causing resource/timeout failures.

## P2 (Watch) ⚠️
- **Silent-skip cluster**: Q, Deployment Incident Monitor, CJS, Label Closed PRs — 0-33% success, zero failure logs
- **Daily Safe Output Tool Optimizer** (#35316): 115 turns / 14.9M tokens runaway loop
- **Ubuntu Actions Image Analyzer** (#35378): intermittent failures
- **Daily Firewall Logs Collector and Reporter** (#35363): intermittent failures
- **Avenger** (#35374/#35532): intermittent failures

## Resolved / Monitoring ✅
- CGO build failures (#35028): stabilizing (note: new CI test failures now present)
- Daily Community Attribution Updater (#35105): resolved
- PR-review cluster #31724: CLOSED

## Do Not Re-File ✅
- Copilot/Codex CLI issue: tracked in #35388
- safe_outputs systemic: tracked in #35351
- LintMonster backlog: tracked in #35368 epic
- Step Name Alignment: filed 2026-05-29 by Workflow Health Manager
- CJS shard 4 CI: filed 2026-05-29 by Workflow Health Manager
- Smoke test failures: tracked in multiple open smoke issues (see P1 above)
