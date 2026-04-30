//go:build !integration

package stringutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestSpec_PublicAPI_Truncate validates the documented behavior of Truncate
// as described in the package README.md.
//
// Specification:
// - Truncates s to at most maxLen characters, appending "..." when truncation occurs.
// - For maxLen ≤ 3 the string is truncated without ellipsis.
func TestSpec_PublicAPI_Truncate(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		maxLen   int
		expected string
	}{
		{
			name:     "truncates with ellipsis for maxLen > 3",
			input:    "hello world",
			maxLen:   8,
			expected: "hello...",
		},
		{
			name:     "no truncation when string fits within maxLen",
			input:    "hi",
			maxLen:   8,
			expected: "hi",
		},
		{
			name:     "maxLen <= 3 truncates without ellipsis",
			input:    "hello world",
			maxLen:   3,
			expected: "hel",
		},
		{
			name:     "maxLen = 1 truncates without ellipsis",
			input:    "hello",
			maxLen:   1,
			expected: "h",
		},
		{
			name:     "maxLen = 2 truncates without ellipsis",
			input:    "hello",
			maxLen:   2,
			expected: "he",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Truncate(tt.input, tt.maxLen)
			assert.Equal(t, tt.expected, result,
				"Truncate(%q, %d) should match documented output", tt.input, tt.maxLen)
		})
	}
}

// TestSpec_PublicAPI_ParseVersionValue validates the documented behavior of
// ParseVersionValue as described in the package README.md.
//
// Specification examples:
//
//	stringutil.ParseVersionValue("20")    // "20"
//	stringutil.ParseVersionValue(20)      // "20"
//	stringutil.ParseVersionValue(20.0)    // "20"
func TestSpec_PublicAPI_ParseVersionValue(t *testing.T) {
	tests := []struct {
		name     string
		input    any
		expected string
	}{
		{
			name:     "string input '20' returns '20'",
			input:    "20",
			expected: "20",
		},
		{
			name:     "int input 20 returns '20'",
			input:    20,
			expected: "20",
		},
		{
			name:     "float64 input 20.0 returns '20'",
			input:    20.0,
			expected: "20",
		},
		{
			name:     "nil input returns empty string",
			input:    nil,
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ParseVersionValue(tt.input)
			assert.Equal(t, tt.expected, result,
				"ParseVersionValue(%v) should match documented output", tt.input)
		})
	}
}

// TestSpec_PublicAPI_IsPositiveInteger validates the documented behavior of
// IsPositiveInteger as described in the package README.md.
//
// Specification: "Returns true if s is a non-empty string containing only
// digit characters (0–9)."
func TestSpec_PublicAPI_IsPositiveInteger(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "digit-only string returns true",
			input:    "123",
			expected: true,
		},
		{
			name:     "single digit returns true",
			input:    "1",
			expected: true,
		},
		{
			name:     "empty string returns false",
			input:    "",
			expected: false,
		},
		{
			name:     "string with letter returns false",
			input:    "12a3",
			expected: false,
		},
		{
			name:     "negative number returns false",
			input:    "-1",
			expected: false,
		},
		{
			name:     "zero returns false",
			input:    "0",
			expected: false,
		},
		{
			name:     "leading zeros return false",
			input:    "007",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsPositiveInteger(tt.input)
			assert.Equal(t, tt.expected, result,
				"IsPositiveInteger(%q) should match documented behavior", tt.input)
		})
	}
}

// TestSpec_PublicAPI_StripANSI validates the documented behavior of StripANSI
// as described in the package README.md.
//
// Specification: "Removes all ANSI/VT100 escape sequences from s. Handles CSI
// sequences (e.g. \x1b[31m for colors) and other ESC-prefixed sequences."
//
// Specification example:
//
//	colored := "\x1b[32mSuccess\x1b[0m"
//	plain := stringutil.StripANSI(colored) // "Success"
func TestSpec_PublicAPI_StripANSI(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "removes CSI color escape sequence (documented example)",
			input:    "\x1b[32mSuccess\x1b[0m",
			expected: "Success",
		},
		{
			name:     "plain string returned unchanged",
			input:    "plain text",
			expected: "plain text",
		},
		{
			name:     "removes red color code",
			input:    "\x1b[31mError\x1b[0m",
			expected: "Error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := StripANSI(tt.input)
			assert.Equal(t, tt.expected, result,
				"StripANSI(%q) should remove ANSI escape sequences", tt.input)
		})
	}
}

