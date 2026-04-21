# Shared Alerts — 2026-04-21T04:45Z

## P0 (Critical)
- **Codex engine 401 auth** (#27127, OPEN, **day 4**): OPENAI_API_KEY missing/expired. Affects all Codex workflows: AI Moderator, Daily Observability Report, Duplicate Code Detector, Schema Feature Coverage. Needs admin credential rotation. **Escalation recommended — 4 days with no resolution.**

## P1 (High)
- **Design Decision Gate max_turns=5** (#27470 OPEN, NEW Apr 21): ADR write+push path requires ≥6 turns; current limit 5 makes ADR generation structurally impossible. Fix: increase to 7 or move mkdir to pre-agent step.
- **AI Moderator chatgpt.com firewall** (#27412 OPEN, NEW Apr 20): Codex v0.121.0 making outbound call to chatgpt.com:443 — blocked by firewall. Separate from 401 auth issue.
- **node: command not found on aw-gpu-runner-T4**: Recurring Daily News + Daily Issues Report. Node.js PATH not available in bash on GPU runners.
- **GitHub App rate limit exhaustion** (#27251 OPEN): Co-scheduled at 23:44 UTC. May recur if not staggered.
- **Smoke Claude** (#27030 OPEN): Failing since Apr 14 — ongoing
- **Smoke Copilot** (#27028 OPEN): Ongoing issues group

## P2 (Watch)
- **Documentation Unbloat cost drain**: $2.46/run, 58 turns, 0 safe outputs. Needs outcome gate.
- **Agent Persona Explorer**: 95% data-gathering in agentic turns, needs pre-agent steps.
- **Safe Outputs SEC-004** (#27235 OPEN): 4 handler files need sanitization
- **Performance regressions** (#27280/#27279/#27278 OPEN): CompileComplexWorkflow +29%, CompileSimpleWorkflow +39%, Validation +96%
- **dev-hawk github-env** (#26933): High severity zizmor finding
- **PR Triage Agent** (#26778 OPEN): 67% success rate
- **MCP gateway long-running drops** (#23153 OPEN): Session not found after 30-45min
- **Copilot reviewer fan-out** (#27130 OPEN): 6 review workflows per Copilot PR push

## Resolved (Recent)
- CLI updates (#27143) CLOSED Apr 20 ✅ / re-updated today via #27484 ✅
- Stale lock files (#27140) resolved ✅

## Ecosystem State
- 197 workflows total (stable)
- 0 stale lock files ✅
- Schedule success rate: ~84% (Codex P0 drag)
- P0 failures: 1 (Codex 401 auth — day 4, escalation needed)
- Overall quality trend: Q:72 (↓-1), E:68 (↓-2)

## Orchestrator Summaries
- Agent Performance (Apr 21 04:45Z): Q:72 E:68. 25 workflows, 31 runs. Codex P0 day 4. DDG max_turns P1 new. Docs Unbloat 0-output cost drain P2.
- Workflow Health (Apr 20 12:14Z): Score 73/100. 197 workflows. 0 stale locks. Codex P0 (day 3). node not found on GPU runner P1.
- Agent Performance (Apr 20 04:46Z): Q:73 E:70. 18 workflows, 33 runs. Codex P0 day 3.

Last updated: 2026-04-21T04:45Z by agent-performance-manager
