# Session Analysis Notes

## Recurring Observation: data_quality = "infrastructure-only"

Across 7 consecutive runs (2026-05-06 → 2026-05-13), conversation transcripts have NEVER been available from the `copilot-session-data-fetch` shared module. This is a known, persistent limitation rather than a regression. The workflow is effectively a CI gate-sweep / orphan-detection monitor, not a behavioral analyzer.

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

## Open Action Items

- [ ] Investigate why conversation transcripts have never been delivered to /tmp/gh-aw/session-data/logs/
- [ ] Consider an "approval bottleneck" severity tier — strict orphan filter misses the dominant failure mode (gates stuck despite agent assignment)
- [ ] Once transcripts arrive, retroactively backfill prompt-quality scoring