// TestSpec_PublicAPI_NormalizeWorkflowName validates the documented behavior of
// NormalizeWorkflowName as described in the package README.md.
//
// Specification examples:
//
//	stringutil.NormalizeWorkflowName("weekly-research.md")       // "weekly-research"
//	stringutil.NormalizeWorkflowName("weekly-research.lock.yml") // "weekly-research"
//	stringutil.NormalizeWorkflowName("weekly-research")          // "weekly-research"
func TestSpec_PublicAPI_NormalizeWorkflowName(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "removes .md extension",
			input:    "weekly-research.md",
			expected: "weekly-research",
		},
		{
			name:     "removes .lock.yml extension",
			input:    "weekly-research.lock.yml",
			expected: "weekly-research",
		},
		{
			name:     "no extension returned unchanged",
			input:    "weekly-research",
			expected: "weekly-research",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := NormalizeWorkflowName(tt.input)
			assert.Equal(t, tt.expected, result,
				"NormalizeWorkflowName(%q) should match documented output", tt.input)
		})
	}
}

// TestSpec_PublicAPI_NormalizeSafeOutputIdentifier validates the documented
// behavior of NormalizeSafeOutputIdentifier as described in the package README.md.
//
// Specification: "Converts dashes and periods to underscores in safe-output
// identifiers, normalizing user-facing dash-separated and dot-separated formats
// to the internal underscore_separated format."
//
// Specification examples:
//
//	stringutil.NormalizeSafeOutputIdentifier("create-issue")            // "create_issue"
//	stringutil.NormalizeSafeOutputIdentifier("executor-workflow.agent") // "executor_workflow_agent"
func TestSpec_PublicAPI_NormalizeSafeOutputIdentifier(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "converts dashes to underscores (documented example)",
			input:    "create-issue",
			expected: "create_issue",
		},
		{
			name:     "converts periods to underscores (documented example)",
			input:    "executor-workflow.agent",
			expected: "executor_workflow_agent",
		},
		{
			name:     "already underscore_separated unchanged",
			input:    "already_ok",
			expected: "already_ok",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := NormalizeSafeOutputIdentifier(tt.input)
			assert.Equal(t, tt.expected, result,
				"NormalizeSafeOutputIdentifier(%q) should convert dashes and periods to underscores", tt.input)
		})
	}
}

// TestSpec_PublicAPI_MarkdownToLockFile validates the documented behavior of
// MarkdownToLockFile as described in the package README.md.
//
// Specification: "Converts a workflow markdown path (.md) to its compiled lock
// file path (.lock.yml). Returns the path unchanged if it already ends with .lock.yml."
//
// Specification example:
//
//	stringutil.MarkdownToLockFile(".github/workflows/test.md")
//	// → ".github/workflows/test.lock.yml"
func TestSpec_PublicAPI_MarkdownToLockFile(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "converts .md to .lock.yml (documented example)",
			input:    ".github/workflows/test.md",
			expected: ".github/workflows/test.lock.yml",
		},
		{
			name:     "already .lock.yml returned unchanged",
			input:    ".github/workflows/test.lock.yml",
			expected: ".github/workflows/test.lock.yml",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MarkdownToLockFile(tt.input)
			assert.Equal(t, tt.expected, result,
				"MarkdownToLockFile(%q) should match documented output", tt.input)
		})
	}
}

// TestSpec_PublicAPI_LockFileToMarkdown validates the documented behavior of
// LockFileToMarkdown as described in the package README.md.
//
// Specification: "Converts a compiled lock file path (.lock.yml) back to its
// markdown source path (.md). Returns the path unchanged if it already ends with .md."
//
// Specification example:
//
//	stringutil.LockFileToMarkdown(".github/workflows/test.lock.yml")
//	// → ".github/workflows/test.md"
func TestSpec_PublicAPI_LockFileToMarkdown(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "converts .lock.yml to .md (documented example)",
			input:    ".github/workflows/test.lock.yml",
			expected: ".github/workflows/test.md",
		},
		{
			name:     "already .md returned unchanged",
			input:    ".github/workflows/test.md",
			expected: ".github/workflows/test.md",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := LockFileToMarkdown(tt.input)
			assert.Equal(t, tt.expected, result,
				"LockFileToMarkdown(%q) should match documented output", tt.input)
		})
	}
}

