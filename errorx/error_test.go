package errorx

import (
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPreferred2(t *testing.T) {
	err := fmt.Errorf("demo error")
	b := IsPreferred(err)
	fmt.Println(b) // false
	err = NewPreferredErr(err)
	fmt.Println(IsPreferred(err)) // true
}

func TestIsPreferred(t *testing.T) {
	err := fmt.Errorf("haha")
	b := IsPreferred(err) // false
	assert.True(t, !b)

	err = NewPreferredErr(fmt.Errorf("haha"))
	b = IsPreferred(err)
	assert.True(t, b)

	err = Prefer500("haha")
	b = IsPreferred(err)
	assert.True(t, b)
}

func ExampleIsPreferred() {
	err := fmt.Errorf("demo error")
	b := IsPreferred(err)
	fmt.Println(b) // false
	err = NewPreferredErr(err)
	fmt.Println(IsPreferred(err)) // true
}

func TestPreferredError(t *testing.T) {
	t.Run("create PreferredError with SetCode", func(t *testing.T) {
		originalErr := errors.New("test error")
		preferredErr := &PreferredError{error: originalErr}
		_ = preferredErr.SetCode(http.StatusBadRequest)

		assert.Equal(t, http.StatusBadRequest, preferredErr.Code())
		assert.Equal(t, "test error", preferredErr.Error())
		assert.True(t, IsPreferred(preferredErr))
	})

	t.Run("PreferredError implements error interface", func(t *testing.T) {
		originalErr := errors.New("test error")
		preferredErr := &PreferredError{error: originalErr, code: http.StatusInternalServerError}

		var err error = preferredErr
		assert.Equal(t, "test error", err.Error())
	})
}

func TestNewPreferredErr(t *testing.T) {
	t.Run("create PreferredError with nil error", func(t *testing.T) {
		err := NewPreferredErr(nil)
		preferredErr, ok := err.(*PreferredError)
		assert.True(t, ok)
		assert.Equal(t, 0, preferredErr.Code())
		assert.Equal(t, "<nil>", preferredErr.Error())
	})

	t.Run("create PreferredError with error", func(t *testing.T) {
		originalErr := errors.New("test error")
		err := NewPreferredErr(originalErr)
		preferredErr, ok := err.(*PreferredError)
		assert.True(t, ok)
		assert.Equal(t, 0, preferredErr.Code())
		assert.Equal(t, "test error", preferredErr.Error())
	})
}

func TestNewPreferredErrCode(t *testing.T) {
	t.Run("create PreferredError with error and code", func(t *testing.T) {
		originalErr := errors.New("test error")
		err := NewPreferredErrCode(originalErr, http.StatusBadRequest)
		preferredErr, ok := err.(*PreferredError)
		assert.True(t, ok)
		assert.Equal(t, http.StatusBadRequest, preferredErr.Code())
		assert.Equal(t, "test error", preferredErr.Error())
	})
}

func TestNewPreferredErrf(t *testing.T) {
	t.Run("create PreferredError with format string", func(t *testing.T) {
		err := NewPreferredErrf("error: %s", "test")
		preferredErr, ok := err.(*PreferredError)
		assert.True(t, ok)
		assert.Equal(t, 0, preferredErr.Code())
		assert.Equal(t, "error: test", preferredErr.Error())
	})
}

func TestNewPreferredCodeErrf(t *testing.T) {
	t.Run("create PreferredError with code and format string", func(t *testing.T) {
		err := NewPreferredCodeErrf(http.StatusBadRequest, "error: %s", "test")
		preferredErr, ok := err.(*PreferredError)
		assert.True(t, ok)
		assert.Equal(t, http.StatusBadRequest, preferredErr.Code())
		assert.Equal(t, "error: test", preferredErr.Error())
	})
}

func TestPreferFunctions(t *testing.T) {
	testCases := []struct {
		name     string
		function func(string, ...any) error
		code     int
		message  string
	}{
		{"Prefer400", Prefer400, http.StatusBadRequest, "bad request"},
		{"Prefer401", Prefer401, http.StatusUnauthorized, "unauthorized"},
		{"Prefer403", Prefer403, http.StatusForbidden, "forbidden"},
		{"Prefer415", Prefer415, http.StatusUnsupportedMediaType, "unsupported media type"},
		{"Prefer423", Prefer423, http.StatusLocked, "locked"},
		{"Prefer426", Prefer426, http.StatusUpgradeRequired, "upgrade required"},
		{"Prefer429", Prefer429, http.StatusTooManyRequests, "too many requests"},
		{"Prefer500", Prefer500, http.StatusInternalServerError, "internal server error"},
		{"Prefer501", Prefer501, http.StatusNotImplemented, "not implemented"},
		{"Prefer503", Prefer503, http.StatusServiceUnavailable, "service unavailable"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.function("test %s", "message")
			preferredErr, ok := err.(*PreferredError)
			assert.True(t, ok)
			assert.Equal(t, tc.code, preferredErr.Code())
			assert.Equal(t, "test message", preferredErr.Error())
			assert.True(t, IsPreferred(err))
		})
	}
}

