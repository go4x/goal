# Limiter Package

A high-performance, thread-safe rate limiting library for Go using the token bucket algorithm.

## Features

- **Token Bucket Algorithm**: Allows bursts of traffic up to bucket capacity while maintaining steady token generation
- **Thread-Safe**: Safe for concurrent use across multiple goroutines
- **Flexible API**: Supports blocking, non-blocking, and timeout-based token acquisition
- **Statistics Tracking**: Built-in monitoring and observability features
- **Interface-Based Design**: Clean abstraction allowing for different rate limiting implementations
- **Graceful Shutdown**: Proper cleanup and signal handling

## Installation

```bash
go get github.com/go4x/goal/limiter
```

## Quick Start

```go
package main

import (
    "fmt"
    "time"
    "github.com/go4x/goal/limiter"
)

func main() {
    // Create a token bucket: capacity=10, rate=5 tokens/second
    limiter := limiter.NewTokenBucket(10, 5, time.Second)
    limiter.Start()
    defer limiter.Stop()

    // Make requests
    for i := 0; i < 15; i++ {
        if limiter.TryTake() {
            fmt.Printf("Request %d: Allowed\n", i+1)
        } else {
            fmt.Printf("Request %d: Rate limited\n", i+1)
        }
        time.Sleep(time.Millisecond * 100)
    }
}
```

## API Reference

### Limiter Interface

The `Limiter` interface provides a common contract for rate limiting implementations:

```go
type Limiter interface {
    Start()                                                    // Start the limiter
    TryTake() bool                                           // Non-blocking token acquisition
    Take()                                                   // Blocking token acquisition
    TakeWithTimeout(timeout time.Duration) bool              // Timeout-based acquisition
    Stat() (total, blocked int64, successRate float64)       // Get statistics
    Stop()                                                   // Stop the limiter
}
```

### TokenBucket

The `TokenBucket` struct implements the `Limiter` interface using the token bucket algorithm.

#### Constructor

```go
func NewTokenBucket(capacity int, rate int, window time.Duration) *TokenBucket
```

**Parameters:**
- `capacity`: Maximum number of tokens the bucket can hold (burst capacity)
- `rate`: Number of tokens to generate per time window
- `window`: Time window for token generation (e.g., 1 second, 1 minute)

**Example:**
```go
// 100 capacity, 10 tokens per second
limiter := limiter.NewTokenBucket(100, 10, time.Second)
```

#### Methods

##### Start()
Starts the token generation process in a background goroutine.

```go
limiter.Start()
```

##### TryTake() bool
Attempts to acquire a token without blocking. Returns `true` if successful, `false` if no token was available.

```go
if limiter.TryTake() {
    // Process request
} else {
    // Handle rate limiting
}
```

##### Take()
Acquires a token, blocking until one becomes available.

```go
limiter.Take() // Will wait until token is available
// Process request
```

##### TakeWithTimeout(timeout time.Duration) bool
Attempts to acquire a token within the specified timeout. Returns `true` if successful, `false` if timeout occurred.

```go
if limiter.TakeWithTimeout(time.Second) {
    // Process request
} else {
    // Handle timeout
}
```

##### Stat() (total, blocked int64, successRate float64)
Returns current statistics:
- `total`: Total number of requests made
- `blocked`: Number of blocked requests
- `successRate`: Success rate percentage (0.0 to 100.0)

```go
total, blocked, rate := limiter.Stat()
fmt.Printf("Total: %d, Blocked: %d, Success Rate: %.2f%%\n", total, blocked, rate)
```

##### Stop()
Stops the token generation process.

```go
limiter.Stop()
```

##### ResetStat()
Resets all statistics counters to zero.

```go
limiter.ResetStat()
```

## Usage Examples

### Basic Rate Limiting

```go
limiter := limiter.NewTokenBucket(10, 5, time.Second)
limiter.Start()
defer limiter.Stop()

// Non-blocking approach
if limiter.TryTake() {
    processRequest()
} else {
    handleRateLimit()
}
```

### Blocking Rate Limiting

```go
limiter := limiter.NewTokenBucket(5, 2, time.Second)
limiter.Start()
defer limiter.Stop()

// Blocking approach - will wait for token
limiter.Take()
processRequest()
```

### Timeout-Based Rate Limiting

```go
limiter := limiter.NewTokenBucket(3, 1, time.Second)
limiter.Start()
defer limiter.Stop()

// Timeout approach - wait up to 5 seconds
if limiter.TakeWithTimeout(5 * time.Second) {
    processRequest()
} else {
    handleTimeout()
}
```

### Statistics Monitoring

```go
limiter := limiter.NewTokenBucket(100, 10, time.Second)
limiter.Start()
defer limiter.Stop()

// Make some requests...

// Check statistics
total, blocked, rate := limiter.Stat()
fmt.Printf("Success rate: %.2f%%\n", rate)

// Reset statistics for new measurement period
limiter.ResetStat()
```

### Interface Usage

```go
// Use interface for polymorphism
var rateLimiter limiter.Limiter = limiter.NewTokenBucket(10, 5, time.Second)
rateLimiter.Start()
defer rateLimiter.Stop()

// All interface methods work the same way
if rateLimiter.TryTake() {
    // Process request
}
```

### Concurrent Usage

```go
limiter := limiter.NewTokenBucket(10, 5, time.Second)
limiter.Start()
defer limiter.Stop()

var wg sync.WaitGroup
for i := 0; i < 100; i++ {
    wg.Add(1)
    go func() {
        defer wg.Done()
        if limiter.TryTake() {
            processRequest()
        }
    }()
}
wg.Wait()
```

## Performance

The limiter is designed for high-performance scenarios:

- **Low Latency**: Non-blocking operations have minimal overhead
- **High Throughput**: Optimized for concurrent access patterns
- **Memory Efficient**: Uses channels for token management
- **Thread-Safe**: All operations are safe for concurrent use

### Benchmark Results

Run benchmarks to see performance characteristics:

```bash
go test -bench=. ./limiter
```

## Error Handling

The limiter handles edge cases gracefully:

- **Zero/Negative Parameters**: Automatically corrected to safe defaults
- **Multiple Start/Stop Calls**: Safe to call multiple times
- **Channel Operations**: Protected against panic conditions

## Best Practices

1. **Always call Start()**: The limiter won't generate tokens until started
2. **Clean up resources**: Always call Stop() when done
3. **Monitor statistics**: Use Stat() to track performance
4. **Choose appropriate capacity**: Balance between burst allowance and memory usage
5. **Use timeouts**: Prefer TakeWithTimeout() over blocking Take() for better control

## Thread Safety

All methods are thread-safe and can be called concurrently from multiple goroutines. The implementation uses:

- Mutex for statistics protection
- Channels for thread-safe token management
- Atomic operations where appropriate

## License

This package is part of the goal project and follows the same license terms.
