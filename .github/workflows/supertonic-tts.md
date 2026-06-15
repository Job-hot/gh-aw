---
emoji: "🔊"
name: Supertonic TTS
description: Converts text to speech using the Supertonic on-device TTS SDK, uploads the generated audio as a run artifact, and posts a comment with the download link on the target issue or pull request
on:
  workflow_dispatch:
    inputs:
      text:
        description: "Text to convert to speech"
        required: true
        type: string
      issue-number:
        description: "Issue number to post the artifact link to (mutually exclusive with pr-number)"
        required: false
        type: string
      pr-number:
        description: "Pull request number to post the artifact link to (mutually exclusive with issue-number)"
        required: false
        type: string

timeout-minutes: 15

permissions:
  contents: read
  issues: read
  pull-requests: read
  copilot-requests: write

engine:
  id: copilot
  bare: true

safe-outputs:
  noop:
  jobs:
    supertonic-tts:
      description: "Install Supertonic TTS, synthesize speech from text, upload as artifact, and post comment"
      runs-on: ubuntu-latest
      permissions:
        contents: read
        issues: write
        pull-requests: write
      inputs:
        text:
          type: string
          required: true
          description: "Text to synthesize"
        issue_number:
          type: string
          required: false
          description: "Issue number to post the artifact link to"
        pr_number:
          type: string
          required: false
          description: "Pull request number to post the artifact link to"
      steps:
        - name: Install Supertonic with server support
          id: install-supertonic
          run: |
            set -e
            mkdir -p /tmp/gh-aw/tts
            python3 -m venv /tmp/gh-aw/tts/venv
            /tmp/gh-aw/tts/venv/bin/pip install --quiet 'supertonic[serve]'
            SUPERTONIC_VERSION=$(/tmp/gh-aw/tts/venv/bin/python3 -c "import supertonic; print(supertonic.__version__)")
            echo "version=$SUPERTONIC_VERSION" >> "$GITHUB_OUTPUT"
            echo "Supertonic $SUPERTONIC_VERSION installed"
        - name: Restore Supertonic model cache
          uses: actions/cache/restore@v5.0.5
          with:
            key: supertonic3-model-${{ runner.os }}-${{ steps.install-supertonic.outputs.version }}
            restore-keys: supertonic3-model-${{ runner.os }}-
            path: ~/.cache/supertonic3
        - name: Start Supertonic HTTP server
          run: |
            set -e
            nohup /tmp/gh-aw/tts/venv/bin/supertonic serve \
              --host 127.0.0.1 --port 7788 \
              > /tmp/gh-aw/tts/supertonic.log 2>&1 &
            echo $! > /tmp/gh-aw/tts/supertonic.pid
            echo "Server PID: $(cat /tmp/gh-aw/tts/supertonic.pid)"
        - name: Wait for Supertonic server readiness
          run: |
            URL="http://127.0.0.1:7788/docs"
            STATUS=""
            for i in $(seq 1 60); do
              STATUS=$(curl -sS -o /dev/null -w "%{http_code}" --connect-timeout 5 --max-time 5 "$URL" || true)
              [ "$STATUS" = "200" ] && echo "Server ready at http://127.0.0.1:7788" && break
              [ -z "$STATUS" ] && STATUS="curl_error"
              echo "Waiting for server... ($i/60, status: $STATUS)" && sleep 3
            done
            if [ "$STATUS" != "200" ]; then
              cat /tmp/gh-aw/tts/supertonic.log || true
              exit 1
            fi
        - name: Synthesize speech
          id: synthesize
          run: |
            set -e
            TEXT=$(jq -r '.items[] | select(.type == "supertonic_tts") | .text' "$GH_AW_AGENT_OUTPUT" | head -1)
            ISSUE_NUMBER=$(jq -r '.items[] | select(.type == "supertonic_tts") | .issue_number // empty' "$GH_AW_AGENT_OUTPUT" | head -1)
            PR_NUMBER=$(jq -r '.items[] | select(.type == "supertonic_tts") | .pr_number // empty' "$GH_AW_AGENT_OUTPUT" | head -1)
            echo "issue_number=$ISSUE_NUMBER" >> "$GITHUB_OUTPUT"
            echo "pr_number=$PR_NUMBER" >> "$GITHUB_OUTPUT"
            curl -sS -X POST http://127.0.0.1:7788/v1/tts \
              -H 'content-type: application/json' \
              --data-binary "$(jq -n --arg t "$TEXT" '{"text": $t, "voice": "F1", "lang": "en", "steps": 8}')" \
              -o /tmp/gh-aw/tts/output.wav
            echo "Audio generated: $(ls -lh /tmp/gh-aw/tts/output.wav)"
        - name: Save Supertonic model cache
          uses: actions/cache/save@v5.0.5
          with:
            key: supertonic3-model-${{ runner.os }}-${{ steps.install-supertonic.outputs.version }}
            path: ~/.cache/supertonic3
        - name: Upload audio artifact
          id: upload
          uses: actions/upload-artifact@v7.0.1
          with:
            name: tts-output
            path: /tmp/gh-aw/tts/output.wav
            retention-days: 30
        - name: Post comment with artifact link
          run: |
            ITEM_NUMBER="${ISSUE_NUMBER:-$PR_NUMBER}"
            if [ -z "$ITEM_NUMBER" ]; then
              echo "No target item number found in agent output, skipping comment"
              exit 0
            fi
            mkdir -p /tmp/gh-aw/tts
            jq -n --arg url "$ARTIFACT_URL" \
              '{"body": ("🔊 **Voice summary ready**\n\n[🎧 Download tts-output.wav](" + $url + ")")}' \
              > /tmp/gh-aw/tts/comment.json
            gh api "repos/$REPO/issues/$ITEM_NUMBER/comments" --input /tmp/gh-aw/tts/comment.json
          env:
            GH_TOKEN: ${{ github.token }}
            ARTIFACT_URL: ${{ steps.upload.outputs.artifact-url }}
            ISSUE_NUMBER: ${{ steps.synthesize.outputs.issue_number }}
            PR_NUMBER: ${{ steps.synthesize.outputs.pr_number }}
            REPO: ${{ github.repository }}
---

# Supertonic TTS

Convert the provided text to speech, upload the WAV as a run artifact, and post a comment with the download link on the target issue or pull request.

## Inputs

- **Text**: ${{ github.event.inputs.text }}
- **Issue number**: ${{ github.event.inputs.issue-number }}
- **PR number**: ${{ github.event.inputs.pr-number }}

## Task

### Step 1 — Validate inputs

Exactly one of `issue-number` or `pr-number` must be provided. Check the inputs:

- If **both** are set: call `noop` with message `"Both issue-number and pr-number were provided. Provide exactly one."` and stop.
- If **neither** is set: call `noop` with message `"Neither issue-number nor pr-number was provided. Provide exactly one."` and stop.
- Otherwise: proceed.

### Step 2 — Submit TTS request

Call the `supertonic_tts` safe-output tool with the text and the target issue or PR number.

If `issue-number` is set (e.g., `42`):
```json
{
  "type": "supertonic_tts",
  "text": "${{ github.event.inputs.text }}",
  "issue_number": "42"
}
```

If `pr-number` is set (e.g., `99`):
```json
{
  "type": "supertonic_tts",
  "text": "${{ github.event.inputs.text }}",
  "pr_number": "99"
}
```
