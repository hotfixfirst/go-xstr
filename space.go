package xstr

import "strings"

// Unicode zero-width and special spaces to remove
var zeroWidthChars = map[rune]bool{
	'\u200B': true, // Zero-width space
	'\uFEFF': true, // BOM
	'\u2060': true, // Word joiner
	'\u200D': true, // Zero-width joiner
	'\u200E': true, // LTR mark
	'\u200F': true, // RTL mark
	'\u00A0': true, // Non-breaking space
}

// RemoveDuplicateSpaces removes zero-width chars, trims leading/trailing whitespace,
// and replaces all internal whitespace (space, tab, newline, etc.) with a single space.
func RemoveDuplicateSpaces(s string) string {
	var b strings.Builder
	b.Grow(len(s))

	inSpace := false

	for _, r := range s {
		if zeroWidthChars[r] {
			continue
		}

		if isWhitespace(r) {
			if !inSpace {
				b.WriteByte(' ')
				inSpace = true
			}
		} else {
			b.WriteRune(r)
			inSpace = false
		}
	}

	return strings.TrimSpace(b.String())
}

func isWhitespace(r rune) bool {
	return r == ' ' || r == '\t' || r == '\n' || r == '\r' || r == '\v' || r == '\f'
}
