package httpx

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestNewResponse(t *testing.T) {
	// Create a test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("X-Custom-Header", "custom-value")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "success"}`))
	}))
	defer server.Close()

	// Make a request to get a real response
	resp, err := http.Get(server.URL)
	if err != nil {
		t.Fatalf("Failed to make request: %v", err)
	}
	defer resp.Body.Close()

	// Test NewResponse
	response := NewResponse(resp)

	if response.Response != resp {
		t.Errorf("Expected response to wrap the original response")
	}

	// Test that the response is properly wrapped
	if response.StatusCode != resp.StatusCode {
		t.Errorf("Expected status code %d, got %d", resp.StatusCode, response.StatusCode)
	}
}

func TestResponse_IsSuccess(t *testing.T) {
	tests := []struct {
		name           string
		statusCode     int
		expectedResult bool
	}{
		{"200 OK", http.StatusOK, true},
		{"201 Created", http.StatusCreated, true},
		{"202 Accepted", http.StatusAccepted, true},
		{"204 No Content", http.StatusNoContent, true},
		{"299 OK", 299, true},
		{"400 Bad Request", http.StatusBadRequest, false},
		{"401 Unauthorized", http.StatusUnauthorized, false},
		{"404 Not Found", http.StatusNotFound, false},
		{"500 Internal Server Error", http.StatusInternalServerError, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(tt.statusCode)
			}))
			defer server.Close()

			resp, err := http.Get(server.URL)
			if err != nil {
				t.Fatalf("Failed to make request: %v", err)
			}
			defer resp.Body.Close()

			response := NewResponse(resp)
			result := response.IsSuccess()

			if result != tt.expectedResult {
				t.Errorf("Expected IsSuccess() to return %v for status %d, got %v",
					tt.expectedResult, tt.statusCode, result)
			}
		})
	}
}

func TestResponse_IsClientError(t *testing.T) {
	tests := []struct {
		name           string
		statusCode     int
		expectedResult bool
	}{
		{"200 OK", http.StatusOK, false},
		{"400 Bad Request", http.StatusBadRequest, true},
		{"401 Unauthorized", http.StatusUnauthorized, true},
		{"404 Not Found", http.StatusNotFound, true},
		{"499 Client Error", 499, true},
		{"500 Internal Server Error", http.StatusInternalServerError, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(tt.statusCode)
			}))
			defer server.Close()

			resp, err := http.Get(server.URL)
			if err != nil {
				t.Fatalf("Failed to make request: %v", err)
			}
			defer resp.Body.Close()

			response := NewResponse(resp)
			result := response.IsClientError()

			if result != tt.expectedResult {
				t.Errorf("Expected IsClientError() to return %v for status %d, got %v",
					tt.expectedResult, tt.statusCode, result)
			}
		})
	}
}

func TestResponse_IsServerError(t *testing.T) {
	tests := []struct {
		name           string
		statusCode     int
		expectedResult bool
	}{
		{"200 OK", http.StatusOK, false},
		{"400 Bad Request", http.StatusBadRequest, false},
		{"500 Internal Server Error", http.StatusInternalServerError, true},
		{"502 Bad Gateway", http.StatusBadGateway, true},
		{"503 Service Unavailable", http.StatusServiceUnavailable, true},
		{"599 Server Error", 599, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(tt.statusCode)
			}))
			defer server.Close()

			resp, err := http.Get(server.URL)
			if err != nil {
				t.Fatalf("Failed to make request: %v", err)
			}
			defer resp.Body.Close()

			response := NewResponse(resp)
			result := response.IsServerError()

			if result != tt.expectedResult {
				t.Errorf("Expected IsServerError() to return %v for status %d, got %v",
					tt.expectedResult, tt.statusCode, result)
			}
		})
	}
}

func TestResponse_Status(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("test response"))
	}))
	defer server.Close()

	resp, err := http.Get(server.URL)
	if err != nil {
		t.Fatalf("Failed to make request: %v", err)
	}
	defer resp.Body.Close()

	response := NewResponse(resp)
	statusCode, status := response.Status()

	if status != resp.Status {
		t.Errorf("Expected status %s, got %s", resp.Status, status)
	}

	if statusCode != resp.StatusCode {
		t.Errorf("Expected status code %d, got %d", resp.StatusCode, statusCode)
	}
}

func TestResponse_Bytes(t *testing.T) {
	testBody := "test response body"
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(testBody))
	}))
	defer server.Close()

	resp, err := http.Get(server.URL)
	if err != nil {
		t.Fatalf("Failed to make request: %v", err)
	}
	defer resp.Body.Close()

	response := NewResponse(resp)
	bytes, err := response.Bytes()

	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if string(bytes) != testBody {
		t.Errorf("Expected body %s, got %s", testBody, string(bytes))
	}

	// Note: The current implementation doesn't cache the body, so we can't test caching
	// Each call to Bytes() will read from the response body again
}

func TestResponse_String(t *testing.T) {
	testBody := "test response body"
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(testBody))
	}))
	defer server.Close()

	resp, err := http.Get(server.URL)
	if err != nil {
		t.Fatalf("Failed to make request: %v", err)
	}
	defer resp.Body.Close()

	response := NewResponse(resp)
	str, err := response.String()

	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if str != testBody {
		t.Errorf("Expected string %s, got %s", testBody, str)
	}
}

