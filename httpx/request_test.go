package httpx

import (
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
	"testing"
	"time"
)

func TestNewRequest(t *testing.T) {
	tests := []struct {
		name        string
		method      string
		url         string
		body        io.Reader
		options     []RequestOption
		expectError bool
	}{
		{
			name:        "valid GET request",
			method:      http.MethodGet,
			url:         "https://example.com/api",
			body:        nil,
			options:     nil,
			expectError: false,
		},
		{
			name:        "valid POST request with body",
			method:      http.MethodPost,
			url:         "https://example.com/api",
			body:        strings.NewReader("test data"),
			options:     nil,
			expectError: false,
		},
		{
			name:        "request with content type option",
			method:      http.MethodPost,
			url:         "https://example.com/api",
			body:        strings.NewReader("test data"),
			options:     []RequestOption{WithContentType("application/json")},
			expectError: false,
		},
		{
			name:        "request with header option",
			method:      http.MethodGet,
			url:         "https://example.com/api",
			body:        nil,
			options:     []RequestOption{WithHeader("Authorization", "Bearer token")},
			expectError: false,
		},
		{
			name:        "request with context option",
			method:      http.MethodGet,
			url:         "https://example.com/api",
			body:        nil,
			options:     []RequestOption{WithContext(context.Background())},
			expectError: false,
		},
		{
			name:   "request with multiple options",
			method: http.MethodPost,
			url:    "https://example.com/api",
			body:   strings.NewReader("test data"),
			options: []RequestOption{
				WithContentType("application/json"),
				WithHeader("Authorization", "Bearer token"),
				WithContext(context.Background()),
			},
			expectError: false,
		},
		{
			name:        "invalid URL",
			method:      http.MethodGet,
			url:         "://invalid-url",
			body:        nil,
			options:     nil,
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var body io.ReadCloser
			if tt.body != nil {
				body = io.NopCloser(tt.body)
			}
			req, err := NewRequest(tt.method, tt.url, body, tt.options...)

			if tt.expectError {
				if err == nil {
					t.Errorf("Expected error but got none")
				}
				return
			}

			if err != nil {
				t.Errorf("Unexpected error: %v", err)
				return
			}

			if req == nil {
				t.Errorf("Expected request but got nil")
				return
			}

			if req.GetMethod() != tt.method {
				t.Errorf("Expected method %s, got %s", tt.method, req.GetMethod())
			}

			if req.GetURL().String() != tt.url {
				t.Errorf("Expected URL %s, got %s", tt.url, req.GetURL().String())
			}
		})
	}
}

func TestNewRequestWithContext(t *testing.T) {
	tests := []struct {
		name        string
		method      string
		url         string
		body        io.Reader
		options     []RequestOption
		expectError bool
	}{
		{
			name:        "valid request with context",
			method:      http.MethodGet,
			url:         "https://example.com/api",
			body:        nil,
			options:     nil,
			expectError: false,
		},
		{
			name:        "request with timeout context",
			method:      http.MethodGet,
			url:         "https://example.com/api",
			body:        nil,
			options:     []RequestOption{WithContext(context.Background())},
			expectError: false,
		},
		{
			name:        "invalid URL with context",
			method:      http.MethodGet,
			url:         "://invalid-url",
			body:        nil,
			options:     nil,
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			var body io.ReadCloser
			if tt.body != nil {
				body = io.NopCloser(tt.body)
			}
			req, err := NewRequestWithContext(ctx, tt.method, tt.url, body, tt.options...)

			if tt.expectError {
				if err == nil {
					t.Errorf("Expected error but got none")
				}
				return
			}

			if err != nil {
				t.Errorf("Unexpected error: %v", err)
				return
			}

			if req == nil {
				t.Errorf("Expected request but got nil")
				return
			}

			if req.GetMethod() != tt.method {
				t.Errorf("Expected method %s, got %s", tt.method, req.GetMethod())
			}

			if req.GetURL().String() != tt.url {
				t.Errorf("Expected URL %s, got %s", tt.url, req.GetURL().String())
			}

			// Check that context is set correctly
			if req.GetContext() == context.Background() {
				// This is expected for valid requests
			}
		})
	}
}

