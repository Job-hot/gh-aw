# Session Analysis Notes

## Persistent Observations

- **Data quality**: `infrastructure-only` / `metadata-only` for ~3 weeks. Conversation transcripts have never been reliably available — `copilot-session-data-fetch` returns OAuth errors. This is the longest-running unresolved workflow risk.
- **Recovery/regression oscillation**: completion rate bounces between near-0% and 40-50% rather than degrading steadily. Recoveries (2026-05-23: 44%, 2026-05-26: 46%) are isolated days, not sustained trends.
- **Branch concentration drives outcomes**: 80-90% of daily sessions are on 2-4 branches. Concentration on agent-assigned PR branches → higher completion. Concentration on agent-less branches → action_required dominance.
- **Orphan rate**: 0% for 7+ consecutive days. The 40% baseline appears outdated; gating is no longer accumulating on un-owned branches.
- **Action_required share**: the volatile metric. Ranges 8% → 98% day-to-day. Tracks inversely with completion.

## Run: 2026-05-26 (latest)

- **46% success** (23/50) — new 7-day high; beats 2026-05-23's 44% outlier
- 82% of sessions on 2 PR branches (copilot/bugfix-create-pull-request-patch: 22, copilot/fix-patch-application-issue: 19), both Copilot-assigned
- Avg 6.35m, median 3.85m, max 18.7m — 0 sessions ≥20m
- Bimodal distribution: 23 sessions <30s + 24 sessions >5m
- 0 spec-orphans; 22 action_required still concentrated on the same PR branches (intermittent permission gating, not systemic)
- Conversation logs unavailable (3rd consecutive day)
- 100%-success workflows: PR-anchored gates (`Addressing comment on PR #34874/#34876`), quality gates (`Test Quality Sentinel`, `PR Code Quality Reviewer`, `Matt Pocock Skills Reviewer`, `Design Decision Gate`)
- Action_required-heavy workflows: short-named CI sweep gates (`Q`, `CJS`, `CGO`, `Smoke CI`)
