# Copilot Session Insights — repo memory

## 2026-06-05 snapshot
- 50 sessions; **8% completion (4/50)** — flat vs 06-04; 2nd day at floor; below 7d avg 14.0%
- **92% action_required (46/50)**; 0 failures/skips
- **Symmetry reversal vs 06-04 bifurcation**: all 3 branches scored ≥1 success → new pattern `branch_symmetric_success_distribution` (yaml-outputs 1/21, close-pr 1/18, cost-mgmt 2/11)
- success_duration_floor holds (4.82/7.9/8.9/15min, all ≥4.82); 46 zero-dur sweeps
- Provenance: cloud-agent 15m, 2× Addressing-comment-PR#37072, Copilot-Code-Review 4.82m
- Concentration: 3 branches; top 42%, top-3 100% (all copilot/*)
- Orphans: **0**; 2 in-progress runs both on main; 4 open PRs all Copilot-assigned, 0 PR gates → orphaned_rate 0% (vs 40% baseline)
- Conversation logs empty 13th+ day (OAuth) — behavioral analysis blocked
- Experimental: `Branch-Symmetric Success Distribution`

## Active patterns
- success_duration_floor: robust (all successes ≥4.82 min)
- branch_symmetric_success_distribution: NEW (06-05), inverse of 06-04 bifurcation
- inverse_gate_count_to_conclusiveness: reconfirmed (Copilot-assigned ⇒ never orphaned)
- recovery_regression_oscillation: completion oscillates 4–46%, no trend
- conversation_log_fetch_failure: 13th+ day (longest unresolved risk)

## Files
- cache history.json — 17 analyses; patterns.json — 16 patterns
- history-entry-2026-06-05.json — today's entry
