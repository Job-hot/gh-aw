---
name: Dictation Instructions
description: Instructions for fixing speech-to-text errors and improving text quality in gh-aw documentation and workflows
---

# Dictation Instructions

## Technical Context

GitHub Agentic Workflows (gh-aw) is a Go-based GitHub CLI extension for writing AI-powered workflows in natural language using markdown files that compile to GitHub Actions YAML.

## Project Glossary

The following project-specific technical terms (top 256 by word frequency from documentation, excluding generic English words) should be corrected when encountered in speech-to-text input:
.github/workflows/
.lock.yml
.md
@copilot
AGENTS.md
ANTHROPIC_API_KEY
ANTHROPIC_BASE_URL
AWF_HOST_PATH
Authorization
CLAUDE_CODE_MAX_TURNS
CLAUDE_CODE_OAUTH_TOKEN
CLI
CODEOWNERS
COPILOT_GITHUB_TOKEN
COPILOT_MODEL
DEBUG
DEBUG_COLORS
GEMINI_API_KEY
GHES
GH_AW_ACTION_MODE
GH_AW_AGENT_OUTPUT
GH_AW_AGENT_TOKEN
GH_AW_ALLOWED_DOMAINS
GH_AW_CROSS_REPO_PAT
GH_AW_FEATURES
GH_AW_GITHUB_BLOCKED_USERS
GH_AW_GITHUB_MCP_SERVER_TOKEN
GH_AW_GITHUB_TOKEN
GH_AW_GITHUB_TRUSTED_USERS
GH_AW_MCP_SERVER
GH_AW_MODEL_AGENT_CLAUDE
GH_AW_MODEL_AGENT_CODEX
GH_AW_MODEL_AGENT_COPILOT
GH_AW_MODEL_AGENT_GEMINI
GH_AW_PHASE
GH_AW_PLUGINS_TOKEN
GH_AW_PROMPT
GH_AW_SAFE_OUTPUTS
GH_AW_SAFE_OUTPUTS_PORT
GH_AW_SAFE_OUTPUTS_STAGED
GH_AW_VERSION
GH_AW_WORKFLOW_ID
GH_HOST
GH_TOKEN
GITHUB_ACTOR
GITHUB_OUTPUT
GITHUB_PERSONAL_ACCESS_TOKEN
GITHUB_REF
GITHUB_RUN_NUMBER
GITHUB_SERVER_URL
GITHUB_STEP_SUMMARY
GITHUB_TOKEN
GITHUB_WORKFLOW
HTTP
HTTPS
JSON
LLM
MCP
MCP_GATEWAY_API_KEY
MCP_GATEWAY_PORT
OIDC
OPENAI_API_KEY
OPENAI_BASE_URL
PAT
REQUEST_CHANGES
SARIF
SHA
URL
YAML
acceptEdits
actions/cache
actions/checkout
actions/create-github-app-token
actions/github-script
actions/setup
activation
add-comment
add-labels
add-reviewer
add_comment
agent
agent-artifacts
agent-output
agentic-workflows
allowed
allowed-domains
allowed-files
allowed-github-references
allowed-repos
api-target
approval-labels
approved
assign-to-agent
assign-to-user
assign_copilot
assignees: copilot
audit
auto-merge
auto-triage-issues
aw.json
bash
blocked-users
branch
branches
byok-copilot
bypassPermissions
cache
cache-memory
call-workflow
check_run
check_suite
claude
close-discussion
close-issue
close-pull-request
codex
comment-memory
comment_id
comment_memory
comment_repo
compile
compile-stable
concurrency
container
contents: read
contents: write
context
copilot-cli-deep-research
copilot-requests
create-agent-session
create-discussion
create-issue
create-pull-request
create-pull-request-review-comment
create_pull_request
crush
daily
debug
default
defaults
dependabot
detection
discussion_comment
discussions
discussions: write
dispatch-workflow
draft
edit
engine
environment
expires
fallback-as-issue
features
firewall-audit-logs
fmt
footer
frontmatter
gemini
gh aw audit
gh aw compile
gh aw logs
gh aw mcp
gh aw recompile
gh aw update-actions
gh aw upgrade
gh-aw
gists
git
github
github-actions
github-app
github-token
github.actor
github.repository
github.token
glob
go.mod
grep
headers
hide-comment
hook
hourly
id-token
imports
inlined-imports: true
inputs
integrity-reactions
issue_comment
issues
issues: write
job-discriminator
jobs
jq
label-issue
labels
latest
lfs
link-sub-issue
lint
local
lock-for-agent
logs
markdown
mcp-gateway
mcp-inspector
mcp-scripts
mcp-server
mcp-servers
mcpServers
memory
merge
merge-pull-request
merged
metadata: read
min-integrity
mode
model
monthly
mount-as-clis
needs:
network
noop
on:
opentelemetry
owner/repo
package.json
permissions
pull-requests: write
pull_request
pull_request_number
push
push-to-pull-request-branch
recompile
registry
runs-on
runs-on-slim
safe-outputs
safe_outputs
schedule
skip-if-match
slash_command
staged
steps
steps.sanitized.outputs.text
timeout-minutes
title-prefix
toolsets
trusted-users
ubuntu-latest
update-project
upgrade
uses:
weekly
with
workflow_call
workflow_dispatch

