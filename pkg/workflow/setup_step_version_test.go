//go:build !integration

package workflow

import (
	"strings"
	"testing"

	"github.com/github/gh-aw/pkg/constants"
)

func TestGetVersionForSetup(t *testing.T) {
	tests := []struct {
		name            string
		data            *WorkflowData
		expectedVersion string
	}{
		{
			name:            "nil data returns empty string",
			data:            nil,
			expectedVersion: "",
		},
		{
			name:            "no engine config returns empty string",
			data:            &WorkflowData{},
			expectedVersion: "",
		},
		{
			name: "explicit version in EngineConfig takes priority",
			data: &WorkflowData{
				EngineConfig: &EngineConfig{ID: "copilot", Version: "1.2.3"},
			},
			expectedVersion: "1.2.3",
		},
		{
			name: "copilot engine uses default version",
			data: &WorkflowData{
				EngineConfig: &EngineConfig{ID: "copilot"},
			},
			expectedVersion: string(constants.DefaultCopilotVersion),
		},
		{
			name: "claude engine uses default version",
			data: &WorkflowData{
				EngineConfig: &EngineConfig{ID: "claude"},
			},
			expectedVersion: string(constants.DefaultClaudeCodeVersion),
		},
		{
			name: "codex engine uses default version",
			data: &WorkflowData{
				EngineConfig: &EngineConfig{ID: "codex"},
			},
			expectedVersion: string(constants.DefaultCodexVersion),
		},
		{
			name: "AI field used when EngineConfig.ID is empty",
			data: &WorkflowData{
				AI: "copilot",
			},
			expectedVersion: string(constants.DefaultCopilotVersion),
		},
		{
			name: "EngineConfig.ID takes priority over AI field",
			data: &WorkflowData{
				AI:           "copilot",
				EngineConfig: &EngineConfig{ID: "claude"},
			},
			expectedVersion: string(constants.DefaultClaudeCodeVersion),
		},
		{
			name: "custom engine returns empty string",
			data: &WorkflowData{
				EngineConfig: &EngineConfig{ID: "custom"},
			},
			expectedVersion: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := getVersionForSetup(tt.data)
			if result != tt.expectedVersion {
				t.Errorf("getVersionForSetup() = %q, want %q", result, tt.expectedVersion)
			}
		})
	}
}

func TestGenerateSetupStepIncludesVersion(t *testing.T) {
	tests := []struct {
		name          string
		data          *WorkflowData
		expectVersion string
		noVersionLine bool
	}{
		{
			name: "copilot engine injects default version",
			data: &WorkflowData{
				Name:         "my-workflow",
				EngineConfig: &EngineConfig{ID: "copilot"},
			},
			expectVersion: string(constants.DefaultCopilotVersion),
		},
		{
			name: "explicit version is injected",
			data: &WorkflowData{
				Name:         "my-workflow",
				EngineConfig: &EngineConfig{ID: "copilot", Version: "1.2.3"},
			},
			expectVersion: "1.2.3",
		},
		{
			name: "custom engine without version does not inject GH_AW_INFO_VERSION",
			data: &WorkflowData{
				Name:         "my-workflow",
				EngineConfig: &EngineConfig{ID: "custom"},
			},
			noVersionLine: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewCompiler()
			lines := c.generateSetupStep(tt.data, "github/gh-aw/actions/setup@abc123", "${{ runner.temp }}/gh-aw", false, "", "", false)
			combined := strings.Join(lines, "")

			if tt.noVersionLine {
				if strings.Contains(combined, "GH_AW_INFO_VERSION") {
					t.Errorf("expected no GH_AW_INFO_VERSION in setup step, but found it:\n%s", combined)
				}
				return
			}

			expectedLine := `GH_AW_INFO_VERSION: "` + tt.expectVersion + `"`
			if !strings.Contains(combined, expectedLine) {
				t.Errorf("expected setup step to contain %q, got:\n%s", expectedLine, combined)
			}
		})
	}
}

