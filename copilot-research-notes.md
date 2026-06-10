# Copilot CLI Research Notes (Last 6 runs)

### 2026-06-10 (Run 27254548925) — This Run
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
- **engine.args**: 0 (PERSISTENT GAP, 17th+ consecutive run)
- **engine.api-target**: 0 (PERSISTENT GAP, 17th consecutive run)
- **engine.harness**: 0 (never used)
- **engine.version** (Copilot pin): 0 (persistent gap)
- **MCP session/tool timeout**: 0 (never used)
- **engine.token-weights**: 0 (never used)
- **Unused agent files**: grumpy-reviewer, interactive-agent-designer, w3c-specification-writer, create-safe-output-type, custom-engine-implementation (5 files — unchanged)
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
- **engine.args**: 0 (PERSISTENT GAP, 16th+ consecutive run)
- **engine.api-target**: 0 (PERSISTENT GAP, 16th consecutive run)
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

### 2026-05-26 (Run 26433217435)
- 236 total workflows; 126 Copilot (53%), 63 Claude, 16 Codex
- **engine.agent**: 26 workflows (grew from 10)
- **engine.args**: 0 (PERSISTENT GAP, 14th consecutive run)
- Discussion created: "Copilot CLI Deep Research - 2026-05-26"

### 2026-05-25 (Run 26384338048)
- 235 total workflows; 97 Copilot (41%)
- **engine.args**: 0 (PERSISTENT GAP, 13th+ consecutive run)
- Discussion created: "Copilot CLI Deep Research - 2026-05-25"

### 2026-05-21 (Run 26206481620)
- 233 total MD workflows; 100 Copilot (43%)
- **engine.args**: 0 (PERSISTENT GAP, 12th consecutive run)
- Discussion created: "Copilot CLI Deep Research - 2026-05-21"

---

## Key Persistent Gaps (Tracked Across All Runs)

1. **engine.args** — 17+ consecutive runs with ZERO usage (custom CLI arguments)
2. **engine.api-target** — 17 consecutive runs with ZERO usage (custom API endpoints)
3. **engine.harness** — Never used (custom harness scripts)
4. **engine.version** (Copilot pin) — Zero explicit Copilot CLI version pins
5. **max-continuations** — Only 6/132 workflows (4.5%) use autopilot mode
6. **MCP session/tool timeout** — Never configured
7. **engine.token-weights** — Never used

## Trends

- `engine.agent` adoption: 10 → 13 → 26 → 34 (clear growth trajectory)
- `copilot-sdk-driver`: 0 → 3 → 5 (growing custom driver usage)
- `max-ai-credits`: 5 → 5 → 5 → 14 (spike this run — good improvement)
- `web-fetch`: 12 → 25 (doubled in 2 days)
- `max-continuations` adoption: 5 → 12 → 5 → 6 (fluctuating, not consistently growing)
- Total workflows: 233 → 236 → 340 → 245 (explosion then cleanup)
- Copilot share: 43% → 53% → 39% → 54% (varies with total count)
- Persistent gaps remain unchanged across 17+ runs — truly non-discovered or purposely unused
