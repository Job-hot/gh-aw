---
name: Dictation Prompt Generator
description: Generates optimized prompts for voice dictation and speech-to-text workflows
on:
  workflow_dispatch:
  schedule:
    - cron: "weekly on sunday around 6:00"  # ~6 AM UTC on Sundays (scattered)

permissions:
  contents: read
  issues: read
  pull-requests: read

engine: copilot

network: defaults

imports:
  - shared/reporting.md

tools:
  mount-as-clis: true
  edit:
  bash:
    - "*"
  github:
    toolsets: [default]

safe-outputs:
  create-pull-request:
    expires: 2d
    title-prefix: "[docs] "
    labels: [documentation, automation]
    draft: false
    auto-merge: true

timeout-minutes: 10
features:
  mcp-cli: true
  copilot-requests: true
---

# Dictation Prompt Generator

Extract technical vocabulary from documentation files and create a concise dictation instruction file for fixing speech-to-text errors and improving text clarity.

## Your Mission

Create a concise dictation instruction file at `skills/dictation/SKILL.md` that:
1. Contains a glossary of exactly 256 project-specific terms selected from the precomputed word frequency histogram
2. Provides instructions for fixing speech-to-text errors (ambiguous terms, spacing, hyphenation)
3. Provides instructions for "agentifying" text: removing filler words (humm, you know, um, uh, like, etc.), improving clarity, and making text more professional
4. Does NOT include planning guidelines or examples (keep it short and focused on error correction and text cleanup)
5. Includes guidelines to NOT plan or provide examples, just focus on fixing speech-to-text errors and improving text quality.

## Task Steps

### 1. Compute Word Frequency Histogram

Run the following Python script to compute a statistical histogram of technical term occurrence across all documentation files:

```bash
python3 - <<'PYEOF'
import os
import re
from collections import Counter

docs_dir = "docs/src/content/docs"
words = Counter()

for root, dirs, files in os.walk(docs_dir):
    for fname in files:
        if fname.endswith(".md"):
            with open(os.path.join(root, fname)) as f:
                text = f.read()
            # Extract terms from inline code backticks
            terms = re.findall(r'`([^`\n]{2,50})`', text)
            # Extract ALL_CAPS identifiers (env vars, constants)
            terms += re.findall(r'\b([A-Z][A-Z0-9_]{2,})\b', text)
            # Extract @mentions (bot names)
            terms += re.findall(r'(@[a-z][a-z0-9-]+)', text)
            words.update(t.strip() for t in terms if t.strip())

# Exclude generic English words and print top 256 project-specific terms
EXCLUDE = {
    'MUST', 'SHOULD', 'NOT', 'MAY', 'SHALL', 'NOTE', 'RECOMMENDED', 'REQUIRED', 'OPTIONAL',
    'IMPORTANT', 'WARNING', 'TIP', 'true', 'false', 'null', 'none', 'all', 'any', 'max', 'min',
    'yes', 'no', 'new', 'old', 'get', 'set', 'add', 'run', 'use', 'see', 'via', 'in', 'out',
}
filtered = [(t, c) for t, c in words.most_common(512)
            if t not in EXCLUDE and len(t) > 2 and not t.isdigit()]
for term, count in filtered[:256]:
    print(f"{count:4d} {term}")
PYEOF
```

The script prints the top 256 project-specific terms by frequency. Use these terms directly as the glossary — they are already ranked and filtered.

### 2. Scan Documentation for Project-Specific Glossary

Use the histogram output from Step 1 plus targeted searches to validate and supplement term selection:

- `search("workflow configuration frontmatter engine permissions")` — core workflow concepts
- `search("safe-outputs create-pull-request tools MCP server")` — tools and integrations
- `search("compilation CLI commands audit logs")` — CLI and developer tools
- `search("network sandbox runtime activation triggers")` — advanced features

**Focus areas for extraction:**
- Configuration: safe-outputs, permissions, tools, cache-memory, toolset, frontmatter
- Engines: @copilot, claude, codex, custom (use @copilot not copilot)
- Bot mentions: @copilot (for GitHub issue assignment — always use @copilot)
- Commands: compile, audit, logs, mcp, recompile
- GitHub concepts: workflow_dispatch, pull_request, issues, discussions
- Repository-specific: agentic workflows, gh-aw, activation, MCP servers
- File formats: markdown, lockfile (.lock.yml), YAML
- Tool types: edit, bash, github, playwright, web-fetch, web-search
- Operations: fmt, lint, test-unit, timeout-minutes, runs-on

**Exclude**: makefile, Astro, starlight (tooling-specific, not user-facing)

### 3. Create the Dictation Instructions File

Create `skills/dictation/SKILL.md` with:
- Frontmatter with name and description fields
- Title: Dictation Instructions
- Technical Context: Brief description of gh-aw
- Project Glossary: exactly 256 terms from the histogram, alphabetically sorted, one per line
- Fix Speech-to-Text Errors: Common misrecognitions → correct terms (use @copilot not copilot)
- Clean Up and Improve Text: Instructions for removing filler words and improving clarity
- Guidelines: General instructions as follows

```markdown
You do not have enough background information to plan or provide code examples.
- do NOT generate code examples
- do NOT plan steps
- focus on fixing speech-to-text errors and improving text quality
- remove filler words (humm, you know, um, uh, like, basically, actually, etc.)
- improve clarity and make text more professional
- maintain the user's intended meaning
```

### 4. Create Pull Request

Use the create-pull-request tool to submit your changes with:
- Title: "[docs] Update dictation skill instructions"
- Description explaining the changes made to skills/dictation/SKILL.md

## Guidelines

- Run the Python NLP histogram script first to identify high-frequency terms
- Scan only `docs/src/content/docs/**/*.md` files for additional context
- Select exactly 256 terms driven by histogram frequency
- Exclude tooling-specific terms (makefile, Astro, starlight) and generic English words
- Prioritize frequently used project-specific terms
- Always use @copilot (not copilot) when referring to the GitHub Copilot bot
- Alphabetize the glossary
- No descriptions in glossary (just term names)
- Focus on fixing speech-to-text errors, not planning or examples

## Success Criteria

- ✅ File `skills/dictation/SKILL.md` exists
- ✅ Contains proper SKILL.md frontmatter (name, description)
- ✅ Contains exactly 256 project-specific terms from the histogram
- ✅ Terms selected using precomputed word frequency histogram
- ✅ Uses @copilot (not copilot) throughout
- ✅ Focuses on fixing speech-to-text errors
- ✅ Includes instructions for removing filler words and improving text clarity
- ✅ Pull request created with changes

{{#import shared/noop-reminder.md}}
