package errorx

import (
	"fmt"
)

// ChainExecutor is an executor that implements a fail-fast pattern for sequential operations.
// It centralizes error handling to avoid redundant error processing. Once an error occurs,
// all subsequent operations are skipped until the error is handled.
//
// The ChainExecutor follows these principles:
//   - Fail-fast: Stop execution on the first error
//   - Centralized: All errors are collected and handled in one place
//   - Chainable: Supports fluent API design with method chaining
//   - Context-aware: Tracks which step failed and preserves intermediate results
//
// Basic Usage:
//
//	executor := errorx.NewChainExecutor()
//	executor.Do(func() error { return step1() }).
//		Do(func() error { return step2() }).
//		Do(func() error { return step3() })
//
//	if executor.HasErr() {
//		fmt.Printf("Failed at step: %s, Error: %v\n", executor.FailedStep(), executor.Err())
//		return executor.Err()
//	}
//
// Advanced Usage with Context:
//
//	executor := errorx.NewChainExecutor()
//	var result1, result2 string
//
//	executor.DoWithContext("validate", func() error {
//		// validation logic
//		return nil
//	}).DoWithResult("process", func() (string, error) {
//		// processing logic
//		return "processed", nil
//	}).DoWithContext("save", func() error {
//		// save logic using result2
//		return nil
//	})
//
//	if executor.HasErr() {
//		return fmt.Errorf("operation failed at %s: %w", executor.FailedStep(), executor.Err())
//	}
type ChainExecutor struct {
	err        error
	failedStep string
	results    map[string]interface{}
}

// NewChainExecutor creates a new ChainExecutor instance.
func NewChainExecutor() *ChainExecutor {
	return &ChainExecutor{
		results: make(map[string]interface{}),
	}
}

// Do executes the given function if no error has occurred yet.
// If an error has already occurred, the function is skipped.
// Returns the executor for method chaining.
func (h *ChainExecutor) Do(f func() error) *ChainExecutor {
	if h.err == nil && f != nil {
		h.err = f()
	}
	return h
}

// DoWithContext executes the given function with a step name for better error tracking.
// If an error has already occurred, the function is skipped.
// Returns the executor for method chaining.
func (h *ChainExecutor) DoWithContext(step string, f func() error) *ChainExecutor {
	if h.err == nil && f != nil {
		h.err = f()
		if h.err != nil {
			h.failedStep = step
		}
	}
	return h
}

// DoWithResult executes a function that returns a result and an error.
// The result is stored and can be retrieved later using GetResult.
// If an error has already occurred, the function is skipped.
// Returns the executor for method chaining.
func (h *ChainExecutor) DoWithResult(step string, f func() (interface{}, error)) *ChainExecutor {
	if h.err == nil && f != nil {
		result, err := f()
		if err != nil {
			h.err = err
			h.failedStep = step
		} else {
			h.results[step] = result
		}
	}
	return h
}

// DoWithTypedResult executes a function that returns a typed result and an error.
// The result is stored with the given key and can be retrieved later using GetTypedResult.
// If an error has already occurred, the function is skipped.
// Returns the executor for method chaining.
func DoWithTypedResult[T any](h *ChainExecutor, step string, f func() (T, error)) *ChainExecutor {
	if h.err == nil && f != nil {
		result, err := f()
		if err != nil {
			h.err = err
			h.failedStep = step
		} else {
			h.results[step] = result
		}
	}
	return h
}

// HasErr returns true if the executor has encountered an error.
func (h *ChainExecutor) HasErr() bool {
	return h.err != nil
}

// Err returns the first error that occurred, or nil if no error has occurred.
func (h *ChainExecutor) Err() error {
	return h.err
}

// FailedStep returns the name of the step where the error occurred.
// Returns an empty string if no error has occurred.
func (h *ChainExecutor) FailedStep() string {
	return h.failedStep
}

// GetResult retrieves a result stored by DoWithResult or DoWithTypedResult.
// Returns the result and true if found, nil and false otherwise.
func (h *ChainExecutor) GetResult(key string) (interface{}, bool) {
	result, exists := h.results[key]
	return result, exists
}

// GetTypedResult retrieves a typed result stored by DoWithTypedResult.
// Returns the result and true if found and of the correct type, zero value and false otherwise.
func GetTypedResult[T any](h *ChainExecutor, key string) (T, bool) {
	var zero T
	result, exists := h.results[key]
	if !exists {
		return zero, false
	}

	typedResult, ok := result.(T)
	if !ok {
		return zero, false
	}

	return typedResult, true
}

// Done executes the given function with the current error state.
// This is useful for cleanup operations or final error handling.
func (h *ChainExecutor) Done(f func(err error)) {
	f(h.err)
}

// DoneWithContext executes the given function with both error and step information.
// This provides more context for error handling and cleanup operations.
func (h *ChainExecutor) DoneWithContext(f func(err error, failedStep string)) {
	f(h.err, h.failedStep)
}

// IsPreferredErr returns true if the current error is a PreferredError.
func (h *ChainExecutor) IsPreferredErr() bool {
	return IsPreferred(h.err)
}

// PreferredOr returns the current error if it's preferred, otherwise returns the given error.
func (h *ChainExecutor) PreferredOr(err error) error {
	if h.IsPreferredErr() {
		return h.err
	}
	return err
}

// Reset clears the error state and results, allowing the executor to be reused.
// This should be used with caution as it discards all previous state.
func (h *ChainExecutor) Reset() *ChainExecutor {
	h.err = nil
	h.failedStep = ""
	h.results = make(map[string]interface{})
	return h
}

// ErrorInfo provides detailed information about the error state.
type ErrorInfo struct {
	Err        error
	FailedStep string
	HasError   bool
}

// GetErrorInfo returns comprehensive error information.
func (h *ChainExecutor) GetErrorInfo() ErrorInfo {
	return ErrorInfo{
		Err:        h.err,
		FailedStep: h.failedStep,
		HasError:   h.err != nil,
	}
}

// String returns a string representation of the executor's state.
func (h *ChainExecutor) String() string {
	if h.err == nil {
		return "ChainExecutor{no error}"
	}
	if h.failedStep != "" {
		return fmt.Sprintf("ChainExecutor{error: %v, failed at: %s}", h.err, h.failedStep)
	}
	return fmt.Sprintf("ChainExecutor{error: %v}", h.err)
}
