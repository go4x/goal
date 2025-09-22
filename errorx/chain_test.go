package errorx

import (
	"errors"
	"fmt"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChainExecutor_BasicOperations(t *testing.T) {
	t.Run("no_errors", func(t *testing.T) {
		executor := NewChainExecutor()

		executor.Do(func() error {
			return nil
		}).Do(func() error {
			return nil
		})

		assert.False(t, executor.HasErr())
		assert.Nil(t, executor.Err())
		assert.Empty(t, executor.FailedStep())
	})

	t.Run("single_error", func(t *testing.T) {
		executor := NewChainExecutor()
		expectedErr := errors.New("test error")

		executor.Do(func() error {
			return expectedErr
		}).Do(func() error {
			t.Error("This should not be executed")
			return nil
		})

		assert.True(t, executor.HasErr())
		assert.Equal(t, expectedErr, executor.Err())
	})

	t.Run("error_in_middle", func(t *testing.T) {
		executor := NewChainExecutor()
		step1Executed := false
		step2Executed := false
		step3Executed := false

		executor.Do(func() error {
			step1Executed = true
			return nil
		}).Do(func() error {
			step2Executed = true
			return errors.New("step2 error")
		}).Do(func() error {
			step3Executed = true
			return nil
		})

		assert.True(t, step1Executed)
		assert.True(t, step2Executed)
		assert.False(t, step3Executed)
		assert.True(t, executor.HasErr())
		assert.Equal(t, "step2 error", executor.Err().Error())
	})
}

func TestChainExecutor_DoWithContext(t *testing.T) {
	t.Run("success_with_context", func(t *testing.T) {
		executor := NewChainExecutor()

		executor.DoWithContext("validate", func() error {
			return nil
		}).DoWithContext("process", func() error {
			return nil
		})

		assert.False(t, executor.HasErr())
		assert.Empty(t, executor.FailedStep())
	})

	t.Run("error_with_context", func(t *testing.T) {
		executor := NewChainExecutor()

		executor.DoWithContext("validate", func() error {
			return nil
		}).DoWithContext("process", func() error {
			return errors.New("processing error")
		}).DoWithContext("save", func() error {
			t.Error("This should not be executed")
			return nil
		})

		assert.True(t, executor.HasErr())
		assert.Equal(t, "process", executor.FailedStep())
		assert.Equal(t, "processing error", executor.Err().Error())
	})
}

func TestChainExecutor_DoWithResult(t *testing.T) {
	t.Run("success_with_result", func(t *testing.T) {
		executor := NewChainExecutor()

		executor.DoWithResult("data", func() (interface{}, error) {
			return "test data", nil
		}).DoWithResult("number", func() (interface{}, error) {
			return 42, nil
		})

		assert.False(t, executor.HasErr())

		result1, exists1 := executor.GetResult("data")
		assert.True(t, exists1)
		assert.Equal(t, "test data", result1)

		result2, exists2 := executor.GetResult("number")
		assert.True(t, exists2)
		assert.Equal(t, 42, result2)
	})

	t.Run("error_with_result", func(t *testing.T) {
		executor := NewChainExecutor()

		executor.DoWithResult("data", func() (interface{}, error) {
			return "test data", nil
		}).DoWithResult("error_step", func() (interface{}, error) {
			return nil, errors.New("result error")
		}).DoWithResult("final", func() (interface{}, error) {
			t.Error("This should not be executed")
			return nil, nil
		})

		assert.True(t, executor.HasErr())
		assert.Equal(t, "error_step", executor.FailedStep())

		// First result should still be available
		result1, exists1 := executor.GetResult("data")
		assert.True(t, exists1)
		assert.Equal(t, "test data", result1)

		// Second result should not be stored
		result2, exists2 := executor.GetResult("error_step")
		assert.False(t, exists2)
		assert.Nil(t, result2)
	})
}

func ExampleNewChainExecutor() {
	executor := NewChainExecutor()
	fmt.Println(executor.HasErr())
	fmt.Println(executor.Err())
	// Output:
	// false
	// <nil>
}

func ExampleChainExecutor_Do() {
	// Repeated error handling
	i, err := f1()
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	j, err := f2()
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	fmt.Printf("i= %d, j = %d\n", i, j)

	// Using ChainExecutor
	var x, y int
	executor := NewChainExecutor()
	executor.Do(func() error {
		x, err = f1()
		return err
	})
	executor.Do(func() error {
		y, err = f2()
		return err
	})
	if executor.HasErr() {
		fmt.Printf("error: %v\n", err)
	}
	fmt.Printf("x = %d, y = %d\n", x, y)
}

func TestChainExecutor_DoWithTypedResult(t *testing.T) {
	t.Run("success_with_typed_result", func(t *testing.T) {
		executor := NewChainExecutor()

		executor = DoWithTypedResult(executor, "string_result", func() (string, error) {
			return "typed string", nil
		})

		executor = DoWithTypedResult(executor, "int_result", func() (int, error) {
			return 123, nil
		})

		assert.False(t, executor.HasErr())

		strResult, exists := GetTypedResult[string](executor, "string_result")
		assert.True(t, exists)
		assert.Equal(t, "typed string", strResult)

		intResult, exists := GetTypedResult[int](executor, "int_result")
		assert.True(t, exists)
		assert.Equal(t, 123, intResult)
	})

	t.Run("error_with_typed_result", func(t *testing.T) {
		executor := NewChainExecutor()

		executor = DoWithTypedResult(executor, "string_result", func() (string, error) {
			return "typed string", nil
		})

		executor = DoWithTypedResult(executor, "error_step", func() (int, error) {
			return 0, errors.New("typed error")
		})

		executor = DoWithTypedResult(executor, "final", func() (string, error) {
			t.Error("This should not be executed")
			return "", nil
		})

		assert.True(t, executor.HasErr())
		assert.Equal(t, "error_step", executor.FailedStep())

		// First result should still be available
		strResult, exists := GetTypedResult[string](executor, "string_result")
		assert.True(t, exists)
		assert.Equal(t, "typed string", strResult)

		// Second result should not be stored
		intResult, exists := GetTypedResult[int](executor, "error_step")
		assert.False(t, exists)
		assert.Equal(t, 0, intResult)
	})
}

func TestChainExecutor_Reset(t *testing.T) {
	executor := NewChainExecutor()

	// Execute with error and results
	executor.DoWithContext("error_step", func() error {
		return errors.New("test error")
	}).DoWithResult("test_key", func() (interface{}, error) {
		return "test_value", nil
	})

	assert.True(t, executor.HasErr())
	assert.Equal(t, "error_step", executor.FailedStep())

	// Reset
	resetExecutor := executor.Reset()
	assert.Same(t, executor, resetExecutor) // Should return same instance

	assert.False(t, executor.HasErr())
	assert.Nil(t, executor.Err())
	assert.Empty(t, executor.FailedStep())

	// Results should be cleared
	_, exists := executor.GetResult("test_key")
	assert.False(t, exists)

	// Should be able to use again
	executor.Do(func() error {
		return nil
	})
	assert.False(t, executor.HasErr())
}

func TestChainExecutor_IsPreferredErr(t *testing.T) {
	t.Run("no_error", func(t *testing.T) {
		executor := NewChainExecutor()
		assert.False(t, executor.IsPreferredErr())
	})

	t.Run("regular_error", func(t *testing.T) {
		executor := NewChainExecutor()
		executor.Do(func() error {
			return errors.New("regular error")
		})
		assert.False(t, executor.IsPreferredErr())
	})

	t.Run("preferred_error", func(t *testing.T) {
		executor := NewChainExecutor()
		executor.Do(func() error {
			return Prefer400("preferred error")
		})
		assert.True(t, executor.IsPreferredErr())
	})
}

func TestChainExecutor_String(t *testing.T) {
	t.Run("no_error", func(t *testing.T) {
		executor := NewChainExecutor()
		str := executor.String()
		assert.Equal(t, "ChainExecutor{no error}", str)
	})

	t.Run("error_without_step", func(t *testing.T) {
		executor := NewChainExecutor()
		expectedErr := errors.New("test error")

		executor.Do(func() error {
			return expectedErr
		})

		str := executor.String()
		assert.Equal(t, "ChainExecutor{error: test error}", str)
	})

	t.Run("error_with_step", func(t *testing.T) {
		executor := NewChainExecutor()
		expectedErr := errors.New("test error")

		executor.DoWithContext("test_step", func() error {
			return expectedErr
		})

		str := executor.String()
		assert.Equal(t, "ChainExecutor{error: test error, failed at: test_step}", str)
	})
}

func TestChainExecutor_EdgeCases(t *testing.T) {
	t.Run("nil_function", func(t *testing.T) {
		executor := NewChainExecutor()

		// This should not panic and should not set an error
		executor.Do(nil)
		assert.False(t, executor.HasErr())
		assert.Nil(t, executor.Err())
	})

	t.Run("empty_step_name", func(t *testing.T) {
		executor := NewChainExecutor()

		executor.DoWithContext("", func() error {
			return errors.New("empty step error")
		})

		assert.True(t, executor.HasErr())
		assert.Equal(t, "", executor.FailedStep())
	})

	t.Run("duplicate_result_keys", func(t *testing.T) {
		executor := NewChainExecutor()

		executor.DoWithResult("key", func() (interface{}, error) {
			return "first", nil
		}).DoWithResult("key", func() (interface{}, error) {
			return "second", nil
		})

		result, exists := executor.GetResult("key")
		assert.True(t, exists)
		assert.Equal(t, "second", result) // Should be overwritten
	})
}

// Benchmark tests
func BenchmarkChainExecutor_Do(b *testing.B) {
	executor := NewChainExecutor()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		executor.Do(func() error {
			return nil
		})
	}
}