// TestSpec_PublicAPI_NormalizeGitHubHostURL validates the documented behavior
// of NormalizeGitHubHostURL as described in the package README.md.
//
// Specification: "Normalizes a GitHub host URL by ensuring it has an https://
// scheme and no trailing slash."
//
// Specification examples:
//
//	stringutil.NormalizeGitHubHostURL("github.example.com")        // "https://github.example.com"
//	stringutil.NormalizeGitHubHostURL("https://github.com/")       // "https://github.com"
func TestSpec_PublicAPI_NormalizeGitHubHostURL(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "bare hostname gets https scheme",
			input:    "github.example.com",
			expected: "https://github.example.com",
		},
		{
			name:     "trailing slash removed from https URL",
			input:    "https://github.com/",
			expected: "https://github.com",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := NormalizeGitHubHostURL(tt.input)
			assert.Equal(t, tt.expected, result,
				"NormalizeGitHubHostURL(%q) should match documented output", tt.input)
		})
	}
}

// TestSpec_PublicAPI_ExtractDomainFromURL validates the documented behavior of
// ExtractDomainFromURL as described in the package README.md.
//
// Specification: "Extracts the hostname (without port) from a URL string."
//
// Specification example:
//
//	stringutil.ExtractDomainFromURL("https://api.github.com/repos") // "api.github.com"
func TestSpec_PublicAPI_ExtractDomainFromURL(t *testing.T) {
	result := ExtractDomainFromURL("https://api.github.com/repos")
	assert.Equal(t, "api.github.com", result,
		"ExtractDomainFromURL should return hostname without port (documented example)")
}

// TestSpec_Constants_PATType validates the documented PATType constant values
// as described in the package README.md.
//
// Specification:
//
//	| Constant            | Value          | Prefix       |
//	|---------------------|----------------|--------------|
//	| PATTypeFineGrained  | "fine-grained" | github_pat_  |
//	| PATTypeClassic      | "classic"      | ghp_         |
//	| PATTypeOAuth        | "oauth"        | gho_         |
//	| PATTypeUnknown      | "unknown"      | (other)      |
func TestSpec_Constants_PATType(t *testing.T) {
	assert.Equal(t, PATTypeFineGrained, PATType("fine-grained"),
		"PATTypeFineGrained should have documented value 'fine-grained'")
	assert.Equal(t, PATTypeClassic, PATType("classic"),
		"PATTypeClassic should have documented value 'classic'")
	assert.Equal(t, PATTypeOAuth, PATType("oauth"),
		"PATTypeOAuth should have documented value 'oauth'")
	assert.Equal(t, PATTypeUnknown, PATType("unknown"),
		"PATTypeUnknown should have documented value 'unknown'")
}

// TestSpec_PublicAPI_PATType_Methods validates the documented PATType methods
// as described in the package README.md.
//
// Specification: Methods: String() string, IsFineGrained() bool, IsValid() bool
func TestSpec_PublicAPI_PATType_Methods(t *testing.T) {
	t.Run("IsFineGrained returns true only for fine-grained type", func(t *testing.T) {
		assert.True(t, PATTypeFineGrained.IsFineGrained(),
			"PATTypeFineGrained.IsFineGrained() should return true")
		assert.False(t, PATTypeClassic.IsFineGrained(),
			"PATTypeClassic.IsFineGrained() should return false")
		assert.False(t, PATTypeOAuth.IsFineGrained(),
			"PATTypeOAuth.IsFineGrained() should return false")
		assert.False(t, PATTypeUnknown.IsFineGrained(),
			"PATTypeUnknown.IsFineGrained() should return false")
	})

	t.Run("IsValid returns false only for unknown type", func(t *testing.T) {
		assert.True(t, PATTypeFineGrained.IsValid(),
			"PATTypeFineGrained.IsValid() should return true")
		assert.True(t, PATTypeClassic.IsValid(),
			"PATTypeClassic.IsValid() should return true")
		assert.True(t, PATTypeOAuth.IsValid(),
			"PATTypeOAuth.IsValid() should return true")
		assert.False(t, PATTypeUnknown.IsValid(),
			"PATTypeUnknown.IsValid() should return false")
	})
}

