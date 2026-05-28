# Shared Alerts — 2026-05-28T05:55Z

## P0 (Critical) 🚨
- **safe_outputs add_comment validation** (#35351): `Target is "*" but no item_number specified` — affects PR Sous Chef, Contribution Check, and others. Ongoing.
  - Root cause: agent omits required target identifier when `target: "*"` is configured
  - **ACTION NEEDED:** Fix safe_outputs target resolution in affected workflows

## P1 (High) 🚨
- **Copilot CLI failures** — Copilot CLI Deep Research Agent (#35388), recurring pattern
  - `Execute Copilot CLI` step failing intermittently; may be platform-level
- **LintMonster** (#35370): Failed today; carrying massive backlog (2417 issues epic #35368)
  - May be resource/timeout related
- **Daily Safe Output Tool Optimizer** (#35316): Claude rate-limit exhaustion (115 turns, 14.9M tokens)

## P2 (Watch) ⚠️
- **Ubuntu Actions Image Analyzer** (#35378): Failed today — investigate root cause
- **Daily Firewall Logs Collector and Reporter** (#35363): Failed today — investigate  
- **Go Logger Enhancement** (#35377): Failed today — investigate
- **Documentation Noob Tester** (#35397): Failed today

## Resolved / Monitoring ✅
- CGO build failures: no new failures today — appears stable
- PR Sous Chef: #35353 auto-filed, likely safe_outputs issue (see P0 above)
- Daily Community Attribution Updater (#35105): Resolved (success today)
- Step Name Alignment: Success today ✅

## Do Not Re-File ✅
- PR-review cluster #31724: CLOSED ✅
- May 14 mass failure: RESOLVED ✅
