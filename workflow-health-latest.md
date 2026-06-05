# Workflow Health — 2026-06-05T05:56Z

Score: 78/100 (↓3 from 81)
Workflows: 240 | Lock files: 240/240 (100% ✅) | Run: §26998143442

## KEY FINDINGS

### Status (June 5)
- **Compilation:** 240/240 workflows have lock files (100% ✅)
- **CJS typecheck** (#36410 OPEN): 100% failure since Jun 1 — DO NOT RE-FILE
- **CGO unit tests** (#35028 OPEN): 100% failure since May 27 — DO NOT RE-FILE
- **Daily Firewall Logs Collector**: 7th consecutive failure (token budget) — issue filed Jun 4. DO NOT RE-FILE
- **Daily BYOK Ollama Test**: Consecutive failures (auth) — issue filed Jun 4. DO NOT RE-FILE
- **Daily Documentation Healer**: NEW regression Jun 5 (effort param rejected). #37039 filed today. DO NOT RE-FILE
- **Daily Model Inventory Checker**: 2nd consecutive failure (BYOK auth). #37039 filed today. DO NOT RE-FILE
- **PR Triage Agent**: 3rd consecutive failure (#37035 filed today by notifier)
- **Daily Sentrux Report**: 2nd consecutive failure (#37019 filed today) — ESCALATING
- **Code Simplifier**: Recurring (#36829, #37057 filed today)
- **Designer Drift Audit**: New failure (#37059 filed today)
- **DataFlow PR & Discussion Dataset Builder**: New failure (#37054 filed today)
- **Smoke tests (PR #37004 branch)**: 5/5 failing (#37015, #37016, #37020, #37022, #37026 filed today)

### Critical Issues (P1) 🚨
- **CJS typecheck**: #36410 OPEN. DO NOT RE-FILE.
- **CGO unit tests**: #35028 OPEN. DO NOT RE-FILE.
- **Daily Firewall Logs Collector**: Filed Jun 4. DO NOT RE-FILE.
- **Daily BYOK Ollama Test**: Filed Jun 4. DO NOT RE-FILE.
- **Daily Documentation Healer**: #37039 filed today. DO NOT RE-FILE.
- **Daily Model Inventory Checker**: #37039 filed today. DO NOT RE-FILE.

### P2 Issues ⚠️
- **PR Triage Agent**: 3rd consecutive. #37035. Monitor for P1 escalation.
- **Daily Sentrux Report**: 2nd consecutive. #37019. Was transient, now pattern.
- **Smoke tests on PR #37004**: 5/5 failing — regression in PR branch. Monitor PR.
- **Linter Miner**: Still intermittent (~2/7) — monitoring.

### Systemic Pattern
- Model/auth config cluster (4 workflows): Documentation Healer + Model Inventory + BYOK Ollama + Firewall Logs
- CI blockage (CJS + CGO)

### Actions Taken This Run
- 0 new issues created (auto-notifiers and aw-failure-investigator covered all)
- Dashboard issue created
- Updated shared memory
- Dashboard issue created (workflow-health + report labels)
