# Workflow Health — 2026-04-28T12:20Z

Score: 57/100 (↓-17 from 74 Apr 27). 204 workflows. Run: §25052372422

## KEY FINDINGS

### Compilation Status
- 204/204 lock files present ✅
- **0 missing lock files** ✅
- **0 stale lock files** ✅

### Today's Failures (Apr 28)
13 scheduled runs failed out of 30 (57% success rate — significant regression from 93% yesterday)

**Category 1: THREAT_DETECTION_RESULT parse failure (systemic)**
- **Dead Code Removal Agent** — detection job: "No THREAT_DETECTION_RESULT found"
- **Daily Testify Uber Super Expert** — detection job: same error
- **Update Astro** — detection job: same error
- Tracked in #28866 ([aw] Detection Runs)

**Category 2: Agent job failures (various)**
- **Daily Fact About gh-aw** — codex engine agent crash (P0 ongoing)
- **Sub-Issue Closer** — agent job crash (no OTEL = possible runtime issue)
- **Daily Team Evolution Insights** — agent job crash
- **Daily AstroStyleLite Markdown Spellcheck** — agent crash (no OTEL)
- **Daily Rendering Scripts Verifier** — agent crash (Docker/Playwright env)
- **Developer Documentation Consolidator** — agent crash (Docker env)
- **Semantic Function Refactoring** — agent crash

**Category 3: Safe outputs failure**
- **Daily Documentation Updater** — safe_outputs job failed

**Category 4: CI**
- **CI** (scheduled integration tests) — 4 integration test jobs failed

### P0 Issues (Active)
- **Daily Fact About gh-aw codex failure** (auto-issues #28703 etc): codex binary/engine crash. Daily recurring.

### P1 Issues (Carry from Apr 27)
- **Documentation Unbloat claude auth failure** (#28659 OPEN): Claude OAuth token issue
- **GitHub Remote MCP Authentication Test** (#27965 OPEN): gpt-5.1-codex-mini model not supported, day 7+
- **THREAT_DETECTION_RESULT parse failure** — NOW SYSTEMIC: affecting ≥3 workflows today (was 1-2 workflows). Tracked #28866.
- **Safe outputs session not found** (#23153 OPEN): long-running workflows
- **GitHub App rate limit** (#27251 OPEN)
- **awf-api-proxy sidecar unhealthy** (#27888 OPEN)
- **CODEX_HOME collision** (#27512 OPEN)

### P2 Issues
- **Safe Outputs SEC-004** (#27235 OPEN)
- **Performance regressions** (#27280/#27279/#27278 OPEN)
- **Daily Documentation Updater protected files** (#27801 OPEN)
- **MCP gateway long-running drops** (#23153 OPEN)

## Issues Created This Run
- None (existing tracking issues cover identified failures; threat detection tracked in #28866)

## Issues Updated
- None

## Positive Notes
- 204/204 workflows compiled, all lock files present
- Smoke Codex ran today (issue #28881) — mostly passing (2/8 checks failed: web-fetch unavailable, comment-memory unavailable)

## Regression Alert
- **Success rate dropped from 93% → 57%**: THREAT_DETECTION_RESULT parse failure expanded to hit multiple new workflows today. Possibly related to a detection model change or outage.
