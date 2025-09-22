package errorx

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// CustomError is a test error type for testing error wrapping
type CustomError struct {
	Code    int
	Message string
}

func (e CustomError) Error() string {
	return e.Message
}

func TestWrap(t *testing.T) {
	t.Run("should return nil for nil error", func(t *testing.T) {
		assert.Nil(t, Wrap(nil))
	})

	t.Run("should wrap error without losing original", func(t *testing.T) {
		originalErr := errors.New("original error")
		wrappedErr := Wrap(originalErr)

		assert.NotNil(t, wrappedErr)
		assert.True(t, errors.Is(wrappedErr, originalErr))
		assert.Equal(t, originalErr.Error(), wrappedErr.Error())
	})

	t.Run("should preserve error chain", func(t *testing.T) {
		originalErr := errors.New("database connection failed")
		wrappedErr := Wrap(originalErr)

		assert.True(t, errors.Is(wrappedErr, originalErr))

		// Test error unwrapping - errors.As should be used with specific error types
		// For general error unwrapping, we can use errors.Unwrap
		unwrapped := errors.Unwrap(wrappedErr)
		assert.Equal(t, originalErr, unwrapped)
	})

	t.Run("should work with wrapped errors", func(t *testing.T) {
		originalErr := errors.New("base error")
		firstWrapped := Wrapf(originalErr, "first level")
		secondWrapped := Wrap(firstWrapped)

		assert.True(t, errors.Is(secondWrapped, originalErr))
		assert.True(t, errors.Is(secondWrapped, firstWrapped))
	})
}

func TestWrapf(t *testing.T) {
	t.Run("should return nil for nil error", func(t *testing.T) {
		assert.Nil(t, Wrapf(nil, "test"))
		assert.Nil(t, Wrapf(nil, "%s", "test"))
	})

	t.Run("should wrap error with simple message", func(t *testing.T) {
		originalErr := errors.New("bar")
		wrappedErr := Wrapf(originalErr, "foo")

		assert.NotNil(t, wrappedErr)
		assert.Equal(t, "foo: bar", wrappedErr.Error())
		assert.True(t, errors.Is(wrappedErr, originalErr))
	})

	t.Run("should wrap error with formatted message", func(t *testing.T) {
		originalErr := errors.New("quz")
		wrappedErr := Wrapf(originalErr, "foo %s", "bar")

		assert.NotNil(t, wrappedErr)
		assert.Equal(t, "foo bar: quz", wrappedErr.Error())
		assert.True(t, errors.Is(wrappedErr, originalErr))
	})

	t.Run("should handle multiple format arguments", func(t *testing.T) {
		originalErr := errors.New("timeout")
		wrappedErr := Wrapf(originalErr, "failed to connect to %s after %d seconds", "database", 30)

		assert.NotNil(t, wrappedErr)
		assert.Equal(t, "failed to connect to database after 30 seconds: timeout", wrappedErr.Error())
		assert.True(t, errors.Is(wrappedErr, originalErr))
	})

	t.Run("should handle empty format string", func(t *testing.T) {
		originalErr := errors.New("error")
		wrappedErr := Wrapf(originalErr, "")

		assert.NotNil(t, wrappedErr)
		assert.Equal(t, ": error", wrappedErr.Error())
		assert.True(t, errors.Is(wrappedErr, originalErr))
	})

	t.Run("should handle format string with no placeholders", func(t *testing.T) {
		originalErr := errors.New("error")
		wrappedErr := Wrapf(originalErr, "simple message")

		assert.NotNil(t, wrappedErr)
		assert.Equal(t, "simple message: error", wrappedErr.Error())
		assert.True(t, errors.Is(wrappedErr, originalErr))
	})

	t.Run("should work without format arguments", func(t *testing.T) {
		originalErr := errors.New("database error")
		wrappedErr := Wrapf(originalErr, "operation failed")

		assert.NotNil(t, wrappedErr)
		assert.Equal(t, "operation failed: database error", wrappedErr.Error())
		assert.True(t, errors.Is(wrappedErr, originalErr))
	})

	t.Run("should preserve error chain with multiple wraps", func(t *testing.T) {
		originalErr := errors.New("base error")
		firstWrapped := Wrapf(originalErr, "level1 error")
		secondWrapped := Wrapf(firstWrapped, "level2 error")

		// All should be able to unwrap to original
		assert.True(t, errors.Is(firstWrapped, originalErr))
		assert.True(t, errors.Is(secondWrapped, originalErr))
		assert.True(t, errors.Is(secondWrapped, firstWrapped))

		// But first wrapped should not be able to unwrap to second
		assert.False(t, errors.Is(firstWrapped, secondWrapped))
	})

	t.Run("should work with PreferredError", func(t *testing.T) {
		preferredErr := Prefer400("bad request")
		wrappedErr := Wrapf(preferredErr, "validation failed")

		assert.NotNil(t, wrappedErr)
		assert.True(t, errors.Is(wrappedErr, preferredErr))

		// Should preserve PreferredError properties
		var preferredErrType *PreferredError
		assert.True(t, errors.As(wrappedErr, &preferredErrType))
		assert.Equal(t, 400, preferredErrType.Code())
	})

	t.Run("should handle complex formatting", func(t *testing.T) {
		originalErr := errors.New("connection failed")
		wrappedErr := Wrapf(originalErr, "Operation %s failed for user %s (ID: %d, retries: %d)",
			"login", "john", 123, 3)

		expected := "Operation login failed for user john (ID: 123, retries: 3): connection failed"
		assert.Equal(t, expected, wrappedErr.Error())
		assert.True(t, errors.Is(wrappedErr, originalErr))
	})
}

