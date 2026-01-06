// Package main demonstrates the usage of the xstr EMVCo QR string parsing functionality.
package main

import (
	"fmt"

	xstr "github.com/hotfixfirst/go-xstr"
)

func main() {
	fmt.Println("=== EMV Co QR Examples ===")
	fmt.Println()

	// Example 1: Parse PromptPay QR with phone number
	fmt.Println("1. ParseEMVCoQRString - PromptPay with Phone Number")
	fmt.Println("----------------------------------------------------")

	qrWithPhone := "00020101021129370016A000000677010111011300668123456785802TH5303764540510.006304EFD4"
	fmt.Printf("QR String: %s\n\n", qrWithPhone)

	info, err := xstr.ParseEMVCoQRString(qrWithPhone)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Println("Parsed Information:")
		fmt.Printf("  Format:           %s\n", info.Format)
		fmt.Printf("  Phone Number:     %s\n", info.PhoneNumber)
		fmt.Printf("  Amount:           %s\n", info.Amount)
		fmt.Printf("  Country Code:     %s\n", info.CountryCode)
		fmt.Printf("  Currency (ISO):   %s\n", info.CurrencyISO4217)
		fmt.Printf("  CRC:              %s\n", info.Crc)
	}

	fmt.Println()

	// Example 2: Error handling - Invalid CRC
	fmt.Println("2. Error Handling - Invalid CRC")
	fmt.Println("---------------------------------")
	invalidCRC := "010201630441C6"
	fmt.Printf("QR String: %s\n", invalidCRC)
	_, err = xstr.ParseEMVCoQRString(invalidCRC)
	if err != nil {
		fmt.Printf("Expected Error: %v\n", err)
	}

	fmt.Println()

	// Example 3: Error handling - Too short
	fmt.Println("3. Error Handling - Too Short")
	fmt.Println("-------------------------------")
	tooShort := "01020163"
	fmt.Printf("QR String: %s\n", tooShort)
	_, err = xstr.ParseEMVCoQRString(tooShort)
	if err != nil {
		fmt.Printf("Expected Error: %v\n", err)
	}

	fmt.Println()
	fmt.Println("=== End of Examples ===")
}
