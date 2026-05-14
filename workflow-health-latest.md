# Workflow Health — 2026-05-14T05:41Z

Score: 62/100 (↓ -1). 225 workflows. Run: §25843932794

## KEY FINDINGS

### Compilation Status
- 225/225 lock files present ✅ (2 new workflows since last run: 223→225)
- 0 missing lock files ✅
- Compile validation unavailable (gh aw extension not in PATH)

### Active Failures Today (2026-05-14)
- **CGO/CJS**: failing on every push to main — regression ongoing (P1, issue #29669 open)
- **Safe Output Health Monitor** (#32063): failed — single occurrence (was success 9/10)
- **Daily Grafana OTel Instrumentation Advisor** (#32066): failed — isolated
- **PR Sous Chef** (#32062): failed — auto-fix in progress via copilot
- **Daily Firewall Logs Collector** (#32045): failed — ongoing
- **Daily Observability Report** (#32035): failed — ongoing
- **Daily Model Inventory Checker** (#32032): failed — ongoing

### Open [aw] Failure Issues: 30 total
- Many from last 24h: AI Moderator, Design Decision Gate, Auto-Triage Issues, etc.
- Issue #31729 (deep-report triage of stale/duplicate aw failures): open

### Active Issues (P1)
- Daily Fact parse failures (#31432, #31524): ongoing
- MCP gateway session timeout (#23153): open
- Performance Regression (#30180): 82.1% slower
- CGO/CJS regression: every push to main fails

### P2 Issues (Watch)
- PR-review cluster (~272 wasted/day): #31724 filed
- Security findings: #31708, #31704 open
- Node.js 20 deprecation: Sep 16, 2026
- aw-portfolio-yield defect (#31456): open

### Actions Taken This Run
- Added comment to dashboard issue #29109
- Updated shared memory
- Added comment to CGO issue #29669

### Trends
- Score: 62/100 (↓ -1 from 63)
- P0 count: 0 ✅
- 225 workflows (was 223 — 2 new workflows)
- CGO/CJS failing on all pushes: persistent regression
- 30 open [aw] failure issues (was ~18)
