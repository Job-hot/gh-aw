---
title: "Weekly Update – June 8, 2026"
description: "This week gh-aw completes a full migration from Effective Tokens to AI Credits, fixes a duplicate success marker in gh aw fix, and stabilizes wasm golden normalization."
authors:
  - copilot
date: 2026-06-08
metadata:
  seoDescription: "gh-aw migrates from Effective Tokens to AI Credits across CLI, OTel, docs, and safe-outputs — plus bug fixes and test improvements."
---

Another week of high-velocity changes in [github/gh-aw](https://github.com/github/gh-aw)! The headline this week is a sweeping migration from **Effective Tokens** to **AI Credits** — a rename that touched the CLI, OpenTelemetry instrumentation, error codes, safe-outputs schema, and documentation all at once. Plus a couple of welcome bug fixes and test stabilization work.

## Effective Tokens → AI Credits

The biggest story this week is a coordinated, multi-PR rename that retired the "Effective Tokens" terminology in favor of **AI Credits** across the entire stack. This change aligns the project with how Copilot and the broader GitHub AI platform surface cost and usage to users.

Here's what landed:

- **[#37692](https://github.com/github/gh-aw/pull/37692) – Migrate docs terminology from Effective Tokens to AI Credits**: Updated the documentation to use AI Credits (AIC) throughout, replacing "effective token" references in user-facing content.
- **[#37693](https://github.com/github/gh-aw/pull/37693) – OpenTelemetry JS: enforce AIC emission and block effective token attribute**: The OTel instrumentation layer now emits the `ai_credits` attribute and rejects the legacy `effective_tokens` attribute, making the rename enforceable at the instrumentation layer.
- **[#37691](https://github.com/github/gh-aw/pull/37691) – Replace `effective_tokens_rate_limit_error` with `ai_credits_rate_limit_error`**: Error codes in workflow outputs are now aligned to AI Credits, so tooling and monitoring rules can be updated consistently.
- **[#37690](https://github.com/github/gh-aw/pull/37690) – Replace EffectiveTokens with AI Credits in CLI audit/logs outputs**: The `gh aw audit` and `gh aw logs` commands now display AI Credits instead of Effective Tokens.
- **[#37673](https://github.com/github/gh-aw/pull/37673) – Remove JavaScript-side effective-token handling in favor of AI credits**: Cleaned up the JS runtime to use AI Credits natively.
- **[#37686](https://github.com/github/gh-aw/pull/37686) – Remove legacy effective-token limits and ET surfacing; keep migration codemods**: Removed the old ET limit enforcement while retaining the `gh aw fix` codemods that help workflows migrate automatically.
- **[#37685](https://github.com/github/gh-aw/pull/37685) – Add fix codemod to migrate safe-outputs ET suffix placeholders to AI Credits**: The `gh aw fix` command now includes a codemod that automatically rewrites `_et_` suffix placeholders in safe-outputs message templates to their `_aic_` equivalents.

If you maintain workflows that reference `effective_tokens` in error handling, rate limit responses, or OTel attributes, run `gh aw fix --write` to apply the migration automatically. A companion blog post — [Migrating from Effective Tokens to AI Credits](/blog/2026-06-08-migrating-from-effective-tokens-to-ai-credits) — walks through the full transition.

## Bug Fixes

- **[#37688](https://github.com/github/gh-aw/pull/37688) – Fix duplicate success marker in `gh aw fix --write` output**: `gh aw fix --write` was printing the success summary twice when changes were applied. It now prints it exactly once — concise and correct.
- **[#37684](https://github.com/github/gh-aw/pull/37684) – Add missing `RuntimeConnection` import in `copilot_sdk_driver_sample_node.cjs`**: A missing import in the sample Copilot SDK driver was causing runtime errors in Node.js environments. Fixed.

## Test & CI Improvements

- **[#37711](https://github.com/github/gh-aw/pull/37711) – Stabilize pinning tests and unify wasm golden normalization**: Addressed three sources of flakiness: a brittle action-pin expectation that hardcoded a SHA, inconsistent wasm-vs-golden normalization between Go and Node runners, and stale wasm golden outputs. CI should be substantially quieter now.

## Documentation

- **[#37671](https://github.com/github/gh-aw/pull/37671)** – Removed stale "25M default budget" language and aligned docs to current guardrail behavior.
- **[#37670](https://github.com/github/gh-aw/pull/37670)** – Surfaced guardrail `direction` and numeric `threshold` fields in experiment documentation.
- **[#37679](https://github.com/github/gh-aw/pull/37679)** – Self-healing documentation fixes from automated issue analysis.

---

## 🤖 Agent of the Week: Necromancer

The undertaker of unprotected code paths — reads every merge-ready PR, traces root-cause issues, and adds regression tests before the branch disappears into `main`.

This month, Necromancer has been busy stalking pull requests across the repo. On PR [#31161](https://github.com/github/gh-aw/pull/31161), it dug through `missing_messages_helper.cjs`, found a gap, and pushed commit `a50a723` with a regression test asserting the `// @ts-check` directive stays firmly at line one — the kind of precise, easy-to-miss invariant that only breaks once and then causes 45 minutes of confusion. It also tightened the all-fields test to assert tools, data, noop, and incomplete sections together. On PR [#31194](https://github.com/github/gh-aw/pull/31194), it modelled the retry-decision matrix for the codex harness — adding explicit tests for `exitCode=143` (SIGTERM) being correctly rejected by `shouldRetryWithContinue()`.

In one memorable run, Necromancer attempted to warm its "featured plugin IDs cache" by calling `chatgpt.com/backend-api/plugins/featured` — and was immediately greeted by a Cloudflare challenge asking it to enable JavaScript and cookies. A robot, politely asked by the internet to prove it's not a robot.

💡 **Usage tip**: Trigger Necromancer with the `/necromancer` label command on any PR that fixes a tricky bug — it'll trace the root cause, find coverage gaps, and add tests that fail without the fix.

→ [View the workflow on GitHub](https://github.com/github/gh-aw/blob/main/.github/workflows/necromancer.md)

---

Grab the latest changes from [github/gh-aw](https://github.com/github/gh-aw) and run `gh aw fix --write` to get your workflows migrated to AI Credits. As always, contributions and feedback are welcome.
