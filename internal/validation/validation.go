package validation

import (
	"math"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var weightFormatRegex = regexp.MustCompile(`^\d+(\.\d+)?$`)

func ValidateAndFormatWeight(input string) (float64, string) {
	normalized := strings.TrimSpace(input)
	normalized = strings.ReplaceAll(normalized, ",", ".")

	if !weightFormatRegex.MatchString(normalized) {
		return 0, ErrWeightInvalidFormat
	}

	parsed, err := strconv.ParseFloat(normalized, 64)
	if err != nil {
		return 0, ErrWeightFailedToParse
	}

	parsed = math.Round(parsed*100) / 100

	if parsed < 10 {
		return 0, ErrWeightTooLow
	}
	if parsed > 200 {
		return 0, ErrWeightTooHigh
	}

	return parsed, ""
}

func ValidateTimestamp(input string) (time.Time, string) {
	parsed, err := time.Parse(time.RFC3339, input)
	if err != nil {
		return time.Time{}, ErrTimestampNotDate
	}

	return parsed, ""
}
