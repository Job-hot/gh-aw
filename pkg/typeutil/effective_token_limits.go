package typeutil

import (
	"math"
	"strconv"
	"strings"

	"github.com/github/gh-aw/pkg/constants"
)

// ParseInt64KMSuffix parses a positive base-10 integer string with an optional
// K/k (×1,000) or M/m (×1,000,000) suffix.
func ParseInt64KMSuffix(raw string) (int64, bool) {
	trimmed := strings.TrimSpace(raw)
	if trimmed == "" {
		return 0, false
	}

	multiplier := int64(1)
	switch last := trimmed[len(trimmed)-1]; last {
	case 'k', 'K':
		multiplier = 1_000
		trimmed = trimmed[:len(trimmed)-1]
	case 'm', 'M':
		multiplier = 1_000_000
		trimmed = trimmed[:len(trimmed)-1]
	}

	if trimmed == "" {
		return 0, false
	}

	parsed, err := strconv.ParseInt(trimmed, 10, 64)
	if err != nil || parsed <= 0 {
		typeutilLog.Printf("Rejected K/M-suffixed value %q: not a positive base-10 integer", raw)
		return 0, false
	}
	if parsed > math.MaxInt64/multiplier {
		typeutilLog.Printf("Rejected K/M-suffixed value %q: would overflow int64 (multiplier=%d)", raw, multiplier)
		return 0, false
	}
	return parsed * multiplier, true
}

// NormalizeInt64KMSuffix returns a canonical base-10 string for a positive
// integer string with an optional K/k or M/m suffix.
func NormalizeInt64KMSuffix(raw string) (string, bool) {
	parsed, ok := ParseInt64KMSuffix(raw)
	if !ok {
		return "", false
	}
	return strconv.FormatInt(parsed, 10), true
}

// ConvertLegacyEffectiveTokensToAICredits converts a legacy effective-token
// budget to its AI-credits counterpart. Positive values are floor-divided by
// the ET→AI-credit ratio, while -1 is preserved as the disable sentinel.
func ConvertLegacyEffectiveTokensToAICredits(limit int64) (int64, bool) {
	if limit == -1 {
		return -1, true
	}
	if limit <= 0 {
		return 0, false
	}

	aiCredits := limit / constants.EffectiveTokensPerAICredit
	if aiCredits <= 0 {
		return 0, false
	}
	return aiCredits, true
}

// ParseLegacyEffectiveTokensAsAICredits parses a legacy ET budget from either a
// numeric value or a numeric string with optional K/M suffix and converts it to
// AI credits.
func ParseLegacyEffectiveTokensAsAICredits(raw any) (int64, bool) {
	if parsed, ok := ParseIntValue(raw); ok && parsed != 0 {
		return ConvertLegacyEffectiveTokensToAICredits(int64(parsed))
	}

	rawStr, ok := raw.(string)
	if !ok {
		return 0, false
	}

	trimmed := strings.TrimSpace(rawStr)
	if trimmed == "-1" {
		return -1, true
	}

	normalized, ok := NormalizeInt64KMSuffix(trimmed)
	if !ok {
		return 0, false
	}

	parsed, err := strconv.ParseInt(normalized, 10, 64)
	if err != nil {
		return 0, false
	}
	return ConvertLegacyEffectiveTokensToAICredits(parsed)
}
