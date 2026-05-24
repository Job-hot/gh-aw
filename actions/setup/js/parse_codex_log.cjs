// @ts-check
/// <reference types="@actions/github-script" />

const { createEngineLogParser, generateConversationMarkdown, generateInformationSection, formatInitializationSummary, formatToolUse, AWF_INFRA_LINE_RE } = require("./log_parser_shared.cjs");

const main = createEngineLogParser({
  parserName: "Codex",
  parseFunction: parseCodexLog,
  supportsDirectories: false,
});

/**
 * Parses Codex `codex exec --json` output (JSONL) from the agent log.
 *
 * The Codex CLI prints one JSON object per line. The exact schema varies across
 * releases, so we defensively normalize any JSON lines that look like:
 * - Canonical gh-aw entries (type: system/assistant/user/result)
 * - OpenAI chat-like messages (role + content/tool_calls)
 * - Tool events (tool_name/tool_id/parameters/output)
 * - Streaming delta events (type ends with ".delta")
 *
 * The output is normalized into the canonical `logEntries` format consumed by
 * `generateConversationMarkdown` and the Copilot CLI-style renderer.
 *
 * @param {string} logContent
 * @returns {{markdown: string, logEntries: Array<any>, mcpFailures: Array<string>, maxTurnsHit: boolean}}
 */
function parseCodexLog(logContent) {
  if (!logContent) {
    return {
      markdown: "## Agent Log Summary\n\nNo log content provided.\n",
      logEntries: [],
      mcpFailures: [],
      maxTurnsHit: false,
    };
  }

  const rawEntries = extractJsonLines(logContent);
  if (rawEntries.length === 0) {
    return {
      markdown: "## Agent Log Summary\n\nLog format not recognized as Codex `--json` output.\n",
      logEntries: [],
      mcpFailures: [],
      maxTurnsHit: false,
    };
  }

  const logEntries = transformCodexEntries(rawEntries);

  const conversationResult = generateConversationMarkdown(logEntries, {
    formatToolCallback: (toolUse, toolResult) => formatToolUse(toolUse, toolResult, { includeDetailedParameters: true }),
    formatInitCallback: initEntry => formatInitializationSummary(initEntry, { includeSlashCommands: false }),
  });

  // Prefer a trailing result entry for stats, else fall back to the last raw entry.
  const lastEntry = findLastStatsEntry(logEntries) || null;

  let markdown = conversationResult.markdown;
  markdown += generateInformationSection(lastEntry);

  return {
    markdown,
    logEntries,
    mcpFailures: [],
    maxTurnsHit: false,
  };
}

/**
 * Extract JSON objects from log content, treating each JSON line as a candidate entry.
 * @param {string} logContent
 * @returns {Array<any>}
 */
function extractJsonLines(logContent) {
  /** @type {Array<any>} */
  const entries = [];
  const lines = logContent.split("\n");

  for (const line of lines) {
    const trimmed = line.trim();
    if (!trimmed) continue;
    if (AWF_INFRA_LINE_RE.test(trimmed)) continue;

    // Most Codex `--json` output is one object per line.
    if (!trimmed.startsWith("{") && !trimmed.startsWith("[")) {
      continue;
    }

    try {
      const parsed = JSON.parse(trimmed);
      if (Array.isArray(parsed)) {
        entries.push(...parsed);
      } else if (parsed && typeof parsed === "object") {
        entries.push(parsed);
      }
    } catch (_e) {
      // Ignore non-JSON lines (or truncated JSON).
    }
  }

  return entries;
}

/**
 * Convert Codex JSON entries into canonical log entries for the shared renderer.
 * @param {Array<any>} rawEntries
 * @returns {Array<any>}
 */
function transformCodexEntries(rawEntries) {
  /** @type {Array<any>} */
  const logEntries = [];

  for (const raw of rawEntries) {
    const normalized = normalizeCodexEntry(raw, logEntries);
    if (!normalized) continue;
    if (Array.isArray(normalized)) {
      logEntries.push(...normalized);
    } else {
      logEntries.push(normalized);
    }
  }

  // Ensure we have a trailing stats entry when Codex provides usage in a non-canonical shape.
  const trailingStats = findTrailingUsage(rawEntries);
  if (trailingStats && !findLastStatsEntry(logEntries)) {
    logEntries.push(trailingStats);
  }

  return logEntries;
}

