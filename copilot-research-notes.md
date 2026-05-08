# Copilot CLI Research Notes (Trimmed - last 2 runs)

### 2026-05-08 (Run 25537169013) — This Run
- 217 total MD workflows; 96 Copilot (44%)
- **max-continuations**: 2 (smoke-copilot, test-quality-sentinel only)
- **engine.api-target**: 0 (persistent gap)
- **engine.harness**: 0 (persistent gap)  
- **engine.version pinning**: 0
- **engine.model**: 3 overrides
- **engine.bare**: 9 total
- **cache-memory**: 30/96 (31%)
- **comment-memory**: 3/96 (3%)
- **mcp-scripts**: 1/96 (1%)
- **sandbox AWF**: 11/96; **network config**: 45/96
- **strict: true**: 62/96 (65%)
- **bash without network**: 30 workflows — security gap
- **safe-outputs without strict**: 28 workflows — injection risk
- Discussion created: "Copilot CLI Deep Research - 2026-05-08"

### 2026-05-06 (Run 25416993511)
- 214 total MD workflows; 93 Copilot (simple form)
- max-continuations: 0 in copilot wfs; engine.model: 12 overrides
- engine.agent: 7; cache-memory: 29/93; repo-memory: 18/93
- sandbox AWF: 11/93; strict: true: 56/93 (60%); network: 45/93
- Discussion created: "Copilot CLI Deep Research - 2026-05-06"
