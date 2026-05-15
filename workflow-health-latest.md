# Workflow Health — 2026-05-15T05:43Z

Score: 64/100 (↑ +2). 229 workflows. Run: §25902421370

## KEY FINDINGS

### Compilation Status
- 229 workflows in workflow-list.txt
- Compile validation unavailable (gh aw extension not in PATH)
- No new missing lock files reported

### Recovery from May 14 Mass Failure Event
- **AI Moderator**: ✅ SUCCESS today (was failed May 14)
- **Content Moderation**: ✅ SUCCESS today (was failed May 14)
- **Agentic Commands**: ✅ SUCCESS today
- **Smoke CI**: ✅ SUCCESS today (was failing)
- **Doc Build - Deploy**: ✅ SUCCESS today (was action_required)
- **Safe Output Health Monitor**: ✅ SUCCESS today
- **License Compliance Check**: ✅ SUCCESS today
- PR #32070 (safe output bundle fix) merged May 14 — appears to have restored many workflows

### Persistent Failures (May 15)
- **CGO** (#29669): 2 failures + 3 action_required today — ongoing push regression (P1)
- **CJS** (#29669 related): 2 failures + 3 action_required today — ongoing push regression (P1)

### Today's Run Summary (May 15)
- Total: 80 runs
- Success: 12
- Failure: 4 (all CGO/CJS)
- action_required: 56 (no-trigger / skipped)
- In-progress: ~4

### Open Tracking Issues
- #29669: CGO/CJS regression — STILL OPEN (persistent)
- #31432, #31524: Daily Fact parse failures — STILL OPEN
- #23153: MCP gateway session timeout — STILL OPEN
- #30180: Performance Regression — STILL OPEN
- #31724: PR-review cluster waste (~272 runs/day at 0%) — STILL OPEN
- ~30+ open [aw] failure issues (mass event May 14, partial recovery today)

### Actions Taken This Run
- Added comment to dashboard issue #29109
- Updated shared memory

### Trends
- Score: 64/100 (↑ +2 from 62 — partial recovery post May-14 mass failure)
- Mass failure recovery: significant ✅ (PR #32070 effective)
- CGO/CJS regression: persistent, no resolution yet
- P0 count: 0 ✅
- 229 workflows (+4 from 225)
