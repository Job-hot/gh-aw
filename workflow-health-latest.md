# Workflow Health - 2026-04-17T12:10Z

Score: 73/100 (↑2 from 71 Apr 16). 194 workflows. Run: §24564212990

## KEY FINDINGS

### Compilation Status (IMPROVED)
- 194/194 lock files present ✅ (was 191)
- 0 stale lock files ✅ (was 16 - major improvement!)
- +3 new workflows added since Apr 16

### P0 Persistent Failures
- **Daily Fact About gh-aw**: 10/10 consecutive schedule failures, auto-issue #26852 (today)
  - Codex engine failure pattern
  - Latest failure: 2026-04-17T11:27Z

### Resolved Since Last Run
- ✅ Daily Issues Report Generator (#26393): Closed not_planned by pelikhan Apr 17
- ✅ Smoke Gemini (#26351): Closed not_planned by pelikhan Apr 17 (API key invalid)

### P1 Watch
- **Daily Community Attribution**: 5/10 failures (50%), copilot engine crash, #26848 today
  - Engine: copilot. Error: engine terminated unexpectedly during README edit
- **Smoke Claude**: Last schedule run Apr 14 = failure. No recent schedule data. Monitor.
- **Smoke Codex**: Last schedule run Apr 14 = success. Recovered.

### Today's Schedule Failures (Apr 17)
- daily-fact (100% streak, Codex engine)
- daily-community-attribution (50% rate, Copilot engine crash)

## Compilation Status
- 194/194 lock files present ✅
- 0 stale lock files ✅ (major improvement from 16 yesterday)

## Open Issues (workflow-related)
- #26848 Daily Community Attribution (P1) - TODAY
- #26852 Daily Fact About gh-aw (P0-like) - TODAY
- #26239 MCP rate-limit circuit breaker (P2)
- #26458 GitHub MCP get_me 403 errors (P2)
- #26777, #26790 Smoke Claude (P1)
- #26803 Copilot v1.0.27 upgrade (P2)

## Actions This Run
- Created: Dashboard issue Apr 17 (see GitHub)
- No new P0/P1 issues created (daily-fact and community-attribution already auto-tracked)

## Engine/Tool Status
- Copilot v1.0.21 active / v1.0.27 available (#26803 open)
- Claude Code 2.1.109 available
- Gemini: API key invalid (issues closed not_planned)
- Codex: Working (smoke-codex recovered)
