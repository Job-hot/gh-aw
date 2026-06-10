#!/usr/bin/env bash
set +o histexpand

# Start Safe Outputs MCP HTTP Server in a Docker container
# This script starts the safe-outputs MCP server inside a node:lts-bookworm container
# and waits for it to become ready.

set -e

# Ensure logs directory exists
mkdir -p /tmp/gh-aw/mcp-logs/safeoutputs

# Create initial server.log file for artifact upload
{
  echo "Safe Outputs MCP Server Log"
  echo "Start time: $(date)"
  echo "==========================================="
  echo ""
} > /tmp/gh-aw/mcp-logs/safeoutputs/server.log

# Build the list of -e flags for environment variables that are already exported in
# the step environment. This covers the static GH_AW_* / GITHUB_* / GH_HOST vars set
# above as well as any dynamic config-derived vars (GH_AW_INPUT_*, GH_AW_SECRET_*, etc.)
DOCKER_ENV_ARGS=()
while IFS= read -r key; do
  DOCKER_ENV_ARGS+=("-e" "$key")
done < <(compgen -e | grep -E "^(GH_AW_|GITHUB_|GH_HOST|DEBUG)$")

# Start the safe-outputs MCP server in a node:lts-bookworm container.
# node:lts-bookworm is based on buildpack-deps:bookworm which includes git,
# so no runtime package installation is needed.
# Mounts:
#   safeoutputs dir  - safe-outputs scripts and config files
#   logs dir         - write server logs back to the host
#   workspace        - enable git operations on the checked-out repository
#   gh binary        - enable gh CLI commands inside the container
SAFE_OUTPUTS_CONTAINER_ID=$(docker run -d \
  --network host \
  "${DOCKER_ENV_ARGS[@]}" \
  -v "${RUNNER_TEMP}/gh-aw/safeoutputs:${RUNNER_TEMP}/gh-aw/safeoutputs:rw" \
  -v "/tmp/gh-aw/mcp-logs/safeoutputs:/tmp/gh-aw/mcp-logs/safeoutputs:rw" \
  -v "${GITHUB_WORKSPACE}:${GITHUB_WORKSPACE}:rw" \
  -v "/usr/bin/gh:/usr/bin/gh:ro" \
  --workdir "${RUNNER_TEMP}/gh-aw/safeoutputs" \
  node:lts-bookworm \
  node mcp-server.cjs >> /tmp/gh-aw/mcp-logs/safeoutputs/server.log 2>&1)
echo "Started safe-outputs MCP server container: ${SAFE_OUTPUTS_CONTAINER_ID}"

# Wait for server to be ready (max 60 seconds)
echo "Waiting for safe-outputs MCP server to become ready..."
for i in $(seq 1 60); do
  if ! docker inspect -f '{{.State.Running}}' "${SAFE_OUTPUTS_CONTAINER_ID}" 2>/dev/null | grep -q true; then
    echo "ERROR: Safe-outputs container ${SAFE_OUTPUTS_CONTAINER_ID} has stopped"
    echo "Container logs:"
    docker logs "${SAFE_OUTPUTS_CONTAINER_ID}" 2>&1 || true
    echo "Server log contents:"
    cat /tmp/gh-aw/mcp-logs/safeoutputs/server.log || true
    exit 1
  fi

  if curl -s -f "http://localhost:${GH_AW_SAFE_OUTPUTS_PORT}/health" > /dev/null 2>&1; then
    echo "Safe Outputs MCP server is ready (attempt ${i}/60)"
    echo "::group::Server Log Contents"
    cat /tmp/gh-aw/mcp-logs/safeoutputs/server.log || true
    echo "::endgroup::"
    break
  fi

  if [ "${i}" -eq 60 ]; then
    echo "ERROR: Safe Outputs MCP server failed to start after 60 seconds"
    echo "Server log contents:"
    cat /tmp/gh-aw/mcp-logs/safeoutputs/server.log || true
    exit 1
  fi

  echo "Waiting for server... (attempt ${i}/60)"
  sleep 1
done

# Output the configuration for the MCP client
{
  echo "port=${GH_AW_SAFE_OUTPUTS_PORT}"
  echo "api_key=${GH_AW_SAFE_OUTPUTS_API_KEY@Q}"
} >> "$GITHUB_OUTPUT"
