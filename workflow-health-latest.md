# Workflow Health — 2026-05-22T05:50Z

Score: 63/100 (→0 from May 21). ~234 workflows. Run: §26270848573

## KEY FINDINGS

### Status (May 22)
- **Compilation:** 234/234 workflows have lock files (100% ✅)
- **Today's Runs:** 100 runs analyzed (38% success, 13% failure, 20% action_required)
- **Health Score:** 63/100 (stable but degraded from critical issues)
- **CGO/CJS critical:** 100% failure rate ongoing (10/10 failures today May 22, 0:00-5:50 UTC)

### Persistent Critical Issues (P0/P1) 🚨
- **CGO/CJS regression** (#29669): **90+ days unresolved, 0% success rate — CRITICAL ESCALATION NEEDED**
  - Continuous failures May 22: 10/10 failures in past 6 hours (100% failure rate)
  - Every push to main triggers multiple CGO/CJS failures
  - Pattern: "action_required" (2x) + "failure" (8x) in last 10 runs
  - **ACTION:** Needs dedicated engineering escalation (90+ day threshold exceeded)
- **Codex OPENAI_API_KEY sandbox exclusion** (#32446): 12 workflows blocked
- **MCP gateway session timeout** (#23153): long-running workflows at risk
- **Performance Regression** (#30180): 82.1% slower validation

### Persistent Warnings (P2) ⚠️
- **ET budget exhaustion** (#32717): multiple daily workflows at risk
- **Engine-fail-after-completion** (#32736): systemic lifecycle bug
- **Step Name Alignment recurring** (#32955): daily recurrence, needs structural fix
- **[aw-compat] Cross-repo warnings** (#32528): compatibility issues

### 🎉 Still Resolved
- #32755 Sergo ✅ CLOSED
- #32754 Step Name Alignment ✅ CLOSED (recurred as #32955)
- PR-review cluster #31724 ✅ RESOLVED (~272 wasted runs/day saved)
- May 14 mass failure ✅ RESOLVED
- AWF Firewall v0.25.47 ✅ CONTAINED

### Systemic Patterns
- **90+ day critical threshold:** CGO/CJS needs immediate dedicated engineering escalation
- **Token budget pressure:** ET exhaustion affecting multiple daily workflows
- **Recurring failure cycle:** Step Name Alignment needs structural fix, not patches
- **CI confidence impact:** CGO/CJS failures on every push to main

### Actions Taken This Run
- Verified 100% lock file coverage (234/234)
- Analyzed today's workflow runs (100 runs, 38% success rate)
- Confirmed CGO/CJS 100% failure rate continues (10/10 failures in 6 hours)
- Updated shared memory coordination files
- **No new issues created** - existing issues cover all known problems
- **No comments added** - issues already well-documented

### Trends
- Score: 63/100 (→0 stable)
- CGO/CJS: 0% success rate (10/10 failures today, ongoing)
- Today's fleet: 38% success, 13% failure, 20% action_required, 21% skipped
- Critical issues: 4 P0/P1 (stable)
- Warnings: 4 P2 (stable)

### Recommendations for Next Run
1. **CRITICAL:** Escalate CGO/CJS to dedicated engineering (#29669) — 90+ days threshold
2. **High:** Fix Codex sandbox OPENAI_API_KEY exclusion (#32446) — 12 workflows blocked
3. **High:** Resolve MCP gateway session timeout (#23153) — long-running workflows at risk
4. **Medium:** Structural fix for Step Name Alignment (#32955) — stop daily recurrence
5. **Medium:** Fleet-wide ET budget audit (#32717) — prevent exhaustion

Last updated: 2026-05-22T05:50:30Z by workflow-health-manager
