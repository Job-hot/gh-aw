import fs from "fs";
import { afterEach, beforeEach, describe, expect, it, vi } from "vitest";

describe("ai_credits_context", () => {
  let resolveAICreditsFailureState;
  beforeEach(async () => {
    vi.resetModules();
    ({ resolveAICreditsFailureState } = await import("./ai_credits_context.cjs"));
  });

  afterEach(() => {
    vi.restoreAllMocks();
    delete process.env.GH_AW_AIC;
    delete process.env.GH_AW_MAX_AI_CREDITS;
    delete process.env.GH_AW_AI_CREDITS_RATE_LIMIT_ERROR;
    delete process.env.GH_AW_AGENT_OUTPUT;
    fs.rmSync("/tmp/gh-aw/agent_usage.json", { force: true });
    fs.rmSync("/tmp/gh-aw/awf-config.json", { force: true });
  });

  it("falls back to downloaded agent usage and awf config files when env data is absent", () => {
    fs.mkdirSync("/tmp/gh-aw", { recursive: true });
    fs.writeFileSync("/tmp/gh-aw/agent_usage.json", JSON.stringify({ ai_credits: 395.044 }));
    fs.writeFileSync("/tmp/gh-aw/awf-config.json", JSON.stringify({ apiProxy: { maxAiCredits: 1000 } }));

    process.env.GH_AW_AI_CREDITS_RATE_LIMIT_ERROR = "true";

    expect(resolveAICreditsFailureState()).toEqual({
      aiCredits: "395.044",
      maxAICredits: "1000",
      aiCreditsRateLimitError: false,
    });
  });
});
