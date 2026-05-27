# Shared Alerts — 2026-05-27T13:51Z

## P0 (Critical) 🚨
- **CGO build regression** (#35028): 37% failure on push to main — blocks CI
  - Previously #29669 (90+ days). New instance #35028 from today.
  - **ACTION NEEDED:** Engineering escalation

## P1 (High) 🚨
- **Safe outputs permission blocking** — systemic: affects PR Sous Chef (#35073, #35067, #35052), Contribution Check, Agentic Maintenance
  - `safeoutputs writes are permission-blocked in this environment` error
  - Root cause unknown — platform-level incident, not per-workflow issue
- **Copilot/Codex CLI failures** — Daily News (5+ days), Daily Issues Report Generator, Daily Fact About gh-aw
  - 100% failure rate, `Execute Copilot CLI` / `Execute Codex CLI` steps failing
  - New issue filed May 27 by Agent Performance Analyzer
- **Repo-memory patch size exceeded** (#35105): Daily Community Attribution Updater
  - Fix: increase `max-patch-size` in workflow frontmatter

## P2 (Watch) ⚠️
- **Step Name Alignment cache_memory miss** (#35135): 67% failure
- **Agentic Maintenance intermittent failures**: 26% (6/23) — no issue yet, monitor
- **Super Linter Report**: 100% fail (2/2) — infrastructure issue
- **Daily AW Cross-Repo Compile Check**: 100% fail — cache-memory git repo setup failure
- **87 inactive workflows**: No recent runs, review for deprecation
- **failure-reporters**: 20 issues/day, 60% duplicate rate — dedup issue filed prior run, validate

## Resolved (Do Not Re-File) ✅
- PR-review cluster #31724: CLOSED ✅
- May 14 mass failure (#32045-#32119): RESOLVED ✅
- Sergo #32755: CLOSED ✅
- Step Name #32754: CLOSED ✅ (recurred as #32955, now #35135)
- CGO #29669: REPLACED by #35028

## Coordination Notes
- Health Score: 90/100 (↑20 from 70/100 — biggest single-week improvement)
- Compilation: 100% (236/236)
- copilot-swe-agent: 67% merge rate (↑6pp) — strong signal, increase campaign assignments
- Quality plateau: 74/100 for 4 weeks — infrastructure fixes are prerequisite to improvement
- Infrastructure-blocking is dominant fleet pattern (5/9 agent groups): P0 priority to fix
