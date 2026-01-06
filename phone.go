package xstr

import (
	"errors"
	"strings"

	"github.com/nyaruka/phonenumbers"
)

// PhoneFormat represents different phone number formatting options.
type PhoneFormat int

const (
	// PhoneFormatE164 represents international E.164 format (+66812345678).
	PhoneFormatE164 PhoneFormat = iota
	// PhoneFormatDomestic represents domestic format (0812345678).
	PhoneFormatDomestic
	// PhoneFormatE164Dashed represents E.164 format with dashes (+66-81-234-5678).
	PhoneFormatE164Dashed
	// PhoneFormatDomesticDashed represents domestic format with dashes (081-234-5678).
	PhoneFormatDomesticDashed
)

// Common phone validation errors.
var (
	ErrInvalidPhoneFormat   = errors.New("invalid phone number format")
	ErrInvalidCountryCode   = errors.New("invalid country code")
	ErrNotMobileNumber      = errors.New("not a mobile phone number")
	ErrCurrencyNotSupported = errors.New("currency not supported")
	ErrPhoneCountryMismatch = errors.New("phone number country doesn't match currency")
)

// CurrencyToCountryMapping maps currency codes to their corresponding country codes.
var CurrencyToCountryMapping = map[string]string{
	"THB": "TH", // Thai Baht -> Thailand
	"USD": "US", // US Dollar -> United States
	"SGD": "SG", // Singapore Dollar -> Singapore
	"MYR": "MY", // Malaysian Ringgit -> Malaysia
	"IDR": "ID", // Indonesian Rupiah -> Indonesia
	"PHP": "PH", // Philippine Peso -> Philippines
	"VND": "VN", // Vietnamese Dong -> Vietnam
	"EUR": "DE", // Euro -> Germany (default, could be multiple countries)
	"GBP": "GB", // British Pound -> United Kingdom
	"JPY": "JP", // Japanese Yen -> Japan
	"KRW": "KR", // South Korean Won -> South Korea
	"CNY": "CN", // Chinese Yuan -> China
	"HKD": "HK", // Hong Kong Dollar -> Hong Kong
	"AUD": "AU", // Australian Dollar -> Australia
	"CAD": "CA", // Canadian Dollar -> Canada
}

// ConvertPhoneByCurrency converts phone number to domestic format based on currency code.
// This function validates that the phone number matches the currency's country.
//
// Examples:
//   - ConvertPhoneByCurrency("+66812345678", "THB") -> "0812345678"
//   - ConvertPhoneByCurrency("66812345678", "THB") -> "0812345678"
//   - ConvertPhoneByCurrency("+15551234567", "USD") -> "5551234567"
//   - ConvertPhoneByCurrency("+6591234567", "SGD") -> "91234567"
func ConvertPhoneByCurrency(phoneNumber, currencyCode string) (string, error) {
	// Get expected country from currency
	expectedCountry, exists := CurrencyToCountryMapping[strings.ToUpper(currencyCode)]
	if !exists {
		return "", ErrCurrencyNotSupported
	}

	// Normalize phone to E.164 first
	e164Phone, err := NormalizePhoneToE164(phoneNumber)
	if err != nil {
		return "", err
	}

	// Validate phone is mobile number
	if !IsMobileNumber(e164Phone) {
		return "", ErrNotMobileNumber
	}

	// Get phone's country
	phoneCountry, err := GetPhoneCountryCode(e164Phone)
	if err != nil {
		return "", err
	}

	// Validate phone country matches currency country
	if phoneCountry != expectedCountry {
		return "", ErrPhoneCountryMismatch
	}

	// Convert to domestic format
	domestic, err := ConvertPhoneFormat(e164Phone, PhoneFormatDomestic)
	if err != nil {
		return "", err
	}

	return domestic, nil
}

