package errorx_test

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/go4x/goal/errorx"
)

// ExamplePreferredError demonstrates how to create and use PreferredError
func ExamplePreferredError() {
	// Create a preferred error with HTTP status code
	err := errorx.Prefer500("database connection failed")

	// Check if it's a preferred error
	if errorx.IsPreferred(err) {
		preferredErr := err.(*errorx.PreferredError)
		fmt.Printf("HTTP Status: %d\n", preferredErr.Code())
		fmt.Printf("Error: %s\n", preferredErr.Error())
	}

	// Output:
	// HTTP Status: 500
	// Error: database connection failed
}

// ExampleChainExecutor demonstrates how to use ErrorHandler for centralized error handling
func ExampleChainExecutor() {
	// Create a new error handler
	executor := errorx.NewChainExecutor()

	// Chain operations with error handling
	executor.Do(func() error {
		// First operation that might fail
		return nil
	}).Do(func() error {
		// Second operation that might fail
		return errorx.Prefer400("invalid input")
	}).Do(func() error {
		// This won't execute because previous operation failed
		fmt.Println("This won't be printed")
		return nil
	})

	// Check if any error occurred
	if executor.HasErr() {
		fmt.Printf("Error occurred: %v\n", executor.Err())
	}

	// Output:
	// Error occurred: invalid input
}

// ExampleChainExecutor_DoWithContext demonstrates how to use ErrorHandler with context tracking
func ExampleChainExecutor_DoWithContext() {
	executor := errorx.NewChainExecutor()

	// Use DoWithContext to track which step failed
	executor.DoWithContext("validate", func() error {
		// Validation logic
		return nil
	}).DoWithContext("process", func() error {
		// Processing logic that fails
		return errors.New("processing failed")
	}).DoWithContext("save", func() error {
		// This won't execute because previous step failed
		return nil
	})

	if executor.HasErr() {
		fmt.Printf("Failed at step '%s': %v\n", executor.FailedStep(), executor.Err())
	}

	// Output:
	// Failed at step 'process': processing failed
}

// ExampleChainExecutor_DoWithResult demonstrates how to use ErrorHandler with result storage
func ExampleChainExecutor_DoWithResult() {
	executor := errorx.NewChainExecutor()

	// Store results from operations
	executor.DoWithResult("calculate", func() (interface{}, error) {
		result := 42 * 2
		return result, nil
	}).DoWithContext("validate", func() error {
		// Use the stored result
		result, _ := executor.GetResult("calculate")
		if result.(int) >= 100 {
			return errors.New("result too small")
		}
		return nil
	})

	if executor.HasErr() {
		fmt.Printf("Validation failed: %v\n", executor.Err())
	} else {
		result, _ := executor.GetResult("calculate")
		fmt.Printf("Calculation result: %v\n", result)
	}

	// Output:
	// Calculation result: 84
}

// ExampleRecover demonstrates how to use Recover for panic recovery
func ExampleRecover() {
	// Simulate a function that might panic
	func() {
		defer errorx.Recover(func(r any) {
			fmt.Println(r)
		}, func() {
			fmt.Println("Cleanup: closing resources")
		})

		// This would normally panic, but we'll simulate success
		fmt.Println("Operation completed successfully")
	}()

	// Output:
	// Operation completed successfully
}

// ExampleRecoverCtx demonstrates how to use RecoverCtx with context
func ExampleRecoverCtx() {
	// Create context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	// Simulate a function that might panic
	func() {
		defer errorx.RecoverCtx(ctx, func(r any) {
			fmt.Println(r)
		}, func() {
			fmt.Println("Cleanup: releasing resources")
		})

		fmt.Println("Operation with context completed")
	}()

	// Output:
	// Operation with context completed
}

// ExampleThrow demonstrates how to use Throw for panic-based error handling
func ExampleThrow() {
	// This will not panic because error is nil
	errorx.Throw(nil)
	fmt.Println("No panic occurred")

	// This would panic if error is not nil
	// errorx.Throw(errors.New("something went wrong"))

	// Output:
	// No panic occurred
}

// ExampleThrowv demonstrates how to use Throwv for value extraction with error handling
func ExampleThrowv() {
	// Simulate a function that returns a value and an error
	getValue := func() (string, error) {
		return "success", nil
	}

	// Extract value, panic if error
	value := errorx.Throwv(getValue())
	fmt.Printf("Value: %s\n", value)

	// Output:
	// Value: success
}

// ExampleWrap demonstrates how to wrap errors with additional context
func ExampleWrap() {
	originalErr := fmt.Errorf("database connection failed")

	// Wrap with additional context
	wrappedErr := errorx.Wrap(originalErr)

	fmt.Printf("Wrapped error: %v\n", wrappedErr)

	// Output:
	// Wrapped error: database connection failed
}

// ExampleWrapf demonstrates how to wrap errors with formatted context
func ExampleWrapf() {
	originalErr := fmt.Errorf("network timeout")

	// Wrap with formatted context
	wrappedErr := errorx.Wrapf(originalErr, "request to %s failed", "api.example.com")

	fmt.Printf("Wrapped error: %v\n", wrappedErr)

	// Output:
	// Wrapped error: request to api.example.com failed: network timeout
}

// ExampleHTTPErrorFactories demonstrates how to use HTTP error factory functions
func ExamplePrefer500() {
	// Create different types of HTTP errors
	serverErr := errorx.Prefer500("service unavailable")
	badReqErr := errorx.Prefer400("invalid JSON format")
	unauthErr := errorx.Prefer401("missing authentication token")
	tooManyErr := errorx.Prefer429("rate limit exceeded")
	forbiddenErr := errorx.Prefer403("access denied")

	// Check their HTTP status codes
	fmt.Printf("Server busy: %d\n", serverErr.(*errorx.PreferredError).Code())
	fmt.Printf("Bad request: %d\n", badReqErr.(*errorx.PreferredError).Code())
	fmt.Printf("Unauthorized: %d\n", unauthErr.(*errorx.PreferredError).Code())
	fmt.Printf("Too many requests: %d\n", tooManyErr.(*errorx.PreferredError).Code())
	fmt.Printf("Forbidden: %d\n", forbiddenErr.(*errorx.PreferredError).Code())

	// Output:
	// Server busy: 500
	// Bad request: 400
	// Unauthorized: 401
	// Too many requests: 429
	// Forbidden: 403
}

// ExampleChainExecutorPreferred demonstrates how to use ErrorHandler with preferred errors
func ExampleChainExecutor_IsPreferredErr() {
	executor := errorx.NewChainExecutor()

	// Chain operations
	executor.Do(func() error {
		// First operation succeeds
		return nil
	}).Do(func() error {
		// Second operation fails with preferred error
		return errorx.Prefer500("temporary service issue")
	})

	// Check if error is preferred
	if executor.IsPreferredErr() {
		fmt.Println("Preferred error occurred")
	}

	// Get preferred error or fallback
	fallbackErr := fmt.Errorf("fallback error")
	finalErr := executor.PreferredOr(fallbackErr)

	fmt.Printf("Final error: %v\n", finalErr)

	// Output:
	// Preferred error occurred
	// Final error: temporary service issue
}
