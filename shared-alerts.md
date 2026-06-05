# Shared Alerts — 2026-06-05T05:56Z

## P0 (Critical) 🚨
_None active_

## P1 (High) 🚨
- **CJS typecheck** (#36410 OPEN): 100% failure since Jun 1. DO NOT re-file.
- **CGO unit tests** (#35028 OPEN): 100% failure since May 27. DO NOT re-file.
- **Daily Firewall Logs Collector**: 7th consecutive failure, token budget exhaustion. Issue filed Jun 4. DO NOT re-file.
- **Daily BYOK Ollama Test**: Consecutive failures, auth. Issue filed Jun 4. DO NOT re-file.
- **Daily Documentation Healer** (NEW - Jun 5): effort param rejected by Claude small-agent. Root-cause in #37039. DO NOT re-file.
- **Daily Model Inventory Checker** (NEW - Jun 5): BYOK auth missing, 2nd consecutive. Root-cause in #37039. DO NOT re-file.

## P2 (Watch) ⚠️
- **PR Triage Agent**: 3rd consecutive failure (Jun 3-5). #37035. Approaching P1.
- **Daily Sentrux Report**: 2nd consecutive failure (Jun 4-5). #37019. Escalating from transient.
- **Code Simplifier**: Recurring failure. #36829, #37057.
- **Designer Drift Audit**: New failure Jun 5. #37059.
- **DataFlow PR & Discussion Dataset Builder**: New failure Jun 5. #37054.
- **Smoke tests on PR #37004 branch**: 5/5 failing with missing tools / no safe output. #37015, #37016, #37020, #37022, #37026. Monitor PR before merge.
- **Safe Outputs Conformance SEC-005** (#36591 OPEN): update_activation_comment.cjs allowlist gap — unchanged.
- **chaos-test PR stall**: 10+ open PRs pending merge — still unresolved.

## Resolved ✅ (since Jun 4)
_None resolved today_

## Systemic Notes
- **Model/auth config cluster**: Documentation Healer (effort param) + Model Inventory (BYOK auth) + BYOK Ollama (auth) + Firewall Logs (token budget) = 4 workflows with bootstrap/config failures. #37039 has root cause for first two.
- **CI blockage**: CJS + CGO both 100% failing — affects all PR validation.
- **Smoke test regression on PR #37004**: All 5 engine smoke tests failing simultaneously on this PR branch — suggests regression. Warrants investigation before merge.
- Health score trending down: 82 → 81 → 78 (3-day drop). New failures outpacing resolutions.
