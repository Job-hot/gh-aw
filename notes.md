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

## Run: 2026-05-20

- 92% action_required (46/50) + **8% in-progress** (4 Running Copilot cloud agent) — 0% conclusive successes/failures. First time the in-progress bucket is non-zero in the 13-run window; the agent fleet was actively working at sampling time.
- 5 unique branches; queue is even flatter than 05-17 — top two tied at **14/50 (28%) each**: copilot-opt-refactor-semantic-function-clustering and copilot/duplicate-code-fix-json-wrappers. Top-4 absorb 49/50 (98%).
- Workflow fingerprint: Agentic Commands + Q = **24/50 (48%)** — concentration *relaxed* vs 64% on both 05-18 and 05-19, reverting to the 05-17 level (48%). The two-day stuck-at-64% is broken.
- Burst shape: **5 bursts of exactly 5 fires each** at 06:52:03 / 06:57:58 / 07:33:18 / 07:33:27 / 07:33:40 — purest gate-sweep cadence observed; no large-burst-plus-drip pattern as on 05-12/05-19. Suggests synchronized template firing rather than commit-triggered cascades.
- 11 open PRs (down from 13 on 05-19, 22 on 05-18) — **continued backlog reduction**; 10/11 have Copilot assigned. The single unassigned PR (enhance-sentry, mnkiefer-authored) is human-owned and brand-new (created during sampling), not a candidate for the orphan filter.
- 0 spec-orphans — **13th consecutive day at zero orphan threshold**. The enhance-sentry PR has only 3 active runs (below the ≥5 threshold) and is <2min old; the orphan filter correctly excludes it.
- Sampling window was 06:46:36Z → 07:52:25Z (1h 6m) — narrowest sampling span in the window; corresponds to the period when the 4 Copilot cloud agents were spinning up. No completed agent durations available.
- data_quality stays `infrastructure-only` for the **12th consecutive run** — conversation transcript path unchanged.
- Per [historical_trend_regression strategy](session-analysis-strategies.json), the 4 in-progress agents started in 06:46–07:46Z window are too new (<= 6min runtime at sample) to predict outcomes — they'll likely close in the >5-min "high-success" band by next run.

## Run: 2026-05-21

- **86% action_required** (43/50) — meaningful relief from 92%/98% steady-state of last 3 days; first day with all three outcome types (success/failure/action_required) since 05-18.
- **12% success** (6/50) — second-best day in window after 05-18 (22%). 4 are real Copilot agent runs (1 cloud agent on top branch + 3 "Addressing comment" runs on 3 distinct PRs); 2 are platform gates (Doc Build + Smoke CI) on the same `parser-workflow` branch.
- **2% failure** (1/50) — single CGO failure on `copilot/refactor-oversized-functions-parser-workflow` at the same timestamp as the Smoke CI success on the same branch. Platform-specific CGO compile issue, not a regression of the change.
- **Top branch dominance**: `copilot/fix-duplicate-regex-patterns` absorbs **25/50 sessions (50%)** — highest single-branch share since 05-12's 70%. The branch's lone success was the Copilot cloud agent run (20.3min, success).
- **Workflow concentration broke**: Agentic Commands + Q = **13/50 (26%)** — lowest in the 13-day window (prior min was 48% on 05-17/05-20). Compositional shift: Smoke CI is now the largest single workflow (11/50 = 22%) — points to genuine pre-merge CI activity rather than agentic-loop polling.
- **Copilot agent durations** (all success): 7.55 / 8.83 / 12.55 / 20.27 min — average 12.3min, median 10.7min. All firmly in the >5-min high-success band per [historical_trend_regression strategy](session-analysis-strategies.json). The 20.3-min cloud agent run on `fix-duplicate-regex-patterns` is the only agent on the 50%-dominant branch.
- **Open PR backlog: 7** (down from 11 on 05-20, 13 on 05-19, 22 on 05-18) — **4 net PRs merged or closed in 24h, lowest backlog observed in the window**.
- **0 spec-orphans** — **14th consecutive day at zero orphan threshold**. Only 2 in-progress runs system-wide (both on `main`: Agentic Maintenance + this workflow). All 7 open PRs have a Copilot assignee.
- **Experimental strategy**: "Inverse Gate-Count to Conclusiveness" — perfect inverse correlation across the 5 active branches today. 3-session branch = 100% conclusive; 25-session branch = 4% conclusive. Strong signal that high gate count indicates *waiting on agent*, not *waiting on CI*. Added to [strategies.json](session-analysis-strategies.json) as High effectiveness.
- **Sampling window**: 06:02:11Z → 06:59:51Z (58 min — narrower than the 1h6m on 05-20).
- **Bursts**: 5 fires at 06:39:29Z (sweep coincident with cloud agent start, same sweep-after-success pattern as 05-16/05-17/05-19) + 4 fires at 06:59:51Z (final-minute sweep on top branch) + 4 fires at 06:15:19Z.
- `data_quality` stays `infrastructure-only` for the **13th consecutive run** — OAuth blocker still in effect; no conversation transcripts available.