func BenchmarkChainExecutor_DoWithContext(b *testing.B) {
	executor := NewChainExecutor()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		executor.DoWithContext("step", func() error {
			return nil
		})
	}
}

func BenchmarkChainExecutor_DoWithResult(b *testing.B) {
	executor := NewChainExecutor()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		executor.DoWithResult("key", func() (interface{}, error) {
			return "value", nil
		})
	}
}

func BenchmarkChainExecutor_GetResult(b *testing.B) {
	executor := NewChainExecutor()
	executor.DoWithResult("key", func() (interface{}, error) {
		return "value", nil
	})

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = executor.GetResult("key")
	}
}

func ExampleChainExecutor_DoWithContext() {
	executor := NewChainExecutor()

	executor.DoWithContext("validate", func() error {
		fmt.Println("Validating...")
		return nil
	}).DoWithContext("process", func() error {
		fmt.Println("Processing...")
		return nil
	})

	if executor.HasErr() {
		fmt.Printf("Failed at: %s, Error: %v\n", executor.FailedStep(), executor.Err())
	} else {
		fmt.Println("All steps completed successfully")
	}
	// Output:
	// Validating...
	// Processing...
	// All steps completed successfully
}

func ExampleChainExecutor_DoWithResult() {
	executor := NewChainExecutor()

	executor.DoWithResult("user_id", func() (interface{}, error) {
		return 123, nil
	}).DoWithResult("user_name", func() (interface{}, error) {
		return "john_doe", nil
	})

	if executor.HasErr() {
		fmt.Printf("Error: %v\n", executor.Err())
		return
	}

	userID, _ := executor.GetResult("user_id")
	userName, _ := executor.GetResult("user_name")
	fmt.Printf("User: %s (ID: %d)\n", userName, userID)
	// Output:
	// User: john_doe (ID: 123)
}