func TestIsPreferredDetailed(t *testing.T) {
	t.Run("nil error", func(t *testing.T) {
		assert.False(t, IsPreferred(nil))
	})

	t.Run("regular error", func(t *testing.T) {
		err := errors.New("regular error")
		assert.False(t, IsPreferred(err))
	})

	t.Run("PreferredError", func(t *testing.T) {
		err := Prefer400("test error")
		assert.True(t, IsPreferred(err))
	})

	t.Run("wrapped PreferredError", func(t *testing.T) {
		err := Prefer400("test error")
		wrapped := errors.New("wrapped: " + err.Error())
		assert.False(t, IsPreferred(wrapped))
	})
}

func TestNew(t *testing.T) {
	t.Run("create error without arguments", func(t *testing.T) {
		err := New("test error")
		assert.Equal(t, "test error", err.Error())
		assert.False(t, IsPreferred(err))
	})

	t.Run("create error with arguments", func(t *testing.T) {
		err := New("error: %s", "test")
		assert.Equal(t, "error: test", err.Error())
		assert.False(t, IsPreferred(err))
	})

	t.Run("create error with multiple arguments", func(t *testing.T) {
		err := New("error %s: %d", "code", 123)
		assert.Equal(t, "error code: 123", err.Error())
		assert.False(t, IsPreferred(err))
	})

	t.Run("create error with empty string", func(t *testing.T) {
		err := New("")
		assert.Equal(t, "", err.Error())
	})
}

func TestPreferredErrorEdgeCases(t *testing.T) {
	t.Run("empty error message", func(t *testing.T) {
		err := Prefer400("")
		preferredErr, ok := err.(*PreferredError)
		assert.True(t, ok)
		assert.Equal(t, http.StatusBadRequest, preferredErr.Code())
		assert.Equal(t, "", preferredErr.Error())
	})

	t.Run("error with nil values", func(t *testing.T) {
		err := NewPreferredErrCode(nil, http.StatusInternalServerError)
		preferredErr, ok := err.(*PreferredError)
		assert.True(t, ok)
		assert.Equal(t, http.StatusInternalServerError, preferredErr.Code())
		assert.Equal(t, "<nil>", preferredErr.Error())
	})

	t.Run("format string with no placeholders", func(t *testing.T) {
		err := Prefer500("simple message")
		preferredErr, ok := err.(*PreferredError)
		assert.True(t, ok)
		assert.Equal(t, http.StatusInternalServerError, preferredErr.Code())
		assert.Equal(t, "simple message", preferredErr.Error())
	})
}

func BenchmarkPreferFunctions(b *testing.B) {
	b.Run("Prefer400", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = Prefer400("test error %d", i)
		}
	})

	b.Run("Prefer500", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = Prefer500("test error %d", i)
		}
	})

	b.Run("IsPreferred", func(b *testing.B) {
		err := Prefer400("test error")
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = IsPreferred(err)
		}
	})

	b.Run("New", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = New("test error %d", i)
		}
	})
}

// Example functions for documentation
func ExamplePreferredError() {
	// Create a preferred error
	err := Prefer400("invalid input: %s", "missing field")

	// Check if it's a preferred error
	if IsPreferred(err) {
		preferredErr := err.(*PreferredError)
		fmt.Printf("HTTP Status: %d\n", preferredErr.Code())
		fmt.Printf("Error: %s\n", preferredErr.Error())
	}

	// Output:
	// HTTP Status: 400
	// Error: invalid input: missing field
}

func ExampleNew() {
	// Create a regular error
	err := New("operation failed: %s", "timeout")
	fmt.Printf("Error: %s\n", err.Error())

	// Output:
	// Error: operation failed: timeout
}

func ExamplePrefer400() {
	// Create a 400 Bad Request error
	err := Prefer400("invalid JSON format")
	preferredErr := err.(*PreferredError)
	fmt.Printf("Status: %d, Message: %s\n", preferredErr.Code(), preferredErr.Error())

	// Output:
	// Status: 400, Message: invalid JSON format
}

func ExamplePrefer401() {
	// Create a 401 Unauthorized error
	err := Prefer401("authentication required")
	preferredErr := err.(*PreferredError)
	fmt.Printf("Status: %d, Message: %s\n", preferredErr.Code(), preferredErr.Error())

	// Output:
	// Status: 401, Message: authentication required
}

func ExamplePrefer500() {
	// Create a 500 Internal Server Error
	err := Prefer500("database connection failed")
	preferredErr := err.(*PreferredError)
	fmt.Printf("Status: %d, Message: %s\n", preferredErr.Code(), preferredErr.Error())

	// Output:
	// Status: 500, Message: database connection failed
}
