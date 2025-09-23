package limiter

import (
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTokenBucketStart(t *testing.T) {
	t.Run("start", func(t *testing.T) {
		// max 2 tokens, only consume 1 token every 2 seconds
		limiter := NewTokenBucket(2, 1, time.Second*2)
		limiter.Start()
		defer limiter.Stop()

		time.Sleep(time.Second * 5)
		assert.Equal(t, 2, len(limiter.tokens))
	})
}

func TestTokenBucketTryTake(t *testing.T) {
	limiter := NewTokenBucket(2, 1, time.Second*2)
	limiter.Start()
	defer limiter.Stop()

	time.Sleep(time.Second * 5)

	if !limiter.TryTake() {
		t.Error("should be able to take the first token")
	}
	if !limiter.TryTake() {
		t.Error("should be able to take the second token")
	}
	if limiter.TryTake() {
		t.Error("bucket is empty, should not be able to take the third token")
	}

	time.Sleep(time.Second * 5)
	if !limiter.TryTake() {
		t.Error("should be able to take the new token after 5 seconds")
	}
	if !limiter.TryTake() {
		t.Error("bucket is empty, should not be able to take the third token")
	}
	if limiter.TryTake() {
		t.Error("bucket is empty, should not be able to take the third token")
	}
}

func TestTokenBucketTake(t *testing.T) {
	limiter := NewTokenBucket(1, 1, time.Second)
	limiter.Start()

	time.Sleep(time.Second * 3)

	limiter.Take() // ok
	t.Log("take the first token")
	limiter.Take() // ok
	t.Log("take the second token")

	go func() {
		time.AfterFunc(time.Second*4, func() {
			t.Log("stop limiter after 4 seconds")
			limiter.Stop()
		})
	}()
	t.Log("taking the third token, should be blocked")
	limiter.Take() // blocked
}

func TestTokenBucketTakeWithTimeout(t *testing.T) {
	limiter := NewTokenBucket(1, 1, time.Second*2)
	limiter.Start()
	defer limiter.Stop()

	start := time.Now()
	limiter.TakeWithTimeout(time.Second * 3) // ok
	t.Log("take the first token")
	limiter.TakeWithTimeout(time.Second * 2) // ok
	t.Log("take the second token")
	used := time.Since(start)
	t.Logf("take 2 tokens, used %s", used)
	// need 4 seconds to take 2 tokens
	if used < time.Second*4 {
		t.Errorf("expect 4 seconds, actual %s", used)
	}
}

func TestTokenBucketLimiterConcurrent(t *testing.T) {
	// Create a token bucket for testing: capacity 1, 1 token per second
	limiter := NewTokenBucket(1, 1, time.Second)
	limiter.Start()
	defer limiter.Stop()

	var wg sync.WaitGroup
	successCount := 0
	var mu sync.Mutex

	// Wait for initial token to be available
	time.Sleep(time.Second * 2)

	// Start 10 concurrent requests
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if limiter.TryTake() {
				mu.Lock()
				successCount++
				mu.Unlock()
			}
		}()
	}

	wg.Wait()

	// Since the capacity is 1 and we waited 2 seconds, we should have 2 tokens available
	// But due to timing, it might be 1 or 2 tokens
	if successCount < 1 || successCount > 2 {
		t.Errorf("Expected 1-2 successful requests, got %d", successCount)
	}
}

func TestTokenBucketLimiterStats(t *testing.T) {
	limiter := NewTokenBucket(1, 1, time.Second)
	limiter.Start()
	defer limiter.Stop()

	// reset stats
	limiter.ResetStat()

	time.Sleep(time.Second * 2)

	// try to take tokens
	limiter.TryTake() // success
	limiter.TryTake() // failed
	limiter.TryTake() // failed

	total, blocked, successRate := limiter.Stat()

	if total != 3 {
		t.Errorf("expect 3 total requests, actual %d", total)
	}
	if blocked != 2 {
		t.Errorf("expect 2 blocked requests, actual %d", blocked)
	}
	if successRate != 33.33333333333333 {
		t.Errorf("expect 33.33%% success rate, actual %f%%", successRate)
	}
}

func TestNewTokenBucket(t *testing.T) {
	t.Run("valid parameters", func(t *testing.T) {
		limiter := NewTokenBucket(10, 5, time.Second)
		assert.NotNil(t, limiter)
		assert.Equal(t, 10, cap(limiter.tokens))
		assert.NotNil(t, limiter.ticker)
		assert.NotNil(t, limiter.stop)
	})

	t.Run("zero capacity", func(t *testing.T) {
		limiter := NewTokenBucket(0, 5, time.Second)
		assert.NotNil(t, limiter)
		assert.Equal(t, 0, cap(limiter.tokens))
	})

	t.Run("zero rate", func(t *testing.T) {
		limiter := NewTokenBucket(10, 0, time.Second)
		assert.NotNil(t, limiter)
		// Zero rate should be handled gracefully by setting minimum rate of 1
		// This prevents division by zero in ticker creation
	})
}

