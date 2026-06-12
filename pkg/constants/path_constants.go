package constants

// Repository (dot-github) path constants.
// These are repository-relative paths used for GitHub configuration files.

// DotGithubDir is the repository-relative path to the .github directory (with trailing slash).
// Use for prefix comparisons and directory operations.
const DotGithubDir = ".github/"

// DotGithubWorkflowsDir is the default repository-relative path to the workflows directory,
// without a trailing slash. Callers that need the trailing slash should append "/".
// Use GetWorkflowDir() to allow override via the GH_AW_WORKFLOWS_DIR environment variable.
const DotGithubWorkflowsDir = ".github/workflows"

// DotGithubAgentsDir is the repository-relative path to the .github/agents directory,
// without a trailing slash. Callers that need the trailing slash should append "/".
const DotGithubAgentsDir = ".github/agents"

// Container-side temporary path constants (/tmp/gh-aw/...).
// These are paths within the agent container filesystem, distinct from the
// runner-side paths rooted at ${RUNNER_TEMP}/gh-aw (see GhAwRootDirShell).

// TmpGhAwDir is the base directory for gh-aw temporary files within the container
// filesystem, without a trailing slash.
const TmpGhAwDir = "/tmp/gh-aw"

// AgentStdioLogPath is the path where the agent process stdout/stderr is logged.
const AgentStdioLogPath = "/tmp/gh-aw/agent-stdio.log"

// AgentDir is the directory for agent-generated output files, without a trailing slash.
const AgentDir = "/tmp/gh-aw/agent"

// AgentPromptFilePath is the path to the prompt file passed to the AI engine.
const AgentPromptFilePath = "/tmp/gh-aw/aw-prompts/prompt.txt"

// MCPLogsDir is the base directory for MCP server logs, without a trailing slash.
const MCPLogsDir = "/tmp/gh-aw/mcp-logs"

// MCPLogsSafeOutputsDir is the directory for safe-outputs MCP server logs.
const MCPLogsSafeOutputsDir = "/tmp/gh-aw/mcp-logs/safeoutputs"

// MCPLogsPlaywrightDir is the directory for Playwright MCP server logs.
const MCPLogsPlaywrightDir = "/tmp/gh-aw/mcp-logs/playwright"

// MCPScriptsLogsDir is the base directory for MCP scripts logs, without a trailing slash.
const MCPScriptsLogsDir = "/tmp/gh-aw/mcp-scripts/logs"

// MCPConfigDir is the directory for MCP configuration files, without a trailing slash.
const MCPConfigDir = "/tmp/gh-aw/mcp-config"

// MCPConfigLogsDir is the directory for MCP configuration logs, without a trailing slash.
const MCPConfigLogsDir = "/tmp/gh-aw/mcp-config/logs"

// MCPConfigServersFilePath is the path to the MCP servers configuration file in the container.
const MCPConfigServersFilePath = "/tmp/gh-aw/mcp-config/mcp-servers.json"

// AwMCPLogsDir is the directory for agentic-workflows MCP server logs.
const AwMCPLogsDir = "/tmp/gh-aw/aw-mcp/logs"

// CommentMemoryDir is the directory for comment memory persistence, without a trailing slash.
const CommentMemoryDir = "/tmp/gh-aw/comment-memory"

// RepoMemoryDir is the base directory for repository memory persistence, without a trailing slash.
const RepoMemoryDir = "/tmp/gh-aw/repo-memory"

// ThreatDetectionLogPath is the path for the threat detection log file.
const ThreatDetectionLogPath = "/tmp/gh-aw/threat-detection/detection.log"

// PiAgentDir is the directory for Pi agent configuration files.
const PiAgentDir = "/tmp/gh-aw/pi-agent-dir"

// AntigravityClientErrorGlob is the glob pattern for Antigravity client error JSON files.
const AntigravityClientErrorGlob = "/tmp/gh-aw/antigravity-client-error-*.json"

// GeminiClientErrorGlob is the glob pattern for Gemini client error JSON files.
const GeminiClientErrorGlob = "/tmp/gh-aw/gemini-client-error-*.json"

// AwPatchGlob is the glob pattern for agent-generated git patch files.
const AwPatchGlob = "/tmp/gh-aw/aw-*.patch"

// AwBundleGlob is the glob pattern for agent-generated git bundle files.
const AwBundleGlob = "/tmp/gh-aw/aw-*.bundle"

// DifcProxyTLSCACertPath is the path to the DIFC proxy TLS CA certificate file.
const DifcProxyTLSCACertPath = "/tmp/gh-aw/difc-proxy-tls/ca.crt"

// ProxyLogsDir is the directory for proxy logs, without a trailing slash.
const ProxyLogsDir = "/tmp/gh-aw/proxy-logs"

// ProxyLogsTLSCACertPath is the path to the proxy TLS CA certificate file.
const ProxyLogsTLSCACertPath = "/tmp/gh-aw/proxy-logs/proxy-tls/ca.crt"

// SandboxAgentLogsDir is the directory for sandbox agent logs, without a trailing slash.
const SandboxAgentLogsDir = "/tmp/gh-aw/sandbox/agent/logs"
