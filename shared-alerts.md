# Shared Alerts — 2026-05-07T13:15Z

## P0 (Critical)
- **Smoke Gemini** (#30175, #29666 OPEN): 100% failure, proxy architecture blocks all agent traffic. 33+ days unresolved.
- **Smoke CI** (#29666 OPEN): CGO/EROFS persistent, 100% action_required.
- **Daily Model Inventory Checker** (#30043 OPEN): Copilot CLI silent startup crash.
- **APM Unpack Systemic Failure** (#30252 OPEN): apm-default.tar.gz exits code 1, affects multiple workflows.

## P1 (High)
- **CJS 100% failure** NEW (May 7 ~11:00 UTC): `determine_automatic_lockdown.test.cjs:70` — vi.fn() mock never called (0 calls). 4 consecutive failures. No issue filed yet.
- **Smoke macOS ARM64**: 100% failure since 2026-02-20 (76 days). **Issue FILED 2026-05-07** ✅
- **CI regression on main** (May 6): `TestStrictModePermissions` failing. Issue filed.
- **config.models unsupported field** (#30307 OPEN): blocks smoke runs.
- **MCP gateway session timeout** (#23153 OPEN): Long-running workflows at risk.
- **Performance Regression in Validation** (#30180): 82.1% slower.

## P2 (Watch)
- **Node.js 20 deprecation** in CI: deadline Sep 16, 2026. Migrate to Node.js 22.
- **9 PR-review agents** on same triggers — 100+ action_required runs/day. Consolidation recommended.
- **Doc Build - Deploy**: 100% action_required (3/3 runs). Deployment permanently stalled.
- **Content Moderation + AI Moderator**: scope-creep, firing on PR diff events outside domain.
- **Resource Summarizer Agent**: 2/2 skipped, zero outputs.

## Resolved (Do Not Re-File)
- #29863 Smoke Copilot regression → RECOVERED ✅
- #30205 Auto-Triage Issues → CLOSED ✅
- #30188 Documentation Unbloat → CLOSED ✅
- #30233 Daily Documentation Healer → CLOSED ✅
- #30069 Step Name Alignment → CLOSED ✅
- #30241 Smoke Claude → CLOSED ✅
- #30244 Smoke Codex → CLOSED ✅
- #30347, #30144 GitHub MCP Structural Analysis → CLOSED ✅
- #30085, #30086, #30087 Safe Outputs Conformance → CLOSED ✅
- #30102 Schema Consistency Checker → CLOSED ✅

## Trends
- 217 workflows (+3), 0 missing lock files
- Health: 61/100 (→ stable)
- Quality/Effectiveness plateau at 74/71 for 6 days
- New regression: CJS test failure (May 7)
