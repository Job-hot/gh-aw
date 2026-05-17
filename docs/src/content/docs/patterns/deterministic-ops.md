---
title: DeterministicOps
description: Combine deterministic computation and data extraction with agentic reasoning in GitHub Agentic Workflows for powerful hybrid automation.
sidebar:
  order: 6
  badge: { text: 'Hybrid', variant: 'caution' }
---

Combine deterministic shell commands with AI reasoning to precompute data, filter triggers, preprocess inputs, post-process outputs, or build multi-stage pipelines. The **DataOps** sub-pattern uses `steps:` to reliably collect and prepare data — fast, cacheable, reproducible — then lets the AI agent generate insights from the results. Common uses: data aggregation, report generation, trend analysis, and auditing.

## Architecture

```text
┌────────────────────────┐
│  Deterministic Jobs    │
│  - Data fetching       │
│  - Preprocessing       │
└───────────┬────────────┘
            │ artifacts/outputs
            ▼
┌────────────────────────┐
│   Agent Job (AI)       │
│   - Reasons & decides  │
└───────────┬────────────┘
            │ safe outputs
            ▼
┌────────────────────────┐
│  Safe Output Jobs      │
│  - GitHub API calls    │
└────────────────────────┘
```

## Precomputation Example

```yaml wrap title=".github/workflows/release-highlights.md"
---
on:
  push:
    tags: ['v*.*.*']
engine: copilot
safe-outputs:
  update-release:

steps:
  - run: |
      gh release view "${GITHUB_REF#refs/tags/}" --json name,tagName,body > /tmp/gh-aw/agent/release.json
      gh pr list --state merged --limit 100 --json number,title,labels > /tmp/gh-aw/agent/prs.json
    env:
      GH_TOKEN: ${{ secrets.GITHUB_TOKEN }}
---

# Release Highlights Generator

Generate release highlights for `${GITHUB_REF#refs/tags/}`. Analyze PRs in `/tmp/gh-aw/agent/prs.json`, categorize changes, and use update-release to prepend highlights to the release notes.
```

Files in `/tmp/gh-aw/agent/` are automatically uploaded as artifacts and available to the AI agent.

## Multi-Job Precomputation

```yaml wrap title=".github/workflows/static-analysis.md"
---
on:
  schedule: daily
engine: claude
safe-outputs:
  create-discussion:

jobs:
  run-analysis:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v6
      - run: ./gh-aw compile --zizmor --poutine > /tmp/gh-aw/agent/analysis.txt

steps:
  - uses: actions/download-artifact@v6
    with:
      name: analysis-results
      path: /tmp/gh-aw/
---

# Static Analysis Report

Parse findings in `/tmp/gh-aw/agent/analysis.txt`, cluster by severity, and create a discussion with fix suggestions.
```

Pass data between jobs via artifacts, job outputs, or environment variables.

## Custom Trigger Filtering

### Inline Steps (`on.steps:`) — Preferred

For lightweight filtering, inject deterministic steps directly into the pre-activation job with `on.steps:` — this saves one workflow job versus a separate filter job:

```yaml wrap title=".github/workflows/smart-responder.md"
---
on:
  issues:
    types: [opened]
  steps:
    - id: check
      env:
        LABELS: ${{ toJSON(github.event.issue.labels.*.name) }}
      run: echo "$LABELS" | grep -q '"bug"'
      # exits 0 (outcome: success) if the label is found, 1 (outcome: failure) if not
engine: copilot
safe-outputs:
  add-comment:

if: needs.pre_activation.outputs.check_result == 'success'
---

# Bug Issue Responder

Triage bug report: "${{ github.event.issue.title }}" and add-comment with a summary of the next steps.
```

Each step with an `id` gets an auto-wired `<id>_result` output (`success` when exit code is 0, `failure` otherwise) — gate the workflow with `needs.pre_activation.outputs.<id>_result == 'success'`.

To pass an explicit value instead of relying on exit codes, re-expose a step output via `jobs.pre-activation.outputs`:

```yaml wrap
jobs:
  pre-activation:
    outputs:
      has_bug_label: ${{ steps.check.outputs.has_bug_label }}

if: needs.pre_activation.outputs.has_bug_label == 'true'
```

When `on.steps:` need GitHub API access, use `on.permissions:` to grant the required scopes to the pre-activation job:

```yaml wrap
on:
  schedule: every 30m
  permissions:
    issues: read
  steps:
    - id: search
      uses: actions/github-script@v8
      with:
        script: |
          const open = await github.rest.issues.listForRepo({ ...context.repo, state: 'open' });
          core.setOutput('has_work', open.data.length > 0 ? 'true' : 'false');

jobs:
  pre-activation:
    outputs:
      has_work: ${{ steps.search.outputs.has_work }}

if: needs.pre_activation.outputs.has_work == 'true'
```

See [Pre-Activation Steps](/gh-aw/reference/triggers/#pre-activation-steps-onsteps) and [Pre-Activation Permissions](/gh-aw/reference/triggers/#pre-activation-permissions-onpermissions) for full documentation.

### Separate Filter Job

When filtering needs heavy tooling (checkout, compiled tools, multiple runners), use a separate `jobs:` entry instead of `on.steps:`:

```yaml wrap title=".github/workflows/smart-responder.md"
---
on:
  issues:
    types: [opened]
engine: copilot
safe-outputs:
  add-comment:

jobs:
  filter:
    runs-on: ubuntu-latest
    outputs:
      should-run: ${{ steps.check.outputs.result }}
    steps:
      - id: check
        env:
          LABELS: ${{ toJSON(github.event.issue.labels.*.name) }}
        run: |
          if echo "$LABELS" | grep -q '"bug"'; then
            echo "result=true" >> "$GITHUB_OUTPUT"
          else
            echo "result=false" >> "$GITHUB_OUTPUT"
          fi

if: needs.filter.outputs.should-run == 'true'
---

# Bug Issue Responder

Triage bug report: "${{ github.event.issue.title }}" and add-comment with a summary of the next steps.
```

The compiler wires the filter job as a dependency of the activation job, so when the condition is false the run is **skipped** rather than failed, keeping the Actions tab clean.

### Simple Context Conditions

For conditions that can be expressed directly with GitHub Actions context, use `if:` without a custom job:

```yaml wrap
---
on:
  pull_request:
    types: [opened, synchronize]
engine: copilot
if: github.event.pull_request.draft == false
---
```

### Query-Based Filtering

For conditions based on GitHub search results, use [`skip-if-match:`](/gh-aw/reference/triggers/#skip-if-match-condition-skip-if-match) or [`skip-if-no-match:`](/gh-aw/reference/triggers/#skip-if-no-match-condition-skip-if-no-match) in the `on:` section — these accept standard [GitHub search query syntax](https://docs.github.com/en/search-github/searching-on-github/searching-issues-and-pull-requests) and are evaluated in the pre-activation job, producing the same skipped-not-failed behavior:

```yaml wrap
---
on:
  issues:
    types: [opened]
  # Skip if a duplicate issue already exists (GitHub search query syntax)
  skip-if-match: 'is:issue is:open label:duplicate'
engine: copilot
---
```

## Post-Processing Pattern

```yaml wrap title=".github/workflows/code-review.md"
---
on:
  pull_request:
    types: [opened]
engine: copilot

safe-outputs:
  jobs:
    format-and-notify:
      description: "Format and post review"
      runs-on: ubuntu-latest
      inputs:
        summary: {required: true, type: string}
      steps:
        - ...
---

# Code Review Agent

Review the pull request and use format-and-notify to post your summary.
```

## Importing Shared Instructions

Define reusable guidance in shared files and import them:

```yaml wrap title=".github/workflows/analysis.md"
---
on:
  schedule: daily
engine: copilot
imports:
  - shared/reporting.md
safe-outputs:
  create-discussion:
---

# Daily Analysis

Follow the report formatting guidelines from shared/reporting.md.
```

For daily discussion-based audit workflows, prefer `shared/daily-audit-base.md` to bundle discussion publishing, reporting guidance, and OTLP observability in a single import.

## DataOps: Scheduled Data Extraction

For scheduled reporting, trend analysis, and auditing, collect data deterministically in `steps:` then let the agent analyze it.

### Example: Weekly PR Activity Summary

````aw wrap
---
on:
  schedule: weekly
engine: copilot
permissions:
  contents: read
  pull-requests: read
safe-outputs:
  create-discussion:
    title-prefix: "[weekly-summary] "
    category: "announcements"
    close-older-discussions: true

steps:
  - name: Fetch and aggregate PRs
    env:
      GH_TOKEN: ${{ secrets.GITHUB_TOKEN }}
    run: |
      mkdir -p /tmp/gh-aw/pr-data
      gh pr list --state all --limit 100 \
        --json number,title,state,author,createdAt,mergedAt,additions,deletions,labels \
        > /tmp/gh-aw/pr-data/recent-prs.json
      jq '{
        total: length,
        merged: [.[] | select(.state == "MERGED")] | length,
        top_authors: ([.[].author.login] | group_by(.) | map({author: .[0], count: length}) | sort_by(-.count) | .[0:5])
      }' /tmp/gh-aw/pr-data/recent-prs.json > /tmp/gh-aw/pr-data/stats.json
---

# Weekly Pull Request Summary

Read `/tmp/gh-aw/pr-data/recent-prs.json` and `/tmp/gh-aw/pr-data/stats.json`. Create a discussion summarizing total PRs, merge rate, code changes, top contributors, and any notable trends.
````

### Data Caching

For workflows that run frequently or process large datasets, cache the data directory to avoid redundant API calls:

```aw wrap
---
cache:
  - key: pr-data-${{ github.run_id }}
    path: /tmp/gh-aw/pr-data
    restore-keys: |
      pr-data-
---
```

### Multi-Source Data

Combine multiple sources into a single dataset before analysis by adding a final `jq` step that merges the per-source JSON files (`gh pr list`, `gh issue list`, `gh run list`, etc.) into `/tmp/gh-aw/combined.json`, then point the agent at the combined file.

## Subagents with Smaller Models

Delegate narrow, repetitive reasoning — categorization, per-item summarization, sentiment scoring — to **inline sub-agents** on a smaller, cheaper model, leaving the main agent to synthesize the final output:

```
steps:          → deterministic shell commands (zero AI cost)
sub-agents:     → small-model agents for per-item analysis (cheap, parallelizable)
main agent:     → orchestrates sub-agents, writes final report (high-reasoning)
```

Enable inline sub-agents by adding `cli-proxy` so they can make authenticated GitHub API calls:

```yaml
tools:
  cli-proxy: true
```

### Example: Issue Triage with Categorization

```aw wrap
# Weekly Issue Triage

The raw issue data is in `/tmp/gh-aw/triage/` — one file per issue (`issue-<number>.json`).

For each `issue-*.json`, call `issue-categorizer` to classify it and `issue-summarizer` to produce a one-sentence summary; write results to `category-<number>.json` and `summary-<number>.json`. Then create a discussion that groups issues by category with their summaries and highlights the top 3 most urgent.

## agent: `issue-categorizer`
---
description: Classifies a GitHub issue into exactly one category
model: claude-haiku-4.5
---
Classify the issue into exactly one of: bug, feature-request, question, documentation, performance, security, or other.
Return `{"number": <issue number>, "category": "<category>"}`.

## agent: `issue-summarizer`
---
description: Produces a one-sentence summary of a GitHub issue
model: claude-haiku-4.5
---
Write a single sentence (≤ 20 words) describing the issue.
Return `{"number": <issue number>, "summary": "<sentence>"}`.
```

| Layer | Model | Work done | Cost driver |
|---|---|---|---|
| `steps:` | — | Fetch + prepare data | GitHub API only |
| `issue-categorizer` | Haiku / small | Classify one issue | ~200 tokens per issue |
| `issue-summarizer` | Haiku / small | Summarize one issue | ~150 tokens per issue |
| Main agent | Full model | Read all results, write report | One high-quality pass |

## Related Documentation

- [Pre-Activation Steps](/gh-aw/reference/triggers/#pre-activation-steps-onsteps) - Inline step injection into the pre-activation job
- [Pre-Activation Permissions](/gh-aw/reference/triggers/#pre-activation-permissions-onpermissions) - Grant additional scopes for `on.steps:` API calls
- [Custom Safe Outputs](/gh-aw/reference/custom-safe-outputs/) - Custom post-processing jobs
- [Frontmatter Reference](/gh-aw/reference/frontmatter/) - Configuration options
- [Compilation Process](/gh-aw/reference/compilation-process/) - How jobs are orchestrated
- [Imports](/gh-aw/reference/imports/) - Sharing configurations across workflows
- [Templating](/gh-aw/reference/templating/) - Using GitHub Actions expressions
