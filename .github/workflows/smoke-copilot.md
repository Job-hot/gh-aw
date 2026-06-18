---
private: true
emoji: "🧪"
description: Smoke Copilot
on: 
  slash_command:
    name: smoke-copilot
    strategy: centralized
    events: [issues, issue_comment, pull_request, pull_request_comment]
  workflow_dispatch:
  label_command:
    name: smoke
    events: [pull_request]
  reaction: "eyes"
  status-comment: true
  github-token: ${{ secrets.GH_AW_GITHUB_TOKEN || secrets.GITHUB_TOKEN }}
permissions:
  contents: read
  pull-requests: read
  issues: read
  discussions: read
  actions: read
name: Smoke Copilot
engine:
  id: copilot
  model: gpt-5.4
  max-continuations: 2
  bare: true
imports:
  - shared/github-guard-policy.md
  - shared/gh.md
  - shared/reporting.md
  - shared/github-queries-mcp-script.md
  - shared/mcp/serena-go.md
  - shared/otlp.md
network:
  allowed:
    - defaults
    - node
    - github
    - playwright
tools:
  agentic-workflows:
  cache-memory: true
  comment-memory: true
  edit:
  bash:
    - "*"
  github:
    mode: gh-proxy
    min-integrity: approved
    trusted-users:
      - pelikhan
  playwright:
    mode: cli
  web-fetch:
  cli-proxy: true
runtimes:
  go:
    version: "1.26"
models:
  providers:
    anthropic:
      models:
        my-custom-claude:
          cost:
            input: "3e-06"
            output: "1.5e-05"
            cache_read: "3e-07"
            cache_write: "3.75e-06"
safe-outputs:
    allowed-domains: [default-safe-outputs]
    upload-artifact:
      max-uploads: 1
      retention-days: 1
      skip-archive: true
    add-comment:
      allowed-repos: ["github/gh-aw"]
      hide-older-comments: true
      max: 2
    create-issue:
      expires: 2h
      group: true
      close-older-issues: true
      close-older-key: "smoke-copilot"
      labels: [automation, testing]
    create-discussion:
      category: announcements
      labels: [ai-generated]
      expires: 2h
      close-older-discussions: true
      close-older-key: "smoke-copilot"
      max: 1
    create-pull-request-review-comment:
      max: 5
    submit-pull-request-review:
    reply-to-pull-request-review-comment:
      max: 5
    add-labels:
      allowed: [smoke-copilot]
      allowed-repos: ["github/gh-aw"]
    remove-labels:
      allowed: [smoke]
    set-issue-type:
    dispatch-workflow:
      workflows:
        - haiku-printer
      max: 1
    create-check-run:
      name: "Smoke Copilot"
      max: 1
    jobs:
      send-slack-message:
        description: "Send a message to Slack (stub for testing)"
        runs-on: ubuntu-latest
        output: "Slack message stub executed!"
        inputs:
          message:
            description: "The message to send"
            required: false
            default: ""
            type: string
        permissions:
          contents: read
        steps:
          - name: Stub Slack message
            run: |
              echo "🎭 This is a stub - not sending to Slack"
              if [ -f "$GH_AW_AGENT_OUTPUT" ]; then
                MESSAGE=$(cat "$GH_AW_AGENT_OUTPUT" | jq -r '.items[] | select(.type == "send_slack_message") | .message')
                echo "Would send to Slack: $MESSAGE"
                {
                  echo "### 📨 Slack Message Stub"
                  echo "**Message:** $MESSAGE"
                  echo ""
                  echo "> ℹ️ This is a stub for testing purposes. No actual Slack message is sent."
                } >> "$GITHUB_STEP_SUMMARY"
              else
                echo "No agent output found"
              fi
    messages:
      append-only-comments: true
      footer: "> 📰 *BREAKING: Report filed by [{workflow_name}]({run_url})*{ai_credits_suffix}{history_link}"
      run-started: "📰 BREAKING: [{workflow_name}]({run_url}) is now investigating this {event_type}. Sources say the story is developing..."
      run-success: "📰 VERDICT: [{workflow_name}]({run_url}) has concluded. All systems operational. This is a developing story. 🎤"
      run-failure: "📰 DEVELOPING STORY: [{workflow_name}]({run_url}) reports {status}. Our correspondents are investigating the incident..."
