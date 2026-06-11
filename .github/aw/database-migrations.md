---
description: Reference pattern for reviewing database migration files on pull requests — detecting risky operations and posting a structured safety summary.
---

# Database Migration Review

Consult this file when creating an agentic workflow that reviews database migration files on pull requests.

## Core Principle: Read the Diff, Never Execute

The agent must **read and analyze** migration file content — it must never connect to a database or attempt to run migrations. All output goes through `safe-outputs`.

## Frontmatter Template

```yaml
engine: copilot
on:
  pull_request:
    types: [opened, synchronize, reopened]
    paths:
      - "db/migrate/**"
      - "migrations/**"
      - "**/*.sql"
      - "**/schema.rb"
      - "**/schema.prisma"
permissions:
  contents: read
  pull-requests: read
network:
  allowed:
    - defaults
    - github
tools:
  github:
    mode: gh-proxy
    toolsets: [default]
safe-outputs:
  add-comment:
    hide-older-comments: true
  noop:
```

## Agent Instructions Pattern

```markdown
Review the migration files changed in this pull request for safety and best practices.

1. **Identify migration files**: List every file under `db/migrate/`, `migrations/`, or matching `*.sql` that was added or modified.
2. **Analyze each migration** for the following risk categories:
   - **Data-loss risk**: `DROP TABLE`, `DROP COLUMN`, `TRUNCATE`, destructive `UPDATE`/`DELETE` without `WHERE`
   - **Lock risk**: `ADD COLUMN NOT NULL` without a default, large table `ALTER`, index creation without `CONCURRENTLY`
   - **Irreversibility**: migrations lacking a `down` or rollback method
   - **Data type changes**: narrowing casts or charset changes that could truncate existing data
   - **Missing index**: foreign key columns added without a corresponding index
3. **Post a comment** with a structured summary:
   - 🟢 Safe — no concerns detected
   - 🟡 Review recommended — low-risk issues or best-practice suggestions
   - 🔴 Attention needed — high-risk operations identified
4. **Use `noop`** if no migration files were changed in this pull request.
```

## Risk Reference

| Operation | Risk Level | Guidance |
|---|---|---|
| `DROP TABLE` / `DROP COLUMN` | 🔴 High | Verify no application code still references the removed object. Prefer soft-delete or feature-flag removal first. |
| `ADD COLUMN NOT NULL` (no default) | 🔴 High | Acquires a full table lock on most engines. Add a `DEFAULT` or use a multi-step migration. |
| `CREATE INDEX` (without `CONCURRENTLY`) | 🟡 Medium | Locks the table for reads and writes. Use `CREATE INDEX CONCURRENTLY` on PostgreSQL. |
| `ALTER COLUMN` type change | 🟡 Medium | Check for implicit casts; narrowing casts can silently truncate data. |
| Foreign key without index | 🟡 Medium | Can cause slow lookups and lock contention. Add an index on the FK column. |
| Missing rollback / `down` method | 🟡 Medium | Without a rollback path, recovery from a failed deployment is manual. |
| `TRUNCATE` | 🔴 High | Irreversible data loss — only acceptable in non-production environments. |
| Bulk `UPDATE` / `DELETE` without `WHERE` | 🔴 High | Confirm the intent; suggest batching to reduce lock duration. |

## Key Design Decisions

- **`paths:` filter** — triggers only when migration files change; avoids false positives on doc-only PRs
- **`gh-proxy` mode** — all GitHub reads are safe-output-routed; no direct API credentials in the agent
- **`hide-older-comments: true`** — collapses the previous review when the PR is updated with a new commit
- **`noop` safe-output** — explicit exit path when no migration files are present
- **Read-only permissions** — the agent never writes to the repository; all output is a PR comment

## When to Add Ecosystem Network Access

If the workflow prompt needs to fetch schema documentation or tool documentation from the internet (rare), add the appropriate ecosystem to `network.allowed`. Most migration reviews require only `defaults` and `github`.

See also: [workflow-patterns.md](workflow-patterns.md) — Database Migration Safety Pattern
