# Shared Alerts — 2026-05-29T05:54Z

## P0 (Critical) 🚨
- **safe_outputs add_comment validation** (#35351): `Target is "*" but no item_number specified` — affects PR Sous Chef, Contribution Check, Sub-Issue Closer, and others. Ongoing.
  - Root cause: agent omits required target identifier when `target: "*"` is configured
  - **ACTION NEEDED:** Fix safe_outputs target resolution in affected workflows

## P1 (High) 🚨
- **CJS shard 4 CI failure** (NEW 2026-05-29): `push_to_pull_request_branch.test.cjs` assertion failures at lines 1027+1056 — every push to main fails CI. New issue filed this run.
  - Impact: every commit to main has red CI — developer experience severely degraded
- **Step Name Alignment** (NEW 2026-05-29): 80% failure rate since May 20 — new issue filed
- **Copilot CLI failures** — Copilot CLI Deep Research Agent (#35388), Daily News, Daily Issues Report Generator
  - 0% success for 5+ consecutive days; platform-level — infra/platform team needed
- **failure-reporters duplication**: 20 issues/day, 60% duplicate rate — dedup gate still unimplemented
- **LintMonster** (#35370, epic #35368): 2218+ finding backlog causing resource/timeout failures
  - Recommendation: shard into bounded batches before next run

## P2 (Watch) ⚠️
- **Silent-skip cluster** (2026-05-28): Q, Deployment Incident Monitor, CJS, Label Closed PRs
  - 0-33% success with zero failure logs across 6-8 runs each — trigger audit needed
- **Daily Safe Output Tool Optimizer** (#35316): 115 turns / 14.9M tokens runaway loop
  - Add early-exit guard on rate-limit detection + max turn budget cap
- **Ubuntu Actions Image Analyzer** (#35378): intermittent failures
- **Daily Firewall Logs Collector and Reporter** (#35363): intermittent failures
- **Avenger** (#35374/#35532): intermittent failures

## Resolved / Monitoring ✅
- CGO build failures (#35028): stabilizing
- Daily Community Attribution Updater (#35105): resolved
- Step Name Alignment: NEWLY FILED #TBD — was incorrectly marked resolved
- PR-review cluster #31724: CLOSED

## Do Not Re-File ✅
- Copilot/Codex CLI issue: tracked in #35388
- safe_outputs systemic: tracked in #35351
- LintMonster backlog: tracked in #35368 epic
