package limiter

import (
	"sync"
	"testing"
	"time"
)

func BenchmarkTokenBucketTryTake(b *testing.B) {
	limiter := NewTokenBucket(1000, 1000, time.Second)
	limiter.Start()
	defer limiter.Stop()

	// Wait for bucket to fill
	time.Sleep(time.Millisecond * 100)

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			limiter.TryTake()
		}
	})
}

func BenchmarkTokenBucketTakeWithTimeout(b *testing.B) {
	limiter := NewTokenBucket(1000, 1000, time.Second)
	limiter.Start()
	defer limiter.Stop()

	// Wait for bucket to fill
	time.Sleep(time.Millisecond * 100)

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			limiter.TakeWithTimeout(time.Millisecond)
		}
	})
}

func BenchmarkTokenBucketStat(b *testing.B) {
	limiter := NewTokenBucket(100, 10, time.Second)
	limiter.Start()
	defer limiter.Stop()

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			limiter.Stat()
		}
	})
}

func BenchmarkTokenBucketConcurrentOperations(b *testing.B) {
	limiter := NewTokenBucket(100, 10, time.Second)
	limiter.Start()
	defer limiter.Stop()

	// Wait for initial tokens
	time.Sleep(time.Millisecond * 100)

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			// Mix of different operations
			if pb.Next() {
				limiter.TryTake()
			} else {
				limiter.Stat()
			}
		}
	})
}

func BenchmarkTokenBucketHighContention(b *testing.B) {
	// Small bucket with high contention
	limiter := NewTokenBucket(1, 1, time.Second)
	limiter.Start()
	defer limiter.Stop()

	// Wait for initial token
	time.Sleep(time.Millisecond * 1200)

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			limiter.TryTake()
		}
	})
}

func BenchmarkTokenBucketLargeCapacity(b *testing.B) {
	// Large bucket with low contention
	limiter := NewTokenBucket(100000, 1000, time.Second)
	limiter.Start()
	defer limiter.Stop()

	// Wait for bucket to fill
	time.Sleep(time.Millisecond * 100)

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			limiter.TryTake()
		}
	})
}

func BenchmarkTokenBucketMixedWorkload(b *testing.B) {
	limiter := NewTokenBucket(100, 10, time.Second)
	limiter.Start()
	defer limiter.Stop()

	// Wait for initial tokens
	time.Sleep(time.Millisecond * 100)

	b.ResetTimer()

	var wg sync.WaitGroup
	goroutines := 10

	for i := 0; i < goroutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < b.N/goroutines; i++ {
				// Mix of operations
				switch i % 4 {
				case 0:
					limiter.TryTake()
				case 1:
					limiter.TakeWithTimeout(time.Millisecond)
				case 2:
					limiter.Stat()
				case 3:
					limiter.ResetStat()
				}
			}
		}()
	}

	wg.Wait()
}

func BenchmarkTokenBucketCreation(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		limiter := NewTokenBucket(100, 10, time.Second)
		_ = limiter
	}
}

func BenchmarkTokenBucketStartStop(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		limiter := NewTokenBucket(100, 10, time.Second)
		limiter.Start()
		limiter.Stop()
	}
}
