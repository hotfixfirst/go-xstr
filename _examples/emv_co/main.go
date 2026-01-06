// Package main demonstrates the usage of the xstr EMV QR Code decoding functionality.
package main

import (
	"encoding/json"
	"fmt"

	xstr "github.com/hotfixfirst/go-xstr"
)

func main() {
	fmt.Println("=== EMV Co Examples ===")
	fmt.Println()

	// Example 1: Decode EMV QR Code
	fmt.Println("1. DecodeEMVQR - Decode EMV QR Code string")
	fmt.Println("-------------------------------------------")

	// Sample Thai PromptPay QR Code
	qrString := "00020101021129370016A000000677010111011300668123456785802TH5303764540510.006304EFD4"

	fmt.Printf("QR String: %s\n\n", qrString)

	emvData, err := xstr.DecodeEMVQR(qrString)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Println("Decoded Fields:")
	fmt.Printf("  Payload Format Indicator: %s\n", emvData.PayloadFormatIndicator)
	fmt.Printf("  Point of Initiation:      %s (%s)\n", emvData.PointOfInitiationMethod, emvData.POIMethodType)
	fmt.Printf("  Country Code:             %s\n", emvData.CountryCode)
	fmt.Printf("  Transaction Currency:     %s\n", emvData.TransactionCurrency)
	fmt.Printf("  Transaction Amount:       %s\n", emvData.TransactionAmount)
	fmt.Printf("  CRC:                      %s\n", emvData.CRC)

	fmt.Println()

	// Example 2: Merchant Account Info
	fmt.Println("2. Merchant Account Information")
	fmt.Println("--------------------------------")
	if len(emvData.MerchantAccountInfo) > 0 {
		for tag, merchant := range emvData.MerchantAccountInfo {
			fmt.Printf("  Tag %s:\n", tag)
			fmt.Printf("    AID:            %s\n", merchant.AID)
			fmt.Printf("    AID Type:       %s\n", merchant.AIDType)
			fmt.Printf("    Payment Scheme: %s\n", merchant.PaymentScheme)
			fmt.Printf("    Merchant ID:    %s\n", merchant.MerchantID)
		}
	} else {
		fmt.Println("  No merchant account info found")
	}

	fmt.Println()

	// Example 3: JSON output
	fmt.Println("3. JSON Output (partial)")
	fmt.Println("-------------------------")
	jsonBytes, _ := json.MarshalIndent(map[string]string{
		"payload_format": emvData.PayloadFormatIndicator,
		"country":        emvData.CountryCode,
		"currency":       emvData.TransactionCurrency,
		"amount":         emvData.TransactionAmount,
		"crc":            emvData.CRC,
	}, "  ", "  ")
	fmt.Printf("  %s\n", string(jsonBytes))

	fmt.Println()
	fmt.Println("=== End of Examples ===")
}
