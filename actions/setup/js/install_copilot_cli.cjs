// @ts-check
// install_copilot_cli.cjs — zero-dependency Copilot CLI resolver
//
// Runs from actions/setup/setup.sh for Copilot-engine workflows.
// Looks for a cached, gh-aw-compatible build of @github/copilot in the runner
// tool cache. On a hit, appends the bin directory to $GITHUB_PATH and writes
// `copilot-cached=true` / `copilot-path=<dir>` to $GITHUB_OUTPUT so the
// compiler-emitted installer step skips itself. On any miss or error, writes
// `copilot-cached=false` and exits 0 — the installer step then runs as before.
//
// Design constraints (see ADR-10093):
//   - No third-party deps (cannot rely on @actions/tool-cache being present).
//   - Never throws / never exits non-zero — fall back to the existing installer.
//   - Resolve bundled compat.json via __dirname (script mode runs from a
//     non-cwd location).
//   - Fetch the live matrix from gh-aw-actions main (best-effort, 5s timeout)
//     and fall back to the bundled snapshot on any error.
//
// Matrix entry format (see actions/setup/compat.json):
//   { "max-gh-aw": "*"|<semver>, "min-agent": <semver>, "max-agent": <semver> }

require("./shim.cjs");

const fs = require("fs");
const path = require("path");

const COMPAT_URL = "https://raw.githubusercontent.com/github/gh-aw-actions/main/.github/aw/compat.json";
const FETCH_TIMEOUT_MS = 5000;

/**
 * @param {string} msg
 */
function log(msg) {
  core.info(`[install_copilot_cli] ${msg}`);
}

/**
 * @param {string} msg
 */
function logErr(msg) {
  core.warning(`[install_copilot_cli] ${msg}`);
}

/**
 * Parsed semantic version: `[major, minor, patch, prerelease]`.
 * `prerelease` is `""` when absent.
 * @typedef {[number, number, number, string]} ParsedSemver
 */

/**
 * Compat matrix row.
 * @typedef {{
 *   "max-gh-aw": string,
 *   "min-agent": string,
 *   "max-agent": string,
 * }} CompatRow
 */

/**
 * Parse a SemVer string into a comparable tuple. Returns null on malformed
 * input so callers can skip the entry rather than crash.
 * @param {unknown} v
 * @returns {ParsedSemver | null}
 */
function parseSemver(v) {
  if (typeof v !== "string") return null;
  const m = v.match(/^(\d+)\.(\d+)\.(\d+)(?:-([0-9A-Za-z.-]+))?$/);
  if (!m) return null;
  return [Number(m[1]), Number(m[2]), Number(m[3]), m[4] || ""];
}

/**
 * Compare two parsed SemVers. Returns -1/0/1. Treats any pre-release as lower
 * than its release counterpart (sufficient for our pinning use case).
 * @param {ParsedSemver} a
 * @param {ParsedSemver} b
 * @returns {-1 | 0 | 1}
 */
function cmpSemver(a, b) {
  for (let i = 0; i < 3; i++) {
    if (a[i] !== b[i]) return a[i] < b[i] ? -1 : 1;
  }
  if (a[3] === b[3]) return 0;
  if (!a[3]) return 1;
  if (!b[3]) return -1;
  return a[3] < b[3] ? -1 : 1;
}

/**
 * Does the matrix row's `max-gh-aw` cover the current gh-aw compiler version?
 * "*" always matches. Otherwise the compiler version must be <= max-gh-aw.
 * Unparseable compiler versions (e.g., "dev") are treated as matching only "*".
 * @param {unknown} row
 * @param {ParsedSemver | null} ghAwSemver
 * @returns {boolean}
 */
function rowMatchesGhAw(row, ghAwSemver) {
  const maxGhAw = row && typeof row === "object" ? /** @type {Record<string, unknown>} */ (row)["max-gh-aw"] : undefined;
  if (maxGhAw === "*") return true;
  if (!ghAwSemver) return false;
  const max = parseSemver(maxGhAw);
  if (!max) return false;
  return cmpSemver(ghAwSemver, max) <= 0;
}

