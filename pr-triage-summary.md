# PR Triage Summary - 2026-05-21T18:51:11Z

### Current Status

- **Total Bot PRs Found**: 9
- **Fork PRs Found**: 0
- **Branch-based PRs**: 9 (skipped per fork-only policy)
- **PRs Triaged**: 0
- **GitHub CLI Status**: ✅ Working

### Fork-Only Policy

Per workflow instructions, this triage agent only processes **fork PRs** where the head repository differs from the base repository (`isCrossRepository: true`).

All 9 open bot PRs are **branch-based** (created from branches within `github/gh-aw` rather than from external forks):

1. #33841 - [WIP] Fix model flag position in Codex lock compiler command
2. #33840 - [WIP] Fix branch link missing in fallback issue for protected files
3. #33839 - [WIP] Emit langfuse.session.id and langfuse.user.id for views
4. #33834 - [linter-miner] Add file-close-not-deferred linter
5. #33831 - Normalize context cancellation checks to `errors.Is` in CLI paths
6. #33830 - feat: extract cache-memory trending pattern into shared component
7. #33829 - Consolidate npm-not-available error surface across wasm and non-wasm builds
8. #33818 - pkg/cli: replace bare fmt.Println/Printf with explicit os.Stdout writes; add CI guard
9. #33219 - Bind Node toolcache into AWF chroot for Copilot-engine workflow startup reliability

### Conclusion

**No fork PRs available for automated triage.** All agent-created PRs are branch-based and excluded from this workflow's scope.

### Next Steps

- Monitor for new fork PRs from external contributors
- Branch-based PRs require manual review (not handled by this automated triage)
- Re-run triage on next scheduled interval

---
*Last Updated: 2026-05-21T18:51:11Z | Run: 26246351754*
