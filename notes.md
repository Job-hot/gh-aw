# Copilot Session Insights — repo memory

## 2026-06-06 snapshot
- 50 sessions; **2% completion (1/50)** — NEW non-zero window low; down from 8% (06-05); 7d avg fell 14.0%→**10.6%**
- **98% action_required (49/50)**; 0 failures, 0 skips — pure gate-sweep day
- **SUSTAINED COMPLETION FLOOR (new pattern)**: 4 consecutive days ≤8% (06-03..06-06) — `recovery_regression_oscillation` has flattened into a low-productivity regime
- Single success = `Running Copilot cloud agent` @13.97m on copilot/aw-failures-doc-healer-model-inventory-sentrux → confirms `copilot_cloud_agent_reliability` + `success_duration_floor` (≥~14m)
- bimodal holds: 49 sweeps ~0min + 1 substantive @13.97m
- Synchronized burst: all 50 in 15-min window 05:56–06:11; 5 copilot/* branches × ~10 CI gate workflows (Q, Agentic Commands, Smoke CI, CJS, CGO)
- Concentration: top fix-safe-output-health-monitor=15 (30%), top-2 56%, top-3 78%; all 5 branches copilot/*
- Orphans: **0**; 9 open PRs ALL Copilot-assigned; only 1 in-progress run repo-wide (this analysis workflow on main); 0 PR gates → orphaned_rate 0% (vs 40% baseline, healthy)
- Conversation logs empty **14th+ day** — behavioral/loop/context analysis still blocked
- Experimental: none (standard run, roll=55)

## Active patterns
- sustained_completion_floor: NEW (06-03..06-06), supersedes oscillation framing
- success_duration_floor: robust (all successes ≥~14min on cloud-agent days)
- copilot_cloud_agent_reliability: only workflow producing successes
- inverse_gate_count_to_conclusiveness: Copilot-assigned ⇒ never orphaned
- conversation_log_fetch_failure: 14th+ day (longest unresolved risk)

## Files
- cache history.json — 18 analyses; patterns.json — 17 patterns
- history-entry-2026-06-06.json — today's entry