func ExampleChainExecutor_Reset() {
	executor := NewChainExecutor()

	// First execution with error
	executor.Do(func() error {
		return errors.New("first error")
	})
	fmt.Printf("Has error: %v\n", executor.HasErr())

	// Reset and try again
	executor.Reset().Do(func() error {
		return nil
	})
	fmt.Printf("Has error after reset: %v\n", executor.HasErr())
	// Output:
	// Has error: true
	// Has error after reset: false
}

func f1() (int, error) {
	n := rand.Intn(100)
	if n%3 == 0 {
		return -1, fmt.Errorf("make an error1")
	}
	return n, nil
}

func f2() (int, error) {
	n := rand.Intn(200)
	if n%3 == 0 {
		return -2, fmt.Errorf("make an error2")
	}
	return n, nil
}

func TestErrorHandlerEnhanced(t *testing.T) {
	t.Run("basic functionality", func(t *testing.T) {
		executor := NewChainExecutor()

		// Test successful execution
		executor.Do(func() error { return nil })
		assert.False(t, executor.HasErr())
		assert.Empty(t, executor.FailedStep())

		// Test error handling
		executor.Do(func() error { return errors.New("test error") })
		assert.True(t, executor.HasErr())
		assert.Equal(t, "test error", executor.Err().Error())

		// Test subsequent operations are skipped
		called := false
		executor.Do(func() error {
			called = true
			return nil
		})
		assert.False(t, called)
	})

	t.Run("DoWithContext tracks failed step", func(t *testing.T) {
		executor := NewChainExecutor()

		executor.DoWithContext("validation", func() error {
			return errors.New("validation failed")
		})

		assert.True(t, executor.HasErr())
		assert.Equal(t, "validation failed", executor.Err().Error())
		assert.Equal(t, "validation", executor.FailedStep())
	})

	t.Run("DoWithResult stores results", func(t *testing.T) {
		executor := NewChainExecutor()

		executor.DoWithResult("process", func() (interface{}, error) {
			return "processed_data", nil
		})

		assert.False(t, executor.HasErr())
		result, exists := executor.GetResult("process")
		assert.True(t, exists)
		assert.Equal(t, "processed_data", result)
	})

	t.Run("DoWithTypedResult with type safety", func(t *testing.T) {
		executor := NewChainExecutor()

		DoWithTypedResult(executor, "calculate", func() (int, error) {
			return 42, nil
		})

		assert.False(t, executor.HasErr())
		result, exists := GetTypedResult[int](executor, "calculate")
		assert.True(t, exists)
		assert.Equal(t, 42, result)

		// Test type mismatch
		_, exists = GetTypedResult[string](executor, "calculate")
		assert.False(t, exists)
	})

	t.Run("chained operations with results", func(t *testing.T) {
		executor := NewChainExecutor()

		executor.
			DoWithContext("validate", func() error { return nil }).
			DoWithResult("process", func() (interface{}, error) { return "data", nil }).
			DoWithContext("save", func() error { return nil })

		assert.False(t, executor.HasErr())
		result, exists := executor.GetResult("process")
		assert.True(t, exists)
		assert.Equal(t, "data", result)
	})

	t.Run("fail-fast behavior with context", func(t *testing.T) {
		executor := NewChainExecutor()

		executor.
			DoWithContext("step1", func() error { return nil }).
			DoWithContext("step2", func() error { return errors.New("step2 failed") }).
			DoWithContext("step3", func() error { return nil })

		assert.True(t, executor.HasErr())
		assert.Equal(t, "step2", executor.FailedStep())

		// step3 should not have been executed
		_, exists := executor.GetResult("step3")
		assert.False(t, exists)
	})
}

