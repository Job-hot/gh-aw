# Copilot Session Insights — repo memory

## 2026-05-28 snapshot

- 50 sessions analyzed; 28% completion rate (3rd consecutive recovery day: 22% → 28%)
- First two-day monotonic improvement in the 9-day window
- 92% concentration on 2 PR branches (api-proxy-reflect-call-resilience 70%, convert-python-to-javascript 22%)
- Per-branch story:
  - api-proxy-reflect-call-resilience: 13/35 success (37%), 0 action_required, 18 skipped, 4 failed
  - convert-python-to-javascript: 1/11 success, 10 action_required (the rate-drag culprit)
- 7 sessions ≥15min — highest since 2026-05-23; mostly successes on lead branch
- Cloud agent 1/1 success at 32.7 min (longest session)
- Orphan rate: 0% (vs ~40% baseline)
- Conversation transcripts empty for 5th consecutive day — behavioral analysis still blocked

## Active patterns
- bimodal_duration_distribution: confirmed (35 <30s + 10 ≥5min)
- success_duration_floor: 11/14 successes ≥3.6 min
- branch_specific_friction: NEW — one branch (convert-py-to-js) producing ~70% of action_required
- orphan_rate_healthy: NEW — 0% today, all active branches have Copilot assigned
- conversation_log_fetch_failure: 5-day streak, longest-running unresolved workflow risk

## Files
- /tmp/gh-aw/cache-memory/session-analysis/history.json — appended 2026-05-28
- /tmp/gh-aw/cache-memory/session-analysis/patterns.json — updated with 2 new patterns
