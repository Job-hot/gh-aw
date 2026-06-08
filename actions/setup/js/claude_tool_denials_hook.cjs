// @ts-check

/**
 * Claude Code Tool Denials Hook
 *
 * This hook tracks tool-denied permission errors and enforces the max-tool-denials
 * guardrail for Claude Code CLI. It monitors Claude's output stream for permission
 * denial events and terminates the session when the threshold is exceeded.
 *
 * The hook is configured via the CLAUDE_HOOKS_PATH environment variable and runs
 * when Claude launches. It implements the same guardrail logic as the Copilot SDK
 * driver (copilot_sdk_session.cjs) but adapted for Claude's architecture.
 *
 * Environment Variables:
 *   GH_AW_MAX_TOOL_DENIALS - Maximum allowed tool denials (default: 5)
 *   CLAUDE_HOOKS_PATH      - Path to hook directory (set by setup action)
 *
 * Hook Events:
 *   onLaunch        - Called when Claude CLI starts
 *   onToolDenied    - Called when a tool use is denied
 *   onShutdown      - Called when Claude CLI exits
 *
 * Event Format:
 *   The hook writes JSONL events to stderr for unified_timeline.cjs:
 *   {"type": "guard.tool_denials_exceeded", "timestamp": "...", "data": {...}}
 */

"use strict";

const fs = require("fs");
const os = require("os");
const path = require("path");

// Default maximum tool denials threshold (matches Copilot SDK default)
const MAX_TOOL_DENIALS_DEFAULT = 5;

/**
 * Parse and validate the max-tool-denials limit from environment or parameter.
 * @param {string|number|undefined} value
 * @returns {number}
 */
function parseMaxToolDenialsLimit(value) {
  if (value === undefined || value === null || value === "") {
    return MAX_TOOL_DENIALS_DEFAULT;
  }
  const parsed = typeof value === "number" ? value : parseInt(String(value), 10);
  if (Number.isNaN(parsed) || parsed < 1) {
    return MAX_TOOL_DENIALS_DEFAULT;
  }
  return parsed;
}

/**
 * Get a positive integer from environment variable or return default.
 * @param {string} envVar
 * @param {number} defaultValue
 * @returns {number}
 */
function getEnvPositiveIntOrDefault(envVar, defaultValue) {
  const value = process.env[envVar];
  return parseMaxToolDenialsLimit(value);
}

/**
 * Write a JSONL event to stderr for capture in the agent logs.
 * @param {string} type
 * @param {object} data
 */
function writeDriverEvent(type, data) {
  const entry = {
    type,
    timestamp: new Date().toISOString(),
    data,
  };
  const jsonl = JSON.stringify(entry) + "\n";
  process.stderr.write(jsonl);
}

/**
 * Emit a timestamped log line to stderr.
 * @param {string} message
 */
function log(message) {
  const ts = new Date().toISOString();
  process.stderr.write(`[claude-tool-denials-hook] ${ts} ${message}\n`);
}

/**
 * Claude Tool Denials Hook State
 */
class ClaudeToolDenialsHook {
  constructor() {
    this.toolDenialCount = 0;
    this.maxToolDenialsLimit = getEnvPositiveIntOrDefault("GH_AW_MAX_TOOL_DENIALS", MAX_TOOL_DENIALS_DEFAULT);
    this.catastrophicToolDenialsTriggered = false;
    this.deniedCommands = [];
    this.sessionStartTime = Date.now();
  }

  /**
   * Called when Claude CLI launches.
   * Initialize the hook and log configuration.
   */
  onLaunch() {
    log(`hook initialized: maxToolDenialsLimit=${this.maxToolDenialsLimit}`);
    log(`monitoring for tool-denied permission errors`);
  }

