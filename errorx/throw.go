package errorx

import (
	"fmt"
)

// Throw panics with the given error if it's not nil.
// This function is useful for propagating errors as panics in scenarios where
// error handling through panic/recover is preferred over traditional error returns.
//
// Parameters:
//   - err: The error to panic with. If nil, no panic occurs.
//
// Example:
//
//	func riskyOperation() {
//		err := someOperation()
//		Throw(err) // Panics if err is not nil
//		// Continue with normal execution if err is nil
//	}
func Throw(err error) {
	if err != nil {
		panic(err)
	}
}

// Throwf panics with a custom formatted message if the given error is not nil.
// The formatted message is created using the provided format string and arguments,
// similar to fmt.Sprintf. If the error is nil, no panic occurs and the format
// string and arguments are ignored.
//
// Parameters:
//   - err: The error to check. If nil, no panic occurs.
//   - format: The format string for the panic message.
//   - v: Arguments to format the message with.
//
// Example:
//
//	func riskyOperation(id int) {
//		err := someOperation()
//		Throwf(err, "operation failed for ID %d", id)
//		// Panics with "operation failed for ID 123" if err is not nil
//	}
func Throwf(err error, format string, v ...any) {
	if err != nil {
		panic(fmt.Sprintf(format, v...))
	}
}

// Throwv returns the given value if the error is nil, otherwise panics with the error.
// This function is useful for functions that return (value, error) pairs where you
// want to panic on error and continue with the value on success.
//
// Parameters:
//   - v: The value to return if error is nil.
//   - err: The error to check. If nil, returns v. If not nil, panics with err.
//
// Returns:
//   - T: The value v if err is nil.
//
// Example:
//
//	result := Throwv(42, nil) // Returns 42
//	// result := Throwv(0, errors.New("failed")) // Panics
func Throwv[T any](v T, err error) T {
	Throw(err)
	return v
}