func TestGenerateSetupStepIncludesAWFVersion(t *testing.T) {
	tests := []struct {
		name            string
		data            *WorkflowData
		expectedVersion string
		expectNoAWFLine bool
	}{
		{
			name: "firewall enabled with explicit version",
			data: &WorkflowData{
				Name: "my-workflow",
				NetworkPermissions: &NetworkPermissions{
					Firewall: &FirewallConfig{
						Enabled: true,
						Version: "v1.2.3-awf",
					},
				},
			},
			expectedVersion: "v1.2.3-awf",
		},
		{
			name: "firewall enabled with default version",
			data: &WorkflowData{
				Name: "my-workflow",
				NetworkPermissions: &NetworkPermissions{
					Firewall: &FirewallConfig{
						Enabled: true,
					},
				},
			},
			expectedVersion: string(constants.DefaultFirewallVersion),
		},
		{
			name: "sandbox agent version overrides firewall version",
			data: &WorkflowData{
				Name: "my-workflow",
				NetworkPermissions: &NetworkPermissions{
					Firewall: &FirewallConfig{
						Enabled: true,
					},
				},
				SandboxConfig: &SandboxConfig{
					Agent: &AgentSandboxConfig{
						Type:    SandboxTypeAWF,
						Version: "v9.9.9-awf",
					},
				},
			},
			expectedVersion: "v9.9.9-awf",
		},
		{
			name: "firewall disabled does not inject GH_AW_INFO_AWF_VERSION",
			data: &WorkflowData{
				Name: "my-workflow",
				NetworkPermissions: &NetworkPermissions{
					Firewall: &FirewallConfig{
						Enabled: false,
					},
				},
			},
			expectNoAWFLine: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewCompiler()
			lines := c.generateSetupStep(tt.data, "github/gh-aw/actions/setup@abc123", "${{ runner.temp }}/gh-aw", false, "", "", false)
			combined := strings.Join(lines, "")

			if tt.expectNoAWFLine {
				if strings.Contains(combined, "GH_AW_INFO_AWF_VERSION") {
					t.Errorf("expected no GH_AW_INFO_AWF_VERSION in setup step, but found it:\n%s", combined)
				}
				return
			}

			expectedLine := `GH_AW_INFO_AWF_VERSION: "` + tt.expectedVersion + `"`
			if !strings.Contains(combined, expectedLine) {
				t.Errorf("expected setup step to contain %q, got:\n%s", expectedLine, combined)
			}
		})
	}
}

func TestGenerateSetupStepIncludesParentSpanID(t *testing.T) {
	c := NewCompiler()
	data := &WorkflowData{Name: "my-workflow"}
	parentExpr := "${{ needs.activation.outputs.setup-span-id }}"

	lines := c.generateSetupStep(data, "github/gh-aw/actions/setup@abc123", "${{ runner.temp }}/gh-aw", false, "", parentExpr, false)
	combined := strings.Join(lines, "")

	if !strings.Contains(combined, "parent-span-id: "+parentExpr) {
		t.Fatalf("expected setup step to include parent-span-id input, got:\n%s", combined)
	}
}

func TestGenerateSetupStepIncludesEngineID(t *testing.T) {
	c := NewCompiler()
	data := &WorkflowData{
		Name:         "my-workflow",
		EngineConfig: &EngineConfig{ID: "copilot"},
	}

	lines := c.generateSetupStep(data, "github/gh-aw/actions/setup@abc123", "${{ runner.temp }}/gh-aw", false, "", "", false)
	combined := strings.Join(lines, "")

	if !strings.Contains(combined, `GH_AW_INFO_ENGINE_ID: "copilot"`) {
		t.Fatalf("expected setup step to include GH_AW_INFO_ENGINE_ID for engine config, got:\n%s", combined)
	}
}

func TestGenerateSetupStepIncludesEngineIDInScriptModeFromAIField(t *testing.T) {
	c := NewCompiler()
	c.SetActionMode(ActionModeScript)
	data := &WorkflowData{
		Name: "my-workflow",
		AI:   "claude",
	}

	lines := c.generateSetupStep(data, "github/gh-aw/actions/setup@abc123", "${{ runner.temp }}/gh-aw", false, "", "", false)
	combined := strings.Join(lines, "")

	if !strings.Contains(combined, `GH_AW_INFO_ENGINE_ID: "claude"`) {
		t.Fatalf("expected setup script step to include GH_AW_INFO_ENGINE_ID from AI field, got:\n%s", combined)
	}
}

// TestGenerateSetupStepEmitsInstallCopilotGate verifies the compiler-controlled
// resolver gate: INPUT_INSTALL_COPILOT='true' and INPUT_GH_AW_VERSION must be
// emitted on the setup step env block when (and only when) the workflow uses
// the Copilot engine AND the caller opts in via installCopilot=true. Other
// jobs that share the setup step but do not invoke the Copilot CLI must not
// emit either env var, so the toolcache resolver in setup.sh stays a no-op
// outside the agent and threat-detection jobs.
func TestGenerateSetupStepEmitsInstallCopilotGate(t *testing.T) {
	c := NewCompiler()

	copilotData := &WorkflowData{
		Name:         "copilot-workflow",
		AI:           "copilot",
		EngineConfig: &EngineConfig{ID: "copilot"},
	}
	claudeData := &WorkflowData{
		Name: "claude-workflow",
		AI:   "claude",
	}

	tests := []struct {
		name           string
		data           *WorkflowData
		installCopilot bool
		wantEmit       bool
	}{
		{name: "copilot engine + opt-in emits both env vars", data: copilotData, installCopilot: true, wantEmit: true},
		{name: "copilot engine without opt-in suppresses env vars", data: copilotData, installCopilot: false, wantEmit: false},
		{name: "non-copilot engine ignores opt-in", data: claudeData, installCopilot: true, wantEmit: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lines := c.generateSetupStep(tt.data, "github/gh-aw/actions/setup@abc123", "${{ runner.temp }}/gh-aw", false, "", "", tt.installCopilot)
			combined := strings.Join(lines, "")
			hasGate := strings.Contains(combined, "INPUT_INSTALL_COPILOT: 'true'")
			hasVersion := strings.Contains(combined, "INPUT_GH_AW_VERSION:")
			if tt.wantEmit {
				if !hasGate {
					t.Errorf("expected INPUT_INSTALL_COPILOT='true' to be emitted, got:\n%s", combined)
				}
				if !hasVersion {
					t.Errorf("expected INPUT_GH_AW_VERSION to be emitted, got:\n%s", combined)
				}
			} else {
				if hasGate {
					t.Errorf("did not expect INPUT_INSTALL_COPILOT to be emitted, got:\n%s", combined)
				}
				if hasVersion {
					t.Errorf("did not expect INPUT_GH_AW_VERSION to be emitted, got:\n%s", combined)
				}
			}
		})
	}
}

