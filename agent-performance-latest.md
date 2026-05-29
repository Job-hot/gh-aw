# Agent Performance Analyzer — Latest Run

**Timestamp:** 2026-05-29T13:46:36Z  
**Workflow:** Agent Performance Analyzer  
**Run ID:** 26640736780  
**Run URL:** https://github.com/github/gh-aw/actions/runs/26640736780

## Executive Summary

- **Agents analyzed:** 14 active workflow groups (~236 total workflows)
- **Quality score:** 72/100 (↓2 from 74 last week)
- **Effectiveness:** 68/100 (↓4 from 72)
- **Ecosystem health:** 76/100 (↓6 from 82)

## Critical Findings

### P0 - Critical
1. **safe_outputs add_comment validation** (#35351): ongoing — PR Sous Chef, Contribution Check, Sub-Issue Closer blocked

### P1 - High Priority
2. **Agentic Commands regression**: 80%→25% success (was most stable) — likely CJS/Step Name Alignment related
3. **Content Moderation regression**: 75%→22% success — new this week
4. **Copilot CLI execution failure** (#35388): 0% success 5+ days
5. **CJS shard 4 CI failure** (filed 2026-05-29): every push to main fails
6. **Step Name Alignment** (filed 2026-05-29): 80% failure rate since May 20

### P2 - Medium Priority
7. **failure-reporters 60% duplicate rate**: chronic, 3+ weeks unresolved
8. **LintMonster backlog overwhelm** (#35368 epic): 2218+ findings, resource timeouts
9. **Daily Safe Output Tool Optimizer** (#35316): 115 turns / 14.9M tokens runaway

## Top Performers

1. spec-enforcer (85/100 quality, 88/100 effectiveness, 100% merge rate)
2. copilot-swe-agent (82/100 quality, 80/100 effectiveness, 67% merge rate)
3. Copilot cloud agent (80/100 quality, 80/100 effectiveness, 100% success — early signal)

## Pattern Detection Results

- **under-creation**: Q, AI Moderator, Smoke CI, LintMonster, Copilot CLI workflows, PR Sous Chef
- **inconsistency**: Content Moderation, Agentic Commands, CGO, LintMonster, Daily Safe Output Tool Optimizer
- **over-creation + repetition**: failure-reporters
- **scope-creep**: Daily Safe Output Tool Optimizer
- **healthy**: copilot-swe-agent, spec-enforcer, Copilot cloud agent

## Issues Created This Run

None filed — regressions root-caused to existing P1 issues tracked by Workflow Health Manager.

## Coordination Notes

### For Campaign Manager
- Block Copilot CLI workflow campaigns (#35388)
- Block PR Sous Chef campaigns (#35351)
- Agentic Commands and Content Moderation at risk — avoid high-stakes campaign assignment until regressions resolved

### For Workflow Health Manager
- NEW: Agentic Commands regression (80%→25%) likely correlated with CJS/Step Name Alignment P1
- NEW: Content Moderation regression (75%→22%) needs log review
- failure-reporters dedup gate: 3rd consecutive escalation — needs owner assignment
