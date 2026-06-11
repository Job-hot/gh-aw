---
description: Guidance for creating agentic workflows that analyze test coverage — prefer reading pre-computed CI artifacts over re-running tests.
---

# Test Coverage Workflow Guidance

Consult this file when creating or updating a workflow that analyzes test coverage.

## Core Principle: Read Artifacts First

**Always prefer fetching pre-computed coverage artifacts from CI over re-running the test suite.** Re-running tests is slow and duplicates work CI has already done.

## Coverage Data Strategy

Include this decision block in every coverage workflow prompt:

```
1. Find the latest successful CI run for this commit:
   `gh run list --commit "$HEAD_SHA" --status success --limit 5 --json databaseId,workflowName`
2. Download the coverage artifact (try names: coverage-report, coverage, test-results):
   `gh run download <run-id> --name coverage-report --dir /tmp/coverage`
3. If found, parse and analyze it — do NOT re-run tests.
4. If not found, run tests with coverage and note in the report that data was computed fresh.
```

## Frontmatter Template

```yaml
engine: copilot
on:
  pull_request:
    types: [opened, synchronize]
permissions:
  actions: read         # download artifacts
network:
  allowed:
    - defaults
    # Artifact-download path only needs defaults.
    # If tests must run as fallback, add the project's package ecosystem.
    # Detect with: gh aw domains <workflow-file>
    # Common additions: node | python | go | java | dotnet | ruby | rust
tools:
  github:
    toolsets: [default, actions]  # actions toolset enables artifact download
safe-outputs:
  add-comment:
    hide-older-comments: true
```

## Fallback: Run Tests

Use **only when** no prior CI artifact exists or CI doesn't upload coverage. Supported commands:

| Language | Command |
|---|---|
| Node.js | `npx jest --coverage --coverageReporters=json-summary` |
| Python | `python -m pytest --cov=src --cov-report=json` |
| Go | `go test ./... -coverprofile=/tmp/coverage.out` |

### Network for Fallback Execution

`network: defaults` alone does not cover package registries. When tests may need to install or resolve dependencies at runtime, extend the `network.allowed` list with the matching ecosystem identifier. Infer the ecosystem from repository files:

| Repository file | Add to `network.allowed` |
|---|---|
| `package.json`, `yarn.lock`, `pnpm-lock.yaml` | `node` |
| `requirements.txt`, `pyproject.toml`, `uv.lock` | `python` |
| `go.mod` | `go` |
| `pom.xml`, `build.gradle` | `java` |
| `*.csproj`, `*.sln` | `dotnet` |
| `Gemfile` | `ruby` |
| `Cargo.toml` | `rust` |

See [network.md](network.md) for the full list of ecosystem identifiers.
