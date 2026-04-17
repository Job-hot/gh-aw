---
title: Repo Memory
description: Guide to using repo-memory for persistent file storage via Git branches with unlimited retention.
sidebar:
  order: 1510
---

Repo memory provides persistent file storage via Git branches with unlimited retention. The compiler auto-configures branch cloning/creation, file access at `/tmp/gh-aw/repo-memory-{id}/`, commits/pushes, and merge conflict resolution (your changes win).

## Enabling Repo Memory

```aw wrap
---
tools:
  repo-memory: true
---
```

Creates branch `memory/default` at `/tmp/gh-aw/repo-memory-default/`; files auto-commit/push after workflow completion.

## Advanced Configuration

```aw wrap
---
tools:
  repo-memory:
    branch-name: memory/custom-agent-for-aw
    branch-prefix: tracking  # Custom prefix instead of "memory"
    description: "Long-term insights"
    file-glob: ["*.md", "*.json"]
    max-file-size: 1048576  # 1MB (default 10KB)
    max-file-count: 50      # default 100
    max-patch-size: 102400  # 100KB max (default 10KB)
    target-repo: "owner/repository"
    create-orphan: true     # default
    allowed-extensions: [".json", ".txt", ".md"]  # Restrict file types (default: empty/all files allowed)
---
```

Notable fields:
- **`branch-prefix`**: Branch name prefix (default: `memory`). Must be 4–32 alphanumeric characters with hyphens/underscores; cannot be `copilot`.
- **`allowed-extensions`**: Restricts storable file types. Files with unlisted extensions trigger validation failures.
- **`max-patch-size`**: Maximum total diff size per push (default: 10KB, max: 100KB). Exceeding this rejects the push.
- **`file-glob`**: Matched against **relative paths** within the artifact directory — not the branch name (e.g. use `*.json`, not `memory/custom/*.json`).

## Multiple Configurations

```aw wrap
---
tools:
  repo-memory:
    - id: insights
      branch-prefix: daily  # Creates daily/insights branch
      file-glob: ["*.md"]
    - id: state
      file-glob: ["*.json"]
      max-file-size: 524288  # 512KB
---
```

Each entry mounts at `/tmp/gh-aw/repo-memory-{id}/`. The required `id` field determines the folder name; `branch-name` defaults to `{branch-prefix}/{id}`.

## Comparison with Cache Memory

| Feature | Cache Memory | Repo Memory |
|---------|--------------|-------------|
| Storage | GitHub Actions Cache | Git Branches |
| Retention | 7 days | Unlimited |
| Size Limit | 10GB/repo | Repository limits |
| Version Control | No | Yes |
| Performance | Fast | Slower |
| Best For | Temporary/sessions | Long-term/history |

For fast 7-day caching without version control, see [Cache Memory](/gh-aw/reference/cache-memory/).

## Troubleshooting

- **Branch not created**: Ensure `create-orphan: true` or create manually.
- **Validation failures**: Match `file-glob`, stay under `max-file-size` (10KB default), `max-file-count` (100 default), and `max-patch-size` (10KB default).
- **Patch too large**: If the total diff exceeds `max-patch-size` (default 10KB), the push is rejected. Reduce the number or size of changes, or increase `max-patch-size` in the configuration.
- **Changes not persisting**: Check directory path, workflow completion, push errors in logs.
- **Merge conflicts**: Uses `-X ours` (your changes win). Read before writing to preserve data.

## Security

Don't store sensitive data; repo memory follows repository permissions. Use private repos, set constraints, and `target-repo` to isolate sensitive branches.

## Examples

See [Deep Report](https://github.com/github/gh-aw/blob/main/.github/workflows/deep-report.md) and [Daily Firewall Report](https://github.com/github/gh-aw/blob/main/.github/workflows/daily-firewall-report.md) for long-term insights and historical data tracking.

## Related Documentation

- [Cache Memory](/gh-aw/reference/cache-memory/) - GitHub Actions cache-based storage with 7-day retention
- [Frontmatter](/gh-aw/reference/frontmatter/) - Complete frontmatter configuration guide
- [Safe Outputs](/gh-aw/reference/safe-outputs/) - Output processing and automation
