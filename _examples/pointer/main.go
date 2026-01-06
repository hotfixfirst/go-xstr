// Package main demonstrates the usage of the xstr pointer functionality.
package main

import (
	"fmt"

	xstr "github.com/hotfixfirst/go-xstr"
)

func stringPtr(s string) *string {
	return &s
}

func main() {
	fmt.Println("=== Pointer Examples ===")
	fmt.Println()

	// Example 1: NormalizeOptionalString
	fmt.Println("1. NormalizeOptionalString - Returns nil if empty")
	fmt.Println("--------------------------------------------------")

	optionalTests := []struct {
		name  string
		input *string
	}{
		{"nil", nil},
		{"empty", stringPtr("")},
		{"whitespace", stringPtr("   ")},
		{"normal", stringPtr("hello")},
		{"with spaces", stringPtr("  hello  ")},
	}

	for _, test := range optionalTests {
		result := xstr.NormalizeOptionalString(test.input)
		inputStr := "<nil>"
		if test.input != nil {
			inputStr = fmt.Sprintf("%q", *test.input)
		}
		resultStr := "<nil>"
		if result != nil {
			resultStr = fmt.Sprintf("%q", *result)
		}
		fmt.Printf("  %-12s: %-16s -> %s\n", test.name, inputStr, resultStr)
	}

	fmt.Println()

	// Example 2: NormalizeUpdateString
	fmt.Println("2. NormalizeUpdateString - Returns empty string ptr if empty")
	fmt.Println("--------------------------------------------------------------")

	for _, test := range optionalTests {
		result := xstr.NormalizeUpdateString(test.input)
		inputStr := "<nil>"
		if test.input != nil {
			inputStr = fmt.Sprintf("%q", *test.input)
		}
		resultStr := "<nil>"
		if result != nil {
			resultStr = fmt.Sprintf("%q", *result)
		}
		fmt.Printf("  %-12s: %-16s -> %s\n", test.name, inputStr, resultStr)
	}

	fmt.Println()

	// Example 3: Comparison
	fmt.Println("3. Key Difference")
	fmt.Println("-----------------")
	fmt.Println("  NormalizeOptionalString: empty/whitespace -> nil")
	fmt.Println("  NormalizeUpdateString:   empty/whitespace -> \"\" (empty string ptr)")

	fmt.Println()
	fmt.Println("=== End of Examples ===")
}
