# Shared Alerts — 2026-05-28T14:05Z

## P0 (Critical) 🚨
- **safe_outputs add_comment validation** (#35351): `Target is "*" but no item_number specified` — affects PR Sous Chef, Contribution Check, Sub-Issue Closer, and others. Ongoing.
  - Root cause: agent omits required target identifier when `target: "*"` is configured
  - **ACTION NEEDED:** Fix safe_outputs target resolution in affected workflows

## P1 (High) 🚨
- **Copilot CLI failures** — Copilot CLI Deep Research Agent (#35388), Daily News, Daily Issues Report Generator
  - 0% success for 5+ consecutive days; platform-level — infra/platform team needed
- **failure-reporters duplication**: 20 issues/day, 60% duplicate rate — dedup gate still unimplemented
- **LintMonster** (#35370, epic #35368): 2417 issue backlog causing resource/timeout failures
  - Recommendation: shard into bounded batches before next run

## P2 (Watch) ⚠️
- **Silent-skip cluster** (NEW — 2026-05-28): Q, Deployment Incident Monitor, CJS, Label Closed PRs
  - 0-33% success with zero failure logs across 6-8 runs each — trigger audit needed
  - New tracking issue filed this run
- **Daily Safe Output Tool Optimizer** (#35316): 115 turns / 14.9M tokens runaway loop
  - Add early-exit guard on rate-limit detection + max turn budget cap
- **Ubuntu Actions Image Analyzer** (#35378): intermittent failures
- **Daily Firewall Logs Collector and Reporter** (#35363): intermittent failures
- **Go Logger Enhancement** (#35377): intermittent failures
- **Documentation Noob Tester** (#35397): intermittent failures

## Resolved / Monitoring ✅
- CGO build failures (#35028): stabilizing — no new failures on push today
- Daily Community Attribution Updater (#35105): resolved
- Step Name Alignment: consistently succeeding
- PR-review cluster #31724: CLOSED
- May 14 mass failure: RESOLVED

## Do Not Re-File ✅
- Copilot/Codex CLI issue: tracked in #35388
- safe_outputs systemic: tracked in #35351
- LintMonster backlog: tracked in #35368 epic
