// This file provides Copilot engine tool permission and error pattern logic.
//
// This file handles three key responsibilities:
//
//  1. Tool Permission Arguments (computeCopilotToolArguments):
//     Converts workflow tool configurations into --allow-tool flags for Copilot CLI.
//     Handles bash/shell tools, edit tools, safe outputs, mcp-scripts, and MCP servers.
//     Supports granular permissions (e.g., "github(get_file)") and server-level wildcards.
//
//  2. Tool Argument Comments (generateCopilotToolArgumentsComment):
//     Generates human-readable comments documenting which tool permissions are granted.
//     Used in compiled workflows for transparency and debugging.
//
//  3. Error Patterns (GetErrorPatterns):
//     Defines regex patterns for extracting error messages from Copilot CLI logs.
//     Includes timestamped log formats, command failures, module errors, and permission issues.
//     Used by log parsers to detect and categorize errors.
//
// These functions are grouped together because they all relate to tool configuration
// and error handling in the Copilot engine.

package workflow

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/github/gh-aw/pkg/console"
	"github.com/github/gh-aw/pkg/constants"
	"github.com/github/gh-aw/pkg/logger"
)

var copilotEngineToolsLog = logger.New("workflow:copilot_engine_tools")

// sanitizeCopilotShellCommand truncates a bash tool command at the first single
// quote to produce a safe prefix for the Copilot CLI --allow-tool shell() argument.
//
// Copilot CLI uses prefix matching for shell() arguments, so shell(jq) matches any
// jq invocation including "jq '.filter' ...". Single quotes in the allow-tool argument
// cause the Copilot CLI to crash at startup because of quoting conflicts in the
// multi-level shell escaping required by the AWF entrypoint.
//
// Returns the sanitized command and whether sanitization was needed.
func sanitizeCopilotShellCommand(cmdStr string) (string, bool) {
	prefix, _, found := strings.Cut(cmdStr, "'")
	if !found {
		return cmdStr, false
	}
	// Trim trailing whitespace from the prefix.
	// shell(jq) prefix-matches any jq invocation, preserving full tool access.
	return strings.TrimRight(prefix, " "), true
}

// computeCopilotToolArguments computes the --allow-tool arguments for Copilot CLI based on tool configurations.
// It handles bash/shell tools, edit tools, safe outputs, mcp-scripts, and MCP server tools.
// Returns a sorted list of arguments ready to be passed to the Copilot CLI.
func (e *CopilotEngine) computeCopilotToolArguments(tools map[string]any, safeOutputs *SafeOutputsConfig, mcpScripts *MCPScriptsConfig, workflowData *WorkflowData) []string {
	copilotEngineToolsLog.Printf("Computing tool arguments: tools=%d", len(tools))
	if tools == nil {
		tools = make(map[string]any)
	}

	var args []string
	hasRestrictedBashAllowlist := false

	// Check if bash has wildcard - if so, use --allow-all-tools instead
	if bashConfig, hasBash := tools["bash"]; hasBash {
		if bashCommands, ok := bashConfig.([]any); ok {
			for _, cmd := range bashCommands {
				if cmdStr, ok := cmd.(string); ok && (cmdStr == ":*" || cmdStr == "*") {
					copilotEngineToolsLog.Print("Bash wildcard detected, using --allow-all-tools")
					return []string{"--allow-all-tools"}
				}
			}
			hasRestrictedBashAllowlist = true
			args = appendSpecificBashCommandArgs(args, bashCommands)
		} else {
			args = append(args, "--allow-tool", "shell")
		}
	}

	args = e.appendMCPMountedBashArgs(args, workflowData, tools, safeOutputs, mcpScripts, hasRestrictedBashAllowlist)

	// Handle edit tools requirement for file write access
	if _, hasEdit := tools["edit"]; hasEdit {
		copilotEngineToolsLog.Print("Edit tool enabled, adding write permission")
		args = append(args, "--allow-tool", "write")
	}

	// Handle safe_outputs MCP server
	if HasSafeOutputsEnabled(safeOutputs) {
		copilotEngineToolsLog.Print("Safe-outputs enabled, adding MCP server permission")
		args = append(args, "--allow-tool", constants.SafeOutputsMCPServerID.String())
	}

	// Handle mcp_scripts MCP server
	if IsMCPScriptsEnabled(mcpScripts) {
		args = append(args, "--allow-tool", constants.MCPScriptsMCPServerID.String())
	}

	// Handle web-fetch builtin tool (Copilot CLI uses web_fetch with underscore)
	if _, hasWebFetch := tools["web-fetch"]; hasWebFetch {
		copilotEngineToolsLog.Print("Web-fetch tool enabled, adding web_fetch permission")
		args = append(args, "--allow-tool", "web_fetch")
	}

	args = appendMCPServerToolArgs(args, tools)

	copilotEngineToolsLog.Printf("Computed %d tool arguments", len(args)/2)
	return sortAndDeduplicateCopilotArgs(args)
}