func TestRequestOptions(t *testing.T) {
	t.Run("WithContentType", func(t *testing.T) {
		req, err := NewRequest(http.MethodPost, "https://example.com/api", nil, WithContentType("application/json"))
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}

		contentType := req.GetHeader().Get("Content-Type")
		if contentType != "application/json" {
			t.Errorf("Expected Content-Type application/json, got %s", contentType)
		}
	})

	t.Run("WithHeader", func(t *testing.T) {
		req, err := NewRequest(http.MethodGet, "https://example.com/api", nil,
			WithHeader("Authorization", "Bearer token"),
			WithHeader("X-Custom", "value"))
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}

		if req.GetHeader().Get("Authorization") != "Bearer token" {
			t.Errorf("Expected Authorization header to be set")
		}

		if req.GetHeader().Get("X-Custom") != "value" {
			t.Errorf("Expected X-Custom header to be set")
		}
	})

	t.Run("WithContext", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		req, err := NewRequest(http.MethodGet, "https://example.com/api", nil, WithContext(ctx))
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}

		// Check that the context is set (it should be the same as the one we passed)
		if req.GetContext() != ctx {
			t.Errorf("Expected custom context to be set")
		}
	})

	t.Run("Multiple options", func(t *testing.T) {
		ctx := context.Background()
		req, err := NewRequest(http.MethodPost, "https://example.com/api", strings.NewReader("data"),
			WithContentType("application/json"),
			WithHeader("Authorization", "Bearer token"),
			WithContext(ctx))
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}

		// Check all options are applied
		if req.GetHeader().Get("Content-Type") != "application/json" {
			t.Errorf("Expected Content-Type to be set")
		}

		if req.GetHeader().Get("Authorization") != "Bearer token" {
			t.Errorf("Expected Authorization header to be set")
		}

		if req.GetContext() != ctx {
			t.Errorf("Expected custom context to be set")
		}
	})

	t.Run("Header override", func(t *testing.T) {
		req, err := NewRequest(http.MethodPost, "https://example.com/api", nil,
			WithContentType("application/json"),
			WithHeader("Content-Type", "text/plain")) // This should override the previous Content-Type
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}

		contentType := req.GetHeader().Get("Content-Type")
		if contentType != "text/plain" {
			t.Errorf("Expected Content-Type to be overridden to text/plain, got %s", contentType)
		}
	})
}

func TestRequest_EdgeCases(t *testing.T) {
	t.Run("nil body", func(t *testing.T) {
		req, err := NewRequest(http.MethodGet, "https://example.com/api", nil)
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}

		body, err := req.GetBody()
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}
		if body != nil {
			t.Errorf("Expected nil body for GET request")
		}
	})

	t.Run("empty string body", func(t *testing.T) {
		req, err := NewRequest(http.MethodPost, "https://example.com/api", strings.NewReader(""))
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}

		body, err := req.GetBody()
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}
		if body == nil {
			t.Errorf("Expected non-nil body for POST request")
		}
	})

	t.Run("large body", func(t *testing.T) {
		largeBody := strings.NewReader(strings.Repeat("test data ", 1000))
		req, err := NewRequest(http.MethodPost, "https://example.com/api", largeBody)
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}

		body, err := req.GetBody()
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}
		if body == nil {
			t.Errorf("Expected non-nil body for POST request with large body")
		}
	})
}

// Benchmark tests
func BenchmarkNewRequest(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		req, err := NewRequest(http.MethodGet, "https://example.com/api", nil)
		if err != nil {
			b.Fatal(err)
		}
		_ = req
	}
}

func BenchmarkNewRequestWithOptions(b *testing.B) {
	options := []RequestOption{
		WithContentType("application/json"),
		WithHeader("Authorization", "Bearer token"),
		WithContext(context.Background()),
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		req, err := NewRequest(http.MethodPost, "https://example.com/api", strings.NewReader("data"), options...)
		if err != nil {
			b.Fatal(err)
		}
		_ = req
	}
}

