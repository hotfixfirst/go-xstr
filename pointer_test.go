package xstr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNormalizeOptionalString(t *testing.T) {
	tests := []struct {
		name     string
		input    *string
		expected *string
	}{
		{
			name:     "nil input returns nil",
			input:    nil,
			expected: nil,
		},
		{
			name:     "empty string returns nil",
			input:    stringPtr(""),
			expected: nil,
		},
		{
			name:     "whitespace only returns nil",
			input:    stringPtr("   "),
			expected: nil,
		},
		{
			name:     "tab and newline returns nil",
			input:    stringPtr("\t\n  \n\t"),
			expected: nil,
		},
		{
			name:     "normal string returns trimmed",
			input:    stringPtr("hello"),
			expected: stringPtr("hello"),
		},
		{
			name:     "string with leading spaces returns trimmed",
			input:    stringPtr("  hello"),
			expected: stringPtr("hello"),
		},
		{
			name:     "string with trailing spaces returns trimmed",
			input:    stringPtr("hello  "),
			expected: stringPtr("hello"),
		},
		{
			name:     "string with both leading and trailing spaces returns trimmed",
			input:    stringPtr("  hello  "),
			expected: stringPtr("hello"),
		},
		{
			name:     "string with tabs returns trimmed",
			input:    stringPtr("\thello\t"),
			expected: stringPtr("hello"),
		},
		{
			name:     "string with mixed whitespace returns trimmed",
			input:    stringPtr(" \t hello world \n "),
			expected: stringPtr("hello world"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := NormalizeOptionalString(tt.input)

			if tt.expected == nil {
				assert.Nil(t, result)
			} else {
				assert.NotNil(t, result)
				assert.Equal(t, *tt.expected, *result)
			}
		})
	}
}

func TestNormalizeUpdateString(t *testing.T) {
	tests := []struct {
		name     string
		input    *string
		expected *string
	}{
		{
			name:     "nil input returns nil",
			input:    nil,
			expected: nil,
		},
		{
			name:     "empty string returns empty string pointer",
			input:    stringPtr(""),
			expected: stringPtr(""),
		},
		{
			name:     "whitespace only returns empty string pointer",
			input:    stringPtr("   "),
			expected: stringPtr(""),
		},
		{
			name:     "tab and newline returns empty string pointer",
			input:    stringPtr("\t\n  \n\t"),
			expected: stringPtr(""),
		},
		{
			name:     "normal string returns trimmed",
			input:    stringPtr("hello"),
			expected: stringPtr("hello"),
		},
		{
			name:     "string with leading spaces returns trimmed",
			input:    stringPtr("  hello"),
			expected: stringPtr("hello"),
		},
		{
			name:     "string with trailing spaces returns trimmed",
			input:    stringPtr("hello  "),
			expected: stringPtr("hello"),
		},
		{
			name:     "string with both leading and trailing spaces returns trimmed",
			input:    stringPtr("  hello  "),
			expected: stringPtr("hello"),
		},
		{
			name:     "string with tabs returns trimmed",
			input:    stringPtr("\thello\t"),
			expected: stringPtr("hello"),
		},
		{
			name:     "string with mixed whitespace returns trimmed",
			input:    stringPtr(" \t hello world \n "),
			expected: stringPtr("hello world"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := NormalizeUpdateString(tt.input)

			if tt.expected == nil {
				assert.Nil(t, result)
			} else {
				assert.NotNil(t, result)
				assert.Equal(t, *tt.expected, *result)
			}
		})
	}
}
