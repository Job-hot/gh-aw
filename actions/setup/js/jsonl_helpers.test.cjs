import { describe, it, expect, vi } from "vitest";
import { parseJsonlContent, emitJSONLEvent, parseMixedConfiguredModelNamesEvent } from "./jsonl_helpers.cjs";

describe("jsonl_helpers", () => {
  describe("parseJsonlContent", () => {
    it("returns parsed JSON entries and skips malformed lines", () => {
      const parsed = parseJsonlContent(['{"event":"token_steering"}', "not-json", "", "   ", '{"event":"request"}'].join("\n"));

      expect(parsed).toEqual([{ event: "token_steering" }, { event: "request" }]);
    });

    it("returns empty array for non-string or empty content", () => {
      expect(parseJsonlContent("")).toEqual([]);
      expect(parseJsonlContent(/** @type {any} */ null)).toEqual([]);
      expect(parseJsonlContent(/** @type {any} */ undefined)).toEqual([]);
    });

    it("supports optional line pre-filtering before JSON parsing", () => {
      const parsed = parseJsonlContent(['{"event":"token_steering"}', '{"event":"request"}', '{"event":"model_steering"}'].join("\n"), line => line.includes("steering"));

      expect(parsed).toEqual([{ event: "token_steering" }, { event: "model_steering" }]);
    });
  });

  describe("emitJSONLEvent", () => {
    it("writes a JSON-serialized line to stderr", () => {
      const write = vi.spyOn(process.stderr, "write").mockImplementation(() => true);
      try {
        emitJSONLEvent({ type: "test.event", value: 1 });
        expect(write).toHaveBeenCalledWith('{"type":"test.event","value":1}\n');
      } finally {
        write.mockRestore();
      }
    });

    it("does not throw on non-serializable values", () => {
      const write = vi.spyOn(process.stderr, "write").mockImplementation(() => true);
      try {
        const circular = /** @type {any} */ ({});
        circular.self = circular;
        expect(() => emitJSONLEvent(circular)).not.toThrow();
      } finally {
        write.mockRestore();
      }
    });
  });

  describe("parseMixedConfiguredModelNamesEvent", () => {
    it("returns null for empty or missing content", () => {
      expect(parseMixedConfiguredModelNamesEvent("")).toBeNull();
      expect(parseMixedConfiguredModelNamesEvent(/** @type {any} */ null)).toBeNull();
    });

    it("parses a plain JSON line without prefix", () => {
      const line = JSON.stringify({
        type: "awf.mixed_configured_model_names",
        engine: "copilot",
        configured_models: ["gpt-5.4", "claude-sonnet-4.6"],
        available_models: ["gpt-4.1", "gpt-5.4"],
        retry_disabled: true,
      });
      const result = parseMixedConfiguredModelNamesEvent(line);
      expect(result).not.toBeNull();
      expect(result?.engine).toBe("copilot");
      expect(result?.configured_models).toEqual(["gpt-5.4", "claude-sonnet-4.6"]);
      expect(result?.available_models).toEqual(["gpt-4.1", "gpt-5.4"]);
    });

    it("parses a JSON line with a timestamp prefix", () => {
      const payload = JSON.stringify({
        type: "awf.mixed_configured_model_names",
        engine: "codex",
        configured_models: ["model-a", "model-b"],
        available_models: ["model-a"],
        retry_disabled: true,
      });
      const line = `2026-04-27T21:45:00.080Z  ${payload}`;
      const result = parseMixedConfiguredModelNamesEvent(line);
      expect(result).not.toBeNull();
      expect(result?.engine).toBe("codex");
      expect(result?.configured_models).toEqual(["model-a", "model-b"]);
    });

    it("returns null when no matching event is present", () => {
      const content = ['{"type":"other.event","value":1}', "not json at all", '{"type":"awf.something_else"}'].join("\n");
      expect(parseMixedConfiguredModelNamesEvent(content)).toBeNull();
    });

    it("defaults engine to AI when engine field is absent", () => {
      const line = JSON.stringify({
        type: "awf.mixed_configured_model_names",
        configured_models: ["model-x", "model-y"],
        available_models: [],
        retry_disabled: true,
      });
      const result = parseMixedConfiguredModelNamesEvent(line);
      expect(result?.engine).toBe("AI");
    });

    it("skips malformed JSON lines and continues scanning", () => {
      const good = JSON.stringify({
        type: "awf.mixed_configured_model_names",
        engine: "copilot",
        configured_models: ["m1", "m2"],
        available_models: [],
        retry_disabled: true,
      });
      const content = ['{"type":"awf.mixed_configured_model_names" BROKEN}', good].join("\n");
      const result = parseMixedConfiguredModelNamesEvent(content);
      expect(result?.engine).toBe("copilot");
    });
  });
});
