import { describe, it, expect, beforeEach, vi } from "vitest";

describe("parse_codex_log.cjs (Codex --json)", () => {
  let mockCore;
  let parseCodexLog;
  let transformCodexEntries;
  let normalizeUsage;

  beforeEach(async () => {
    mockCore = {
      debug: vi.fn(),
      info: vi.fn(),
      warning: vi.fn(),
      error: vi.fn(),
      setFailed: vi.fn(),
      setOutput: vi.fn(),
      summary: {
        addRaw: vi.fn().mockReturnThis(),
        write: vi.fn().mockResolvedValue(),
      },
    };
    global.core = mockCore;

    const module = await import("./parse_codex_log.cjs");
    parseCodexLog = module.parseCodexLog;
    transformCodexEntries = module.transformCodexEntries;
    normalizeUsage = module.normalizeUsage;
  });

  it("returns a helpful message when log has no JSON entries", () => {
    const result = parseCodexLog("OpenAI Codex v0.133.0\nworkdir: /tmp\nhello\n");
    expect(result.markdown).toContain("Log format not recognized as Codex");
    expect(result.logEntries).toEqual([]);
  });

  it("normalizes chat-like tool calls and tool results", () => {
    const logContent = [
      "[INFO] Containers started successfully",
      "non-json noise",
      JSON.stringify({ role: "assistant", content: "I'll list open issues." }),
      JSON.stringify({
        role: "assistant",
        content: null,
        tool_calls: [{ id: "call_1", type: "function", function: { name: "mcp__github__list_issues", arguments: '{"state":"open"}' } }],
      }),
      JSON.stringify({ role: "tool", tool_call_id: "call_1", content: '{"items":[{"number":1}]}' }),
      JSON.stringify({ usage: { prompt_tokens: 5, completion_tokens: 7 }, duration_ms: 1200, num_turns: 3 }),
    ].join("\n");

    const result = parseCodexLog(logContent);
    expect(result.logEntries.length).toBeGreaterThan(0);

    const toolUse = result.logEntries.find(e => e.type === "assistant" && e.message?.content?.some(c => c.type === "tool_use"));
    expect(toolUse).toBeDefined();
    expect(toolUse.message.content[0].name).toBe("mcp__github__list_issues");
    expect(toolUse.message.content[0].input).toEqual({ state: "open" });

    const toolResult = result.logEntries.find(e => e.type === "user" && e.message?.content?.some(c => c.type === "tool_result"));
    expect(toolResult).toBeDefined();
    expect(toolResult.message.content[0].tool_use_id).toBe("call_1");

    const trailingResult = result.logEntries[result.logEntries.length - 1];
    expect(trailingResult.type).toBe("result");
    expect(trailingResult.usage).toEqual({ input_tokens: 5, output_tokens: 7 });
    expect(trailingResult.duration_ms).toBe(1200);
    expect(trailingResult.num_turns).toBe(3);
  });

  it("merges streaming delta events into a single assistant text entry", () => {
    const rawEntries = [
      { type: "response.output_text.delta", delta: "Hello " },
      { type: "response.output_text.delta", delta: "world" },
    ];

    const logEntries = transformCodexEntries(rawEntries);
    const assistantTexts = logEntries.filter(e => e.type === "assistant" && e.message?.content?.[0]?.type === "text");
    expect(assistantTexts).toHaveLength(1);
    expect(assistantTexts[0].message.content[0].text).toBe("Hello world");
  });

  it("maps multiple usage shapes into canonical token usage", () => {
    expect(normalizeUsage({ input_tokens: 1, output_tokens: 2 })).toEqual({ input_tokens: 1, output_tokens: 2, cache_creation_input_tokens: 0, cache_read_input_tokens: 0 });
    expect(normalizeUsage({ prompt_tokens: 3, completion_tokens: 4 })).toEqual({ input_tokens: 3, output_tokens: 4 });
    expect(normalizeUsage({ total_tokens: 9 })).toEqual({ input_tokens: 0, output_tokens: 9 });
  });
});
