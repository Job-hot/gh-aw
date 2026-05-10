# Workflow Health — 2026-05-10T05:39Z

Score: 61/100 (→ stable). 218 workflows. Run: §25620950271

## KEY FINDINGS

### Compilation Status
- 218/218 lock files present ✅ (no change)
- 0 missing lock files ✅
- Compile validation unavailable (gh aw extension not in PATH)

### Recent Run Observations (last ~100 runs)
- PR review cluster (Q, Scout, Archie, cloclo, Grumpy, Security Review, PR Nitpick, PR Code Quality) — action_required on PRs (expected skip behavior)
- Deployment Incident Monitor: 8x skipped (zombie pattern continues)
- Successes: Semgrep, Grafana OTel, Quality Sentinel, Skill Optimizer, Code Simplifier, Monster Issue, Metrics Collector, Community Attribution, Decision Gate, Dependency Cleaner, Compiler Quality, Compiler Threat, Copilot Cloud, Noob Tester, Name Alignment, etc.
- `printf` allow-list workflow: 2 success, 8 action_required (mixed, expected for PR-triggered)
- GitHub API JSON mode unavailable this run (gh run list --json failed)

### P0 Issues (Active — carried forward)
- **Smoke Gemini** (#29666 OPEN): 100% failure, proxy/API-key blocks. 35+ days.
- **Smoke CI** (#29666 OPEN): CGO/EROFS persistent
- **Daily Model Inventory Checker** (#30043 OPEN): Copilot CLI silent startup crash
- **APM Unpack** (#30252 OPEN): apm-default.tar.gz exit code 1
- **config.models** (#30307 OPEN): unsupported AWF config field

### P1 Issues (Active — carried forward)
- **Smoke macOS ARM64**: Issue filed 2026-05-07 ✅
- **CI regression** TestStrictModePermissions: Issue filed 2026-05-06
- **MCP gateway session timeout** (#23153 OPEN)
- **Performance Regression** (#30180): 82.1% slower
- **Deployment Incident Monitor**: zombie pattern — 8x skipped today

### Actions Taken This Run
- Updated shared memory files
- Added comment to dashboard issue #29109
- No new issues created (no new critical failures identified)

### Trends
- Score: 61/100 (→ stable, day 4)
- 218 workflows (stable)
- Overall system health stable; no new P0/P1 regressions detected