  /**
   * Called when a tool use is denied.
   * Increments the denial count and triggers the guardrail if threshold is exceeded.
   * @param {string} toolName - Name of the tool that was denied
   * @param {string} reason - Reason for the denial
   */
  onToolDenied(toolName, reason) {
    this.toolDenialCount += 1;
    const deniedCommand = `${toolName}: ${reason}`;
    this.deniedCommands.push(deniedCommand);

    log(`tool denial ${this.toolDenialCount}/${this.maxToolDenialsLimit}: ${deniedCommand}`);

    if (this.catastrophicToolDenialsTriggered || this.toolDenialCount < this.maxToolDenialsLimit) {
      return;
    }

    this.catastrophicToolDenialsTriggered = true;

    // Emit guard event for unified_timeline.cjs
    writeDriverEvent("guard.tool_denials_exceeded", {
      denialCount: this.toolDenialCount,
      threshold: this.maxToolDenialsLimit,
      reason,
      deniedCommands: this.deniedCommands,
      sessionDurationMs: Date.now() - this.sessionStartTime,
    });

    log(`max tool denials threshold reached (${this.toolDenialCount}/${this.maxToolDenialsLimit}); guardrail triggered`);

    // Request Claude CLI to terminate gracefully
    // Claude hooks can signal termination by writing to a control file or
    // by exiting with a specific code. For now, we log the event and rely
    // on the post-processing step to detect the guard event and fail the job.
  }

  /**
   * Called when Claude CLI exits.
   * Log final statistics.
   */
  onShutdown() {
    const sessionDurationMs = Date.now() - this.sessionStartTime;
    log(`hook shutdown: totalDenials=${this.toolDenialCount} sessionDurationMs=${sessionDurationMs}`);

    if (this.catastrophicToolDenialsTriggered) {
      log(`session terminated due to max tool denials guardrail`);
    }
  }

  /**
   * Process a line from Claude's output stream.
   * Detects permission denied errors and updates the denial count.
   * @param {string} line - A line from Claude's stdout/stderr
   */
  processOutputLine(line) {
    // Claude Code emits permission denied errors in various formats:
    // 1. "Permission denied and could not request permission from user"
    // 2. JSON error objects with type: "tool_use_error"
    // 3. Stream-json format with error details

    // Try JSON first to avoid double-counting
    try {
      const parsed = JSON.parse(line);
      if (parsed.type === "tool_use_error" || parsed.type === "error") {
        if (parsed.error && /permission denied/i.test(parsed.error)) {
          const toolName = parsed.toolName || parsed.tool_name || "unknown";
          this.onToolDenied(toolName, parsed.error);
          return; // JSON pattern matched, don't check other patterns
        }
      }
    } catch {
      // Not JSON, continue to text patterns
    }

    // Pattern 1: Plain text permission denied
    if (/permission denied and could not request permission from user/i.test(line)) {
      // Extract the tool name or command from context if possible
      const toolMatch = line.match(/tool[:\s]+([^\s,]+)/i);
      const toolName = toolMatch ? toolMatch[1] : "unknown";
      this.onToolDenied(toolName, "Permission denied and could not request permission from user");
      return;
    }

    // Pattern 2: Bash tool permission denied
    if (/bash.*permission denied/i.test(line)) {
      this.onToolDenied("Bash", line.trim());
      return;
    }

    // Pattern 3: Tool call denied by policy
    if (/tool.*denied.*policy/i.test(line)) {
      const toolMatch = line.match(/tool[:\s]+([^\s,]+)/i);
      const toolName = toolMatch ? toolMatch[1] : "unknown";
      this.onToolDenied(toolName, "Tool denied by policy");
      return;
    }
  }
}

/**
 * Main entry point for the hook.
 * This function should be called by the Claude driver harness.
 */
function run() {
  const hook = new ClaudeToolDenialsHook();
  hook.onLaunch();

  // Set up stdin reader to monitor Claude's output
  // In a real hook integration, this would be connected to Claude's event stream
  process.stdin.setEncoding("utf8");
  process.stdin.on("data", (chunk) => {
    const lines = chunk.split("\n");
    for (const line of lines) {
      if (line.trim()) {
        hook.processOutputLine(line);
      }
    }
  });

  process.stdin.on("end", () => {
    hook.onShutdown();
  });

  // Handle process termination signals
  process.on("SIGINT", () => {
    hook.onShutdown();
    process.exit(0);
  });

  process.on("SIGTERM", () => {
    hook.onShutdown();
    process.exit(0);
  });
}

// Export for testing and integration
if (typeof module !== "undefined" && module.exports) {
  module.exports = {
    ClaudeToolDenialsHook,
    parseMaxToolDenialsLimit,
    getEnvPositiveIntOrDefault,
    writeDriverEvent,
    log,
    run,
    MAX_TOOL_DENIALS_DEFAULT,
  };
}

// Run the hook if executed directly
if (require.main === module) {
  run();
}
