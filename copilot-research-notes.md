# Copilot CLI Research Notes (Last 6 runs)

### 2026-05-20 (Run 26142362877) — This Run
- 231 total MD workflows (+1 from last run); 100 Copilot (43% - DOWN from 55%)
- **SIGNIFICANT DROP**: Copilot usage dropped from 126 (55%) to 100 (43%) — likely measurement refinement
- **max-continuations**: 5 workflows (stable: contribution-check:20, test-quality-sentinel:15, mattpocock-skills-reviewer:10, smoke:2, smoke-otel-backends:1)
- **max-runs**: 2 workflows (daily-safe-output-optimizer:200, linter-miner:1000) — SEVERELY underused feature
- **bare mode**: 11 workflows (stable pattern, mostly smoke tests and experimental workflows)
- **cache-memory**: 73 workflows (32% — stable, widely adopted)
- **repo-memory**: 23 workflows (10% — moderate but growing)
- **sandbox AWF**: 16 workflows (7% — firewall protection for untrusted input)
- **web-fetch**: 21 workflows (9% — up from 2, strong growth!)
- **web-search**: 2 workflows (minimal adoption)
- **model overrides**: 48 total (mostly "model: small", 3 "model: large")
- **playwright**: 11 workflows (browser automation growing)
- **engine.agent**: 11 workflows (custom agent files)
- **engine.args**: 0 (PERSISTENT GAP, 11th+ consecutive run with zero usage)
- **engine.env**: 0 (persistent gap)
- **engine.api-target**: 0 (PERSISTENT GAP, 12th consecutive run)
- **engine.harness**: 0 (persistent gap, zero workflows using custom harness)
- **engine.version pinning**: 0 (no explicit version pinning found)
- **mcp-scripts (frontmatter)**: 0 (used in prompts but not as frontmatter config)
- **BYOK**: 0 (persistent gap, no bring-your-own-key configurations)
- **experiments**: 16 workflows (A/B testing features)
- **Unused agent files**: grumpy-reviewer, interactive-agent-designer, w3c-specification-writer, create-safe-output-type, custom-engine-implementation (5 files unused)
- **Key findings**: Major gaps in advanced features (args, env, api-target, harness, BYOK, version pinning); max-runs severely underused despite being valuable for long-running workflows; web-fetch showing strong adoption growth

### 2026-05-18 (Run 26014468484)
- 230 total MD workflows (+1); 126 Copilot (114 simple + 29 extended block = 55%)
- **engine.agent**: 11 workflows (drop from 25 — volatility in count due to AWF vs engine field ambiguity)
- **max-continuations**: 5 workflows (stable)
- **bare mode**: 11 workflows (stable)
- **cache-memory**: 73 workflows (32% — stable)
- **repo-memory**: 23 workflows (10%)
- **sandbox AWF**: 16 workflows
- **web-search/fetch**: 21 workflows (up significantly from 2!)
- **model overrides**: 48 total (43 model:small, 3 model:large)
- **engine.args**: 0 (PERSISTENT GAP, 10th+ consecutive run)
- **engine.api-target**: 0 (PERSISTENT GAP, 11th consecutive run)
- **engine.harness**: 0
- **mcp-scripts (frontmatter)**: 0
- **BYOK**: 0
- **max-runs**: 2 workflows only
- Discussion created: "Copilot CLI Deep Research - 2026-05-18"

[Previous entries from 2026-05-17, 2026-05-16, 2026-05-14, 2026-05-13 available in git history]

---

## Key Persistent Gaps (Tracked Across All Runs)

1. **engine.args** — 11+ consecutive runs with ZERO usage (custom CLI arguments)
2. **engine.api-target** — 12 consecutive runs with ZERO usage (custom API endpoints)
3. **engine.harness** — Never used (custom harness scripts for engine wrapping)
4. **BYOK (Bring Your Own Key)** — Never used (custom provider credentials)
5. **max-runs** — Only 2 workflows use this despite being valuable for rate-limited long-running workflows
6. **engine.version pinning** — Zero explicit version pins (workflows use "latest" implicitly)
7. **mcp-scripts frontmatter** — Used in prompts but never configured in frontmatter
8. **Unused agent files** — 5 agent files created but never referenced

## Adoption Trends

- **Growing**: web-fetch (2→21), repo-memory (slow steady growth), playwright (browser automation)
- **Stable**: max-continuations (5), cache-memory (73), bare mode (11), sandbox (16)
- **Declining**: engine.agent volatility (25→11), possibly due to measurement changes
- **Stagnant**: engine.args, engine.env, engine.api-target, engine.harness, BYOK, max-runs

## High-Impact Recommendations

1. Explore max-runs for long-running workflows (daily generators, research agents)
2. Document engine.args patterns for common CLI customizations
3. Create examples of BYOK configurations for enterprise users
4. Investigate why 5 agent files remain unused (are they obsolete or undiscovered?)
5. Promote web-fetch adoption (successful growth from 2→21 shows demand)
