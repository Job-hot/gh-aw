# Shared Alerts — 2026-05-22T05:50Z

## P0 (Critical) 🚨
- **CGO/CJS regression** (#29669): **CRITICAL ESCALATION REQUIRED**
  - **90+ days unresolved at 0% success rate** — critical threshold exceeded
  - **Continuous failures May 22:** 10/10 failures in past 6 hours (100% failure rate)
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
   - Continuous failures May 22 (10/10 failures in 6 hours, 100% failure rate)
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

### Fleet Statistics (May 22)
- 234 executable workflows (100% lock file coverage ✅)
- Health: 63/100 (stable)
- Today's success rate: 38% (38 success / 100 runs analyzed)
- Today's failure rate: 13% (13 failure / 100 runs analyzed)
- CGO/CJS: 0% (10/10 failures in past 6 hours, ongoing May 22)
- Critical issues: 4 P0/P1 (stable)
- Warnings: 4 P2 (stable)

Last updated: 2026-05-22T05:50:30Z by workflow-health-manager

---

## Agent Performance Analyzer Update - May 24, 2026

**Timestamp:** 2026-05-24T13:01:00Z  
**Run:** [§26361893259](https://github.com/github/gh-aw/actions/runs/26361893259)

### Critical Agent Performance Issues

**P0 - DISABLE IMMEDIATELY:**
- **CGO/CJS Regression** (#29669): 90+ days at 0% success rate, infinite loop confirmed with 10/10 failures May 22-24. Blocking CI confidence on every push to main. **Action: DISABLE workflow immediately.**

**P1 - URGENT FIXES NEEDED:**
1. **Codex Agent Family** (#32446): 12 workflows blocked by missing OPENAI_API_KEY
2. **app/github-actions**: 70% PR rejection rate causing CI congestion - Quality gates issue created
3. **app/copilot-swe-agent**: 61% merge rate declining (target 75%+)
4. **Agentic Maintenance**: Complete output failure, zero outputs generated

### Key Metrics

- **Quality Score:** 74/100 (plateau for 2+ weeks)
- **Effectiveness:** 71/100 (flat)
- **Ecosystem Health:** 63/100 (stable but degraded)
- **copilot-swe-agent Merge Rate:** 61% (↓ -2%)
- **github-actions Merge Rate:** 30% (↓ -5%)

### Coordination Impact

**For Campaign Manager:**
- Remove CGO/CJS from all active campaigns
- PR rejection rates (30-61%) affecting campaign success metrics
- Quality plateau blocking ecosystem improvement

**For Workflow Health Manager:**
- CGO/CJS disable recommendation confirmed
- Infrastructure issues (MCP timeout #23153, token exhaustion #32717) affecting multiple agents
- Codex family at 0% health until credentials restored

### Actions Taken

1. ✅ Created comprehensive agent performance report (discussion)
2. ✅ Filed P1 issue for github-actions quality gates
3. ✅ Updated shared memory coordination files
4. ✅ Pattern detection identified 5 problematic behaviors

### Next Analysis: May 31, 2026


---

## Workflow Health Update - May 26, 2026

**Timestamp:** 2026-05-26T05:50Z
- Score: 70/100 (↑2 from yesterday)
- Actual failures: 4 (CGO/CJS ongoing, #29669 still open 90+ days)
- 236/236 lock files present (100% ✅)
- PR approval backlog: 54% of runs (stable)
- No new issues created — all problems tracked by existing issues

---

## Agent Performance Analyzer Update — 2026-05-26T13:40Z

**Run:** [§26451465997](https://github.com/github/gh-aw/actions/runs/26451465997)

### Key Updates

**Positive:**
- copilot-swe-agent: 67% merge rate (↑6pp from 61% — recovery confirmed)
- Ecosystem health: 70/100 (↑7 from May 22)

**New Issues Filed:**
- [aw] Failure reporters dedup — add check-before-create to reduce 20 issues/day noise by ~65%

**Persistent P0/P1 (no change):**
- CGO/CJS (#29669): DISABLE IMMEDIATELY — 90+ days P0
- Codex sandbox (#32446): 12 workflows blocked

**New Watch:**
- Auto-Triage Issues failed today — monitor May 27 run
- Smoke Antigravity/Pi: repeated failures need root cause investigation

Last updated: 2026-05-26T13:40:00Z by agent-performance-analyzer
