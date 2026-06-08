# Copilot CLI Research Notes (Last 6 runs)

### 2026-06-08 (Run 27117423076) — This Run
- **340 total workflows** (from 236 on 05-31, +104 or +44% in 8 days — explosive growth!)
- 132 Copilot (39%), 64 Claude, 16 Codex
- **copilot-sdk: 63 workflows (48%)** — MASSIVE new feature, effectively viral adoption
  - 3 workflows use custom SDK drivers: Python (.py), Node.js (.cjs), TypeScript (.ts)
- **BYOK: 2 workflows** (new: Azure OpenAI smoke test `smoke-copilot-aoai-apikey.md` added alongside Ollama)
- **max-turns (top-level)**: ~13 Copilot workflows, growing adoption
- **max-ai-credits**: only 5 workflows (125 missing out of 130)
- **sandbox AWF**: 22 workflows (stable, `sandbox: agent: awf` format confirmed)
- **network open (no config)**: 65 of 130 Copilot workflows (50%) have zero network restrictions
- **strict: true**: 77/130 (59%, growing)
- **engine.agent**: ~13 workflows using custom agents
- **Unused agent files**: grumpy-reviewer, interactive-agent-designer, w3c-specification-writer, create-safe-output-type, custom-engine-implementation (5 files - unchanged)
- **engine.args**: 0 (PERSISTENT GAP, 16th+ consecutive run)
- **engine.api-target**: 0 (PERSISTENT GAP, 16th consecutive run)
- **engine.harness**: 0 (never used)
- **engine.version** (Copilot pin): 0 (persistent gap)
- **MCP session/tool timeout**: 0 (never used)
- **engine.token-weights**: 0 (never used)
- Discussion created: "Copilot CLI Deep Research - 2026-06-08"

### 2026-05-31 (Run 26703913319)
- 236 total workflows (same as 05-27); 97 Copilot (41%), 51 Claude, 9 Codex
- **max-continuations**: 5 workflows (unchanged — persistently underused)
- **bare mode**: 12 workflows (slight growth)
- **cache-memory**: 116 workflows (all engines — significant growth from 95)
- **sandbox AWF**: 23 workflows (growth from 16)
- **model: small**: 11 copilot-specific; model: large: 1; 86 no explicit model
- **mcp-scripts**: 12 workflows (stable)
- **engine.agent**: 20+ workflows using custom agent files
- **engine.args**: 0 (PERSISTENT GAP, 15th consecutive run)
- **engine.api-target**: 0 (PERSISTENT GAP, 15th consecutive run)
- **engine.harness**: 0 (never used)
- **BYOK**: 1 workflow (daily-byok-ollama-test)
- **network configured copilot**: 49/97 workflows (49% with network restrictions, 51% without)
- **github toolsets [default]**: 34 workflows; [default, issues]: 7; [default, discussions]: 7
- **Unused agent files**: grumpy-reviewer, interactive-agent-designer, w3c-specification-writer, create-safe-output-type, custom-engine-implementation (5 files unchanged)
- Discussion created: "Copilot CLI Deep Research - 2026-05-31"

### 2026-05-27 (Run 26491933777)
- 236 total workflows; 125 Copilot (53%) [96 simple + 29 block form], 63 Claude, 16 Codex
- **max-continuations**: 5 workflows (stable)
- **cache-memory**: 95 workflows (stable)
- **engine.agent**: used: adr-writer, ci-cleaner, contribution-checker, developer.instructions, technical-doc-writer, agentic-workflows, awf(x16)
- **bare mode**: 9 workflows (↓ from 11 - slight decrease)
- **engine.args**: 0 (PERSISTENT GAP, 15th consecutive run)
- **engine.api-target**: 0 (PERSISTENT GAP, 15th consecutive run)
- **engine.env**: 0 (PERSISTENT GAP, 15th consecutive run)
- **engine.version pinning**: 0 (no copilot-specific version pins)
- **BYOK**: 0 (persistent gap)
- **harness**: 0 (persistent gap)
- **network restrictions**: 62 of 125 copilot workflows missing (50%)
- **sandbox**: 110 of 125 copilot workflows missing (88%)
- Discussion created: "Copilot CLI Deep Research - 2026-05-27"

### 2026-05-26 (Run 26433217435)
- 236 total workflows (+1); 126 Copilot (53%), 63 Claude, 16 Codex, 1 Gemini
- **max-continuations**: 5 workflows
- **cache-memory**: 95 workflows
- **engine.agent**: 26 workflows (grew from 10)
- **strict mode**: 146 workflows (62% coverage)
- **engine.args**: 0 (PERSISTENT GAP, 14th consecutive run)
- **engine.api-target**: 0 (PERSISTENT GAP, 14th consecutive run)
- Discussion created: "Copilot CLI Deep Research - 2026-05-26"

### 2026-05-25 (Run 26384338048)
- 235 total workflows (+2); 97 Copilot (41%), 51 Claude, 10 Codex
- **max-continuations**: 12 workflows (↑ significant growth)
- **sandbox AWF**: 22 workflows
- **engine.args**: 0 (PERSISTENT GAP, 13th+ consecutive run)
- Discussion created: "Copilot CLI Deep Research - 2026-05-25"

### 2026-05-21 (Run 26206481620)
- 233 total MD workflows (+1); 100 Copilot (43%)
- **max-continuations**: 5 workflows
- **engine.args**: 0 (PERSISTENT GAP, 12th consecutive run)
- **engine.api-target**: 0 (12th consecutive run)
- Discussion created: "Copilot CLI Deep Research - 2026-05-21"

---

## Key Persistent Gaps (Tracked Across All Runs)

1. **engine.args** — 16+ consecutive runs with ZERO usage (custom CLI arguments)
2. **engine.api-target** — 16 consecutive runs with ZERO usage (custom API endpoints)
3. **engine.harness** — Never used (custom harness scripts)
4. **engine.version** (Copilot pin) — Zero explicit Copilot CLI version pins
5. **max-continuations** — Only 6/130 workflows (5%) use autopilot mode
6. **MCP session/tool timeout** — Never configured
7. **engine.token-weights** — Never used
8. **max-ai-credits** — Only 5/130 Copilot workflows set this (125 workflows using default limit)
9. **Network open** — 65/130 Copilot workflows (50%) have no network restrictions at all

## Trends

- `copilot-sdk` adoption: 0 → 63 (EXPLOSIVE in one cycle — entirely new capability)
- `copilot-sdk-driver`: 3 custom drivers (Python, Node.js, TypeScript) exploring the feature
- `sandbox AWF`: 16 → 22 → 22 (stable)
- `BYOK`: 0 → 1 → 2 (slow but growing)
- `max-continuations` adoption: 5 → 12 → 5 → 6 (fluctuating, not consistently growing)
- `strict: true`: ~62% → 59% (stable/slight dip due to many new unstricited workflows)
- Persistent gaps remain unchanged across 16+ runs — truly non-discovered or purposely unused
- Repository doubled in size: 236 → 340 workflows in 8 days (likely bulk workflow import/generation)
