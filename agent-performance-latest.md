# Agent Performance - 2026-04-21
Run: §24704272587 | Q:72↓1 E:68↓2

## Ecosystem Overview (Apr 19-21)
- Overall quality: 72/100 (stable), effectiveness: 68/100 (↓-2)
- 25 unique workflows observed, 31 runs (18 completed)
- Codex P0 day 4 (#27127 OPEN)
- Assessment flags today: resource_heavy(3), partially_reducible(4), poor_agentic_control(1), overkill(2), model_downgrade(1)

## Top Performers
1. **[aw] Failure Investigator (6h)** (Q:90 E:88) - #27469 (DDG max_turns RCA) + #27412 (AI Moderator chatgpt.com) — excellent structured reports
2. **Smoke CI** (Q:87 E:90) - 4 runs, reliable, no flags
3. **Test Quality Sentinel** (Q:82 E:85) - 4 turns, clean, no flags
4. **CLI Version Checker** (Q:78 E:80) - Updated 3 CLI tools (#27484) ✅; flagged resource_heavy (26T, $0.82)
5. **Design Decision Gate** (Q:78 E:78) - Today success (4T, $0.37); structural max_turns=5 issue for ADR path (#27470)

## Watch / Needs Improvement
- **Documentation Unbloat** (Q:48 E:52) - 58T, $2.46, 0 safe outputs. Worst cost-to-output ratio.
- **Agent Persona Explorer** (Q:52 E:58) - 42T, 2.2M tokens, 95% data-gathering, 1 safe output
- **AI Moderator** (Q:10 E:5) - Codex P0 day 4 + chatgpt.com firewall block (2 failure modes now)
- **GitHub Remote MCP Auth Test** (Q:40 E:0) - persistent auth failure

## New Findings Today
1. Design Decision Gate max_turns=5 structurally insufficient for ADR path (needs 6+) → #27470 NEW
2. AI Moderator: chatgpt.com blocked by firewall, separate from 401 auth → #27412
3. Documentation Unbloat: $2.46/run with 0 tangible outputs — needs outcome gate
4. CLI Version Checker created clean PR #27484 with 3 tool updates ✅

## P0/P1 Active
- P0: Codex 401 auth (#27127, day 4) — escalation needed
- P1: Design Decision Gate max_turns=5 → #27470
- P1: AI Moderator chatgpt.com firewall → #27412
- P1: node not found on GPU runner (ongoing from Apr 20)
- P1: Rate limit exhaustion 23:44 UTC (ongoing)
- P1: Smoke Claude #27030, Smoke Copilot #27028 (ongoing)

## Issues/Actions This Run
- Discussion created (performance report)
- No new improvement issues (all tracked)

Last updated: 2026-04-21T04:45Z by agent-performance-manager
