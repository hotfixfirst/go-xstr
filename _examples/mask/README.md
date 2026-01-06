# Mask Example

This example demonstrates the `xstr` mask functionality for secure logging.

## Run

```bash
cd _examples/mask
go run main.go
```

## Features Demonstrated

| #   | Feature                | Function          |
|-----|------------------------|-------------------|
| 1   | Mask sensitive strings | `MaskSensitive()` |
| 2   | Mask email addresses   | `MaskEmail()`     |
| 3   | Mask phone numbers     | `MaskPhone()`     |

## Sample Output

```text
=== Mask Examples ===

1. MaskSensitive - Mask sensitive string data
-----------------------------------------------
  Input: ""                   -> Output: ""
  Input: "1234"               -> Output: "****"
  Input: "12345678"           -> Output: "****78"
  Input: "1234567890123456"   -> Output: "1234****3456"
  Input: "A1B2C3D4E5F6G7H8"   -> Output: "A1B2****G7H8"

2. MaskEmail - Mask email addresses
------------------------------------
  Input: ""                           -> Output: ""
  Input: "a@example.com"              -> Output: "a***@example.com"
  Input: "ab@example.com"             -> Output: "a***@example.com"
  Input: "john@example.com"           -> Output: "j****n@example.com"
  Input: "john.doe@company.co.th"     -> Output: "j****e@company.co.th"

3. MaskPhone - Mask phone numbers
----------------------------------
  Input: ""                 -> Output: ""
  Input: "081234"           -> Output: "****"
  Input: "0812345678"       -> Output: "****5678"
  Input: "+66812345678"     -> Output: "+66****5678"
  Input: "+1234567890123"   -> Output: "+12****0123"

=== End of Examples ===
```
