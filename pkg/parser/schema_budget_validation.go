package parser

import (
	"fmt"
	"strings"

	"github.com/github/gh-aw/pkg/typeutil"
)

func validateLegacyBudgetMigration(frontmatter map[string]any) error {
	legacyValue, hasLegacy := frontmatter["max-effective-tokens"]
	_, hasModern := frontmatter["max-ai-credits"]

	if hasLegacy && hasModern {
		return fmt.Errorf("'max-effective-tokens' cannot be used with 'max-ai-credits' in the same workflow; remove the legacy field and keep 'max-ai-credits'")
	}

	if !hasLegacy {
		return nil
	}

	if legacyExpressionValue(legacyValue) {
		return fmt.Errorf("'max-effective-tokens' no longer supports expressions because legacy ET budgets cannot be converted automatically; replace it with 'max-ai-credits'")
	}

	if _, ok := typeutil.ParseLegacyEffectiveTokensAsAICredits(legacyValue); ok {
		return nil
	}

	return fmt.Errorf("'max-effective-tokens' value cannot be converted to 'max-ai-credits'; use a numeric value of at least 10000 ET, use -1 to disable, or replace it with 'max-ai-credits'")
}

func legacyExpressionValue(raw any) bool {
	rawStr, ok := raw.(string)
	if !ok {
		return false
	}
	trimmed := strings.TrimSpace(rawStr)
	return strings.HasPrefix(trimmed, "${{") && strings.HasSuffix(trimmed, "}}")
}