timeout-minutes: 15
strict: false
experiments:
  caveman: [yes, no]
  subagent_model: [small, large]

---

# Smoke Test: Copilot Engine Validation

{{#if experiments.caveman }}
Talk like a caveman in all your responses and outputs. Use short, broken sentences. Me test. You run.
{{/if}}

**IMPORTANT: Keep all outputs extremely short and concise. Use single-line responses where possible. No verbose explanations.**

## Hard Limit: `add_comment` Budget

`safe-outputs.add-comment.max` is `2`. Never exceed 2 total `add_comment` calls in this run.

- Call #1 is required for the discussion interaction test (comment on latest discussion).
- Call #2 depends on trigger:
  - `pull_request` event: post the brief PR summary comment, and **skip** the fun discussion follow-up comment.
  - non-`pull_request` event: **skip** the PR summary comment and post the fun discussion follow-up comment.

## Test Requirements

Run these checks and mark each as ✅/❌:

1. GitHub MCP: review 2 merged PRs in `${{ github.repository }}`.
2. `mcpscripts-gh`: `pr list --repo ${{ github.repository }} --limit 2 --json number,title,author`.
3. Serena (bash): `serena activate_project --path ${{ github.workspace }}`; `serena find_symbol --name_path <symbol>`; confirm ≥3 symbols.
4. Playwright (bash): `playwright-cli open https://github.com`; `playwright-cli screenshot`; confirm navigation succeeded.
5. Web fetch: fetch `https://github.com`; confirm response contains `GitHub`.
6. File+bash: create `/tmp/gh-aw/agent/smoke-test-copilot-${{ github.run_id }}.txt` with timestamped success text; `cat` it.
7. Discussion comment: get latest discussion (`limit=1`, `jq=".[0]"`); `add_comment` to it.
8. Build: `GOCACHE=/tmp/gh-aw/agent/go-cache GOMODCACHE=/tmp/gh-aw/agent/go-mod make build`.
9. Artifact (build-pass only): stage `./gh-aw` at `$RUNNER_TEMP/gh-aw/safeoutputs/upload-artifacts/gh-aw`; call `upload_artifact` with `path: "gh-aw"`.
10. Discussion create: `create_discussion` in `announcements`, label `ai-generated`, title `copilot was here`, temp ID `aw_smoke_discussion`.
11. Workflow dispatch: `dispatch_workflow` for `haiku-printer` with an original testing/automation haiku.
12. PR review: 1–2 `create_pull_request_review_comment` inline comments; submit review with `COMMENT`; reply to most recent existing review comment when available.
13. Comment memory: append an original 3-line haiku to `/tmp/gh-aw/comment-memory/*.md`.
14. Sub-agent: use `file-summarizer` on `README.md`.
15. Check run: `create_check_run` with `conclusion=success`, title `Smoke Copilot - Run ${{ github.run_id }}`, summary `All smoke tests completed.`, text `Detailed results attached.`

## Output

- **Create issue** (`aw_smoke1`): title `Smoke Test: Copilot - ${{ github.run_id }}`; body includes ✅/❌ per test, overall PASS/FAIL, run URL, timestamp, PR author and assignees. Then call `set_issue_type` with `issue_number: "aw_smoke1"` and `issue_type: "Bug"`.
- **Comment** (max 2 total): if `pull_request` event → brief PR comment (≤10 lines, ✅/❌ per test, PASS/FAIL, mention author/assignees); else → fun creative comment on `aw_smoke_discussion`. Call `send_slack_message` with a one-line summary.
- **Labels** (PR event, all-pass only): `add_labels` → `smoke-copilot`; `remove_labels` → `smoke`.

{{#runtime-import shared/noop-reminder.md}}

## agent: `file-summarizer`
---
model: ${{ experiments.subagent_model }}
description: Summarizes the content of a file in a few concise sentences
---
You are a file summarization assistant. When given a file path, read the file and return a brief summary (2–4 sentences) describing its purpose and key contents. Be concise and factual.