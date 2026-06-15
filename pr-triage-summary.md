# PR Triage Memory
Last: 2026-06-15T01:29:52Z | Run: 27518690601 | Fork PRs: 0

## PR Triage Summary - 2026-06-15T01:29Z (Run 27518690601)

### Status
**No fork PRs to triage** (fork-only policy; all 8 open agent PRs are branch-based)

### Details
- **Run ID**: 27518690601
- **Open Agent PRs Found**: 8 (all branch-based, skipped per fork-only policy)
- **Fork PRs Found**: 0
- **PRs Triaged**: 0
- **New Since Last Run**: #39282, #39300, #39299
- **Closed Since Last Run**: #39243
- **Unchanged**: #39266, #39164, #39156, #39100, #38911

### Branch PRs Skipped (fork-only policy, scores informational)
| # | Title | Cat | Risk | Score | Action | CI | Review |
|---|---|---|---|---|---|---|---|
| #39100 | Run safe-outputs MCP in gh-aw node container | feature | high | 72 | fast_track | pending | bot-approved |
| #39282 | Add Copilot SDK idle-hang watchdog | feature | high | 66 | fast_track | ✅ | commented |
| #38911 | [ARC/DinD] chroot.binariesSourcePath + identity | feature | high | 65 | batch_review | ✅ 19/19 | CHANGES_REQUESTED |
| #39266 | Branch-aware cache-miss semantics | bug | high | 61 | batch_review | pending | dismissed |
| #39156 | Stop Codex retries on router failures | bug | medium | 56 | batch_review | pending | CHANGES_REQUESTED |
| #39300 | Wildcard validation: PR tools (DRAFT) | bug | low | 46 | batch_review | pending | none |
| #39299 | Wildcard validation: review comment (DRAFT) | bug | low | 46 | batch_review | pending | none |
| #39164 | supertonic TTS standalone dispatcher (DRAFT) | feature | low | 28 | defer | pending | none |

### Batches
- batch-wildcard-validation: #39299, #39300
- batch-changes-requested: #38911, #39156
- batch-bulk-workflow: #39282, #39266

### Notable
- #39100 removes 12,551 lines; human review recommended
- #39282 fixes 24.5-min silent hang; CI passing; fast-track candidate
- #38911 CI all green (19/19) but CHANGES_REQUESTED; 3d stale
