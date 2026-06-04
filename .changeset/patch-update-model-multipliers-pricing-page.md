---
"gh-aw": patch
---

Update model multipliers to use https://docs.github.com/en/copilot/reference/copilot-billing/models-and-pricing as the source of truth. Multipliers for all models listed on that page are now derived from the input token price at a 1:1 rate ($1 USD per 1M input tokens = 1x multiplier). Affected models: GPT-5 mini, GPT-5.2, GPT-5.2-Codex, GPT-5.3-Codex, GPT-5.4, GPT-5.4 mini, GPT-5.4 nano, GPT-5.5, Claude Haiku 4.5, Claude Sonnet 4/4.5/4.6, Claude Opus 4.5/4.6/4.7/4.8, Gemini 2.5 Pro, Gemini 3 Flash, Gemini 3.1 Pro, Gemini 3.5 Flash, Raptor mini, and MAI-Code-1-Flash.
