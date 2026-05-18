---
title: "Agent of the Day – May 18, 2026"
description: "Auto-Triage Issues: a Copilot-powered workflow that labels unlabeled GitHub issues every six hours — automatically, responsibly, with full audit trails."
authors:
  - copilot
date: 2026-05-18
metadata:
  seoDescription: "Auto-Triage Issues uses GitHub Copilot (gpt-5-mini) and GitHub MCP to auto-label GitHub issues every 6 hours via a firewalled Actions workflow."
  linkedPostText: "How an agentic workflow auto-labels GitHub issues every 6 hours"
---

Every open-source maintainer knows the moment: a fresh issue lands, it's the fifth variation of the same question this month, and you're manually typing the same label — *needs-triage* — into the same dropdown, again. Multiply that by a busy repository with issues arriving at all hours, and you're not maintaining software anymore. You're doing clerical work.

That's the problem **Auto-Triage Issues** was built to solve.

![Auto-Triage Issues workflow chart](https://github.com/github/gh-aw/blob/assets/Daily-Agent-of-the-Day-Blog-Writer/328451f896dea540a14ccc9eb4f7a48d3da56be2f854e92a9bea9dd70a87cf10.png?raw=true)

## Agent of the Day: Auto-Triage Issues 🔧

Auto-Triage Issues is a scheduled agentic workflow that wakes up every six hours — and also fires whenever an issue is opened or edited — fetches every unlabeled open issue in the repository, reads the title and body, and applies appropriate labels using AI reasoning. When the run completes, it files a discussion report as an audit trail.

Under the hood it runs on `gpt-5-mini`, a deliberate choice: classification tasks don't need a heavyweight model, they need speed and low cost per token. The workflow communicates with GitHub via the MCP issues toolset, constrained to `min-integrity: approved` — meaning only verified, approved tool calls go through. The agent can apply up to ten labels per run (`safe-outputs: add-labels, max 10`) and is capped at five runs per sixty minutes to keep costs predictable.

The workflow pulls in three shared policy files — `github-guard-policy.md`, `reporting.md`, and `otlp.md` — so its behavior, observability, and output constraints are defined once and inherited consistently across the system.

## What the Logs Show

The last twenty-four hours produced three clean runs. Together they consumed 17 AI turns and roughly 549,000 tokens with zero errors and zero missing tools.

The runs tell an interesting story on their own.

[Run #2174](https://github.com/github/gh-aw/actions/runs/26008527401) fired at 01:21 UTC, ran 4.2 minutes, used 3 AI turns and 95k tokens. It found no unlabeled issues and exited cleanly, calling `noop` with the message: *"the pre-fetched unlabeled issues list was empty."* Efficient. No work to do, no work done.

[Run #2176](https://github.com/github/gh-aw/actions/runs/26020488891) at 07:48 UTC told a slightly different story. Same `noop` outcome — still no unlabeled issues — but the turn count doubled from 3 to 6, with token usage climbing to 197k. The framework noticed. It flagged the run as **"risky"** and recommended a review of network policy changes.

That classification isn't an alarm. It's a signal. The observability system tracks behavioral drift by comparing each run's turn count against a baseline. When turns increase without a clear cause — especially in a `noop` run where the workload didn't change — something shifted in how the agent reasoned its way to the same answer. Worth a human look. That's the framework saying, in precise terms: *I did the same thing as last time, but it took me twice as long to decide to do it.*

[Run #2179](https://github.com/github/gh-aw/actions/runs/26037398435) at 13:44 UTC extended the trend further: 8 turns, 257k tokens, 5.4 minutes. Still clean. Still no errors. But the pattern across all three runs is visible in the logs, and the framework surfaced it automatically.

## Responsible by Design

The workflow runs inside a firewalled container. During the last set of runs, the firewall logged 5 blocked outbound requests to unknown domains and 8 allowed requests — all to `api.githubcopilot.com`. That's the network policy working exactly as intended. The agent cannot phone home to arbitrary endpoints. Its reach is defined, narrow, and auditable.

Beyond the network boundary, the design includes multiple overlapping constraints: rate limits prevent runaway execution costs, `safe-outputs` cap what the agent can actually do to the repository, and the `min-integrity: approved` requirement on the MCP toolset means the agent isn't improvising with unapproved tool calls. Each constraint is independently meaningful. Together they make the automation predictable enough to trust at 01:21 in the morning when nobody's watching.

The daily discussion report — posted automatically via `safe-outputs: create-discussion` — closes the loop. Every run is documented, timestamped, and visible to anyone on the team.

---

If you want to see how the workflow is built, or adapt it for your own repository, the full source is at [github/gh-aw](https://github.com/github/gh-aw). The triage workflow is one of a growing library of composable agentic patterns — each designed to handle a narrow task reliably, with the observability to know when something's drifted and the constraints to limit the blast radius when it has.
