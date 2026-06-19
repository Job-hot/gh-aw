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

**Manifest structure** (`manifest.json`):
```json
{
  "created": "2024-01-15-10-30-00",
  "pages": ["home", "dashboard", "settings"],
  "screenshots": {
    "home": "home-2024-01-15-10-30-00.png",
    "dashboard": "dashboard-2024-01-15-10-30-00.png"
  }
}
```

### Updating Baselines

Baselines must be updated whenever the UI intentionally changes (new feature, redesign). Add a companion `workflow_dispatch` workflow for manual baseline updates:

```markdown
---
description: Manually update visual regression baselines for a given base branch
on:
  workflow_dispatch:
    inputs:
      base-branch:
        description: "Base branch to update baselines for (e.g. main)"
        required: true
        default: "main"
permissions:
  contents: read
tools:
  playwright:
    allowed_domains:
      - localhost
      - 127.0.0.1
  cache-memory:
    key: visual-regression-baselines-${{ github.event.inputs.base-branch }}
    retention-days: 30
    allowed-extensions: [".png", ".json"]
  bash:
    - "mkdir *"
    - "cp *"
    - "date *"
safe-outputs:
  add-comment:
    max: 1
---

Build and serve the app locally. Capture full-page screenshots of all key pages into `/tmp/visual-regression/current/`.

Copy the screenshots to `/tmp/gh-aw/cache-memory/baselines/` and write a new `manifest.json` with the current timestamp. Post a comment: "Baselines updated for `${{ github.event.inputs.base-branch }}` — N pages captured."
```

> **Tip**: Run the baseline-update workflow on `main` after merging a PR that intentionally changes the UI. The PR regression workflow will then compare against the updated baselines on subsequent PRs.

## Artifact Retention

Upload diff screenshots as workflow artifacts so reviewers can inspect them without re-running:

```yaml
safe-outputs:
  upload-artifact:
    max-uploads: 5
    retention-days: 14          # keep diffs for 2 weeks (range: 1-90)
    allowed-paths:
      - "visual-regression-diffs/**"  # workspace-relative; agent copies diffs here before uploading
```

- Set `retention-days` to match your PR review cadence (14 days covers most review cycles).
- Pair the `upload_artifact` call with an `add-comment` that links the artifact URL: `${{ steps.process_safe_outputs.outputs.upload_artifact_url }}`.
- Use `upload-asset` instead of `upload-artifact` when diff images need stable embeddable URLs in the comment body (images embedded in GitHub markdown require a raw URL, not a zip archive link).

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
