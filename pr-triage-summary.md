# PR Triage Summary

## Latest Run: 2026-05-21T01:18:55Z

**Status**: ❌ System constraint - gh CLI authentication failure

**PRs Triaged**: 0  
**Fork PRs Found**: 0  
**Authentication Error**: HTTP 403

## Issue

gh CLI is not authenticated in safe-output workflows, despite workflow documentation stating "gh CLI is pre-authenticated". The runtime environment returns HTTP 403 when attempting to access GitHub API.

**Error Details:**
- Error Code: HTTP 403: 403 Forbidden
- Endpoint: https://localhost:18443/api/v3/meta
- Consecutive Failures: 6 runs

## Recommendations

1. **Fix gh CLI authentication** in safe-output workflow environment
2. **Provide alternative method** to fetch PR data (GitHub MCP tools, authenticated API)
3. **Update workflow documentation** if gh CLI authentication is intentionally disabled
4. **Remove fork-only restriction** for github/gh-aw if same-repo PRs should be triaged

## Historical Failures

- 26199592140 (2026-05-21 01:18 UTC): HTTP 403 ← Current run
- 26183688727 (2026-05-20 19:03 UTC): HTTP 403
- 26148138573 (2026-05-20 07:32 UTC): HTTP 403
- 26135250674: HTTP 403
- 26118240951: HTTP 403
- 26082965672: HTTP 403

## Next Steps

This workflow cannot function without authenticated GitHub API access. The issue requires infrastructure-level resolution.