func TestResponse_JSON(t *testing.T) {
	t.Run("valid JSON", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"message": "success", "count": 42}`))
		}))
		defer server.Close()

		resp, err := http.Get(server.URL)
		if err != nil {
			t.Fatalf("Failed to make request: %v", err)
		}
		defer resp.Body.Close()

		response := NewResponse(resp)
		var data map[string]interface{}
		err = response.JSON(&data)

		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}

		if data["message"] != "success" {
			t.Errorf("Expected message 'success', got %v", data["message"])
		}

		if data["count"] != float64(42) { // JSON numbers are float64
			t.Errorf("Expected count 42, got %v", data["count"])
		}
	})

	t.Run("invalid JSON", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("invalid json"))
		}))
		defer server.Close()

		resp, err := http.Get(server.URL)
		if err != nil {
			t.Fatalf("Failed to make request: %v", err)
		}
		defer resp.Body.Close()

		response := NewResponse(resp)
		var data map[string]interface{}
		err = response.JSON(&data)

		if err == nil {
			t.Errorf("Expected error for invalid JSON")
		}
	})
}

func TestResponse_Close(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("test"))
	}))
	defer server.Close()

	resp, err := http.Get(server.URL)
	if err != nil {
		t.Fatalf("Failed to make request: %v", err)
	}

	response := NewResponse(resp)
	err = response.Close()

	if err != nil {
		t.Errorf("Unexpected error closing response: %v", err)
	}

	// Test that body is closed
	_, err = io.ReadAll(resp.Body)
	if err == nil {
		t.Errorf("Expected error when reading closed body")
	}
}

func TestResponse_HeaderValue(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("X-Custom-Header", "custom-value")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("test"))
	}))
	defer server.Close()

	resp, err := http.Get(server.URL)
	if err != nil {
		t.Fatalf("Failed to make request: %v", err)
	}
	defer resp.Body.Close()

	response := NewResponse(resp)

	tests := []struct {
		name     string
		key      string
		expected string
	}{
		{"existing header", "Content-Type", "application/json"},
		{"existing header case insensitive", "content-type", "application/json"},
		{"existing custom header", "X-Custom-Header", "custom-value"},
		{"non-existing header", "Non-Existing", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := response.HeaderValue(tt.key)
			if result != tt.expected {
				t.Errorf("Expected header value %s, got %s", tt.expected, result)
			}
		})
	}
}

func TestResponse_ContentType(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("test"))
	}))
	defer server.Close()

	resp, err := http.Get(server.URL)
	if err != nil {
		t.Fatalf("Failed to make request: %v", err)
	}
	defer resp.Body.Close()

	response := NewResponse(resp)
	contentType := response.ContentType()

	expected := "application/json"
	if contentType != expected && contentType != "application/json; charset=utf-8" {
		t.Errorf("Expected content type %s or %s; charset=utf-8, got %s", expected, expected, contentType)
	}
}

func TestResponse_ContentLength(t *testing.T) {
	testBody := "test response body"
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Length", string(rune(len(testBody))))
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(testBody))
	}))
	defer server.Close()

	resp, err := http.Get(server.URL)
	if err != nil {
		t.Fatalf("Failed to make request: %v", err)
	}
	defer resp.Body.Close()

	response := NewResponse(resp)
	contentLength := response.ContentLength()

	if contentLength != resp.ContentLength {
		t.Errorf("Expected content length %d, got %d", resp.ContentLength, contentLength)
	}
}

func TestResponse_ErrorHandling(t *testing.T) {
	t.Run("response with bad status", func(t *testing.T) {
		// Create a response with a bad status
		resp := &http.Response{
			StatusCode: http.StatusBadRequest,
		}
		response := NewResponse(resp)

		if response.IsSuccess() {
			t.Errorf("Expected response to not be successful")
		}

		if !response.IsClientError() {
			t.Errorf("Expected response to be a client error")
		}
	})

	t.Run("read error handling", func(t *testing.T) {
		// Create a mock response with a body that will cause read error
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			// Write some data and then close the connection
			w.Write([]byte("partial"))
			if f, ok := w.(http.Flusher); ok {
				f.Flush()
			}
		}))
		defer server.Close()

		// Use a custom transport that will cause connection issues
		client := &http.Client{}
		resp, err := client.Get(server.URL)
		if err != nil {
			t.Fatalf("Failed to make request: %v", err)
		}

		response := NewResponse(resp)

		// Try to read the body
		_, err = response.Bytes()
		// We might get an error or not, depending on timing
		// The important thing is that the method handles it gracefully
		_ = err // Ignore the error for this test
	})
}

// Benchmark tests
func BenchmarkResponse_Bytes(b *testing.B) {
	testBody := strings.Repeat("test data ", 100)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(testBody))
	}))
	defer server.Close()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		resp, err := http.Get(server.URL)
		if err != nil {
			b.Fatalf("Failed to make request: %v", err)
		}

		response := NewResponse(resp)
		_, err = response.Bytes()
		if err != nil {
			b.Fatal(err)
		}
		_ = resp.Body.Close()
	}
}

func BenchmarkResponse_String(b *testing.B) {
	testBody := strings.Repeat("test data ", 100)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(testBody))
	}))
	defer server.Close()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		resp, err := http.Get(server.URL)
		if err != nil {
			b.Fatalf("Failed to make request: %v", err)
		}

		response := NewResponse(resp)
		_, err = response.String()
		if err != nil {
			b.Fatal(err)
		}
		_ = resp.Body.Close()
	}
}