// ConvertPhoneByCurrencyToFormat converts phone number to specified format based on currency code.
// This function validates that the phone number matches the currency's country.
//
// Examples:
//   - ConvertPhoneByCurrencyToFormat("+66812345678", "THB", PhoneFormatDomestic) -> "0812345678"
//   - ConvertPhoneByCurrencyToFormat("+66812345678", "THB", PhoneFormatDomesticDashed) -> "081-234-5678"
//   - ConvertPhoneByCurrencyToFormat("0812345678", "THB", PhoneFormatE164) -> "+66812345678"
func ConvertPhoneByCurrencyToFormat(phoneNumber, currencyCode string, format PhoneFormat) (string, error) {
	// Get expected country from currency
	expectedCountry, exists := CurrencyToCountryMapping[strings.ToUpper(currencyCode)]
	if !exists {
		return "", ErrCurrencyNotSupported
	}

	// Normalize phone to E.164 first
	e164Phone, err := NormalizePhoneToE164(phoneNumber)
	if err != nil {
		return "", err
	}

	// Validate phone is mobile number
	if !IsMobileNumber(e164Phone) {
		return "", ErrNotMobileNumber
	}

	// Get phone's country
	phoneCountry, err := GetPhoneCountryCode(e164Phone)
	if err != nil {
		return "", err
	}

	// Validate phone country matches currency country
	if phoneCountry != expectedCountry {
		return "", ErrPhoneCountryMismatch
	}

	// Convert to specified format
	result, err := ConvertPhoneFormat(e164Phone, format)
	if err != nil {
		return "", err
	}

	return result, nil
}

// ValidatePhoneCurrency validates if phone number is compatible with currency code.
// Returns nil if valid, error if incompatible.
//
// Examples:
//   - ValidatePhoneCurrency("+66812345678", "THB") -> nil (valid)
//   - ValidatePhoneCurrency("+66812345678", "USD") -> ErrPhoneCountryMismatch
//   - ValidatePhoneCurrency("+1234567890", "THB") -> ErrPhoneCountryMismatch
func ValidatePhoneCurrency(phoneNumber, currencyCode string) error {
	// Get expected country from currency
	expectedCountry, exists := CurrencyToCountryMapping[strings.ToUpper(currencyCode)]
	if !exists {
		return ErrCurrencyNotSupported
	}

	// Normalize phone to E.164 first
	e164Phone, err := NormalizePhoneToE164(phoneNumber)
	if err != nil {
		return err
	}

	// Validate phone is mobile number
	if !IsMobileNumber(e164Phone) {
		return ErrNotMobileNumber
	}

	// Get phone's country
	phoneCountry, err := GetPhoneCountryCode(e164Phone)
	if err != nil {
		return err
	}

	// Validate phone country matches currency country
	if phoneCountry != expectedCountry {
		return ErrPhoneCountryMismatch
	}

	return nil
}

// NormalizePhoneToE164 converts any phone format to E.164 format.
// Supports input formats:
//   - +66812345678 (E.164 - returns as is)
//   - 66812345678 (international without +)
//   - 0812345678 (Thai domestic)
//   - 812345678 (9 digits, assumes Thai)
//
// Returns E.164 format (+66812345678) or error if invalid.
func NormalizePhoneToE164(phoneNumber string) (string, error) {
	// Clean input
	phone := cleanPhoneInput(phoneNumber)
	if phone == "" {
		return "", ErrInvalidPhoneFormat
	}

	// If already in E.164 format, validate and return
	if strings.HasPrefix(phone, "+") {
		if isValidE164Format(phone) {
			return phone, nil
		}
		return "", ErrInvalidPhoneFormat
	}

	// Try to parse without + prefix
	if strings.HasPrefix(phone, "66") && len(phone) >= 10 {
		e164 := "+" + phone
		if isValidE164Format(e164) {
			return e164, nil
		}
	}

	// Handle Thai domestic format (0812345678)
	if strings.HasPrefix(phone, "0") && len(phone) == 10 {
		e164 := "+66" + phone[1:]
		if isValidE164Format(e164) {
			return e164, nil
		}
	}

	// Handle 9 digits (assume Thai mobile without country code and leading 0)
	if len(phone) == 9 && isDigits(phone) {
		e164 := "+66" + phone
		if isValidE164Format(e164) {
			return e164, nil
		}
	}

	return "", ErrInvalidPhoneFormat
}

// ConvertPhoneFormat converts phone number from E.164 to specified format.
// Input must be valid E.164 format (+66812345678).
func ConvertPhoneFormat(e164Phone string, format PhoneFormat) (string, error) {
	// Validate input is E.164
	if !isValidE164Format(e164Phone) {
		return "", ErrInvalidPhoneFormat
	}

	// Parse to get region info
	num, err := phonenumbers.Parse(e164Phone, "")
	if err != nil {
		return "", ErrInvalidPhoneFormat
	}

	regionCode := phonenumbers.GetRegionCodeForNumber(num)

	switch format {
	case PhoneFormatE164:
		return e164Phone, nil

	case PhoneFormatDomestic:
		return convertToDomestic(e164Phone, regionCode)

	case PhoneFormatE164Dashed:
		return formatE164WithDashes(e164Phone, regionCode)

	case PhoneFormatDomesticDashed:
		domestic, err := convertToDomestic(e164Phone, regionCode)
		if err != nil {
			return "", err
		}
		return formatDomesticWithDashes(domestic, regionCode)

	default:
		return e164Phone, nil
	}
}

