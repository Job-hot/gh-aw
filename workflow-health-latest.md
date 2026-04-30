# Workflow Health — 2026-04-30T23:33Z (PM update)

Score: 72/100 (↓-5 from morning 77). 205 workflows. Run: §25194518483

## KEY FINDINGS

### Compilation Status
- 205/205 lock files present ✅
- 0 missing lock files ✅

### Today's Full Day Failures (Apr 30)
AM batch (10:43-11:51): 8 failures - codex crash + 7 safe_outputs
Afternoon: GitHub MCP Structural Analysis (agent timeout, claude killed)
           Super Linter Report (Node.js crash)
           CI: 6 failures (integration tests throughout day)
Evening: CJS test failure (CI-related)

**Category 1: Codex engine crash (P0 ongoing)**
- Daily Fact About gh-aw — `codex: command not found` (#29088)

**Category 2: Safe outputs failures (7 workflows)**
- Daily Rendering Scripts Verifier
- Developer Documentation Consolidator
- Daily AstroStyleLite Markdown Spellcheck
- Instructions Janitor
- Daily AW Cross-Repo Compile Check
- Daily Community Attribution Updater
- Daily Issues Report Generator (new)

**Category 3: CI failures (P1 escalated)**
- CI integration tests: 6 failures today (50% of CI runs)
- CJS tests: 1 failure

**Category 4: Agent timeout (new)**
- GitHub MCP Structural Analysis — claude process killed

**Category 5: Super Linter (P2)**
- Super Linter Report failure - Node.js runtime crash

### P0 Issues (Active)
- **Codex binary missing** (#29088): Daily Fact About gh-aw fails daily

### P1 Issues (Carry forward + CI escalated)
- **CI integration tests failing** (6 runs today 50% fail rate) — ESCALATED
- **Documentation Unbloat claude auth** (#28659) OPEN
- **GitHub Remote MCP Authentication Test** (#27965) OPEN day 10
- **MCP gateway session timeout** (#23153) OPEN
- **awf-api-proxy sidecar unhealthy** (#27888) OPEN
- **GitHub App rate limit** (#27251) OPEN
- **CODEX_HOME collision** (#27512) OPEN

### P2 Issues
- **Safe Outputs SEC-004** (#27235)
- **Node.js 20 deprecation** in CI (deadline Sep 16, 2026)

### Dashboard Issue
- Apr 30 AM dashboard: #29304 (updated this run)
