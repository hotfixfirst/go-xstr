// Package main demonstrates the usage of the xstr phone functionality.
package main

import (
	"fmt"

	xstr "github.com/hotfixfirst/go-xstr"
)

func main() {
	fmt.Println("=== Phone Examples ===")
	fmt.Println()

	// Example 1: Normalize to E.164
	fmt.Println("1. NormalizePhoneToE164 - Convert to international format")
	fmt.Println("-----------------------------------------------------------")
	phones := []string{
		"+66812345678",
		"66812345678",
		"0812345678",
		"812345678",
	}
	for _, phone := range phones {
		result, err := xstr.NormalizePhoneToE164(phone)
		if err != nil {
			fmt.Printf("  Input: %-16q -> Error: %v\n", phone, err)
		} else {
			fmt.Printf("  Input: %-16q -> Output: %q\n", phone, result)
		}
	}

	fmt.Println()

	// Example 2: Convert Phone Format
	fmt.Println("2. ConvertPhoneFormat - Convert between formats")
	fmt.Println("------------------------------------------------")
	e164Phone := "+66812345678"
	formats := []struct {
		name   string
		format xstr.PhoneFormat
	}{
		{"E164", xstr.PhoneFormatE164},
		{"Domestic", xstr.PhoneFormatDomestic},
		{"E164Dashed", xstr.PhoneFormatE164Dashed},
		{"DomesticDashed", xstr.PhoneFormatDomesticDashed},
	}
	fmt.Printf("  Source: %s\n", e164Phone)
	for _, f := range formats {
		result, err := xstr.ConvertPhoneFormat(e164Phone, f.format)
		if err != nil {
			fmt.Printf("    %-16s -> Error: %v\n", f.name, err)
		} else {
			fmt.Printf("    %-16s -> %s\n", f.name, result)
		}
	}

	fmt.Println()

	// Example 3: Check Mobile Number
	fmt.Println("3. IsMobileNumber - Check if phone is mobile")
	fmt.Println("---------------------------------------------")
	mobileTests := []string{
		"+66812345678",
		"+6621234567",
	}
	for _, phone := range mobileTests {
		isMobile := xstr.IsMobileNumber(phone)
		fmt.Printf("  Phone: %-16s -> IsMobile: %v\n", phone, isMobile)
	}

	fmt.Println()

	// Example 4: Get Country Code
	fmt.Println("4. GetPhoneCountryCode - Get country from phone")
	fmt.Println("------------------------------------------------")
	countryTests := []string{
		"+66812345678",
		"+6591234567",
		"+15551234567",
	}
	for _, phone := range countryTests {
		country, err := xstr.GetPhoneCountryCode(phone)
		if err != nil {
			fmt.Printf("  Phone: %-16s -> Error: %v\n", phone, err)
		} else {
			fmt.Printf("  Phone: %-16s -> Country: %s\n", phone, country)
		}
	}

	fmt.Println()
	fmt.Println("=== End of Examples ===")
}
