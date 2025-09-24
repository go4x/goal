package retry_test

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/go4x/goal/retry"
)

// ExampleDo demonstrates basic retry functionality
func ExampleDo() {
	attempts := 0

	// Define a function that will fail twice then succeed
	f := retry.F(func() (bool, error) {
		attempts++
		fmt.Printf("Attempt %d\n", attempts)

		if attempts < 3 {
			return false, errors.New("temporary failure")
		}
		return true, nil // Success
	})

	// Retry up to 5 times with default exponential backoff
	err := retry.Do(f, retry.Times(5))
	if err != nil {
		log.Printf("Operation failed: %v", err)
	} else {
		fmt.Printf("Operation succeeded after %d attempts\n", attempts)
	}
	// Output:
	// Attempt 1
	// Attempt 2
	// Attempt 3
	// Operation succeeded after 3 attempts
}

// ExampleDo_withCallback demonstrates retry with callback
func ExampleDo_withCallback() {
	attempts := 0

	f := retry.F(func() (bool, error) {
		attempts++
		if attempts < 2 {
			return false, errors.New("network error")
		}
		return true, nil
	})

	// Retry with callback to monitor attempts
	err := retry.Do(f,
		retry.Times(3),
		retry.Callback(func(n uint, err error) {
			fmt.Printf("Retry attempt %d failed: %v\n", n, err)
		}),
	)

	if err != nil {
		log.Printf("Final error: %v", err)
	} else {
		fmt.Printf("Success after %d attempts\n", attempts)
	}
	// Output:
	// Retry attempt 1 failed: network error
	// Success after 2 attempts
}

// ExampleDo_withConstantInterval demonstrates retry with constant interval
func ExampleDo_withConstantInterval() {
	attempts := 0

	f := retry.F(func() (bool, error) {
		attempts++
		if attempts < 2 {
			return false, errors.New("service unavailable")
		}
		return true, nil
	})

	start := time.Now()
	err := retry.Do(f,
		retry.Times(3),
		retry.Interval(retry.ConstantInterval(100*time.Millisecond)),
	)
	elapsed := time.Since(start)

	if err != nil {
		log.Printf("Operation failed: %v", err)
	} else {
		fmt.Printf("Success after %d attempts in %v\n", attempts, elapsed)
	}
	// Output: Success after 2 attempts in 100ms
}

// ExampleDo_httpRequest demonstrates HTTP request retry
func ExampleDo_httpRequest() {
	// Simulate HTTP request function
	makeRequest := func() (bool, error) {
		// In real usage, this would be an actual HTTP request
		// For this example, we'll simulate different scenarios
		return false, errors.New("server error 500")
	}

	// Retry HTTP requests with exponential backoff
	err := retry.Do(makeRequest,
		retry.Times(3),
		retry.Callback(func(n uint, err error) {
			fmt.Printf("HTTP request attempt %d failed: %v\n", n, err)
		}),
	)

	if err != nil {
		fmt.Printf("HTTP request failed after retries: %v\n", err)
	}
	// Output:
	// HTTP request attempt 1 failed: server error 500
	// HTTP request attempt 2 failed: server error 500
	// HTTP request attempt 3 failed: server error 500
	// HTTP request failed after retries: server error 500
}

// ExampleDo_conditionalRetry demonstrates conditional retry logic
func ExampleDo_conditionalRetry() {
	// Simulate an operation that returns different types of errors
	operation := func() (bool, error) {
		// Simulate different error scenarios
		return false, errors.New("temporary network issue")
	}

	// Only retry for temporary errors
	err := retry.Do(operation,
		retry.Times(2),
		retry.Callback(func(n uint, err error) {
			fmt.Printf("Attempt %d: %v\n", n, err)
		}),
	)

	if err != nil {
		fmt.Printf("Operation failed: %v\n", err)
	}
	// Output:
	// Attempt 1: temporary network issue
	// Attempt 2: temporary network issue
	// Operation failed: temporary network issue
}

