# Copilot Session Insights — repo memory

## 2026-06-03 snapshot

- 50 sessions; **4% completion (2/50)** — collapse from 16% (06-02); lowest since 05-25 (0%); 7d avg 17.7%
- **96% action_required (48/50)** — pure gate-sweep day, all zero-duration
- Both successes = `Running Copilot cloud agent` workflow (14.23m + 16.2m); gate workflows produced 0 successes → new strategy `Success-Workflow Provenance Mapping`
- success_duration_floor holds (both ≥14min); 48 sub-30s gate firings
- 4 agent branches: duplicate-code-fix 22 (44%), add-license 12, refactor-copilot-sdk-driver 8, implement-experiments-notify 8; top-2=68%
- duplicate-code-fix fired 21 gates in 54-min window but PR #36587 is Copilot-assigned → NOT orphaned (inverse_gate_count_to_conclusiveness reconfirmed)
- Orphan escalations: 0; only 1 in-progress run (on main); 1/5 open PRs unassigned = 20% (idle auto-update branch #36593, 0 gates) — well below 40% baseline
- Conversation logs empty for 11th+ consecutive day — behavioral analysis still blocked

## Active patterns
- recovery_regression_oscillation: confirmed — 2-day monotonic recovery broken on day 3
- success_duration_floor: continues (7 successes all ≥2.4 min)
- branch_specific_friction: convert-py-to-js replaced by remove-features-copilot-requests-field as drag culprit
- synchronized_burst_saturation: NEW — bursts = 100% action_required, never produce successes
- concentrated_branch_activity: peak concentration (100% top-2 share, 2 unique branches)
- conversation_log_fetch_failure: 6-day streak (now longest-running unresolved workflow risk)

## Files
- /tmp/gh-aw/cache-memory/session-analysis/history.json — appended 2026-05-29
- /tmp/gh-aw/cache-memory/session-analysis/patterns.json — updated; +1 new pattern (synchronized_burst_saturation)