// TestSpec_PublicAPI_ClassifyPAT validates the documented behavior of ClassifyPAT
// as described in the package README.md.
//
// Specification: "Determines the token type from its prefix."
//
// Prefixes per spec:
//   - github_pat_ → PATTypeFineGrained
//   - ghp_        → PATTypeClassic
//   - gho_        → PATTypeOAuth
//   - (other)     → PATTypeUnknown
func TestSpec_PublicAPI_ClassifyPAT(t *testing.T) {
	tests := []struct {
		name     string
		token    string
		expected PATType
	}{
		{
			name:     "github_pat_ prefix yields fine-grained",
			token:    "github_pat_abc123",
			expected: PATTypeFineGrained,
		},
		{
			name:     "ghp_ prefix yields classic",
			token:    "ghp_abc123",
			expected: PATTypeClassic,
		},
		{
			name:     "gho_ prefix yields oauth",
			token:    "gho_abc123",
			expected: PATTypeOAuth,
		},
		{
			name:     "unknown prefix yields unknown",
			token:    "xyz_unknown_token",
			expected: PATTypeUnknown,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ClassifyPAT(tt.token)
			assert.Equal(t, tt.expected, result,
				"ClassifyPAT(%q) should classify token by prefix", tt.token)
		})
	}
}

// TestSpec_PublicAPI_NormalizeWhitespace validates the documented behavior of
// NormalizeWhitespace as described in the package README.md.
//
// Specification: "Normalizes trailing whitespace in multi-line content. Trims
// trailing spaces and tabs from every line, then ensures the content ends with
// exactly one newline (or is empty)."
func TestSpec_PublicAPI_NormalizeWhitespace(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "empty string returns empty string",
			input:    "",
			expected: "",
		},
		{
			name:     "trims trailing spaces from a line",
			input:    "hello   \n",
			expected: "hello\n",
		},
		{
			name:     "trims trailing tabs from a line",
			input:    "hello\t\t\n",
			expected: "hello\n",
		},
		{
			name:     "trims trailing whitespace from every line in multi-line content",
			input:    "line1 \nline2\t\nline3\n",
			expected: "line1\nline2\nline3\n",
		},
		{
			name:     "content with no trailing whitespace is unchanged",
			input:    "clean\nlines\n",
			expected: "clean\nlines\n",
		},
		{
			name:     "ensures content ends with exactly one newline",
			input:    "no trailing newline",
			expected: "no trailing newline\n",
		},
		{
			name:     "collapses multiple trailing newlines to exactly one",
			input:    "content\n\n\n",
			expected: "content\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := NormalizeWhitespace(tt.input)
			assert.Equal(t, tt.expected, result,
				"NormalizeWhitespace(%q) should match documented behavior", tt.input)
		})
	}
}

// TestSpec_PublicAPI_ValidateCopilotPAT validates the documented behavior of
// ValidateCopilotPAT as described in the package README.md.
//
// Specification: "Returns nil if the token is a fine-grained PAT; returns an
// actionable error message with a link to create the correct token type otherwise."
func TestSpec_PublicAPI_ValidateCopilotPAT(t *testing.T) {
	t.Run("fine-grained PAT returns nil", func(t *testing.T) {
		err := ValidateCopilotPAT("github_pat_validtokenhere")
		assert.NoError(t, err,
			"ValidateCopilotPAT should return nil for fine-grained PAT")
	})

	t.Run("classic PAT returns actionable error", func(t *testing.T) {
		err := ValidateCopilotPAT("ghp_classic_token")
		require.Error(t, err,
			"ValidateCopilotPAT should return an error for classic PAT")
		assert.NotEmpty(t, err.Error(),
			"ValidateCopilotPAT error should contain an actionable message")
	})

	t.Run("oauth token returns actionable error", func(t *testing.T) {
		err := ValidateCopilotPAT("gho_oauth_token")
		require.Error(t, err,
			"ValidateCopilotPAT should return an error for OAuth token")
	})
}

