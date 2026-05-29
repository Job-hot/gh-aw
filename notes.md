# Copilot Session Insights — repo memory

## 2026-05-29 snapshot

- 50 sessions; 14% completion rate — recovery streak broken (22%→28%→14%)
- 100% top-2 branch share, only 2 unique branches active — highest concentration in 10-day window
- copilot/update-daily-outcome-report-guidance dominated: 34 sessions (68%), 6/34 success (17.6%)
- copilot/remove-features-copilot-requests-field: 16 sessions (32%), 1/16 success (6.3%) — the rate-drag culprit
- 7 long-tail successes (2.4–31.9 min); the success-duration-floor pattern continues to hold (all ≥2.4 min)
- 6 burst events (≥3 sessions sharing same branch+timestamp) covering 31/50 sessions — 100% action_required
- Non-burst sessions had 36.8% success vs 0% inside bursts → bursts are pure retry-wave noise
- Cloud agent 1/1 success at 9.0 min
- Orphan escalations: 0 (only 2 active runs, both on main)
- Unassigned-PR rate: 7/15 = 46.7% (slightly elevated vs ~40% baseline, no CI-waste impact)
- Conversation transcripts empty for 6th consecutive day — behavioral analysis still blocked

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
