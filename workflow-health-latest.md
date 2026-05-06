# Workflow Health — 2026-05-06T05:37Z

Score: 61/100 (↓ from 63 — new CI regression on main). 214 workflows. Run: §25418403606

## KEY FINDINGS

### Compilation Status
- 214/214 lock files present ✅ (stable, +1 from 213)
- 0 missing lock files ✅

### P0 Issues (Active, unchanged)
- **Smoke Gemini** (#30175, #29666 OPEN): 100% failure, proxy blocks traffic (31+ days)
- **Smoke CI** (#29666 OPEN): Chronic action_required, EROFS crash
- **Daily Model Inventory Checker** (#30043 OPEN): Copilot CLI silent startup crash
- **APM Unpack Systemic Failure** (#30252 OPEN): apm-default.tar.gz unpack exit code 1

### P1 Issues (Active)
- **CI regression on main** (NEW — no issue #): TestStrictModePermissions/no_permissions_specified failing since ~01:16 UTC May 6. Playwright MCP deprecation error in strict mode. Filed via safeoutputs.
- **config.models unsupported field** (#30307 OPEN): blocks smoke runs
- **Smoke macOS ARM64** (NO ISSUE FILED): 100% failure since Feb 2026

### P1 Issues Resolved This Run ✅
- #30205 Auto-Triage Issues → CLOSED ✅
- #30188 Documentation Unbloat → CLOSED ✅
- #30233 Daily Documentation Healer → CLOSED ✅
- #30069 Step Name Alignment → CLOSED ✅
- #30241 Smoke Claude → CLOSED ✅
- #30244 Smoke Codex → CLOSED ✅

### P2 Issues
- Node.js 20 deprecation deadline Sep 16, 2026
- MCP gateway session timeout (#23153)
- 6+ PR-review agents on same triggers (~100 action_required/day)
- Performance regression in Validation (#30180)

### Actions Taken This Run
- Created CI regression issue (via safeoutputs) for TestStrictModePermissions failure on main
- Added health dashboard comment to #29109
- Updated shared memory

### Trends
- Score: 61/100 (↓-2 from CI regression on main)
- 6 P1 issues resolved ✅ (good recovery)
- New CI integration test regression since ~01:16 UTC May 6
- Gemini still completely broken (31+ days, P0 unresolved)
- macOS ARM64 chronic failure since Feb 2026 — no issue filed (still)
