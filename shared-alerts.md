# Shared Alerts — 2026-06-11T14:15Z

## P0 (Critical) 🔴
- **Failure Investigator (6h) AI Credits** (#38559, NEW OPEN): Meta-monitor itself hitting 1K AIC limit (over by 25.1). Creates blind spot in failure detection. Fix: raise max-ai-credits to 2000 in aw-failure-investigator. DO NOT RE-FILE #38559.

## P1 (High) 🚨
- **AI Credits Cluster June 11 — EXPANDED** (prev #38520 CLOSED; no open tracker): Now 6+ workflows. New members: Package Specification Extractor (1,023.9 AIC — largest single consumer), Failure Investigator (6h) (1,025 AIC). Original 4: Code Simplifier, Ubuntu Analyzer, Workflow Skill Extractor, Matt Pocock. Fix: raise max-ai-credits to 2000 for all analysis-heavy workflows. Individual failures: #38499, #38500, #38501, #38497, #38576, #38559 — DO NOT RE-FILE.
- **Daily News Node.js chroot** (#38379, Jun 10 OPEN): Node.js unreachable in AWF chroot (Day 3). Assigned @zarenner. DO NOT RE-FILE.
- **Daily Safe Outputs Git Simulator** (#aw_gitsim10, ongoing): memory/git-simulator branch missing signed-commit seed. DO NOT RE-FILE.

## P2 (Watch) ⚠️
- **GitHub MCP Structural Analysis** (#38614, Jun 11): Failed today — investigate root cause.
- **PR Sous Chef** (#38618, Jun 11): Failed today — 1 occurrence, monitor.
- **Daily MCP Tool Concurrency Analysis** (#38578, Jun 11): Failed today.
- **Daily Safe Outputs Conformance Checker** (#38540, Jun 11): Failed today.
- **Smoke test flakiness** (#38605, #38602, #38582, #38581, #38599, #38597): Multiple engines — likely transient/test-env.
- **jsweep tool denial** (#38505, Jun 11): Guardrail fired, one-off. Monitor.

## Resolved (Jun 9→11) ✅
- #38520 AI Credits Cluster (4-workflow tracker) — CLOSED
- #38389 Failure cascade — CLOSED
- Jun 10 bulk individual failures — most CLOSED on credits reset

## Systemic Notes
- **Health score trend:** 68→83→87→85→83 (AI credits cluster persisting, expanding)
- **AI credits persistent pattern:** Analysis-heavy workflows ROUTINELY exceed 1K. Cluster grew to 6+. Failure Investigator is now also affected — meta-monitor reliability at risk. Fix: raise to 2K. FOURTH day of same pattern.
- **Failure Investigator blind spot (NEW):** Meta-monitor failing from AI credits creates a ~6h detection gap each day. Priority fix needed.
- **Package Specification Extractor:** Largest single-run AIC consumer (1,023.9). Needs budget tuning independent of cluster fix.
- **Smoke tests:** Multiple engine failures — likely flaky/transient. Monitor for 2+ days before escalating.

## Do Not Re-File (Active Issues)
- #38559: Failure Investigator (6h) AI credits
- #38499: Code Simplifier failed
- #38500: Ubuntu Actions Image Analyzer failed
- #38501: Workflow Skill Extractor failed
- #38497: Matt Pocock Skills Reviewer failed
- #38576: Package Specification Extractor failed
- #38379: Daily News chroot (assigned @zarenner)
- #38505: jsweep tool denial
