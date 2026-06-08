# Shared Alerts — 2026-06-08T14:02Z

## P0 (Critical) 🚨
- **Failure Cascade** (#37721 OPEN Jun 8): awf-cli-proxy container exit (1) + LLM max_runs_exceeded (50/50). Root cause NOT patched. Possibly related to firewall bump #37707 (v0.25.66). DO NOT RE-FILE.

## P1 (High) 🚨
- **CJS typecheck** (#aw_cjs8 filed Jun 8): Re-regression after #37503 closed prematurely Jun 7. Still failing on main. DO NOT RE-FILE.
- **CGO unit tests** (#35028 OPEN): 100% failing Jun 8. DO NOT RE-FILE.
- **Daily Compiler Quality Check** (#37730 OPEN, 3rd day): Excessive tool denials (5/5). Escalated to P1 via comment Jun 8. DO NOT RE-FILE.

## P2 (Watch) ⚠️
- **AI Moderator** (#37723, cascade-suspected): Still failing. Covered under cascade #37721. Monitor.
- **Safe Output Health Monitor** (#37759 OPEN Jun 8): Re-failing after Jun 7 success. Token exhaustion. DO NOT RE-FILE.
- **Code Simplifier** (#37733 OPEN): Re-failed Jun 8 after Jun 7 success. First recurrence — monitor.
- **Daily Model Inventory Checker**: Failed Jun 8 (auto-filed #37683, then closed). Auto-cycle continuing. Monitor.
- **Issue Lifecycle Gap** (#aw_isg_jun8 OPEN Jun 8): Systemic process issue — premature P1 closure. Action: improvement issue filed.

## Resolved ✅ (since Jun 7)
- **Daily Documentation Healer**: SUCCESS Jun 8 ✅ — model pinning fix (#37505) confirmed
- **Daily Sentrux Report**: Continuing healthy ✅
- **PR Sous Chef**: Healthy ✅

## Systemic Notes
- **awf-cli-proxy cascade**: 10+ smoke + agentic workflows affected; batch-close path needs root cause fix first
- **CI blockage cluster**: CJS + CGO both failing on main — all PR branches lack full CI validation
- **Tool denial cluster (Day 3)**: Compiler Quality + Safe Output Health both hitting tool denial limit
- **Issue lifecycle gap (2nd occurrence)**: CJS #37503 closed prematurely again — systemic improvement issue filed (#aw_isg_jun8)
- **Health score trend**: 82→81→78→74→71→68 (6-day decline) — cascade significantly worsening trajectory
- **copilot-swe-agent**: Continued healthy throughput; 7 merges Jun 8; counterbalancing quality decline
