---
description: Demonstrates the `rate-limit` schema field
on:
  workflow_dispatch:
permissions:
  contents: read
engine: codex
rate-limit:
  window: 60
  events:
    - workflow_dispatch
timeout-minutes: 5
---

# Schema Demo: `rate-limit`

This workflow was auto-generated to demonstrate usage of the `rate-limit` field in
the gh-aw frontmatter schema. It exists solely to achieve 100% schema feature
coverage.

## What `rate-limit` Does

Legacy alias for `user-rate-limit`.

## Task

Call `noop` -- this is a coverage-only demo workflow.

**Important**: Always call the `noop` safe-output tool.

```json
{"noop": {"message": "Coverage demo for `rate-limit` -- no action needed."}}
```
