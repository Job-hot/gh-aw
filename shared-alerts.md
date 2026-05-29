# Shared Alerts — 2026-05-29T13:46Z

## P0 (Critical) 🚨
- **safe_outputs add_comment validation** (#35351): `Target is "*" but no item_number specified` — affects PR Sous Chef, Contribution Check, Sub-Issue Closer, and others. Ongoing.
  - Root cause: agent omits required target identifier when `target: "*"` is configured
  - **ACTION NEEDED:** Fix safe_outputs target resolution in affected workflows

## P1 (High) 🚨
- **Agentic Commands regression** (NEW 2026-05-29): 80%→25% success week-over-week — was "most stable" workflow. Likely correlated with CJS shard 4 CI failure or Step Name Alignment P1.
- **Content Moderation regression** (NEW 2026-05-29): 75%→22% success week-over-week. Root cause unclear — review last 3 failure logs.
- **CJS shard 4 CI failure** (filed 2026-05-29): `push_to_pull_request_branch.test.cjs` lines 1027+1056 fail on EVERY push to main — CI permanently broken.
- **Step Name Alignment** (filed 2026-05-29): 80% failure rate (8/10 recent runs), chronic since May 20.
- **Copilot CLI failures** (#35388): 0% success for 5+ consecutive days — platform-level, infra team needed.
- **failure-reporters duplication**: 20 issues/day, 60% duplicate rate — dedup gate unimplemented (3rd consecutive escalation, needs owner assignment).
- **LintMonster** (#35370, epic #35368): 2218+ finding backlog causing resource/timeout failures.

## P2 (Watch) ⚠️
- **Silent-skip cluster**: Q, Deployment Incident Monitor, CJS, Label Closed PRs — 0-33% success, zero failure logs
- **Daily Safe Output Tool Optimizer** (#35316): 115 turns / 14.9M tokens runaway loop
- **Ubuntu Actions Image Analyzer** (#35378): intermittent failures
- **Daily Firewall Logs Collector and Reporter** (#35363): intermittent failures
- **Avenger** (#35374/#35532): intermittent failures
- **Campaign Manager**: latest memory file absent — may have stopped running

## Resolved / Monitoring ✅
- CGO build failures (#35028): stabilizing
- Daily Community Attribution Updater (#35105): resolved
- PR-review cluster #31724: CLOSED

## Do Not Re-File ✅
- Copilot/Codex CLI issue: tracked in #35388
- safe_outputs systemic: tracked in #35351
- LintMonster backlog: tracked in #35368 epic
- Step Name Alignment: filed 2026-05-29 by Workflow Health Manager
- CJS shard 4 CI: filed 2026-05-29 by Workflow Health Manager
