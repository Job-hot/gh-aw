# Workflow Health — 2026-04-21T12:14Z

Score: 70/100 (↓-3 from 73 Apr 20). 198 workflows. Run: §24721544063

## KEY FINDINGS

### Compilation Status
- 198/198 lock files present ✅ (up 1 from 197)
- **15 stale lock files** ⚠️ (up from 0 yesterday — needs `make recompile`)
  - ai-moderator, artifacts-summary, daily-compiler-quality, daily-otel-instrumentation-advisor
  - daily-safe-output-integrator, deep-report, developer-docs-consolidator, functional-pragmatist
  - jsweep, refactoring-cadence, smoke-agent-all-none, smoke-claude, smoke-crush, spec-enforcer, step-name-alignment

### P0 Issues
- None today (Codex 401 #27127 CLOSED Apr 20)

### P1 Issues Today
- **MCP CLI → MCP Gateway startup failure** (NEW): 3 Codex+MCP CLI server workflows failing
  - Duplicate Code Detector (#27557): codex + serena → "Start MCP Gateway" failed
  - AI Moderator: codex + mempalace → "Start MCP Gateway" failed
  - Daily Fact About gh-aw (#27550): codex + mempalace → failed
  - Pattern: ALL use MCP CLI servers (`serena`, `mempalace`) + Codex engine
  - Previously tracked as #27317 (closed Apr 20 as not_planned) — recurring
- **15 stale lock files** (#new): workflows modified since last compile
  - Risk: compiled behavior may not match source markdown
- **node: command not found on aw-gpu-runner-T4** (#27534): Daily Issues Report (recurring)
- **Design Decision Gate max_turns=5** (#27470 OPEN): structurally impossible ADR generation
- **GitHub App rate limit** (#27251 OPEN): co-scheduled at 23:44 UTC

### P2 Issues
- Daily Community Attribution Updater crash (#27546): Copilot engine terminated mid-run
- Safe Outputs SEC-004 conformance (#27235 OPEN): 4 handler files
- Performance regressions (#27280/#27279/#27278 OPEN)
- Serena migration codemod missing shared/mcp/serena.md (#27545 NEW)
- Copilot reviewer fan-out (#27130 OPEN)
- MCP gateway long-running drops (#23153 OPEN)

### Resolved (Since Last Run)
- Codex 401 auth (#27127) CLOSED Apr 20 ✅ — but Codex+MCP CLI failures suggest partial resolution
- Daily Fact MCP Gateway (#27317) CLOSED Apr 20 ✅ — but pattern recurred today

### Schedule Success Rate (Today, Apr 21 by 12:14Z)
- ~80% success (5 confirmed failures from 30 scheduled runs observed)
- Healthy: most event-triggered and daily workflows ✅

## Open Issues (workflow-health)
- #27557 Duplicate Code Detector failed (P1, auto-created today)
- #27550 Daily Fact About gh-aw failed (P1, auto-created today)
- #27546 Daily Community Attribution failed (P2, auto-created today)
- #27534 Daily Issues Report failed (P1 GPU/node, auto-created today)
- #27512 CODEX_HOME variable collision (P1, Apr 21)
- #27470 Design Decision Gate max_turns=5 (P1)
- #27412 AI Moderator chatgpt.com firewall (P1)
- #27251 Rate limit exhaustion co-scheduled (P1)
- #27235 Safe Outputs SEC-004 (P2)
- #27030 Smoke Claude (P1)
- #27028 Smoke Copilot (P1)
- #23153 MCP gateway session drops (P2)

## Engine/Tool Status
- Copilot: mostly ✅ (occasional crashes - #27546)
- Claude: mostly ✅ (Smoke Claude #27030 ongoing)
- Codex: ⚠️ MCP CLI server + Gateway startup failures (3 workflows today)
- AWF/MCP Gateway: ⚠️ startup fails when MCP CLI servers (serena/mempalace) used with Codex

## Actions This Run
- Stale locks P1 tracker created
- MCP Gateway Codex+CLI server failure tracker created
- Memory files updated

Last updated: 2026-04-21T12:14Z by workflow-health-manager
