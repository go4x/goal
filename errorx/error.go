// Package errorx provides enhanced error handling utilities for Go applications.
//
// This package offers a comprehensive set of error handling tools designed to
// improve error management in Go applications, particularly web services and
// APIs that need to handle HTTP status codes and centralized error processing.
//
// Key Features:
//
//   - PreferredError: Custom error type with HTTP status codes for web applications
//   - ErrorHandler: Centralized error handling to avoid redundant error processing
//   - Panic recovery utilities with cleanup functions and context support
//   - Error wrapping and formatting utilities with additional context
//   - HTTP error factory functions for common status codes (400, 401, 403, 429, 500)
//   - Panic-based error handling for simplified error propagation
//
// Basic Usage:
//
//	// Create preferred errors with HTTP status codes
//	err := errorx.Prefer500("service temporarily unavailable")
//	err = errorx.Prefer400("invalid request format")
//
//	// Use error handler for centralized error processing
//	handler := errorx.NewHandler()
//	handler.Do(func() error { return riskyOperation1() })
//	handler.Do(func() error { return riskyOperation2() })
//	if handler.HasErr() {
//		log.Printf("Operation failed: %v", handler.Err())
//	}
//
//	// Recover from panics with cleanup functions
//	defer errorx.Recover(logger, func() {
//		cleanupResources()
//	})
//
//	// Wrap errors with additional context
//	wrapped := errorx.Wrapf(originalErr, "failed to process user request")
//
//	// Panic-based error handling (use with caution)
//	result := errorx.Throwv(riskyOperation())
//
// HTTP Status Code Support:
//
// The PreferredError type supports standard HTTP status codes:
//   - 400 Bad Request
//   - 401 Unauthorized
//   - 403 Forbidden
//   - 415 Unsupported Media Type
//   - 423 Locked
//   - 426 Upgrade Required
//   - 429 Too Many Requests
//   - 500 Internal Server Error
//   - 501 Not Implemented
//   - 503 Service Unavailable
//
// Error Handler Design:
//
// The ErrorHandler implements a fail-fast pattern where the first error
// stops subsequent operations. This is useful for operations that should
// not continue if any step fails. Use the Done method to add cleanup hooks.
//
// Panic Recovery:
//
// Both Recover and RecoverCtx functions provide panic recovery with optional
// cleanup functions. RecoverCtx additionally respects context cancellation
// and timeout during cleanup operations.
package errorx

import (
	"errors"
	"fmt"
	"net/http"
)

// PreferredError is a struct to describe a preferred error, which will be processed specially, such as
// response it to http
type PreferredError struct {
	error
	code int
}

// SetCode set a http status code to PreferredError. To maintain encapsulation, prefer to use
// NewPreferredErrCode/NewPreferredCodeErrf to create a PreferredError
func (e *PreferredError) SetCode(code int) *PreferredError {
	e.code = code
	return e
}

func (e *PreferredError) Error() string {
	if e.error == nil {
		return "<nil>"
	}
	return e.error.Error()
}

func (e *PreferredError) Code() int {
	return e.code
}

// NewPreferredErr create a PreferredError with given error, code is 0, call SetCode to set a http status code
// Deprecated: prefer to NewPreferredErrCode
func NewPreferredErr(err error) error {
	return &PreferredError{error: err}
}

// NewPreferredErrCode create a PreferredError with given error and code
func NewPreferredErrCode(err error, code int) error {
	return &PreferredError{error: err, code: code}
}

// NewPreferredErrf create a PreferredError with error message, code is 0, call SetCode to set a http status code
// Deprecated: prefer to NewPreferredErrCodeErrf
func NewPreferredErrf(format string, args ...any) error {
	return &PreferredError{error: fmt.Errorf(format, args...)}
}

// NewPreferredCodeErrf create a PreferredError with given code and error message
func NewPreferredCodeErrf(code int, format string, args ...any) error {
	return &PreferredError{code: code, error: fmt.Errorf(format, args...)}
}

// Prefer400 create a PreferredError with http.StatusBadRequest
func Prefer400(format string, args ...any) error {
	return &PreferredError{code: http.StatusBadRequest, error: fmt.Errorf(format, args...)}
}

// Prefer401 create a PreferredError with http.StatusUnauthorized
func Prefer401(format string, args ...any) error {
	return &PreferredError{code: http.StatusUnauthorized, error: fmt.Errorf(format, args...)}
}

// Prefer403 create a PreferredError with http.StatusForbidden
func Prefer403(format string, args ...any) error {
	return &PreferredError{code: http.StatusForbidden, error: fmt.Errorf(format, args...)}
}

// Prefer415 create a PreferredError with http.StatusUnsupportedMediaType
func Prefer415(format string, args ...any) error {
	return &PreferredError{code: http.StatusUnsupportedMediaType, error: fmt.Errorf(format, args...)}
}

// Prefer423 create a PreferredError with http.StatusLocked
func Prefer423(format string, args ...any) error {
	return &PreferredError{code: http.StatusLocked, error: fmt.Errorf(format, args...)}
}

// Prefer426 create a PreferredError with http.StatusUpgradeRequired
func Prefer426(format string, args ...any) error {
	return &PreferredError{code: http.StatusUpgradeRequired, error: fmt.Errorf(format, args...)}
}

// Prefer429 create a PreferredError with http.StatusTooManyRequests
func Prefer429(format string, args ...any) error {
	return &PreferredError{code: http.StatusTooManyRequests, error: fmt.Errorf(format, args...)}
}

// Prefer500 create a PreferredError with http.StatusInternalServerError
func Prefer500(format string, args ...any) error {
	return &PreferredError{code: http.StatusInternalServerError, error: fmt.Errorf(format, args...)}
}

// Prefer501 create a PreferredError with http.StatusNotImplemented
func Prefer501(format string, args ...any) error {
	return &PreferredError{code: http.StatusNotImplemented, error: fmt.Errorf(format, args...)}
}

// Prefer503 create a PreferredError with http.StatusServiceUnavailable
func Prefer503(format string, args ...any) error {
	return &PreferredError{code: http.StatusServiceUnavailable, error: fmt.Errorf(format, args...)}
}

// IsPreferred checks if the given err is a PreferredError
func IsPreferred(err error) bool {
	if err == nil {
		return false
	}
	_, ok := err.(*PreferredError)
	return ok
}

// New method will create an error using the given string and arguments.
func New(s string, v ...any) error {
	if len(v) == 0 {
		return errors.New(s)
	}
	return fmt.Errorf(s, v...)
}
