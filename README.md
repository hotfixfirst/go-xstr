# go-xstr

A lightweight Go string utility library with pointer helpers, safe nil-handling, masking, phone number formatting, EMV QR code parsing, and common transformations.

## Installation

```bash
go get github.com/hotfixfirst/go-xstr
```

## Quick Start

```go
import xstr "github.com/hotfixfirst/go-xstr"

// Mask sensitive data
masked := xstr.MaskSensitive("1234567890123456") // "1234****3456"

// Normalize phone number
phone, _ := xstr.NormalizePhoneToE164("0812345678") // "+66812345678"

// Remove duplicate spaces
clean := xstr.RemoveDuplicateSpaces("hello   world") // "hello world"
```

## Features

| Feature                       | Description                            | Documentation                      |
| ----------------------------- | -------------------------------------- | ---------------------------------- |
| [Mask](#mask)                 | Mask sensitive data for logging        | [Examples](./_examples/mask/)      |
| [Phone](#phone)               | Phone number parsing and formatting    | [Examples](./_examples/phone/)     |
| [Pointer](#pointer)           | String pointer normalization           | [Examples](./_examples/pointer/)   |
| [Space](#space)               | Whitespace and duplicate space removal | [Examples](./_examples/space/)     |
| [EMV Co](#emv-co)             | EMV QR Code decoding                   | [Examples](./_examples/emv_co/)    |
| [EMV Co QR](#emv-co-qr)       | EMVCo QR string parsing                | [Examples](./_examples/emv_co_qr/) |

---

## Mask

Mask sensitive data for secure logging.

| Function                       | Description                                              |
| ------------------------------ | -------------------------------------------------------- |
| `MaskSensitive(data string)`   | Mask sensitive strings (shows first 4 and last 4 chars)  |
| `MaskEmail(email string)`      | Mask email addresses                                     |
| `MaskPhone(phone string)`      | Mask phone numbers                                       |

```go
xstr.MaskSensitive("1234567890123456") // "1234****3456"
xstr.MaskEmail("john@example.com")     // "j***@example.com"
xstr.MaskPhone("+66812345678")         // "+66****5678"
```

---

## Phone

Phone number parsing, validation, and formatting with multi-country support.

| Function                                               | Description                        |
| ------------------------------------------------------ | ---------------------------------- |
| `NormalizePhoneToE164(phone string)`                   | Convert to E.164 format            |
| `ConvertPhoneFormat(phone string, format PhoneFormat)` | Convert between formats            |
| `IsMobileNumber(phone string)`                         | Check if phone is mobile           |
| `GetPhoneCountryCode(phone string)`                    | Get country code from phone        |
| `ConvertPhoneByCurrency(phone, currency string)`       | Convert based on currency          |
| `ValidatePhoneCurrency(phone, currency string)`        | Validate phone matches currency    |

**Supported Formats:**

| Format                      | Example           |
| --------------------------- | ----------------- |
| `PhoneFormatE164`           | `+66812345678`    |
| `PhoneFormatDomestic`       | `0812345678`      |
| `PhoneFormatE164Dashed`     | `+66-81-234-5678` |
| `PhoneFormatDomesticDashed` | `081-234-5678`    |

```go
// Normalize to E.164
phone, _ := xstr.NormalizePhoneToE164("0812345678") // "+66812345678"

// Convert format
domestic, _ := xstr.ConvertPhoneFormat("+66812345678", xstr.PhoneFormatDomestic) // "0812345678"

// Check mobile
isMobile := xstr.IsMobileNumber("+66812345678") // true
```

---

## Pointer

String pointer normalization utilities for handling optional fields.

| Function                                   | Description                       |
| ------------------------------------------ | --------------------------------- |
| `NormalizeOptionalString(value *string)`   | Returns nil if empty/whitespace   |
| `NormalizeUpdateString(value *string)`     | Returns empty string ptr if empty |

**Key Differences:**

| Input     | `NormalizeOptionalString` | `NormalizeUpdateString` |
| --------- | ------------------------- | ----------------------- |
| `nil`     | `nil`                     | `nil`                   |
| `""`      | `nil`                     | `""` (empty ptr)        |
| `"   "`   | `nil`                     | `""` (empty ptr)        |
| `"hello"` | `"hello"`                 | `"hello"`               |

```go
empty := ""
result := xstr.NormalizeOptionalString(&empty) // nil
result := xstr.NormalizeUpdateString(&empty)   // ptr to ""
```

---

## Space

Whitespace normalization and zero-width character removal.

| Function                          | Description                                                   |
| --------------------------------- | ------------------------------------------------------------- |
| `RemoveDuplicateSpaces(s string)` | Remove duplicate spaces, tabs, newlines, and zero-width chars |

**Zero-Width Characters Removed:**

- `\u200B` Zero-width space
- `\uFEFF` BOM
- `\u2060` Word joiner
- `\u200D` Zero-width joiner
- `\u00A0` Non-breaking space

```go
xstr.RemoveDuplicateSpaces("hello   world")     // "hello world"
xstr.RemoveDuplicateSpaces("hello\t\nworld")    // "hello world"
xstr.RemoveDuplicateSpaces("  hello  world  ")  // "hello world"
```

---

## EMV Co

EMV QR Code decoding with support for multiple payment schemes.

| Function                         | Description                            |
| -------------------------------- | -------------------------------------- |
| `DecodeEMVQR(qrString string)`   | Decode EMV QR code to structured data  |

**Supported Payment Schemes:**

| Scheme              | Country     |
| ------------------- | ----------- |
| `QRSchemePromptPay` | Thailand    |
| `QRSchemeQRIS`      | Indonesia   |
| `QRSchemeDuitNow`   | Malaysia    |
| `QRSchemeUPI`       | India       |
| `QRSchemeNETS`      | Singapore   |

```go
emvData, err := xstr.DecodeEMVQR(qrString)
if err != nil {
    log.Fatal(err)
}
fmt.Println(emvData.TransactionAmount)
fmt.Println(emvData.CountryCode)
```

---

## EMV Co QR

Simple EMVCo QR string parsing for Thai PromptPay.

| Function                                | Description            |
| --------------------------------------- | ---------------------- |
| `ParseEMVCoQRString(qrString string)`   | Parse EMVCo QR string  |

```go
info, err := xstr.ParseEMVCoQRString(qrString)
if err != nil {
    log.Fatal(err)
}
fmt.Println(info.PhoneNumber)
fmt.Println(info.Amount)
```

---

## Running Examples

See the [_examples](./_examples/) directory for runnable examples.

```bash
# Run from project root
go run ./_examples/mask/main.go
go run ./_examples/phone/main.go
go run ./_examples/pointer/main.go
go run ./_examples/space/main.go
go run ./_examples/emv_co/main.go
go run ./_examples/emv_co_qr/main.go
```

## License

MIT License - see [LICENSE](./LICENSE) for details.
