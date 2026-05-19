# Session Analysis Notes

## Recurring Observation: data_quality = "infrastructure-only"

Across 8 consecutive runs (2026-05-06 → 2026-05-16), conversation transcripts have NEVER been available from the `copilot-session-data-fetch` shared module. Today's `20260516-conversation.txt` is a 1-line OAuth error from `gh auth login`, confirming the auth-gated access path is the blocker. This is a known, persistent limitation rather than a regression. The workflow is effectively a CI gate-sweep / orphan-detection monitor, not a behavioral analyzer.

## Persistent Patterns (7-day window)

- Orphan rate: 0% every day for 5+ consecutive days (baseline is 40% per spec)
- Action_required share is the volatile metric: 8% → 16% → 100% (2026-05-13 spike)
- Branch concentration drives queue size: 2-7 unique branches; max gates on one branch has ranged 16-35
- Copilot agent success rate (when transcripts surface): 100% on most days, 57.9% on 2026-05-07

## Run: 2026-05-13

- 100% action_required (highest in window) — unusual spike, was 58% on 2026-05-12
- Workflow concentration: Scout/Q/Agentic Commands fire 12x each = 72% of queue
- 4 active branches: refactor-extract-safe-outputs-config (18), aw-failures-fix-daily-report-generator (14), 2x (9)
- All branches in warning band (1-2h wait), none critical yet

## Run: 2026-05-16

- 92% action_required (46/50) — back near the May-13 pattern after May 14/15 gap
- 6 active branches; top 2 absorb 66% of queue: update-claude-code-and-mcp-gateway (17), refactor-git-patch-header-parsing (16)
- 4 success runs: 2x "Running Copilot cloud agent" (top 2 branches) + 2x "Haiku Printer"
- 0 spec-orphans; 5 chaos/* PRs have no assignee but zero gates so they don't escalate
- Workflow fingerprint: Agentic Commands + Q = 26 firings (52%) — same concentration as May 13 but spread across 2 agent branches instead of 4
- Notable: both top branches got Copilot success runs AND a 5-fire gate burst at the success timestamp — sweep-after-success pattern

## Run: 2026-05-17

- 92% action_required (46/50) — identical share to 2026-05-16 (two-day-in-a-row run-stuck pattern)
- 6 active branches; queue is flatter than yesterday: top branch (scan-repeated-permission-denied-issues) holds only 11 of 50 (22%), no branch dominates
- 4 success runs are all real Copilot agent sessions on 4 distinct PR branches — every "investigate-safe-output*" branch + the agentic-workflow + the comment-addressing run on PR #32759 all completed
- Copilot agent duration spread: 8 / 11 / 15 / 20 min (avg 13.5 min) — wider than recent days but all in the >5-min "high-success" band per [historical_trend_regression strategy](session-analysis-strategies.json)
- Workflow fingerprint: Agentic Commands + Q = 24 firings (48%) — concentration relaxed slightly vs May-13/16
- 0 spec-orphans: only 2 in-progress runs system-wide and both on `main` (this workflow + Failure Investigator). The 3 chaos/* PRs from 05:54 have no assignee but zero in-progress gates, so they don't trip escalation. The `docs/update-dictation-glossary-*` PR is brand-new (06:37) and below the 1h waiting threshold.
- Notable: 3 of 4 success runs were close-paired (within ~6 min) with their `action_required` gate bursts on the same branch — same sweep-after-success pattern as 05-16 but tighter time spacing
- data_quality stays `infrastructure-only` for the 9th consecutive run (2026-05-06 → 2026-05-17). Conversation logs dir empty again.

## Run: 2026-05-18

- 74% action_required (37/50) — meaningful drop from the 92% level of 05-16 and 05-17 (first relief in the run-stuck pattern)
- 22% success (11/50) — first material lift above the 8% baseline; doubled the previous high (16% on 05-12)
- 13 active (non-zero-duration) sessions, up from 4 the prior two days — 3.25x more real agent work
- Median active duration 8.0 min (down from 15.1 on 05-17) — faster turnaround when work actually starts
- Long-running sessions (>10 min): 6 absolute (up from 3) but 46% of active vs 75% prior day — share dropped even as count grew
- Workflow fingerprint: Agentic Commands + Q = 32 firings (64%) — concentration *increased* despite better completion rate, so the win is on the active-session side, not noise reduction
- 4 "Running Copilot cloud agent" runs, all completed (durations 11.5 / 13.2 / 13.5 / 17.6 min — top end of recent range)
- 4 "Addressing comment on PR #XXXXX" runs all finished 8-12 min — tight scope is paying off
- 2 outright `failure` conclusions (first non-zero failures in window) — worth a follow-up
- 0 spec-orphans across 22 open PRs; 3 in-progress runs (all on main) — same orphan picture as the 10-day window
- data_quality stays `infrastructure-only` for the 10th consecutive run; `gh auth login` OAuth-error on conversation export persists

## Run: 2026-05-19

- 98% action_required (49/50) — **highest action_required share in the 12-day window**, reversing the 05-18 relief (74%) and exceeding even the 05-13 spike (100%) in concentration. The 05-18 lift was not sustained.
- 2% success (1/50) — single Copilot cloud agent run on `copilot/fix-duplicate-code-safe-output`, completing in **22.3 min** (longest single-session duration recorded in the window; prior max was 31.7 min daily-average on 05-07)
- 5 unique copilot/* branches absorb the queue; top branch holds 18/50 (36%) — same dominance shape as 05-13, but spread across one fewer branch
- Workflow fingerprint: Agentic Commands + Q = 32/50 (64%) — **identical concentration to 05-18** despite opposite completion outcome. Workflow fingerprint is therefore not a predictor of daily success.
- Two large bursts only: 5 fires at 06:39:48Z (success-timestamp gate burst — same sweep-after-success pattern as 05-16) and 4 fires at 05:56:00Z (initial push)
- 0 spec-orphans across 13 open PRs — **12th consecutive day at zero orphan threshold**. 5 unassigned chaos/* PRs from 06:02 UTC have zero in-progress runs (same pattern as 05-17 and 05-18)
- Open PR backlog fell from 22 → 13 (net 9 PRs merged or closed in 24h) — meaningful throughput despite low *daily* completion rate. The completion-rate-pct metric is a poor proxy for actual merge throughput.
- data_quality stays `infrastructure-only` for the 11th consecutive run — gh auth login OAuth blocker still present
- Per [historical_trend_regression strategy](session-analysis-strategies.json), 22.3 min duration is firmly in the >15-min "100% success" band — consistent with the day's only agent succeeding

## Open Action Items

- [ ] Investigate why conversation transcripts have never been delivered to /tmp/gh-aw/session-data/logs/
- [ ] Consider an "approval bottleneck" severity tier — strict orphan filter misses the dominant failure mode (gates stuck despite agent assignment)
- [ ] Once transcripts arrive, retroactively backfill prompt-quality scoring
- [ ] **Replace daily-completion-rate-pct as the headline metric** — 05-19 shows it can read 2% on a day where the open-PR backlog dropped 41%. Net PR throughput would be more informative.
- [ ] Track whether the 05-18 → 05-19 reversal is a 2-day oscillation or whether 92%+ action_required is the steady state
