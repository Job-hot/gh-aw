# Workflow Health — 2026-05-07T05:38Z

Score: 61/100 (→ stable). 217 workflows (+3). Run: §25477981073

## KEY FINDINGS

### Compilation Status
- 217/217 lock files present ✅ (+3 from yesterday's 214)
- 0 missing lock files ✅

### P0 Issues (Active, unchanged)
- **Smoke Gemini** (#30175, #29666 OPEN): 100% failure, proxy blocks traffic (32+ days)
- **Smoke CI** (#29666 OPEN): Chronic action_required, EROFS crash
- **Daily Model Inventory Checker** (#30043 OPEN): Copilot CLI silent startup crash
- **APM Unpack Systemic Failure** (#30252 OPEN): apm-default.tar.gz unpack exit code 1

### P1 Issues (Active)
- **Smoke macOS ARM64**: 100% failure since Feb 20 2026 (**76 days**) — **ISSUE FILED THIS RUN** ✅
- **CI regression** (TestStrictModePermissions): Filed 2026-05-06; status unknown this run (API unavailable)
- **config.models unsupported field** (#30307 OPEN): blocks smoke runs
- **MCP gateway session timeout** (#23153 OPEN): Long-running workflows at risk
- **Performance Regression in Validation** (#30180): 82.1% slower

### P2 Issues
- Node.js 20 deprecation deadline Sep 16, 2026
- 9 PR-review agents on same triggers (~100 action_required/day)

### Actions Taken This Run
- Added health dashboard comment to #29109
- Filed P1 issue for Smoke macOS ARM64 (76+ days overdue)
- Updated shared memory

### Trends
- Score: 61/100 (→ stable, no new regressions)
- 217 workflows (+3 new), all compiled
- P0s persist (Gemini 32+ days, CI CGO chronic)
- macOS ARM64 issue finally filed after 76 days
- Network-restricted run: GitHub API unavailable, analysis from pre-computed data
