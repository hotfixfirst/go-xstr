// Package main demonstrates the usage of the xstr space functionality.
package main

import (
	"fmt"

	xstr "github.com/hotfixfirst/go-xstr"
)

func main() {
	fmt.Println("=== Space Examples ===")
	fmt.Println()

	// Example 1: Basic duplicate space removal
	fmt.Println("1. RemoveDuplicateSpaces - Basic usage")
	fmt.Println("---------------------------------------")
	basicTests := []string{
		"hello world",
		"hello  world",
		"hello   world   foo",
		"  hello  world  ",
	}
	for _, input := range basicTests {
		result := xstr.RemoveDuplicateSpaces(input)
		fmt.Printf("  Input: %-28q -> Output: %q\n", input, result)
	}

	fmt.Println()

	// Example 2: Tab and newline handling
	fmt.Println("2. Tab and newline handling")
	fmt.Println("----------------------------")
	whitespaceTests := []string{
		"hello\tworld",
		"hello\nworld",
		"hello\t\n\t  world",
	}
	for _, input := range whitespaceTests {
		result := xstr.RemoveDuplicateSpaces(input)
		fmt.Printf("  Input: %-28q -> Output: %q\n", input, result)
	}

	fmt.Println()

	// Example 3: Zero-width character removal
	fmt.Println("3. Zero-width character removal")
	fmt.Println("--------------------------------")
	fmt.Println("  Removes: Zero-width space, BOM, Word joiner, etc.")
	zeroWidthTests := []struct {
		name  string
		input string
	}{
		{"Zero-width space", "hello\u200Bworld"},
		{"BOM character", "\uFEFFhello world"},
		{"Non-breaking space", "hello\u00A0world"},
	}
	for _, test := range zeroWidthTests {
		result := xstr.RemoveDuplicateSpaces(test.input)
		fmt.Printf("  %-20s: %q -> %q\n", test.name, test.input, result)
	}

	fmt.Println()
	fmt.Println("=== End of Examples ===")
}
