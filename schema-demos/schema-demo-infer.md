---
description: Demonstrates the `infer` schema field
on:
  workflow_dispatch:
permissions:
  contents: read
engine: codex
infer: false
timeout-minutes: 5
---

# Schema Demo: `infer`

This workflow was auto-generated to demonstrate usage of the `infer` field in the
gh-aw frontmatter schema. It exists solely to achieve 100% schema feature coverage.

## What `infer` Does

Maintains backward compatibility for the deprecated model-invocation control field.

## Task

Call `noop` -- this is a coverage-only demo workflow.

**Important**: Always call the `noop` safe-output tool.

```json
{"noop": {"message": "Coverage demo for `infer` -- no action needed."}}
```
