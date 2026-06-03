package workflow

import (
	"encoding/json"
	"strings"

	"github.com/github/gh-aw/pkg/logger"
)

type singleJSONResponse struct {
	Response string         `json:"response"`
	Stats    map[string]any `json:"stats"`
}

// parseSingleJSONResponseLogMetrics parses line-delimited logs where each
// successful turn can emit a full JSON response object containing response and stats.
func parseSingleJSONResponseLogMetrics(logContent string, verbose bool, log *logger.Logger, engineDisplayName string) LogMetrics {
	log.Printf("Parsing %s log metrics: log_size=%d bytes, verbose=%v", engineDisplayName, len(logContent), verbose)

	metrics := LogMetrics{
		Turns:      0,
		TokenUsage: 0,
		ToolCalls:  []ToolCallInfo{},
	}
	toolCallCounts := make(map[string]int)

	lines := strings.SplitSeq(logContent, "\n")
	for line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		var response singleJSONResponse
		if err := json.Unmarshal([]byte(line), &response); err != nil {
			continue
		}

		if response.Response != "" {
			metrics.Turns = 1
		}

		if response.Stats != nil {
			if models, ok := response.Stats["models"].(map[string]any); ok {
				for _, modelStats := range models {
					if stats, ok := modelStats.(map[string]any); ok {
						if inputTokens, ok := stats["input_tokens"].(float64); ok {
							metrics.TokenUsage += int(inputTokens)
						}
						if outputTokens, ok := stats["output_tokens"].(float64); ok {
							metrics.TokenUsage += int(outputTokens)
						}
					}
				}
			}

			if tools, ok := response.Stats["tools"].(map[string]any); ok {
				for toolName := range tools {
					toolCallCounts[toolName]++
				}
			}
		}

		log.Printf("Parsed JSON response: response_len=%d, stats_present=%v", len(response.Response), response.Stats != nil)
	}

	for toolName, count := range toolCallCounts {
		metrics.ToolCalls = append(metrics.ToolCalls, ToolCallInfo{
			Name:      toolName,
			CallCount: count,
		})
	}

	log.Printf("Parsed metrics: turns=%d, token_usage=%d, tool_calls=%d",
		metrics.Turns, metrics.TokenUsage, len(metrics.ToolCalls))

	return metrics
}
