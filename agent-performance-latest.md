# Agent Performance - 2026-04-17
Run: §24547915508 | Q:77↑1 E:75↑6

## Ecosystem Overview (Apr 16-17, 120 scheduled runs across 82 workflows)
- Overall success rate: 83.3% (↑ from ~76% Apr 16)
- 54/82 (66%) workflows ran with 100% success
- 28/82 workflows had at least one failure

## Top Performers
1. **Agentic Maintenance** (Q:87 E:100) - 9/9 runs, 100% success, most reliable high-freq agent
2. **Issue Monster** (Q:88 E:95) - 20/21 runs, 95% success — 1 fail (#26775 opened)
3. **Bot Detection** (Q:82 E:100) - 3/3 runs, 100%
4. **DeepReport** (Q:85 E:100) - 1/1 run, 100%
5. **Architecture Guardian** (Q:82 E:100) - 1/1 run, 100%
6. **Daily Compiler Quality Check** (Q:80 E:100) - 1/1 run, 100%
7. **Workflow Health Manager** (Q:82 E:100) - 1/1 run, 100%

## Watch List
- **Contribution Check** (Q:65 E:50) - 2/4 runs success, improving but inconsistent
- **Auto-Triage Issues** (Q:55 E:67) - 2/3 runs, intermittent failures (#26364)
- **PR Triage Agent** (Q:55 E:67) - 2/3 runs, 1 fail (#26778)
- **[aw] Failure Investigator (6h)** (Q:55 E:50) - 1/2 runs, 50%

## P0 Persistent Failures (Unchanged)
- **Daily Issues Report Generator** (#26393 OPEN) - node:not-found, 10+ day streak
- **Smoke Gemini** (#26351 OPEN) - Gemini API proxy crash, 14+ day streak

## New Infrastructure Issues
- **Super Linter Report** - EACCES permission error on super-linter.log upload (structural bug)
- **Daily Copilot PR Merged Report** - agent timeout (orphan processes terminated)

## Copilot Tool Upgrade Status
- Copilot v1.0.27 available, #26803 open (was #26367). Active: v1.0.21
- Gemini 0.38.0 available (may fix Smoke Gemini)
- Claude Code 2.1.109 / Codex 0.120.0 available

## Issues Created This Run
- No new issues created (failure patterns already tracked)

## Key Improvements Since Apr 16
- GitHub Remote MCP Auth Test: recovered ✅ (was 100% fail)
- Documentation Unbloat: 100% today (had failed Apr 16)
- Ecosystem success rate improved from ~76% to 83.3%

Last updated: 2026-04-17T04:39Z by agent-performance-manager
