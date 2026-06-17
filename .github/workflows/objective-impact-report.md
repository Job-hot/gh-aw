---
emoji: 📊
description: Executive impact efficiency report from workflow outcomes tied to tracked objectives.
on:
  workflow_dispatch:
permissions:
  contents: read
  pull-requests: read
  actions: read
  issues: read
cache:
  - key: objective-impact-report-cache-${{ github.run_id }}
    name: Save objective impact report dataset cache
    path: /tmp/gh-aw/agent/objective-impact-report
    restore-keys: |
      objective-impact-report-cache-
tools:
  cli-proxy: true
  github:
    mode: gh-proxy
    read-only: true
    toolsets: [default]
  bash:
    - "cat /tmp/gh-aw/agent/objective-impact-report/*.json"
    - "cat /tmp/gh-aw/agent/objective-impact-report/*.jsonl"
    - "jq *"
    - "ls /tmp/gh-aw/agent/objective-impact-report"
    - "head -n * /tmp/gh-aw/agent/objective-impact-report/*.json /tmp/gh-aw/agent/objective-impact-report/*.jsonl"
pre-agent-steps:
  - name: Prepare deterministic impact datasets
    env:
      GH_TOKEN: ${{ secrets.GITHUB_TOKEN }}
      GITHUB_TOKEN: ${{ secrets.GITHUB_TOKEN }}
      EXPR_GITHUB_REPOSITORY: ${{ github.repository }}
    run: bash scripts/prepare-objective-impact-report-dataset.sh
safe-outputs:
  close-issue:
    required-title-prefix: "Impact Efficiency Report - "
    target: "*"
    max: 1
  create-issue:
    title-prefix: "Impact Efficiency Report - "
    max: 1
---

# Impact Efficiency Report

## Required Inputs (already precomputed)

Use these deterministic files first:

- /tmp/gh-aw/agent/objective-impact-report/dataset-manifest.json
- /tmp/gh-aw/agent/objective-impact-report/run-context.json
- /tmp/gh-aw/agent/objective-impact-report/objective-mapping.json
- /tmp/gh-aw/agent/objective-impact-report/workflow-logs.json
- /tmp/gh-aw/agent/objective-impact-report/aic-by-workflow.json
- /tmp/gh-aw/agent/objective-impact-report/merged-prs-linked.json
- /tmp/gh-aw/agent/objective-impact-report/closed-unmerged-prs-linked.json
- /tmp/gh-aw/agent/objective-impact-report/safe-output-issue-evaluations.jsonl
- /tmp/gh-aw/agent/objective-impact-report/safe-output-issue-summary.json

`safe-output-issue-evaluations.jsonl` provides deterministic outcome status and workflow attribution, but it does **not** include precomputed objective attribution fields. If optional `*-with-objective.json` files exist and are non-empty, use them; otherwise compute objective mapping from linked/root issue labels.

Do **not** re-fetch these datasets with GitHub tools unless a required file is missing, empty, or fails JSON parsing.

## Goal

Produce a comprehensive executive report on what work was performed, what AIC tokens were spent on, which outcomes delivered the highest impact, and which workflows contributed that impact. The report must clearly answer: *What did we build, fix, and ship, what was the most impactful work in the repository, which workflows drove that impact, and was it worth the cost?*

Focus only on **pull request outcomes** and **safe output outcomes** (issues created or closed via the safe-output mechanism). Other outcome types are excluded because their acceptance criteria are not yet well-defined and most remain pending.

Use this model:

```text
Outcome = a PR or safe-output issue produced by a GitHub Agentic Workflow run
Objective Value = numeric value from the repository objective-mapping configuration applied to traced root labels
Outcome Indicator = 1 for accepted/delivered outcomes, 0 otherwise
Outcome Value = Outcome Indicator × Objective Value
Impact Efficiency = Σ Outcome Value / AI Credits
```

Treat AI Credits as total model-credit cost aggregated per workflow across the full analysis window, not just the subset of runs that produced the analyzed outcomes.
Start with `/tmp/gh-aw/agent/objective-impact-report/aic-by-workflow.json` as the primary AIC source, and `/tmp/gh-aw/agent/objective-impact-report/workflow-logs.json` and `/tmp/gh-aw/agent/objective-impact-report/dataset-manifest.json` as additional context for run details and source provenance.
When available, use deterministic precomputed run data that already includes each run's `aic` field.
Prefer existing gh-aw outputs that already surface `aic`, such as pre-downloaded `gh aw logs --json` data or audit/log artifacts derived from the same run summaries.
Only fall back to MCP or other live retrieval if deterministic precomputed AIC inputs are unavailable or the manifest says the fallback is still required.
Use the same time window for AIC as for outcomes.

