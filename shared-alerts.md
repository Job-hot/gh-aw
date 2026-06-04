# Shared Alerts — 2026-06-04T06:09Z

## P0 (Critical) 🚨
_None active_

## P1 (High) 🚨
- **Daily Firewall Logs Collector** (NEW - Jun 4): 6 consecutive failures, token budget exhaustion. Issue filed today. DO NOT re-file.
- **Daily BYOK Ollama Test** (NEW - Jun 4): 5 consecutive failures, agent execution failure. Issue filed today. DO NOT re-file.
- **CGO unit tests** (#35028 OPEN): Persistent 100% failure since May 27. DO NOT re-file.
- **Code Simplifier** (#36829 OPEN): Failure Jun 4. DO NOT re-file.
- **Workflow Skill Extractor** (#36837 OPEN): Failure Jun 4. DO NOT re-file.

## P2 (Watch) ⚠️
- **PR Triage Agent**: 2 failures on Jun 4. Inconsistent — monitor for pattern.
- **Linter Miner**: Intermittent (~2/7 days). Monitor.
- **Daily Sentrux Report**: 1 failure Jun 4 (was reliable). Likely transient.
- **Daily Model Inventory Checker**: 1 failure Jun 4. Likely transient.
- **chaos-test PR stall**: 10+ open PRs pending merge — still unresolved
- **Safe Outputs Conformance SEC-005** (#36591 OPEN): update_activation_comment.cjs allowlist gap

## Resolved ✅ (since Jun 3)
- **Sergo** (#36574 CLOSED): back to success today
- **Auto-Triage Issues**: passing

## Systemic Notes
- Token budget exhaustion is a recurring theme (Firewall, CGO, potentially others)
- Compilation health excellent: 240/240 lock files
