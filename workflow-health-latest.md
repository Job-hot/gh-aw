# Workflow Health — 2026-06-11T06:07Z

Score: 85/100 (↓2 from 87)
Workflows: 245 | Lock files: 245/245 (100% ✅) | Run: §27327316752

## KEY FINDINGS

### Status (June 11)
- **Compilation:** 245/245 workflows have lock files (100% ✅)
- **AI Credits Cluster Day 3+**: Code Simplifier (#38499), Ubuntu Analyzer (#38500), Workflow Skill Extractor (#38501), Matt Pocock (#38497) — all hitting AI credits limits. New tracker: #aw_aic11
- **Daily Safe Outputs Git Simulator** (P1, ongoing, #aw_gitsim10): memory/git-simulator branch missing signed-commit seed — DO NOT RE-FILE
- **Daily News** (P1, ongoing, #38379): Node.js unreachable in AWF chroot — assigned @zarenner — DO NOT RE-FILE

### Critical Issues (P0/P1) 🚨
- **AI Credits Cluster** (#aw_aic11, new June 11): 4 workflows — Code Simplifier (1K), Ubuntu Analyzer (1K), Workflow Skill Extractor (1K), Matt Pocock (5K daily). Fix: raise max-ai-credits to 2000 for analysis-heavy workflows.
- **Daily Safe Outputs Git Simulator** (#aw_gitsim10, ongoing): memory/git-simulator branch missing. DO NOT RE-FILE.
- **Daily News** (#38379, ongoing): Node.js chroot issue. DO NOT RE-FILE.

### Warnings (P2) ⚠️
- **jsweep** (#38505): Tool denial (rm -rf), guardrail working, one-off occurrence. Monitor.

### Resolved Since Last Run ✅
- Many June 10 individual failures auto-closed after AI credits reset
- Failure cascade #38389 — CLOSED
- Several daily workflow failures closed

### P2/P3 Open Issues (From Other Workflows)
- #38145: Consolidate AI credits cluster under batch-close (quick-win)
- #38144: Enable cancel-in-progress for PR-gating reviewers (quick-win)
- #38379: Daily News chroot bug (Node.js missing)
- #38434: Ambient Context Optimizer report (non-critical)

### Systemic Patterns
- **AI credits 1K limit**: Analysis-heavy workflows routinely exceed 1K max-ai-credits; need 2K
- **memory/* bootstrap**: New memory branches need signed-commit seed on first push
- **Daily News chroot**: Node.js PATH resolution failing in AWF chroot environment

### Actions Taken This Run
- 1 issue created: #aw_aic11 (AI Credits Cluster June 11 — P1, 4 workflows)
- 1 comment added to #29109 (health dashboard)

## Do Not Re-File
- #aw_aic11: AI Credits Cluster June 11 (4 workflows: Code Simplifier, Ubuntu, Skill Extractor, Matt Pocock)
- #aw_gitsim10: Daily Safe Outputs Git Simulator (memory/git-simulator branch needs seed)
- #38379: Daily News Node.js chroot (assigned @zarenner)
- #38499: Code Simplifier failed
- #38500: Ubuntu Actions Image Analyzer failed
- #38501: Workflow Skill Extractor failed
- #38497: Matt Pocock Skills Reviewer failed
- #38505: jsweep tool denial
