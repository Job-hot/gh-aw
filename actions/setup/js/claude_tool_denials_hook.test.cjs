// @ts-check

const { describe, it, beforeEach, afterEach } = require("node:test");
const assert = require("node:assert");
const {
  ClaudeToolDenialsHook,
  parseMaxToolDenialsLimit,
  getEnvPositiveIntOrDefault,
  MAX_TOOL_DENIALS_DEFAULT,
} = require("./claude_tool_denials_hook.cjs");

describe("claude_tool_denials_hook", () => {
  describe("parseMaxToolDenialsLimit", () => {
    it("returns default when value is undefined", () => {
      assert.strictEqual(parseMaxToolDenialsLimit(undefined), MAX_TOOL_DENIALS_DEFAULT);
    });

    it("returns default when value is null", () => {
      assert.strictEqual(parseMaxToolDenialsLimit(null), MAX_TOOL_DENIALS_DEFAULT);
    });

    it("returns default when value is empty string", () => {
      assert.strictEqual(parseMaxToolDenialsLimit(""), MAX_TOOL_DENIALS_DEFAULT);
    });

    it("returns default when value is NaN", () => {
      assert.strictEqual(parseMaxToolDenialsLimit("invalid"), MAX_TOOL_DENIALS_DEFAULT);
    });

    it("returns default when value is less than 1", () => {
      assert.strictEqual(parseMaxToolDenialsLimit(0), MAX_TOOL_DENIALS_DEFAULT);
      assert.strictEqual(parseMaxToolDenialsLimit(-5), MAX_TOOL_DENIALS_DEFAULT);
    });

    it("parses valid numeric strings", () => {
      assert.strictEqual(parseMaxToolDenialsLimit("10"), 10);
      assert.strictEqual(parseMaxToolDenialsLimit("100"), 100);
    });

    it("handles numeric values", () => {
      assert.strictEqual(parseMaxToolDenialsLimit(10), 10);
      assert.strictEqual(parseMaxToolDenialsLimit(100), 100);
    });
  });

  describe("getEnvPositiveIntOrDefault", () => {
    const originalValue = process.env.TEST_ENV_VAR;
    const originalGHAW = process.env.GH_AW_MAX_TOOL_DENIALS;

    beforeEach(() => {
      delete process.env.TEST_ENV_VAR;
      delete process.env.GH_AW_MAX_TOOL_DENIALS;
    });

    afterEach(() => {
      if (originalValue !== undefined) {
        process.env.TEST_ENV_VAR = originalValue;
      } else {
        delete process.env.TEST_ENV_VAR;
      }
      if (originalGHAW !== undefined) {
        process.env.GH_AW_MAX_TOOL_DENIALS = originalGHAW;
      } else {
        delete process.env.GH_AW_MAX_TOOL_DENIALS;
      }
    });

    it("returns default when env var is not set", () => {
      // Note: parseMaxToolDenialsLimit is used internally and defaults to MAX_TOOL_DENIALS_DEFAULT (5)
      // when the value is invalid, so the actual return is 5, not the passed default of 42
      assert.strictEqual(getEnvPositiveIntOrDefault("TEST_ENV_VAR", MAX_TOOL_DENIALS_DEFAULT), MAX_TOOL_DENIALS_DEFAULT);
    });

    it("parses valid env var value", () => {
      process.env.TEST_ENV_VAR = "100";
      assert.strictEqual(getEnvPositiveIntOrDefault("TEST_ENV_VAR", MAX_TOOL_DENIALS_DEFAULT), 100);
      delete process.env.TEST_ENV_VAR;
    });

    it("returns default for invalid env var value", () => {
      process.env.TEST_ENV_VAR = "invalid";
      // parseMaxToolDenialsLimit returns MAX_TOOL_DENIALS_DEFAULT for invalid values
      assert.strictEqual(getEnvPositiveIntOrDefault("TEST_ENV_VAR", MAX_TOOL_DENIALS_DEFAULT), MAX_TOOL_DENIALS_DEFAULT);
      delete process.env.TEST_ENV_VAR;
    });
  });

  describe("ClaudeToolDenialsHook", () => {
    /** @type {ClaudeToolDenialsHook} */
    let hook;
    /** @type {string[]} */
    let stderrWrites;
    /** @type {typeof process.stderr.write} */
    let originalWrite;
    const originalGHAW = process.env.GH_AW_MAX_TOOL_DENIALS;

    beforeEach(() => {
      // Reset environment
      delete process.env.GH_AW_MAX_TOOL_DENIALS;

      // Capture stderr writes
      stderrWrites = [];
      originalWrite = process.stderr.write;
      // @ts-ignore - Mock stderr.write
      process.stderr.write = (chunk) => {
        stderrWrites.push(String(chunk));
        return true;
      };

      hook = new ClaudeToolDenialsHook();
    });

    afterEach(() => {
      // Restore stderr
      process.stderr.write = originalWrite;
      // Restore environment
      if (originalGHAW !== undefined) {
        process.env.GH_AW_MAX_TOOL_DENIALS = originalGHAW;
      } else {
        delete process.env.GH_AW_MAX_TOOL_DENIALS;
      }
    });

    it("initializes with default max denials limit", () => {
      assert.strictEqual(hook.maxToolDenialsLimit, MAX_TOOL_DENIALS_DEFAULT);
      assert.strictEqual(hook.toolDenialCount, 0);
      assert.strictEqual(hook.catastrophicToolDenialsTriggered, false);
    });

    it("respects GH_AW_MAX_TOOL_DENIALS environment variable", () => {
      process.env.GH_AW_MAX_TOOL_DENIALS = "10";
      const customHook = new ClaudeToolDenialsHook();
      assert.strictEqual(customHook.maxToolDenialsLimit, 10);
    });

    it("increments denial count when tool is denied", () => {
      hook.onToolDenied("Bash", "Permission denied");
      assert.strictEqual(hook.toolDenialCount, 1);
      assert.strictEqual(hook.deniedCommands.length, 1);
      assert.strictEqual(hook.deniedCommands[0], "Bash: Permission denied");
    });

    it("does not trigger guardrail before threshold", () => {
      for (let i = 0; i < MAX_TOOL_DENIALS_DEFAULT - 1; i++) {
        hook.onToolDenied(`Tool${i}`, "Permission denied");
      }
      assert.strictEqual(hook.catastrophicToolDenialsTriggered, false);
    });

    it("triggers guardrail when threshold is reached", () => {
      for (let i = 0; i < MAX_TOOL_DENIALS_DEFAULT; i++) {
        hook.onToolDenied(`Tool${i}`, "Permission denied");
      }
      assert.strictEqual(hook.catastrophicToolDenialsTriggered, true);

      // Check that guard event was emitted
      const guardEvent = stderrWrites.find((line) => line.includes("guard.tool_denials_exceeded"));
      assert.ok(guardEvent, "Expected guard.tool_denials_exceeded event to be emitted");

      const parsed = JSON.parse(guardEvent);
      assert.strictEqual(parsed.type, "guard.tool_denials_exceeded");
      assert.strictEqual(parsed.data.denialCount, MAX_TOOL_DENIALS_DEFAULT);
      assert.strictEqual(parsed.data.threshold, MAX_TOOL_DENIALS_DEFAULT);
    });

    it("does not emit duplicate guard events", () => {
      // Trigger threshold
      for (let i = 0; i < MAX_TOOL_DENIALS_DEFAULT + 5; i++) {
        hook.onToolDenied(`Tool${i}`, "Permission denied");
      }

      // Count guard events
      const guardEvents = stderrWrites.filter((line) => line.includes("guard.tool_denials_exceeded"));
      assert.strictEqual(guardEvents.length, 1, "Expected exactly one guard event");
    });

    describe("processOutputLine", () => {
      beforeEach(() => {
        // Reset hook state for each processOutputLine test
        hook.toolDenialCount = 0;
        hook.deniedCommands = [];
        hook.catastrophicToolDenialsTriggered = false;
      });

      it("detects plain text permission denied errors", () => {
        hook.processOutputLine("Error: Permission denied and could not request permission from user");
        assert.strictEqual(hook.toolDenialCount, 1);
      });

      it("detects JSON tool_use_error", () => {
        const jsonError = JSON.stringify({
          type: "tool_use_error",
          toolName: "Bash",
          error: "Permission denied",
        });
        hook.processOutputLine(jsonError);
        assert.strictEqual(hook.toolDenialCount, 1);
      });

      it("detects bash permission denied errors", () => {
        hook.processOutputLine("bash: /usr/bin/foo: Permission denied");
        assert.strictEqual(hook.toolDenialCount, 1);
      });

      it("detects tool denied by policy errors", () => {
        hook.processOutputLine("Tool Edit denied by policy");
        assert.strictEqual(hook.toolDenialCount, 1);
      });

      it("ignores non-error lines", () => {
        hook.processOutputLine("Normal output line");
        hook.processOutputLine('{"type": "assistant.message", "content": "Hello"}');
        assert.strictEqual(hook.toolDenialCount, 0);
      });

      it("handles multiple permission denials in sequence", () => {
        hook.processOutputLine("Permission denied and could not request permission from user");
        hook.processOutputLine("bash: Permission denied");
        hook.processOutputLine("Tool Edit denied by policy");
        assert.strictEqual(hook.toolDenialCount, 3);
      });
    });

    describe("onLaunch and onShutdown", () => {
      it("logs initialization on launch", () => {
        hook.onLaunch();
        const launchLog = stderrWrites.find((line) => line.includes("hook initialized"));
        assert.ok(launchLog, "Expected launch log");
      });

      it("logs statistics on shutdown", () => {
        hook.onToolDenied("Bash", "Permission denied");
        hook.onShutdown();
        const shutdownLog = stderrWrites.find((line) => line.includes("hook shutdown"));
        assert.ok(shutdownLog, "Expected shutdown log");
        assert.ok(shutdownLog.includes("totalDenials=1"), "Expected denial count in shutdown log");
      });
    });
  });
});
