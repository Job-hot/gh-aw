# Agent Performance Analyzer — Latest Run

**Timestamp:** 2026-05-30T13:01:31Z  
**Workflow:** Agent Performance Analyzer  
**Run ID:** 26684413824  
**Run URL:** https://github.com/github/gh-aw/actions/runs/26684413824

## Executive Summary

- **Agents analyzed:** 14 active workflow groups (~236 total workflows)
- **Quality score:** 71/100 (↓1 from 72 yesterday)
- **Effectiveness:** 66/100 (↓2 from 68 yesterday)
- **Ecosystem health:** 74/100 (↓2 from 76 — smoke + CGO worsening)

## Critical Findings

### P0 - Critical
1. **safe_outputs add_comment validation** (#35351): partial fix merged (PR #35901 — cross-repo support for safe-output handlers, merged 2026-05-30T12:15Z). Verify if core `item_number` validation is also resolved.

### P1 - High Priority
2. **Smoke tests** (#35829, #35832, #35856, #35864, #35866+): ALL variants 100% failing — systemic infrastructure failure
3. **Copilot CLI engine failure** (#35388): 0% success since May 28 — engine step broken
4. **CGO CI broken**: unit tests + custom linters failing (run 26675666376)
5. **Agentic Commands regression**: 80%→25% success (CJS correlation)
6. **Step Name Alignment** (filed 2026-05-29): 80% failure rate since May 20
7. **CJS shard 4 CI** (filed 2026-05-29): CI blocker on every push to main

### P2 - Medium Priority
8. **failure-reporters 60% duplicate rate**: chronic, 3+ weeks unresolved (~12 duplicates/day)
9. **LintMonster backlog overwhelm** (#35368 epic): 2218+ findings, resource timeouts
10. **Daily Safe Output Tool Optimizer** (#35316): 14.9M tokens / 115 turns runaway

## Top Performers

1. copilot-swe-agent (84/100 quality, 82/100 effectiveness, 94% merge rate, 36 merges in 7d)
2. spec-enforcer (85/100 quality, 88/100 effectiveness, 100% merge rate)
3. docs-updater (78/100 quality, 72/100 effectiveness, 70% merge rate, consistent cadence)

## Pattern Detection Results

- **blocked (critical)**: smoke-tests, copilot-cli-workflows, pr-sous-chef
- **scope-creep (critical)**: daily-safe-output-optimizer
- **degraded (critical)**: step-name-alignment
- **inconsistency (critical)**: agentic-commands
- **over-creation (critical)**: chaos-test (0 merges, 5 open PRs), lint-monster (2218 backlog)
- **repetition (critical)**: failure-reporters (60% duplicate rate)
- **over-creation (warn)**: copilot-swe-agent (high volume but high quality)
- **healthy**: spec-enforcer, docs-updater, ab-advisor, deep-report

## Issues Created This Run

None — all critical issues already tracked. See Do Not Re-File list in shared-alerts.md.

## Coordination Notes

### For Campaign Manager
- Block Copilot CLI workflow campaigns (#35388 — ongoing)
- Block PR Sous Chef campaigns (#35351 — partial fix in PR #35901, verify resolution)
- Chaos-test: 5 open unmerged PRs — do not assign new campaigns until reviewed
- Agentic Commands + Step Name Alignment: avoid high-stakes work until resolved

### For Workflow Health Manager
- NEW: PR #35901 (safe output handlers) merged today — needs verification if #35351 fully resolved
- Smoke tests worsening: now 6+ issues tracked, systemic infrastructure fault
- CGO CI: new unit test failures in addition to ongoing issues
- failure-reporters dedup gate: 4th consecutive escalation — needs immediate owner assignment
