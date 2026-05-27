# Agent Performance Analyzer — Latest Run

**Timestamp:** 2026-05-27T13:51:00Z  
**Workflow:** Agent Performance Analyzer  
**Run ID:** 26515287616  
**Run URL:** https://github.com/github/gh-aw/actions/runs/26515287616

## Executive Summary

- **Agents analyzed:** 14 active workflow groups (~236 total workflows)
- **Quality score:** 74/100 (↔ plateau — 4th consecutive flat week)
- **Effectiveness:** 72/100 (↑1)
- **Ecosystem health:** 90/100 (↑20 from 70/100 last week 🎉)

## Critical Findings

### P0 - Critical
1. **CGO build regression** (#35028): 37% fail rate on push — engineering escalation needed

### P1 - High Priority
2. **Safe-outputs permission blocking** — systemic: PR Sous Chef, Contribution Check, Agentic Maintenance
3. **Copilot/Codex CLI execution failure** — Daily News (5+ days), Daily Issues Report Generator, Daily Fact — new issue filed this run
4. **failure-reporters** 60% duplicate rate — dedup issue from prior run, needs validation

### Positive Signals
- **Ecosystem health: 90/100** (↑20) — biggest single-week improvement recorded
- **copilot-swe-agent**: 67% merge rate (↑6pp), 6 merges May 27 alone
- **Compilation: 100%** (236/236)

## Top Performers

1. Lint Monster (90/100 quality, 85/100 effectiveness)
2. copilot-swe-agent (88/100 quality, 84/100 effectiveness)
3. spec-enforcer/extractor (82/100 quality, 80/100 effectiveness, 100% merge rate)

## Pattern Detection Results

- **Dominant pattern:** Infrastructure-blocking (5/9 agent groups affected)
- **No scope creep detected** across fleet
- **over-creation/duplication:** failure-reporters (60% duplicate rate)

## Issues Created This Run

1. **Copilot/Codex CLI persistent execution failure** — Daily News 5+ days offline

## Coordination Notes

### For Campaign Manager
- copilot-swe-agent strong: 67% merge rate, 6 merges May 27 — campaigns can increase assignment rate
- CGO #35028 — no campaign assignments until resolved
- Failure-reporter noise (20 issues/day, 60% dupe) still polluting health signals

### For Workflow Health Manager
- safe-outputs permission blocking is systemic — root cause investigation needed
- Copilot/Codex CLI failures: new issue filed, 5+ days offline
- 87/236 inactive workflows — deprecation pass recommended
