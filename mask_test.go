package xstr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaskSensitive(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "empty string",
			input:    "",
			expected: "",
		},
		{
			name:     "single character",
			input:    "a",
			expected: "****",
		},
		{
			name:     "two characters",
			input:    "ab",
			expected: "****",
		},
		{
			name:     "four characters",
			input:    "abcd",
			expected: "****",
		},
		{
			name:     "five characters",
			input:    "abcde",
			expected: "****de",
		},
		{
			name:     "eight characters",
			input:    "abcdefgh",
			expected: "****gh",
		},
		{
			name:     "nine characters",
			input:    "abcdefgh9",
			expected: "abcd****fgh9",
		},
		{
			name:     "long api key",
			input:    "pk_live_1234567890abcde9",
			expected: "pk_l****cde9",
		},
		{
			name:     "jwt token",
			input:    "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVC9",
			expected: "eyJh****XVC9",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MaskSensitive(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestMaskEmail(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "empty string",
			input:    "",
			expected: "",
		},
		{
			name:     "standard email",
			input:    "user@example.com",
			expected: "u***@example.com",
		},
		{
			name:     "short local part",
			input:    "a@example.com",
			expected: "a***@example.com",
		},
		{
			name:     "single char local part",
			input:    "u@example.com",
			expected: "u***@example.com",
		},
		{
			name:     "long email",
			input:    "very.long.email.address@example.com",
			expected: "v****s@example.com",
		},
		{
			name:     "no at symbol",
			input:    "not9ane1mail",
			expected: "not9****mail",
		},
		{
			name:     "empty local part",
			input:    "@example.com",
			expected: "****@example.com",
		},
		{
			name:     "multiple at symbols",
			input:    "user@@example.com",
			expected: "u***@@example.com",
		},
		{
			name:     "four char local part",
			input:    "test@example.com",
			expected: "t***@example.com",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MaskEmail(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestMaskPhone(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "empty string",
			input:    "",
			expected: "",
		},
		{
			name:     "thai mobile number",
			input:    "+66812345678",
			expected: "+66****5678",
		},
		{
			name:     "us mobile number",
			input:    "+15551234567",
			expected: "+15****4567",
		},
		{
			name:     "short international number",
			input:    "+661234",
			expected: "+66****",
		},
		{
			name:     "very short international number",
			input:    "+66123",
			expected: "+66****",
		},
		{
			name:     "local number long",
			input:    "0812345678",
			expected: "****5678",
		},
		{
			name:     "local number short",
			input:    "123456",
			expected: "****",
		},
		{
			name:     "very short local number",
			input:    "12345",
			expected: "****",
		},
		{
			name:     "single digit",
			input:    "1",
			expected: "****",
		},
		{
			name:     "landline number",
			input:    "021234567",
			expected: "****4567",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MaskPhone(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestMaskSensitive_APIKeys(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "stripe api key",
			input:    "pk_live_1234567890abcde9",
			expected: "pk_l****cde9",
		},
		{
			name:     "short key",
			input:    "key123",
			expected: "****23",
		},
		{
			name:     "very short key",
			input:    "abc",
			expected: "****",
		},
		{
			name:     "empty key",
			input:    "",
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MaskSensitive(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestMaskSensitive_Tokens(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "jwt token",
			input:    "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpVCJ9",
			expected: "eyJh****VCJ9",
		},
		{
			name:     "bearer token",
			input:    "Bearer abc123def456",
			expected: "Bear****f456",
		},
		{
			name:     "empty token",
			input:    "",
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MaskSensitive(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestMaskSensitive_Signatures(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "hmac signature",
			input:    "a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6",
			expected: "a1b2****o5p6",
		},
		{
			name:     "short signature",
			input:    "abc123",
			expected: "****23",
		},
		{
			name:     "empty signature",
			input:    "",
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MaskSensitive(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestMaskSensitive_EdgeCases(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "special characters",
			input:    "!@#$%^&*()",
			expected: "!@#$****&*()",
		},
		{
			name:     "mixed alphanumeric",
			input:    "abc123DEF",
			expected: "abc1****3DEF",
		},
		{
			name:     "spaces in string",
			input:    "hello world test",
			expected: "hell****test",
		},
		{
			name:     "six characters",
			input:    "abc123",
			expected: "****23",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MaskSensitive(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestMaskEmail_EdgeCases(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "subdomain email",
			input:    "user@mail.example.com",
			expected: "u***@mail.example.com",
		},
		{
			name:     "plus addressing",
			input:    "user+tag@example.com",
			expected: "u****g@example.com",
		},
		{
			name:     "dots in local part",
			input:    "first.last@example.com",
			expected: "f****t@example.com",
		},
		{
			name:     "numbers in email",
			input:    "user123@example.com",
			expected: "u****3@example.com",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MaskEmail(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestMaskPhone_EdgeCases(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "country code with spaces",
			input:    "+66 81 234 5678",
			expected: "+66****5678",
		},
		{
			name:     "country code with dashes",
			input:    "+66-81-234-5678",
			expected: "+66****5678",
		},
		{
			name:     "extension number",
			input:    "+6681234567890",
			expected: "+66****7890",
		},
		{
			name:     "local with parentheses",
			input:    "(081)234-5678",
			expected: "****5678",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MaskPhone(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}