/**
 * Normalize a single Codex JSON entry.
 * @param {any} raw
 * @param {Array<any>} existingEntries - Used for delta merging.
 * @returns {any|Array<any>|null}
 */
function normalizeCodexEntry(raw, existingEntries) {
  if (!raw || typeof raw !== "object") return null;

  // 1) Already canonical gh-aw log entries (Claude/Copilot style).
  if (typeof raw.type === "string" && ["system", "assistant", "user", "result"].includes(raw.type)) {
    return raw;
  }

  // 2) Pi-style JSONL entries (tool_use/tool_result/init/result/assistant).
  if (typeof raw.type === "string" && ["init", "assistant", "tool_use", "tool_result", "result"].includes(raw.type)) {
    return normalizePiLikeEntry(raw, existingEntries);
  }

  // 3) OpenAI chat-like message objects (role + content/tool_calls).
  const message = extractMessageEnvelope(raw);
  if (message) {
    return normalizeChatLikeMessage(message, existingEntries);
  }

  // 4) Streaming delta events (best-effort).
  if (typeof raw.type === "string" && raw.type.endsWith(".delta")) {
    return normalizeDeltaEvent(raw, existingEntries);
  }

  return null;
}

/**
 * @param {any} raw
 * @param {Array<any>} existingEntries
 * @returns {any|null}
 */
function normalizePiLikeEntry(raw, existingEntries) {
  if (raw.type === "init") {
    return {
      type: "system",
      subtype: "init",
      model: raw.model,
      session_id: raw.session_id,
      tools: raw.tools,
      cwd: raw.cwd,
      mcp_servers: raw.mcp_servers,
    };
  }

  if (raw.type === "assistant") {
    const text = typeof raw.content === "string" ? raw.content : typeof raw.text === "string" ? raw.text : "";
    if (!text.trim()) return null;
    return appendAssistantText(existingEntries, text, raw.delta === true);
  }

  if (raw.type === "tool_use") {
    return {
      type: "assistant",
      message: {
        content: [
          {
            type: "tool_use",
            id: raw.tool_id || raw.id,
            name: raw.tool_name || raw.name,
            input: raw.parameters || raw.input || {},
          },
        ],
      },
    };
  }

  if (raw.type === "tool_result") {
    const toolUseID = raw.tool_id || raw.tool_use_id || raw.tool_call_id;
    const content = typeof raw.output === "string" ? raw.output : typeof raw.content === "string" ? raw.content : JSON.stringify(raw.output ?? raw.content ?? "");
    return {
      type: "user",
      message: {
        content: [
          {
            type: "tool_result",
            tool_use_id: toolUseID,
            content,
            is_error: raw.status != null && raw.status !== "success" && raw.status !== "ok",
          },
        ],
      },
    };
  }

  if (raw.type === "result") {
    // Permit both Codex-like and Pi-like stats shapes.
    const stats = raw.stats && typeof raw.stats === "object" ? raw.stats : raw;
    const usage = normalizeUsage(stats.usage || stats);
    if (!usage) return null;
    return {
      type: "result",
      usage,
      duration_ms: stats.duration_ms || stats.durationMs || undefined,
      num_turns: stats.num_turns || stats.turns || undefined,
      total_cost_usd: stats.total_cost_usd || stats.totalCostUsd || undefined,
    };
  }

  return null;
}

/**
 * Extract an OpenAI chat-like message object from a variety of Codex envelopes.
 * @param {any} raw
 * @returns {any|null}
 */
function extractMessageEnvelope(raw) {
  // Some Codex builds wrap messages in { message: { role, content, ... } }.
  if (raw.message && typeof raw.message === "object" && typeof raw.message.role === "string") {
    return raw.message;
  }
  // Other variants use { data: { role, content, ... } }.
  if (raw.data && typeof raw.data === "object" && typeof raw.data.role === "string") {
    return raw.data;
  }
  // Some emit role at top-level.
  if (typeof raw.role === "string") {
    return raw;
  }
  return null;
}

/**
 * Normalize an OpenAI chat-like message into canonical entries.
 * @param {any} msg
 * @param {Array<any>} existingEntries
 * @returns {any|Array<any>|null}
 */
