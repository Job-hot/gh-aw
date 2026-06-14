---
# Supertonic TTS HTTP Server
#
# Shared workflow that installs the Supertonic Python SDK, downloads the model,
# starts the HTTP server on port 7788, and provides curl instructions.
#
# Usage:
#   imports:
#     - shared/supertonic.md
#
# This import provides:
# - Supertonic TTS HTTP server on http://127.0.0.1:7788
# - /v1/tts endpoint (native Supertonic API)
# - /v1/audio/speech endpoint (OpenAI-compatible)
# - 10 built-in voices (M1–M5, F1–F5), 31 languages
# - WAV, FLAC, and OGG output formats
#
# Prerequisites:
# - Python 3.8+ available on runner
# - Network access to pypi.org (python ecosystem)
# - Sufficient disk space (~400MB for model download on first run)
# - Recommended timeout-minutes: ≥ 20 to allow model download and server startup
#
# Note: The Supertonic model (~400MB) is downloaded from Hugging Face on the first
# run and cached in ~/.cache/supertonic3/. The model cache is stored in
# actions/cache keyed by OS and package version, so subsequent runs are fast.

tools:
  bash:
    - "curl *"
    - "cat *"
    - "echo *"
    - "kill *"
    - "cp *"
    - "mkdir *"
    - "ls *"
    - "python3 *"
    - "python *"

network:
  allowed:
    - python

steps:
  - name: Install Supertonic with server support
    id: install-supertonic
    run: |
      set -e
      mkdir -p /tmp/gh-aw/agent
      if [ ! -d /tmp/gh-aw/agent/venv ]; then
        python3 -m venv /tmp/gh-aw/agent/venv
      fi
      echo "/tmp/gh-aw/agent/venv/bin" >> "$GITHUB_PATH"
      /tmp/gh-aw/agent/venv/bin/pip install --quiet 'supertonic[serve]'
      SUPERTONIC_VERSION=$(/tmp/gh-aw/agent/venv/bin/python3 -c "import supertonic; print(supertonic.__version__)")
      echo "version=$SUPERTONIC_VERSION" >> "$GITHUB_OUTPUT"
      echo "Supertonic $SUPERTONIC_VERSION installed with server support"

  - name: Restore Supertonic model cache
    uses: actions/cache/restore@v5.0.5
    with:
      key: supertonic3-model-${{ runner.os }}-${{ steps.install-supertonic.outputs.version }}
      restore-keys: supertonic3-model-${{ runner.os }}-
      path: ~/.cache/supertonic3

  - name: Start Supertonic HTTP server
    run: |
      set -e
      # Start server in background; first run downloads the model (~400MB) from HuggingFace
      nohup /tmp/gh-aw/agent/venv/bin/supertonic serve \
        --host 127.0.0.1 --port 7788 \
        > /tmp/gh-aw/agent/supertonic.log 2>&1 &
      STTC_PID=$!
      echo $STTC_PID > /tmp/gh-aw/agent/supertonic.pid
      echo "Supertonic server PID: $STTC_PID — logs at /tmp/gh-aw/agent/supertonic.log"

  - name: Wait for Supertonic server readiness
    run: |
      # Allow up to 3 minutes (60 iterations × 3s each) to accommodate model download on first run
      URL="http://127.0.0.1:7788/docs"
      STATUS=""
      for i in $(seq 1 60); do
        STATUS=$(curl -sS -o /dev/null -w "%{http_code}" \
          --connect-timeout 5 --max-time 5 "$URL" || true)
        [ "$STATUS" = "200" ] && echo "Supertonic server ready at http://127.0.0.1:7788" && break
        if [ -z "$STATUS" ]; then STATUS="curl_error"; fi
        echo "Waiting for Supertonic... ($i/60) (status: $STATUS)" && sleep 3
      done
      if [ "$STATUS" != "200" ]; then
        echo "Supertonic server failed to start after 3 minutes (final status: $STATUS)"
        cat /tmp/gh-aw/agent/supertonic.log || true
        exit 1
      fi

  - name: Save Supertonic model cache
    uses: actions/cache/save@v5.0.5
    with:
      key: supertonic3-model-${{ runner.os }}-${{ steps.install-supertonic.outputs.version }}
      path: ~/.cache/supertonic3
---

## Supertonic TTS HTTP Server Ready

The Supertonic text-to-speech HTTP server is running at `http://127.0.0.1:7788`.

