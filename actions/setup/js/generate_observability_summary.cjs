// @ts-check
/// <reference types="@actions/github-script" />

const { main: exportCopilotOtelTraces } = require("./export_copilot_otel_traces.cjs");
const { buildObservabilitySummary, collectRuntimeObservabilityData } = require("./runtime_observability.cjs");

async function main(core) {
  await exportCopilotOtelTraces(core);
  const data = collectRuntimeObservabilityData();
  const markdown = buildObservabilitySummary(data);
  await core.summary.addRaw(markdown).write();
  core.info("Generated observability summary in step summary");
}

module.exports = {
  buildObservabilitySummary,
  collectObservabilityData: collectRuntimeObservabilityData,
  main,
};
