# Copilot CLI Research Notes (Last 6 runs)

### 2026-05-25 (Run 26384338048) — This Run
- 235 total workflows (+2); 97 Copilot (41%), 51 Claude, 10 Codex
- **max-continuations**: 12 workflows (↑ from 5 — significant growth, teams embracing longer autonomous runs)
- **max-runs**: 2 workflows (severely underused, persistent gap)
- **bare mode**: 11 workflows (stable)
- **cache-memory**: 30 workflows (copilot-only count; higher across all engines)
- **sandbox AWF**: 22 workflows (growing from 16)
- **web-fetch**: 20 workflows (stable)
- **model: small**: 23 workflows; model: large: 4; explicit models: ~8
- **mcp-scripts**: 3 workflows (minimal adoption)
- **engine.agent**: 10 workflows (awf most common; 5 custom files used: technical-doc-writer, ci-cleaner, contribution-checker, adr-writer, developer.instructions)
- **engine.args**: 0 (PERSISTENT GAP, 13th+ consecutive run with ZERO usage)
- **engine.api-target**: 0 (PERSISTENT GAP, 13th consecutive run)
- **engine.version pinning**: 0 (no explicit version pins)
- **BYOK**: 0 (persistent gap)
- **toolsets: [all]**: 5 workflows (over-permissioned)
- **missing timeout**: ~20 workflows without timeout-minutes
- **Unused agent files**: grumpy-reviewer, interactive-agent-designer, w3c-specification-writer, create-safe-output-type, custom-engine-implementation (5 files unused)
- Discussion created: "Copilot CLI Deep Research - 2026-05-25"

### 2026-05-21 (Run 26206481620)
- 233 total MD workflows (+1); 100 Copilot (43%)
- **max-continuations**: 5 workflows
- **max-runs**: 2 workflows
- **bare mode**: 11 workflows
- **cache-memory**: 73 workflows (all engines)
- **sandbox AWF**: 16 workflows
- **web-search/fetch**: 21 workflows
- **model overrides**: 48 total
- **engine.agent**: 11 workflows
- **engine.args**: 0 (PERSISTENT GAP, 12th consecutive run)
- **engine.api-target**: 0 (12th consecutive run)
- **BYOK**: 0
- Discussion created: "Copilot CLI Deep Research - 2026-05-21"

### 2026-05-18 (Run 26014468484)
- 230 total MD workflows; 126 Copilot (55% — different counting method)
- **web-fetch**: 21 workflows (strong growth from ~2)
- **engine.args**: 0 (10th+ consecutive run)
- **engine.api-target**: 0 (11th consecutive run)

### 2026-05-14 (Run 25842508637)
- 225 total MD workflows; copilot: 97 simple + 24 extended block
- **mcp-scripts**: 12 workflows
- **cache-memory**: 92 workflows (all engines)
- **network.allowed**: 116 workflows
- **engine.args**: 0 (persistent gap)
- **engine.api-target**: 0 (persistent gap)

[Earlier entries available in git history]

---

## Key Persistent Gaps (Tracked Across All Runs)

1. **engine.args** — 13+ consecutive runs with ZERO usage (custom CLI arguments)
2. **engine.api-target** — 13 consecutive runs with ZERO usage (custom API endpoints)
3. **engine.harness** — Never used (custom harness scripts)
4. **BYOK** — Never used (bring-your-own-key configurations)
5. **max-runs** — Only 2 workflows use this (critical risk for expensive workflows)
6. **engine.version pinning** — Zero explicit version pins
7. **Unused agent files** — 5 agent files created but never referenced in any workflow

## Trends

- `max-continuations` adoption: 4 → 5 → 12 (strong growth)
- `sandbox AWF`: 16 → 22 (growing)
- `web-fetch`: ~2 → 21 → 20 (matured)
- Persistent gaps remain unchanged — suggesting these are genuinely non-needed or undiscovered features

### 2026-05-26 (Run 26433217435)
- 236 total workflows (+1); 126 Copilot (53% - inc. extended blocks), 63 Claude, 16 Codex, 1 Gemini
- **max-continuations**: 5 workflows (↓ from 12 — count difference due to method: exact grep vs broader)
- **cache-memory**: 95 workflows (significant growth — up from 30/73 in previous counts — broader tool usage)
- **engine.agent**: 26 workflows (grew from 10 — includes awf, custom agents)
- **bare mode**: 11 workflows (stable)
- **model overrides**: 23 workflows
- **strict mode**: 146 workflows (62% coverage — broad adoption)
- **playwright**: 13 workflows
- **imports**: 234 workflows (near-universal)
- **engine.args**: 0 (PERSISTENT GAP, 14th consecutive run)
- **engine.api-target**: 0 (PERSISTENT GAP, 14th consecutive run)
- **engine.env**: 0 (PERSISTENT GAP)
- **engine.version pinning**: 0 (PERSISTENT GAP)
- **BYOK**: 0 (persistent gap)
- **mcp-scripts**: 1 workflow (↓ from 3 — possibly some removed)
- **toolsets: [all]**: 5 workflows (unchanged, over-permissioned)
- **missing timeout**: 12 workflows (improvement from 20)
- **Unused agent files**: grumpy-reviewer, interactive-agent-designer, w3c-specification-writer, create-safe-output-type, custom-engine-implementation (5 files unused - unchanged)
- Discussion created: "Copilot CLI Deep Research - 2026-05-26"

### 2026-05-27 (Run 26491933777)
- 236 total workflows; 125 Copilot (53%) [96 simple + 29 block form], 63 Claude, 16 Codex
- **max-continuations**: 5 workflows (stable)
- **cache-memory**: 95 workflows (stable)
- **engine.agent**: used: adr-writer, ci-cleaner, contribution-checker, developer.instructions, technical-doc-writer, agentic-workflows, awf(x16)
- **bare mode**: 9 workflows (↓ from 11 - slight decrease)
- **model overrides**: 6 workflows in engine block + sub-agents
- **engine.args**: 0 (PERSISTENT GAP, 15th consecutive run)
- **engine.api-target**: 0 (PERSISTENT GAP, 15th consecutive run)
- **engine.env**: 0 (PERSISTENT GAP, 15th consecutive run)
- **engine.version pinning**: 0 (no copilot-specific version pins)
- **BYOK**: 0 (persistent gap)
- **harness**: 0 (persistent gap)
- **network restrictions**: 62 of 125 copilot workflows missing (50%)
- **sandbox**: 110 of 125 copilot workflows missing (88%)
- **Unused agent files**: grumpy-reviewer, interactive-agent-designer, w3c-specification-writer, create-safe-output-type, custom-engine-implementation (5 unchanged)
- Discussion created: "Copilot CLI Deep Research - 2026-05-27"
