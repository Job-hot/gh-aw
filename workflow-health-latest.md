# Workflow Health — 2026-05-25T05:55Z

Score: 68/100 (↑5 from May 22). ~235 workflows. Run: §26385705325

## KEY FINDINGS

### Status (May 25)
- **Compilation:** 235/235 workflows have lock files (100% ✅)
- **Today's Runs:** 100 runs analyzed (14% success, 9% actual failure, 74% action_required/pending PRs)
- **Actual failures (excl. PR approvals):** 9 runs
- **Health Score:** 68/100 (slight improvement — fewer actual failures, PR approval backlog skewing stats)
- **CGO/CJS critical:** Ongoing — 6+ failures today (Build gh-aw failures)

### Persistent Critical Issues (P0/P1) 🚨
- **CGO/CJS regression** (#29669, #34574): Build gh-aw failures causing cascading workflow failures
  - Today: Train Log Pattern Weights, Safe Output Health Monitor, PR Sous Chef, Copilot cloud agent (4x) all failed due to build failure
  - **Still unresolved — 90+ days critical threshold exceeded**
- **Step Name Alignment** (#34582): Recurring daily failure — Execute Claude Code CLI
- **Copilot engine deprecated beta header** (#34556): User-reported — `anthropic-beta: context-1m-2025-08-07` rejected by Anthropic API (400 error)

### Closed/Stable Issues
- #34587 Safe Output Health Monitor - issued today (build failure / CGO)
- #34586 PR Sous Chef - issued today (build failure / CGO)
- #34582 Step Name Alignment - issued today (recurring)
- Codex OPENAI_API_KEY sandbox (#32446): still tracked

### PR Approval Backlog
- 74/100 runs were `action_required` — all from PRs needing approval
- Not actual workflow failures — expected behavior for fork PRs

### Actions Taken This Run
- Verified 235/235 lock file coverage (100% ✅)
- Analyzed 100 runs (9 actual failures, all related to CGO/CJS build issue)
- Confirmed all failures already have open issues
- **No new issues created** — all problems covered by existing issues
- Updated shared memory

### Recommendations for Next Run
1. **CRITICAL:** Escalate CGO/CJS to dedicated engineering (#29669) — 90+ days threshold
2. **High:** Fix Copilot deprecated beta header (#34556)
3. **Medium:** Structural fix for Step Name Alignment (#34582) — daily recurrence
4. **Medium:** Reduce PR approval backlog — 74% of runs blocked

### Trends
- Score: 68/100 (↑5 from May 22, 63/100)
- Actual failures: 9 (down from ~13 May 22)
- PR approval backlog: 74 runs pending
- CGO/CJS: 0% success rate on impacted workflows (ongoing)

Last updated: 2026-05-25T05:55:59Z by workflow-health-manager
