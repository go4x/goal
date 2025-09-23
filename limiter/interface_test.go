package limiter

import (
	"testing"
	"time"
)

// TestTokenBucketImplementsLimiter verifies that TokenBucket implements the Limiter interface.
// This test validates the interface implementation at compile time.
func TestTokenBucketImplementsLimiter(t *testing.T) {
	// Compile-time interface implementation check
	var _ Limiter = &TokenBucket{}

	// Create instance and test all interface methods
	bucket := NewTokenBucket(10, 5, time.Second)

	// Test TryTake
	_ = bucket.TryTake()

	// Test Take (in goroutine to avoid blocking)
	go func() {
		bucket.Take()
	}()

	// Test TakeWithTimeout
	_ = bucket.TakeWithTimeout(100 * time.Millisecond)

	// Test Stat
	total, blocked, successRate := bucket.Stat()
	_ = total
	_ = blocked
	_ = successRate

	t.Log("TokenBucket successfully implements Limiter interface")
}

// TestLimiterInterface tests interface polymorphism.
func TestLimiterInterface(t *testing.T) {
	// Use interface type
	var limiter Limiter = NewTokenBucket(10, 5, time.Second)

	// Test all interface methods
	_ = limiter.TryTake()
	_ = limiter.TakeWithTimeout(100 * time.Millisecond)
	total, blocked, successRate := limiter.Stat()

	_ = total
	_ = blocked
	_ = successRate

	t.Log("Interface polymorphism works correctly")
}
