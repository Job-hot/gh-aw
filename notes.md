# Copilot Session Insights — repo memory

## 2026-06-09 snapshot
- 50 sessions; **0% completion** — NEW 21d FLOOR (4%→0%); 3rd zero-day in window (05-20,05-25,06-09), first since 05-25; 7d avg 11.7%→**9.4%** (prior 7d 20.3%)
- **PURE GATE-SWEEP**: 50/50 action_required, 0 success/fail/cancel; every session 0-duration (created==updated) — snapshot caught all runs pre-completion
- Window 06:36–06:46 (10 min); 4 copilot/* branches × 7 gate workflows (Q 11, Agentic Commands 11, Smoke CI 7, Doc Build 7, CGO 7, CWI 6, CJS 1)
- Branch concentration: fix-sortslice-precision 18 (36%), static-analysis-report 14, update-docs-for-sortslice 12, aw-fix-tool-denial-cluster 6; top-3 88%; **3 of 4 = sortslice cluster**
- Orphans: **0**; 15 open PRs (3 unassigned housekeeping, 0 gates); 3 in-progress runs all on main → orphan_rate 0% (vs 40% baseline, healthy)
- Conversation logs empty **17th+ day** — behavioral analysis still blocked
- Experimental: none (standard run, roll=91)

## 2026-06-07 snapshot (condensed)
- 40% completion (recovery, broke 4-day floor); provenance_inversion observed (gate/review workflows produced successes); success_duration_floor broken low (1.78m); 0 orphans; logs empty 15th day.

## Active patterns
- provenance_inversion: NEW (06-07) — on a recovery day gate/moderator workflows DO produce successes; weakens success_workflow_provenance_mapping
- sustained_completion_floor: BROKEN 06-07 (40% after 4 days <=8%)
- success_duration_floor: WEAKENED 06-07 (1.78m success observed)
- copilot_cloud_agent_reliability: holds but no longer sole success source
- inverse_gate_count_to_conclusiveness: holds; Copilot-assigned ⇒ never orphaned
- conversation_log_fetch_failure: 15th+ day (longest unresolved risk)

## Files
- cache history.json — 19 analyses; patterns.json — patterns
- history-entry-2026-06-07.json — today's entry
