// @ts-check

/**
 * Detect agent engine errors in the agent stdio log.
 *
 * Scans the agent stdio log for known error patterns and sets GitHub Actions
 * output variables for each detected error class:
 *
 *   - inference_access_error: The COPILOT_GITHUB_TOKEN does not have valid
 *     access to inference (e.g., "Access denied by policy settings").
 *   - mcp_policy_error: MCP servers were blocked by enterprise/organization
 *     policy (e.g., "MCP servers were blocked by policy: 'github', 'safeoutputs'").
 *   - agentic_engine_timeout: The agentic engine process was killed by a
 *     signal (SIGTERM/SIGKILL/SIGINT), typically due to the step
 *     timeout-minutes limit being reached.
 *   - model_not_supported_error: The configured model is invalid or unsupported
 *     for the selected engine/account (for example unknown model name, model not
 *     found, or model unavailable for the plan).
 *
 * This replaces the individual bash scripts (detect_inference_access_error.sh,
 * detect_mcp_policy_error.sh) with a single JavaScript step.
 *
 * Exit codes:
 *   0 — Always succeeds (uses continue-on-error in the workflow step)
 */

"use strict";

const fs = require("fs");
const { resolveAICreditsFailureState } = require("./ai_credits_context.cjs");
const { logSpan } = require("./otlp.cjs");

const LOG_FILE = "/tmp/gh-aw/agent-stdio.log";

// Pattern: Copilot CLI inference access denied
const INFERENCE_ACCESS_ERROR_PATTERN = /Access denied by policy settings|invalid access to inference/;

// Pattern: MCP servers blocked by enterprise/organization policy
const MCP_POLICY_BLOCKED_PATTERN = /MCP servers were blocked by policy:/;

// Pattern: Agentic engine process killed by signal (timeout).
// When GitHub Actions cancels a step due to timeout-minutes, the runner sends
// SIGINT/SIGTERM/SIGKILL to the process group.  The copilot_harness.cjs (and
// other engine wrappers) log the signal in their close handlers:
//   [copilot-harness] attempt 1: process closed exitCode=1 signal=SIGTERM ...
// The pattern matches any "signal=SIG(TERM|KILL|INT)" occurrence in the log,
// making it engine-agnostic.
const AGENTIC_ENGINE_TIMEOUT_PATTERN = /signal=SIG(?:TERM|KILL|INT)/;

