package errorx

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestThrow(t *testing.T) {
	t.Run("should panic on non-nil error", func(t *testing.T) {
		testErr := errors.New("test error")
		assert.Panics(t, func() {
			Throw(testErr)
		})
	})

	t.Run("should not panic on nil error", func(t *testing.T) {
		assert.NotPanics(t, func() {
			Throw(nil)
		})
	})

	t.Run("should panic with correct error value", func(t *testing.T) {
		testErr := errors.New("specific error message")
		assert.PanicsWithValue(t, testErr, func() {
			Throw(testErr)
		})
	})

	t.Run("should handle wrapped errors", func(t *testing.T) {
		originalErr := errors.New("original error")
		wrappedErr := fmt.Errorf("wrapped: %w", originalErr)

		assert.PanicsWithValue(t, wrappedErr, func() {
			Throw(wrappedErr)
		})
	})
}

func TestThrowf(t *testing.T) {
	t.Run("should panic with custom message on non-nil error", func(t *testing.T) {
		assert.PanicsWithValue(t, "custom error message", func() {
			Throwf(errors.New("test error"), "custom error message")
		})
	})

	t.Run("should panic with formatted message", func(t *testing.T) {
		assert.PanicsWithValue(t, "error code: 123", func() {
			Throwf(errors.New("test error"), "error code: %d", 123)
		})
	})

	t.Run("should not panic on nil error", func(t *testing.T) {
		assert.NotPanics(t, func() {
			Throwf(nil, "should not panic")
		})
	})

	t.Run("should handle multiple format arguments", func(t *testing.T) {
		expected := "User 123 failed operation: timeout after 30s"
		assert.PanicsWithValue(t, expected, func() {
			Throwf(errors.New("test"), "User %d failed operation: %s after %ds", 123, "timeout", 30)
		})
	})

	t.Run("should handle empty format string", func(t *testing.T) {
		assert.PanicsWithValue(t, "", func() {
			Throwf(errors.New("test"), "")
		})
	})

	t.Run("should handle format string with no placeholders", func(t *testing.T) {
		assert.PanicsWithValue(t, "simple message", func() {
			Throwf(errors.New("test"), "simple message")
		})
	})

	t.Run("should ignore format arguments when error is nil", func(t *testing.T) {
		// This should not panic even with invalid format arguments
		assert.NotPanics(t, func() {
			Throwf(nil, "invalid format: %s", "not a number")
		})
	})
}

func TestThrowv(t *testing.T) {
	t.Run("should panic on error", func(t *testing.T) {
		assert.Panics(t, func() {
			Throwv("result", errors.New("test error"))
		})
	})

	t.Run("should return value on nil error", func(t *testing.T) {
		result := Throwv("success", nil)
		assert.Equal(t, "success", result)
	})

	t.Run("should work with different types", func(t *testing.T) {
		result := Throwv(42, nil)
		assert.Equal(t, 42, result)
	})

	t.Run("should work with struct", func(t *testing.T) {
		type testStruct struct {
			Value int
			Name  string
		}
		expected := testStruct{Value: 123, Name: "test"}
		result := Throwv(expected, nil)
		assert.Equal(t, expected, result)
	})

	t.Run("should work with pointer types", func(t *testing.T) {
		value := 42
		result := Throwv(&value, nil)
		assert.Equal(t, &value, result)
		assert.Equal(t, 42, *result)
	})

	t.Run("should work with slice types", func(t *testing.T) {
		slice := []int{1, 2, 3}
		result := Throwv(slice, nil)
		assert.Equal(t, slice, result)
	})

	t.Run("should work with map types", func(t *testing.T) {
		m := map[string]int{"a": 1, "b": 2}
		result := Throwv(m, nil)
		assert.Equal(t, m, result)
	})

	t.Run("should work with interface types", func(t *testing.T) {
		var iface interface{} = "interface value"
		result := Throwv(iface, nil)
		assert.Equal(t, iface, result)
	})

	t.Run("should work with zero values", func(t *testing.T) {
		var zeroInt int
		result := Throwv(zeroInt, nil)
		assert.Equal(t, 0, result)

		var zeroString string
		result2 := Throwv(zeroString, nil)
		assert.Equal(t, "", result2)
	})

	t.Run("should panic with correct error value", func(t *testing.T) {
		testErr := errors.New("specific error")
		assert.PanicsWithValue(t, testErr, func() {
			Throwv("value", testErr)
		})
	})
}

// Edge cases and integration tests
func TestThrowIntegration(t *testing.T) {
	t.Run("Throw with PreferredError", func(t *testing.T) {
		preferredErr := Prefer400("bad request")
		assert.PanicsWithValue(t, preferredErr, func() {
			Throw(preferredErr)
		})
	})

	t.Run("Throwf with PreferredError", func(t *testing.T) {
		preferredErr := Prefer500("server error")
		assert.PanicsWithValue(t, "Operation failed: server error", func() {
			Throwf(preferredErr, "Operation failed: %s", preferredErr.Error())
		})
	})

	t.Run("Throwv with PreferredError", func(t *testing.T) {
		preferredErr := Prefer403("forbidden")
		assert.PanicsWithValue(t, preferredErr, func() {
			Throwv("data", preferredErr)
		})
	})
}

// Benchmark tests
func BenchmarkThrow(b *testing.B) {
	err := errors.New("benchmark error")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		func() {
			defer func() { recover() }()
			Throw(err)
		}()
	}
}

func BenchmarkThrowf(b *testing.B) {
	err := errors.New("benchmark error")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		func() {
			defer func() { recover() }()
			Throwf(err, "error: %s", err.Error())
		}()
	}
}

func BenchmarkThrowv(b *testing.B) {
	err := errors.New("benchmark error")
	value := "test value"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		func() {
			defer func() { recover() }()
			Throwv(value, err)
		}()
	}
}

// Example functions
func ExampleThrow() {
	// This will panic if err is not nil
	Throw(nil) // Does nothing
	// Throw(errors.New("error")) // This would panic
}

func ExampleThrowf() {
	// This will panic with custom message if err is not nil
	Throwf(nil, "operation failed") // Does nothing
	// Throwf(errors.New("error"), "operation failed") // This would panic with "operation failed"
}

func ExampleThrowv() {
	// Get result and handle error in one line
	result := Throwv("success", nil) // Returns "success"
	_ = result

	// This would panic if there's an error
	// result := Throwv("", errors.New("error"))
}

func ExampleThrow_riskyOperation() {
	// Simulate a risky operation that might fail
	riskyOperation := func() error {
		// Simulate some operation that might fail
		return nil // or errors.New("operation failed")
	}

	// Use Throw to panic on error
	err := riskyOperation()
	Throw(err) // Panics if err is not nil
	fmt.Println("Operation completed successfully")
}

func ExampleThrowf_withContext() {
	// Simulate operation with context
	operation := func(id int) error {
		// Simulate operation that might fail
		return nil // or errors.New("operation failed")
	}

	id := 123
	err := operation(id)
	Throwf(err, "operation failed for ID %d", id)
	fmt.Println("Operation completed successfully")
}

func ExampleThrowv_valueExtraction() {
	// Simulate a function that returns (value, error)
	getValue := func() (string, error) {
		// Simulate getting a value
		return "important data", nil // or errors.New("failed to get data")
	}

	// Extract value and panic on error in one line
	value := Throwv(getValue())
	fmt.Printf("Got value: %s\n", value)
}