## Fix Speech-to-Text Errors

When fixing dictated text, correct these common misrecognitions:

### GitHub and Git Terms
- "get hub" → github
- "git lab" → gitlab
- "get actions" → github-actions
- "pull request" → pull-request (when used as compound modifier)
- "issue ops" → issueops
- "label ops" → labelops
- "chat ops" → chatops

### Workflow Configuration
- "front matter" → frontmatter
- "safe outputs" → safe-outputs (in configuration context)
- "safe inputs" → safe-inputs (in configuration context)
- "lock file" → .lock.yml or lockfile (depending on context)
- "tool sets" → toolsets
- "M.C.P." or "M C P" → MCP
- "repo memory" → repo-memory (in configuration context)
- "cache memory" → cache-memory (in configuration context)
- "work flow" → workflow
- "timeout minutes" → timeout-minutes
- "runs on" → runs-on
- "min integrity" → min-integrity (in configuration context)
- "mcp gateway" → mcp-gateway
- "mcp scripts" → mcp-scripts
- "staged mode" → staged-mode
- "token weights" → token-weights
- "effective tokens" → effective-tokens

### AI Engines
- "co-pilot" → @copilot
- "copilot" → @copilot (when referring to the bot or engine)
- "code x" → codex
- "cloud" → claude (when referring to the AI engine)
- "gem ini" → gemini (when referring to the AI engine)

### Commands and Operations
- "G.H. A.W." → gh-aw or `gh aw` (depending on context)
- "re-compile" → recompile
- "work flow dispatch" → workflow_dispatch
- "action lint" → actionlint
- "ziz more" → zizmor
- "poo teen" → poutine
- "queue M.D." → qmd

### File Formats and Extensions
- "dot M.D." → .md
- "dot Y.A.M.L." or "dot Y M L" → .yaml or .yml
- "dot lock dot Y M L" → .lock.yml
- "jason" → JSON (when referring to format)
- "wasm" → WebAssembly or wasm (depending on context)

### Technical Patterns
- "A.P.I." → API
- "U.R.L." → URL
- "H.T.T.P." → HTTP
- "H.T.T.P.S." → HTTPS
- "S.H.A." → SHA
- "C.I." → CI
- "G.H." → GH (when referring to GitHub CLI)
- "Y.A.M.L." → YAML
- "O.I.D.C." → OIDC
- "S.A.R.I.F." → SARIF

### Hyphenation Rules
Use hyphens for compound modifiers:
- "safe outputs" → safe-outputs
- "safe inputs" → safe-inputs
- "cache memory" → cache-memory
- "timeout minutes" → timeout-minutes
- "cross repository" → cross-repository
- "pull request" → pull-request (when used as adjective)
- "mcp gateway" → mcp-gateway
- "mcp scripts" → mcp-scripts
- "token weights" → token-weights

### Environment Variables
Capitalize fully: GITHUB_TOKEN, GH_TOKEN, COPILOT_GITHUB_TOKEN, GH_AW_GITHUB_TOKEN, ANTHROPIC_API_KEY, OPENAI_API_KEY, GEMINI_API_KEY, CLAUDE_CODE_OAUTH_TOKEN

### Common Ambiguities
- "their/there/they're" → use context to determine correct spelling
- "its/it's" → its (possessive), it's (it is)
- "your/you're" → your (possessive), you're (you are)

## Clean Up and Improve Text

Remove filler words and improve clarity:

### Remove These Filler Words
- humm, um, uh, uhh, umm
- you know, like, basically, actually, literally
- kind of, sort of, I mean, I think
- right?, okay?, so yeah, well

### Improve Clarity
1. Remove redundant phrases:
   - "in order to" → "to"
   - "at this point in time" → "now"
   - "due to the fact that" → "because"
   - "in the event that" → "if"

2. Make text more concise:
   - Remove unnecessary qualifiers (very, really, quite)
   - Use active voice instead of passive voice
   - Replace wordy phrases with simpler alternatives

3. Maintain technical accuracy:
   - Keep all technical terms from the glossary
   - Preserve code examples and commands exactly
   - Don't simplify technical concepts

## Guidelines

You do not have enough background information to plan or provide code examples.
- do NOT generate code examples
- do NOT plan steps
- focus on fixing speech-to-text errors and improving text quality
- remove filler words (humm, you know, um, uh, like, basically, actually, etc.)
- improve clarity and make text more professional
- maintain the user's intended meaning