/**
 * Fetch the live matrix from gh-aw-actions. Resolves to parsed JSON or null
 * on any error (network, timeout, non-200, malformed JSON). Never throws.
 * @returns {Promise<unknown>}
 */
async function fetchLiveMatrix() {
  try {
    const res = await fetch(COMPAT_URL, { signal: AbortSignal.timeout(FETCH_TIMEOUT_MS) });
    if (!res.ok) {
      logErr(`live matrix fetch returned HTTP ${res.status}`);
      return null;
    }
    return JSON.parse(await res.text());
  } catch (e) {
    logErr(`live matrix fetch failed: ${e instanceof Error ? e.message : String(e)}`);
    return null;
  }
}

/**
 * Load the bundled fallback matrix from disk, resolved via __dirname so script
 * mode (running from /tmp/gh-aw/actions-source/...) and dev/release mode both
 * find it next to the setup action.
 * @returns {unknown}
 */
function loadBundledMatrix() {
  try {
    const p = path.join(__dirname, "..", "compat.json");
    return JSON.parse(fs.readFileSync(p, "utf8"));
  } catch (e) {
    logErr(`bundled matrix load failed: ${e instanceof Error ? e.message : String(e)}`);
    return null;
  }
}

/**
 * Extract the copilot row list from a matrix document. Returns [] if the
 * document is malformed (treated as "no compatible versions").
 * @param {unknown} matrix
 * @returns {CompatRow[]}
 */
function copilotRows(matrix) {
  if (!matrix || typeof matrix !== "object") return [];
  const v1 = /** @type {Record<string, unknown>} */ (matrix)["agent-compat-v1"];
  if (!v1 || typeof v1 !== "object") return [];
  const rows = /** @type {Record<string, unknown>} */ (v1)["copilot"];
  return Array.isArray(rows) ? rows : [];
}

/**
 * Pick the resolution range [min, max] from the first row whose max-gh-aw
 * covers the current compiler version. Returns null when no row matches.
 * @param {CompatRow[]} rows
 * @param {ParsedSemver | null} ghAwSemver
 * @returns {{min: ParsedSemver, max: ParsedSemver} | null}
 */
function pickRange(rows, ghAwSemver) {
  for (const row of rows) {
    if (!rowMatchesGhAw(row, ghAwSemver)) continue;
    const min = parseSemver(row["min-agent"]);
    const max = parseSemver(row["max-agent"]);
    if (!min || !max) continue;
    return { min, max };
  }
  return null;
}

/**
 * Map process.arch to the runner-images tool-cache arch directory name.
 * @returns {string}
 */
function detectArch() {
  switch (process.arch) {
    case "x64":
      return "x64";
    case "arm64":
      return "arm64";
    default:
      return process.arch;
  }
}

/**
 * Find the highest cached Copilot CLI version in [min, max] under the runner
 * tool cache. Returns { version, dir, binDir } on hit, null on miss. Only
 * considers entries with a sibling .complete marker (matches @actions/tool-cache).
 * @param {string} toolCacheRoot
 * @param {string} arch
 * @param {{min: ParsedSemver, max: ParsedSemver}} range
 * @returns {{version: string, dir: string, binDir: string} | null}
 */
function findCachedCopilot(toolCacheRoot, arch, range) {
  const baseDir = path.join(toolCacheRoot, "copilot-cli");
  /** @type {string[]} */
  let entries;
  try {
    entries = fs.readdirSync(baseDir);
  } catch (e) {
    if (/** @type {NodeJS.ErrnoException} */ (e).code !== "ENOENT") {
      logErr(`tool cache scan failed: ${e instanceof Error ? e.message : String(e)}`);
    }
    return null;
  }

  /** @type {{parsed: ParsedSemver, version: string, dir: string, binDir: string} | null} */
  let best = null;
  for (const entry of entries) {
    const v = parseSemver(entry);
    if (!v) continue;
    if (cmpSemver(v, range.min) < 0) continue;
    if (cmpSemver(v, range.max) > 0) continue;

    const archDir = path.join(baseDir, entry, arch);
    const marker = `${archDir}.complete`;
    if (!fs.existsSync(marker)) continue;

    const binDir = path.join(archDir, "bin");
    const binFile = path.join(binDir, "copilot");
    if (!fs.existsSync(binFile)) continue;

    if (!best || cmpSemver(v, best.parsed) > 0) {
      best = { parsed: v, version: entry, dir: archDir, binDir };
    }
  }

  if (!best) return null;
  return { version: best.version, dir: best.dir, binDir: best.binDir };
}

