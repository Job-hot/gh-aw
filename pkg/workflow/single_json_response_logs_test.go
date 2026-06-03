//go:build !integration

package workflow

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSingleJSONResponseEnginesParseLogMetrics(t *testing.T) {
	logContent := `not-json
{"response":"first","stats":{"models":{"m1":{"input_tokens":10,"output_tokens":5}},"tools":{"read_file":{}}}}
{"response":"second","stats":{"models":{"m2":{"input_tokens":3,"output_tokens":2}},"tools":{"read_file":{},"glob":{}}}}`

	testCases := []struct {
		name   string
		parser func(string, bool) LogMetrics
	}{
		{
			name: "gemini",
			parser: func(content string, verbose bool) LogMetrics {
				return NewGeminiEngine().ParseLogMetrics(content, verbose)
			},
		},
		{
			name: "antigravity",
			parser: func(content string, verbose bool) LogMetrics {
				return NewAntigravityEngine().ParseLogMetrics(content, verbose)
			},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			metrics := tt.parser(logContent, false)

			assert.Equal(t, 1, metrics.Turns)
			assert.Equal(t, 20, metrics.TokenUsage)

			toolCounts := map[string]int{}
			for _, tool := range metrics.ToolCalls {
				toolCounts[tool.Name] = tool.CallCount
			}
			assert.Equal(t, 2, toolCounts["read_file"])
			assert.Equal(t, 1, toolCounts["glob"])
		})
	}
}
