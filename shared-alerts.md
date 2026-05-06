# Shared Alerts — 2026-05-06T13:14Z

## P0 (Critical)
- **Smoke Gemini** (#30175, #29666 OPEN): 100% failure, proxy architecture blocks all agent traffic. 31+ days unresolved.
- **Smoke CI** (#29666 OPEN): CGO/EROFS persistent, 100% action_required.
- **Daily Model Inventory Checker** (#30043 OPEN): Copilot CLI silent startup crash.
- **APM Unpack Systemic Failure** (#30252 OPEN): apm-default.tar.gz exits code 1, affects multiple workflows.

## P1 (High)
- **CI regression on main** (May 6): `TestStrictModePermissions/no_permissions_specified_allowed_in_strict_mode` failing since ~01:16 UTC May 6. Playwright MCP deprecation error in strict mode integration test. Issue filed via safeoutputs.
- **config.models unsupported field** (#30307 OPEN): blocked 10 smoke runs in config.models sweep.
- **Smoke macOS ARM64**: 100% failure since 2026-02-20 (~75 days). NO ISSUE FILED — needs one urgently.
- **MCP gateway session timeout** (#23153 OPEN): Long-running workflows at risk.
- **Performance Regression in Validation** (#30180): 82.1% slower.

## P2 (Watch)
- **Node.js 20 deprecation** in CI: deadline Sep 16, 2026. Migrate to Node.js 22.
- **9 PR-review agents** on same triggers — 100+ action_required runs/day (Scout, Archie, /cloclo, Q, AI Moderator, Content Moderation, Grumpy, Security, PR Nitpick). Consolidation recommended.

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
- 214 workflows, 0 missing lock files
- Health: 61/100 (↓-2, CI regression on main)
- Quality: 74/100 stable (plateau day 5)
- Effectiveness: 71/100 stable (plateau day 5)
- 6 P1 issues resolved this cycle ✅ (good recovery)