function normalizeChatLikeMessage(msg, existingEntries) {
  const role = typeof msg.role === "string" ? msg.role : "";

  if (role === "tool") {
    const toolUseID = msg.tool_call_id || msg.tool_use_id || msg.tool_call?.id;
    const content = typeof msg.content === "string" ? msg.content : JSON.stringify(msg.content ?? "");
    return {
      type: "user",
      message: {
        content: [
          {
            type: "tool_result",
            tool_use_id: toolUseID,
            content,
            is_error: false,
          },
        ],
      },
    };
  }

  /** @type {Array<any>} */
  const assistantContent = [];

  const thinking = extractThinkingText(msg);
  if (thinking) {
    assistantContent.push({ type: "thinking", thinking });
  }

  const text = extractTextContent(msg);
  if (text) {
    assistantContent.push({ type: "text", text });
  }

  const toolUses = extractToolUses(msg);
  if (toolUses.length > 0) {
    assistantContent.push(...toolUses);
  }

  if (assistantContent.length === 0) {
    return null;
  }

  // If this message is purely a streaming delta text chunk, merge it.
  if (assistantContent.length === 1 && assistantContent[0].type === "text" && msg.delta === true) {
    return appendAssistantText(existingEntries, assistantContent[0].text, true);
  }

  // For tool-only messages, keep a single assistant entry with tool_use blocks.
  return {
    type: role === "system" ? "system" : "assistant",
    ...(role === "system" ? { subtype: "init" } : {}),
    message: role === "system" ? undefined : { content: assistantContent },
  };
}

/**
 * Best-effort normalization of streaming delta events.
 * @param {any} raw
 * @param {Array<any>} existingEntries
 * @returns {any|null}
 */
function normalizeDeltaEvent(raw, existingEntries) {
  const deltaText = typeof raw.delta === "string" ? raw.delta : typeof raw.text === "string" ? raw.text : null;
  if (!deltaText) return null;
  return appendAssistantText(existingEntries, deltaText, true);
}

/**
 * Append or merge assistant text into the canonical log entry list.
 * @param {Array<any>} existingEntries
 * @param {string} text
 * @param {boolean} merge
 * @returns {any|null}
 */
function appendAssistantText(existingEntries, text, merge) {
  const trimmed = text;
  if (!trimmed || !trimmed.trim()) return null;

  const last = existingEntries[existingEntries.length - 1];
  if (merge && isMergeableAssistantTextEntry(last)) {
    last.message.content[0].text += trimmed;
    return null;
  }

  return {
    type: "assistant",
    message: {
      content: [{ type: "text", text: trimmed }],
    },
  };
}

/**
 * @param {any} entry
 * @returns {boolean}
 */
function isMergeableAssistantTextEntry(entry) {
  return entry && entry.type === "assistant" && entry.message && Array.isArray(entry.message.content) && entry.message.content.length === 1 && entry.message.content[0].type === "text";
}

/**
 * @param {any} msg
 * @returns {string}
 */
function extractThinkingText(msg) {
  if (typeof msg.thinking === "string") return msg.thinking;
  if (typeof msg.reasoning === "string") return msg.reasoning;
  if (Array.isArray(msg.content)) {
    const parts = msg.content
      .map(part => {
        if (!part || typeof part !== "object") return "";
        if (part.type === "thinking" && typeof part.thinking === "string") return part.thinking;
        if (part.type === "reasoning" && typeof part.text === "string") return part.text;
        return "";
      })
      .filter(Boolean);
    return parts.join("");
  }
  return "";
}

/**
 * @param {any} msg
 * @returns {string}
 */
function extractTextContent(msg) {
  if (typeof msg.content === "string") return msg.content;
  if (typeof msg.text === "string") return msg.text;
  if (Array.isArray(msg.content)) {
    const parts = msg.content
      .map(part => {
        if (!part || typeof part !== "object") return "";
        if (typeof part.text === "string") return part.text;
        if (typeof part.content === "string") return part.content;
        return "";
      })
      .filter(Boolean);
    return parts.join("");
  }
  return "";
}

/**
 * Extract tool calls from message-like shapes.
 * @param {any} msg
 * @returns {Array<any>}
 */
