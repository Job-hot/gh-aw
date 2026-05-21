# PR Triage Summary - 2026-05-21

## Fork-Only Policy Applied

Per workflow instructions, this triage agent only processes **fork PRs** (PRs where `head.repo.full_name` differs from `base.repo.full_name`).

### Current Status

- **Total Open Agent PRs**: 7
- **Fork PRs**: 0
- **Branch-based PRs (excluded)**: 7
- **PRs Triaged**: 0

### Excluded PRs (Branch-based)

All 7 agent PRs are branch-based (opened from branches within `github/gh-aw`):

1. #33716 - Deduplicate workflow expression regex usage
2. #33702 - Fix pull_request_reviewer double-trigger  
3. #33687 - refactor(create_pull_request): extract helpers
4. #33685 - Standardize pkg debug logging
5. #33683 - refactor(parser): break up oversized functions
6. #33664 - chore: bump default gh-aw-mcpg
7. #33219 - Bind Node toolcache into AWF chroot

### Historical Context

Previous runs (6 consecutive) encountered HTTP 403 authentication errors. This run successfully fetched PR data but found no fork PRs to triage.

### Next Run

The triage agent will continue monitoring for fork PRs. Branch-based PRs should be handled by standard review processes.

---
*Last Updated: 2026-05-21T07:36:00Z*
