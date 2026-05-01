---
title: DispatchOps
description: Manually trigger and test agentic workflows with custom inputs using workflow_dispatch
sidebar:
  badge: { text: 'Manual', variant: 'tip' }
---

DispatchOps enables manual workflow execution via the GitHub Actions UI or CLI, perfect for on-demand tasks, testing, and workflows that need human judgment about timing. The `workflow_dispatch` trigger lets you run workflows with custom inputs whenever needed.

Use DispatchOps for research tasks, operational commands, testing workflows during development, debugging production issues, or any task that doesn't fit a schedule or event trigger.

## How Workflow Dispatch Works

### With Input Parameters

Define inputs to customize workflow behavior at runtime:

```yaml
on:
  workflow_dispatch:
    inputs:
      topic:
        description: 'Research topic'
        required: true
        type: string
      priority:
        description: 'Task priority'
        required: false
        type: choice
        options:
          - low
          - medium
          - high
        default: medium
      deploy_target:
        description: 'Deployment environment'
        required: false
        type: environment
        default: staging
```

Supported input types: `string` (text), `boolean` (checkbox), `choice` (dropdown), `environment` (GitHub environments dropdown).

### Environment Input Type

The `environment` type auto-populates from repository Settings → Environments, returning the selected name as a string. No `options` list is needed; specify a `default` matching an existing environment name. The type does not enforce protection rules — use `manual-approval:` for approval gates (see [Environment Approval Gates](#environment-approval-gates)).

## Security Model

### Permission Requirements

Manual workflow execution respects the same security model as other triggers:

- **Repository permissions** - User must have write access or higher to trigger workflows
- **Role-based access** - Use the `roles:` field to restrict who can run workflows:

```yaml
on:
  workflow_dispatch:
roles: [admin, maintainer]
```

- **Bot authorization** - Use the `bots:` field to allow specific bot accounts:

```yaml
on:
  workflow_dispatch:
bots: ["dependabot[bot]", "github-actions[bot]"]
```

Unlike issue/PR triggers, `workflow_dispatch` only executes in the repository where it's defined — forks cannot trigger workflows in the parent repository.

### Environment Approval Gates

Require manual approval before execution using GitHub environment protection rules:

```yaml
on:
  workflow_dispatch:
manual-approval: production
```

Configure approval rules, required reviewers, and wait timers in repository Settings → Environments. See [GitHub's environment documentation](https://docs.github.com/en/actions/deployment/targeting-different-environments/using-environments-for-deployment) for setup details.

## Running Workflows from GitHub.com

Go to the **Actions** tab, select the workflow from the sidebar, click **Run workflow**, fill in any inputs, and confirm. Only workflows with `workflow_dispatch:` in their `on:` section appear in the dropdown — if yours is missing, verify it has been compiled and the `.lock.yml` pushed to the repository.

## Running Workflows with CLI

The `gh aw run` command provides a faster way to trigger workflows from the command line.

### Basic Usage

```bash
gh aw run workflow
```

This matches workflows by filename prefix, validates `workflow_dispatch:`, and returns the run URL immediately.

### With Input Parameters

Pass inputs using the `--raw-field` or `-f` flag in `key=value` format:

```bash
gh aw run research --raw-field topic="quantum computing"
```

```bash
gh aw run scout \
  --raw-field topic="AI safety research" \
  --raw-field priority=high
```

### Wait for Completion

Monitor workflow execution and wait for results:

```bash
gh aw run research --raw-field topic="AI agents" --wait
```

`--wait` monitors progress in real-time and exits with a success/failure code on completion. Additional flags:

```bash
gh aw run research --ref feature-branch              # Run from specific branch
gh aw run workflow --repo owner/repository           # Run in another repository
gh aw run research --raw-field topic="AI" --verbose  # Verbose output
```

## Declaring and Referencing Inputs

### Referencing Inputs in Markdown

Access input values using GitHub Actions expression syntax:

```aw wrap
---
on:
  workflow_dispatch:
    inputs:
      topic:
        description: 'Research topic'
        required: true
        type: string
      depth:
        description: 'Analysis depth'
        type: choice
        options:
          - brief
          - detailed
        default: brief
permissions:
  contents: read
safe-outputs:
  create-discussion:
---

# Research Assistant

Research the following topic: "${{ github.event.inputs.topic }}"

Analysis depth requested: ${{ github.event.inputs.depth }}

Provide a ${{ github.event.inputs.depth }} analysis with key findings and recommendations.
```

Reference inputs with `${{ github.event.inputs.INPUT_NAME }}`—values are interpolated at compile time throughout the workflow.

### Conditional Logic Based on Inputs

Use Handlebars conditionals to change behavior based on input values:

```markdown
{{#if (eq github.event.inputs.include_code "true")}}
Include actual code snippets in your analysis.
{{else}}
Describe code patterns without including actual code.
{{/if}}

{{#if (eq github.event.inputs.priority "high")}}
URGENT: Prioritize speed over completeness.
{{/if}}
```

## Development Pattern: Branch Testing

Add `workflow_dispatch:` to feature branches for testing before merging. Use [trial mode](/gh-aw/patterns/trial-ops/) for isolated testing without affecting the production repository, or run from a branch directly:

```bash
gh aw trial ./research.md --raw-field topic="test query"  # isolated, no side effects
gh aw run research --ref feature/improve-workflow          # runs against live repo
```

## Common Use Cases

- **On-demand research**: trigger with a `topic` input — `gh aw run research --raw-field topic="AI safety"`
- **Manual operations**: `choice` input for predefined tasks (cleanup, sync, audit)
- **Testing**: add `workflow_dispatch` to event-triggered workflows to run without creating real events
- **Scheduled testing**: combine `schedule` with `workflow_dispatch` for immediate on-demand runs

## Troubleshooting

| Problem | Solution |
|---------|----------|
| Workflow not in GitHub UI | Verify `workflow_dispatch:` exists in `on:`, recompile, push both `.md` + `.lock.yml`; refresh the page |
| "Workflow not found" | Use filename without `.md` extension (`research` not `research.md`); verify it's compiled |
| "Workflow cannot be run" | Add `workflow_dispatch:` to `on:`, recompile, push `.lock.yml` |
| Permission denied | Verify write access; check `roles:` in frontmatter; confirm org role |
| Inputs not appearing | Check YAML indentation (2 spaces); verify valid types; recompile and push |
| Wrong branch context | Use `--ref branch-name` in CLI or select the correct branch in the UI dropdown |

## Related Documentation

- [Manual Workflows Example](/gh-aw/examples/manual/) - Example manual workflows
- [Triggers Reference](/gh-aw/reference/triggers/) - Complete trigger syntax including workflow_dispatch
- [TrialOps](/gh-aw/patterns/trial-ops/) - Testing workflows in isolation
- [CLI Commands](/gh-aw/setup/cli/) - Complete gh aw run command reference
- [Templating](/gh-aw/reference/templating/) - Using expressions and conditionals
- [Security Best Practices](/gh-aw/introduction/architecture/) - Securing workflow execution
- [Quick Start](/gh-aw/setup/quick-start/) - Getting started with agentic workflows
