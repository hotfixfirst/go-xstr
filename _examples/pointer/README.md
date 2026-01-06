# Pointer Example

This example demonstrates the `xstr` string pointer normalization utilities.

## Run

```bash
cd _examples/pointer
go run main.go
```

## Features Demonstrated

| #   | Feature                    | Function                    |
|-----|----------------------------|-----------------------------|
| 1   | Normalize optional strings | `NormalizeOptionalString()` |
| 2   | Normalize update strings   | `NormalizeUpdateString()`   |

## Key Differences

| Input                  | `NormalizeOptionalString` | `NormalizeUpdateString` |
|------------------------|---------------------------|-------------------------|
| `nil`                  | `nil`                     | `nil`                   |
| `""` (empty)           | `nil`                     | `""` (empty pointer)    |
| `"   "` (whitespace)   | `nil`                     | `""` (empty pointer)    |
| `"hello"`              | `"hello"`                 | `"hello"`               |

## Sample Output

```text
=== Pointer Examples ===

1. NormalizeOptionalString - Returns nil if empty
--------------------------------------------------
  nil         : <nil>            -> <nil>
  empty       : ""               -> <nil>
  whitespace  : "   "            -> <nil>
  normal      : "hello"          -> "hello"
  with spaces : "  hello  "      -> "hello"

2. NormalizeUpdateString - Returns empty string ptr if empty
--------------------------------------------------------------
  nil         : <nil>            -> <nil>
  empty       : ""               -> ""
  whitespace  : "   "            -> ""
  normal      : "hello"          -> "hello"
  with spaces : "  hello  "      -> "hello"

3. Key Difference
-----------------
  NormalizeOptionalString: empty/whitespace -> nil
  NormalizeUpdateString:   empty/whitespace -> "" (empty string ptr)

=== End of Examples ===
```
