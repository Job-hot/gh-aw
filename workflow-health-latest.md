# Workflow Health - 2026-04-16T12:10Z

Score: 71/100 (↓1 from 72 Apr 15). 191 workflows. Run: §24509370401

## KEY FINDINGS

### Stale Lock Files
- 16/191 stale lock files (down from 18 Apr 15 — some recompiled)
- Same-second timestamps suggest checkout artifacts, not real staleness
- List: archie, brave, craft, daily-cli-performance, daily-file-diet, daily-integrity-analysis, daily-multi-device-docs-tester, daily-performance-summary, daily-safe-output-optimizer, docs-noob-tester, firewall-escape, github-mcp-tools-report, refiner, research, test-project-url-default, ubuntu-image-analyzer

### P0 Persistent Failures (Unchanged)
- **Daily Issues Report Generator**: 10+ consecutive day streak (#26393 OPEN)
  - node:command not found pattern in Copilot runner
- **Smoke Gemini**: 14+ day 100% failure (#26351 OPEN, #26456 deep-report OPEN)
  - Gemini API proxy handler crash

### Intermittent Failures Today (Apr 16)
- ~18 workflows failed in sample of ~35 (50% rate)  
- Most are one-off: Instructions Janitor (1/5), AI Moderator (2x issues-triggered)
- Auto-Triage Issues: recovered (failed 01:07Z, success 07:04Z)
- High Copilot engine flakiness suspected

### P1 Watch — Smoke Claude
- ~60% success rate recent schedule runs (3 fail in last 5)
- No open issue — monitor

### Successful Today (sample)
- Daily Rendering Scripts Verifier, Update Astro, Agentic Maintenance, Sub-Issue Closer
- Daily Syntax Error Quality Check, Package Specification Extractor
- Issue Monster, Glossary Maintainer, Content Moderation, Terminal Stylist

## Compilation Status
- 191/191 lock files present ✅
- 16 stale lock files ⚠️ (checkout artifacts, down from 18)

## Open Issues (workflow-related)
- #26393 Daily Issues Report (P0)
- #26351 Smoke Gemini (P0)
- #26456 Gemini proxy handler deep-report (P0)
- #26239 MCP rate-limit circuit breaker (P2)
- #26458 GitHub MCP get_me 403 errors (P2)

## Actions This Run
- Created: [workflow-health] Dashboard Apr 16 issue (#aw_whdapr16)
- No new P0/P1 issues created (all already tracked)

## Engine/Tool Status
- Copilot v1.0.27 available (was v1.0.21 active) → #26367 open
- Claude Code 2.1.109 available
- Gemini 0.38.0 available (not yet deployed)
