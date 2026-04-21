# Shared Alerts — 2026-04-21T12:14Z

## P0 (Critical)
- None currently (Codex 401 #27127 CLOSED Apr 20)

## P1 (High)
- **MCP Gateway startup failures — Codex + MCP CLI servers** (#aw_mcpgw NEW Apr 21): 3 workflows failing at "Start MCP Gateway" — Duplicate Code Detector (serena), AI Moderator (mempalace), Daily Fact (mempalace). Recurring pattern (was #27317, closed Apr 20).
- **15 stale lock files** (#aw_stalocks NEW Apr 21): 15 workflow .lock.yml files outdated — run `make recompile` to fix. 0→15 regression since Apr 20.
- **Design Decision Gate max_turns=5** (#27470 OPEN, Apr 21): ADR write+push requires ≥6 turns; current limit 5 makes ADR generation structurally impossible. Fix: increase to 7 or move mkdir to pre-agent step.
- **AI Moderator chatgpt.com firewall** (#27412 OPEN, Apr 20): Codex v0.121.0 making outbound call to chatgpt.com:443 — blocked by firewall.
- **node: command not found on aw-gpu-runner-T4** (#27534 auto Apr 21): Daily Issues Report. Node.js PATH not available in bash on GPU runners. Recurring.
- **GitHub App rate limit exhaustion** (#27251 OPEN): Co-scheduled at 23:44 UTC. May recur.
- **CODEX_HOME variable collision** (#27512 OPEN, Apr 21): cp same-file error breaks Codex workflows with MCP config.
- **Smoke Claude** (#27030 OPEN): Failing since Apr 14 — ongoing
- **Smoke Copilot** (#27028 OPEN): Ongoing issues group

## P2 (Watch)
- **Daily Community Attribution crash** (#27546 auto Apr 21): Copilot engine terminated mid-run (tool output reading crash). Not systemic — single occurrence.
- **Serena migration codemod missing shared/mcp/serena.md** (#27545 NEW Apr 21): Post-codemod compile fails in external repos.
- **Safe Outputs SEC-004** (#27235 OPEN): 4 handler files need sanitization
- **Performance regressions** (#27280/#27279/#27278 OPEN)
- **dev-hawk github-env** (#26933): High severity zizmor finding
- **PR Triage Agent** (#26778 OPEN): 67% success rate
- **MCP gateway long-running drops** (#23153 OPEN): Session not found after 30-45min
- **Copilot reviewer fan-out** (#27130 OPEN): 6 review workflows per Copilot PR push

## Resolved (Recent)
- CLI updates (#27143) CLOSED Apr 20 ✅ / re-updated via #27484 ✅
- Stale lock files (#27140) resolved ✅ (but re-emerged Apr 21 — #aw_stalocks)
- Codex 401 auth (#27127) CLOSED Apr 20 ✅ (but Codex+MCP CLI Gateway failures persist)
- Daily Fact MCP Gateway (#27317) CLOSED Apr 20 ✅ (but pattern recurred today)

## Ecosystem State
- 198 workflows total (+1 from Apr 20)
- 15 stale lock files ⚠️ (was 0 on Apr 20)
- Schedule success rate: ~80% (MCP Gateway + node PATH failures)
- P0 failures: 0
- P1 failures: systemic MCP Gateway (3 workflows), stale locks (15), node PATH (recurring)
- Overall quality trend: Q:70 (↓-3 from 73), recompile needed

## Orchestrator Summaries
- Workflow Health (Apr 21 12:14Z): Score 70/100. 198 workflows. 15 stale locks. MCP Gateway P1 (codex+CLI servers). node not found GPU runner P1.
- Agent Performance (Apr 21 04:45Z): Q:72 E:68. 25 workflows, 31 runs. DDG max_turns P1. Docs Unbloat 0-output cost drain P2.
- Workflow Health (Apr 20 12:14Z): Score 73/100. 197 workflows. 0 stale locks. node not found on GPU runner P1.
- Agent Performance (Apr 20 04:46Z): Q:73 E:70. 18 workflows, 33 runs. Codex P0 day 3.

Last updated: 2026-04-21T12:14Z by workflow-health-manager
