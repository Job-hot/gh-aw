// @ts-check
"use strict";

const fs = require("fs");

const AW_INFO_PATH = "/tmp/gh-aw/aw_info.json";
const AGENT_OUTPUT_PATH = "/tmp/gh-aw/agent_output.json";
const AGENT_USAGE_PATH = "/tmp/gh-aw/agent_usage.json";
const AGENT_STDIO_LOG_PATH = "/tmp/gh-aw/agent-stdio.log";
const GATEWAY_EVENT_PATHS = ["/tmp/gh-aw/mcp-logs/gateway.jsonl", "/tmp/gh-aw/mcp-logs/rpc-messages.jsonl"];

function readJSONIfExists(path) {
  try {
    return JSON.parse(fs.readFileSync(path, "utf8"));
  } catch {
    return null;
  }
}

function countBlockedRequests() {
  let total = 0;

  for (const path of GATEWAY_EVENT_PATHS) {
    try {
      const lines = fs.readFileSync(path, "utf8").split("\n");
      for (const raw of lines) {
        const line = raw.trim();
        if (!line) continue;
        try {
          const entry = JSON.parse(line);
          if (entry && entry.type === "DIFC_FILTERED") {
            total += 1;
          }
        } catch {
          // Skip malformed lines.
        }
      }
    } catch {
      // Missing gateway logs are normal for many runs.
    }
  }

  return total;
}

function uniqueCreatedItemTypes(items) {
  const types = new Set();

  for (const item of items) {
    if (item && typeof item.type === "string" && item.type.trim() !== "") {
      types.add(item.type);
    }
  }

  return [...types].sort();
}

function readAgentRuntimeMetrics() {
  const metrics = { turns: undefined, estimatedCostUsd: undefined, warningCount: 0 };

  try {
    const content = fs.readFileSync(AGENT_STDIO_LOG_PATH, "utf8");
    for (const rawLine of content.split("\n")) {
      const line = rawLine.trim();
      if (!line) {
        continue;
      }

      if (/^(?:\[WARN\]|npm warn\b)/i.test(line)) {
        metrics.warningCount += 1;
      }

      const jsonStart = line.indexOf("{");
      if (jsonStart < 0) {
        continue;
      }

      try {
        const parsed = JSON.parse(line.slice(jsonStart));
        if (!parsed || parsed.type !== "result") {
          continue;
        }

        if (typeof parsed.num_turns === "number" && parsed.num_turns >= 0) {
          metrics.turns = parsed.num_turns;
        }
        if (typeof parsed.total_cost_usd === "number" && Number.isFinite(parsed.total_cost_usd) && parsed.total_cost_usd >= 0) {
          metrics.estimatedCostUsd = parsed.total_cost_usd;
        }
      } catch {
        // Ignore non-JSON and truncated log lines.
      }
    }
  } catch {
    return metrics;
  }

  return metrics;
}

function formatPercent(value) {
  return `${Math.round(value * 100)}%`;
}

function deriveRuntimeStatus(raw) {
  if (raw.outputErrorCount > 0) {
    return "error";
  }
  if (raw.blockedRequests > 0) {
    return "degraded";
  }
  if (raw.warningCount > 0) {
    return "warning";
  }
  return "ok";
}

function deriveTokenIntensity(totalTokens) {
  if (typeof totalTokens !== "number" || !Number.isFinite(totalTokens) || totalTokens <= 0) {
    return "unknown";
  }
  if (totalTokens >= 1000000) {
    return "very_high";
  }
  if (totalTokens >= 250000) {
    return "high";
  }
  if (totalTokens >= 100000) {
    return "elevated";
  }
  return "normal";
}

function computeRuntimeRiskScore(raw) {
  let score = 0;
  score += Math.min(raw.outputErrorCount * 25, 50);
  score += Math.min(raw.warningCount * 5, 20);
  score += Math.min(raw.blockedRequests, 50) * 0.5;

  if (typeof raw.turnCount === "number") {
    if (raw.turnCount >= 20) {
      score += 10;
    } else if (raw.turnCount >= 10) {
      score += 5;
    }
  }

  return Math.round(Math.min(score, 100));
}