func TestTokenBucketEdgeCases(t *testing.T) {
	t.Run("empty bucket initially", func(t *testing.T) {
		limiter := NewTokenBucket(1, 1, time.Second)
		limiter.Start()
		defer limiter.Stop()

		// Initially bucket should be empty
		assert.False(t, limiter.TryTake())
	})

	t.Run("bucket full scenario", func(t *testing.T) {
		limiter := NewTokenBucket(2, 10, time.Second) // High rate
		limiter.Start()
		defer limiter.Stop()

		// Wait for bucket to fill
		time.Sleep(time.Millisecond * 200)

		// Should be able to take tokens (might be 1 or 2 depending on timing)
		successCount := 0
		if limiter.TryTake() {
			successCount++
		}
		if limiter.TryTake() {
			successCount++
		}
		// Third attempt should fail
		assert.False(t, limiter.TryTake())

		// Should have taken at least 1 token
		assert.Greater(t, successCount, 0)
	})

	t.Run("high frequency requests", func(t *testing.T) {
		limiter := NewTokenBucket(1, 1, time.Millisecond*100)
		limiter.Start()
		defer limiter.Stop()

		// Wait for initial token
		time.Sleep(time.Millisecond * 150)

		successCount := 0
		for i := 0; i < 10; i++ {
			if limiter.TryTake() {
				successCount++
			}
		}

		// Should have some successful requests
		assert.Greater(t, successCount, 0)
	})
}

func TestTokenBucketStatistics(t *testing.T) {
	t.Run("initial statistics", func(t *testing.T) {
		limiter := NewTokenBucket(1, 1, time.Second)
		total, blocked, rate := limiter.Stat()
		assert.Equal(t, int64(0), total)
		assert.Equal(t, int64(0), blocked)
		assert.Equal(t, float64(0), rate)
	})

	t.Run("statistics after mixed operations", func(t *testing.T) {
		limiter := NewTokenBucket(1, 1, time.Second)
		limiter.Start()
		defer limiter.Stop()

		// Wait for token
		time.Sleep(time.Millisecond * 1200)

		// Mix of successful and failed operations
		limiter.TryTake() // Success
		limiter.TryTake() // Failure
		limiter.TryTake() // Failure

		total, blocked, rate := limiter.Stat()
		assert.Equal(t, int64(3), total)
		assert.Equal(t, int64(2), blocked)
		assert.InDelta(t, 33.33, rate, 0.1)
	})

	t.Run("reset statistics", func(t *testing.T) {
		limiter := NewTokenBucket(1, 1, time.Second)
		limiter.Start()
		defer limiter.Stop()

		// Generate some statistics
		limiter.TryTake()
		limiter.TryTake()

		// Reset
		limiter.ResetStat()

		total, blocked, rate := limiter.Stat()
		assert.Equal(t, int64(0), total)
		assert.Equal(t, int64(0), blocked)
		assert.Equal(t, float64(0), rate)
	})
}

func TestTokenBucketTimeout(t *testing.T) {
	t.Run("timeout with no tokens", func(t *testing.T) {
		limiter := NewTokenBucket(1, 1, time.Second)
		limiter.Start()
		defer limiter.Stop()

		// Try to take token with very short timeout
		start := time.Now()
		result := limiter.TakeWithTimeout(time.Millisecond * 100)
		duration := time.Since(start)

		assert.False(t, result)
		assert.GreaterOrEqual(t, duration, time.Millisecond*100)
		assert.LessOrEqual(t, duration, time.Millisecond*150)
	})

	t.Run("timeout with available token", func(t *testing.T) {
		limiter := NewTokenBucket(1, 1, time.Millisecond*100)
		limiter.Start()
		defer limiter.Stop()

		// Wait for token to be available
		time.Sleep(time.Millisecond * 150)

		start := time.Now()
		result := limiter.TakeWithTimeout(time.Second)
		duration := time.Since(start)

		assert.True(t, result)
		assert.Less(t, duration, time.Millisecond*50)
	})

	t.Run("zero timeout", func(t *testing.T) {
		limiter := NewTokenBucket(1, 1, time.Second)
		limiter.Start()
		defer limiter.Stop()

		result := limiter.TakeWithTimeout(0)
		assert.False(t, result)
	})
}