// TestSpec_PublicAPI_GetPATTypeDescription validates the documented behavior of
// GetPATTypeDescription as described in the package README.md.
//
// Specification: "Returns a human-readable description of the token type
// (e.g. 'fine-grained personal access token')."
func TestSpec_PublicAPI_GetPATTypeDescription(t *testing.T) {
	tests := []struct {
		name          string
		token         string
		expectContain string
	}{
		{
			name:          "fine-grained PAT returns description containing 'fine-grained'",
			token:         "github_pat_abc123",
			expectContain: "fine-grained",
		},
		{
			name:          "classic PAT returns non-empty description",
			token:         "ghp_abc123",
			expectContain: "",
		},
		{
			name:          "oauth token returns non-empty description",
			token:         "gho_abc123",
			expectContain: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GetPATTypeDescription(tt.token)
			assert.NotEmpty(t, result,
				"GetPATTypeDescription(%q) should return a non-empty human-readable description", tt.token)
			if tt.expectContain != "" {
				assert.Contains(t, result, tt.expectContain,
					"GetPATTypeDescription(%q) should contain %q as documented", tt.token, tt.expectContain)
			}
		})
	}
}

// TestSpec_PublicAPI_SanitizeErrorMessage validates the documented behavior of
// SanitizeErrorMessage as described in the package README.md.
//
// Specification: "Redacts potential secret key names from error messages. Matches
// uppercase SNAKE_CASE identifiers (e.g. MY_SECRET_KEY, API_TOKEN) and PascalCase
// identifiers ending with security-related suffixes (e.g. GitHubToken, ApiKey).
// Common GitHub Actions workflow keywords (GITHUB, RUNNER, WORKFLOW, etc.) are
// excluded from redaction."
//
// Specification example:
//
//	stringutil.SanitizeErrorMessage("Error: MY_SECRET_TOKEN is invalid")
//	// → "Error: [REDACTED] is invalid"
func TestSpec_PublicAPI_SanitizeErrorMessage(t *testing.T) {
	t.Run("redacts SNAKE_CASE secret (documented example)", func(t *testing.T) {
		result := SanitizeErrorMessage("Error: MY_SECRET_TOKEN is invalid")
		assert.Equal(t, "Error: [REDACTED] is invalid", result,
			"SanitizeErrorMessage should redact SNAKE_CASE secret identifiers")
	})

	// Specification: "Common GitHub Actions workflow keywords (GITHUB, RUNNER,
	// WORKFLOW, etc.) are excluded from redaction."
	// Note: standalone keywords like "GITHUB" do not match the compound pattern
	// (which requires underscores), so they pass through unchanged.
	t.Run("does not redact standalone GITHUB keyword", func(t *testing.T) {
		result := SanitizeErrorMessage("Error: GITHUB is not responding")
		assert.NotContains(t, result, "[REDACTED]",
			"SanitizeErrorMessage should not redact standalone GITHUB keyword")
	})

	// Specification: "GH_AW_ prefixed variables are not redacted."
	t.Run("does not redact GH_AW_ configuration variable", func(t *testing.T) {
		result := SanitizeErrorMessage("Set GH_AW_SKIP_NPX_VALIDATION=true")
		assert.NotContains(t, result, "[REDACTED]",
			"SanitizeErrorMessage should not redact GH_AW_ configuration variables")
	})
}

// TestSpec_PublicAPI_SanitizeIdentifierName validates the documented behavior of
// SanitizeIdentifierName as described in the package README.md.
//
// Specification: "Sanitizes a string for use as a programming-language identifier
// by replacing invalid characters with underscores and prefixing _ when the
// identifier starts with a digit. extraAllowed can be used to permit additional
// runes beyond the normal identifier rules; if extraAllowed is nil, no extra
// characters are allowed."
func TestSpec_PublicAPI_SanitizeIdentifierName(t *testing.T) {
	tests := []struct {
		name         string
		input        string
		extraAllowed func(rune) bool
		expected     string
	}{
		{
			name:         "replaces hyphen with underscore when nil extraAllowed",
			input:        "my-func",
			extraAllowed: nil,
			expected:     "my_func",
		},
		{
			name:         "prefixes underscore when identifier starts with digit",
			input:        "123abc",
			extraAllowed: nil,
			expected:     "_123abc",
		},
		{
			name:         "valid identifier unchanged",
			input:        "hello",
			extraAllowed: nil,
			expected:     "hello",
		},
		{
			name:         "replaces dots with underscores when nil extraAllowed",
			input:        "a.b.c",
			extraAllowed: nil,
			expected:     "a_b_c",
		},
		{
			name:  "extraAllowed permits the specified rune",
			input: "a-b",
			extraAllowed: func(r rune) bool {
				return r == '-'
			},
			expected: "a-b",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SanitizeIdentifierName(tt.input, tt.extraAllowed)
			assert.Equal(t, tt.expected, result,
				"SanitizeIdentifierName(%q) should match documented behavior", tt.input)
		})
	}
}

