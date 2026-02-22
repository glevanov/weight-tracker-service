package validation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateAndFormatWeight(t *testing.T) {
	tests := []struct {
		input    string
		expected float64
		wantErr  bool
	}{
		{"82", 82, false},
		{"82,2", 82.2, false},
		{"82.2", 82.2, false},
		{"83,40", 83.4, false},
		{"83.40", 83.4, false},
		{"83,405", 83.41, false},
		{"83.405", 83.41, false},
		{"83,404", 83.4, false},
		{"83.404", 83.4, false},
		{"103", 103, false},
		{"103.4", 103.4, false},
		{"103.405", 103.41, false},
		{"199.993", 199.99, false},
		{"5", 0, true},
		{"9", 0, true},
		{"500", 0, true},
		{"250", 0, true},
		{"44.1.1", 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result, err := ValidateAndFormatWeight(tt.input)
			if tt.wantErr {
				assert.NotEmpty(t, err, "expected error for input %q", tt.input)
			} else {
				assert.Empty(t, err, "unexpected error for input %q: %s", tt.input, err)
				assert.Equal(t, tt.expected, result)
			}
		})
	}
}

func TestValidateTimestamp(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr bool
	}{
		{"valid ISO 8601", "2024-01-15T10:30:00Z", false},
		{"valid with timezone", "2024-01-15T10:30:00+01:00", false},
		{"valid with milliseconds", "2024-01-15T10:30:00.123Z", false},
		{"invalid format", "2024-01-15", true},
		{"invalid format - slash date", "2024/01/15", true},
		{"empty string", "", true},
		{"not a date", "not-a-timestamp", true},
		{"partial timestamp", "2024-01-15T10:30", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ValidateTimestamp(tt.input)
			if tt.wantErr {
				assert.NotEmpty(t, err, "expected error for input %q", tt.input)
			} else {
				assert.Empty(t, err, "unexpected error for input %q: %s", tt.input, err)
				assert.False(t, result.IsZero(), "expected non-zero time for valid input %q", tt.input)
			}
		})
	}
}
