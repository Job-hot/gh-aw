# Workflow Health — 2026-05-26T05:50Z

Score: 70/100 (↑2 from May 25). ~236 workflows. Run: §26434802866

## KEY FINDINGS

### Status (May 26)
- **Compilation:** 236/236 workflows have lock files (100% ✅)
- **Last 48h Runs:** 100 runs analyzed (28% success, 4% failure, 54% action_required PR approvals)
- **Actual failures (excl. PR approvals):** 4 runs (CGO failure, CGO/CJS cancelled, Copilot cloud agent cancelled)
- **Health Score:** 70/100 (slight improvement — fewer failures, PR approval backlog stable)

### Persistent Critical Issues (P0/P1) 🚨
- **CGO/CJS regression** (#29669): Build gh-aw failures still occurring — 2 CGO/1 CJS runs failed/cancelled today
  - **Still unresolved — 90+ days critical threshold exceeded**
  - Impact: cancels dependent workflows, degrades CI confidence
- **Copilot cloud agent cancellation**: 1 cancelled run (likely cascading from CGO build failure)

### Stable/No Change
- #34582 Step Name Alignment — still open, recurring
- #34556 Copilot deprecated beta header — still open
- Codex OPENAI_API_KEY sandbox (#32446) — still tracked

### PR Approval Backlog
- 54/100 runs were `action_required` — PRs needing approval
- Not actual failures — expected behavior for fork PRs

### Actions Taken This Run
- Verified 236/236 lock file coverage (100% ✅)
- Analyzed 100 runs (4 actual failures, all related to CGO/CJS build issue)
- All problems covered by existing open issues
- **No new issues created** — all problems already tracked

### Trends
- Score: 70/100 (↑2 from May 25, ↑7 from May 22)
- Actual failures: 4 (down from 9 May 25)
- CGO/CJS: ongoing (P0 #29669, unresolved 90+ days)

Last updated: 2026-05-26T05:50:00Z by workflow-health-manager
