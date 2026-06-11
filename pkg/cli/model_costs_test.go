//go:build !integration

package cli

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFindModelPricing(t *testing.T) {
	pricing, ok := findModelPricing("anthropic", "claude-sonnet-4.6")
	require.True(t, ok)
	assert.InDelta(t, 0.000003, pricing["input"], 1e-12)
}

func TestComputeModelInferenceAIC(t *testing.T) {
	aic := computeModelInferenceAIC("anthropic", "claude-sonnet-4.6", 1000, 200, 400, 50, 25)
	assert.InDelta(t, 0.54825, aic, 1e-9)
}

func TestNormalizeCatalogProvider(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"github", "github-copilot"},
		{"copilot", "github-copilot"},
		{"github_models", "github-copilot"},
		{"GITHUB_MODELS", "github-copilot"},
		{"anthropic", "anthropic"},
		{"openai", "openai"},
		{"", ""},
	}
	for _, tt := range tests {
		name := tt.input
		if name == "" {
			name = "<empty>"
		}
		t.Run(name, func(t *testing.T) {
			got := normalizeCatalogProvider(tt.input)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestComputeModelInferenceAICGitHubModels(t *testing.T) {
	// provider="github_models" is written by the AWF proxy for Copilot engine runs;
	// it must normalize to "github-copilot" so pricing is found and AIC is non-zero.
	aicViaGitHubModels := computeModelInferenceAIC("github_models", "claude-sonnet-4.6", 1000, 200, 0, 0, 0)
	aicViaGitHubCopilot := computeModelInferenceAIC("github-copilot", "claude-sonnet-4.6", 1000, 200, 0, 0, 0)
	assert.Greater(t, aicViaGitHubModels, 0.0, "github_models provider should produce non-zero AIC")
	assert.InDelta(t, aicViaGitHubCopilot, aicViaGitHubModels, 1e-9, "github_models and github-copilot should yield identical AIC")
}

func TestComputeModelInferenceAICCopilotAlias(t *testing.T) {
	// provider="copilot" is another accepted alias for "github-copilot".
	aicViaCopilot := computeModelInferenceAIC("copilot", "claude-sonnet-4.6", 1000, 200, 0, 0, 0)
	aicViaGitHubCopilot := computeModelInferenceAIC("github-copilot", "claude-sonnet-4.6", 1000, 200, 0, 0, 0)
	assert.Greater(t, aicViaCopilot, 0.0, "copilot provider alias should produce non-zero AIC")
	assert.InDelta(t, aicViaGitHubCopilot, aicViaCopilot, 1e-9, "copilot and github-copilot should yield identical AIC")
}

func TestComputeModelInferenceAICLiveCopilotData(t *testing.T) {
	// Live sample from:
	// https://github.com/github/gh-aw/actions/runs/27355745225/job/80833046748
	// usage artifact: agent/token_usage.jsonl
	tests := []struct {
		name         string
		inputTokens  int
		outputTokens int
		wantAIC      float64
	}{
		{name: "request_1", inputTokens: 20314, outputTokens: 970, wantAIC: 6.5335},
		{name: "request_2", inputTokens: 24162, outputTokens: 1534, wantAIC: 8.3415},
		{name: "request_3", inputTokens: 54255, outputTokens: 4411, wantAIC: 20.18025},
		{name: "request_4", inputTokens: 65750, outputTokens: 224, wantAIC: 16.7735},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := computeModelInferenceAIC("copilot", "gpt-5.4-2026-03-05", tt.inputTokens, tt.outputTokens, 0, 0, 0)
			assert.InDelta(t, tt.wantAIC, got, 1e-9)
		})
	}
}

func TestComputeModelInferenceAICCopilotCacheRead(t *testing.T) {
	// gpt-5.4 pricing in data/models.json:
	// input=0.0000025, output=0.000015, cache_read=0.00000025 (USD/token)
	// expected USD: 1000*0.0000025 + 200*0.000015 + 400*0.00000025 = 0.0056 → 0.56 AIC
	got := computeModelInferenceAIC("copilot", "gpt-5.4-2026-03-05", 1000, 200, 400, 0, 0)
	assert.InDelta(t, 0.56, got, 1e-9)
}
