# Goal

[![Go Version](https://img.shields.io/badge/go-1.24+-blue.svg)](https://golang.org/)
[![License](https://img.shields.io/badge/license-apache2.0-green.svg)](LICENSE)
[![Version](https://img.shields.io/badge/version-v1.0.0-brightgreen.svg)](https://github.com/go4x/goal/releases/tag/v1.0.0)

[中文](./README_CN.md)

A comprehensive Go utility library that provides a rich set of packages for common programming tasks, data structures, and system operations.

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

| Package | Description | Features |
|---------|-------------|----------|
| [`assert`](./assert/) | Testing assertions | Type-safe assertions, custom messages |
| [`value`](./value/) | Value handling utilities | Null checks, conditional logic, safe operations |
| [`ptr`](./ptr/) | Pointer utilities | Safe dereferencing, pointer operations |
| [`conv`](./conv/) | Type conversion | Safe conversions, format validation |

### Collections

| Package | Description | Features |
|---------|-------------|----------|
| [`col/set`](./col/set/) | Set implementations | HashSet, ArraySet, LinkedSet with O(1) operations |
| [`col/mapx`](./col/mapx/) | Map implementations | Regular Map, ArrayMap, LinkedMap with order preservation |
| [`col/slicex`](./col/slicex/) | Enhanced slice operations | Immutable operations, functional programming |

### String & Text

| Package | Description | Features |
|---------|-------------|----------|
| [`stringx`](./stringx/) | String utilities | Case conversion, blurring, constants, builders |
| [`color`](./color/) | Color manipulation | RGB operations, color conversion |
| [`jsonx`](./jsonx/) | JSON utilities | Enhanced JSON operations, validation |

### System & I/O

| Package | Description | Features |
|---------|-------------|----------|
| [`cmd`](./cmd/) | Command execution | Async execution, timeout handling, streaming |
| [`iox`](./iox/) | I/O utilities | File operations, directory walking, path handling |
| [`httpx`](./httpx/) | HTTP client | Async client, request/response handling |

### Cryptography & Security

| Package | Description | Features |
|---------|-------------|----------|
| [`ciphers`](./ciphers/) | Cryptographic functions | AES encryption, hashing, data compression |
| [`uuid`](./uuid/) | UUID generation | Standard UUIDs, distributed IDs (Sonyflake) |

### Error Handling

| Package | Description | Features |
|---------|-------------|----------|
| [`errorx`](./errorx/) | Error utilities | Error chaining, wrapping, recovery |

### Mathematics & Statistics

| Package | Description | Features |
|---------|-------------|----------|
| [`mathx`](./mathx/) | Mathematical utilities | Advanced math operations, calculations |
| [`prob`](./prob/) | Probability functions | Statistical operations, probability calculations |
| [`random`](./random/) | Random generation | Number generation, string generation |

### Utilities

| Package | Description | Features |
|---------|-------------|----------|
| [`timex`](./timex/) | Time utilities | Time formatting, parsing, operations |
| [`limiter`](./limiter/) | Rate limiting | Token bucket, rate limiting algorithms |
| [`retry`](./retry/) | Retry mechanisms | Exponential backoff, retry strategies |
| [`reflectx`](./reflectx/) | Reflection utilities | Type inspection, reflection helpers |
| [`printer`](./printer/) | Printing utilities | Formatted output, pretty printing |

## 🚀 Quick Start

### Installation

```bash
go get github.com/go4x/goal
```

### Basic Usage

```go
package main

import (
    "fmt"
    "github.com/go4x/goal/stringx"
    "github.com/go4x/goal/value"
    "github.com/go4x/goal/col/set"
)

func main() {
    // String operations
    result := stringx.ToCamelCase("hello_world")
    fmt.Println(result) // "helloWorld"
    
    // Value handling
    safeValue := value.Or("", "", "fallback")
    fmt.Println(safeValue) // "fallback"
    
    // Set operations
    mySet := set.New[string]()
    mySet.Add("apple").Add("banana")
    fmt.Println(mySet.Contains("apple")) // true
}
```

## 📚 Package Examples

### String Operations

```go
import "github.com/go4x/goal/stringx"

// Case conversion
camel := stringx.ToCamelCase("hello_world")     // "helloWorld"
pascal := stringx.ToPascalCase("hello_world")   // "HelloWorld"
kebab := stringx.ToKebabCase("HelloWorld")      // "hello-world"

// String blurring
blurred := stringx.BlurEmail("user@example.com") // "u****@example.com"

// String building
builder := stringx.NewBuilder()
builder.WriteString("Hello ").WriteString("World")
result := builder.String() // "Hello World"
```

### Value Handling

```go
import "github.com/go4x/goal/value"

// Conditional logic
result := value.IfElse(age >= 18, "adult", "minor")

// Null/empty checks
if value.IsNotEmpty(data) {
    // Process data
}

// Value coalescing
fallback := value.Or("", "", "default")

// Safe operations
safeValue := value.Must(strconv.Atoi("123"))
```

### Collections

```go
import "github.com/go4x/goal/col/set"
import "github.com/go4x/goal/col/mapx"

// Set operations
mySet := set.New[string]()
mySet.Add("apple").Add("banana").Add("apple") // Duplicates ignored
fmt.Println(mySet.Size()) // 2

// Map operations
myMap := mapx.New[string, int]()
myMap.Put("apple", 1).Put("banana", 2)
value, exists := myMap.Get("apple")
fmt.Println(value, exists) // 1 true
```

### HTTP Client

```go
import "github.com/go4x/goal/httpx"

// Simple HTTP request
response, err := httpx.Get("https://api.example.com/data")
if err != nil {
    log.Fatal(err)
}
defer response.Close()

// Async HTTP request
client := httpx.NewAsyncClient()
future := client.GetAsync("https://api.example.com/data")
response, err := future.Get()
```

### Encryption

```go
import "github.com/go4x/goal/ciphers"

// AES encryption
data := []byte("sensitive data")
key := []byte("your-32-byte-key-here-123456789012")
iv := []byte("random-16-byte-iv")

encrypted, err := ciphers.AES.Encrypt(data, key, ciphers.CBC, iv)
if err != nil {
    log.Fatal(err)
}

decrypted, err := ciphers.AES.Decrypt(encrypted, key, ciphers.CBC, iv)
if err != nil {
    log.Fatal(err)
}
```

### Command Execution

```go
import "github.com/go4x/goal/cmd"

// Execute command with timeout
result, err := cmd.ExecWithTimeout("ls -la", 5*time.Second)
if err != nil {
    log.Fatal(err)
}

// Async command execution
future := cmd.ExecAsync("long-running-command")
result, err := future.Get()
```

## 🏗️ Architecture

### Design Principles

1. **Modularity**: Each package is self-contained and focused on a specific domain
2. **Generics First**: Modern Go generics are used throughout for type safety
3. **Simplicity**: Each package is simple and easy to use
4. **Few Dependencies**: Few external dependencies, easy to maintain
5. **Performance**: Optimized for speed and memory efficiency
6. **Thread Safety**: All packages are designed for concurrent use
7. **Documentation**: Comprehensive documentation with examples

### Package Structure

```
goal/
├── assert/          # Testing assertions
├── ciphers/         # Cryptographic functions
├── cmd/             # Command execution
├── col/             # Collections
│   ├── mapx/        # Map implementations
│   ├── set/         # Set implementations
│   └── slicex/      # Enhanced slice operations
├── color/           # Color manipulation
├── conv/            # Type conversion
├── errorx/          # Error handling
├── httpx/           # HTTP client
├── iox/             # I/O utilities
├── jsonx/           # JSON utilities
├── limiter/         # Rate limiting
├── mathx/           # Mathematical utilities
├── printer/         # Printing utilities
├── prob/            # Probability functions
├── ptr/             # Pointer utilities
├── random/          # Random generation
├── reflectx/        # Reflection utilities
├── retry/           # Retry mechanisms
├── stringx/         # String utilities
├── timex/           # Time utilities
├── uuid/            # UUID generation
└── value/           # Value handling
```

## 🔧 Development

### Requirements

- Go 1.24.0 or later
- Git

### Building

```bash
# Clone the repository
git clone https://github.com/go4x/goal.git
cd goal

# Run tests
go test ./...

# Run benchmarks
go test -bench=. ./...

# Check coverage
go test -cover ./...
```

### Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests for new functionality
5. Ensure all tests pass
6. Submit a pull request

## 📊 Performance

### Benchmarks

The library is designed for high performance:

- **Collections**: O(1) operations for most set/map operations
- **String Operations**: Optimized string manipulation
- **HTTP Client**: Async operations with connection pooling
- **Encryption**: Hardware-accelerated crypto operations

### Memory Usage

- **Efficient**: Minimal memory allocations
- **Pooled**: Connection and buffer pooling where appropriate
- **Immutable**: Immutable operations to prevent side effects

## 📖 Documentation

### Documentation Structure

Each package includes:

- **README.md**: English documentation
- **README_CN.md**: Chinese documentation
- **Examples**: Comprehensive usage examples
- **API Reference**: Complete API documentation
- **Performance Notes**: Performance characteristics and tips

### Getting Help

- **GitHub Issues**: Report bugs and request features
- **Documentation**: Comprehensive package documentation
- **Examples**: Extensive code examples
- **Community**: Join the discussion

## 🤝 Contributing

We welcome contributions! Please see our [Contributing Guide](CONTRIBUTING.md) for details.

### Areas for Contribution

- **New Packages**: Suggest new utility packages
- **Performance**: Optimize existing code
- **Documentation**: Improve documentation
- **Examples**: Add more usage examples
- **Testing**: Improve test coverage

## 📄 License

This project is licensed under the Apache License 2.0 - see the [LICENSE](LICENSE) file for details.

## 🙏 Acknowledgments

- **Go Team**: For the excellent Go programming language
- **Community**: For feedback and contributions
- **Dependencies**: For the excellent third-party packages we use

### Version History

- **v1.0.0**: Initial release with core packages

## 📞 Support

- **GitHub Issues**: [Report Issues](https://github.com/go4x/goal/issues)

---

**Goal** - Making Go development more productive and enjoyable! 🎯
