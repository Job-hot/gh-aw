# Workflow Health — 2026-04-30T12:18Z

Score: 77/100 (↑+4 from 73 Apr 29). 205 workflows. Run: §25164831125

## KEY FINDINGS

### Compilation Status
- 205/205 lock files present ✅
- 0 missing lock files ✅

### Today's Failures (Apr 30 - morning)
7 scheduled runs failed out of 30 (77% success rate — slight improvement from 73% yesterday)

**Category 1: Codex engine crash (P0 ongoing)**
- **Daily Fact About gh-aw** — `codex: command not found` (#29088)

**Category 2: Safe outputs failures (6 workflows - P2 recurring)**
- Daily Rendering Scripts Verifier — safe_outputs failed
- Developer Documentation Consolidator — safe_outputs failed
- Daily AstroStyleLite Markdown Spellcheck — safe_outputs failed
- Instructions Janitor — safe_outputs failed
- Daily AW Cross-Repo Compile Check — safe_outputs failed
- Daily Community Attribution Updater — safe_outputs failed (new today)

### P0 Issues (Active)
- **Codex binary missing** (#29088): Daily Fact About gh-aw fails daily

### P1 Issues (Carry forward)
- **CI integration tests failing** (Apr 29): ERR_API fetch audit-workflows.md
- **Documentation Unbloat claude auth** (#28659) OPEN
- **GitHub Remote MCP Auth Test** (#27965) OPEN day 9
- **MCP gateway session timeout** (#23153) OPEN
- **awf-api-proxy sidecar unhealthy** (#27888) OPEN
- **GitHub App rate limit** (#27251) OPEN
- **CODEX_HOME collision** (#27512) OPEN

### P2 Issues
- **Safe Outputs SEC-004** (#27235)
- **Node.js 20 deprecation** in CI (deadline Sep 16, 2026)

### Dashboard Issue
- Apr 29 dashboard: #29108 (expiring Apr 30 12:29 UTC)
- Apr 30 dashboard: created this run
