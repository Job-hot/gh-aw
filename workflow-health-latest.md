# Workflow Health — 2026-05-16T05:35Z

Score: 64/100 (→ stable). 229 workflows. Run: §25953979404

## KEY FINDINGS

### Today's New Issues (May 16)
- **AWF Firewall v0.25.47 broken** (#32522): CONTAINED — `oidc-token-provider-base` module missing in Node bundle. PR #32503 closed without merging; main is unaffected. Smoke gate worked as designed. ✅
- **Smoke cluster failure at 01:00-01:02Z**: All smokes (Claude, Copilot, Gemini, OTEL, Agent Container, OTEL Backends) failed — root cause was AWF firewall v0.25.47 image on PR #32503. Now closed. Recovery in progress (Smoke OTEL ✅ at 05:28Z).
- **Smoke Codex failed** (#32561): still open (May 16)
- **Smoke Pi failed** (#32553): still open (May 16)
- **AW Compat**: 18/19 repos green, 1 hard failure (microsoft/aspire) — #32526
- **Open [aw] failure issues**: 19 (down from 30+ post-May 14 mass event)

### Persistent Issues (Unchanged)
- **CGO/CJS regression** (#29669): failing every push to main (P1)
- **Daily Fact parse failures** (#31432, #31524): still open (P2)
- **PR-review cluster waste** (#31724): ~272 wasted runs/day at 0% (P2, highest ROI fix)
- **MCP gateway session timeout** (#23153): still open (P2)
- **Performance Regression** (#30180): still open (P2)

### Actions Taken This Run
- Added comment to dashboard issue #29109
- Updated shared memory

### Trends
- Score: 64/100 (→ stable, day 2 post-recovery)
- Open [aw] failures: 19 (↓ from 30+ on May 14 peak; recovering)
- AWF firewall issue: isolated, contained by smoke gate ✅
- No new P0 cascading failures
- 229 workflows (stable)
