# EMV Co Example

This example demonstrates the `xstr` EMV QR Code decoding functionality.

## Run

```bash
cd _examples/emv_co
go run main.go
```

## Features Demonstrated

| #   | Feature               | Function/Type            |
|-----|-----------------------|--------------------------|
| 1   | Decode EMV QR         | `DecodeEMVQR()`          |
| 2   | Merchant Account Info | `MerchantAccount` struct |

## QR Payment Types

| Type               | Description                 |
|--------------------|-----------------------------||
| `QRTypeC2C`        | Consumer-to-Consumer Transfer |
| `QRTypeC2B`        | Consumer-to-Business        |
| `QRTypeBillPayment`| Bill Payment                |
| `QRTypeCrossBorder`| Cross-Border Payment        |

## Payment Schemes

| Scheme             | Description |
|--------------------|-------------|
| `QRSchemePromptPay`| Thailand    |
| `QRSchemeQRIS`     | Indonesia   |
| `QRSchemeDuitNow`  | Malaysia    |
| `QRSchemeUPI`      | India       |

## Sample Output

```text
=== EMV Co Examples ===

1. DecodeEMVQR - Decode EMV QR Code string
-------------------------------------------
QR String: 00020101021129370016A000000677010111011300668123456785802TH5303764540510.006304EFD4

Decoded Fields:
  Payload Format Indicator: 01
  Point of Initiation:      11 (dynamic)
  Country Code:             TH
  Transaction Currency:     764
  Transaction Amount:       10.00
  CRC:                      EFD4

2. Merchant Account Information
--------------------------------
  Tag 29:
    AID:            A000000677010111
    AID Type:       C2C
    Payment Scheme: PromptPay
    Merchant ID:    0066812345678

=== End of Examples ===
```
