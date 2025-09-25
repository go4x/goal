# ErrorX - Enhanced Error Handling for Go

ErrorX is a comprehensive error handling package for Go that provides enhanced error management utilities, particularly useful for web applications and APIs that need to handle HTTP status codes and centralized error processing.

## Features

- **PreferredError**: Custom error type with HTTP status codes for web applications
- **ChainExecutor**: Centralized error handling with fail-fast pattern
- **Panic Recovery**: Utilities with cleanup functions and context support
- **Error Wrapping**: Enhanced error context with formatting support
- **HTTP Error Factories**: Pre-configured errors for common HTTP status codes
- **Panic-based Handling**: Simplified error propagation (use with caution)

## Installation

```bash
go get github.com/go4x/goal/errorx
```

## Quick Start

### Basic Error Handling

```go
package main

import (
    "fmt"
    "github.com/go4x/goal/errorx"
)

func main() {
    // Create preferred errors with HTTP status codes
    err := errorx.Prefer500("service temporarily unavailable")
    if errorx.IsPreferred(err) {
        preferredErr := err.(*errorx.PreferredError)
        fmt.Printf("HTTP Status: %d\n", preferredErr.Code())
    }
}
```

### Centralized Error Handling

```go
package main

import (
    "fmt"
    "github.com/go4x/goal/errorx"
)

func main() {
    // Use error handler for centralized error processing
    executor := errorx.NewChainExecutor()
    
    executor.Do(func() error {
        return riskyOperation1()
    }).Do(func() error {
        return riskyOperation2()
    })
    
    if executor.HasErr() {
        fmt.Printf("Operation failed: %v\n", executor.Err())
    }
}
```

### Panic Recovery

```go
package main

import (
    "github.com/go4x/logx"
    "github.com/go4x/goal/errorx"
)

func main() {
    logger := logx.New()
    
    defer errorx.Recover(logger, func() {
        cleanupResources()
    })
    
    // Your risky operations here
    riskyOperation()
}
```

### Context-aware Panic Recovery

```go
package main

import (
    "context"
    "time"
    
    "github.com/go4x/logx"
    "github.com/go4x/goal/errorx"
)

func main() {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    
    logger := logx.New()
    
    defer errorx.RecoverCtx(ctx, logger, func() {
        cleanupResources()
    })
    
    // Your operations here
}
```

## API Reference

### PreferredError

A custom error type that includes HTTP status codes for web applications.

```go
type PreferredError struct {
    error
    code int
}
```

**Methods:**
- `Code() int` - Returns the HTTP status code
- `Error() string` - Returns the error message

### ChainExecutor

Centralized error handling with fail-fast pattern, context tracking, and result storage.

```go
type ChainExecutor struct {
    err        error
    failedStep string
    results    map[string]interface{}
}
```

**Core Methods:**
- `NewHandler() *ChainExecutor` - Creates a new error handler
- `Do(f func() error) *ChainExecutor` - Executes function if no previous error
- `DoWithContext(step string, f func() error) *ChainExecutor` - Executes with step tracking
- `DoWithResult(step string, f func() (interface{}, error)) *ChainExecutor` - Stores operation results
- `DoWithTypedResult[T any](step string, f func() (T, error)) *ChainExecutor` - Type-safe result storage

**State Methods:**
- `HasErr() bool` - Checks if handler has an error
- `Err() error` - Returns the current error
- `FailedStep() string` - Returns the step where error occurred
- `GetResult(key string) (interface{}, bool)` - Retrieves stored result
- `GetTypedResult[T any](key string) (T, bool)` - Retrieves typed result

**Utility Methods:**
- `Done(f func(err error))` - Adds cleanup hook
- `DoneWithContext(f func(err error, failedStep string))` - Context-aware cleanup
- `IsPreferredErr() bool` - Checks if error is preferred
- `PreferredOr(err error) error` - Returns preferred error or fallback
- `Reset() *ChainExecutor` - Clears state for reuse
- `GetErrorInfo() ErrorInfo` - Returns comprehensive error information

### Recovery Functions

