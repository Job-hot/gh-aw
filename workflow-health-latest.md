# Workflow Health — 2026-06-06T05:56Z

Score: 74/100 (↓4 from 78)
Workflows: 241 | Lock files: 241/241 (100% ✅) | Run: §27054065038

## KEY FINDINGS

### Status (June 6)
- **Compilation:** 241/241 workflows have lock files (100% ✅) — 1 new workflow added
- **CJS typecheck** (#36410 OPEN): 100% failure since Jun 1 — DO NOT RE-FILE
- **CGO unit tests** (#35028 OPEN): 100% failure since May 27 — DO NOT RE-FILE
- **Daily Documentation Healer**: Still failing Jun 6 after #37039 CLOSED. NEW issue filed today.
- **Daily Model Inventory Checker**: Still failing Jun 6 after #37039 CLOSED. Covered by today's issue.
- **Daily Sentrux Report**: 3rd+ consecutive failure, #37019 CLOSED. Covered by today's issue.
- **Daily BYOK Ollama Test**: Still failing, #37211 OPEN — DO NOT RE-FILE
- **Daily Firewall Logs Collector**: Still failing Jun 6 (token budget, filed Jun 4) — DO NOT RE-FILE
- **Safe Output Health Monitor**: Token budget failure, #37264 OPEN — DO NOT RE-FILE
- **PR Sous Chef**: #37216 OPEN — DO NOT RE-FILE

### Critical Issues (P1) 🚨
- **CJS typecheck**: #36410 OPEN. DO NOT RE-FILE.
- **CGO unit tests**: #35028 OPEN. DO NOT RE-FILE.
- **Doc Healer + Model Inventory + Sentrux**: NEW issue filed today. DO NOT RE-FILE.
- **Daily BYOK Ollama Test**: #37211 OPEN. DO NOT RE-FILE.
- **Daily Firewall Logs Collector**: filed Jun 4. DO NOT RE-FILE.

### P2 Issues ⚠️
- **Safe Output Health Monitor**: #37264 OPEN. DO NOT RE-FILE.
- **AI Moderator**: Still blocked (0% success). Monitor.
- **PR Sous Chef**: #37216 OPEN. DO NOT RE-FILE.
- **Code Simplifier**: recurring (#36829). Monitor.
- **jsweep/Daily Compiler Quality Check**: token budget. Monitor.

### New 1x Failures (Jun 6, possibly transient)
- Copilot CLI Deep Research Agent (1x) — aw-failure-investigator may file
- Daily Agent of the Day Blog Writer (1x) — aw-failure-investigator may file
- Auto-Triage Issues (1x) — transient?
- Daily Security Red Team Agent (1x) — transient?
- Daily Observability Report for AWF Firewall (1x) — transient?
- Semantic Function Refactoring (1x) — transient?

### Systemic Pattern
- Model/auth config cluster (4 workflows): Documentation Healer + Model Inventory + BYOK Ollama + Firewall Logs
- Token budget exhaustion cluster (4 workflows): Safe Output Health Monitor, jsweep, Daily Compiler, Firewall Logs
- CI blockage (CJS + CGO): affects all PR validation

### Actions Taken This Run
- 2 issues created: Dashboard 2026-06-06, Doc Healer/Model Inventory/Sentrux re-failure
- Updated shared memory