func TestTokenBucketConcurrentSafety(t *testing.T) {
	t.Run("concurrent statistics access", func(t *testing.T) {
		limiter := NewTokenBucket(10, 5, time.Millisecond*100)
		limiter.Start()
		defer limiter.Stop()

		// Wait for some tokens
		time.Sleep(time.Millisecond * 200)

		// Concurrent access to statistics
		done := make(chan bool)
		for i := 0; i < 10; i++ {
			go func() {
				for j := 0; j < 10; j++ {
					limiter.TryTake()
					limiter.Stat()
				}
				done <- true
			}()
		}

		// Wait for all goroutines to complete
		for i := 0; i < 10; i++ {
			<-done
		}

		// Statistics should be consistent
		total, blocked, rate := limiter.Stat()
		assert.Greater(t, total, int64(0))
		assert.GreaterOrEqual(t, blocked, int64(0))
		assert.GreaterOrEqual(t, rate, float64(0))
		assert.LessOrEqual(t, rate, float64(100))
	})

	t.Run("concurrent reset and access", func(t *testing.T) {
		limiter := NewTokenBucket(10, 5, time.Millisecond*100)
		limiter.Start()
		defer limiter.Stop()

		// Wait for some tokens
		time.Sleep(time.Millisecond * 200)

		done := make(chan bool)

		// Some goroutines doing operations
		for i := 0; i < 5; i++ {
			go func() {
				for j := 0; j < 10; j++ {
					limiter.TryTake()
				}
				done <- true
			}()
		}

		// Some goroutines resetting statistics
		for i := 0; i < 2; i++ {
			go func() {
				time.Sleep(time.Millisecond * 10)
				limiter.ResetStat()
				done <- true
			}()
		}

		// Wait for all goroutines
		for i := 0; i < 7; i++ {
			<-done
		}

		// Should not panic and should have consistent state
		total, blocked, rate := limiter.Stat()
		assert.GreaterOrEqual(t, total, int64(0))
		assert.GreaterOrEqual(t, blocked, int64(0))
		assert.GreaterOrEqual(t, rate, float64(0))
	})
}

func TestTokenBucketStartStop(t *testing.T) {
	t.Run("multiple start calls", func(t *testing.T) {
		limiter := NewTokenBucket(1, 1, time.Second)

		// Multiple start calls should not cause issues
		limiter.Start()
		limiter.Start()
		limiter.Start()

		defer limiter.Stop()

		// Should still work
		time.Sleep(time.Millisecond * 1200)
		assert.True(t, limiter.TryTake())
	})

	t.Run("stop before start", func(t *testing.T) {
		limiter := NewTokenBucket(1, 1, time.Second)

		// Stop before start should not panic
		limiter.Stop()

		// Should still be able to start after (create new limiter for this test)
		limiter2 := NewTokenBucket(1, 1, time.Second)
		limiter2.Start()
		defer limiter2.Stop()

		// Should work normally
		time.Sleep(time.Millisecond * 1200)
		assert.True(t, limiter2.TryTake())
	})

	t.Run("multiple stop calls", func(t *testing.T) {
		limiter := NewTokenBucket(1, 1, time.Second)
		limiter.Start()

		// Multiple stop calls should not panic
		limiter.Stop()
		limiter.Stop()
		limiter.Stop()
	})
}

func TestTokenBucketRegisterExitHandler(t *testing.T) {
	t.Run("register exit handler", func(t *testing.T) {
		limiter := NewTokenBucket(1, 1, time.Second)

		// Should not panic
		limiter.RegisterExitHandler()

		// Clean up any potential goroutines
		time.Sleep(time.Millisecond * 10)
	})
}

func TestTokenBucketBoundaryValues(t *testing.T) {
	t.Run("negative capacity", func(t *testing.T) {
		limiter := NewTokenBucket(-1, 1, time.Second)
		assert.NotNil(t, limiter)
		assert.Equal(t, 0, cap(limiter.tokens))
	})

	t.Run("negative rate", func(t *testing.T) {
		limiter := NewTokenBucket(10, -1, time.Second)
		assert.NotNil(t, limiter)
		// Should handle gracefully with minimum rate of 1
	})

	t.Run("zero window", func(t *testing.T) {
		limiter := NewTokenBucket(10, 1, 0)
		assert.NotNil(t, limiter)
		// Should handle gracefully with default window of 1 second
	})

	t.Run("very small window", func(t *testing.T) {
		limiter := NewTokenBucket(1, 1, time.Microsecond)
		limiter.Start()
		defer limiter.Stop()

		// Should not panic
		time.Sleep(time.Millisecond)
		limiter.TryTake()
	})

	t.Run("very large capacity", func(t *testing.T) {
		limiter := NewTokenBucket(1000000, 1, time.Second)
		limiter.Start()
		defer limiter.Stop()

		// Should not panic
		time.Sleep(time.Millisecond * 1200)
		assert.True(t, limiter.TryTake())
	})

	t.Run("very high rate", func(t *testing.T) {
		limiter := NewTokenBucket(10, 1000, time.Second)
		limiter.Start()
		defer limiter.Stop()

		// Should not panic
		time.Sleep(time.Millisecond * 10)
		assert.True(t, limiter.TryTake())
	})
}
