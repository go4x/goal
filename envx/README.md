# envx

`envx` provides small helpers for reading environment variables.

The package uses only the Go standard library. It does not load `.env` files,
bind structs, watch configuration, or introduce external dependencies.

## Install

```bash
go get github.com/go4x/goal/envx
```

## Usage

```go
package main

import (
	"fmt"
	"time"

	"github.com/go4x/goal/envx"
)

func main() {
	host := envx.GetDefault("APP_HOST", "127.0.0.1")
	port := envx.GetInt("APP_PORT", 8080)
	debug := envx.GetBool("APP_DEBUG", false)
	timeout := envx.GetDuration("APP_TIMEOUT", 5*time.Second)
	services := envx.GetSlice("APP_SERVICES", ",", []string{"api"})

	fmt.Println(host, port, debug, timeout, services)
}
```

## API

| Function | Description |
| --- | --- |
| `Get(key string) string` | Returns the raw value, or an empty string when unset. |
| `Exists(key string) bool` | Reports whether the variable is set, even if it is empty. |
| `GetDefault(key, fallback string) string` | Returns fallback only when the variable is unset. |
| `Require(key string) (string, error)` | Returns an error when the variable is unset. |
| `GetInt(key string, fallback int) int` | Parses an int or returns fallback. |
| `GetBool(key string, fallback bool) bool` | Parses a bool or returns fallback. |
| `GetDuration(key string, fallback time.Duration) time.Duration` | Parses a Go duration or returns fallback. |
| `GetSlice(key, sep string, fallback []string) []string` | Splits a string value and trims empty items. |

## Notes

- `GetDefault` treats an empty but set variable as a real value.
- `Require` only fails when the variable is unset.
- Typed helpers return fallback for both missing and invalid values.
- `GetSlice` uses `,` when `sep` is empty.
