// This file provides command-line interface functionality for gh-aw.
// This file (logs_parsing_runner.go) contains shared Node.js execution
// scaffolding used by the firewall and agent log parsers.

package cli

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/github/gh-aw/pkg/constants"
)

// jsCoreMock is a minimal @actions/core mock suitable for use in CLI-mode
// Node.js log parsers.  It exposes summary.addRaw/write, setFailed, and a
// silent info stub so parser scripts run outside of GitHub Actions without
// modification.
const jsCoreMock = `const core = {
	summary: {
		addRaw: function(content) {
			this._content = content;
			return this;
		},
		write: function() {
			console.log(this._content);
		},
		_content: ''
	},
	setFailed: function(message) {
		console.error('FAILED:', message);
		process.exit(1);
	},
	info: function(message) {
		// Silent in CLI mode
	}
};`

// runNodeScript writes script to a temporary parser.js file, executes it with
// node, and writes the trimmed combined output (stdout and stderr) to
// outputPath.  The temporary directory is cleaned up automatically.
func runNodeScript(script, outputPath string) error {
	tempDir, err := os.MkdirTemp("", "node_parser")
	if err != nil {
		return fmt.Errorf("failed to create temp dir: %w", err)
	}
	defer os.RemoveAll(tempDir)

	nodeFile := filepath.Join(tempDir, "parser.js")
	if err := os.WriteFile(nodeFile, []byte(script), constants.FilePermPublic); err != nil {
		return fmt.Errorf("failed to write node script: %w", err)
	}

	// #nosec G204 -- nodeFile is an absolute path to a file written by this
	// process to tempDir; using exec.Command with separate arguments (not shell
	// execution) prevents shell injection.
	cmd := exec.Command("node", nodeFile)
	cmd.Dir = tempDir
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to execute node script: %w\nOutput: %s", err, strings.TrimSpace(string(output)))
	}

	if err := os.WriteFile(outputPath, []byte(strings.TrimSpace(string(output))), constants.FilePermPublic); err != nil {
		return fmt.Errorf("failed to write %s: %w", filepath.Base(outputPath), err)
	}
	return nil
}
