package xstr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvertPhoneByCurrency(t *testing.T) {
	tests := []struct {
		name          string
		phone         string
		currency      string
		expected      string
		expectError   bool
		expectedError error
	}{
		{
			name:        "valid Thai phone with THB",
			phone:       "+66812345678",
			currency:    "THB",
			expected:    "0812345678",
			expectError: false,
		},
		{
			name:        "valid Thai phone domestic format with THB",
			phone:       "0812345678",
			currency:    "THB",
			expected:    "0812345678",
			expectError: false,
		},
		{
			name:        "valid US phone with USD",
			phone:       "+12025551234",
			currency:    "USD",
			expected:    "2025551234",
			expectError: false,
		},
		{
			name:        "valid Singapore phone with SGD",
			phone:       "+6591234567",
			currency:    "SGD",
			expected:    "91234567",
			expectError: false,
		},
		{
			name:          "Thai phone with wrong currency",
			phone:         "+66812345678",
			currency:      "USD",
			expected:      "",
			expectError:   true,
			expectedError: ErrPhoneCountryMismatch,
		},
		{
			name:          "US phone with wrong currency",
			phone:         "+12025551234",
			currency:      "THB",
			expected:      "",
			expectError:   true,
			expectedError: ErrPhoneCountryMismatch,
		},
		{
			name:          "invalid currency",
			phone:         "+66812345678",
			currency:      "XYZ",
			expected:      "",
			expectError:   true,
			expectedError: ErrCurrencyNotSupported,
		},
		{
			name:          "invalid phone format",
			phone:         "invalid-phone",
			currency:      "THB",
			expected:      "",
			expectError:   true,
			expectedError: ErrInvalidPhoneFormat,
		},
		{
			name:        "case insensitive currency",
			phone:       "+66812345678",
			currency:    "thb",
			expected:    "0812345678",
			expectError: false,
		},
		{
			name:        "phone with spaces and dashes",
			phone:       "+66 81-234-5678",
			currency:    "THB",
			expected:    "0812345678",
			expectError: false,
		},
		{
			name:        "different Thai mobile prefix 06",
			phone:       "+66612345678",
			currency:    "THB",
			expected:    "0612345678",
			expectError: false,
		},
		{
			name:        "different Thai mobile prefix 09",
			phone:       "+66912345678",
			currency:    "THB",
			expected:    "0912345678",
			expectError: false,
		},
		{
			name:        "Thai phone international without plus",
			phone:       "66812345678",
			currency:    "THB",
			expected:    "0812345678",
			expectError: false,
		},
		{
			name:        "Thai phone 9 digits",
			phone:       "812345678",
			currency:    "THB",
			expected:    "0812345678",
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ConvertPhoneByCurrency(tt.phone, tt.currency)

			if tt.expectError {
				assert.Error(t, err)
				if tt.expectedError != nil {
					assert.Equal(t, tt.expectedError, err)
				}
				assert.Equal(t, tt.expected, result)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, result)
			}
		})
	}
}

func TestConvertPhoneByCurrencyToFormat(t *testing.T) {
	tests := []struct {
		name        string
		phone       string
		currency    string
		format      PhoneFormat
		expected    string
		expectError bool
	}{
		{
			name:        "Thai phone to domestic format",
			phone:       "+66812345678",
			currency:    "THB",
			format:      PhoneFormatDomestic,
			expected:    "0812345678",
			expectError: false,
		},
		{
			name:        "Thai phone to E164 format",
			phone:       "0812345678",
			currency:    "THB",
			format:      PhoneFormatE164,
			expected:    "+66812345678",
			expectError: false,
		},
		{
			name:        "Thai phone to domestic dashed format",
			phone:       "+66812345678",
			currency:    "THB",
			format:      PhoneFormatDomesticDashed,
			expected:    "081-234-5678",
			expectError: false,
		},
		{
			name:        "Thai phone to E164 dashed format",
			phone:       "+66812345678",
			currency:    "THB",
			format:      PhoneFormatE164Dashed,
			expected:    "+66-81-234-5678",
			expectError: false,
		},
		{
			name:        "US phone to domestic format",
			phone:       "+12025551234",
			currency:    "USD",
			format:      PhoneFormatDomestic,
			expected:    "2025551234",
			expectError: false,
		},
		{
			name:        "Singapore phone to domestic format",
			phone:       "+6591234567",
			currency:    "SGD",
			format:      PhoneFormatDomestic,
			expected:    "91234567",
			expectError: false,
		},
		{
			name:        "Singapore phone to E164 dashed format",
			phone:       "+6591234567",
			currency:    "SGD",
			format:      PhoneFormatE164Dashed,
			expected:    "+65-9123-4567",
			expectError: false,
		},
		{
			name:        "Thai phone with wrong currency should error",
			phone:       "+66812345678",
			currency:    "USD",
			format:      PhoneFormatDomestic,
			expected:    "",
			expectError: true,
		},
		{
			name:        "invalid currency should error",
			phone:       "+66812345678",
			currency:    "XYZ",
			format:      PhoneFormatDomestic,
			expected:    "",
			expectError: true,
		},
		{
			name:        "invalid phone format should error",
			phone:       "invalid",
			currency:    "THB",
			format:      PhoneFormatDomestic,
			expected:    "",
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ConvertPhoneByCurrencyToFormat(tt.phone, tt.currency, tt.format)

			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, result)
			}
		})
	}
}

