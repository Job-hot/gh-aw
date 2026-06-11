---
# Codebase Memory MCP — Structural Code Intelligence
# Wraps the codebase-memory-mcp server for fast code-graph queries,
# semantic search, and architecture analysis across 159 languages.
# The knowledge-graph index is stored in cache-memory so it survives
# between workflow runs; only incremental re-indexing is needed after that.
#
# See: https://github.com/DeusData/codebase-memory-mcp
#
# Usage:
#   network:
#     allowed:
#       - node                     # Required for npm install -g codebase-memory-mcp
#   tools:
#     cache-memory: true           # REQUIRED: already set by this import
#   imports:
#     - shared/mcp/codebase-memory-mcp.md

tools:
  cache-memory: true

mcp-servers:
  codebase-memory:
    command: "codebase-memory-mcp"
    args: []
    allowed:
      - index_repository
      - index_status
      - list_projects
      - search_graph
      - trace_path
      - detect_changes
      - query_graph
      - get_graph_schema
      - get_code_snippet
      - get_architecture
      - search_code
      - ingest_traces

steps:
  - name: Install codebase-memory-mcp
    run: |
      npm install -g codebase-memory-mcp@0.7.0 --ignore-scripts
      # Explicitly run the binary-download postinstall after reviewing it above.
      node "$(npm root -g)/codebase-memory-mcp/install.js"

  - name: Set up codebase-memory index cache
    run: |
      # Symlink ~/.cache/codebase-memory-mcp into cache-memory so the SQLite
      # knowledge-graph index persists automatically across workflow runs.
      mkdir -p /tmp/gh-aw/cache-memory/codebase-memory-mcp-store
      mkdir -p ~/.cache

      if [ ! -e ~/.cache/codebase-memory-mcp ]; then
        ln -s /tmp/gh-aw/cache-memory/codebase-memory-mcp-store ~/.cache/codebase-memory-mcp
        echo "🔗 Created ~/.cache/codebase-memory-mcp → cache-memory/codebase-memory-mcp-store"
      elif [ -d ~/.cache/codebase-memory-mcp ] && [ ! -L ~/.cache/codebase-memory-mcp ]; then
        # Plain directory present (e.g. first run after adding this import) — migrate
        if ! cp -r ~/.cache/codebase-memory-mcp/. /tmp/gh-aw/cache-memory/codebase-memory-mcp-store/ 2>&1; then
          echo "warning: migration copy failed; the existing store may be incomplete. Starting fresh." >&2
        fi
        rm -rf ~/.cache/codebase-memory-mcp
        ln -s /tmp/gh-aw/cache-memory/codebase-memory-mcp-store ~/.cache/codebase-memory-mcp
        echo "🔗 Migrated existing store → cache-memory/codebase-memory-mcp-store"
      else
        echo "✅ ~/.cache/codebase-memory-mcp already linked to cache-memory store"
      fi

  - name: Index repository (incremental)
    run: |
      # Index or incrementally update the knowledge graph.
      # On the first run a full graph is built; subsequent runs detect the
      # existing store and only re-index changed files (git-diff-driven).
      echo "Indexing ${GITHUB_WORKSPACE}..."
      if ! codebase-memory-mcp cli index_repository "{\"repo_path\": \"${GITHUB_WORKSPACE}\"}" 2>&1; then
        echo "error: index_repository failed — check the output above for details." >&2
        exit 1
      fi
      # index_status is best-effort: the server may not yet have flushed all
      # metadata, so a non-zero exit here is not a fatal condition.
      codebase-memory-mcp cli index_status "{\"repo_path\": \"${GITHUB_WORKSPACE}\"}" 2>&1 || \
        echo "warning: index_status returned a non-zero exit code (non-fatal)"
      echo "✅ Codebase memory index ready"
---

<!--
## Codebase Memory MCP

Provides fast structural code intelligence via codebase-memory-mcp with
159-language tree-sitter parsing, Hybrid LSP semantic type resolution, and
a persistent SQLite knowledge graph backed by GitHub Actions cache-memory.
The graph is built on the first run and updated incrementally on subsequent runs.

Documentation: https://github.com/DeusData/codebase-memory-mcp

## Available Tools

The following read-oriented tools are enabled by default:

- `index_repository` — Full or incremental index of the workspace
- `index_status` — Indexing status and node/edge counts for a project
- `list_projects` — All indexed projects with node/edge counts
- `search_graph` — Structural search by label, name pattern, file pattern, or degree
- `trace_path` — BFS call-graph traversal (inbound/outbound/both, depth 1–5)
- `detect_changes` — Map git diff to affected symbols with risk classification
- `query_graph` — Read-only openCypher-like graph queries (`MATCH`, `WHERE`, `RETURN`, ...)
- `get_graph_schema` — Node/edge schema, counts, and property definitions
- `get_code_snippet` — Retrieve source for a function by qualified name
- `get_architecture` — Architecture overview: languages, packages, routes, clusters, ADR
- `search_code` — Grep-like text search over indexed project files
- `ingest_traces` — Ingest runtime traces to validate `HTTP_CALLS` edges

## Excluded Write Tools

These tools are excluded from the default allowlist:

- `delete_project` — Permanently removes a project and all graph data
- `manage_adr` — Creates, updates, and deletes Architecture Decision Records

To enable `manage_adr`, add it explicitly in the importing workflow:

```yaml
mcp-servers:
  codebase-memory:
    allowed:
      - manage_adr
```

## Persistence

The knowledge-graph index lives at
`/tmp/gh-aw/cache-memory/codebase-memory-mcp-store/` (symlinked from
`~/.cache/codebase-memory-mcp/`). The `cache-memory` tool saves and restores
this directory across runs automatically.

## Requirements

The importing workflow must include:

- `node` in `network.allowed` — needed for `npm install -g codebase-memory-mcp`
- `tools.cache-memory: true` — already set by this import; do not disable it

## Example Queries

```
# Architecture overview (entry points, routes, hotspots, clusters)
get_architecture {}

# Find all HTTP handler functions
search_graph {"label": "Function", "name_pattern": ".*Handler.*"}

# Trace callers of a function, depth 3
trace_path {"function_name": "ProcessOrder", "direction": "inbound", "depth": 3}

# Cypher query — functions that call a specific method
query_graph {"query": "MATCH (f:Function)-[:CALLS]->(g:Function) WHERE g.name = 'SaveRecord' RETURN f.name LIMIT 20"}

# Symbols affected by the current git diff
detect_changes {"repo_path": "."}

# Text search within indexed files
search_code {"query": "TODO", "file_pattern": "*.go"}
```
-->
