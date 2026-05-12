# Copilot CLI Research Notes (Trimmed - last 4 runs)

### 2026-05-11 (Run 25651194663) — This Run
- 218 total MD workflows; ~115 Copilot (95 simple + 20 block with id: copilot)
- **engine.agent**: 18 workflows (grew from 14 last run)
- **model overrides**: 13 workflows (grew from 5 last run - diverse models now)
- **max-continuations**: 2 (smoke-copilot:2, test-quality-sentinel:40) — persistent gap
- **engine.api-target**: 0 (persistent gap, 4th consecutive run)
- **harness script custom**: 0 (auto-applied; custom override unused)
- **BYOK/COPILOT_PROVIDER_***: 0 (persistent gap)
- **mcp-scripts**: 1/218 (only security-review.md) — severely underutilized
- **version pinning**: 2 (smoke-copilot + smoke-copilot-arm pin v1.25)
- **bare**: 9/218; **sandbox AWF**: 19/218; **strict: true**: 130/218 (60%)
- **cache-memory**: 89/218 (41%, stable)
- **network.allowed**: 115/218 (53%)
- Discussion created: "Copilot CLI Deep Research - 2026-05-11"

### 2026-05-10 (Run 25620196538)
- 218 total MD workflows; 96 Copilot (44%)
- **max-continuations**: 2; **engine.api-target**: 0; **engine.harness**: 0
- **engine.agent**: 14 workflows; **model:small**: 5 overrides
- **cache-memory**: 89/218 (massive growth); **sandbox AWF**: 19/218 (+73%)
- **strict: true**: 62/96 copilot (65%, stable)
- Discussion created: "Copilot CLI Deep Research - 2026-05-10"

### 2026-05-08 (Run 25537169013)
- 217 total MD workflows; 96 Copilot (44%)
- **max-continuations**: 2; **engine.api-target**: 0; **engine.harness**: 0
- **cache-memory**: 30/96; **sandbox AWF**: 11/96; **strict**: 62/96 (65%)
- Discussion created: "Copilot CLI Deep Research - 2026-05-08"

### 2026-05-06 (Run 25416993511)
- 214 total MD workflows; 93 Copilot (43%)
- max-continuations: 0; engine.agent: 7; cache-memory: 29/93; sandbox AWF: 11/93
- strict: 56/93 (60%); network: 45/93
- Discussion created: "Copilot CLI Deep Research - 2026-05-06"

### 2026-05-12 (Run 25714049123) — This Run
- 219 total MD workflows; 96 Copilot (44%)
- **max-continuations**: 4 (contribution-check:20, test-quality-sentinel:15, mattpocock-skills-reviewer:10, smoke-copilot:2) — growing!
- **model overrides**: 27 workflows — significant jump, driven by multi-agent `model: small` sub-agents
- **engine.agent**: 7 workflows (contribution-checker, adr-writer, technical-doc-writer x2, agentic-workflows, ci-cleaner, developer.instructions)
- **4 unused agent files**: grumpy-reviewer, interactive-agent-designer, w3c-specification-writer, create-safe-output-type
- **mcp-scripts**: 5 workflows (grew from 1) — accelerating
- **web-fetch**: 20 workflows (grew from 8)
- **engine.api-target**: 0 (persistent gap, 5th consecutive run)
- **engine.harness**: 0 (persistent gap)
- **BYOK**: 0 (persistent gap)
- **cache-memory**: 10 (different counting from previous; previous may have included imports)
- **sandbox AWF**: 20; **bare**: 9; **network.allowed**: 114
- Discussion created: "Copilot CLI Deep Research - 2026-05-12"
