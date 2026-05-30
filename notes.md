# Copilot Session Insights — repo memory

## 2026-05-30 snapshot

- 50 sessions; 26% completion (13/50) — rebound after 05-29 dip (14%→26%)
- 2 branches, 100% top-2: fix-pi-agent-configuration (37, 12/37=32%) + share-runner-temp-env-var (13, 1/13=8%)
- action_required 22%, skipped 34%, failure 18%
- BURST PATTERN BROKEN: 25-sess burst on fix-pi = 7 success/16 skipped/2 fail (NOT action_required); 7-sess share-runner burst = 100% action_required. Outcome is branch-dependent → new pattern burst_outcome_branch_dependence (supersedes synchronized_burst_saturation).
- Bimodal: median 0.175m vs mean 2.23m; 37 sub-30s, 8 ≥5min, 3 ≥15min (max 18.95m)
- Orphan escalations: 0; unassigned PRs 6/8=75% but all idle (0 gates) → not orphaned
- Conversation logs empty 7th consecutive day

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
