# Agent Performance Analyzer — Latest Run

**Timestamp:** 2026-05-28T14:05:00Z  
**Workflow:** Agent Performance Analyzer  
**Run ID:** 26579184217  
**Run URL:** https://github.com/github/gh-aw/actions/runs/26579184217

## Executive Summary

- **Agents analyzed:** 16 active workflow groups (~236 total workflows)
- **Quality score:** 74/100 (↔ plateau — 5th consecutive flat week)
- **Effectiveness:** 72/100 (↔ stable)
- **Ecosystem health:** 82/100 (↓8 from 90 — more failures today vs yesterday)

## Critical Findings

### P0 - Critical
1. **safe_outputs add_comment validation** (#35351): 3+ workflows fully blocked (PR Sous Chef, Contribution Check, Sub-Issue Closer) — ongoing

### P1 - High Priority
2. **Copilot CLI execution failure** (#35388): 0% success for 5+ consecutive days (Daily News, Deep Research Agent, Issues Report Generator) — platform-level
3. **failure-reporters 60% duplicate rate**: ~20 issues/day, 60% duplicates — dedup gate needed
4. **LintMonster backlog overwhelm** (#35368 epic: 2417 issues, #35370 latest failure): resource/timeout failure pattern

### P2 - Medium Priority
5. **Silent-skip cluster**: Q, Deployment Incident Monitor, CJS, Label Closed PRs — 0-33% success with no failure logs
6. **Daily Safe Output Tool Optimizer** (#35316): 115 turns / 14.9M tokens — runaway reasoning loop

### Positive Signals
- **copilot-swe-agent**: Active (remove-yield-feature PR in progress), 67% merge rate
- **Compilation: 100%** (236/236)
- **Agentic Commands**: 80% pass rate — most stable PR-triggered workflow
- **CGO**: Stabilizing after #35028 regression

## Top Performers

1. copilot-swe-agent (88/100 quality, 84/100 effectiveness, 67% merge rate)
2. spec-enforcer/extractor (82/100 quality, 80/100 effectiveness, 100% merge rate)
3. Agentic Commands (80% pass rate, no failures)
4. Content Moderation (75% pass, clean behavioral profile)

## Pattern Detection Results (via pattern-detector)

- **Under-creation** (blocked): safe-outputs cluster, Copilot CLI cluster, Agentic Maintenance
- **Over-creation + repetition**: failure-reporters (60% dupe rate)
- **Over-creation + inconsistency**: LintMonster, Daily Safe Output Tool Optimizer
- **Inconsistency**: CGO, CJS, PR Description Updater
- **Healthy (no patterns)**: copilot-swe-agent, spec-enforcer, Agentic Commands, Content Moderation

## Issues Created This Run

1. **Silent-skip cluster audit** — Q, Deployment Monitor, CJS, Label Closed PRs (new issue filed)

## Coordination Notes

### For Campaign Manager
- copilot-swe-agent still strong — active on remove-yield-feature, campaigns can continue high assignment rate
- Block all Copilot CLI workflow campaigns until #35388 resolved
- safe_outputs P0 (#35351) blocks 3+ workflows — do not assign campaigns to affected agents

### For Workflow Health Manager
- Silent-skip cluster is a new finding — 4 workflows with 0-33% success and zero failure logs need trigger audit
- failure-reporters deduplication remains unresolved — continues to pollute health signals
- Daily Safe Output Tool Optimizer (#35316) token cap: add early-exit on rate-limit detection