// appendSpecificBashCommandArgs appends --allow-tool shell(...) flags for each specific bash
// command in the allowlist. It normalises trailing wildcards and sanitizes single quotes.
func appendSpecificBashCommandArgs(args []string, bashCommands []any) []string {
	for _, cmd := range bashCommands {
		cmdStr, ok := cmd.(string)
		if !ok {
			continue
		}
		cmdStr, _ = normalizeBashCommand(cmdStr)
		if !strings.Contains(cmdStr, ":") && !strings.Contains(cmdStr, " ") && constants.CopilotStemCommands[cmdStr] {
			args = append(args, "--allow-tool", fmt.Sprintf("shell(%s:*)", cmdStr))
		} else {
			sanitized, wasSanitized := sanitizeCopilotShellCommand(cmdStr)
			if wasSanitized {
				fmt.Fprintln(os.Stderr, console.FormatWarningMessage(
					fmt.Sprintf("bash tool %q contains single quotes that crash Copilot CLI; "+
						"truncated to safe prefix %q for shell() prefix-matching. "+
						"Use %q in your workflow to silence this warning.",
						cmdStr, sanitized, sanitized)))
			}
			args = append(args, "--allow-tool", fmt.Sprintf("shell(%s)", sanitized))
		}
	}
	return args
}

// appendMCPMountedBashArgs ensures mounted MCP CLI commands and special tools (playwright-cli,
// gh) are executable via shell(<server>:*) when a restricted bash allowlist is active.
func (e *CopilotEngine) appendMCPMountedBashArgs(args []string, workflowData *WorkflowData, tools map[string]any, safeOutputs *SafeOutputsConfig, mcpScripts *MCPScriptsConfig, hasRestrictedBashAllowlist bool) []string {
	if !hasRestrictedBashAllowlist {
		return args
	}
	effectiveWorkflowData := buildCLIWorkflowDataForMounts(workflowData, tools, safeOutputs, mcpScripts)
	for _, serverName := range getMountedCLIServerNamesIfBashRestricted(effectiveWorkflowData, tools, safeOutputs, mcpScripts) {
		args = append(args, "--allow-tool", fmt.Sprintf("shell(%s:*)", serverName))
	}
	if workflowData != nil && isPlaywrightCLIMode(workflowData.Tools) {
		args = append(args, "--allow-tool", "shell(playwright-cli:*)")
	}
	if isGitHubCLIModeEnabled(effectiveWorkflowData) {
		args = append(args, "--allow-tool", "shell(gh:*)")
	}
	return args
}

