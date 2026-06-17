package workflow

import (
	"slices"
	"strings"
)

func hasAgentMatrix(data *WorkflowData) bool {
	return data != nil && data.Matrix != nil && len(data.Matrix.StrategyMatrix) > 0
}

func matrixArtifactSuffixExpression(data *WorkflowData) string {
	if !hasAgentMatrix(data) {
		return ""
	}
	keys := make([]string, 0, len(data.Matrix.StrategyMatrix))
	for key := range data.Matrix.StrategyMatrix {
		if key == "include" || key == "exclude" {
			continue
		}
		keys = append(keys, key)
	}
	if len(keys) == 0 {
		return ""
	}
	slices.Sort(keys)
	var b strings.Builder
	for _, key := range keys {
		b.WriteString("_${{ matrix." + key + " }}")
	}
	return b.String()
}
