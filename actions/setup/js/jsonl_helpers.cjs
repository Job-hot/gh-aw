// @ts-check

/**
 * Parse JSONL content into an array of entries.
 * Malformed lines are ignored so callers can safely consume partially written logs.
 *
 * @param {string} content
 * @param {(line: string) => boolean} [lineFilter] Optional predicate to pre-filter lines before JSON parsing.
 * @returns {unknown[]}
 */
function parseJsonlContent(content, lineFilter) {
  if (typeof content !== "string" || content.length === 0) {
    return [];
  }

  /** @type {unknown[]} */
  const entries = [];
  for (const rawLine of content.split("\n")) {
    const line = rawLine.trim();
    if (!line) {
      continue;
    }
    if (typeof lineFilter === "function" && !lineFilter(line)) {
      continue;
    }
    try {
      entries.push(JSON.parse(line));
    } catch {
      // Ignore malformed JSONL lines.
    }
  }
  return entries;
}

/**
 * Emit one dedicated JSONL event line to stderr for downstream failure-report parsing.
 * Best-effort: serialization failures are silently ignored.
 *
 * @param {Record<string, unknown>} event
 */
function emitJSONLEvent(event) {
  try {
    process.stderr.write(JSON.stringify(event) + "\n");
  } catch {
    // Best-effort diagnostics only; ignore serialization failures.
  }
}

/**
 * Parse a dedicated mixed-configured-models JSONL event from engine log content.
 *
 * Log lines may be prefixed with a timestamp or other non-JSON text
 * (e.g. "2026-04-27T21:45:00.080Z  {JSON}"), so the parser finds the first
 * "{" on each matching line rather than requiring the line to start with "{".
 *
 * @param {string} logContent
 * @returns {{engine: string, configured_models: string[], available_models: string[]} | null}
 */
function parseMixedConfiguredModelNamesEvent(logContent) {
  if (!logContent) return null;
  const lines = logContent.split("\n");
  for (const line of lines) {
    if (!line.includes('"awf.mixed_configured_model_names"')) {
      continue;
    }
    const jsonStart = line.indexOf("{");
    if (jsonStart === -1) continue;
    try {
      const parsed = JSON.parse(line.slice(jsonStart));
      if (parsed?.type !== "awf.mixed_configured_model_names") {
        continue;
      }
      const configuredModels = Array.isArray(parsed.configured_models)
        ? parsed.configured_models
            .map(m => String(m || "").trim())
            .filter(Boolean)
        : [];
      const availableModels = Array.isArray(parsed.available_models)
        ? parsed.available_models
            .map(m => String(m || "").trim())
            .filter(Boolean)
        : [];
      return {
        engine:
          typeof parsed.engine === "string" && parsed.engine.trim()
            ? parsed.engine.trim()
            : "AI",
        configured_models: configuredModels,
        available_models: availableModels,
      };
    } catch {
      // Ignore malformed lines and keep scanning.
    }
  }
  return null;
}

module.exports = {
  parseJsonlContent,
  emitJSONLEvent,
  parseMixedConfiguredModelNamesEvent,
};
