# Shared Alerts — 2026-05-03T13:00Z

## P0 (Critical)
- **Smoke Gemini** (#29459, #29852 OPEN): 100% failure, API_KEY_INVALID. Every PR sees red. 30+ days unresolved.
- **Smoke Copilot/Claude regression** (#29863, #29864 OPEN): New regression wave ~03:37 UTC May 3. Copilot 2-3/5, Claude 3/5 failures.
- **Smoke Codex**: 1/5 failures today (minor regression in same wave)
- **Smoke CI** (#29666 OPEN): CGO/EROFS, 4/5 action_required persisting.
- **CGO build** (#29669 OPEN): ongoing failures.

## P1 (High)
- **Q prompt instability**: 0→72 turn variance (avg 16) — resource unpredictability risk
- **GitHub API Consumption Report**: 25.6m runtime — MCP timeout risk (>5m inactivity limit)
- **Design Decision Gate**: 50% failure rate (May 2) — Claude errors, undiagnosed
- **MCP gateway session timeout** (#23153 OPEN): Long-running workflows at risk.

## P2 (Watch)
- **Node.js 20 deprecation** in CI: deadline Sep 16, 2026. Migrate to Node.js 22.
- **YAMLGeneration regression** (#29779): 21.7% slower — no dedicated monitoring agent.
- **Safe Outputs SEC-004** (#27235 OPEN).
- **6 PR-review agents** on same triggers — evaluate redundancy (Scout, Archie, /cloclo, Q, AI Moderator, Content Moderation)

## Resolved (Do Not Re-File)
- #29088 Codex crash → CLOSED
- #28659 Doc Unbloat claude auth → CLOSED
- #27965 GitHub Remote MCP Auth → CLOSED
- #27888 awf-api-proxy sidecar → CLOSED
- #27251 GitHub App rate limit → CLOSED
- #27512 CODEX_HOME collision → CLOSED

## Trends
- 209 workflows, 0 missing lock files
- Quality: 74/100 (stable 7-day plateau)
- Effectiveness: 71/100 (stable)
- New smoke regression wave May 3 03:37 UTC — Copilot, Claude, Codex
- Gemini still completely broken (30+ days, P0 unresolved)
