# PR Triage Summary

## Latest Run: 2026-05-20T19:03:14Z

**Status**: ⚠️ System constraint - gh CLI not authenticated

**PRs Triaged**: 0  
**Fork PRs Found**: 0  
**Authentication Error**: HTTP 403

## Issue

gh CLI is not authenticated in safe-output workflows, preventing PR data retrieval.

## Recommendations

1. Add GitHub read tools to safe-output workflows
2. Remove fork-only restriction for github/gh-aw
3. Use different workflow environment with gh auth

## Historical Failures

- 26183688727 (2026-05-20 19:03 UTC): HTTP 403
- 26148138573 (2026-05-20 07:32 UTC): HTTP 403
- 26135250674: HTTP 403
- 26118240951: HTTP 403
- 26082965672: HTTP 403