## Run: 2026-05-22

- **98% action_required (49/50) + 2% success (1/50)** — identical conclusiveness shape to 05-19; reverses the brief 05-21 relief (86%/12%/2%). The single success is the Copilot cloud agent run on `copilot/refactor-semantic-function-clustering-please-work` (18.22 min, well in the >15-min "high-success" band per [historical_trend_regression strategy](session-analysis-strategies.json)).
- **Sampling window: 06:02:08Z → 06:21:14Z = 19 minutes** — narrowest sampling window observed in the 14-day series (previous narrowest was 58 min on 05-21, 66 min on 05-20). The window is shrinking month-over-month — possible the upstream `copilot-session-data-fetch` module is becoming faster or the system is more idle outside the sweep window.
- **6 unique branches** absorb the queue; top branch (`refactor-semantic-function-clustering-please-work`) holds only **16/50 (32%)** — meaningfully flatter than 05-21 (50%) and 05-12 (70%). Branch diversity is at its 14-day high.
- **Workflow fingerprint**: Agentic Commands + Q = **26/50 (52%)** — climbs back from 05-21's 26% (the 14-day low) toward the 48% baseline of 05-17/05-20. The 05-21 dip was anomalous, not a regime shift. CGO + Smoke CI + Doc Build = 19/50 (38%) — heavy CI activity, similar to 05-21.
- **Bursts**: 8 fires at 06:11:32Z (largest single burst since 05-12's 14) coincide with the creation of PR #33954 on `add-request-review-mode`. Next-largest are 5 fires at 06:02:11Z (5 chaos/* PRs created simultaneously) and 5 at 06:16:36 / 06:20:45. Same "push + drip" pattern as 05-12.
- **Open PR backlog: 12** (up from 7 on 05-21, +5 net) — the increase is entirely the 5 unassigned `chaos/*` PRs created within seconds at 06:00-06:01Z (paranoid-reviewer-r48-amend, code-archaeologist-r48-two-commits, selective-stager-r48-staged-subset, minor-renamer-r48-rename, line-ending-normalizer-r48). This matches the chaos-PR pattern seen on 05-17 and 05-19 but at twice the volume.
- **0 spec-orphans** — **15th consecutive day at zero orphan threshold**. The 5 chaos/* PRs are unassigned but have **zero active in-progress runs** in the 6h lookback (the only in-progress runs are 3 on `main`: Daily Workflow Updater, Failure Investigator, this workflow). Filter correctly excludes them. However, the chaos PRs sit at the 2h-warning threshold edge — worth tracking whether they ever attract gate activity.
- **0 failure conclusions** — 2nd consecutive day after the 1 CGO failure on 05-21. Failure rate remains very low across the window.
- **data_quality** stays `infrastructure-only` for the **14th consecutive run** — `/tmp/gh-aw/session-data/logs/` is empty again. OAuth blocker has been continuous since 05-06.
- Per [inverse-gate-count-to-conclusiveness strategy](session-analysis-strategies.json), today the 16-session top branch has 6.25% conclusive rate (1/16) — consistent with the model (>15-gate branches are waiting on agent action, not on green CI). The cloud agent run at 06:03:04Z (success) is the only conclusive event, and it sits on the dominant branch.

## Open Action Items

- [ ] Investigate why conversation transcripts have never been delivered to /tmp/gh-aw/session-data/logs/ — 14 consecutive runs blocked
- [ ] Consider an "approval bottleneck" severity tier — strict orphan filter misses the dominant failure mode (gates stuck despite agent assignment)
- [ ] Once transcripts arrive, retroactively backfill prompt-quality scoring
- [ ] **Replace daily-completion-rate-pct as the headline metric** — 05-19 and 05-22 both read 2% on days that are operationally normal. Net PR throughput would be more informative.
- [ ] Track whether the 05-21 → 05-22 reversal (86% → 98% action_required) is a 2-day oscillation or whether the 05-21 relief was a one-off
- [ ] Watch the 5 chaos/* PRs from 06:00-06:01Z on 05-22 — they sit at the edge of the 2h-warning band but lack active gates; if gates start firing on them within the next sampling window, they'd become true orphans (first time in 15 days)