function extractToolUses(msg) {
  /** @type {Array<any>} */
  const toolUses = [];

  const toolCalls = Array.isArray(msg.tool_calls) ? msg.tool_calls : [];
  for (const toolCall of toolCalls) {
    if (!toolCall || typeof toolCall !== "object") continue;
    const id = toolCall.id || toolCall.tool_call_id;
    const fn = toolCall.function && typeof toolCall.function === "object" ? toolCall.function : null;
    const name = fn && typeof fn.name === "string" ? fn.name : typeof toolCall.name === "string" ? toolCall.name : null;
    const args = fn && typeof fn.arguments === "string" ? fn.arguments : typeof toolCall.arguments === "string" ? toolCall.arguments : null;
    if (!name) continue;
    toolUses.push({
      type: "tool_use",
      id: id || `tool_${toolUses.length}`,
      name,
      input: parseMaybeJSON(args) || {},
    });
  }

  // Older "function_call" shape.
  if (msg.function_call && typeof msg.function_call === "object" && typeof msg.function_call.name === "string") {
    toolUses.push({
      type: "tool_use",
      id: msg.function_call.id || `tool_${toolUses.length}`,
      name: msg.function_call.name,
      input: parseMaybeJSON(msg.function_call.arguments) || {},
    });
  }

  return toolUses;
}

/**
 * @param {any} value
 * @returns {Record<string, any>|null}
 */
function parseMaybeJSON(value) {
  if (typeof value !== "string" || !value.trim()) return null;
  try {
    const parsed = JSON.parse(value);
    return parsed && typeof parsed === "object" && !Array.isArray(parsed) ? parsed : null;
  } catch (_e) {
    return null;
  }
}

/**
 * Find the last canonical stats entry to feed into generateInformationSection.
 * @param {Array<any>} entries
 * @returns {any|null}
 */
function findLastStatsEntry(entries) {
  for (let i = entries.length - 1; i >= 0; i--) {
    const entry = entries[i];
    if (!entry || typeof entry !== "object") continue;
    if (entry.type === "result") return entry;
    if (entry.usage || entry.duration_ms || entry.num_turns) return entry;
  }
  return null;
}

/**
 * Attempt to build a trailing canonical result entry from raw usage-like entries.
 * @param {Array<any>} rawEntries
 * @returns {any|null}
 */
function findTrailingUsage(rawEntries) {
  for (let i = rawEntries.length - 1; i >= 0; i--) {
    const raw = rawEntries[i];
    if (!raw || typeof raw !== "object") continue;
    const usage = normalizeUsage(raw.usage || raw.stats || raw);
    if (!usage) continue;
    return {
      type: "result",
      usage,
      duration_ms: raw.duration_ms || raw.durationMs || raw.stats?.duration_ms || undefined,
      num_turns: raw.num_turns || raw.turns || raw.stats?.turns || undefined,
      total_cost_usd: raw.total_cost_usd || raw.totalCostUsd || undefined,
    };
  }
  return null;
}

/**
 * Normalize multiple common token usage shapes into the canonical {input_tokens, output_tokens}.
 * @param {any} usage
 * @returns {any|null}
 */
function normalizeUsage(usage) {
  if (!usage || typeof usage !== "object") return null;

  // Canonical shape already.
  if (typeof usage.input_tokens === "number" || typeof usage.output_tokens === "number") {
    return {
      input_tokens: usage.input_tokens || 0,
      output_tokens: usage.output_tokens || 0,
      cache_creation_input_tokens: usage.cache_creation_input_tokens || 0,
      cache_read_input_tokens: usage.cache_read_input_tokens || 0,
    };
  }

  // OpenAI chat completions / responses-like shapes.
  const hasPromptTokens = typeof usage.prompt_tokens === "number" || typeof usage.inputTokens === "number";
  const hasCompletionTokens = typeof usage.completion_tokens === "number" || typeof usage.outputTokens === "number";
  if (hasPromptTokens || hasCompletionTokens) {
    const promptTokens = typeof usage.prompt_tokens === "number" ? usage.prompt_tokens : typeof usage.inputTokens === "number" ? usage.inputTokens : 0;
    const completionTokens = typeof usage.completion_tokens === "number" ? usage.completion_tokens : typeof usage.outputTokens === "number" ? usage.outputTokens : 0;
    return {
      input_tokens: typeof promptTokens === "number" ? promptTokens : 0,
      output_tokens: typeof completionTokens === "number" ? completionTokens : 0,
    };
  }

  // Aggregate totals.
  if (typeof usage.total_tokens === "number") {
    return {
      input_tokens: 0,
      output_tokens: usage.total_tokens,
    };
  }

  return null;
}

// Export for testing
if (typeof module !== "undefined" && module.exports) {
  module.exports = {
    main,
    parseCodexLog,
    extractJsonLines,
    transformCodexEntries,
    normalizeUsage,
  };
}