// Pattern: Configured model is invalid or unavailable.
// Covers common engine/provider variants:
//   - "The requested model is not supported"
//   - "invalid model name '...'"
//   - "unknown model <id>"
//   - "model ... not found"
//   - "model ... does not exist"
const MODEL_NOT_SUPPORTED_PATTERN =
  /(?:The requested model is not supported|invalid model(?:\s+name)?\s+['"`]?[a-z0-9._:/@-]+['"`]?(?=(?:\s*$|\s*[\n\r.,;:!?)]))|unknown model\s+['"`]?[a-z0-9._:/@-]+['"`]?(?=(?:\s*$|\s*[\n\r.,;:!?)]))|model(?:\s+name)?\s+['"`]?[a-z0-9._:/@-]+['"`]?\s+(?:is\s+)?(?:not found|does not exist|not supported|not available|unavailable))/i;

/**
 * Detect known error patterns in a log string and return detection results.
 * @param {string} logContent - Contents of the agent stdio log
 * @returns {{ inferenceAccessError: boolean, mcpPolicyError: boolean, agenticEngineTimeout: boolean, modelNotSupportedError: boolean, aiCredits: string }}
 */
function detectErrors(logContent) {
  const { aiCredits } = resolveAICreditsFailureState();
  return {
    inferenceAccessError: INFERENCE_ACCESS_ERROR_PATTERN.test(logContent),
    mcpPolicyError: MCP_POLICY_BLOCKED_PATTERN.test(logContent),
    agenticEngineTimeout: AGENTIC_ENGINE_TIMEOUT_PATTERN.test(logContent),
    modelNotSupportedError: MODEL_NOT_SUPPORTED_PATTERN.test(logContent),
    aiCredits,
  };
}

/**
 * Write GitHub Actions outputs to $GITHUB_OUTPUT.
 * @param {{ inferenceAccessError: boolean, mcpPolicyError: boolean, agenticEngineTimeout: boolean, modelNotSupportedError: boolean, aiCredits: string }} results
 */
function writeOutputs(results) {
  const outputFile = process.env.GITHUB_OUTPUT;
  if (!outputFile) {
    process.stderr.write("[detect-agent-errors] GITHUB_OUTPUT not set — skipping output\n");
    return;
  }

  const lines = [
    `inference_access_error=${results.inferenceAccessError}`,
    `mcp_policy_error=${results.mcpPolicyError}`,
    `agentic_engine_timeout=${results.agenticEngineTimeout}`,
    `model_not_supported_error=${results.modelNotSupportedError}`,
    `ai_credits=${results.aiCredits}`,
  ];
  fs.appendFileSync(outputFile, lines.join("\n") + "\n");
}

/**
 * @param {{ inferenceAccessError: boolean, mcpPolicyError: boolean, agenticEngineTimeout: boolean, modelNotSupportedError: boolean, aiCredits: string }} results
 * @returns {Record<string, string | number | boolean>}
 */
function buildOTLPAttributes(results) {
  const attributes = {
    "detect-agent-errors.inference_access_error": results.inferenceAccessError,
    "detect-agent-errors.mcp_policy_error": results.mcpPolicyError,
    "detect-agent-errors.agentic_engine_timeout": results.agenticEngineTimeout,
    "detect-agent-errors.model_not_supported_error": results.modelNotSupportedError,
  };
  const aiCreditsNumber = Number.parseFloat(results.aiCredits);
  if (Number.isFinite(aiCreditsNumber) && aiCreditsNumber >= 0) {
    return { ...attributes, "gh-aw.aic": aiCreditsNumber };
  }
  return attributes;
}

async function main() {
  let logContent = "";

  if (fs.existsSync(LOG_FILE)) {
    logContent = fs.readFileSync(LOG_FILE, "utf8");
  } else {
    process.stderr.write(`[detect-agent-errors] Log file not found: ${LOG_FILE}\n`);
  }

  const results = detectErrors(logContent);

  if (results.inferenceAccessError) {
    process.stderr.write("[detect-agent-errors] Detected inference access error in agent log\n");
  }
  if (results.mcpPolicyError) {
    process.stderr.write("[detect-agent-errors] Detected MCP policy error in agent log\n");
  }
  if (results.agenticEngineTimeout) {
    process.stderr.write("[detect-agent-errors] Detected timeout: engine process was killed by signal (step timeout-minutes likely exceeded)\n");
  }
  if (results.modelNotSupportedError) {
    process.stderr.write("[detect-agent-errors] Detected model configuration error: configured model is invalid or unavailable for this engine/account\n");
  }
  process.stderr.write(`[detect-agent-errors] AI Credits: ${results.aiCredits || "unavailable"}\n`);

  writeOutputs(results);
  await logSpan("detect-agent-errors", buildOTLPAttributes(results), {
    isError: results.inferenceAccessError || results.mcpPolicyError || results.agenticEngineTimeout || results.modelNotSupportedError,
  });
}

if (require.main === module) {
  main().catch(error => {
    process.stderr.write(`[detect-agent-errors] Execution warning: ${error instanceof Error ? error.message : String(error)}\n`);
  });
}

module.exports = { detectErrors, buildOTLPAttributes, INFERENCE_ACCESS_ERROR_PATTERN, MCP_POLICY_BLOCKED_PATTERN, AGENTIC_ENGINE_TIMEOUT_PATTERN, MODEL_NOT_SUPPORTED_PATTERN };
