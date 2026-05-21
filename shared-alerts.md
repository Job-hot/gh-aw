# Shared Alerts — 2026-05-21T05:52Z

## P0 (Critical) 🚨
- **CGO/CJS regression** (#29669): **CRITICAL ESCALATION REQUIRED**
  - **90+ days unresolved at 0% success rate** — critical threshold exceeded
  - **Continuous failures May 21:** 10+ failures in last few hours
  - **Impact:** Every push to main triggers multiple failures, blocking CI confidence
  - **ACTION NEEDED:** Dedicated engineering escalation (beyond workflow automation scope)

## P1 (High) 🚨
- **Codex OPENAI_API_KEY sandbox exclusion** (#32446): blocking all Codex workflows (12 workflows)
- **MCP gateway session timeout** (#23153): long-running workflows at risk, ~5min inactivity timeout
- **Performance Regression in Validation** (#30180): 82.1% slower

## P2 (Watch) ⚠️
- **ET budget exhaustion**: Multiple daily workflows at risk. Audit `max-effective-tokens`. #32717
- **Engine-fail-after-completion** pattern persists (#32736) — systemic engine lifecycle bug
- **Step Name Alignment recurring daily**: #32955 (was #32754 closed May 17) — **needs structural fix, not patches**
- **[aw-compat] Cross-repo warnings** (#32528): P2

## Resolved (Do Not Re-File) ✅
- PR-review cluster #31724: CLOSED ✅ (was ~272 wasted runs/day)
- May 14 mass failure (#32045-#32119): RESOLVED ✅
- AWF Firewall v0.25.47 #32522: CONTAINED ✅
- Sergo #32755: CLOSED ✅
- Step Name #32754: CLOSED ✅ (recurred as #32955)

## Outlook & Coordination Notes

### Critical Actions Needed
1. **CRITICAL:** Escalate CGO/CJS (#29669) to dedicated engineering — 90+ days at 0% = critical threshold
   - Continuous failures May 21 (10+ failures in hours)
   - Every push to main triggers failures
   - Beyond workflow automation scope — needs dedicated engineering resources
2. **HIGH:** Fix Codex sandbox configuration (#32446) — 12 workflows blocked
3. **HIGH:** Resolve MCP gateway session timeout (#23153) — long-running workflows fail
4. **MEDIUM:** Structural fix for Step Name Alignment (#32955) — daily noise, stop patching
5. **MEDIUM:** Fleet-wide ET budget audit (#32717) — prevent token exhaustion

### Cross-Orchestrator Impact
- **CGO/CJS critical threshold** → CI confidence degraded, every push triggers failures
- **Token budget pressure** → Multiple workflows at risk of ET exhaustion
- **Quality stable** at 63/100 but cannot improve until critical P0/P1 resolved

### Fleet Statistics (May 21)
- 233 executable workflows (100% lock file coverage ✅)
- Health: 63/100 (stable)
- Estimated success rate: ~77% (historical, API limited)
- CGO/CJS: 0% (100% failure, ongoing May 21)
- Critical issues: 4 P0/P1 (stable)
- Warnings: 4 P2 (stable)

### Network Constraints (May 21)
GitHub API access limited during this run. Analysis based on:
- Pre-computed workflow inventory (233 workflows)
- Recent workflow run queries (CGO/CJS confirmed ongoing)
- Lock file verification (100% coverage confirmed)
- Historical shared memory (May 19)

Full live analysis will resume when API access restored.

Last updated: 2026-05-21T05:52:20Z by workflow-health-manager