func TestGenerateSetupStepIncludesOTLPOIDCMintingBeforeSetup(t *testing.T) {
	c := NewCompiler()
	data := &WorkflowData{
		Name: "my-workflow",
		RawFrontmatter: map[string]any{
			"observability": map[string]any{
				"otlp": map[string]any{
					"github-app": map[string]any{
						"audience": "https://example.com/collector",
					},
				},
			},
		},
	}

	lines := c.generateSetupStep(data, "github/gh-aw/actions/setup@abc123", "${{ runner.temp }}/gh-aw", false, "", "", false)
	combined := strings.Join(lines, "")

	if !strings.Contains(combined, "id: mint-otlp-oidc-token") {
		t.Fatalf("expected setup step to include OTLP OIDC mint step, got:\n%s", combined)
	}
	if !strings.Contains(combined, "otlp-oidc-token: ${{ steps.mint-otlp-oidc-token.outputs.token }}") {
		t.Fatalf("expected setup action input to include minted OTLP OIDC token, got:\n%s", combined)
	}

	mintPos := strings.Index(combined, "id: mint-otlp-oidc-token")
	setupPos := strings.Index(combined, "id: setup")
	if mintPos < 0 || setupPos < 0 || mintPos > setupPos {
		t.Fatalf("expected OTLP OIDC mint step to appear before setup step, got:\n%s", combined)
	}
}

func TestGenerateSetupStepIncludesOTLPOIDCMintingFromParsedFrontmatter(t *testing.T) {
	c := NewCompiler()
	data := &WorkflowData{
		Name: "my-workflow",
		ParsedFrontmatter: &FrontmatterConfig{
			Observability: &ObservabilityConfig{
				OTLP: &OTLPConfig{
					GitHubApp: &OTLPGitHubAppConfig{
						Audience: "https://example.com/collector",
					},
				},
			},
		},
	}

	lines := c.generateSetupStep(data, "github/gh-aw/actions/setup@abc123", "${{ runner.temp }}/gh-aw", false, "", "", false)
	combined := strings.Join(lines, "")

	if !strings.Contains(combined, "id: mint-otlp-oidc-token") {
		t.Fatalf("expected setup step to include OTLP OIDC mint step from parsed frontmatter, got:\n%s", combined)
	}
	if !strings.Contains(combined, "GH_AW_OTLP_OIDC_AUDIENCE") {
		t.Fatalf("expected mint step to include OTLP OIDC audience env from parsed frontmatter, got:\n%s", combined)
	}
	if !strings.Contains(combined, "https://example.com/collector") {
		t.Fatalf("expected mint step to include parsed frontmatter OTLP OIDC audience value, got:\n%s", combined)
	}
}

func TestGenerateSetupStepIncludesOTLPOIDCTokenInScriptMode(t *testing.T) {
	c := NewCompiler()
	c.SetActionMode(ActionModeScript)
	data := &WorkflowData{
		Name: "my-workflow",
		RawFrontmatter: map[string]any{
			"observability": map[string]any{
				"otlp": map[string]any{
					"github-app": map[string]any{},
				},
			},
		},
	}

	lines := c.generateSetupStep(data, "github/gh-aw/actions/setup@abc123", "${{ runner.temp }}/gh-aw", false, "", "", false)
	combined := strings.Join(lines, "")

	if !strings.Contains(combined, "id: mint-otlp-oidc-token") {
		t.Fatalf("expected script mode to include OTLP OIDC mint step, got:\n%s", combined)
	}
	if !strings.Contains(combined, "INPUT_OTLP_OIDC_TOKEN: ${{ steps.mint-otlp-oidc-token.outputs.token }}") {
		t.Fatalf("expected setup.sh env to include minted OTLP OIDC token, got:\n%s", combined)
	}
}
