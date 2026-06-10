// @ts-check

const { parseJsonlContent } = require("./jsonl_helpers.cjs");

/**
 * Pre-filter pattern: only parse lines that contain the word "steering".
 * This avoids JSON.parse on unrelated log entries.
 *
 * @type {RegExp}
 */
const STEERING_EVENT_PATTERN = /steering/i;

/** Message prefix emitted by the AWF api-proxy for token budget warnings. */
const AWF_TOKEN_WARNING_PREFIX = "[AWF TOKEN WARNING]";

/** Message prefix emitted by the AWF api-proxy for time budget warnings. */
const AWF_TIME_WARNING_PREFIX = "[AWF TIME WARNING]";

/**
 * Resolve an event name from a firewall proxy event entry.
 *
 * Supports four top-level field variants used across AWF api-proxy log schema versions:
 *   - `event`:      `{ event: "token_steering", ... }`
 *   - `type`:       `{ type: "model_steering", ... }`
 *   - `event_name`: `{ event_name: "timeout_steering", ... }`
 *   - `eventName`:  `{ eventName: "timeout_steering", ... }`
 *
 * Also handles the legacy nested payload schema:
 *   - `{ payload: { event: "steering" } }`
 *
 * @param {unknown} entry
 * @returns {string}
 */
function getApiProxyEventName(entry) {
  if (!entry || typeof entry !== "object" || Array.isArray(entry)) {
    return "";
  }
  if ("event" in entry && typeof entry.event === "string") {
    return entry.event;
  }
  if ("type" in entry && typeof entry.type === "string") {
    return entry.type;
  }
  if ("event_name" in entry && typeof entry.event_name === "string") {
    return entry.event_name;
  }
  if ("eventName" in entry && typeof entry.eventName === "string") {
    return entry.eventName;
  }
  if ("payload" in entry) {
    const payload = entry.payload;
    if (payload && typeof payload === "object" && !Array.isArray(payload)) {
      if ("event" in payload && typeof payload.event === "string") {
        return payload.event;
      }
      if ("type" in payload && typeof payload.type === "string") {
        return payload.type;
      }
    }
  }
  return "";
}

/**
 * Count steering events in proxy event-log JSONL content.
 *
 * Known steering event names: "steering", "token_steering", "model_steering".
 * Any event whose name is exactly "steering" or ends with "_steering" is counted.
 *
 * @param {string} jsonlContent
 * @returns {number}
 */
function countSteeringEventsInApiProxyJsonl(jsonlContent) {
  let count = 0;
  for (const parsed of parseJsonlContent(jsonlContent, line => STEERING_EVENT_PATTERN.test(line))) {
    const eventName = getApiProxyEventName(parsed).toLowerCase();
    // Known steering events: "steering", "token_steering", "model_steering".
    if (eventName === "steering" || eventName.endsWith("_steering")) {
      count += 1;
    }
  }
  return count;
}

/**
 * Extracts spec-compliant AWF token/timeout steering entries from api-proxy JSONL content.
 *
 * Only entries whose event name is "token_steering" or "timeout_steering" and whose
 * message starts with the corresponding AWF warning prefix are returned. This mirrors
 * the filtering logic in the Go `isSteeringEvent` helper in `pkg/cli/token_usage.go`.
 *
 * @param {string} jsonlContent
 * @returns {Array<{eventName: string, message: string, timestamp: string}>}
 */
function extractSteeringEntriesFromApiProxyJsonl(jsonlContent) {
  /** @type {Array<{eventName: string, message: string, timestamp: string}>} */
  const entries = [];
  for (const parsed of parseJsonlContent(jsonlContent, line => STEERING_EVENT_PATTERN.test(line))) {
    const eventName = getApiProxyEventName(parsed).toLowerCase();
    if (!parsed || typeof parsed !== "object" || Array.isArray(parsed)) continue;
    const obj = /** @type {Record<string, unknown>} */ parsed;
    const message = typeof obj["message"] === "string" ? obj["message"] : "";
    const isTokenSteering = eventName === "token_steering" && message.startsWith(AWF_TOKEN_WARNING_PREFIX);
    const isTimeoutSteering = eventName === "timeout_steering" && message.startsWith(AWF_TIME_WARNING_PREFIX);
    if (!isTokenSteering && !isTimeoutSteering) continue;
    const timestamp = typeof obj["timestamp"] === "string" ? obj["timestamp"] : "";
    entries.push({ eventName, message, timestamp });
  }
  return entries;
}

module.exports = {
  STEERING_EVENT_PATTERN,
  AWF_TOKEN_WARNING_PREFIX,
  AWF_TIME_WARNING_PREFIX,
  getApiProxyEventName,
  countSteeringEventsInApiProxyJsonl,
  extractSteeringEntriesFromApiProxyJsonl,
};
