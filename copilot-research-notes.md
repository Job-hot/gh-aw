# Copilot CLI Research Notes (Last 6 runs)

### 2026-06-15 (Run 27525865107) — This Run
- **246 total workflows** (up from 245 on 06-10 — +1 net)
- 133 Copilot (54%), 47 Claude, 10 Codex
- **⚠️ REGRESSIONS vs last run:**
  - `engine.agent`: 34 → 21 (-13 workflows removed agent references!)
  - `max-ai-credits`: 14 → 6 (-8 workflows lost budget guard)
  - `sandbox`: 20 → 15 (-5 workflows lost sandbox)
- **max-continuations: 7** (up from 6 — slight growth)
- **copilot-sdk: 63** (stable — 47% of all copilot workflows)
- **copilot-sdk-driver: 3** (stable)
- **strict: true: 79** (stable at 59%)
- **network: 67** (50% — 66/133 still have NO network config)
- **min-integrity: 22/133** (17% — 83% missing)
- **max-ai-credits: 6/133** (5% — 95% missing)
- **max-tool-denials: 0/63 SDK** (0% — NEW PERSISTENT GAP for SDK workflows)
- **engine.args**: 0 (PERSISTENT GAP, 18th consecutive run)
- **engine.api-target**: 0 (PERSISTENT GAP, 18th consecutive run)
- **engine.harness**: 0 (never used)
- **engine.token-weights**: 0 (never used)
- **Unused agent files**: grumpy-reviewer, interactive-agent-designer, w3c-specification-writer, create-safe-output-type, custom-engine-implementation (5 files — unchanged from previous runs)
- Discussion created: "Copilot CLI Deep Research - 2026-06-15"

### 2026-06-10 (Run 27254548925)
- **245 total workflows** (down from 340 on 06-08 — repo cleanup/consolidation removed ~95 workflows)
- 132 Copilot (54%), 64 Claude, 16 Codex
- **engine.agent: 34 workflows** (up from 13! — +161% growth — major improvement)
- **copilot-sdk-driver: 5** (up from 3 — 2 new custom drivers added)
- **max-ai-credits: 14** (up from 5 — +180% improvement, but 118/132 still missing)
- **web-fetch: 25** (up from 12 — +108% growth)
- **bare mode: 15** (stable)
- **copilot-sdk: 63** (stable)
- **max-continuations: 6** (stable — persistently underused at 4.5%)
- **sandbox AWF: 15-20 workflows** (stable — 112/132 still unprotected, 85%)
- **network open: 65/132** (49% still no restrictions)
- **strict: true: 79/132** (60% — 53 workflows still missing)
- **BYOK: 2** (stable)
- **engine.args**: 0 (PERSISTENT GAP, 17th consecutive run)
- **engine.api-target**: 0 (PERSISTENT GAP, 17th consecutive run)
- Discussion created: "Copilot CLI Deep Research - 2026-06-10"

### 2026-06-08 (Run 27117423076)
- **340 total workflows** (from 236 on 05-31, +104 or +44% in 8 days)
- 132 Copilot (39%), 64 Claude, 16 Codex
- **copilot-sdk: 63 workflows (48%)** — MASSIVE adoption
- **BYOK: 2 workflows** (new: Azure OpenAI smoke test added)
- **engine.agent**: ~13 workflows
- **max-ai-credits**: only 5 workflows
- **sandbox AWF**: 22 workflows
- **network open**: 65 of 130 Copilot workflows (50%)
- **strict: true**: 77/130 (59%)
- **engine.args**: 0 (PERSISTENT GAP, 16th consecutive run)
- Discussion created: "Copilot CLI Deep Research - 2026-06-08"

### 2026-05-31 (Run 26703913319)
- 236 total workflows; 97 Copilot (41%), 51 Claude, 9 Codex
- **max-continuations**: 5 workflows (unchanged)
- **cache-memory**: 116 workflows (significant growth)
- **sandbox AWF**: 23 workflows
- **engine.args**: 0 (PERSISTENT GAP, 15th consecutive run)
- Discussion created: "Copilot CLI Deep Research - 2026-05-31"

### 2026-05-27 (Run 26491933777)
- 236 total workflows; 125 Copilot (53%)
- **max-continuations**: 5 workflows (stable)
- **engine.args**: 0 (PERSISTENT GAP, 15th consecutive run)
- Discussion created: "Copilot CLI Deep Research - 2026-05-27"

### 2026-05-21 (Run 26206481620)
- 233 total MD workflows; 100 Copilot (43%)
- **engine.args**: 0 (PERSISTENT GAP, 12th+ consecutive run)
- Discussion created: "Copilot CLI Deep Research - 2026-05-21"

---

## Key Persistent Gaps (Tracked Across All Runs)

1. **engine.args** — 18+ consecutive runs with ZERO usage (custom CLI arguments)
2. **engine.api-target** — 18 consecutive runs with ZERO usage (custom API endpoints)
3. **engine.harness** — Never used (custom harness scripts)
4. **engine.version** (Copilot pin) — Zero explicit Copilot CLI version pins
5. **max-continuations** — Only 7/133 workflows (5%) use autopilot mode
6. **MCP session/tool timeout** — Never configured
7. **engine.token-weights** — Never used
8. **max-tool-denials** — 0/63 SDK workflows (NEW: should pair with copilot-sdk: true)

## Trends

- `engine.agent` adoption: 25 → 25 → 13 → 34 → 21 (very volatile — not sustainably growing)
- `copilot-sdk`: 0 → 63 (exploded in June 2026)
- `copilot-sdk-driver`: 0 → 3 → 3 (stable small usage)
- `max-ai-credits`: 0 → 5 → 14 → 6 (inconsistent — recent regression)
- `web-fetch`: 21 → 12 → 25 → 10 (volatile)
- `max-continuations` adoption: 5 → 5 → 6 → 6 → 7 (very slow growth)
- Total workflows: 233 → 236 → 340 → 245 → 246 (stabilized)
- Copilot share: 43% → 53% → 39% → 54% → 54% (stabilized at 54%)
- Persistent gaps remain unchanged across 18+ runs — truly non-discovered or purposely unused
