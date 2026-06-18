# Copilot Session Insights — repo memory

## 2026-06-18 snapshot
- 50 sessions; **4% completion** (2 success) — uptick from 0% (06-17); recent 7d avg ~11.1% vs prior 7d ~12.0% (flat)
- **PURE GATE-SWEEP**: 48/50 action_required, 48 zero-duration; window 06:24–07:10Z (46 min)
- Both successes AGENTIC: "Running Copilot cloud agent" 12.8m + "Addressing comment on PR #39959" 12.3m
- 9 copilot/* branches; top-3=64% (feature-ai-authorship-footer 11, fix-extract-base-branch-regex 11, clean-up-gh-aw-containers 10)
- Orphans: **0**; 8 open PRs ALL Copilot-assigned → orphan_rate 0% (vs 40% baseline, NORMAL)
- Conversation logs empty **18th+ day**; experimental none (roll=40)

## Active patterns
- provenance_inversion (06-07): on recovery days gate/moderator workflows CAN produce successes; 06-18 reverts — both successes agentic (cloud-agent + PR-comment), NOT gate workflows
- copilot_cloud_agent_reliability: holds — cloud agent remains a consistent success source
- inverse_gate_count_to_conclusiveness: holds; Copilot-assigned ⇒ never orphaned (0 orphans again 06-18)
- gate_sweep_zero_duration: recurring — snapshots routinely catch 48-50 action_required 0-duration runs
- conversation_log_fetch_failure: 18th+ day (longest unresolved risk)
