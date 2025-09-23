// Package limiter provides rate limiting functionality using the token bucket algorithm.
// The token bucket algorithm allows bursts of traffic up to the bucket capacity,
// while maintaining a steady rate of token generation.
package limiter

import (
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

// TokenBucket implements a token bucket rate limiter.
// It controls the rate of requests by maintaining a bucket of tokens.
// Tokens are added to the bucket at a steady rate, and requests consume tokens.
// When the bucket is empty, requests are either blocked or rejected.
//
// The limiter provides both blocking and non-blocking methods for acquiring tokens,
// along with statistics tracking for monitoring purposes.
//
// TokenBucket implements the Limiter interface, providing:
//   - TryTake() bool: Non-blocking token acquisition
//   - Take(): Blocking token acquisition
//   - TakeWithTimeout(timeout time.Duration) bool: Timeout-based token acquisition
//   - Stat() (total, blocked int64, successRate float64): Statistics retrieval
type TokenBucket struct {
	tokens chan struct{} // Channel that holds available tokens
	ticker *time.Ticker  // Timer that generates new tokens at regular intervals
	stop   chan struct{} // Channel used to signal the limiter to stop

	// Statistics tracking (protected by mu)
	mu              sync.Mutex // Mutex protecting statistics fields
	totalRequests   int64      // Total number of requests made
	blockedRequests int64      // Number of requests that were blocked/rejected
	lastResetTime   time.Time  // Time when statistics were last reset
}

// Start begins the token generation process in a separate goroutine.
// This method starts a background process that periodically adds tokens to the bucket
// at the configured rate. The process continues until Stop() is called.
//
// The token generation follows this logic:
//   - Tokens are added to the bucket at regular intervals based on the ticker
//   - If the bucket is full, new tokens are dropped (bucket doesn't overflow)
//   - The process can be stopped by sending a signal to the stop channel
func (l *TokenBucket) Start() {
	go func() {
		for {
			select {
			case <-l.ticker.C:
				select {
				case l.tokens <- struct{}{}:
					// Successfully added a token to the bucket
				default:
					// Bucket is full, drop the token
				}
			case <-l.stop:
				return
			}
		}
	}()
}

// RegisterExitHandler sets up signal handling for graceful shutdown.
// This method registers handlers for SIGINT and SIGTERM signals,
// allowing the limiter to be stopped gracefully when the application receives these signals.
//
// When a signal is received:
//   - The limiter is stopped using Stop()
//   - The application exits with code 0
//
// This is useful for ensuring proper cleanup when the application is terminated.
func (l *TokenBucket) RegisterExitHandler() {
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)
		<-c
		l.Stop()
		os.Exit(0)
	}()
}

// Stop gracefully shuts down the token bucket limiter.
// This method:
//   - Stops the ticker that generates new tokens
//   - Closes the stop channel to signal the background goroutine to exit
//
// After calling Stop(), no new tokens will be generated,
// but existing tokens in the bucket can still be consumed.
//
// This method is safe to call multiple times.
func (l *TokenBucket) Stop() {
	l.ticker.Stop()

	// Use select to avoid closing an already closed channel
	select {
	case <-l.stop:
		// Channel already closed
	default:
		close(l.stop)
	}
}

// TryTake attempts to acquire a token without blocking.
// This is a non-blocking operation that immediately returns whether a token was available.
//
// Returns:
//   - true: A token was successfully acquired
//   - false: No token was available (bucket is empty)
//
// This method also updates the request statistics:
//   - Increments totalRequests counter
//   - Increments blockedRequests counter if no token was available
func (l *TokenBucket) TryTake() bool {
	l.mu.Lock()
	l.totalRequests++
	l.mu.Unlock()

	select {
	case <-l.tokens:
		return true
	default:
		l.mu.Lock()
		l.blockedRequests++
		l.mu.Unlock()
		return false
	}
}

// Take acquires a token, blocking until one becomes available.
// This method will wait indefinitely until a token is available in the bucket.
// It should be used when you need to ensure the request is processed
// regardless of how long it takes to get a token.
//
// Note: This method does not update statistics automatically.
// If you need statistics tracking, use TryTake() or TakeWithTimeout() instead.
func (l *TokenBucket) Take() {
	<-l.tokens
}

// TakeWithTimeout attempts to acquire a token within the specified timeout duration.
// This method provides a balance between blocking and non-blocking behavior.
//
// Parameters:
//   - timeout: Maximum time to wait for a token to become available
//
// Returns:
//   - true: A token was successfully acquired within the timeout
//   - false: Timeout occurred before a token became available
//
// This method is useful when you want to limit how long you're willing to wait
// for rate limiting, but still prefer to process the request if possible.
func (l *TokenBucket) TakeWithTimeout(timeout time.Duration) bool {
	select {
	case <-l.tokens:
		return true
	case <-time.After(timeout):
		return false
	}
}

// Stat retrieves the current statistics for the token bucket limiter.
// This method provides insights into the limiter's performance and usage patterns.
//
// Returns:
//   - total: Total number of requests made (including both successful and blocked)
//   - blocked: Number of requests that were blocked due to no available tokens
//   - successRate: Percentage of requests that were successful (0.0 to 100.0)
//
// The success rate is calculated as: (total - blocked) / total * 100
// If no requests have been made, the success rate will be 0.
func (l *TokenBucket) Stat() (total, blocked int64, successRate float64) {
	l.mu.Lock()
	defer l.mu.Unlock()

	if l.totalRequests > 0 {
		successRate = float64(l.totalRequests-l.blockedRequests) / float64(l.totalRequests) * 100
	}

	return l.totalRequests, l.blockedRequests, successRate
}

// ResetStat clears all statistics and resets the counters to zero.
// This method is useful for starting fresh measurements or clearing old data.
//
// After calling this method:
//   - totalRequests is reset to 0
//   - blockedRequests is reset to 0
//   - lastResetTime is set to the current time
//
// This allows you to track statistics for specific time periods
// or reset counters after maintenance operations.
func (l *TokenBucket) ResetStat() {
	l.mu.Lock()
	defer l.mu.Unlock()

	l.totalRequests = 0
	l.blockedRequests = 0
	l.lastResetTime = time.Now()
}

// NewTokenBucket creates a new token bucket limiter with the specified parameters.
// This constructor sets up the limiter but does not start the token generation process.
// You must call Start() to begin token generation.
//
// Parameters:
//   - capacity: Maximum number of tokens the bucket can hold (burst capacity)
//   - rate: Number of tokens to generate per time window
//   - window: Time window for token generation (e.g., 1 second, 1 minute)
//
// The token generation interval is calculated as: window / rate
// For example, with rate=10 and window=1 second, tokens are generated every 100ms
//
// Returns:
//   - *TokenBucket: A new limiter instance (not started)
//
// Example:
//
//	// Basic usage
//	limiter := NewTokenBucket(100, 10, time.Second) // 100 capacity, 10 tokens/sec
//	limiter.Start()
//
//	// Method chaining
//	limiter := NewTokenBucket(100, 10, time.Second).Start()
func NewTokenBucket(capacity int, rate int, window time.Duration) *TokenBucket {
	// Handle edge cases
	if capacity < 0 {
		capacity = 0
	}
	if rate <= 0 {
		rate = 1 // Minimum rate of 1 to avoid division by zero
	}
	if window <= 0 {
		window = time.Second // Default window
	}

	return &TokenBucket{
		tokens: make(chan struct{}, capacity),
		ticker: time.NewTicker(window / time.Duration(rate)),
		stop:   make(chan struct{}),
	}
}