func TestErrorHandlerResults(t *testing.T) {
	t.Run("GetResult with non-existent key", func(t *testing.T) {
		executor := NewChainExecutor()

		result, exists := executor.GetResult("nonexistent")
		assert.False(t, exists)
		assert.Nil(t, result)
	})

	t.Run("GetTypedResult with non-existent key", func(t *testing.T) {
		executor := NewChainExecutor()

		result, exists := GetTypedResult[string](executor, "nonexistent")
		assert.False(t, exists)
		assert.Equal(t, "", result)
	})

	t.Run("multiple results with different types", func(t *testing.T) {
		executor := NewChainExecutor()

		DoWithTypedResult(executor, "string", func() (string, error) { return "hello", nil })
		DoWithTypedResult(executor, "int", func() (int, error) { return 123, nil })
		DoWithTypedResult(executor, "bool", func() (bool, error) { return true, nil })

		str, exists := GetTypedResult[string](executor, "string")
		assert.True(t, exists)
		assert.Equal(t, "hello", str)

		num, exists := GetTypedResult[int](executor, "int")
		assert.True(t, exists)
		assert.Equal(t, 123, num)

		flag, exists := GetTypedResult[bool](executor, "bool")
		assert.True(t, exists)
		assert.True(t, flag)
	})
}

