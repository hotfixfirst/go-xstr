package xstr

import "strings"

// NormalizeOptionalString trims whitespace from a string pointer.
// Returns nil if the input is nil or the trimmed string is empty.
func NormalizeOptionalString(value *string) *string {
	if value == nil {
		return nil
	}
	trimmed := strings.TrimSpace(*value)
	if trimmed == "" {
		return nil
	}
	return &trimmed
}

// NormalizeUpdateString trims whitespace from a string pointer.
// Returns nil if the input is nil.
// Returns a pointer to an empty string if the trimmed result is empty.
// This is useful for update operations where empty string clears the field.
func NormalizeUpdateString(value *string) *string {
	if value == nil {
		return nil
	}
	trimmed := strings.TrimSpace(*value)
	if trimmed == "" {
		empty := ""
		return &empty
	}
	return &trimmed
}
