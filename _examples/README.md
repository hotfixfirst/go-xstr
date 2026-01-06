# Examples

Runnable examples demonstrating `go-xstr` features.

## Table of Contents

| Example                     | Description                               | Run                              |
|-----------------------------|-------------------------------------------|----------------------------------|
| [mask](./mask/)             | Masking sensitive data for secure logging | `cd mask && go run main.go`      |
| [phone](./phone/)           | Phone number parsing and formatting       | `cd phone && go run main.go`     |
| [pointer](./pointer/)       | String pointer normalization utilities    | `cd pointer && go run main.go`   |
| [space](./space/)           | Whitespace and duplicate space removal    | `cd space && go run main.go`     |
| [emv_co](./emv_co/)         | EMV QR Code decoding and parsing          | `cd emv_co && go run main.go`    |
| [emv_co_qr](./emv_co_qr/)   | EMVCo QR string parsing                   | `cd emv_co_qr && go run main.go` |

## Quick Start

```bash
# Clone the repository
git clone https://github.com/hotfixfirst/go-xstr.git
cd go-xstr/_examples

# Run a specific example
cd mask && go run main.go
```

## Adding New Examples

When adding a new feature example:

1. Create a new directory: `_examples/{feature}/`
2. Add `main.go` with runnable code
3. Add `README.md` with documentation
4. Update this file's table of contents
