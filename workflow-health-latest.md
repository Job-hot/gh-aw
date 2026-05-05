# Workflow Health — 2026-05-05T05:32Z

Score: 63/100 (↓ from 65 — new APM systemic failure). 213 workflows. Run: §25359456500

## KEY FINDINGS

### Compilation Status
- 213/213 lock files present ✅ (stable)
- 0 missing lock files ✅

### P0 Issues (Active)
- **Smoke Gemini** (#30175, #29666, #30242 OPEN): 100% failure, proxy blocks traffic (chronic 30+ days)
- **Smoke CI** (#29666 OPEN): 100% action_required, EROFS crash (chronic)
- **Daily Model Inventory Checker** (#30043 OPEN): Copilot CLI silent startup crash
- **APM Unpack Systemic Failure** (#30252 OPEN): NEW — apm-default.tar.gz unpack exit code 1, affects Smoke Claude and multiple workflows

### P1 Issues (Active)
- **Smoke Claude** (#30241 OPEN): failed 00:38 UTC May 5 due to APM unpack
- **Smoke Codex** (#30244 OPEN): issue filed (cache_memory miss) but run succeeded ✅
- **Smoke macOS ARM64** (NO ISSUE): 100% failure since Feb 2026 — still needs dedicated issue
- **Auto-Triage Issues** (#30205 OPEN)
- **Documentation Unbloat** (#30188 OPEN)
- **Daily Documentation Healer** (#30233 OPEN)
- **Matt Pocock Skills Reviewer** (#30272, #30252 OPEN): APM unpack failure
- **Step Name Alignment** (#30069 OPEN)

### P2 Issues
- Node.js 20 deprecation deadline Sep 16, 2026
- MCP gateway session timeout (#23153)
- 6 PR-review agents with potential redundancy (Scout, Archie, /cloclo, Q, AI Moderator, Content Moderation)
- Performance regression in Validation: 82.1% slower (#30180)

### Recovery
- Smoke Codex: SUCCESS at 00:38 UTC ✅
- Smoke Copilot: SUCCESS ✅

### Systemic Observation
- APM unpack failure (#30252) is a NEW systemic issue affecting multiple workflows on different PRs/firewall versions — root cause investigation needed urgently.

### Actions Taken This Run
- Updated shared memory
- Added health dashboard comment to #29109

### Trends
- Score: 63/100 (↓ -2 from APM systemic)
- APM unpack regression: newly impacts CI reliability
- Gemini still completely broken (30+ days, P0 unresolved)
- macOS ARM64 chronic failure since Feb 2026 — no issue filed
