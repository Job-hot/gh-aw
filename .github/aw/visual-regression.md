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
timeout-minutes: 30
---

Build and serve the app locally, then use Playwright to capture full-page screenshots of key pages into `/tmp/visual-regression/current/`.

Use filesystem-safe timestamps (no colons — colons break artifact uploads):
`date -u "+%Y-%m-%d-%H-%M-%S"`

If `/tmp/gh-aw/cache-memory/baselines/manifest.json` does not exist, copy screenshots there as new baselines and post: "Baselines initialized — N pages captured."

Otherwise compare each screenshot to its baseline. Post a comment summarizing: pages unchanged / pages with diffs. If nothing changed, use the `noop` safe-output.
```

## Key Design Decisions

- **`cache-memory` key per base branch** — scopes baselines to `main`, `develop`, etc. independently
- **`allowed_domains: [localhost, 127.0.0.1]`** — prevents SSRF; app must be served locally
- **`retention-days: 30`** — keeps baselines beyond the default 7-day cache expiry
- **Filesystem-safe timestamps** — `YYYY-MM-DD-HH-MM-SS` format; colons are invalid in artifact filenames
- **Minimal permissions** — all PR writes go through `safe-outputs`, not GitHub tools

## Optional: Screenshot Diff Artifacts

For detailed diff images that are too large for a PR comment, upload them as workflow artifacts and link to them from the comment. Add an `actions: write` permission and a `bash` step after the agent job:

```yaml
permissions:
  contents: read
  pull-requests: read
  actions: write        # required to upload artifacts
```

Add this `steps:` block after the agent step to collect and upload any diffs the agent wrote to `/tmp/visual-regression/diffs/`:

```yaml
steps:
  - name: Upload screenshot diffs
    if: always()
    uses: actions/upload-artifact@v4
    with:
      name: visual-regression-diffs-${{ github.event.pull_request.number }}
      path: /tmp/visual-regression/diffs/
      retention-days: 14
      if-no-files-found: ignore
```

Keep PR comments concise: include only the pass/fail counts and a link to the artifact. Example comment format:

```
Visual regression: 2 pages changed, 5 unchanged.
Screenshot diffs: [view artifacts](<artifact-url>)
```

> ⚠️ The agent cannot directly generate the artifact URL at comment-post time. Either omit the link or use a pre-step to compute the run URL (`${{ github.server_url }}/${{ github.repository }}/actions/runs/${{ github.run_id }}`).