// Edge cases and integration tests
func TestWrapIntegration(t *testing.T) {
	t.Run("error unwrapping chain", func(t *testing.T) {
		baseErr := errors.New("base error")
		level1 := Wrapf(baseErr, "level 1")
		level2 := Wrapf(level1, "level 2")
		level3 := Wrap(level2)

		// Test unwrapping chain
		assert.True(t, errors.Is(level3, baseErr))
		assert.True(t, errors.Is(level3, level1))
		assert.True(t, errors.Is(level3, level2))

		// Test error messages
		assert.Contains(t, level3.Error(), "level 2")
		assert.Contains(t, level3.Error(), "level 1")
		assert.Contains(t, level3.Error(), "base error")
	})

	t.Run("with custom error types", func(t *testing.T) {
		customErr := CustomError{Code: 123, Message: "custom error"}
		wrappedErr := Wrapf(customErr, "operation failed")

		assert.True(t, errors.Is(wrappedErr, customErr))

		// Test error unwrapping
		var unwrapped CustomError
		assert.True(t, errors.As(wrappedErr, &unwrapped))
		assert.Equal(t, 123, unwrapped.Code)
		assert.Equal(t, "custom error", unwrapped.Message)
	})
}

// Benchmark tests
func BenchmarkWrap(b *testing.B) {
	err := errors.New("benchmark error")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Wrap(err)
	}
}

func BenchmarkWrapf(b *testing.B) {
	err := errors.New("benchmark error")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Wrapf(err, "operation failed: %d", i)
	}
}

func BenchmarkWrapfSimple(b *testing.B) {
	err := errors.New("benchmark error")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Wrapf(err, "simple message")
	}
}

func BenchmarkWrapNil(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Wrap(nil)
	}
}

// Example functions
func ExampleWrap() {
	originalErr := errors.New("database connection failed")
	wrappedErr := Wrap(originalErr)

	fmt.Printf("Original: %s\n", originalErr.Error())
	fmt.Printf("Wrapped: %s\n", wrappedErr.Error())
	fmt.Printf("Is same error: %t\n", errors.Is(wrappedErr, originalErr))

	// Output:
	// Original: database connection failed
	// Wrapped: database connection failed
	// Is same error: true
}

func ExampleWrapf() {
	originalErr := errors.New("connection timeout")
	wrappedErr := Wrapf(originalErr, "failed to connect to %s", "database")

	fmt.Printf("Wrapped error: %s\n", wrappedErr.Error())
	fmt.Printf("Is same error: %t\n", errors.Is(wrappedErr, originalErr))

	// Output:
	// Wrapped error: failed to connect to database: connection timeout
	// Is same error: true
}

func ExampleWrap_errorChain() {
	// Create a chain of wrapped errors
	baseErr := errors.New("base error")
	level1 := Wrapf(baseErr, "level 1 context")
	level2 := Wrapf(level1, "level 2 context")
	level3 := Wrap(level2)

	// All levels can unwrap to the base error
	fmt.Printf("Level 3 contains base: %t\n", errors.Is(level3, baseErr))
	fmt.Printf("Level 3 contains level 1: %t\n", errors.Is(level3, level1))
	fmt.Printf("Level 3 contains level 2: %t\n", errors.Is(level3, level2))

	// Output:
	// Level 3 contains base: true
	// Level 3 contains level 1: true
	// Level 3 contains level 2: true
}

func ExampleWrapf_withPreferredError() {
	// Wrap a PreferredError
	preferredErr := Prefer400("bad request")
	wrappedErr := Wrapf(preferredErr, "validation failed for user %s", "john")

	fmt.Printf("Wrapped error: %s\n", wrappedErr.Error())
	fmt.Printf("Is PreferredError: %t\n", errors.Is(wrappedErr, preferredErr))

	// Unwrap to get the PreferredError
	var preferredErrType *PreferredError
	if errors.As(wrappedErr, &preferredErrType) {
		fmt.Printf("HTTP Status Code: %d\n", preferredErrType.Code())
	}

	// Output:
	// Wrapped error: validation failed for user john: bad request
	// Is PreferredError: true
	// HTTP Status Code: 400
}

func ExampleWrapf_complexFormatting() {
	originalErr := errors.New("permission denied")
	wrappedErr := Wrapf(originalErr, "User %s (ID: %d) %s access to %s",
		"alice", 123, "denied", "admin panel")

	fmt.Printf("Error: %s\n", wrappedErr.Error())

	// Output:
	// Error: User alice (ID: 123) denied access to admin panel: permission denied
}
