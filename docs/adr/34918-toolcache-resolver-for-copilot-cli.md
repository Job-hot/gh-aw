# ADR-34918: Toolcache Resolver for Copilot CLI in `actions/setup`

**Date**: 2026-05-28
**Status**: Draft
**Deciders**: salmanmkc (PR #34918)

---

## Part 1 — Narrative (Human-Friendly)

### Context

Every Copilot-engine workflow currently runs `npm install -g @github/copilot` during `actions/setup`, adding ~10s of latency per job and creating a hard dependency on the npm registry being reachable. Hosted runner images have started baking `@github/copilot` into `$RUNNER_TOOL_CACHE/copilot-cli/<version>/<arch>/`, but `actions/setup` does not currently look there, so the cached binary is wasted. We also need the lookup to remain compatible with whichever gh-aw compiler version is in use, because not every `@github/copilot` release is compatible with every gh-aw release. The change must be transparent for non-Copilot workflows (no golden diff outside Copilot fixtures) and must not introduce new runtime dependencies on top of Node.

### Decision

We will add a small Node.js toolcache resolver (`actions/setup/js/install_copilot_cli.cjs`, ~270 LOC, zero deps) that runs from `setup.sh` whenever the compiler emits `INPUT_GH_AW_VERSION`. The resolver consults a compatibility matrix (live at `gh-aw-actions/.github/aw/compat.json` with a bundled `actions/setup/compat.json` fallback) to pick the highest cached `@github/copilot` version compatible with the running gh-aw compiler, prepends its `bin/` to `$GITHUB_PATH`, and exposes `copilot-cached` / `copilot-path` step outputs. The compiler then gates the bash `npm install -g @github/copilot` step on `steps.setup.outputs.copilot-cached != 'true'`, so the bash installer runs only on a cache miss.

### Alternatives Considered

#### Alternative 1: Keep the runtime `npm install -g @github/copilot` for every job

Continue the status quo — install the CLI every run. Rejected because it permanently pays the ~10s install cost on every Copilot job and leaves jobs exposed to npm registry outages, even though the runner image already ships a compatible CLI for free.

#### Alternative 2: Delegate to a third-party `setup-copilot` composite action via `actions/cache`

Wrap the resolution in a separate published action. Rejected because publishing and pinning another action increases surface area for supply-chain review, multiplies version bumps the gh-aw repo must track, and doesn't actually do anything the in-tree `setup.sh` cannot do directly with Node 24's built-in `fetch`/`fs`/`path`. Keeping the resolver in-tree also keeps the compat-matrix lookup adjacent to the compiler that emits `INPUT_GH_AW_VERSION`.

#### Alternative 3: Pin a single `@github/copilot` version per gh-aw release at compile time

Have the compiler hard-code the exact cached path it expects. Rejected because the cache contents on a runner image change asynchronously from gh-aw releases — a hard-coded version would either force a gh-aw release every time a new `@github/copilot` is baked in, or silently fall back to `npm install` whenever the pin drifts. A version *range* resolved at job time via a published compat matrix is the looser coupling we actually want.

### Consequences

#### Positive
- Saves ~10s per Copilot-engine workflow run on a cache hit, which is the common case on hosted runners.
- Eliminates npm-registry reachability as a setup-time failure mode on a cache hit.
- No new runtime dependencies — uses only Node 24 built-ins (`fetch`, `fs`, `path`).
- Backward-compatible: non-Copilot workflows are byte-identical in compiler output; Copilot workflows on cache miss fall back to the existing bash installer; Copilot workflows on a node-less self-hosted runner skip the resolver via `command -v node` and also fall back.

#### Negative
- Adds ~270 LOC of resolver plus 23 vitest cases that must be maintained alongside the compiler.
- Introduces a runtime dependency on `gh-aw-actions/.github/aw/compat.json` being kept up to date; the bundled fallback `compat.json` can drift from upstream and would need a deliberate refresh cadence.
- Each Copilot job now performs an HTTP fetch (5s timeout) against the compat-matrix URL during setup, adding a small but non-zero network call.
- Two code paths now exist for "install Copilot CLI" (cached resolver vs. bash installer), increasing the cognitive surface for anyone debugging setup failures.

#### Neutral
- Adds two new outputs (`copilot-cached`, `copilot-path`) to `actions/setup/action.yml`.
- Adds the `INPUT_GH_AW_VERSION` env var to the setup step, but only for `engineID == "copilot"`.
- The compiler-emitted Copilot installer step now carries an `if:` gate; this is visible in Copilot golden fixtures only.

---

## Part 2 — Normative Specification (RFC 2119)

> The key words **MUST**, **MUST NOT**, **REQUIRED**, **SHALL**, **SHALL NOT**, **SHOULD**, **SHOULD NOT**, **RECOMMENDED**, **MAY**, and **OPTIONAL** in this section are to be interpreted as described in [RFC 2119](https://www.rfc-editor.org/rfc/rfc2119).

### Compiler Behavior

1. When `engineID == "copilot"`, the compiler **MUST** emit `INPUT_GH_AW_VERSION=<compiler version>` on the `actions/setup` step.
2. When `engineID != "copilot"`, the compiler **MUST NOT** emit `INPUT_GH_AW_VERSION`, and the compiled YAML output **MUST** remain byte-identical to the pre-change output for non-Copilot fixtures.
3. The compiler **MUST** wrap the bash `npm install -g @github/copilot` installer step in a conditional `if: steps.setup.outputs.copilot-cached != 'true'` for Copilot workflows.

### Resolver Behavior

1. The resolver **MUST** run only when `INPUT_GH_AW_VERSION` is set, and `setup.sh` **MUST** guard the invocation with `command -v node` so it is a no-op when Node is unavailable.
2. The resolver **MUST** fetch the compat matrix from the published `gh-aw-actions/.github/aw/compat.json` URL with a timeout of 5 seconds via `AbortSignal.timeout`.
3. On any network failure, timeout, or unparseable response, the resolver **MUST** fall back to the bundled `actions/setup/compat.json` rather than fail.
4. The resolver **MUST** select the first matrix row whose `max-gh-aw` covers the running compiler version, treating `*` as always matching, and then **MUST** select the highest semver-comparable cached version in `$RUNNER_TOOL_CACHE/copilot-cli/` within `[min-agent, max-agent]`.
5. On a successful selection (cache hit), the resolver **MUST** append `<dir>/bin` to `$GITHUB_PATH` and **MUST** set step outputs `copilot-cached=true` and `copilot-path=<dir>`.
6. On any miss, error, malformed input, or missing marker/binary, the resolver **MUST** set `copilot-cached=false`, **MUST** exit with status `0`, and **MUST NOT** modify `$GITHUB_PATH`.
7. The resolver **MUST NOT** depend on any npm package outside the Node 24 standard library (specifically: only `fetch`, `fs`, `path` are permitted).

### Action Surface

1. `actions/setup/action.yml` **MUST** declare exactly two new outputs — `copilot-cached` and `copilot-path` — and **MUST NOT** declare any new inputs as part of this change.
2. The bundled fallback `actions/setup/compat.json` **MUST** be present at the path consumed by the resolver and **SHOULD** be kept in sync with the upstream `gh-aw-actions/.github/aw/compat.json`.

### Conformance

An implementation is considered conformant with this ADR if it satisfies all **MUST** and **MUST NOT** requirements above. Failure to meet any **MUST** or **MUST NOT** requirement constitutes non-conformance — in particular, any path by which a resolver error, network failure, or cache miss could fail the `actions/setup` step (rather than fall back to the bash installer) is non-conformant.

---

*This is a DRAFT ADR generated by the [Design Decision Gate](https://github.com/github/gh-aw/actions/runs/26582322159) workflow. The PR author must review, complete, and finalize this document before the PR can merge.*
