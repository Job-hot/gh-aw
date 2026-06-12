package workflow

import "github.com/github/gh-aw/pkg/constants"

// SetupActionDestination is the path where the setup action copies script files
// on the agent runner (e.g. ${{ runner.temp }}/gh-aw/actions).
// Used in YAML `with:` fields that need GitHub Actions expression resolution.
const SetupActionDestination = constants.GhAwRootDir + "/actions"

// SetupActionDestinationShell is the same as SetupActionDestination but uses the
// shell env var form for use inside `run:` blocks.
const SetupActionDestinationShell = constants.GhAwRootDirShell + "/actions"

// SafeOutputsDir is the directory for safe-outputs files on the runner.
// Used in YAML `with:` and `env:` fields that need GitHub Actions expression resolution.
const SafeOutputsDir = constants.GhAwRootDir + "/safeoutputs"

// SafeOutputsDirShell is the same as SafeOutputsDir but uses the shell env var form.
const SafeOutputsDirShell = constants.GhAwRootDirShell + "/safeoutputs"

// GhAwMCPScriptsDir is the directory for MCP scripts files on the runner
const GhAwMCPScriptsDir = constants.GhAwRootDirShell + "/mcp-scripts"

// GhAwBinaryPath is the path to the gh-aw binary on the runner
const GhAwBinaryPath = constants.GhAwRootDirShell + "/gh-aw"

// SafeJobsDownloadDir is the directory for safe job files on the runner (shell env var form).
const SafeJobsDownloadDir = constants.GhAwRootDirShell + "/safe-jobs/"

// SafeJobsDirGHAExpr is the same as SafeJobsDownloadDir but uses the GitHub Actions
// expression form for use in YAML `with:` fields.
const SafeJobsDirGHAExpr = constants.GhAwRootDir + "/safe-jobs/"

// MCPConfigServersFilePathShell is the path to the MCP servers configuration file on the
// runner (shell env var form), for use inside `run:` blocks.
const MCPConfigServersFilePathShell = constants.GhAwRootDirShell + "/mcp-config/mcp-servers.json"

// MCPConfigServersFilePathGHAExpr is the same as MCPConfigServersFilePathShell but uses the
// GitHub Actions expression form for use in YAML `with:` and `env:` fields.
const MCPConfigServersFilePathGHAExpr = constants.GhAwRootDir + "/mcp-config/mcp-servers.json"

// MCPConfigTomlPathGHAExpr is the path to the Codex MCP configuration TOML file on the
// runner, in GitHub Actions expression form.
const MCPConfigTomlPathGHAExpr = constants.GhAwRootDir + "/mcp-config/config.toml"

// SafeOutputsUploadArtifactsDirShell is the path to the safe-outputs upload-artifacts
// staging directory on the runner (shell env var form).
const SafeOutputsUploadArtifactsDirShell = SafeOutputsDirShell + "/upload-artifacts"
