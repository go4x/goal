# Goal

[![Go Version](https://img.shields.io/badge/go-1.24+-blue.svg)](https://go.dev/)
[![License](https://img.shields.io/badge/license-Apache--2.0-green.svg)](LICENSE)

[中文](./README_CN.md)

Goal is a Go utility library with focused packages for collections, conversion,
errors, HTTP, I/O, JSON, random data, retry, time, UUIDs, and value handling.

## 🚀 Features

- **Comprehensive Coverage**: 20+ specialized packages covering various domains
- **Generic Type Support**: Modern Go generics throughout the codebase
- **Few Dependencies**: Few external dependencies, easy to maintain
- **Performance Optimized**: Carefully designed for efficiency and speed
- **Test Coverage**: Comprehensive test coverage
- **Well Documented**: Extensive documentation with examples
- **Production Ready**: Battle-tested in real-world applications

## 📦 Packages Overview

### Core Utilities

| Package                | Description                   | Features                                                     |
| ---------------------- | ----------------------------- | ------------------------------------------------------------ |
| [`assert`](./assert/)  | Testing assertions            | Type-safe assertions, custom messages                        |
| [`value`](./value/)    | Value handling utilities      | Null checks, conditional logic, safe operations              |
| [`ptr`](./ptr/)        | Pointer utilities             | Safe dereferencing, pointer operations                       |
| [`conv`](./conv/)      | Type conversion               | Safe conversions, format validation                          |
| [`envx`](./envx/)      | Environment variables         | Default values, required values, typed parsing               |
| [`is`](./is/) (v1.1.0) | Value checking and comparison | Boolean operations, zero/nil/empty checks, value comparisons |

### Collections

| Package                       | Description               | Features                                                 |
| ----------------------------- | ------------------------- | -------------------------------------------------------- |
| [`col/set`](./col/set/)       | Set implementations       | HashSet, ArraySet, LinkedSet with O(1) operations        |
| [`col/mapx`](./col/mapx/)     | Map implementations       | Regular Map, ArrayMap, LinkedMap with order preservation |
| [`col/slicex`](./col/slicex/) | Enhanced slice operations | Immutable operations, functional programming             |

### String & Text

| Package                 | Description        | Features                                       |
| ----------------------- | ------------------ | ---------------------------------------------- |
| [`stringx`](./stringx/) | String utilities   | Case conversion, blurring, constants, builders |
| [`color`](./color/)     | Color manipulation | RGB operations, color conversion               |
| [`jsonx`](./jsonx/)     | JSON utilities     | Enhanced JSON operations, validation           |

### System & I/O

| Package             | Description       | Features                                          |
| ------------------- | ----------------- | ------------------------------------------------- |
| [`cmd`](./cmd/)     | Command execution | Async execution, timeout handling, streaming      |
| [`iox`](./iox/)     | I/O utilities     | File operations, directory walking, path handling |
| [`httpx`](./httpx/) | HTTP client       | Async client, request/response handling           |

### Cryptography & Security

| Package                 | Description             | Features                                    |
| ----------------------- | ----------------------- | ------------------------------------------- |
| [`ciphers`](./ciphers/) | Cryptographic functions | AES encryption, hashing, data compression   |
| [`uuid`](./uuid/)       | UUID generation         | Standard UUIDs, distributed IDs (Sonyflake) |

### Error Handling

| Package               | Description     | Features                           |
| --------------------- | --------------- | ---------------------------------- |
| [`errorx`](./errorx/) | Error utilities | Error chaining, wrapping, recovery |

### Mathematics & Statistics

| Package               | Description           | Features                                         |
| --------------------- | --------------------- | ------------------------------------------------ |
| [`prob`](./prob/)     | Probability functions | Statistical operations, probability calculations |
| [`random`](./random/) | Random generation     | Number generation, string generation             |

### Utilities

| Package                   | Description          | Features                               |
| ------------------------- | -------------------- | -------------------------------------- |
| [`timex`](./timex/)       | Time utilities       | Time formatting, parsing, operations   |
| [`limiter`](./limiter/)   | Rate limiting        | Token bucket, rate limiting algorithms |
| [`retry`](./retry/)       | Retry mechanisms     | Exponential backoff, retry strategies  |
| [`reflectx`](./reflectx/) | Reflection utilities | Type inspection, reflection helpers    |
| [`printer`](./printer/)   | Printing utilities   | Formatted output, pretty printing      |

## 🚀 Quick Start

### Installation

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
