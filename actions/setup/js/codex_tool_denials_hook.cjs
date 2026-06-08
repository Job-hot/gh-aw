// @ts-check

/**
 * Codex tool-denials guardrail hook.
 *
 * The Codex CLI can emit repeated "tool denied" / "permission denied by workflow tool permissions"
 * diagnostics when running with non-interactive approval policies. In some failure modes the agent
 * can loop indefinitely attempting blocked tools. This hook provides a small, reusable guardrail
 * that the Codex harness can plug into to stop execution once a maximum number of tool-denied
 * events has been observed.
 */

"use strict";

const fs = require("fs");
const path = require("path");

const MAX_TOOL_DENIALS_DEFAULT = 5;

/**
 * @param {unknown} value
 * @returns {number | undefined}
 */
function parseStrictPositiveInteger(value) {
  if (typeof value === "number" && Number.isSafeInteger(value) && value > 0) {
    return value;
  }
  if (typeof value === "string") {
    const trimmed = value.trim();
    if (/^\d+$/.test(trimmed)) {
      const parsed = Number.parseInt(trimmed, 10);
      if (Number.isSafeInteger(parsed) && parsed > 0) {
        return parsed;
      }
    }
  }
  return undefined;
}

/**
 * @param {unknown} value
 * @returns {number}
 */
function parseMaxToolDenialsLimit(value) {
  return parseStrictPositiveInteger(value) ?? MAX_TOOL_DENIALS_DEFAULT;
}

/**
 * Attempt to extract a concise denial reason from a single output line.
 *
 * Supports both plain-text harness logs and JSONL events.
 *
 * @param {string} line
 * @returns {string | null}
 */
function extractToolDeniedReason(line) {
  const trimmed = (line || "").trim();
  if (!trimmed) return null;

  // JSONL support: best-effort parse and look for common denial fields.
  if (trimmed.startsWith("{") && trimmed.endsWith("}")) {
    try {
      const parsed = JSON.parse(trimmed);
      if (parsed && typeof parsed === "object") {
        // Common shapes seen across agent harnesses:
        // - { type: "...", data: { reason: "..." } }
        // - { error: { message: "...", code: "tool_denied" } }
        // - { message: "...tool denied..." }
        const message =
          typeof parsed.message === "string"
            ? parsed.message
            : parsed.error && typeof parsed.error.message === "string"
              ? parsed.error.message
              : parsed.data && typeof parsed.data.reason === "string"
                ? parsed.data.reason
                : parsed.data && typeof parsed.data.message === "string"
                  ? parsed.data.message
                  : "";

        const code = parsed.error && typeof parsed.error.code === "string" ? parsed.error.code : parsed.data && typeof parsed.data.code === "string" ? parsed.data.code : "";

        const type = typeof parsed.type === "string" ? parsed.type : "";

        const haystack = `${type} ${code} ${message}`.toLowerCase();
        if (haystack.includes("tool") && haystack.includes("denied")) {
          return message.trim() || "tool denied";
        }
        if (haystack.includes("permission denied") && (haystack.includes("tool") || haystack.includes("workflow tool"))) {
          return message.trim() || "permission denied by workflow tool permissions";
        }
      }
    } catch {
      // Not JSON; fall through to text patterns.
    }
  }

  // Text patterns (keep these conservative to avoid counting filesystem EACCES).
  const patterns = [
    // Copilot SDK driver style.
    { re: /\btool denial \d+\/\d+:\s*(.+?)\s*$/i, group: 1 },
    // Permission handler style.
    { re: /\bpermission denied by workflow tool permissions:\s*(.+?)\s*$/i, group: 1 },
    // Generic.
    { re: /\btool denied\b:?[\s-]*(.+?)\s*$/i, group: 1 },
  ];

  for (const p of patterns) {
    const m = trimmed.match(p.re);
    if (m && m[p.group] && String(m[p.group]).trim()) {
      return String(m[p.group]).trim();
    }
  }

  // Last resort: only count explicit "tool denied" lines without an extractable suffix.
  if (/\btool denied\b/i.test(trimmed)) {
    return "tool denied";
  }
  return null;
}

/**
 * @typedef {{
 *   denialCount: number,
 *   threshold: number,
 *   reason: string,
 * }} ToolDenialsExceededEvent
 */

/**
 * @param {string} eventsPath
 * @param {ToolDenialsExceededEvent} event
 */
function appendToolDenialsExceededEvent(eventsPath, event) {
  if (!eventsPath) return;
  const dir = path.dirname(eventsPath);
  try {
    fs.mkdirSync(dir, { recursive: true });
  } catch {
    // best-effort
  }
  const entry = { type: "guard.tool_denials_exceeded", timestamp: new Date().toISOString(), data: event };
  try {
    fs.appendFileSync(eventsPath, JSON.stringify(entry) + "\n", "utf8");
  } catch {
    // best-effort
  }
}

/**
 * Create a guard instance that tracks tool-denied output lines and triggers once.
 *
 * @param {{
 *   threshold: number,
 *   eventsPath?: string,
 *   logger?: (message: string) => void,
 * }} options
 */
function createToolDenialsGuard(options) {
  const threshold = options && Number.isFinite(options.threshold) && options.threshold > 0 ? Math.floor(options.threshold) : MAX_TOOL_DENIALS_DEFAULT;
  const eventsPath = options && typeof options.eventsPath === "string" ? options.eventsPath : "";
  const logger = options && typeof options.logger === "function" ? options.logger : null;

  let denialCount = 0;
  let triggered = false;
  /** @type {string} */
  let lastReason = "";

  /**
   * @param {string} line
   * @returns {{denied: boolean, triggered: boolean, event?: ToolDenialsExceededEvent}}
   */
  function observeLine(line) {
    const reason = extractToolDeniedReason(line);
    if (!reason) {
      return { denied: false, triggered };
    }

    denialCount += 1;
    lastReason = reason;
    if (logger) {
      logger(`tool denial ${denialCount}/${threshold}: ${reason}`);
    }

    if (triggered || denialCount < threshold) {
      return { denied: true, triggered };
    }

    triggered = true;
    const event = { denialCount, threshold, reason: lastReason || "tool denied" };
    appendToolDenialsExceededEvent(eventsPath, event);
    return { denied: true, triggered, event };
  }

  return {
    observeLine,
    getState: () => ({ denialCount, threshold, triggered, lastReason }),
  };
}

module.exports = {
  MAX_TOOL_DENIALS_DEFAULT,
  parseStrictPositiveInteger,
  parseMaxToolDenialsLimit,
  extractToolDeniedReason,
  createToolDenialsGuard,
};