Perform direct workflow attribution for every analyzed outcome.
Outcomes deliver value.
Objectives provide context and importance.
Workflows explain where delivered value came from.
AI Credits provide cost.
Do not use an LLM judge.

## AIC Source of Truth

Resolve AI Credits in this order:

1. **Primary: `/tmp/gh-aw/agent/objective-impact-report/aic-by-workflow.json`** — aggregated per-workflow AIC from daily token-audit memory snapshots covering the analysis window. Each entry has `workflow_name`, `total_aic`, and `run_count`. Use this as the denominator for overall and per-workflow Impact Efficiency. Check `dataset-manifest.json` for `aic_by_workflow_source` and `aic_by_workflow_snapshot_count` to understand coverage. The `aic-by-workflow.json` data is pre-aggregated across all available daily snapshots within the window and is the most reliable AIC source.
2. Deterministic precomputed `/tmp/gh-aw/agent/objective-impact-report/workflow-logs.json` data with per-run `aic` (use only when `aic-by-workflow.json` is unavailable or has `source: "none"`)
3. MCP or other live retrieval only as a documented fallback

When computing total AI Credits for the report:
- Sum `total_aic` across all entries in `aic-by-workflow.json` for the repository-wide total AIC
- For per-workflow AIC, look up the workflow by name in `aic-by-workflow.json`
- If a workflow has no entry in `aic-by-workflow.json`, treat its AIC as unknown (not zero) and add a note in the Data Quality section of the report listing which workflows had no AIC data available.

If a run's `aic` field is missing or null, treat it as `0` and count it as missing-cost data in the report.

## Scope

Analyze only the following outcome types from the last 180 days:

- **Pull request outcomes**: PRs created by GitHub Agentic Workflow runs **that have a linked closing issue** (`Closes #N`). Accepted = merged. Rejected = closed without merge. Skip open (pending) PRs. **Exclude entirely** any PR without a traceable linked issue — do not fall back to PR labels.
- **Safe output outcomes**: issues created or closed by workflow runs via the safe-output mechanism. Accepted = issue successfully created/closed. Skip any with unresolved state.

Exclude all other outcome types (direct issue outcomes, comments, discussions, etc.). These are omitted because their acceptance criteria are incomplete and most are left pending, which would distort the metric.

## Objective value mapping

Objective values should be resolved from deterministic inputs whenever available.

For pull request outcomes, use:

- `/tmp/gh-aw/agent/objective-impact-report/merged-prs-linked.json`
- `/tmp/gh-aw/agent/objective-impact-report/closed-unmerged-prs-linked.json`

If optional `/tmp/gh-aw/agent/objective-impact-report/merged-prs-with-objective.json` and `/tmp/gh-aw/agent/objective-impact-report/closed-unmerged-prs-with-objective.json` exist and are non-empty, use their precomputed `objective_value`, `objective_labels`, `root_issue_numbers`, `root_issue_labels`, and `attribution_source` fields directly.

For safe-output issue outcomes, use outcome identity and status from:

- `/tmp/gh-aw/agent/objective-impact-report/safe-output-issue-evaluations.jsonl`

The mapping uses the `outcome/` root-level resolver (mirrors `intent.Resolver.ResolvePullRequest` in Go):
- For PRs with exactly one closing issue (`attribution_source: "closing_issue"`): objective computed from the closing issue's labels.
- For PRs with no closing issue (`attribution_source: "artifact_labels"` or `"none"`): **exclude from analysis** — do not fall back to PR labels.
- For safe-output issues (`attribution_source: "issue_labels"`): objective computed from the issue's own labels fetched from the issue record.

If `objective_value` is `0` and the entry has root/issue labels present, mark the outcome as `unmapped` (no matching label in the mapping). If there are no root labels at all, mark it as excluded.

Do not invent fallback scoring rules such as milestone bonuses, project bonuses, or priority-to-points heuristics.

## Outcome association rules

For each in-scope outcome, use the precomputed root-tracing results:

1. For pull-request outcomes, start from `linked_issue_numbers` in `merged-prs-linked.json` / `closed-unmerged-prs-linked.json`.
2. If an optional `*-with-objective.json` dataset exists and is non-empty, use its `attribution_source`, `root_issue_numbers`, `root_issue_labels`, `objective_value`, and `objective_labels` fields directly.
3. Otherwise, if `linked_issue_numbers` is empty, **exclude the PR from analysis entirely**. Do not fall back to PR labels. Count it in the "PRs excluded (no linked issue)" total.
4. Otherwise, resolve root issue labels from the linked issues (deterministic cache first; live issue lookup only when labels are missing), then compute objective mapping from those root labels.
5. For safe-output issue outcomes, use `/tmp/gh-aw/agent/objective-impact-report/safe-output-issue-evaluations.jsonl` as the primary source for outcome state, workflow attribution, and issue identity; fetch issue labels to compute objective mapping.
6. Record traced issue numbers (`root_issue_numbers` when present, otherwise `linked_issue_numbers`) in the report as the audit trail.
7. If `objective_value` is `0` and labels are present, mark the outcome as `unmapped`, exclude it from `Σ Outcome Value`, and report it separately.

## Computation

For each in-scope outcome:

```text
Outcome Indicator:
  PR outcome:           1 if merged, 0 if closed without merge (open PRs excluded)
  Safe output outcome:  1 if successfully created/closed, 0 otherwise
Outcome Value = Outcome Indicator × Objective Value
```

Then compute:

```text
Accepted Outcome Count = count(outcomes where Outcome Indicator = 1)
Total Outcome Value    = sum(Outcome Value)
AI Credits             = sum(run.aic across analyzed runs)
Impact Efficiency      = Total Outcome Value / AI Credits  (value units per AI Credit; undefined when AI Credits = 0)
```

Also compute per-workflow attribution using the workflow run that directly produced each analyzed outcome:

```text
Workflow Contributed Value = sum(Outcome Value for outcomes produced by that workflow)
Workflow Accepted Outcomes = count(accepted, mapped outcomes produced by that workflow)
Workflow AI Credits        = sum(run.aic for analyzed runs of that workflow)
Workflow Value Share       = Workflow Contributed Value / Total Outcome Value
Workflow Impact Efficiency = Workflow Contributed Value / Workflow AI Credits
```

Use the workflow name from the producing run as the attribution key. If multiple runs from the same workflow produced analyzed outcomes, aggregate them together. If the workflow name cannot be resolved for an analyzed outcome, report it as unattributed and exclude it from workflow ranking totals while still counting the outcome in overall delivered-value totals.

If AI Credits is missing or zero, report that Impact Efficiency is not computable and explain whether credits data was unavailable or no credits were consumed in the analysis window.
If only some runs are missing `aic`, still compute the metric from the available values and explicitly report how many runs had missing cost data.

## Report

Before creating the new report, search for an existing open issue titled:

```text
Impact Efficiency Report - YYYY-MM-DD
```

If one already exists for today:

1. Close that issue first with a brief comment explaining that it is being replaced by a freshly generated report for the same day.
2. Then create the new report issue.

Create one issue titled:

```text
Impact Efficiency Report - YYYY-MM-DD
```

Use **progressive disclosure** for the report body. Place the Executive Summary first as plain text (no collapsible wrapper). Wrap every other section individually in an HTML `<details>` block using the canonical structure (`<details>` on its own line, `<summary>…</summary>` on its own line, then a blank line before the section body, ending with `</details>`). Use descriptive summary labels that include the section's most important number where possible (e.g. `📋 Summary — 3 accepted outcomes, 43,055 AIC, IE 0.00476`, `🎯 Agentic Work by Objective — top: bug (170 value)`, `📉 Data Quality — 2 gaps`).

The report must include:

### Executive Summary

Write 2–4 sentences that directly answer: *What did the agent work on, what was the highest-impact agentic work, which workflows contributed most to that impact, how efficiently were AIC tokens spent, and what high-impact work was delivered outside agentic workflows (if any)?* Highlight the most impactful objective categories, the workflows contributing the most value, and any significant gaps (e.g., large AIC spend with no mapped objective value).

The Executive Summary must **not** be wrapped in a `<details>` block — it is always visible.

### Summary

*(Wrap this section in `<details><summary>📋 Summary — …</summary>…</details>`.)*

| Metric | Value |
|---|---:|

When a metric includes sub-counts, format the Value as `merged: X, closed: Y, open excluded: Z`.

Include:
- PRs analyzed with linked issue (merged / closed / excluded open)
- PRs excluded (no linked closing issue)
- Safe output outcomes analyzed
- Outcomes mapped to objectives
- Unmapped outcomes
- Accepted outcome count
- Total outcome value
- AI Credits
- Impact Efficiency

### Agentic Work by Objective

*(Wrap this section in `<details><summary>🎯 Agentic Work by Objective — …</summary>…</details>`.)*