// ExampleConstantInterval demonstrates creating a constant interval strategy
func ExampleConstantInterval() {
	// Create a constant interval of 200ms
	interval := retry.ConstantInterval(200 * time.Millisecond)

	attempts := 0
	f := retry.F(func() (bool, error) {
		attempts++
		if attempts < 2 {
			return false, errors.New("retry needed")
		}
		return true, nil
	})

	start := time.Now()
	err := retry.Do(f, retry.Times(3), retry.Interval(interval))
	elapsed := time.Since(start)

	if err != nil {
		fmt.Printf("Failed: %v\n", err)
	} else {
		fmt.Printf("Success in %v\n", elapsed)
	}
	// Output: Success in 200ms
}

// ExampleDefaultInterval demonstrates using the default exponential backoff
func ExampleDefaultInterval() {
	// Get the default interval strategy
	interval := retry.DefaultInterval()

	attempts := 0
	f := retry.F(func() (bool, error) {
		attempts++
		if attempts < 3 {
			return false, errors.New("service busy")
		}
		return true, nil
	})

	start := time.Now()
	err := retry.Do(f, retry.Times(3), retry.Interval(interval))
	elapsed := time.Since(start)

	if err != nil {
		fmt.Printf("Failed: %v\n", err)
	} else {
		fmt.Printf("Success after %d attempts in %v\n", attempts, elapsed)
	}
	// Output: Success after 3 attempts in 3s
}

// ExampleExponentialBackoffWithJitter demonstrates jitter functionality
func ExampleExponentialBackoffWithJitter() {
	attempts := 0

	f := retry.F(func() (bool, error) {
		attempts++
		if attempts < 2 {
			return false, errors.New("service unavailable")
		}
		return true, nil
	})

	// Use exponential backoff with 30% jitter
	start := time.Now()
	err := retry.Do(f,
		retry.Times(3),
		retry.Interval(retry.ExponentialBackoffWithJitter(100*time.Millisecond, 0.3)),
	)
	elapsed := time.Since(start)

	if err != nil {
		fmt.Printf("Failed: %v\n", err)
	} else {
		fmt.Printf("Success after %d attempts in %v\n", attempts, elapsed)
	}
	// Output: Success after 2 attempts in 100ms
}

// ExampleExponentialBackoffWithJitter_jitterComparison demonstrates the difference between jitter and no jitter
func ExampleExponentialBackoffWithJitter_jitterComparison() {
	// Without jitter - all clients retry at the same time
	fmt.Println("Without jitter:")
	attempts := 0
	f1 := retry.F(func() (bool, error) {
		attempts++
		if attempts < 2 {
			return false, errors.New("service busy")
		}
		return true, nil
	})

	start1 := time.Now()
	err1 := retry.Do(f1, retry.Times(2), retry.Interval(retry.ExponentialBackoffWithJitter(50*time.Millisecond, 0.0)))
	elapsed1 := time.Since(start1)
	if err1 != nil {
		fmt.Printf("Error: %v\n", err1)
	} else {
		fmt.Printf("Completed in %v\n", elapsed1)
	}

	// With jitter - clients retry at different times
	fmt.Println("With jitter:")
	attempts = 0
	f2 := retry.F(func() (bool, error) {
		attempts++
		if attempts < 2 {
			return false, errors.New("service busy")
		}
		return true, nil
	})

	start2 := time.Now()
	err2 := retry.Do(f2, retry.Times(2), retry.Interval(retry.ExponentialBackoffWithJitter(50*time.Millisecond, 0.5)))
	elapsed2 := time.Since(start2)
	if err2 != nil {
		fmt.Printf("Error: %v\n", err2)
	} else {
		fmt.Printf("Completed in %v\n", elapsed2)
	}
	// Output:
	// Without jitter:
	// Completed in 50ms
	// With jitter:
	// Completed in 75ms
}
