import { describe, it, expect } from "vitest";
import { getApiProxyEventName, countSteeringEventsInApiProxyJsonl, extractSteeringEntriesFromApiProxyJsonl, AWF_TOKEN_WARNING_PREFIX, AWF_TIME_WARNING_PREFIX } from "./steering_helpers.cjs";

describe("steering_helpers", () => {
  describe("getApiProxyEventName", () => {
    it("returns empty string for non-object inputs", () => {
      expect(getApiProxyEventName(null)).toBe("");
      expect(getApiProxyEventName(undefined)).toBe("");
      expect(getApiProxyEventName("string")).toBe("");
      expect(getApiProxyEventName(42)).toBe("");
      expect(getApiProxyEventName([])).toBe("");
    });

    it("returns top-level event field", () => {
      expect(getApiProxyEventName({ event: "token_steering", request_id: "r1" })).toBe("token_steering");
    });

    it("returns top-level type field when event is absent", () => {
      expect(getApiProxyEventName({ type: "model_steering", request_id: "r2" })).toBe("model_steering");
    });

    it("returns top-level event_name field when event and type are absent", () => {
      expect(getApiProxyEventName({ event_name: "timeout_steering", request_id: "r3" })).toBe("timeout_steering");
    });

    it("returns top-level eventName field when event, type, and event_name are absent", () => {
      expect(getApiProxyEventName({ eventName: "timeout_steering", request_id: "r4" })).toBe("timeout_steering");
    });

    it("prefers event over type and event_name", () => {
      expect(getApiProxyEventName({ event: "token_steering", type: "other", event_name: "another" })).toBe("token_steering");
    });

    it("prefers type over event_name when event is absent", () => {
      expect(getApiProxyEventName({ type: "model_steering", event_name: "other" })).toBe("model_steering");
    });

    it("prefers event_name over eventName when event and type are absent", () => {
      expect(getApiProxyEventName({ event_name: "timeout_steering", eventName: "other" })).toBe("timeout_steering");
    });

    it("returns payload.event when top-level fields are absent", () => {
      expect(getApiProxyEventName({ payload: { event: "steering" }, request_id: "r5" })).toBe("steering");
    });

    it("returns payload.type when payload.event is absent", () => {
      expect(getApiProxyEventName({ payload: { type: "token_steering" }, request_id: "r6" })).toBe("token_steering");
    });

    it("returns empty string when payload is not a plain object", () => {
      expect(getApiProxyEventName({ payload: null })).toBe("");
      expect(getApiProxyEventName({ payload: [] })).toBe("");
      expect(getApiProxyEventName({ payload: "str" })).toBe("");
    });

    it("returns empty string when no recognized field is present", () => {
      expect(getApiProxyEventName({ request_id: "r7", status: "ok" })).toBe("");
    });
  });

  describe("countSteeringEventsInApiProxyJsonl", () => {
    it("counts events with exact 'steering' name", () => {
      const content = '{"event":"steering","request_id":"r1"}\n';
      expect(countSteeringEventsInApiProxyJsonl(content)).toBe(1);
    });

    it("counts events ending with _steering", () => {
      const content = ['{"event":"token_steering"}', '{"event":"model_steering"}'].join("\n");
      expect(countSteeringEventsInApiProxyJsonl(content)).toBe(2);
    });

    it("matches event names case-insensitively", () => {
      const content = ['{"event":"TOKEN_STEERING"}', '{"event":"Steering"}'].join("\n");
      expect(countSteeringEventsInApiProxyJsonl(content)).toBe(2);
    });

    it("counts steering events via top-level type field", () => {
      const content = '{"type":"token_steering","request_id":"r2"}\n';
      expect(countSteeringEventsInApiProxyJsonl(content)).toBe(1);
    });

    it("counts steering events via payload.event field", () => {
      const content = '{"payload":{"event":"token_steering"},"request_id":"r3"}\n';
      expect(countSteeringEventsInApiProxyJsonl(content)).toBe(1);
    });

    it("counts steering events via payload.type field", () => {
      const content = '{"payload":{"type":"model_steering"},"request_id":"r4"}\n';
      expect(countSteeringEventsInApiProxyJsonl(content)).toBe(1);
    });

    it("ignores non-steering events", () => {
      const content = ['{"event":"request"}', '{"event":"response"}', '{"type":"auth"}'].join("\n");
      expect(countSteeringEventsInApiProxyJsonl(content)).toBe(0);
    });

    it("ignores malformed JSONL lines", () => {
      const content = ['{"event":"steering"}', "not-json", '{"event":"token_steering"}'].join("\n");
      expect(countSteeringEventsInApiProxyJsonl(content)).toBe(2);
    });

    it("returns 0 for empty content", () => {
      expect(countSteeringEventsInApiProxyJsonl("")).toBe(0);
    });

    it("counts steering events via event_name field", () => {
      const content = '{"event_name":"token_steering","message":"[AWF TOKEN WARNING] test"}\n';
      expect(countSteeringEventsInApiProxyJsonl(content)).toBe(1);
    });

    it("counts steering events via eventName field", () => {
      const content = '{"eventName":"timeout_steering","message":"[AWF TIME WARNING] test"}\n';
      expect(countSteeringEventsInApiProxyJsonl(content)).toBe(1);
    });
  });

  describe("extractSteeringEntriesFromApiProxyJsonl", () => {
    const TOKEN_MSG = `${AWF_TOKEN_WARNING_PREFIX} You have used 80% of your effective token budget.`;
    const TIMEOUT_MSG = `${AWF_TIME_WARNING_PREFIX} You have used 80% of your allotted run time.`;

    it("returns empty array for empty content", () => {
      expect(extractSteeringEntriesFromApiProxyJsonl("")).toEqual([]);
    });

    it("extracts a token_steering entry via event field", () => {
      const content = `{"event":"token_steering","message":"${TOKEN_MSG}","timestamp":"2024-01-15T10:05:00.000Z"}`;
      const entries = extractSteeringEntriesFromApiProxyJsonl(content);
      expect(entries).toHaveLength(1);
      expect(entries[0]).toEqual({ eventName: "token_steering", message: TOKEN_MSG, timestamp: "2024-01-15T10:05:00.000Z" });
    });

    it("extracts a timeout_steering entry via type field", () => {
      const content = `{"type":"timeout_steering","message":"${TIMEOUT_MSG}","timestamp":"2024-01-15T10:06:00.000Z"}`;
      const entries = extractSteeringEntriesFromApiProxyJsonl(content);
      expect(entries).toHaveLength(1);
      expect(entries[0].eventName).toBe("timeout_steering");
      expect(entries[0].message).toBe(TIMEOUT_MSG);
    });

    it("extracts a token_steering entry via event_name field", () => {
      const content = `{"event_name":"token_steering","message":"${TOKEN_MSG}"}`;
      const entries = extractSteeringEntriesFromApiProxyJsonl(content);
      expect(entries).toHaveLength(1);
      expect(entries[0].eventName).toBe("token_steering");
    });

    it("extracts a timeout_steering entry via eventName field", () => {
      const content = `{"eventName":"timeout_steering","message":"${TIMEOUT_MSG}"}`;
      const entries = extractSteeringEntriesFromApiProxyJsonl(content);
      expect(entries).toHaveLength(1);
      expect(entries[0].eventName).toBe("timeout_steering");
    });

    it("ignores token_steering entry with wrong message prefix", () => {
      const content = `{"event":"token_steering","message":"warn 90%"}`;
      expect(extractSteeringEntriesFromApiProxyJsonl(content)).toEqual([]);
    });

    it("ignores timeout_steering entry with wrong message prefix", () => {
      const content = `{"event":"timeout_steering","message":"[AWF TOKEN WARNING] wrong prefix"}`;
      expect(extractSteeringEntriesFromApiProxyJsonl(content)).toEqual([]);
    });

    it("ignores non-spec event names even with valid message prefix", () => {
      const content = `{"event":"budget_steering","message":"${TOKEN_MSG}"}`;
      expect(extractSteeringEntriesFromApiProxyJsonl(content)).toEqual([]);
    });

    it("ignores non-steering events", () => {
      const content = ['{"event":"request","message":"foo"}', '{"event":"response","message":"bar"}'].join("\n");
      expect(extractSteeringEntriesFromApiProxyJsonl(content)).toEqual([]);
    });

    it("returns empty timestamp string when timestamp is absent", () => {
      const content = `{"event":"token_steering","message":"${TOKEN_MSG}"}`;
      const entries = extractSteeringEntriesFromApiProxyJsonl(content);
      expect(entries[0].timestamp).toBe("");
    });

    it("extracts multiple entries in order", () => {
      const lines = [`{"event":"token_steering","message":"${TOKEN_MSG}","timestamp":"2024-01-15T10:05:00Z"}`, `{"event":"timeout_steering","message":"${TIMEOUT_MSG}","timestamp":"2024-01-15T10:06:00Z"}`];
      const entries = extractSteeringEntriesFromApiProxyJsonl(lines.join("\n"));
      expect(entries).toHaveLength(2);
      expect(entries[0].eventName).toBe("token_steering");
      expect(entries[1].eventName).toBe("timeout_steering");
    });

    it("ignores malformed JSONL lines", () => {
      const content = [`{"event":"token_steering","message":"${TOKEN_MSG}"}`, "not-json"].join("\n");
      expect(extractSteeringEntriesFromApiProxyJsonl(content)).toHaveLength(1);
    });
  });
});
