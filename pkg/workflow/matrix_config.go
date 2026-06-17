package workflow

import (
	"maps"
	"strings"
)

const (
	MatrixSafeOutputsPolicySelectFirst = "select-first"
	MatrixSafeOutputsPolicyConcatAll   = "concat-all"
)

// extractAgentMatrixConfig extracts top-level matrix configuration for agent fan-out.
// It accepts GitHub Actions matrix syntax and a gh-aw extension field:
//
//	matrix:
//	  os: [ubuntu-latest, windows-latest]
//	  include: [{...}]
//	  safe-outputs: select-first|concat-all
func (c *Compiler) extractAgentMatrixConfig(frontmatter map[string]any) *AgentMatrixConfig {
	rawMatrix, ok := frontmatter["matrix"]
	if !ok {
		return nil
	}
	matrixMap, ok := rawMatrix.(map[string]any)
	if !ok || len(matrixMap) == 0 {
		return nil
	}

	strategyMatrix := maps.Clone(matrixMap)
	mergePolicy := MatrixSafeOutputsPolicySelectFirst
	if rawPolicy, hasPolicy := strategyMatrix["safe-outputs"]; hasPolicy {
		if policyStr, isString := rawPolicy.(string); isString {
			normalized := strings.TrimSpace(strings.ToLower(policyStr))
			if normalized == MatrixSafeOutputsPolicyConcatAll {
				mergePolicy = MatrixSafeOutputsPolicyConcatAll
			}
		}
		delete(strategyMatrix, "safe-outputs")
	}
	if len(strategyMatrix) == 0 {
		return nil
	}

	return &AgentMatrixConfig{
		StrategyMatrix: strategyMatrix,
		MergePolicy:    mergePolicy,
	}
}
