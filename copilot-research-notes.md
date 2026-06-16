# Copilot CLI Research Notes (Last 7 runs)

### 2026-06-16 (Run 27596173580) — This Run
- **249 total workflows** (up from 246 on 06-15 — +3 net)
- 136 Copilot (~55%): 40 scalar form + 96 object form; 63 copilot-sdk
- **�� IMPROVEMENTS vs last run:**
  - `max-ai-credits`: 6 → 18 (+200% — budget guardrails finally spreading!)
  - `min-integrity`: 22 → 34 (+12, improving security posture)
  - `total workflows`: 246 → 249 (+3)
  - `strict: true`: 79 → 151 (may include broader engine count now)
- **⚠️ REGRESSIONS vs last run:**
  - `engine.agent`: 21 → 8 (further decline from 34 peak on 06-10)
- **max-continuations: 7** (stable — 5% of copilot workflows)
- **copilot-sdk: 63** (stable — 46% of all copilot workflows)
- **copilot-sdk-driver: 3** (stable)
- **network configured: 132** (all copilot workflows now have some network config!)
- **sandbox: 23** (17% — 113 without sandbox)
- **min-integrity: 34** (only 6 in copilot-specific workflows)
- **max-ai-credits: 18/136** (13% — 118/136 still missing)
- **max-tool-denials: 0/63 SDK** (PERSISTENT GAP — 0% adoption)
- **experiments: 41 workflows** (new tracking — A/B testing feature well adopted)
- **mcp.session-timeout**: 0 (PERSISTENT GAP)
- **mcp.tool-timeout**: 0 (PERSISTENT GAP)
- **blocked-domains**: 1 (new feature, nearly unadopted)
- **startup-timeout**: 1/249 (barely used)
- **engine.args**: 0 (PERSISTENT GAP, 19th consecutive run)
- **engine.api-target**: 0 (PERSISTENT GAP, 19th consecutive run)
- **engine.harness**: 0 (never used)
- **engine.token-weights**: 0 (never used)
- **Unused agent files**: grumpy-reviewer, interactive-agent-designer, w3c-specification-writer, create-safe-output-type, custom-engine-implementation (5 files — unchanged)
- Discussion created: "Copilot CLI Deep Research - 2026-06-16"

### 2026-06-15 (Run 27525865107)
- **246 total workflows** (up from 245 on 06-10 — +1 net)
- 133 Copilot (54%), 47 Claude, 10 Codex
- **⚠️ REGRESSIONS vs last run:**
  - `engine.agent`: 34 → 21 (-13 workflows removed agent references!)
  - `max-ai-credits`: 14 → 6 (-8 workflows lost budget guard)
  - `sandbox`: 20 → 15 (-5 workflows lost sandbox)
- **max-continuations: 7** (up from 6 — slight growth)
- **copilot-sdk: 63** (stable — 47% of all copilot workflows)
- **engine.args**: 0 (PERSISTENT GAP, 18th consecutive run)
- **engine.api-target**: 0 (PERSISTENT GAP, 18th consecutive run)
- Discussion created: "Copilot CLI Deep Research - 2026-06-15"

### 2026-06-10 (Run 27254548925)
- **245 total workflows** (down from 340 on 06-08 — repo cleanup/consolidation removed ~95 workflows)
- 132 Copilot (54%), 64 Claude, 16 Codex
- **engine.agent: 34 workflows** (up from 13! — +161% growth — major improvement, then reverted)
- **copilot-sdk-driver: 5** (up from 3 — 2 new custom drivers added)
- **max-ai-credits: 14** (up from 5 — +180% improvement, but 118/132 still missing)
- **engine.args**: 0 (PERSISTENT GAP, 17th consecutive run)
- Discussion created: "Copilot CLI Deep Research - 2026-06-10"

### 2026-06-08 (Run 27117423076)
- **340 total workflows** (from 236 on 05-31, +104 or +44% in 8 days)
- 132 Copilot (39%), 64 Claude, 16 Codex
- **copilot-sdk: 63 workflows (48%)** — MASSIVE adoption
- **BYOK: 2 workflows** (new: Azure OpenAI smoke test added)
- Discussion created: "Copilot CLI Deep Research - 2026-06-08"

### 2026-05-31 (Run 26703913319)
- 236 total workflows; 97 Copilot (41%), 51 Claude, 9 Codex
- **max-continuations**: 5 workflows (unchanged)
- **cache-memory**: 116 workflows (significant growth)
- **sandbox AWF**: 23 workflows
- Discussion created: "Copilot CLI Deep Research - 2026-05-31"

### 2026-05-27 (Run 26491933777)
- 236 total workflows; 125 Copilot (53%)
- Discussion created: "Copilot CLI Deep Research - 2026-05-27"

### 2026-05-21 (Run 26206481620)
- 233 total MD workflows; 100 Copilot (43%)
- Discussion created: "Copilot CLI Deep Research - 2026-05-21"

---

## Key Persistent Gaps (Tracked Across All Runs)

1. **engine.args** — 19+ consecutive runs with ZERO usage (custom CLI arguments)
2. **engine.api-target** — 19 consecutive runs with ZERO usage (custom API endpoints)
3. **engine.harness** — Never used (custom harness scripts)
4. **engine.version** (Copilot pin) — Zero explicit Copilot CLI version pins (50 with version: in engine block, may be other fields)
5. **max-continuations** — Only 7/136 copilot workflows (5%) use autopilot mode
6. **MCP session/tool timeout** — Never configured (engine.mcp.session-timeout, engine.mcp.tool-timeout)
7. **engine.token-weights** — Never used
8. **max-tool-denials** — 0/63 SDK workflows (should pair with copilot-sdk: true)
9. **startup-timeout** — Only 1/249 workflows
10. **blocked-domains** — Only 1/249 (brand new feature)

## Trends

- `engine.agent` adoption: 25→25→13→34→21→8 (very volatile, declining trend — concerning)
- `copilot-sdk`: 0 → 63 (stabilized at 63 since June 2026)
- `copilot-sdk-driver`: 0 → 3 → 3 → 3 (stable small usage)
- `max-ai-credits`: 0→5→14→6→18 (recovering after last week's regression)
- `min-integrity`: 22→34 (improving slowly)
- `experiments`: 41 (new tracking — strong adoption of A/B testing)
- `max-continuations` adoption: 5→5→6→6→7 (very slow growth)
- Total workflows: 233→236→340→245→246→249 (stabilized at ~249)
- Copilot share: 43%→53%→39%→54%→54%→55% (very stable)
- Persistent gaps unchanged across 19+ runs — likely not on developer radar
