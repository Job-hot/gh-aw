# Workflow Health - 2026-04-11T12:00Z

Score: 73/100 (↓2 from 75 yesterday). 187 workflows. Run: §24281978807

## KEY FINDINGS

### NEW: Smoke Claude Engine Crash (#25727)
- Claude engine terminated exitCode 1, no output (Apr 11 00:44 UTC)
- Single failure - monitor for recurrence
- Previous state: ✅ healthy

### NEW: Daily Rendering Scripts Verifier - Placeholder False Positive
- `validate_prompt_placeholders.sh` incorrectly flags literal `__GH_AW_TRUE__`/`__GH_AW_FALSE__` strings
- These are test inputs embedded in the workflow description for render_template.cjs tests
- New issue created for tracking (#aw_render → actual number from safe-outputs)
- Run: §24281115984

## P2 Issues (Active)
- Smoke Claude: NEW failure (#25727, Apr 11, engine crash exitCode 1)
- Rendering Scripts Verifier: NEW bug (prompt placeholder false positive)
- Design Decision Gate: ongoing (#25548, --print empty prompt bug)
- Documentation Unbloat: ~$55/week Claude, 0 safe outputs
- Daily Issues Report: recurring (#25265, #25503, Copilot agent crash)
- Smoke Gemini: 100% failure (Gemini CLI 0.37.0 compatibility)

## P3 Issues (Ongoing)
- Daily Firewall Logs: safe_outputs process failure (#25456)
- Contribution Check: report_incomplete every run
- GitHub Remote MCP Auth Test: 100% failure

## Copilot Status
- v1.0.20 PINNED (stable)
- v1.0.24 bump PR #25752 CLOSED as draft (Apr 11 09:33)
- --no-ask-user PR #25772 OPEN

## Open Failure Issues (~18)
Key: #25727 (Smoke Claude, TODAY), #25265 (Daily Issues Report), #25456 (Daily Firewall)
#25503 (Deep-Report: Issues Report investigation), #25548 (Design Decision Gate)

## Score Breakdown
- Compilation: 187/187 ✅: +35
- Copilot v1.0.20 stable: +20
- Smoke Claude crash (new): -5
- Rendering Scripts false positive (new): -3
- DDG + Unbloat + Gemini: -7
- ~18 open failure issues: -8
- v1.0.24 upgrade blocked: -5
- Net: 73/100

## Score Trend
68 → 71 → 73 → 71 → 70 → 75 → 73 (Apr 5–11)

## Dashboard Issue
Created new issue #aw_dash411 (Apr 11)

Last updated: 2026-04-11T12:00Z
