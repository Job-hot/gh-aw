---
emoji: "🔊"
name: Supertonic TTS
description: Converts text to speech using the Supertonic on-device TTS server, uploads the generated audio as a run artifact, and posts a comment with the download link on the target issue or pull request
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

imports:
  - shared/supertonic.md

network:
  allowed:
    - defaults
    - python

tools:
  bash:
    - "curl *"
    - "jq *"
    - "cp *"
    - "mkdir *"
    - "ls *"
    - "echo *"
    - "cat *"

safe-outputs:
  upload-artifact:
    max-uploads: 1
    retention-days: 30
    skip-archive: true
  add-comment:
    max: 1
  noop:
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

### Step 2 — Synthesize speech

The Supertonic TTS server is running at `http://127.0.0.1:7788`. Use `jq` to safely encode the input text and avoid shell escaping issues:

```bash
mkdir -p /tmp/gh-aw/agent

# Write the text to a file to avoid here-doc quoting issues
cat > /tmp/gh-aw/agent/tts-input.txt << 'TTS_EOF'
${{ github.event.inputs.text }}
TTS_EOF

SUMMARY_TEXT=$(cat /tmp/gh-aw/agent/tts-input.txt)
curl -sS -X POST http://127.0.0.1:7788/v1/tts \
  -H 'content-type: application/json' \
  --data-binary "{\"text\": $(jq -n --arg t "$SUMMARY_TEXT" '$t'), \"voice\": \"F1\", \"lang\": \"en\", \"steps\": 8}" \
  -o /tmp/gh-aw/agent/tts-output.wav
echo "Audio generated: $(ls -lh /tmp/gh-aw/agent/tts-output.wav)"
```

### Step 3 — Upload as a run artifact

Stage the WAV file and call the `upload_artifact` safe-output tool:

```bash
mkdir -p "$RUNNER_TEMP/gh-aw/safeoutputs/upload-artifacts"
cp /tmp/gh-aw/agent/tts-output.wav \
  "$RUNNER_TEMP/gh-aw/safeoutputs/upload-artifacts/tts-output.wav"
```

Call the safe-output tool:
```json
{ "type": "upload_artifact", "path": "tts-output.wav" }
```

The tool returns `slot_N_artifact_url` — the direct download link to the uploaded file.

### Step 4 — Post comment with download link

Use `add_comment` to post the artifact link on the target item.

If `issue-number` is set (e.g., `42`):
```json
{
  "type": "add_comment",
  "issue_number": 42,
  "body": "🔊 **Voice summary ready**\n\n[🎧 Download tts-output.wav](<ARTIFACT_URL>)"
}
```

If `pr-number` is set (e.g., `99`):
```json
{
  "type": "add_comment",
  "pr_number": 99,
  "body": "🔊 **Voice summary ready**\n\n[🎧 Download tts-output.wav](<ARTIFACT_URL>)"
}
```

Replace `<ARTIFACT_URL>` with the `slot_N_artifact_url` returned by the `upload_artifact` tool.
