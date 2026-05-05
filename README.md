# Goal

[![Go Version](https://img.shields.io/badge/go-1.24+-blue.svg)](https://go.dev/)
[![License](https://img.shields.io/badge/license-Apache--2.0-green.svg)](LICENSE)

[中文](./README_CN.md)

Goal is a Go utility library with focused packages for collections, conversion,
errors, HTTP, I/O, JSON, random data, retry, time, UUIDs, and value handling.

## Install

```bash
go get github.com/go4x/goal
```

## Example

```go
package main

import (
	"fmt"

	"github.com/go4x/goal/col/set"
	"github.com/go4x/goal/is"
	"github.com/go4x/goal/value"
)

func main() {
	names := set.New[string]().Add("alice").Add("bob")

	fmt.Println(names.Contains("alice"))
	fmt.Println(value.Or("", "fallback"))
	fmt.Println(is.Empty([]string{}))
}
```

## Packages

| Package | Purpose |
| --- | --- |
| [`assert`](./assert/) | Test assertions |
| [`ciphers`](./ciphers/) | AES, hash, Base64, and encoding helpers |
| [`cmd`](./cmd/) | Command execution |
| [`col/mapx`](./col/mapx/) | Map implementations |
| [`col/set`](./col/set/) | Set implementations |
| [`col/slicex`](./col/slicex/) | Slice helpers |
| [`conv`](./conv/) | Type conversions |
| [`errorx`](./errorx/) | Error helpers, wrapping, and recover helpers |
| [`httpx`](./httpx/) | HTTP request and async client helpers |
| [`iox`](./iox/) | File, directory, path, and walker helpers |
| [`is`](./is/) | Value checks and comparisons |
| [`jsonx`](./jsonx/) | JSON helpers |
| [`limiter`](./limiter/) | Token bucket limiter |
| [`ptr`](./ptr/) | Pointer helpers |
| [`random`](./random/) | Random numbers and strings |
| [`reflectx`](./reflectx/) | Reflection helpers |
| [`retry`](./retry/) | Retry helpers |
| [`stringx`](./stringx/) | String helpers |
| [`timex`](./timex/) | Time helpers |
| [`uuid`](./uuid/) | UUID and distributed ID helpers |
| [`value`](./value/) | Value selection and Must-style helpers |

See package directories for API examples and package-specific notes.

## Development

```bash
go test ./...
go test -race ./...
go test -bench=. ./...
```

Some APIs intentionally panic when used in `Must` or `Force` style. Prefer
error-returning APIs in library code when the input may be invalid.

## Documentation

- [Changelog](./CHANGELOG.md)
- Package-specific README files in package directories
- Workspace-level docs, when this module is checked out inside go4x, live in
  `../docs/`

## License

Apache License 2.0. See [LICENSE](./LICENSE).
