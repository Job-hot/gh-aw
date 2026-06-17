package workflow

import (
	"fmt"
	"strings"

	"github.com/github/gh-aw/pkg/constants"
)

func (c *Compiler) buildJudgeJob(data *WorkflowData) (*Job, error) {
	if !hasAgentMatrix(data) {
		return nil, nil
	}

	policy := MatrixSafeOutputsPolicySelectFirst
	if data.Matrix != nil && data.Matrix.MergePolicy == MatrixSafeOutputsPolicyConcatAll {
		policy = MatrixSafeOutputsPolicyConcatAll
	}
	agentArtifactPrefix := artifactPrefixExprForDownstreamJob(data)

	steps := []string{
		"      - name: Download matrix agent artifacts\n",
		"        continue-on-error: true\n",
		fmt.Sprintf("        uses: %s\n", c.getActionPin("actions/download-artifact")),
		"        with:\n",
		fmt.Sprintf("          pattern: %sagent_*\n", agentArtifactPrefix),
		"          path: /tmp/gh-aw/matrix-artifacts\n",
		"      - name: Merge matrix safe outputs\n",
		"        env:\n",
		fmt.Sprintf("          GH_AW_MATRIX_MERGE_POLICY: %q\n", policy),
		"        run: |\n",
		"          node <<'EOF'\n",
		"          const fs = require('fs');\n",
		"          const path = require('path');\n",
		"          const artifactsRoot = '/tmp/gh-aw/matrix-artifacts';\n",
		"          const outRoot = '/tmp/gh-aw/judge-agent';\n",
		"          fs.mkdirSync(outRoot, { recursive: true });\n",
		"          function walk(dir) {\n",
		"            if (!fs.existsSync(dir)) return [];\n",
		"            const out = [];\n",
		"            for (const entry of fs.readdirSync(dir, { withFileTypes: true })) {\n",
		"              const full = path.join(dir, entry.name);\n",
		"              if (entry.isDirectory()) out.push(...walk(full));\n",
		"              else out.push(full);\n",
		"            }\n",
		"            return out;\n",
		"          }\n",
		"          const allFiles = walk(artifactsRoot);\n",
		"          const safeoutputsFiles = allFiles.filter((f) => path.basename(f) === 'safeoutputs.jsonl');\n",
		"          const agentOutputFiles = allFiles.filter((f) => path.basename(f) === 'agent_output.json');\n",
		"          const tokenUsageFiles = allFiles.filter((f) => path.basename(f) === 'token-usage.jsonl');\n",
		"          const patchFiles = allFiles.filter((f) => /^aw-.+\\.(patch|bundle)$/.test(path.basename(f)));\n",
		"          if (agentOutputFiles.length > 0) {\n",
		"            const baseDir = path.dirname(agentOutputFiles.sort()[0]);\n",
		"            for (const file of walk(baseDir)) {\n",
		"              const rel = path.relative(baseDir, file);\n",
		"              const dst = path.join(outRoot, rel);\n",
		"              fs.mkdirSync(path.dirname(dst), { recursive: true });\n",
		"              fs.copyFileSync(file, dst);\n",
		"            }\n",
		"          }\n",
		"          const policy = process.env.GH_AW_MATRIX_MERGE_POLICY || 'select-first';\n",
		"          let mergedLines = [];\n",
		"          if (policy === 'concat-all') {\n",
		"            for (const file of safeoutputsFiles.sort()) {\n",
		"              const lines = fs.readFileSync(file, 'utf8').split(/\\r?\\n/).filter(Boolean);\n",
		"              mergedLines.push(...lines);\n",
		"            }\n",
		"          } else {\n",
		"            const selected = safeoutputsFiles.sort().find((file) => fs.readFileSync(file, 'utf8').trim() !== '');\n",
		"            if (selected) mergedLines = fs.readFileSync(selected, 'utf8').split(/\\r?\\n/).filter(Boolean);\n",
		"          }\n",
		"          const mergedSafeOutputsPath = path.join(outRoot, 'safeoutputs.jsonl');\n",
		"          fs.writeFileSync(mergedSafeOutputsPath, mergedLines.length > 0 ? `${mergedLines.join('\\n')}\\n` : '', 'utf8');\n",
		"          const items = [];\n",
		"          const outputTypes = new Set();\n",
		"          for (const line of mergedLines) {\n",
		"            try {\n",
		"              const item = JSON.parse(line);\n",
		"              items.push(item);\n",
		"              if (item && typeof item.type === 'string') outputTypes.add(item.type);\n",
		"            } catch {}\n",
		"          }\n",
		"          fs.writeFileSync(path.join(outRoot, 'agent_output.json'), JSON.stringify({ items, errors: [] }), 'utf8');\n",
		"          for (const patchFile of patchFiles) {\n",
		"            const dst = path.join(outRoot, path.basename(patchFile));\n",
		"            if (!fs.existsSync(dst)) fs.copyFileSync(patchFile, dst);\n",
		"          }\n",
		"          const aggregateDir = path.join(outRoot, 'matrix-token-usage');\n",
		"          fs.mkdirSync(aggregateDir, { recursive: true });\n",
		"          tokenUsageFiles.sort().forEach((file, idx) => {\n",
		"            fs.copyFileSync(file, path.join(aggregateDir, `token-usage-${idx}.jsonl`));\n",
		"          });\n",
		"          const githubOutput = process.env.GITHUB_OUTPUT;\n",
		"          if (githubOutput) {\n",
		"            fs.appendFileSync(githubOutput, `output_types=${Array.from(outputTypes).join(',')}\\n`);\n",
		"            fs.appendFileSync(githubOutput, `has_patch=${patchFiles.length > 0 ? 'true' : 'false'}\\n`);\n",
		"          }\n",
		"          EOF\n",
		"      - name: Upload merged agent artifact\n",
		"        if: always()\n",
		"        continue-on-error: true\n",
		fmt.Sprintf("        uses: %s\n", c.getActionPin("actions/upload-artifact")),
		"        with:\n",
		fmt.Sprintf("          name: %s%s\n", agentArtifactPrefix, constants.AgentArtifactName),
		"          path: /tmp/gh-aw/judge-agent\n",
		"          if-no-files-found: ignore\n",
	}

	outputs := map[string]string{
		"output_types": "${{ steps.merge-matrix.outputs.output_types }}",
		"has_patch":    "${{ steps.merge-matrix.outputs.has_patch }}",
	}
	if hasWorkflowCallTrigger(data.On) {
		outputs[constants.ArtifactPrefixOutputName] = "${{ needs.activation.outputs.artifact_prefix }}"
	}

	// add explicit step ID to output-producing step
	for i, step := range steps {
		if strings.HasPrefix(step, "      - name: Merge matrix safe outputs") {
			steps = append(steps[:i+1], append([]string{"        id: merge-matrix\n"}, steps[i+1:]...)...)
			break
		}
	}

	return &Job{
		Name:        string(constants.JudgeJobName),
		RunsOn:      c.formatFrameworkJobRunsOn(data),
		Needs:       []string{string(constants.AgentJobName), string(constants.ActivationJobName)},
		If:          fmt.Sprintf("always() && needs.%s.result != 'skipped'", constants.AgentJobName),
		Permissions: NewPermissionsContentsRead().RenderToYAML(),
		Steps:       steps,
		Outputs:     outputs,
	}, nil
}
