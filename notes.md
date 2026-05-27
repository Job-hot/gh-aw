# Session Analysis Notes

## Persistent Observations

- **Data quality**: `infrastructure-only` / `metadata-only` for 4+ weeks. Conversation transcripts have never been reliably available — `copilot-session-data-fetch` returns OAuth errors or simply produces an empty `logs/` directory. This is the longest-running unresolved workflow risk.
- **Recovery/regression oscillation**: completion rate bounces between near-0% and 40-50% rather than degrading steadily. Three "good" days in 8 (44%, 46%, 22%) interleaved with five sub-12% days. Recoveries are not sustained trends.
- **Branch concentration drives outcomes**: 70-90% of daily sessions concentrate on 2-4 branches. Concentration on agent-assigned PR branches → higher completion. Concentration on agent-less branches → action_required dominance.
- **Orphan rate**: 0% for 8+ consecutive days. The 40% historical baseline is clearly outdated; gating is no longer accumulating on un-owned branches.
- **Action_required share**: the volatile metric. Ranges 8% → 98% day-to-day. Tracks inversely with completion (today: 70% action_required + 22% success).

## Run: 2026-05-27 (latest)

- **22% success** (11/50) — partial recovery: above slump-day baseline (0–2%) but well below 2026-05-26's 46%
- 70% of sessions on a single PR branch (`copilot/fix-stale-npm-metadata`, Copilot+mnkiefer assigned); 88% on top 2 PR branches (both Copilot-assigned)
- Avg 2.85m, median 0m, max 15.0m — 0 sessions ≥20m; 36/50 finished in <30s
- Success durations cluster between 9–15m (8 of 11 successes); shortest success 1.9m
- 1 Running Copilot cloud agent session — 15.0m, **succeeded** (contrast with 2026-05-25's 2 failures)
- 100%-success workflows today: `Running Copilot Code Review` (2/2), `Running Copilot cloud agent` (1/1)
- 0%-success high-volume workflows today: `Q` (0/9), `Agentic Commands` (2/9), `CGO` (1/5), `Smoke CI` (1/5)
- 0 spec-orphans / 0 escalation candidates (4 in-progress runs all on `main`)
- Conversation logs unavailable (4th consecutive day; logs/ directory empty)

## Cross-Day Workflow Reliability Pattern

`Running Copilot Code Review` is a candidate "always-conclusive" gate — 2/2 success on 2026-05-27 and consistently conclusive on prior good days. Worth tracking as a stability anchor.
