# Retry Package

A flexible and powerful retry mechanism for Go applications with configurable retry strategies, intervals, and callbacks.

## Features

- **Flexible Retry Logic**: Support for custom retry conditions and error handling
- **Configurable Intervals**: Built-in exponential backoff and constant interval strategies
- **Extensible Design**: Easy to implement custom interval strategies
- **Callback Support**: Monitor retry attempts with custom callback functions
- **Functional Options**: Clean API using functional options pattern
- **Type Safety**: Fully typed with comprehensive error handling

## Installation

```bash
go get github.com/go4x/goal/retry
```

## Quick Start

### Basic Usage

```go
package main

import (
    "errors"
    "fmt"
    "time"
    
    "github.com/go4x/goal/retry"
)

func main() {
    // Simple retry with exponential backoff
    err := retry.Do(func() (bool, error) {
        // Your operation here
        return false, errors.New("temporary failure")
    }, retry.Times(3))
    
    if err != nil {
        fmt.Printf("Operation failed after retries: %v\n", err)
    }
}
```

### Advanced Usage

```go
// Retry with custom interval and callback
err := retry.Do(func() (bool, error) {
    // Your operation here
    return false, someError
}, 
    retry.Times(5),
    retry.Interval(retry.ConstantInterval(100 * time.Millisecond)),
    retry.Callback(func(n uint, err error) {
        log.Printf("Retry attempt %d failed: %v", n, err)
    }),
)
```

## API Reference

### Core Types

#### `F` Function Type
```go
type F func() (bool, error)
```

The function type to be retried. Returns:
- `bool`: `true` to stop retrying, `false` to continue retrying
- `error`: `nil` for success, non-nil for failure

**Retry Behavior:**
- If error is `nil`, retry stops immediately regardless of the bool value
- If error is not `nil`, the bool value determines whether to continue retrying

#### `Intervaler` Interface
```go
type Intervaler interface {
    Interval(n uint)
}
```

Defines the interface for retry interval strategies. Implementations determine how long to wait between retry attempts.

### Configuration Functions

#### `Times(n uint)`
Sets the maximum number of retry attempts.

```go
retry.Times(3) // Retry up to 3 times
```

#### `Interval(s Intervaler)`
Sets the interval strategy for determining wait time between retries.

```go
// Use constant interval
retry.Interval(retry.ConstantInterval(time.Second))

// Use default exponential backoff
retry.Interval(retry.DefaultInterval())
```

#### `Callback(c func(n uint, err error))`
Sets a callback function called after each retry attempt.

```go
retry.Callback(func(n uint, err error) {
    log.Printf("Retry attempt %d failed: %v", n, err)
})
```

### Main Function

#### `Do(f F, pf ...settings) error`
Executes the given function with retry logic.

**Parameters:**
- `f`: The function to be retried
- `pf`: Optional settings to configure retry behavior

**Returns:**
- `error`: `nil` if the function succeeded, or the last error if all retries failed

## Built-in Interval Strategies

### Default Exponential Backoff
```go
retry.DefaultInterval()
```

Implements exponential backoff with jitter: sleeps for 2^n seconds plus random jitter between retry attempts to prevent thundering herd problems.

### Constant Interval
```go
retry.ConstantInterval(duration)
```

Sleeps for the same duration between all retry attempts.

```go
retry.ConstantInterval(100 * time.Millisecond)
```

### Exponential Backoff with Jitter
```go
retry.ExponentialBackoffWithJitter(base, jitter)
```

Implements exponential backoff with configurable jitter to prevent thundering herd problems.

**Parameters:**
- `base`: Base duration for the first retry
- `jitter`: Jitter factor (0.0 to 1.0), where 0.0 = no jitter, 1.0 = 100% jitter

```go
// 10% jitter
retry.ExponentialBackoffWithJitter(time.Second, 0.1)

// 50% jitter  
retry.ExponentialBackoffWithJitter(time.Second, 0.5)
```

## Custom Interval Strategies

You can implement your own interval strategies by implementing the `Intervaler` interface:

```go
type customInterval struct {
    base time.Duration
}

func (c *customInterval) Interval(n uint) {
    // Custom logic here
    time.Sleep(c.base * time.Duration(n))
}

// Usage
err := retry.Do(func() (bool, error) {
    return false, someError
}, 
    retry.Times(3),
    retry.Interval(&customInterval{base: 100 * time.Millisecond}),
)
```

## Examples

### HTTP Request Retry
```go
func makeHTTPRequest() (bool, error) {
    resp, err := http.Get("https://api.example.com/data")
    if err != nil {
        return false, err // Continue retrying
    }
    defer resp.Body.Close()
    
    if resp.StatusCode >= 500 {
        return false, fmt.Errorf("server error: %d", resp.StatusCode)
    }
    
    if resp.StatusCode >= 400 {
        return true, fmt.Errorf("client error: %d", resp.StatusCode) // Stop retrying
    }
    
    return true, nil // Success
}

err := retry.Do(makeHTTPRequest, 
    retry.Times(5),
    retry.Interval(retry.ConstantInterval(1 * time.Second)),
    retry.Callback(func(n uint, err error) {
        log.Printf("HTTP request attempt %d failed: %v", n, err)
    }),
)
```

### Database Operation Retry
```go
func saveToDatabase(data interface{}) (bool, error) {
    err := db.Save(data)
    if err != nil {
        // Check if it's a retryable error
        if isRetryableError(err) {
            return false, err // Continue retrying
        }
        return true, err // Stop retrying
    }
    return true, nil // Success
}

err := retry.Do(func() (bool, error) {
    return saveToDatabase(myData)
}, retry.Times(3))
```

### Conditional Retry Logic
```go
func processWithCondition() (bool, error) {
    result, err := someOperation()
    if err != nil {
        // Only retry for specific errors
        if errors.Is(err, ErrTemporary) {
            return false, err
        }
        return true, err // Don't retry for permanent errors
    }
    
    // Check result and decide whether to retry
    if result.NeedsRetry {
        return false, errors.New("result needs retry")
    }
    
    return true, nil // Success
}
```

## Best Practices

1. **Choose Appropriate Retry Count**: Don't retry too many times for operations that are likely to fail permanently
2. **Use Exponential Backoff**: For distributed systems, use exponential backoff to avoid thundering herd problems
3. **Handle Different Error Types**: Distinguish between retryable and non-retryable errors
4. **Set Timeouts**: Consider overall timeout for retry operations
5. **Log Retry Attempts**: Use callbacks to monitor and log retry behavior
6. **Test Retry Logic**: Ensure your retry logic works correctly in various scenarios

## Error Handling

The retry package returns the last error encountered if all retries are exhausted. Make sure to handle this appropriately:

```go
err := retry.Do(myFunction, retry.Times(3))
if err != nil {
    // Handle the final error
    log.Printf("Operation failed after 3 retries: %v", err)
}
```

## Performance Considerations

- **Memory Usage**: The package is lightweight with minimal memory overhead
- **Goroutine Safety**: All functions are safe for concurrent use
- **CPU Usage**: Exponential backoff helps reduce CPU usage during retries
- **Network Efficiency**: Use appropriate intervals to avoid overwhelming remote services

## License

This package is part of the goal project. See the main project license for details.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.