func TestErrorHandlerDoneMethods(t *testing.T) {
	t.Run("Done with error", func(t *testing.T) {
		executor := NewChainExecutor()
		executor.DoWithContext("test", func() error { return errors.New("test error") })

		var capturedErr error
		executor.Done(func(err error) {
			capturedErr = err
		})

		assert.Equal(t, "test error", capturedErr.Error())
	})

	t.Run("DoneWithContext provides step information", func(t *testing.T) {
		executor := NewChainExecutor()
		executor.DoWithContext("validation", func() error { return errors.New("validation failed") })

		var capturedErr error
		var capturedStep string
		executor.DoneWithContext(func(err error, step string) {
			capturedErr = err
			capturedStep = step
		})

		assert.Equal(t, "validation failed", capturedErr.Error())
		assert.Equal(t, "validation", capturedStep)
	})
}

func TestErrorHandlerReset(t *testing.T) {
	t.Run("Reset clears all state", func(t *testing.T) {
		executor := NewChainExecutor()

		// Set up some state - first add result, then error
		executor.DoWithResult("data", func() (interface{}, error) { return "data", nil })
		executor.DoWithContext("test", func() error { return errors.New("test error") })

		assert.True(t, executor.HasErr())
		assert.Equal(t, "test", executor.FailedStep())
		result, exists := executor.GetResult("data")
		assert.True(t, exists)
		assert.Equal(t, "data", result.(string))

		// Reset
		executor.Reset()

		assert.False(t, executor.HasErr())
		assert.Empty(t, executor.FailedStep())
		_, existsAfterReset := executor.GetResult("data")
		assert.False(t, existsAfterReset)
	})
}

func TestErrorHandlerErrorInfo(t *testing.T) {
	t.Run("ErrorInfo with error", func(t *testing.T) {
		executor := NewChainExecutor()
		executor.DoWithContext("test", func() error { return errors.New("test error") })

		info := executor.GetErrorInfo()
		assert.True(t, info.HasError)
		assert.Equal(t, "test error", info.Err.Error())
		assert.Equal(t, "test", info.FailedStep)
	})

	t.Run("ErrorInfo without error", func(t *testing.T) {
		executor := NewChainExecutor()

		info := executor.GetErrorInfo()
		assert.False(t, info.HasError)
		assert.Nil(t, info.Err)
		assert.Empty(t, info.FailedStep)
	})
}

func TestErrorHandlerString(t *testing.T) {
	t.Run("String representation with error and step", func(t *testing.T) {
		executor := NewChainExecutor()
		executor.DoWithContext("test", func() error { return errors.New("test error") })

		str := executor.String()
		assert.Contains(t, str, "test error")
		assert.Contains(t, str, "failed at: test")
	})

	t.Run("String representation with error only", func(t *testing.T) {
		executor := NewChainExecutor()
		executor.Do(func() error { return errors.New("test error") })

		str := executor.String()
		assert.Contains(t, str, "test error")
		assert.NotContains(t, str, "failed at:")
	})

	t.Run("String representation without error", func(t *testing.T) {
		executor := NewChainExecutor()

		str := executor.String()
		assert.Equal(t, "ChainExecutor{no error}", str)
	})
}

func TestErrorHandlerPreferredError(t *testing.T) {
	t.Run("IsPreferredErr with PreferredError", func(t *testing.T) {
		executor := NewChainExecutor()
		executor.Do(func() error { return Prefer400("bad request") })

		assert.True(t, executor.IsPreferredErr())
	})

	t.Run("PreferredOr returns preferred error", func(t *testing.T) {
		executor := NewChainExecutor()
		executor.Do(func() error { return Prefer500("server error") })

		fallbackErr := errors.New("fallback")
		result := executor.PreferredOr(fallbackErr)

		assert.Equal(t, "server error", result.Error())
	})

	t.Run("PreferredOr returns fallback for non-preferred error", func(t *testing.T) {
		executor := NewChainExecutor()
		executor.Do(func() error { return errors.New("regular error") })

		fallbackErr := errors.New("fallback")
		result := executor.PreferredOr(fallbackErr)

		assert.Equal(t, "fallback", result.Error())
	})
}