function computeOptimizationScore(raw, derived) {
  let score = derived.runtimeRiskScore;

  if (typeof raw.totalTokens === "number") {
    if (raw.totalTokens >= 1000000) {
      score += 25;
    } else if (raw.totalTokens >= 250000) {
      score += 15;
    } else if (raw.totalTokens >= 100000) {
      score += 10;
    }
  }

  if (typeof raw.cacheEfficiency === "number" && raw.cacheEfficiency < 0.5) {
    score += 10;
  }

  return Math.round(Math.min(score, 100));
}

function buildRuntimeObservabilityInsights(raw, derived) {
  const insights = [];

  if (typeof raw.totalTokens === "number" && raw.totalTokens >= 250000) {
    insights.push({
      category: "optimization",
      severity: raw.totalTokens >= 1000000 ? "high" : "medium",
      title: "Token-heavy run",
      summary: `The run consumed ${raw.totalTokens} total tokens.`,
    });
  }

  if (typeof raw.cacheEfficiency === "number") {
    insights.push({
      category: "cache",
      severity: raw.cacheEfficiency < 0.5 ? "medium" : "info",
      title: raw.cacheEfficiency < 0.5 ? "Low cache efficiency" : "Healthy cache efficiency",
      summary: `Cache efficiency was ${formatPercent(raw.cacheEfficiency)}.`,
    });
  }

  if (raw.turnCount >= 12) {
    insights.push({
      category: "execution",
      severity: "medium",
      title: "Exploratory execution path",
      summary: `The run required ${raw.turnCount} turns.`,
    });
  }

  if (derived.runtimeRiskScore > 0) {
    insights.push({
      category: "risk",
      severity: derived.runtimeRiskScore >= 50 ? "high" : "medium",
      title: "Runtime risk detected",
      summary: `The run scored ${derived.runtimeRiskScore} on the runtime risk scale.`,
    });
  }

  return insights;
}

function collectRuntimeObservabilityData(options = {}) {
  const awInfo = options.awInfo || readJSONIfExists(AW_INFO_PATH) || {};
  const agentOutput = options.agentOutput || readJSONIfExists(AGENT_OUTPUT_PATH) || { items: [], errors: [] };
  const agentUsage = options.agentUsage || readJSONIfExists(AGENT_USAGE_PATH) || {};
  const runtimeMetrics = options.runtimeMetrics || readAgentRuntimeMetrics();
  const items = Array.isArray(agentOutput.items) ? agentOutput.items : [];
  const errors = Array.isArray(agentOutput.errors) ? agentOutput.errors : [];
  const inputTokens = typeof agentUsage.input_tokens === "number" && agentUsage.input_tokens > 0 ? agentUsage.input_tokens : undefined;
  const outputTokens = typeof agentUsage.output_tokens === "number" && agentUsage.output_tokens > 0 ? agentUsage.output_tokens : undefined;
  const cacheReadTokens = typeof agentUsage.cache_read_tokens === "number" && agentUsage.cache_read_tokens > 0 ? agentUsage.cache_read_tokens : undefined;
  const cacheWriteTokens = typeof agentUsage.cache_write_tokens === "number" && agentUsage.cache_write_tokens > 0 ? agentUsage.cache_write_tokens : undefined;
  const effectiveTokens =
    typeof options.effectiveTokens === "number"
      ? options.effectiveTokens
      : typeof agentUsage.effective_tokens === "number" && agentUsage.effective_tokens > 0
        ? agentUsage.effective_tokens
        : typeof process.env.GH_AW_EFFECTIVE_TOKENS === "string" && process.env.GH_AW_EFFECTIVE_TOKENS.trim() !== ""
          ? parseInt(process.env.GH_AW_EFFECTIVE_TOKENS, 10)
          : undefined;
  const totalTokens = typeof effectiveTokens === "number" && !Number.isNaN(effectiveTokens) ? effectiveTokens : typeof inputTokens === "number" || typeof outputTokens === "number" ? (inputTokens || 0) + (outputTokens || 0) : undefined;
  const cacheEfficiencyBase = (inputTokens || 0) + (cacheReadTokens || 0);
  const cacheEfficiency = cacheEfficiencyBase > 0 && typeof cacheReadTokens === "number" ? cacheReadTokens / cacheEfficiencyBase : undefined;
  const warningCount = typeof options.warningCount === "number" ? options.warningCount : runtimeMetrics.warningCount;

  const raw = {
    firewallEnabled: awInfo.firewall_enabled === true,
    blockedRequests: countBlockedRequests(),
    createdItemCount: items.length,
    createdItemTypes: uniqueCreatedItemTypes(items),
    outputErrorCount: errors.length,
    warningCount,
    turnCount: typeof runtimeMetrics.turns === "number" ? runtimeMetrics.turns : 0,
    estimatedCostUsd: typeof runtimeMetrics.estimatedCostUsd === "number" ? runtimeMetrics.estimatedCostUsd : undefined,
    actionMinutes: typeof options.durationMs === "number" ? options.durationMs / 60000 : undefined,
    totalTokens,
    inputTokens,
    outputTokens,
    cacheReadTokens,
    cacheWriteTokens,
    cacheEfficiency,
  };

  const derived = {
    schemaVersion: 2,
    posture: raw.createdItemCount > 0 ? "write-capable" : "read-only",
    runtimeStatus: deriveRuntimeStatus(raw),
    tokenIntensity: deriveTokenIntensity(raw.totalTokens),
    runtimeRiskScore: computeRuntimeRiskScore(raw),
  };
  derived.optimizationScore = computeOptimizationScore(raw, derived);
  const insights = buildRuntimeObservabilityInsights(raw, derived);

  return {
    metadata: {
      workflowName: awInfo.workflow_name || "",
      engineId: awInfo.engine_id || "",
      traceId: process.env.GITHUB_AW_OTEL_TRACE_ID || (awInfo.context ? awInfo.context.otel_trace_id || "" : ""),
      staged: awInfo.staged === true,
    },
    raw,
    derived,
    insights,
  };
}

