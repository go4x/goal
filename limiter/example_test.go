package limiter_test

import (
	"fmt"
	"time"

	"github.com/go4x/goal/limiter"
)

// ExampleTokenBucket demonstrates basic usage of the token bucket rate limiter.
func ExampleTokenBucket() {
	// Create a token bucket with capacity 10, generating 5 tokens per second
	limiter := limiter.NewTokenBucket(10, 5, time.Second)
	limiter.Start()
	defer limiter.Stop()

	// Simulate some requests
	for i := 0; i < 15; i++ {
		if limiter.TryTake() {
			fmt.Printf("Request %d: Allowed\n", i+1)
		} else {
			fmt.Printf("Request %d: Rate limited\n", i+1)
		}
		time.Sleep(time.Millisecond * 100)
	}
}

// ExampleTokenBucket_blocking demonstrates blocking token acquisition.
func ExampleTokenBucket_blocking() {
	// Create a token bucket with capacity 2, generating 1 token per second
	limiter := limiter.NewTokenBucket(2, 1, time.Second)
	limiter.Start()
	defer limiter.Stop()

	// Wait for initial tokens
	time.Sleep(time.Millisecond * 1200)

	// First two requests should succeed immediately
	for i := 0; i < 2; i++ {
		start := time.Now()
		limiter.Take() // This will block until a token is available
		duration := time.Since(start)
		fmt.Printf("Request %d took %v\n", i+1, duration)
	}

	// Third request will wait for the next token
	start := time.Now()
	limiter.Take()
	duration := time.Since(start)
	fmt.Printf("Request 3 took %v (waited for next token)\n", duration)
}

// ExampleTokenBucket_timeout demonstrates timeout-based token acquisition.
func ExampleTokenBucket_timeout() {
	// Create a token bucket with capacity 1, generating 1 token per 2 seconds
	limiter := limiter.NewTokenBucket(1, 1, time.Second*2)
	limiter.Start()
	defer limiter.Stop()

	// Wait for initial token
	time.Sleep(time.Millisecond * 2100)

	// Try to get token with timeout
	if limiter.TakeWithTimeout(time.Second) {
		fmt.Println("Got token within timeout")
	} else {
		fmt.Println("Timeout waiting for token")
	}

	// Try again with shorter timeout
	if limiter.TakeWithTimeout(time.Millisecond * 100) {
		fmt.Println("Got token within short timeout")
	} else {
		fmt.Println("Timeout with short duration")
	}
}

// ExampleTokenBucket_statistics demonstrates statistics tracking.
func ExampleTokenBucket_statistics() {
	// Create a token bucket with capacity 5, generating 2 tokens per second
	limiter := limiter.NewTokenBucket(5, 2, time.Second)
	limiter.Start()
	defer limiter.Stop()

	// Wait for initial tokens
	time.Sleep(time.Millisecond * 3000)

	// Make some requests
	for i := 0; i < 10; i++ {
		limiter.TryTake()
		time.Sleep(time.Millisecond * 100)
	}

	// Check statistics
	total, blocked, successRate := limiter.Stat()
	fmt.Printf("Total requests: %d\n", total)
	fmt.Printf("Blocked requests: %d\n", blocked)
	fmt.Printf("Success rate: %.2f%%\n", successRate)

	// Reset statistics
	limiter.ResetStat()
	total, blocked, successRate = limiter.Stat()
	fmt.Printf("After reset - Total: %d, Blocked: %d, Success rate: %.2f%%\n",
		total, blocked, successRate)
}

// ExampleTokenBucket_interface demonstrates using the limiter through the interface.
func ExampleTokenBucket_interface() {
	// Use the Limiter interface for polymorphism
	var rateLimiter limiter.Limiter = limiter.NewTokenBucket(3, 1, time.Second)
	rateLimiter.Start()
	defer rateLimiter.Stop()

	// Wait for initial token
	time.Sleep(time.Millisecond * 1100)

	// Use interface methods
	if rateLimiter.TryTake() {
		fmt.Println("Token acquired through interface")
	}

	if rateLimiter.TakeWithTimeout(time.Millisecond * 500) {
		fmt.Println("Token acquired with timeout through interface")
	}

	total, blocked, rate := rateLimiter.Stat()
	fmt.Printf("Interface stats - Total: %d, Blocked: %d, Rate: %.2f%%\n",
		total, blocked, rate)
}

// ExampleTokenBucket_concurrent demonstrates concurrent usage.
func ExampleTokenBucket_concurrent() {
	// Create a token bucket with capacity 10, generating 5 tokens per second
	limiter := limiter.NewTokenBucket(10, 5, time.Second)
	limiter.Start()
	defer limiter.Stop()

	// Wait for initial tokens
	time.Sleep(time.Millisecond * 2500)

	// Simulate concurrent requests
	results := make(chan string, 20)

	for i := 0; i < 20; i++ {
		go func(id int) {
			if limiter.TryTake() {
				results <- fmt.Sprintf("Goroutine %d: Success", id)
			} else {
				results <- fmt.Sprintf("Goroutine %d: Rate limited", id)
			}
		}(i)
	}

	// Collect results
	for i := 0; i < 20; i++ {
		fmt.Println(<-results)
	}
}
