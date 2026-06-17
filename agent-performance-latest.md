# Agent Performance — Latest Run

**Timestamp:** 2026-06-17T14:00:00Z | **Run:** [§27694092266](https://github.com/github/gh-aw/actions/runs/27694092266)

## Summary: 55/100 Quality | 53/100 Effectiveness | 66/100 Health | AIC crisis Day 10

## Top Performers
1. Bot Detection (Q:82, E:91) — 100% success, ~13s
2. Agentic Maintenance (Q:82, E:88)
3. Auto-Triage Issues (Q:81, E:84) — consistent
4. Issue Monster (Q:78, E:85)
5. copilot-swe-agent (Q:78, E:82, 67% merge rate — 30/45 PRs)
6. Dependabot (E:100% — 1/1 today)
7. Content Moderation (Q:76, E:80, 88% success recent)
8. PR Sous Chef — ✅ FULLY RECOVERED
9. Avenger — Healthy
10. AI Moderator — Recovering (consecutive successes Jun 17)

## Underperformers
- Code Simplifier (Q:10, E:5) P1 Day 10: api-proxy cap + HTTP 429. #39199/#39489/#39729. DO NOT RE-FILE.
- Daily Model Inventory (Q:35, E:25) Day 8: session.idle 60s (#39471). DO NOT RE-FILE.
- Dictation Prompt Generator (Q:30, E:20): session.idle 570s (#39196/#39200). DO NOT RE-FILE.
- Smoke Gemini (Q:25, E:15): preview model (#39172/#39559). DO NOT RE-FILE.
- GitHub Remote MCP Auth (Q:30, E:20): missing tool (#39193/#39505). DO NOT RE-FILE.
- Tool Denial Cluster (now 7+ workflows): SYSTEMIC P1. New member: Daily MCP Tool Concurrency Analysis (#39763). DO NOT RE-FILE.
- Incomplete Result Cluster NEW Jun 17: Package Spec Enforcer (#39787), Package Spec Extractor (#39762), Daily Documentation Updater (#39775), Daily Workflow Updater (#39753). Issue filed #aw_inc_result.

## New Patterns Today (Jun 17)
- Incomplete Result Cluster: 4 workflows returning incomplete — new systemic pattern (issue filed)
- Tool Denial Cluster expanding: Daily MCP Tool Concurrency Analysis (#39763) joins list
- Design Decision Gate: 2 issues filed (#39779 failed + #39776 no safe outputs) — same workflow
- New individual failures: Test Quality Sentinel (#39782), Matt Pocock Skills Reviewer (#39781), Glossary Maintainer (#39769), Daily News (#39758), Daily Compiler Quality Check (#39724), Metrics Collector Infrastructure Agent (#39727), Instructions Janitor (#39757)
- Metrics Collector failing (affects shared memory infrastructure) — needs monitoring

## PR Merge Rates (latest)
- dependabot: 100% (1/1 today)
- copilot-swe-agent: 67% (30/45 PRs in window) ↓ from 77%
- github-actions workflows: ~75%

## Do Not Re-File
#39199/#39489/#39729 Code Simplifier, #39352/#39330 PR Sous Chef (RESOLVED), #38998 upload_artifact, #39024 Git Simulator, #39471 Model Inventory, #39172 Smoke Gemini, #39193/#39505/#39757 Remote MCP Auth/Instructions Janitor, #39196/#39200 Dictation, #38999 Smoke Trigger, #39037 Failure Investigator, #38870-38872 Perf regression, #39077 AIC Crisis, #39502 Tool denial cluster (systemic), #39477 Safe Output Integrator, #39476 BYOK Ollama, #39452 AI Moderator, #39451 Cache Strategy Analyzer, #39343 Compiler Threat Spec Optimizer, #39560-39562 Smoke Pi/Codex/Antigravity, #39497 AIC Impact Efficiency, #39511 LintMonster, #39507 pr-code-quality-reviewer compile, #39787 Pkg Spec Enforcer, #39762 Pkg Spec Extractor, #39775 Daily Docs Updater, #39753 Daily Workflow Updater, #39782 Test Quality Sentinel, #39781 Matt Pocock, #39779/#39776 Design Decision Gate, #39769 Glossary Maintainer, #39758 Daily News, #39724 Daily Compiler Quality, #39727 Metrics Collector

## Issues Filed This Run
- 1 systemic issue: "Incomplete Result Cluster" (4 workflows, new pattern Jun 17)
