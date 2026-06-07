package parser

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/github/gh-aw/pkg/constants"
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

	if _, ok := convertLegacyEffectiveTokensToAICredits(legacyValue); ok {
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

func convertLegacyEffectiveTokensToAICredits(raw any) (int64, bool) {
	if parsed, ok := typeutil.ParseIntValue(raw); ok && parsed != 0 {
		return convertLegacyEffectiveTokensInt64(int64(parsed))
	}

	rawStr, ok := raw.(string)
	if !ok {
		return 0, false
	}

	trimmed := strings.TrimSpace(rawStr)
	if trimmed == "-1" {
		return -1, true
	}

	normalized, ok := typeutil.NormalizeInt64KMSuffix(trimmed)
	if !ok {
		return 0, false
	}

	parsed, err := strconv.ParseInt(normalized, 10, 64)
	if err != nil {
		return 0, false
	}
	return convertLegacyEffectiveTokensInt64(parsed)
}

func convertLegacyEffectiveTokensInt64(value int64) (int64, bool) {
	if value == -1 {
		return -1, true
	}
	if value <= 0 {
		return 0, false
	}
	aiCredits := value / constants.EffectiveTokensPerAICredit
	if aiCredits <= 0 {
		return 0, false
	}
	return aiCredits, true
}
