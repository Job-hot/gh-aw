# Agent Performance — Latest Run

**Timestamp:** 2026-06-18T13:47:00Z | **Run:** [§27763811011](https://github.com/github/gh-aw/actions/runs/27763811011)

## Summary: 57/100 Quality (+2) | 55/100 Effectiveness (+2) | 67/100 Health (→) | AIC crisis Day 12

## Top Performers
1. copilot-swe-agent (Q:80, E:87) — 41 PRs today, 76% close-merge rate. Same-day fix #40035 for critical safe-outputs bug.
2. Bot Detection (Q:82, E:91) — 100% success today
3. Agentic Maintenance (Q:82, E:88) — running today
4. Auto-Triage Issues (Q:81, E:84) — success today
5. Content Moderation (Q:76, E:80) — running today
6. PR Sous Chef — FULLY RECOVERED (ongoing)
7. Avenger — Healthy (consecutive successes)

## Recoveries This Run ✅
- Daily Documentation Updater: FULLY RECOVERED
- Daily Workflow Updater: FULLY RECOVERED
- Instructions Janitor: RECOVERED
- Glossary Maintainer: RECOVERED

## Underperformers
- Code Simplifier (Q:10, E:5) P1 Day 12: api-proxy cap + HTTP 429. Re-filed as #39968. DO NOT RE-FILE.
- Tool Denial Cluster (7+ workflows, systemic): Q:20, E:15. Systemic issue filed. DO NOT RE-FILE.
- Daily Model Inventory (Q:35, E:25) Day 9: session.idle. DO NOT RE-FILE.
- Daily News (Q:30, E:20) Day 6+: GPU runner tool path. #40033 PR open. DO NOT RE-FILE.
- Daily Compiler Quality Check (Q:35, E:20) Day 2: gpt-5-mini config error. DO NOT RE-FILE.
- Smoke cluster (Q:25, E:15): 5 engines, 2 batches today (10 issues total). DO NOT RE-FILE.
- Daily BYOK Ollama Test Day 10: transient_bad_request. DO NOT RE-FILE.
- Daily Safe Output Integrator Day 10: tool denial. DO NOT RE-FILE.

## New Patterns Detected (Jun 18)
- **Duplicate Issue Re-Filing**: Daily News #40023 re-filed despite #39758 open; Code Simplifier #39968 re-filed despite existing tracking. Failure detection lacks dedup. Issue filed #aw_dedup.
- **Safe outputs target:triggering bug**: Filed #40017, FIXED same-day #40035 — excellent responsiveness.
- **Smoke cluster 2nd batch**: #40042-#40047 join morning batch #39986-#39994 (10 issues total, same regression).

## PR Merge Rates (Jun 18)
- copilot-swe-agent: 76% (25/33 closed PRs today) ↑ from 67% yesterday
- spec-enforcer/spec-extractor: 100% (2/2 today)
- docs/glossary workflows: 100% (2/2 today)
- github-actions workflows: ~60% (some pending/closed)

## Issues Filed This Run
- 1 systemic issue: Failure Detection Deduplication (#aw_dedup)

## Do Not Re-File (additions)
#40023 Daily News (re-file of #39758), #39968 Code Simplifier (re-file of #39199 series), #40042-#40047 Smoke cluster 2nd batch