function buildObservabilitySummary(data) {
  const lines = [];

  lines.push("<details>");
  lines.push("<summary>Observability</summary>");
  lines.push("");

  if (data.metadata.workflowName) {
    lines.push(`- **workflow**: ${data.metadata.workflowName}`);
  }
  if (data.metadata.engineId) {
    lines.push(`- **engine**: ${data.metadata.engineId}`);
  }
  if (data.metadata.traceId) {
    lines.push(`- **trace id**: ${data.metadata.traceId}`);
  }

  lines.push(`- **posture**: ${data.derived.posture}`);
  lines.push(`- **runtime status**: ${data.derived.runtimeStatus}`);
  if (typeof data.raw.totalTokens === "number") {
    lines.push(`- **total tokens**: ${data.raw.totalTokens}`);
  }
  if (typeof data.raw.estimatedCostUsd === "number") {
    lines.push(`- **estimated cost usd**: ${data.raw.estimatedCostUsd}`);
  }
  if (typeof data.raw.turnCount === "number") {
    lines.push(`- **turns**: ${data.raw.turnCount}`);
  }
  if (typeof data.raw.cacheEfficiency === "number") {
    lines.push(`- **cache efficiency**: ${formatPercent(data.raw.cacheEfficiency)}`);
  }
  lines.push(`- **runtime risk score**: ${data.derived.runtimeRiskScore}`);
  lines.push(`- **optimization score**: ${data.derived.optimizationScore}`);
  lines.push(`- **created items**: ${data.raw.createdItemCount}`);
  lines.push(`- **blocked requests**: ${data.raw.blockedRequests}`);
  lines.push(`- **agent output errors**: ${data.raw.outputErrorCount}`);
  lines.push(`- **firewall enabled**: ${data.raw.firewallEnabled}`);
  lines.push(`- **staged**: ${data.metadata.staged}`);

  if (data.raw.createdItemTypes.length > 0) {
    lines.push("- **item types**:");
    for (const itemType of data.raw.createdItemTypes) {
      lines.push(`  - ${itemType}`);
    }
  }

  lines.push("");
  lines.push("</details>");

  return lines.join("\n") + "\n";
}

module.exports = {
  buildObservabilitySummary,
  collectRuntimeObservabilityData,
  readAgentRuntimeMetrics,
  readJSONIfExists,
};
