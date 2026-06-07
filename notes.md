# Copilot Session Insights — repo memory

## 2026-06-07 snapshot
- 50 sessions; **40% completion (20/50)** — MAJOR RECOVERY; breaks the 4-day `sustained_completion_floor` (<=8% on 06-03..06-06); highest since 05-26 (46%); 7d avg 10.6%→**12.3%**
- **50% action_required** (25/50, all 0-min) — no longer a pure gate-sweep day
- **2 failures + 1 cancelled, all CGO gate** across 2 branches (isolated build-gate flakiness); 2 skipped
- **PROVENANCE INVERSION (new pattern)**: 17/20 successes from gate/moderator/review workflows (Test Quality Sentinel, PR Code Quality Reviewer, Design Decision Gate, Matt Pocock, Agentic Commands, Smoke CI, Addressing-comment x4, Code Review) — only 3 from Running-Copilot-cloud-agent. CONTRADICTS `success_workflow_provenance_mapping` (06-03).
- **success_duration_floor BROKEN low**: 1.78m Code-Review success today; success durations span 1.78–25.85m (avg 10.32m)
- Concentration: 6 branches, top aw-migrate-max-effective-tokens 30%, top-3 88%; all copilot/*
- Synchronized burst: all 50 in 05:57–06:46 window
- Orphans: **0**; 1 open PR (bot actions/update) + 1 in-progress run (this analysis on main); 0 PR gates → orphan_rate 0% (vs 40% baseline, healthy, 4th consecutive 0-orphan day)
- Conversation logs empty **15th+ day** — behavioral/loop/context analysis still blocked
- Experimental: none (standard run, roll=74)

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