// appendGitHubMCPToolArgs appends --allow-tool flags for the github tool configuration,
// handling both wildcard (server-level) and specific allowed-tool permissions.
func appendGitHubMCPToolArgs(args []string, toolConfig any) []string {
	toolConfigMap, ok := toolConfig.(map[string]any)
	if !ok {
		// GitHub tool exists but is not a map (e.g., github: null) - allow entire server
		return append(args, "--allow-tool", "github")
	}
	allowed, hasAllowed := toolConfigMap["allowed"]
	if !hasAllowed {
		// No allowed field specified - allow entire GitHub MCP server
		return append(args, "--allow-tool", "github")
	}
	allowedList, ok := allowed.([]any)
	if !ok {
		return append(args, "--allow-tool", "github")
	}
	hasWildcard := false
	for _, allowedTool := range allowedList {
		toolStr, ok := allowedTool.(string)
		if !ok {
			continue
		}
		if toolStr == "*" {
			hasWildcard = true
		} else {
			args = append(args, "--allow-tool", fmt.Sprintf("github(%s)", toolStr))
		}
	}
	if hasWildcard {
		args = append(args, "--allow-tool", "github")
	}
	return args
}

// appendMCPServerToolArgs appends --allow-tool flags for all MCP server tools in the tools map,
// skipping built-in tools (bash, edit, web-search, playwright) that are handled separately.
func appendMCPServerToolArgs(args []string, tools map[string]any) []string {
	// Built-in tool names skipped here (github and web-fetch need explicit handling elsewhere)
	builtInTools := map[string]bool{
		"bash":       true,
		"edit":       true,
		"web-search": true,
		"playwright": true,
	}
	for toolName, toolConfig := range tools {
		if builtInTools[toolName] {
			continue
		}
		if toolName == "github" {
			args = appendGitHubMCPToolArgs(args, toolConfig)
			continue
		}
		toolConfigMap, ok := toolConfig.(map[string]any)
		if !ok {
			continue
		}
		if hasMcp, _ := hasMCPConfig(toolConfigMap); hasMcp {
			copilotEngineToolsLog.Printf("Adding custom MCP server permission: %s", toolName)
			args = append(args, "--allow-tool", toolName)
			if allowed, hasAllowed := toolConfigMap["allowed"]; hasAllowed {
				if allowedList, ok := allowed.([]any); ok {
					for _, allowedTool := range allowedList {
						if toolStr, ok := allowedTool.(string); ok {
							args = append(args, "--allow-tool", fmt.Sprintf("%s(%s)", toolName, toolStr))
						}
					}
				}
			}
		}
	}
	return args
}

// sortAndDeduplicateCopilotArgs sorts and deduplicates the --allow-tool values in args,
// rebuilding the flag-value pairs in sorted order. Deduplication is needed because
// sanitizeCopilotShellCommand can truncate multiple commands to the same safe prefix.
func sortAndDeduplicateCopilotArgs(args []string) []string {
	if len(args) == 0 {
		return args
	}
	var values []string
	for i := 1; i < len(args); i += 2 {
		values = append(values, args[i])
	}
	sort.Strings(values)
	newArgs := make([]string, 0, len(args))
	prev := ""
	for _, value := range values {
		if value == prev {
			continue
		}
		newArgs = append(newArgs, "--allow-tool", value)
		prev = value
	}
	return newArgs
}

// generateCopilotToolArgumentsComment generates a multi-line comment showing each tool argument.
// This is used to document which tool permissions are being granted in the compiled workflow.
func (e *CopilotEngine) generateCopilotToolArgumentsComment(tools map[string]any, safeOutputs *SafeOutputsConfig, mcpScripts *MCPScriptsConfig, workflowData *WorkflowData, indent string) string {
	toolArgs := e.computeCopilotToolArguments(tools, safeOutputs, mcpScripts, workflowData)
	if len(toolArgs) == 0 {
		return ""
	}

	var comment strings.Builder
	comment.WriteString(indent + "# Copilot CLI tool arguments (sorted):\n")

	// Group flag-value pairs for better readability
	for i := 0; i < len(toolArgs); i += 2 {
		if i+1 < len(toolArgs) {
			fmt.Fprintf(&comment, "%s# %s %s\n", indent, toolArgs[i], toolArgs[i+1])
		}
	}

	return comment.String()
}