### Generate Audio with curl

#### Native `/v1/tts` endpoint — full Supertonic parameter set

```bash
curl -X POST http://127.0.0.1:7788/v1/tts \
  -H 'content-type: application/json' \
  -d '{
    "text": "Hello, this is a test of the Supertonic TTS system.",
    "voice": "M1",
    "lang": "en",
    "steps": 8,
    "speed": 1.05,
    "response_format": "wav"
  }' \
  -o /tmp/gh-aw/agent/output.wav
echo "Audio saved to /tmp/gh-aw/agent/output.wav"
```

Response headers include:
- `X-Audio-Duration` — duration of the generated audio in seconds
- `X-Sample-Rate` — sample rate (44100 Hz)
- `X-Supertonic-Version` — model version

#### OpenAI-compatible `/v1/audio/speech` endpoint

```bash
curl -X POST http://127.0.0.1:7788/v1/audio/speech \
  -H 'content-type: application/json' \
  -d '{
    "model": "supertonic-3",
    "input": "Hello, this is a test of the Supertonic TTS system.",
    "voice": "M1",
    "response_format": "wav"
  }' \
  -o /tmp/gh-aw/agent/output.wav
```

### Available Endpoints

| Endpoint | Method | Description |
|---|---|---|
| `/v1/tts` | POST | Native synthesis — full Supertonic parameter set |
| `/v1/audio/speech` | POST | OpenAI-compatible speech endpoint |
| `/v1/styles` | GET | List all available voice styles |
| `/v1/styles/import` | POST | Upload a custom voice-style JSON (`multipart/form-data`) |
| `/docs` | GET | Interactive OpenAPI documentation |

### Built-in Voices (10)

| Category | Voices |
|---|---|
| **Male** | M1, M2, M3, M4, M5 |
| **Female** | F1, F2, F3, F4, F5 |

### Supported Languages (31)

`ar`, `bg`, `hr`, `cs`, `da`, `nl`, `en`, `et`, `fi`, `fr`, `de`, `el`, `hi`, `hu`, `id`, `it`, `ja`, `ko`, `lv`, `lt`, `pl`, `pt`, `ro`, `ru`, `sk`, `sl`, `es`, `sv`, `tr`, `uk`, `vi`

Pass `"lang": "na"` for automatic language detection.

### Quality and Speed Controls

| Parameter | Range | Default | Notes |
|---|---|---|---|
| `steps` | 5–12 | 8 | Higher = better quality, slower synthesis |
| `speed` | 0.7–2.0 | 1.05 | Higher = faster speech |

### Multilingual Example

```bash
# English
curl -X POST http://127.0.0.1:7788/v1/tts \
  -H 'content-type: application/json' \
  -d '{"text":"Hello from the Supertonic TTS system!","voice":"F1","lang":"en"}' \
  -o /tmp/gh-aw/agent/english.wav

# Spanish
curl -X POST http://127.0.0.1:7788/v1/tts \
  -H 'content-type: application/json' \
  -d '{"text":"¡Hola desde el sistema Supertonic!","voice":"F1","lang":"es"}' \
  -o /tmp/gh-aw/agent/spanish.wav

# Auto-detect language
curl -X POST http://127.0.0.1:7788/v1/tts \
  -H 'content-type: application/json' \
  -d '{"text":"こんにちは、スーパートニックです。","voice":"M2","lang":"na"}' \
  -o /tmp/gh-aw/agent/japanese.wav
```

### Upload Generated Audio as a Run Artifact

Stage the WAV file and call the `upload_artifact` safe-output tool:

```bash
# Stage the file for upload
mkdir -p "$RUNNER_TEMP/gh-aw/safeoutputs/upload-artifacts"
cp /tmp/gh-aw/agent/output.wav \
  "$RUNNER_TEMP/gh-aw/safeoutputs/upload-artifacts/voice-summary.wav"
```

Then call the safe-output tool with:
```json
{ "type": "upload_artifact", "path": "voice-summary.wav" }
```

The tool returns `slot_N_artifact_url` — a direct download link to the uploaded artifact.
Include this URL in any report or discussion to give readers access to the audio file.

### Stop the Server (Optional Cleanup)

```bash
kill $(cat /tmp/gh-aw/agent/supertonic.pid) 2>/dev/null || true
rm -f /tmp/gh-aw/agent/supertonic.pid /tmp/gh-aw/agent/supertonic.log
```