**Recover(logger, cleanups...)**
- Basic panic recovery with cleanup functions

**RecoverCtx(ctx, logger, cleanups...)**
- Context-aware panic recovery with timeout support

### Error Wrapping

**Wrap(err, msg)**
- Wraps error with additional context

**Wrapf(err, format, args...)**
- Wraps error with formatted context

### Panic-based Handling

**Throw(err)**
- Panics if error is not nil

**Throwf(err, format, args...)**
- Panics with custom message if error is not nil

**Throwv(value, err)**
- Returns value or panics if error is not nil

**New(format, args...)**
- Creates new error with formatting

### HTTP Error Factories

Pre-configured error creators for common HTTP status codes:

- `Prefer400(format, args...)` - 400 Bad Request
- `Prefer401(format, args...)` - 401 Unauthorized
- `Prefer403(format, args...)` - 403 Forbidden
- `Prefer415(format, args...)` - 415 Unsupported Media Type
- `Prefer423(format, args...)` - 423 Locked
- `Prefer426(format, args...)` - 426 Upgrade Required
- `Prefer429(format, args...)` - 429 Too Many Requests
- `Prefer500(format, args...)` - 500 Internal Server Error
- `Prefer501(format, args...)` - 501 Not Implemented
- `Prefer503(format, args...)` - 503 Service Unavailable

## Best Practices

### 1. Use ChainExecutor for Sequential Operations

```go
// Basic usage
executor := errorx.NewChainExecutor()
executor.Do(func() error { return validateInput() })
executor.Do(func() error { return processData() })
executor.Do(func() error { return saveToDatabase() })

if executor.HasErr() {
    return executor.Err()
}

// Advanced usage with context tracking
executor := errorx.NewChainExecutor()
executor.DoWithContext("validate", func() error { return validateInput() })
executor.DoWithContext("process", func() error { return processData() })
executor.DoWithContext("save", func() error { return saveToDatabase() })

if executor.HasErr() {
    return fmt.Errorf("failed at %s: %w", executor.FailedStep(), executor.Err())
}
```

### 1.1. Use ChainExecutor with Result Storage

```go
executor := errorx.NewChainExecutor()

// Store intermediate results
executor.DoWithResult("user", func() (interface{}, error) {
    return createUser(userData)
}).DoWithResult("profile", func() (interface{}, error) {
    user, _ := executor.GetResult("user")
    return createProfile(user.(*User).ID)
}).DoWithContext("notify", func() error {
    user, _ := executor.GetResult("user")
    profile, _ := executor.GetResult("profile")
    return sendNotification(user, profile)
})

if executor.HasErr() {
    return executor.Err()
}
```

### 2. Use PreferredError for HTTP APIs

```go
func handleRequest() error {
    if invalidInput {
        return errorx.Prefer400("invalid input format")
    }
    
    if notAuthorized {
        return errorx.Prefer401("authentication required")
    }
    
    return nil
}
```

### 3. Use Recover for Resource Cleanup

```go
func riskyOperation() {
    defer errorx.Recover(logger, func() {
        closeConnections()
        releaseResources()
    })
    
    // Risky operations
}
```

### 4. Use Context-aware Recovery for Timeouts

```go
func operationWithTimeout() {
    ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
    defer cancel()
    
    defer errorx.RecoverCtx(ctx, logger, cleanup)
    
    // Long-running operations
}
```

## Testing

Run the test suite:

```bash
go test -v ./errorx
```

Run tests with coverage:

```bash
go test -cover ./errorx
```

Run benchmarks:

```bash
go test -bench=. ./errorx
```

## Examples

See the `example_test.go` file for comprehensive usage examples.

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests for new functionality
5. Ensure all tests pass
6. Submit a pull request

## License

This project is licensed under the Apache License 2.0 - see the LICENSE file for details.

## Changelog

### v1.0.0
- Initial release
- PreferredError with HTTP status codes
- ChainExecutor for centralized error processing
- Panic recovery utilities
- Error wrapping functions
- HTTP error factory functions
- Comprehensive test coverage
- Full documentation and examples