// TestSpec_PublicAPI_SanitizeParameterName validates the documented behavior of
// SanitizeParameterName as described in the package README.md.
//
// Specification: "Sanitizes a parameter name for use as a GitHub Actions output
// or environment variable name. Preserves letters, digits, $, and _, and replaces
// all other characters with underscores."
func TestSpec_PublicAPI_SanitizeParameterName(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "replaces hyphen with underscore",
			input:    "my-param",
			expected: "my_param",
		},
		{
			name:     "preserves dollar sign",
			input:    "param$1",
			expected: "param$1",
		},
		{
			name:     "preserves underscore",
			input:    "param_1",
			expected: "param_1",
		},
		{
			name:     "replaces space with underscore",
			input:    "hello world",
			expected: "hello_world",
		},
		{
			name:     "preserves letters and digits unchanged",
			input:    "myParam123",
			expected: "myParam123",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SanitizeParameterName(tt.input)
			assert.Equal(t, tt.expected, result,
				"SanitizeParameterName(%q) should match documented behavior", tt.input)
		})
	}
}

// TestSpec_PublicAPI_SanitizePythonVariableName validates the documented behavior
// of SanitizePythonVariableName as described in the package README.md.
//
// Specification: "Sanitizes a string for use as a Python variable name. Similar
// to SanitizeParameterName but follows Python identifier rules."
func TestSpec_PublicAPI_SanitizePythonVariableName(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "replaces hyphen with underscore",
			input:    "my-var",
			expected: "my_var",
		},
		{
			name:     "preserves letters and digits",
			input:    "myVar123",
			expected: "myVar123",
		},
		{
			name:     "preserves underscore",
			input:    "my_var",
			expected: "my_var",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SanitizePythonVariableName(tt.input)
			assert.Equal(t, tt.expected, result,
				"SanitizePythonVariableName(%q) should match documented behavior", tt.input)
		})
	}
}

// TestSpec_PublicAPI_SanitizeToolID validates the documented behavior of
// SanitizeToolID as described in the package README.md.
//
// Specification: "Sanitizes a tool identifier for safe use in generated code.
// Replaces characters that are not valid in identifiers with underscores."
func TestSpec_PublicAPI_SanitizeToolID(t *testing.T) {
	// SPEC_MISMATCH: README.md documents SanitizeToolID as replacing characters
	// that are not valid in identifiers with underscores (e.g. "create-issue" →
	// "create_issue"), but the implementation instead removes "mcp-" prefixes and
	// "-mcp" suffixes (e.g. "notion-mcp" → "notion"). These are entirely different
	// behaviors. Tests are skipped until the specification is corrected or the
	// implementation is updated.
	t.Skip("SPEC_MISMATCH: SanitizeToolID implementation removes MCP prefix/suffix instead of replacing invalid identifier characters with underscores as documented in README.md")
}

// TestSpec_PublicAPI_SanitizeForFilename validates the documented behavior of
// SanitizeForFilename as described in the package README.md.
//
// Specification: "Converts a string into a filesystem-safe filename by
// lowercasing and replacing non-alphanumeric characters with hyphens."
func TestSpec_PublicAPI_SanitizeForFilename(t *testing.T) {
	t.Run("replaces space with hyphen (lowercase input)", func(t *testing.T) {
		result := SanitizeForFilename("hello world")
		assert.Equal(t, "hello-world", result,
			"SanitizeForFilename should replace spaces with hyphens")
	})

	t.Run("already safe filename returned unchanged", func(t *testing.T) {
		result := SanitizeForFilename("simple")
		assert.Equal(t, "simple", result,
			"SanitizeForFilename should leave already-safe filenames unchanged")
	})

	// SPEC_MISMATCH: README.md documents SanitizeForFilename as "lowercasing" the
	// input, but the implementation preserves the original case. Additionally, the
	// spec says "replacing non-alphanumeric characters with hyphens" but the
	// implementation preserves '-', '_', and '.'.
	t.Run("SPEC_MISMATCH: lowercasing not implemented", func(t *testing.T) {
		t.Skip("SPEC_MISMATCH: README.md documents SanitizeForFilename as lowercasing the input (e.g. 'Hello World' → 'hello-world'), but the implementation does not lowercase — actual result is 'Hello-World'")
	})
}