// GetPhoneCountryCode extracts country code from E.164 phone number.
// Returns ISO country code (e.g., "TH", "US", "SG").
func GetPhoneCountryCode(e164Phone string) (string, error) {
	if !isValidE164Format(e164Phone) {
		return "", ErrInvalidPhoneFormat
	}

	num, err := phonenumbers.Parse(e164Phone, "")
	if err != nil {
		return "", err
	}

	return phonenumbers.GetRegionCodeForNumber(num), nil
}

// IsMobileNumber validates if the E.164 phone number is a mobile number.
func IsMobileNumber(e164Phone string) bool {
	if !isValidE164Format(e164Phone) {
		return false
	}

	num, err := phonenumbers.Parse(e164Phone, "")
	if err != nil {
		return false
	}

	numberType := phonenumbers.GetNumberType(num)
	return numberType == phonenumbers.MOBILE || numberType == phonenumbers.FIXED_LINE_OR_MOBILE
}

// Helper functions

// cleanPhoneInput removes spaces, dashes, and other non-essential characters.
func cleanPhoneInput(phone string) string {
	phone = strings.TrimSpace(phone)
	phone = strings.ReplaceAll(phone, " ", "")
	phone = strings.ReplaceAll(phone, "-", "")
	phone = strings.ReplaceAll(phone, "(", "")
	phone = strings.ReplaceAll(phone, ")", "")
	return phone
}

// isValidE164Format validates E.164 format using libphonenumber.
func isValidE164Format(phone string) bool {
	num, err := phonenumbers.Parse(phone, "")
	if err != nil {
		return false
	}
	return phonenumbers.IsValidNumber(num)
}

// isDigits checks if string contains only digits.
func isDigits(s string) bool {
	for _, r := range s {
		if r < '0' || r > '9' {
			return false
		}
	}
	return true
}

// convertToDomestic converts E.164 to domestic format based on country.
func convertToDomestic(e164Phone, regionCode string) (string, error) {
	num, err := phonenumbers.Parse(e164Phone, "")
	if err != nil {
		return "", err
	}

	// Format as national number and clean up
	domestic := phonenumbers.Format(num, phonenumbers.NATIONAL)
	domestic = cleanPhoneInput(domestic)

	// Add leading zero for certain countries if needed
	switch regionCode {
	case "TH":
		if !strings.HasPrefix(domestic, "0") {
			domestic = "0" + domestic
		}
	}

	return domestic, nil
}

// formatE164WithDashes formats E.164 number with dashes for readability.
func formatE164WithDashes(e164Phone, regionCode string) (string, error) {
	switch regionCode {
	case "TH":
		// +66-81-234-5678
		if len(e164Phone) == 12 {
			return e164Phone[:3] + "-" + e164Phone[3:5] + "-" + e164Phone[5:8] + "-" + e164Phone[8:], nil
		}
	case "US", "CA":
		// +1-555-123-4567
		if len(e164Phone) == 12 {
			return e164Phone[:2] + "-" + e164Phone[2:5] + "-" + e164Phone[5:8] + "-" + e164Phone[8:], nil
		}
	case "SG":
		// +65-9123-4567
		if len(e164Phone) == 11 {
			return e164Phone[:3] + "-" + e164Phone[3:7] + "-" + e164Phone[7:], nil
		}
	}

	// Fallback: generic formatting
	return e164Phone, nil
}

// formatDomesticWithDashes formats domestic number with dashes.
func formatDomesticWithDashes(domestic, regionCode string) (string, error) {
	switch regionCode {
	case "TH":
		// 081-234-5678
		if len(domestic) == 10 {
			return domestic[:3] + "-" + domestic[3:6] + "-" + domestic[6:], nil
		}
	case "US", "CA":
		// 555-123-4567
		if len(domestic) == 10 {
			return domestic[:3] + "-" + domestic[3:6] + "-" + domestic[6:], nil
		}
	case "SG":
		// 9123-4567
		if len(domestic) == 8 {
			return domestic[:4] + "-" + domestic[4:], nil
		}
	}

	return domestic, nil
}