func TestValidatePhoneCurrency(t *testing.T) {
	tests := []struct {
		name          string
		phone         string
		currency      string
		expectedError error
	}{
		{
			name:          "valid Thai phone with THB",
			phone:         "+66812345678",
			currency:      "THB",
			expectedError: nil,
		},
		{
			name:          "valid US phone with USD",
			phone:         "+12025551234",
			currency:      "USD",
			expectedError: nil,
		},
		{
			name:          "valid Singapore phone with SGD",
			phone:         "+6591234567",
			currency:      "SGD",
			expectedError: nil,
		},
		{
			name:          "Thai phone with wrong currency",
			phone:         "+66812345678",
			currency:      "USD",
			expectedError: ErrPhoneCountryMismatch,
		},
		{
			name:          "US phone with wrong currency",
			phone:         "+12025551234",
			currency:      "THB",
			expectedError: ErrPhoneCountryMismatch,
		},
		{
			name:          "invalid currency",
			phone:         "+66812345678",
			currency:      "XYZ",
			expectedError: ErrCurrencyNotSupported,
		},
		{
			name:          "invalid phone format",
			phone:         "invalid-phone",
			currency:      "THB",
			expectedError: ErrInvalidPhoneFormat,
		},
		{
			name:          "case insensitive currency - valid",
			phone:         "+66812345678",
			currency:      "thb",
			expectedError: nil,
		},
		{
			name:          "domestic format phone - valid",
			phone:         "0812345678",
			currency:      "THB",
			expectedError: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidatePhoneCurrency(tt.phone, tt.currency)

			if tt.expectedError != nil {
				assert.Error(t, err)
				assert.Equal(t, tt.expectedError, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestNormalizePhoneToE164(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		expected    string
		expectError bool
	}{
		{
			name:        "E164 format already",
			input:       "+66812345678",
			expected:    "+66812345678",
			expectError: false,
		},
		{
			name:        "international without plus",
			input:       "66812345678",
			expected:    "+66812345678",
			expectError: false,
		},
		{
			name:        "Thai domestic format",
			input:       "0812345678",
			expected:    "+66812345678",
			expectError: false,
		},
		{
			name:        "9 digits Thai mobile",
			input:       "812345678",
			expected:    "+66812345678",
			expectError: false,
		},
		{
			name:        "with spaces",
			input:       "+66 81 234 5678",
			expected:    "+66812345678",
			expectError: false,
		},
		{
			name:        "with dashes",
			input:       "+66-81-234-5678",
			expected:    "+66812345678",
			expectError: false,
		},
		{
			name:        "with mixed separators",
			input:       "+66 81-234 5678",
			expected:    "+66812345678",
			expectError: false,
		},
		{
			name:        "different Thai mobile prefix 06",
			input:       "+66612345678",
			expected:    "+66612345678",
			expectError: false,
		},
		{
			name:        "different Thai mobile prefix 09",
			input:       "0912345678",
			expected:    "+66912345678",
			expectError: false,
		},
		{
			name:        "empty string",
			input:       "",
			expected:    "",
			expectError: true,
		},
		{
			name:        "invalid format",
			input:       "invalid",
			expected:    "",
			expectError: true,
		},
		{
			name:        "too short",
			input:       "123456",
			expected:    "",
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := NormalizePhoneToE164(tt.input)

			if tt.expectError {
				assert.Error(t, err)
				assert.Equal(t, tt.expected, result)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, result)
			}
		})
	}
}

func TestConvertPhoneFormat(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		format      PhoneFormat
		expected    string
		expectError bool
	}{
		{
			name:        "E164 to domestic - Thai",
			input:       "+66812345678",
			format:      PhoneFormatDomestic,
			expected:    "0812345678",
			expectError: false,
		},
		{
			name:        "E164 to E164 - no change",
			input:       "+66812345678",
			format:      PhoneFormatE164,
			expected:    "+66812345678",
			expectError: false,
		},
		{
			name:        "E164 to domestic dashed - Thai",
			input:       "+66812345678",
			format:      PhoneFormatDomesticDashed,
			expected:    "081-234-5678",
			expectError: false,
		},
		{
			name:        "E164 to E164 dashed - Thai",
			input:       "+66812345678",
			format:      PhoneFormatE164Dashed,
			expected:    "+66-81-234-5678",
			expectError: false,
		},
		{
			name:        "E164 to domestic - US",
			input:       "+12025551234",
			format:      PhoneFormatDomestic,
			expected:    "2025551234",
			expectError: false,
		},
		{
			name:        "E164 to domestic - Singapore",
			input:       "+6591234567",
			format:      PhoneFormatDomestic,
			expected:    "91234567",
			expectError: false,
		},
		{
			name:        "invalid E164 format",
			input:       "invalid",
			format:      PhoneFormatDomestic,
			expected:    "",
			expectError: true,
		},
		{
			name:        "domestic format input should error",
			input:       "0812345678",
			format:      PhoneFormatDomestic,
			expected:    "",
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ConvertPhoneFormat(tt.input, tt.format)

			if tt.expectError {
				assert.Error(t, err)
				assert.Equal(t, tt.expected, result)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, result)
			}
		})
	}
}

