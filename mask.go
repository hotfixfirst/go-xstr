package xstr

// MaskSensitive masks sensitive string data for secure logging.
// Shows first 4 and last 4 characters for strings longer than 8 characters.
// For shorter strings, shows only asterisks for security.
func MaskSensitive(data string) string {
	// Handle empty or very short strings
	switch {
	case len(data) == 0:
		return ""
	case len(data) <= 4:
		return "****"
	case len(data) <= 8:
		return "****" + data[len(data)-2:]
	default:
		return data[:4] + "****" + data[len(data)-4:]
	}
}

// MaskEmail masks email addresses for secure logging.
// Shows first character and domain while masking the local part.
func MaskEmail(email string) string {
	if len(email) == 0 {
		return ""
	}
	atIndex := -1
	for i, char := range email {
		if char == '@' {
			atIndex = i
			break
		}
	}
	if atIndex == -1 {
		return MaskSensitive(email)
	}
	localPart := email[:atIndex]
	domain := email[atIndex:]
	switch {
	case len(localPart) == 0:
		return "****" + domain
	case len(localPart) == 1:
		return localPart + "***" + domain
	case len(localPart) <= 4:
		return localPart[:1] + "***" + domain
	default:
		return localPart[:1] + "****" + localPart[len(localPart)-1:] + domain
	}
}

// MaskPhone masks phone numbers for secure logging.
// Shows country code and last 4 digits for international numbers.
func MaskPhone(phone string) string {
	if len(phone) == 0 {
		return ""
	}
	if phone[0] == '+' {
		if len(phone) <= 7 {
			return phone[:3] + "****"
		}
		return phone[:3] + "****" + phone[len(phone)-4:]
	}
	if len(phone) <= 6 {
		return "****"
	}
	return "****" + phone[len(phone)-4:]
}
