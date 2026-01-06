# Phone Example

This example demonstrates the `xstr` phone number parsing and formatting functionality.

## Run

```bash
cd _examples/phone
go run main.go
```

## Features Demonstrated

| #   | Feature              | Function                 |
|-----|----------------------|--------------------------|
| 1   | Normalize to E.164   | `NormalizePhoneToE164()` |
| 2   | Convert phone format | `ConvertPhoneFormat()`   |
| 3   | Check mobile number  | `IsMobileNumber()`       |
| 4   | Get country code     | `GetPhoneCountryCode()`  |

## Supported Formats

| Format                      | Example            |
|-----------------------------|--------------------||
| `PhoneFormatE164`           | `+66812345678`     |
| `PhoneFormatDomestic`       | `0812345678`       |
| `PhoneFormatE164Dashed`     | `+66-81-234-5678`  |
| `PhoneFormatDomesticDashed` | `081-234-5678`     |

## Sample Output

```text
=== Phone Examples ===

1. NormalizePhoneToE164 - Convert to international format
-----------------------------------------------------------
  Input: "+66812345678"    -> Output: "+66812345678"
  Input: "66812345678"     -> Output: "+66812345678"
  Input: "0812345678"      -> Output: "+66812345678"
  Input: "812345678"       -> Output: "+66812345678"

2. ConvertPhoneFormat - Convert between formats
------------------------------------------------
  Source: +66812345678
    E164             -> +66812345678
    Domestic         -> 0812345678
    E164Dashed       -> +66-81-234-5678
    DomesticDashed   -> 081-234-5678

3. IsMobileNumber - Check if phone is mobile
---------------------------------------------
  Phone: +66812345678     -> IsMobile: true
  Phone: +6621234567      -> IsMobile: false

4. GetPhoneCountryCode - Get country from phone
------------------------------------------------
  Phone: +66812345678     -> Country: TH
  Phone: +6591234567      -> Country: SG
  Phone: +15551234567     -> Country: US

=== End of Examples ===
```
