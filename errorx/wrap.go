package errorx

import "fmt"

// Wrap wraps the given error with additional context.
// If the error is nil, it returns nil. Otherwise, it wraps the error
// using fmt.Errorf with %w verb to preserve error chain information.
//
// This function is useful when you want to add context to an error
// without losing the original error information, allowing errors.Is
// and errors.As to work correctly.
//
// Parameters:
//   - err: The error to wrap. If nil, returns nil.
//
// Returns:
//   - error: A wrapped error that preserves the original error chain.
//
// Example:
//
//	originalErr := errors.New("database connection failed")
//	wrappedErr := Wrap(originalErr)
//	fmt.Println(errors.Is(wrappedErr, originalErr)) // true
func Wrap(err error) error {
	if err == nil {
		return nil
	}
	return fmt.Errorf("%w", err)
}

// Wrapf wraps the given error with a formatted message.
// If the error is nil, it returns nil. Otherwise, it wraps the error
// with the formatted message, preserving error chain information.
//
// The format string and arguments work the same way as fmt.Sprintf.
// The wrapped error will be in the format: "message: original error".
// If no format arguments are provided, the format string is used as-is.
//
// Parameters:
//   - err: The error to wrap. If nil, returns nil.
//   - format: The format string for the wrapper message.
//   - args: Arguments for the format string (optional).
//
// Returns:
//   - error: A wrapped error with the formatted message.
//
// Example:
//
//	originalErr := errors.New("connection timeout")
//	wrappedErr := Wrapf(originalErr, "failed to connect to %s", "database")
//	// wrappedErr.Error() will be "failed to connect to database: connection timeout"
//	fmt.Println(errors.Is(wrappedErr, originalErr)) // true
//
//	// Simple message without formatting
//	wrappedErr2 := Wrapf(originalErr, "database error")
//	// wrappedErr2.Error() will be "database error: connection timeout"
func Wrapf(err error, format string, args ...any) error {
	if err == nil {
		return nil
	}

	if len(args) == 0 {
		return fmt.Errorf("%s: %w", format, err)
	}
	return fmt.Errorf("%s: %w", fmt.Sprintf(format, args...), err)
}
