---
name: visual-regression
description: Reference prompt for visual regression testing using playwright + cache-memory for baseline screenshot storage across pull requests
---

# Visual Regression Testing

Use `playwright` for screenshots and `cache-memory` to persist baselines between PR runs.

## Example Workflow

```markdown
---
description: Capture screenshots on every PR and compare against cached baselines to detect visual regressions
on:
  pull_request:
    types: [opened, synchronize, reopened]
permissions:
  contents: read
  pull-requests: read
engine: copilot
tools:
  playwright:
    allowed_domains:
      - localhost
      - 127.0.0.1
  cache-memory:
    key: visual-regression-baselines-${{ github.event.pull_request.base.ref }}
    retention-days: 30
    allowed-extensions: [".png", ".json"]
  bash:
    - "mkdir *"
    - "cp *"
    - "diff *"
    - "date *"
safe-outputs:
  add-comment:
    max: 1
  upload-artifact:
    max-uploads: 5
    retention-days: 14
    allowed-paths:
      - "visual-regression-diffs/**"
timeout-minutes: 30
---

Build and serve the app locally, then use Playwright to capture full-page screenshots of key pages into `/tmp/visual-regression/current/`.

Use filesystem-safe timestamps (no colons — colons break artifact uploads):
`date -u "+%Y-%m-%d-%H-%M-%S"`

If `/tmp/gh-aw/cache-memory/baselines/manifest.json` does not exist, copy screenshots there as new baselines and post: "Baselines initialized — N pages captured."

Otherwise compare each screenshot to its baseline. Copy any diff images to `visual-regression-diffs/` (workspace-relative). Upload them using `upload_artifact` (name: `visual-regression-diffs-${{ github.run_id }}`). Post a comment summarizing: pages unchanged / pages with diffs, and link to the uploaded artifact when diffs exist. If nothing changed, use the `noop` safe-output.
```

## Baseline Storage

Baselines are stored in `cache-memory` under `/tmp/gh-aw/cache-memory/baselines/`. The key is scoped per base branch so `main` and `develop` maintain independent baseline sets.

### Updating Baselines

Add a companion `workflow_dispatch` workflow that captures fresh screenshots and writes them to `cache-memory` — using the same `key` pattern but with the target branch as an input. Run it on `main` after merging any PR that intentionally changes the UI.

## Artifact Retention

- Set `retention-days` to match your PR review cadence (14 days covers most review cycles).
- Pair the `upload_artifact` call with an `add-comment` that links the artifact URL: `${{ steps.process_safe_outputs.outputs.upload_artifact_url }}`.
- Use `upload-asset` instead of `upload-artifact` when diff images need stable embeddable URLs in the comment body.

## Key Design Decisions

- **`cache-memory` key per base branch** — scopes baselines to `main`, `develop`, etc.
- **`allowed_domains: [localhost, 127.0.0.1]`** — prevents SSRF; serve app locally
- **`retention-days: 30`** — beyond the default 7-day cache expiry
- **Filesystem-safe timestamps** — `YYYY-MM-DD-HH-MM-SS`; colons break artifact filenames
- **Minimal permissions** — all PR writes go through `safe-outputs`
- **`upload-artifact` for diffs** — run-scoped; expires after `retention-days`; use for downloadable diff archives
- **`upload-asset` for embedded images** — persists to orphaned branch; use when the comment body must inline diff images

## Network-Minimization Reminders

- Prefer local preview (`localhost`/`127.0.0.1`) over external preview environments.
- If external previews are required, allowlist exact domains (no broad wildcards).
- Enable `network.node` only when installing/building Node deps; scope to registries and preview hosts.
- Keep Playwright navigation limited to app-under-test URLs.
