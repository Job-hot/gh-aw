// @ts-check
/// <reference types="@actions/github-script" />

/**
 * conclusion_runner.cjs
 *
 * Orchestrates all conclusion job handlers in a single github-script step.
 * Previously these were separate github-script steps:
 *   - Process no-op messages    (handle_noop_message.cjs)
 *   - Log detection run         (handle_detection_runs.cjs)
 *   - Record missing tool       (missing_tool.cjs)
 *   - Record incomplete         (report_incomplete_handler.cjs)
 *   - Handle agent failure      (handle_agent_failure.cjs)
 *   - Update reaction comment   (notify_comment_error.cjs, when status-comment: true)
 *
 * Merging them into one step:
 *   - calls setupGlobals() once, sharing execution context across all handlers
 *   - eliminates redundant github-script action startup overhead per step
 *   - allows handlers to share the same authenticated github client
 *
 * Each handler is guarded by an environment variable flag set by the compiler
 * when the corresponding feature is enabled in workflow frontmatter.
 */
async function main() {
  // Process no-op messages (enabled when safe-outputs.noop is configured)
  if (process.env.GH_AW_NOOP_ENABLED === "true") {
    const { main: handleNoop } = require("./handle_noop_message.cjs");
    await handleNoop();
  }

  // Log detection run (enabled when threat detection is configured)
  if (process.env.GH_AW_DETECTION_ENABLED === "true") {
    const { main: handleDetectionRuns } = require("./handle_detection_runs.cjs");
    await handleDetectionRuns();
  }

  // Record missing tool (enabled when safe-outputs.missing-tool is configured)
  if (process.env.GH_AW_MISSING_TOOL_ENABLED === "true") {
    const { main: handleMissingTool } = require("./missing_tool.cjs");
    await handleMissingTool();
  }

  // Record incomplete (enabled when safe-outputs.report-incomplete is configured)
  if (process.env.GH_AW_REPORT_INCOMPLETE_ENABLED === "true") {
    const { main: handleReportIncomplete } = require("./report_incomplete_handler.cjs");
    await handleReportIncomplete();
  }

  // Handle agent failure — always runs to surface failures and update issue tracking
  const { main: handleAgentFailure } = require("./handle_agent_failure.cjs");
  await handleAgentFailure();

  // Update reaction comment with completion status (enabled when status-comment: true)
  if (process.env.GH_AW_STATUS_COMMENT_ENABLED === "true") {
    const { main: notifyCommentError } = require("./notify_comment_error.cjs");
    await notifyCommentError();
  }
}

module.exports = { main };
