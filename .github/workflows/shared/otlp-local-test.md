---
observability:
  otlp:
    endpoint: http://localhost:4318

pre-agent-steps:
  - name: Start local OTLP test server
    run: |
      mkdir -p /tmp/gh-aw

      cat > /tmp/gh-aw/otlp-test-server.py << 'PYEOF'
      import http.server
      import json

      JSONL = "/tmp/gh-aw/otlp-test.jsonl"
      PORT  = 4318

      class OTLPHandler(http.server.BaseHTTPRequestHandler):
          def do_POST(self):
              n = int(self.headers.get("Content-Length", 0))
              body = self.rfile.read(n)
              try:
                  data = json.loads(body)
              except Exception:
                  data = {"_raw": body.decode("utf-8", errors="replace")}
              with open(JSONL, "a") as f:
                  f.write(json.dumps(data) + "\n")
              self.send_response(200)
              self.send_header("Content-Type", "application/json")
              self.end_headers()
              self.wfile.write(b"{}")

      http.server.HTTPServer(("0.0.0.0", PORT), OTLPHandler).serve_forever()
      PYEOF

      python3 /tmp/gh-aw/otlp-test-server.py \
        >> /tmp/gh-aw/otlp-test-server.log 2>&1 &
      SERVER_PID=$!
      disown "$SERVER_PID"
      echo "$SERVER_PID" > /tmp/gh-aw/otlp-test-server.pid
      echo "OTLP test server started on port 4318 (PID: $SERVER_PID)"

  - name: Wait for OTLP test server to be ready
    run: |
      for i in {1..30}; do
        CODE=$(curl -sS -o /dev/null -w "%{http_code}" \
          --connect-timeout 2 --max-time 2 \
          -X POST -H "Content-Type: application/json" -d '{}' \
          http://localhost:4318/v1/traces || true)
        if [ -z "$CODE" ]; then CODE="curl_error"; fi
        [ "$CODE" = "200" ] && echo "OTLP test server is ready on port 4318" && break
        if [ "$i" = "30" ]; then
          echo "OTLP test server did not become ready after 30 seconds (last: $CODE)"
          cat /tmp/gh-aw/otlp-test-server.log || true
          exit 1
        fi
        echo "Waiting for OTLP test server ($i/30, status: $CODE)..."
        sleep 1
      done

post-steps:
  - name: Stop local OTLP test server
    if: always()
    run: |
      if [ -f /tmp/gh-aw/otlp-test-server.pid ]; then
        kill "$(cat /tmp/gh-aw/otlp-test-server.pid)" 2>/dev/null || true
        rm -f /tmp/gh-aw/otlp-test-server.pid
        echo "OTLP test server stopped"
      fi
---

<!--
## OTLP Local Test Server

Starts a minimal OTLP HTTP/JSON receiver for end-to-end testing of the OpenTelemetry
integration. No secrets or configuration parameters are required.

### What this import does

1. Configures `observability.otlp.endpoint` to export telemetry to `http://localhost:4318`.
2. Starts a Python HTTP server bound to `0.0.0.0:4318` in `pre-agent-steps`, before the
   agent container launches.
3. Waits up to 30 seconds for the server to accept requests; fails the step if it does not
   become ready in time.
4. Appends every received `/v1/traces` POST body as a JSON line to
   `/tmp/gh-aw/otlp-test.jsonl`.
5. Stops the server in `post-steps` (runs with `if: always()`) so cleanup occurs even when
   the workflow fails.

The server binds to `0.0.0.0` so it is reachable from the agent container. The agent
container uses `--network host`, meaning `localhost:4318` on the runner is visible to
processes inside the container.

### Usage

```yaml
imports:
  - shared/otlp-local-test.md
```

### Inspecting captured spans

After the agent completes, test steps can read `/tmp/gh-aw/otlp-test.jsonl`:

```bash
# Count received OTLP export batches
wc -l /tmp/gh-aw/otlp-test.jsonl

# Extract span names
jq -r '.resourceSpans[]?.scopeSpans[]?.spans[]?.name' /tmp/gh-aw/otlp-test.jsonl

# Filter for gh-aw spans
jq '.resourceSpans[]?.scopeSpans[]?.spans[]? | select(.name | startswith("gh-aw."))' \
  /tmp/gh-aw/otlp-test.jsonl

# Check that the conclusion span was exported
jq 'select(.resourceSpans[]?.scopeSpans[]?.spans[]?.name | test("conclusion"))' \
  /tmp/gh-aw/otlp-test.jsonl
```

### Combining with production OTLP

This import merges cleanly with `shared/otlp.md`. The compiler combines the endpoints so
telemetry is fanned out to both the local test server and any configured production
backends simultaneously.

### Limitations

- Port 4318 (the standard OTLP HTTP port) must be free on the runner.
- The `gh-aw.<job>.setup` span is emitted by `actions/setup` before this import's server
  starts, so it will not appear in `otlp-test.jsonl`. The conclusion span and any custom
  spans emitted via `logSpan` after the server starts are captured.
- The local JSONL mirror at `/tmp/gh-aw/otel.jsonl` (written unconditionally by the
  framework) is separate from `/tmp/gh-aw/otlp-test.jsonl` produced by the test server.
- This import is for integration testing only. For production observability use
  `shared/otlp.md` instead.
-->
