---
description: Demonstrates the `run-install-scripts` schema field
on:
  workflow_dispatch:
permissions:
  contents: read
engine: codex
run-install-scripts: false
timeout-minutes: 5
---

# Schema Demo: `run-install-scripts`

This workflow was auto-generated to demonstrate usage of the
`run-install-scripts` field in the gh-aw frontmatter schema. It exists solely to
achieve 100% schema feature coverage.

## What `run-install-scripts` Does

Allows npm pre/post install scripts to execute during package installation.

## Task

Call `noop` -- this is a coverage-only demo workflow.

**Important**: Always call the `noop` safe-output tool.

```json
{"noop": {"message": "Coverage demo for `run-install-scripts` -- no action needed."}}
```
