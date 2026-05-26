# Agent Performance Analyzer — Latest Run

**Timestamp:** 2026-05-26T13:40:00Z  
**Workflow:** Agent Performance Analyzer  
**Run ID:** 26451465997  
**Run URL:** https://github.com/github/gh-aw/actions/runs/26451465997

## Executive Summary

- **Agents analyzed:** 14 active workflow groups (~236 total workflows)
- **Quality score:** 74/100 (↔ plateau — 3rd consecutive flat week)
- **Effectiveness:** 72/100 (↑1)
- **Ecosystem health:** 70/100 (↑7 from May 22)

## Critical Findings

### P0 - Critical
1. **CGO/CJS** (#29669): 90+ days at 0% success rate — DISABLE IMMEDIATELY

### P1 - High Priority
2. **Codex family** (#32446): 12 workflows blocked by OPENAI_API_KEY missing
3. **[aw] Failure reporters**: 20 issues/day, 60%+ duplicates — dedup issue filed
4. **Smoke Antigravity/Pi**: Persistent failures, no root cause investigation

### Positive Signal
- **copilot-swe-agent**: 67% merge rate (↑6pp from 61%), 6 merges May 26 alone
- **Workflow Health**: 70/100 (↑7 from May 22)

## Top Performers

1. copilot-swe-agent (88/100 quality, 84/100 effectiveness)
2. Lint Monster (90/100 quality, 85/100 effectiveness)
3. Contribution Check (86/100 quality, 84/100 effectiveness)
4. Static Analysis (85/100 quality, 83/100 effectiveness)
5. Code Simplifier (84/100 quality, 88/100 effectiveness)

## Trends

- Quality: 74/100 (↔)
- Effectiveness: 72/100 (↑1)
- copilot-swe-agent merge rate: 67% (↑6pp) ✅
- github-actions PR merge rate: 19% (↓11pp) ⚠️
- [aw] failure issues/day: ~20 (↑43%) ⚠️
- Ecosystem health: 70/100 (↑7) ✅

## Issues Created This Run

1. **[aw] Failure reporters deduplication** — add check-before-create logic

## Coordination Notes

### For Campaign Manager
- copilot-swe-agent recovering — 67% merge rate, monitor for sustained improvement
- CGO/CJS must be removed from all campaigns immediately (#29669)
- [aw] failure reporter noise (20 issues/day) polluting campaign health signals

### For Workflow Health Manager
- 20 [aw] failure issues on May 26 — partially noise from dedup gap
- Auto-Triage failed today — watch for recovery May 27
- Smoke Antigravity/Pi root cause investigation needed