func TestGetPhoneCountryCode(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		expected    string
		expectError bool
	}{
		{
			name:        "Thai phone number",
			input:       "+66812345678",
			expected:    "TH",
			expectError: false,
		},
		{
			name:        "US phone number",
			input:       "+12025551234",
			expected:    "US",
			expectError: false,
		},
		{
			name:        "Singapore phone number",
			input:       "+6591234567",
			expected:    "SG",
			expectError: false,
		},
		{
			name:        "invalid E164 format",
			input:       "0812345678",
			expected:    "",
			expectError: true,
		},
		{
			name:        "invalid phone",
			input:       "invalid",
			expected:    "",
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := GetPhoneCountryCode(tt.input)

			if tt.expectError {
				assert.Error(t, err)
				assert.Equal(t, tt.expected, result)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, result)
			}
		})
	}
}

func TestIsMobileNumber(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "valid Thai mobile",
			input:    "+66812345678",
			expected: true,
		},
		{
			name:     "valid US mobile",
			input:    "+12025551234",
			expected: true,
		},
		{
			name:     "valid Singapore mobile",
			input:    "+6591234567",
			expected: true,
		},
		{
			name:     "invalid E164 format",
			input:    "0812345678",
			expected: false,
		},
		{
			name:     "invalid phone",
			input:    "invalid",
			expected: false,
		},
		{
			name:     "empty string",
			input:    "",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsMobileNumber(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestCurrencyToCountryMapping(t *testing.T) {
	tests := []struct {
		currency        string
		expectedCountry string
		shouldExist     bool
	}{
		{"THB", "TH", true},
		{"USD", "US", true},
		{"SGD", "SG", true},
		{"MYR", "MY", true},
		{"IDR", "ID", true},
		{"PHP", "PH", true},
		{"VND", "VN", true},
		{"EUR", "DE", true},
		{"GBP", "GB", true},
		{"JPY", "JP", true},
		{"KRW", "KR", true},
		{"CNY", "CN", true},
		{"HKD", "HK", true},
		{"AUD", "AU", true},
		{"CAD", "CA", true},
		{"XYZ", "", false},
		{"INVALID", "", false},
	}

	for _, tt := range tests {
		t.Run("currency_"+tt.currency, func(t *testing.T) {
			country, exists := CurrencyToCountryMapping[tt.currency]

			assert.Equal(t, tt.shouldExist, exists)
			if tt.shouldExist {
				assert.Equal(t, tt.expectedCountry, country)
			}
		})
	}
}
