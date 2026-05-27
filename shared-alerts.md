# Shared Alerts — 2026-05-27T05:54Z

## P0 (Critical) 🚨
- **CGO build regression** (#35028): 37% failure on push to main — blocks CI
  - Previously #29669 (90+ days). New instance #35028 from today.
  - **ACTION NEEDED:** Engineering escalation

## P1 (High) 🚨
- **Safe outputs permission blocking** — systemic: affects PR Sous Chef (#35073, #35067, #35052) and Contribution Check
  - `safeoutputs writes are permission-blocked in this environment` error
  - Root cause unknown — may be repo permissions/token scope
- **Copilot/Codex CLI failures** — Daily News, Daily Issues Report Generator, Daily Fact About gh-aw
  - 100% failure rate (2/2 runs), `Execute Copilot CLI` / `Execute Codex CLI` steps failing
  - No open issue yet — **FILE IF PERSISTS**
- **Repo-memory patch size exceeded** (#35105): Daily Community Attribution Updater
  - Fix: increase `max-patch-size` in workflow frontmatter

## P2 (Watch) ⚠️
- **Step Name Alignment cache_memory miss** (#35135): 67% failure, opened today
- **Agentic Maintenance intermittent failures**: 26% (6/23) — no issue yet, monitor
- **Super Linter Report**: 100% fail (2/2) — infrastructure issue, `Run super-linter` step
- **Daily AW Cross-Repo Compile Check**: 100% fail — cache-memory git repo setup failure
- **87 inactive workflows**: No recent runs, review for deprecation

## Resolved (Do Not Re-File) ✅
- PR-review cluster #31724: CLOSED ✅
- May 14 mass failure (#32045-#32119): RESOLVED ✅
- Sergo #32755: CLOSED ✅
- Step Name #32754: CLOSED ✅ (recurred as #32955, now #35135)
- CGO #29669: REPLACED by #35028

## Coordination Notes
- Health Score: 90/100 (stable May 27)
- Compilation: 100% (236/236)
- 500 scheduled runs analyzed
