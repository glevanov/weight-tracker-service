package validation

import (
	"testing"
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
				if err == "" {
					t.Errorf("expected error for input %q, got none", tt.input)
				}
			} else {
				if err != "" {
					t.Errorf("unexpected error for input %q: %s", tt.input, err)
				}
				if result != tt.expected {
					t.Errorf("ValidateAndFormatWeight(%q) = %v, want %v", tt.input, result, tt.expected)
				}
			}
		})
	}
}
