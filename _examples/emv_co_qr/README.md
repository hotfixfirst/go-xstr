# EMV Co QR Example

This example demonstrates the `xstr` EMVCo QR string parsing functionality.

## Run

```bash
cd _examples/emv_co_qr
go run main.go
```

## Features Demonstrated

| #   | Feature            | Function                 |
|-----|--------------------|--------------------------|
| 1   | Parse PromptPay QR | `ParseEMVCoQRString()`   |
| 2   | Error handling     | CRC validation           |

## EMVCoQRInfo Structure

```go
type EMVCoQRInfo struct {
    Format          string
    MerchantAccount string
    Amount          string
    PhoneNumber     string
    CountryCode     string
    CurrencyISO4217 string
    Crc             string
    BillerID        string
    Ref1, Ref2, Ref3 string
}
```

## Sample Output

```text
=== EMV Co QR Examples ===

1. ParseEMVCoQRString - PromptPay with Phone Number
----------------------------------------------------
QR String: 00020101021129370016A000000677010111011300668123456785802TH5303764540510.006304EFD4

Parsed Information:
  Format:           01
  Phone Number:     0066812345678
  Amount:           10.00
  Country Code:     TH
  Currency (ISO):   764
  CRC:              EFD4

2. Error Handling - Invalid CRC
---------------------------------
QR String: 010201630441C6
Expected Error: invalid crc: expected 41C6, got 41C5

3. Error Handling - Too Short
-------------------------------
QR String: 01020163
Expected Error: qr string too short

=== End of Examples ===
```
