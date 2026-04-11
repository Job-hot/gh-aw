# Shared Alerts — 2026-04-11T12:00Z

## P2 (High)
- **Smoke Claude engine crash** (NEW, #25727, Apr 11 00:44 UTC): Claude engine exitCode 1, no output. Single failure — monitor for recurrence. Was previously healthy.
- **Daily Rendering Scripts Verifier false positive** (NEW, issue created Apr 11): `validate_prompt_placeholders.sh` incorrectly flags literal `__GH_AW_TRUE__`/`__GH_AW_FALSE__` in render_template test inputs. Blocks daily rendering validation. Fix in workflow or validator needed.
- **Design Decision Gate broken** (#25548): Empty prompt when --print flag used. Architecture decisions blocked. Fix documented, awaiting PR.
- **Documentation Unbloat zero-output** (ongoing): Claude workflow ~$55/week, 0 safe outputs. Agent runs but never calls output tools.
- **Smoke Gemini failing**: 100% failure (Gemini CLI 0.37.0 compat). No tracking issue visible.
- **Daily Issues Report recurring failure** (#25265, #25503): Copilot agent killed (orphan process). Copilot v1.0.20 related.

## P3 (Watch)
- **Contribution Check report_incomplete**: Every run. Permission/network issue.
- **GitHub Remote MCP Auth Test**: 100% failure. #24829 closed not_planned but test still failing.
- **Daily Firewall Logs safe_outputs failure** (#25456): Process Safe Outputs step fails (likely report_incomplete), artifact 287 bytes.
- **Workflow Normalizer deduplication gap**: Created 3 identical issues in 24h (still monitoring).

## Copilot Version Status
- v1.0.20 PINNED (stable since Apr 10 repin)
- v1.0.24 bump PR #25752: CLOSED as draft Apr 11 09:33 — NO NEW BUMP DEPLOYED
- --no-ask-user flag PR #25772: OPEN (Copilot engine improvement)

## Recent Fixes / Changes
- Apr 10: v1.0.20 re-pinned after v1.0.22 regression fully stabilized
- Apr 11 09:17: PR #25772 opened (--no-ask-user for Copilot engine)

## Ecosystem State
- 187 compiled workflows. Health: ~73/100 (↓2 from 75). Score dipped: Smoke Claude crash + Rendering Scripts false positive.
- Engine split: ~124 copilot, ~41 claude, ~18 codex, ~4 others
- v1.0.20 currently pinned as stable Copilot version
- Claude/Codex engines: mostly resilient (Smoke Claude 1 new failure today)

## Orchestrator Summaries (Apr 11)
- Workflow Health (Apr 11 12:00): Score 73/100 ↓2. New: Smoke Claude crash, Rendering Scripts false positive.
- Agent Performance (Apr 11 04:31): Q:70↑5 E:60↓6. Recovery trend continuing.

Last updated: 2026-04-11T12:00Z by workflow-health-manager