/**
 * Append a line to a GitHub Actions runner file (e.g., $GITHUB_PATH or
 * $GITHUB_OUTPUT). No-ops when the path env var is unset so the resolver runs
 * in local tests without polluting the workflow.
 * @param {string} envVar
 * @param {string} line
 */
function appendRunnerFile(envVar, line) {
  const p = process.env[envVar];
  if (!p) return;
  try {
    fs.appendFileSync(p, line.endsWith("\n") ? line : line + "\n", "utf8");
  } catch (e) {
    logErr(`failed to append to ${envVar}: ${e instanceof Error ? e.message : String(e)}`);
  }
}

/**
 * @param {string} name
 * @param {string} value
 */
function writeOutput(name, value) {
  appendRunnerFile("GITHUB_OUTPUT", `${name}=${value}`);
}

/**
 * @param {string} dir
 */
function addToPath(dir) {
  appendRunnerFile("GITHUB_PATH", dir);
}

/**
 * @returns {Promise<void>}
 */
async function resolve() {
  const ghAwVersionRaw = process.env.INPUT_GH_AW_VERSION || "";
  const ghAwSemver = parseSemver(ghAwVersionRaw);
  if (ghAwVersionRaw && !ghAwSemver) {
    log(`gh-aw version "${ghAwVersionRaw}" is not SemVer; only wildcard rows will match`);
  }

  const toolCacheRoot = process.env.RUNNER_TOOL_CACHE || process.env.AGENT_TOOLSDIRECTORY;
  if (!toolCacheRoot) {
    log("RUNNER_TOOL_CACHE not set; treating as cache miss");
    writeOutput("copilot-cached", "false");
    writeOutput("copilot-path", "");
    return;
  }

  // Try live matrix first, fall back to bundled. Either may be null.
  let matrix = await fetchLiveMatrix();
  if (!matrix) {
    log("falling back to bundled compat.json");
    matrix = loadBundledMatrix();
  }
  const rows = copilotRows(matrix);
  if (rows.length === 0) {
    log("no copilot rows in compat matrix; treating as cache miss");
    writeOutput("copilot-cached", "false");
    writeOutput("copilot-path", "");
    return;
  }

  const range = pickRange(rows, ghAwSemver);
  if (!range) {
    log(`no compat row matches gh-aw version "${ghAwVersionRaw}"; treating as cache miss`);
    writeOutput("copilot-cached", "false");
    writeOutput("copilot-path", "");
    return;
  }

  const arch = detectArch();
  const hit = findCachedCopilot(toolCacheRoot, arch, range);
  if (!hit) {
    log(`no cached copilot in [${range.min.slice(0, 3).join(".")}, ${range.max.slice(0, 3).join(".")}] for arch ${arch}`);
    writeOutput("copilot-cached", "false");
    writeOutput("copilot-path", "");
    return;
  }

  log(`cache hit: copilot ${hit.version} at ${hit.binDir}`);
  addToPath(hit.binDir);
  writeOutput("copilot-cached", "true");
  writeOutput("copilot-path", hit.binDir);
}

// Top-level: never throw, always exit 0. Any unexpected error is logged and
// becomes a cache miss so the bash installer step takes over. Only runs when
// invoked as a script so the module can be require()d safely from tests.
if (require.main === module) {
  resolve()
    .catch(e => {
      logErr(`unexpected error: ${e && e.stack ? e.stack : e}`);
      try {
        writeOutput("copilot-cached", "false");
        writeOutput("copilot-path", "");
      } catch {
        // best effort
      }
    })
    .then(() => process.exit(0));
}

module.exports = {
  parseSemver,
  cmpSemver,
  rowMatchesGhAw,
  copilotRows,
  pickRange,
  detectArch,
  findCachedCopilot,
  resolve,
};
