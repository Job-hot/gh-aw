---
title: Multi-Repository Examples
description: Complete examples for managing workflows across multiple GitHub repositories, including feature synchronization, cross-repo tracking, and organization-wide updates.
---

Multi-repository operations coordinate work across GitHub repositories while preserving security boundaries.

## Featured Examples

- **[Feature Synchronization](/gh-aw/examples/multi-repo/feature-sync/)** — Sync code from a source repo to downstream repos via pull requests, with change detection, path filters, and bidirectional support.
- **[Cross-Repository Issue Tracking](/gh-aw/examples/multi-repo/issue-tracking/)** — Centralize tracking issues in a hub repo with status synchronization across components.

## Authentication

Cross-repo writes require credentials scoped to the **target** repositories only (not the source repo where the workflow runs). Two options:

**Personal Access Token (PAT)** — fastest to set up:

```bash
gh aw secrets set GH_AW_CROSS_REPO_PAT --value "ghp_your_token_here"
```

Grant the minimum write scopes needed (`contents`, `issues`, `pull-requests`) on target repos only.

**GitHub App** — preferred for production: tokens are minted on-demand and revoked after each job. See [Using a GitHub App for Authentication](/gh-aw/reference/auth/#using-a-github-app-for-authentication) for repo-scoped and org-wide setup.

## Cross-Repository Safe Outputs

Most safe output types support the `target-repo` parameter for cross-repository operations. **Without `target-repo`, these safe outputs operate on the repository where the workflow is running.**

| Safe Output | Cross-Repo Support | Example Use Case |
|-------------|-------------------|------------------|
| `create-issue` | ✅ | Create tracking issues in central repo |
| `add-comment` | ✅ | Comment on issues in other repos |
| `update-issue` | ✅ | Update issue status across repos |
| `add-labels` | ✅ | Label issues in target repos |
| `create-pull-request` | ✅ | Create PRs in downstream repos |
| `create-discussion` | ✅ | Create discussions in any repo |
| `create-agent-session` | ✅ | Create tasks in target repos |
| `update-release` | ✅ | Update release notes across repos |
| `update-project` | ✅ (`target_repo`) | Update project items from other repos |

**Configuration Example:**

```yaml wrap
safe-outputs:
  github-token: ${{ secrets.GH_AW_CROSS_REPO_PAT }}
  create-issue:
    target-repo: "org/tracking-repo"  # Cross-repo: creates in tracking-repo
    title-prefix: "[component] "
  add-comment:
    # No target-repo: operates on current repository
```

## GitHub API Tools for Multi-Repo Access

Enable the toolsets your agent needs to query other repositories:

```yaml wrap
tools:
  github:
    toolsets: [repos, issues, pull_requests, actions]
```

See [GitHub Tools](/gh-aw/reference/github-tools/) for the per-toolset capability list.

## Direct Checkout for Deterministic Workflows

When the agent needs direct file access rather than API calls, check out the secondary repo as a workflow step:

```yaml wrap
engine:
  id: claude
  steps:
    - name: Checkout main repo
      uses: actions/checkout@v6
      with:
        path: main-repo

    - name: Checkout secondary repo
      uses: actions/checkout@v6
      with:
        repository: org/secondary-repo
        token: ${{ secrets.GH_AW_CROSS_REPO_PAT }}
        path: secondary-repo
```

## See Also

- [MultiRepoOps Design Pattern](/gh-aw/patterns/multi-repo-ops/) — full pattern documentation
- [Cross-Repository Operations](/gh-aw/reference/cross-repository/) — checkout and `target-repo` configuration
- [Authentication](/gh-aw/reference/auth/) — PAT and GitHub App setup
- [Reusing Workflows](/gh-aw/guides/packaging-imports/) — sharing imports across repos