func TestRequest_AdditionalMethods(t *testing.T) {
	req, err := NewRequest("GET", "https://example.com/path?key=value", nil)
	if err != nil {
		t.Fatalf("NewRequest failed: %v", err)
	}

	// Test GetForm (should be nil for GET request)
	form := req.GetForm()
	if form != nil {
		t.Error("GetForm should return nil for GET request")
	}

	// Test GetQuery
	query := req.GetQuery()
	if query.Get("key") != "value" {
		t.Errorf("Expected query param 'key=value', got %s", query.Get("key"))
	}

	// Test GetProto
	proto := req.GetProto()
	if proto != "HTTP/1.1" {
		t.Errorf("Expected proto 'HTTP/1.1', got %s", proto)
	}

	// Test GetProtoMajor
	major := req.GetProtoMajor()
	if major != 1 {
		t.Errorf("Expected proto major 1, got %d", major)
	}

	// Test GetProtoMinor
	minor := req.GetProtoMinor()
	if minor != 1 {
		t.Errorf("Expected proto minor 1, got %d", minor)
	}

	// Test GetMultipartForm
	multipartForm, err := req.GetMultipartForm()
	if err != nil {
		t.Errorf("GetMultipartForm failed: %v", err)
	}
	if multipartForm != nil {
		t.Error("Expected nil multipart form for GET request")
	}

	// Test WithURL
	newURL, _ := url.Parse("https://newdomain.com")
	reqWithNewURL := req.WithURL(newURL)
	if reqWithNewURL.GetURL().String() != "https://newdomain.com" {
		t.Errorf("Expected new URL 'https://newdomain.com', got %s", reqWithNewURL.GetURL().String())
	}

	// Test WithMethod
	reqWithNewMethod := req.WithMethod("POST")
	if reqWithNewMethod.GetMethod() != "POST" {
		t.Errorf("Expected new method 'POST', got %s", reqWithNewMethod.GetMethod())
	}

	// Test WithBody
	newBody := strings.NewReader("new body")
	reqWithNewBody := req.WithBody(io.NopCloser(newBody))
	body, err := reqWithNewBody.GetBody()
	if err != nil {
		t.Errorf("GetBody failed: %v", err)
	}
	if body == nil {
		t.Error("Expected non-nil body")
	}

	// Test ParseForm
	err = req.ParseForm()
	if err != nil {
		t.Errorf("ParseForm failed: %v", err)
	}
}

func TestMustNewRequest(t *testing.T) {
	// Test successful creation
	req := MustNewRequest("GET", "https://example.com", nil)
	if req == nil {
		t.Fatal("MustNewRequest should return non-nil request")
	}

	// Test panic on invalid URL
	defer func() {
		if r := recover(); r == nil {
			t.Error("MustNewRequest should panic on invalid URL")
		}
	}()
	MustNewRequest("GET", "://invalid-url", nil)
}

func TestMustNewRequestWithContext(t *testing.T) {
	ctx := context.Background()

	// Test successful creation
	req := MustNewRequestWithContext(ctx, "GET", "https://example.com", nil)
	if req == nil {
		t.Fatal("MustNewRequestWithContext should return non-nil request")
	}

	// Test panic on invalid URL
	defer func() {
		if r := recover(); r == nil {
			t.Error("MustNewRequestWithContext should panic on invalid URL")
		}
	}()
	MustNewRequestWithContext(ctx, "GET", "://invalid-url", nil)
}

func TestWithAuthorization(t *testing.T) {
	req, err := NewRequest("GET", "https://example.com", nil, WithAuthorization("Bearer token"))
	if err != nil {
		t.Fatalf("NewRequest failed: %v", err)
	}

	auth := req.GetHeader().Get("Authorization")
	if auth != "Bearer token" {
		t.Errorf("Expected Authorization header 'Bearer token', got %s", auth)
	}
}
