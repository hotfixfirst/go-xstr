// Package main demonstrates the usage of the xstr mask functionality.
package main

import (
	"fmt"

	xstr "github.com/hotfixfirst/go-xstr"
)

func main() {
	fmt.Println("=== Mask Examples ===")
	fmt.Println()

	// Example 1: Mask Sensitive Data
	fmt.Println("1. MaskSensitive - Mask sensitive string data")
	fmt.Println("-----------------------------------------------")
	sensitiveData := []string{
		"",
		"1234",
		"12345678",
		"1234567890123456",
		"A1B2C3D4E5F6G7H8",
	}
	for _, data := range sensitiveData {
		masked := xstr.MaskSensitive(data)
		fmt.Printf("  Input: %-20q -> Output: %q\n", data, masked)
	}

	fmt.Println()

	// Example 2: Mask Email Addresses
	fmt.Println("2. MaskEmail - Mask email addresses")
	fmt.Println("------------------------------------")
	emails := []string{
		"",
		"a@example.com",
		"ab@example.com",
		"john@example.com",
		"john.doe@company.co.th",
	}
	for _, email := range emails {
		masked := xstr.MaskEmail(email)
		fmt.Printf("  Input: %-28q -> Output: %q\n", email, masked)
	}

	fmt.Println()

	// Example 3: Mask Phone Numbers
	fmt.Println("3. MaskPhone - Mask phone numbers")
	fmt.Println("----------------------------------")
	phones := []string{
		"",
		"081234",
		"0812345678",
		"+66812345678",
		"+1234567890123",
	}
	for _, phone := range phones {
		masked := xstr.MaskPhone(phone)
		fmt.Printf("  Input: %-18q -> Output: %q\n", phone, masked)
	}

	fmt.Println()
	fmt.Println("=== End of Examples ===")
}