Group all **accepted, mapped** outcomes by objective category (the highest-value objective label from the mapping). For each category, list:

- Objective category name and its mapping value
- Number of accepted outcomes in this category
- Total outcome value contributed
- AIC consumed by outcomes in this category — use per-workflow AIC from `aic-by-workflow.json` only for workflows that are attributed to outcomes in this category; if attribution is unavailable, show `—` (not `N/A`) and add a note that workflow attribution is required to compute per-category AIC
- Impact Efficiency for this category (total outcome value / AIC consumed) — show `—` if AIC is unknown for this category
- Representative examples (up to 3 linked outcomes)

Sort categories by total outcome value descending. Also call out separately which category consumed the **most AIC** (highest denominator cost), so readers can see where budget was spent regardless of value delivered.

This section should make the most impactful work in the repository obvious at a glance.

### Which Workflows Drove That Impact

*(Wrap this section in `<details><summary>⚙️ Which Workflows Drove That Impact — …</summary>…</details>`.)*

Group all analyzed outcomes by the workflow that directly produced them. For each workflow, list:

- Workflow name
- Number of accepted, mapped outcomes attributed to this workflow
- Total outcome value contributed
- Share of total delivered outcome value
- AIC consumed by this workflow's analyzed runs
- Workflow Impact Efficiency (contributed value / AIC consumed)
- Top objective categories this workflow contributed to
- Representative examples (up to 3 linked outcomes)

Sort workflows by total outcome value descending. Also call out separately:

- which workflow contributed the **most total value**
- which workflow contributed the **largest share** of delivered value
- which workflow consumed the **most AIC**

If any analyzed outcomes cannot be attributed to a workflow, report an unattributed bucket with counts and total outcome value, but do not rank it alongside named workflows.

### Top outcomes by outcome value

*(Wrap this section in `<details><summary>🏆 Top Outcomes by Value — top outcome value: …</summary>…</details>`.)*

| Outcome | Workflow | Type | Root / Associated Objective | Objective Value | Outcome Value |
|---|---|---|---|---:|---:|

List the top 15 outcomes with highest Outcome Value. Include a link to the PR or issue.

### Unmapped outcomes

*(Wrap this section in `<details><summary>❓ Unmapped Outcomes — N unmapped</summary>…</details>`. If there are no unmapped outcomes, use the summary label `❓ Unmapped Outcomes — none`.)*

| Outcome | Type | Reason objective was not mapped |
|---|---|---|

Only include outcomes that were in scope (linked-issue PRs and safe-output issues) but had no matching label in the objective mapping. Do not include PRs that were excluded for lacking a linked issue — those are already counted in "PRs excluded".

### Interpretation

*(Wrap this section in `<details><summary>💡 Interpretation</summary>…</details>`.)*

Compare:

- accepted outcome count alone
- Impact Efficiency

Explain which one better reflects meaningful delivered value relative to cost.

Also compute and report a **scope-adjusted IE**: using only AIC from workflows that produced at least one PR or safe-output outcome in the analysis window (i.e., workflows listed in `aic-by-workflow.json` whose name matches any workflow that attributed at least one analyzed outcome). This gives a fairer picture of efficiency for outcome-producing workflows. Present both the full-denominator IE and the scope-adjusted IE; label them clearly.

Do **not** describe IE as "artificially" depressed. Instead, explain concretely what the denominator includes (all workflows, including reporting and analysis workflows that produced no PR outcomes), and present the scope-adjusted IE as the supplemental view.

Call out the most significant findings:

- which objective category delivered the most value per AIC
- which workflow contributed the most delivered value
- which workflows consumed AIC with little or no mapped value

### Data quality

*(Wrap this section in `<details><summary>📊 Data Quality — N issues</summary>…</details>`. Count the number of ⚠️ and ❌ items and use that count in the summary label.)*

Mention missing or weak links in:

- PR root tracing and linked-closing-issue coverage (count of PRs excluded for lacking a linked issue)
- safe-output issue label mapping coverage in `.github/objective-mapping.json`
- workflow attribution coverage and any unattributed analyzed outcomes
- AI Credits availability

State whether AI Credits came from deterministic precomputed data or from a live fallback path.

If AI Credits are unavailable, still produce the delivered-value analysis and clearly state that the cost-normalized Impact Efficiency metric could not be computed.

### Recommendations

