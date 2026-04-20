# Agent Performance - 2026-04-20
Run: §24648853175 | Q:73↓3 E:70↓5

## Ecosystem Overview (Apr 19-20)
- Overall success rate: 67% obs. / 87% excl. AI Moderator P0 drag
- 18 unique workflows, 33 run observations
- 11 safe output items produced
- 33 agentic assessment flags: 11 resource_heavy, 10 model_downgrade, 9 partially_reducible, 4 overkill, 3 poor_agentic_control

## Top Performers
1. **[aw] Failure Investigator (6h)** (Q:88 E:88) - RCA on rate limit exhaustion from co-scheduled workflows; created issue + updated #27128
2. **Smoke CI** (Q:85 E:90) - 2/2 success, minimal overhead
3. **Design Decision Gate 🏗️** (Q:83 E:78) - 3/3 success, efficient 2.7 avg turns
4. **Issue Monster** (Q:82 E:88) - 3/3 success, 6 safe output items, consistent
5. **Test Quality Sentinel** (Q:80 E:85) - 3/3 success, 2 safe outputs

## Watch / Needs Improvement
- **AI Moderator** (Q:10 E:0) - 7/7 failures today (ongoing P0 Codex 401 #27127), ~9 min wasted per run
- **GitHub Remote MCP Auth Test** (Q:40 E:0) - persistent auth failure
- **Documentation Unbloat** (Q:65 E:60) - 58 turns, 4.88M tokens, $2.40, 0 safe outputs
- **Daily Repository Chronicle** (Q:65 E:72) - 66 turns, 96% data-gathering, poor_agentic_control
- **Agent Persona Explorer** (Q:70 E:75) - recovered ✅ but resource_heavy + 93% data-gathering

## P0 Active
- **Codex 401 auth** (#27127, OPEN): Blocking AI Moderator (7 daily failures)

## Key Issues/Recommendations
- Ecosystem #1 problem: data-gathering in agent turns (Daily Chronicle 96%, Persona Explorer 93%, Daily PR Report 81%)
- Model downgrades needed: AI Moderator + PR Triage → gpt-4.1-mini
- Schema Consistency Checker + jsweep: silent startup failures (conclusion="")
- Stagger 23:44 UTC cron schedules (rate limit exhaustion Apr 19)

## Issues Created This Run
- Discussion created (performance report)
- No new improvement issues (all patterns already tracked)

Last updated: 2026-04-20T04:46Z by agent-performance-manager
