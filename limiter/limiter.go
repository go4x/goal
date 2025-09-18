package limiter

import "time"

// Limiter defines the interface for rate limiting implementations.
// This interface provides a common contract for different rate limiting algorithms,
// allowing clients to use any limiter implementation without being coupled to specific types.
//
// The interface supports both blocking and non-blocking operations for acquiring permits,
// along with statistics tracking for monitoring and observability purposes.
//
// Implementations should be thread-safe and suitable for concurrent use.
type Limiter interface {
	// Start starts the limiter.
	Start()

	// TryTake attempts to acquire a permit without blocking.
	// This is a non-blocking operation that immediately returns whether a permit was available.
	//
	// Returns:
	//   - true: A permit was successfully acquired
	//   - false: No permit was available (rate limit exceeded)
	//
	// This method should update internal statistics for monitoring purposes.
	TryTake() bool

	// Take acquires a permit, blocking until one becomes available.
	// This method will wait indefinitely until a permit is available.
	// It should be used when you need to ensure the request is processed
	// regardless of how long it takes to get a permit.
	//
	// Note: This method may not update statistics automatically.
	// Check the specific implementation for statistics behavior.
	Take()

	// TakeWithTimeout attempts to acquire a permit within the specified timeout duration.
	// This method provides a balance between blocking and non-blocking behavior.
	//
	// Parameters:
	//   - timeout: Maximum time to wait for a permit to become available
	//
	// Returns:
	//   - true: A permit was successfully acquired within the timeout
	//   - false: Timeout occurred before a permit became available
	//
	// This method is useful when you want to limit how long you're willing to wait
	// for rate limiting, but still prefer to process the request if possible.
	TakeWithTimeout(timeout time.Duration) bool

	// Stat retrieves the current statistics for the rate limiter.
	// This method provides insights into the limiter's performance and usage patterns.
	//
	// Returns:
	//   - total: Total number of requests made (including both successful and blocked)
	//   - blocked: Number of requests that were blocked due to rate limiting
	//   - successRate: Percentage of requests that were successful (0.0 to 100.0)
	//
	// The success rate is calculated as: (total - blocked) / total * 100
	// If no requests have been made, the success rate will be 0.
	//
	// This method should be thread-safe and provide a consistent snapshot
	// of the current statistics state.
	Stat() (total, blocked int64, successRate float64)

	// Stop stops the limiter.
	Stop()
}
