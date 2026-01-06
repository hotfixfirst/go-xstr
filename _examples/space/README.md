# Space Example

This example demonstrates the `xstr` whitespace and duplicate space removal functionality.

## Run

```bash
cd _examples/space
go run main.go
```

## Features Demonstrated

| #   | Feature                       | Function                  |
|-----|-------------------------------|---------------------------|
| 1   | Remove duplicate spaces       | `RemoveDuplicateSpaces()` |
| 2   | Handle tabs and newlines      | `RemoveDuplicateSpaces()` |
| 3   | Remove zero-width characters  | `RemoveDuplicateSpaces()` |

## Zero-Width Characters Removed

| Character  | Unicode  | Description             |
|------------|----------|-------------------------|
| `\u200B`   | U+200B   | Zero-width space        |
| `\uFEFF`   | U+FEFF   | Byte Order Mark (BOM)   |
| `\u2060`   | U+2060   | Word joiner             |
| `\u200D`   | U+200D   | Zero-width joiner       |
| `\u00A0`   | U+00A0   | Non-breaking space      |

## Sample Output

```text
=== Space Examples ===

1. RemoveDuplicateSpaces - Basic usage
---------------------------------------
  Input: "hello world"                -> Output: "hello world"
  Input: "hello  world"               -> Output: "hello world"
  Input: "hello   world   foo"        -> Output: "hello world foo"
  Input: "  hello  world  "           -> Output: "hello world"

2. Tab and newline handling
----------------------------
  Input: "hello\tworld"               -> Output: "hello world"
  Input: "hello\nworld"               -> Output: "hello world"
  Input: "hello\t\n\t  world"         -> Output: "hello world"

3. Zero-width character removal
--------------------------------
  Removes: Zero-width space, BOM, Word joiner, etc.
  Zero-width space    : "hello​world" -> "helloworld"
  BOM character       : "﻿hello world" -> "hello world"
  Non-breaking space  : "hello world" -> "hello world"

=== End of Examples ===
```