*(Wrap this section in `<details><summary>✅ Recommendations — N action items</summary>…</details>`. Count only the action items that apply this cycle (i.e., have concrete evidence from this report's data) and use that count in the summary label.)*

Generate **actionable, evidence-grounded recommendations** to improve the accuracy, determinism, and attribution of the next report. Base each recommendation directly on a gap identified in the Data Quality section or a finding from the analysis — do not produce generic advice.

For each recommendation, include:

- A short title (one line)
- The specific gap or finding from this cycle that motivates it (with concrete numbers where available, e.g., "993 PRs excluded for lacking `Closes #N`")
- The expected effect on report accuracy: which metric would improve and by how much, if estimable from available data (e.g., "would convert up to N excluded PRs into analyzable outcomes")
- The owner or mechanism (workflow name, script, PR template change, etc.)

Generate a recommendation for each of the following gap categories **only when the gap is confirmed by this cycle's data**:

1. **Workflow attribution**: If any analyzed outcomes are unattributed (no `workflow_run_id`/`workflow_name` link), recommend the specific change needed to emit that field — e.g., adding a `workflow_run_id` to PR bodies at creation time, or extending the dataset preparation script to join on head-branch naming conventions. Specify which workflow or script would need the change.

2. **Linked-issue coverage**: If a significant number of PRs were excluded for lacking a `Closes #N` link, recommend adding a PR body template or linter that enforces the link for agentic-workflow-created PRs. Include the exclusion count to quantify the impact.

3. **Objective label coverage**: If any in-scope outcomes were unmapped (objective value = 0 with labels present), recommend adding the missing labels to `.github/objective-mapping.json`. List the specific unmapped labels observed this cycle.

4. **AIC per-outcome attribution**: If per-category or per-workflow AIC could not be computed because attribution was missing, recommend the minimal dataset join that would enable it (e.g., matching workflow log entries to outcome PRs by branch name or run ID).

5. **PR dataset cap**: If the merged or closed-unmerged PR dataset was capped at fewer records than the window might contain (check `dataset-manifest.json` for cap warnings), recommend increasing `OBJECTIVE_IMPACT_PR_LIST_LIMIT` or paginating the fetch.

6. **Likely-agentic reclassification**: If a significant number of PRs were moved to the "likely agentic (attribution gap)" bucket in the Human Work section, recommend the specific attribution improvement (e.g., stamping the producing workflow name in the PR body) that would collapse that bucket in the next cycle.

Only include recommendations that are directly supported by data from this cycle. Omit any category where the data shows no gap (e.g., if attribution is fully resolved, omit recommendation 1). Sort recommendations by expected impact on report accuracy, highest first.

### Human Work

*(Wrap this section in `<details><summary>👤 Human Work — N merged PRs</summary>…</details>`.)*

This section is independent of AIC and the agentic efficiency analysis above. It captures pull requests merged in the analysis window that could not be attributed to any GitHub Agentic Workflow run in the deterministic logs.

Identify merged PRs from `/tmp/gh-aw/agent/objective-impact-report/merged-prs-linked.json` that have **no** matching run in `/tmp/gh-aw/agent/objective-impact-report/workflow-logs.json` (i.e., PRs whose author or head branch cannot be linked to any workflow run that produced an outcome).

Before reporting these as human-authored, apply the following filter to identify **likely-agentic PRs** that may appear human due to attribution gaps:

- PR title matches patterns such as `[docs]`, `[linter-miner]`, `[fix]`, `[refactor]`, `[chore]`, or other known bot-prefixes used in this repository.
- PR author is a bot account (login ending in `[bot]` or known agentic accounts).

Report likely-agentic PRs in a separate sub-table labelled **"Likely agentic (attribution gap)"** rather than counting them in the human total. This prevents attribution gaps from inflating the human work count. Explicitly note how many PRs were reclassified.

For the remaining PRs classified as human-authored, treat them as human contributions for reporting. Explicitly note that missing log coverage or attribution gaps can still inflate this count.

For each human-authored merged PR that has a linked closing issue (non-empty `linked_issue_numbers`), use precomputed objective fields from `merged-prs-with-objective.json` when available; otherwise resolve issue labels from linked issues and apply `objective-mapping.json`. Group results by objective category (highest-value mapped label) and report:

- Objective category name and its mapping value
- Number of human-authored merged PRs in this category
- Total objective value contributed
- Representative examples (up to 3 linked PRs)

Also report:

- Total number of merged PRs in the dataset
- Of those: likely-agentic (attribution gap), confirmed-human
- Of confirmed-human: with linked closing issue vs. without
- Of confirmed-human with linked issue: mapped to objective vs. unmapped

Sort categories by total objective value descending. Do **not** compute AIC or Impact Efficiency for this section — human work has no associated AI Credits cost.

## Safe output

Use only `close-issue` and `create-issue`.