func BenchmarkErrorHandler(b *testing.B) {
	b.Run("Do", func(b *testing.B) {
		executor := NewChainExecutor()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			executor.Do(func() error { return nil })
		}
	})

	b.Run("DoWithContext", func(b *testing.B) {
		executor := NewChainExecutor()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			executor.DoWithContext("test", func() error { return nil })
		}
	})

	b.Run("DoWithResult", func(b *testing.B) {
		executor := NewChainExecutor()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			executor.DoWithResult("test", func() (interface{}, error) { return "result", nil })
		}
	})

	b.Run("DoWithTypedResult", func(b *testing.B) {
		executor := NewChainExecutor()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			DoWithTypedResult(executor, "test", func() (string, error) { return "result", nil })
		}
	})

	b.Run("GetTypedResult", func(b *testing.B) {
		executor := NewChainExecutor()
		DoWithTypedResult(executor, "test", func() (int, error) { return 42, nil })
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_, _ = GetTypedResult[int](executor, "test")
		}
	})
}

// Example functions for documentation
func ExampleChainExecutor_basic() {
	executor := NewChainExecutor()

	executor.Do(func() error {
		// First operation
		fmt.Println("Executing step 1")
		return nil
	}).Do(func() error {
		// Second operation
		fmt.Println("Executing step 2")
		return nil
	}).Do(func() error {
		// Third operation
		fmt.Println("Executing step 3")
		return nil
	})

	if executor.HasErr() {
		fmt.Printf("Failed at step: %s, Error: %v\n", executor.FailedStep(), executor.Err())
	} else {
		fmt.Println("All operations completed successfully")
	}

	// Output:
	// Executing step 1
	// Executing step 2
	// Executing step 3
	// All operations completed successfully
}

func ExampleChainExecutor_withContext() {
	executor := NewChainExecutor()

	executor.DoWithContext("validate", func() error {
		fmt.Println("Validating input")
		return nil
	}).DoWithContext("process", func() error {
		fmt.Println("Processing data")
		return errors.New("processing failed")
	}).DoWithContext("save", func() error {
		fmt.Println("Saving results") // This won't execute
		return nil
	})

	if executor.HasErr() {
		fmt.Printf("Failed at step: %s, Error: %v\n", executor.FailedStep(), executor.Err())
	}

	// Output:
	// Validating input
	// Processing data
	// Failed at step: process, Error: processing failed
}

func ExampleChainExecutor_withResults() {
	executor := NewChainExecutor()

	executor.DoWithResult("calculate", func() (interface{}, error) {
		result := 42 * 2
		return result, nil
	}).DoWithContext("validate", func() error {
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

func ExampleChainExecutor_typedResults() {
	executor := NewChainExecutor()

	DoWithTypedResult(executor, "username", func() (string, error) {
		return "john_doe", nil
	})
	DoWithTypedResult(executor, "age", func() (int, error) {
		return 25, nil
	})
	executor.DoWithContext("create_user", func() error {
		username, _ := GetTypedResult[string](executor, "username")
		age, _ := GetTypedResult[int](executor, "age")
		fmt.Printf("Creating user: %s, age: %d\n", username, age)
		return nil
	})

	if !executor.HasErr() {
		fmt.Println("User created successfully")
	}

	// Output:
	// Creating user: john_doe, age: 25
	// User created successfully
}

func ExampleChainExecutor_doneWithContext() {
	executor := NewChainExecutor()

	executor.DoWithContext("database_operation", func() error {
		return errors.New("connection timeout")
	})

	executor.DoneWithContext(func(err error, step string) {
		if err != nil {
			fmt.Printf("Cleanup for failed step '%s': %v\n", step, err)
			// Perform cleanup operations based on the failed step
		}
	})

	// Output:
	// Cleanup for failed step 'database_operation': connection timeout
}
