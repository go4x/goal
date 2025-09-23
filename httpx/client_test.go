package httpx

import (
	"context"
	"net/http"
	"net/url"
	"strings"
	"testing"
	"time"
)

func TestNewClient(t *testing.T) {
	// Test with nil client
	client := NewRestClient(nil)
	if client == nil {
		t.Fatal("NewClient with nil should return a client")
	}
	if client.Client == nil {
		t.Fatal("Client should not be nil")
	}
	if client.Timeout != 60*time.Second {
		t.Fatal("Default timeout should be 60 seconds")
	}

	// Test with custom client
	customClient := &http.Client{Timeout: 10 * time.Second}
	client2 := NewRestClient(customClient)
	if client2.Client != customClient {
		t.Fatal("Should use the provided client")
	}
}

func TestRestClient_Do(t *testing.T) {
	client := NewRestClient(nil)

	// Create a simple request
	req, err := http.NewRequest("GET", "https://httpbin.org/get", nil)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		t.Fatalf("Expected status 200, got %d", resp.StatusCode)
	}
}

func TestRestClient_buildUrl(t *testing.T) {
	client := NewRestClient(nil)

	tests := []struct {
		name     string
		url      string
		params   url.Values
		expected string
		hasError bool
	}{
		{
			name:     "simple URL without params",
			url:      "https://example.com/api",
			params:   nil,
			expected: "https://example.com/api",
			hasError: false,
		},
		{
			name:     "URL with single param",
			url:      "https://example.com/api",
			params:   url.Values{"page": []string{"1"}},
			expected: "https://example.com/api?page=1",
			hasError: false,
		},
		{
			name:     "URL with multiple params",
			url:      "https://example.com/api",
			params:   url.Values{"page": []string{"1"}, "limit": []string{"10"}},
			expected: "https://example.com/api?limit=10&page=1", // order may vary
			hasError: false,
		},
		{
			name:     "URL with multiple values for same param",
			url:      "https://example.com/api",
			params:   url.Values{"tags": []string{"tech", "science"}},
			expected: "https://example.com/api?tags=tech&tags=science",
			hasError: false,
		},
		{
			name:     "URL with existing query params",
			url:      "https://example.com/api?existing=value",
			params:   url.Values{"page": []string{"1"}},
			expected: "https://example.com/api?existing=value&page=1",
			hasError: false,
		},
		{
			name:     "invalid URL",
			url:      "://invalid-url",
			params:   nil,
			expected: "://invalid-url",
			hasError: true, // buildUrl should return an error for invalid URLs
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := client.buildUrl(tt.url, tt.params)

			if tt.hasError && err == nil {
				t.Errorf("Expected error but got none")
			}
			if !tt.hasError && err != nil {
				t.Errorf("Unexpected error: %v", err)
			}

			// For URLs with multiple params, just check that all params are present
			if len(tt.params) > 1 {
				parsedURL, _ := url.Parse(result)
				query := parsedURL.Query()
				for key, values := range tt.params {
					if !query.Has(key) {
						t.Errorf("Missing parameter %s", key)
					}
					for _, value := range values {
						if !contains(query[key], value) {
							t.Errorf("Missing value %s for parameter %s", value, key)
						}
					}
				}
			} else {
				if result != tt.expected && tt.name != "URL with multiple params" {
					t.Errorf("Expected %s, got %s", tt.expected, result)
				}
			}
		})
	}
}

// Helper function to check if slice contains a value
func contains(slice []string, value string) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

func TestRestClient_Get(t *testing.T) {
	client := NewRestClient(nil)

	t.Run("GET with params", func(t *testing.T) {
		params := url.Values{}
		params.Set("test", "value")
		params.Set("page", "1")

		resp, err := client.Get("https://httpbin.org/get", params)
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}
		defer func() { _ = resp.Close() }()

		if !resp.IsSuccess() {
			t.Errorf("Expected success response, got status %d", resp.StatusCode)
		}
	})

	t.Run("GET without params", func(t *testing.T) {
		resp, err := client.Get("https://httpbin.org/get", nil)
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}
		defer func() { _ = resp.Close() }()

		if !resp.IsSuccess() {
			t.Errorf("Expected success response, got status %d", resp.StatusCode)
		}
	})
}

func TestRestClient_GetWithBody(t *testing.T) {
	client := NewRestClient(nil)

	t.Run("GET with JSON body", func(t *testing.T) {
		jsonBody := strings.NewReader(`{"test": "value"}`)

		resp, err := client.GetWithBody("https://httpbin.org/get", jsonBody)
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}
		defer func() { _ = resp.Close() }()

		if !resp.IsSuccess() {
			t.Errorf("Expected success response, got status %d", resp.StatusCode)
		}
	})

	t.Run("GET with nil body", func(t *testing.T) {
		resp, err := client.GetWithBody("https://httpbin.org/get", nil)
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}
		defer func() { _ = resp.Close() }()

		if !resp.IsSuccess() {
			t.Errorf("Expected success response, got status %d", resp.StatusCode)
		}
	})
}

func TestRestClient_GetJson(t *testing.T) {
	client := NewRestClient(nil)

	jsonBody := strings.NewReader(`{"query": "test"}`)

	resp, err := client.GetJson("https://httpbin.org/get", jsonBody)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	defer func() { _ = resp.Close() }()

	if !resp.IsSuccess() {
		t.Errorf("Expected success response, got status %d", resp.StatusCode)
	}

	// Note: We can't easily check request headers from response in this test
	// The Content-Type is set on the request, not the response
}

func TestRestClient_GetForm(t *testing.T) {
	client := NewRestClient(nil)

	formData := url.Values{}
	formData.Set("field1", "value1")
	formData.Set("field2", "value2")
	formBody := strings.NewReader(formData.Encode())

	resp, err := client.GetForm("https://httpbin.org/get", formBody)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	defer func() { _ = resp.Close() }()

	if !resp.IsSuccess() {
		t.Errorf("Expected success response, got status %d", resp.StatusCode)
	}

	// Note: Content-Type is set on the request, not the response
}

func TestRestClient_Post(t *testing.T) {
	client := NewRestClient(nil)

	jsonBody := strings.NewReader(`{"test": "data"}`)

	resp, err := client.Post("https://httpbin.org/post", jsonBody)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	defer func() { _ = resp.Close() }()

	if !resp.IsSuccess() {
		t.Errorf("Expected success response, got status %d", resp.StatusCode)
	}
}

func TestRestClient_PostJson(t *testing.T) {
	client := NewRestClient(nil)

	jsonBody := strings.NewReader(`{"name": "test", "value": 123}`)

	resp, err := client.PostJson("https://httpbin.org/post", jsonBody)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	defer func() { _ = resp.Close() }()

	if !resp.IsSuccess() {
		t.Errorf("Expected success response, got status %d", resp.StatusCode)
	}

	// Note: Content-Type is set on the request, not the response
}

func TestRestClient_PostForm(t *testing.T) {
	client := NewRestClient(nil)

	formData := url.Values{}
	formData.Set("name", "test")
	formData.Set("value", "123")
	formBody := strings.NewReader(formData.Encode())

	resp, err := client.PostForm("https://httpbin.org/post", formBody)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	defer func() { _ = resp.Close() }()

	if !resp.IsSuccess() {
		t.Errorf("Expected success response, got status %d", resp.StatusCode)
	}

	// Note: Content-Type is set on the request, not the response
}

func TestRestClient_Put(t *testing.T) {
	client := NewRestClient(nil)

	jsonBody := strings.NewReader(`{"id": 1, "name": "updated"}`)

	resp, err := client.Put("https://httpbin.org/put", jsonBody)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	defer func() { _ = resp.Close() }()

	if !resp.IsSuccess() {
		t.Errorf("Expected success response, got status %d", resp.StatusCode)
	}
}

func TestRestClient_PutJson(t *testing.T) {
	client := NewRestClient(nil)

	jsonBody := strings.NewReader(`{"id": 1, "name": "updated"}`)

	resp, err := client.PutJson("https://httpbin.org/put", jsonBody)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	defer func() { _ = resp.Close() }()

	if !resp.IsSuccess() {
		t.Errorf("Expected success response, got status %d", resp.StatusCode)
	}

	// Note: Content-Type is set on the request, not the response
}

func TestRestClient_PutForm(t *testing.T) {
	client := NewRestClient(nil)

	formData := url.Values{}
	formData.Set("id", "1")
	formData.Set("name", "updated")
	formBody := strings.NewReader(formData.Encode())

	resp, err := client.PutForm("https://httpbin.org/put", formBody)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	defer func() { _ = resp.Close() }()

	if !resp.IsSuccess() {
		t.Errorf("Expected success response, got status %d", resp.StatusCode)
	}

	// Note: Content-Type is set on the request, not the response
}

func TestRestClient_Patch(t *testing.T) {
	client := NewRestClient(nil)

	jsonBody := strings.NewReader(`{"name": "patched"}`)

	resp, err := client.Patch("https://httpbin.org/patch", jsonBody)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	defer func() { _ = resp.Close() }()

	if !resp.IsSuccess() {
		t.Errorf("Expected success response, got status %d", resp.StatusCode)
	}
}

func TestRestClient_PatchJson(t *testing.T) {
	client := NewRestClient(nil)

	jsonBody := strings.NewReader(`{"name": "patched"}`)

	resp, err := client.PatchJson("https://httpbin.org/patch", jsonBody)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	defer func() { _ = resp.Close() }()

	if !resp.IsSuccess() {
		t.Errorf("Expected success response, got status %d", resp.StatusCode)
	}

	// Note: Content-Type is set on the request, not the response
}

func TestRestClient_PatchForm(t *testing.T) {
	client := NewRestClient(nil)

	formData := url.Values{}
	formData.Set("name", "patched")
	formBody := strings.NewReader(formData.Encode())

	resp, err := client.PatchForm("https://httpbin.org/patch", formBody)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	defer func() { _ = resp.Close() }()

	if !resp.IsSuccess() {
		t.Errorf("Expected success response, got status %d", resp.StatusCode)
	}

	// Note: Content-Type is set on the request, not the response
}

func TestRestClient_Delete(t *testing.T) {
	client := NewRestClient(nil)

	resp, err := client.Delete("https://httpbin.org/delete")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	defer func() { _ = resp.Close() }()

	if !resp.IsSuccess() {
		t.Errorf("Expected success response, got status %d", resp.StatusCode)
	}
}

func TestRestClient_DeleteJson(t *testing.T) {
	client := NewRestClient(nil)

	jsonBody := strings.NewReader(`{"confirmCode": "ABC123"}`)

	resp, err := client.DeleteJson("https://httpbin.org/delete", jsonBody)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	defer func() { _ = resp.Close() }()

	if !resp.IsSuccess() {
		t.Errorf("Expected success response, got status %d", resp.StatusCode)
	}

	// Note: Content-Type is set on the request, not the response
}

func TestRestClient_DeleteForm(t *testing.T) {
	client := NewRestClient(nil)

	formData := url.Values{}
	formData.Set("confirmCode", "ABC123")
	formBody := strings.NewReader(formData.Encode())

	resp, err := client.DeleteForm("https://httpbin.org/delete", formBody)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	defer func() { _ = resp.Close() }()

	if !resp.IsSuccess() {
		t.Errorf("Expected success response, got status %d", resp.StatusCode)
	}

	// Note: Content-Type is set on the request, not the response
}

func TestRestClient_Options(t *testing.T) {
	client := NewRestClient(nil)

	// Test basic OPTIONS request
	resp, err := client.Options("https://httpbin.org/anything")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	defer func() { _ = resp.Close() }()

	if !resp.IsSuccess() {
		t.Errorf("Expected success response, got status %d", resp.StatusCode)
	}

	// OPTIONS request should return 200 OK for httpbin.org
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200, got %d", resp.StatusCode)
	}

	// Test OPTIONS request with custom headers
	resp2, err := client.Options("https://httpbin.org/anything", WithHeader("Custom-Header", "test-value"))
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	defer func() { _ = resp2.Close() }()

	if !resp2.IsSuccess() {
		t.Errorf("Expected success response, got status %d", resp2.StatusCode)
	}
}

func TestRestClient_RequestOptions(t *testing.T) {
	client := NewRestClient(nil)

	t.Run("WithContentType option", func(t *testing.T) {
		jsonBody := strings.NewReader(`{"test": "data"}`)

		resp, err := client.Post("https://httpbin.org/post", jsonBody, WithContentType("application/custom"))
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}
		defer func() { _ = resp.Close() }()

		if !resp.IsSuccess() {
			t.Errorf("Expected success response, got status %d", resp.StatusCode)
		}

		// Note: Content-Type is set on the request, not the response
	})

	t.Run("WithHeader option", func(t *testing.T) {
		jsonBody := strings.NewReader(`{"test": "data"}`)

		resp, err := client.Post("https://httpbin.org/post", jsonBody,
			WithHeader("X-Custom-Header", "custom-value"),
			WithHeader("X-Another-Header", "another-value"))
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}
		defer func() { _ = resp.Close() }()

		if !resp.IsSuccess() {
			t.Errorf("Expected success response, got status %d", resp.StatusCode)
		}
	})

	t.Run("WithContext option", func(t *testing.T) {
		jsonBody := strings.NewReader(`{"test": "data"}`)

		// Use a proper context instead of nil
		ctx := context.Background()
		resp, err := client.Post("https://httpbin.org/post", jsonBody, WithContext(ctx))
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}
		defer func() { _ = resp.Close() }()

		if !resp.IsSuccess() {
			t.Errorf("Expected success response, got status %d", resp.StatusCode)
		}
	})

	t.Run("Multiple options", func(t *testing.T) {
		jsonBody := strings.NewReader(`{"test": "data"}`)

		resp, err := client.Post("https://httpbin.org/post", jsonBody,
			WithContentType("application/json"),
			WithHeader("X-Test", "value"),
			WithHeader("Authorization", "Bearer token"))
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}
		defer func() { _ = resp.Close() }()

		if !resp.IsSuccess() {
			t.Errorf("Expected success response, got status %d", resp.StatusCode)
		}

		// Note: Content-Type is set on the request, not the response
	})
}

func TestRestClient_EdgeCases(t *testing.T) {
	client := NewRestClient(nil)

	t.Run("nil body handling", func(t *testing.T) {
		resp, err := client.Post("https://httpbin.org/post", nil)
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}
		defer func() { _ = resp.Close() }()

		if !resp.IsSuccess() {
			t.Errorf("Expected success response, got status %d", resp.StatusCode)
		}
	})

	t.Run("empty body handling", func(t *testing.T) {
		emptyBody := strings.NewReader("")
		resp, err := client.Post("https://httpbin.org/post", emptyBody)
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}
		defer func() { _ = resp.Close() }()

		if !resp.IsSuccess() {
			t.Errorf("Expected success response, got status %d", resp.StatusCode)
		}
	})

	t.Run("large body handling", func(t *testing.T) {
		// Create a large JSON body
		largeData := strings.Repeat(`{"field": "value", "number": 123},`, 1000)
		largeBody := strings.NewReader(`[` + largeData[:len(largeData)-1] + `]`)

		resp, err := client.Post("https://httpbin.org/post", largeBody)
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}
		defer func() { _ = resp.Close() }()

		if !resp.IsSuccess() {
			t.Errorf("Expected success response, got status %d", resp.StatusCode)
		}
	})
}

// Benchmark tests
func BenchmarkRestClient_Get(b *testing.B) {
	client := NewRestClient(nil)
	params := url.Values{}
	params.Set("test", "benchmark")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		resp, err := client.Get("https://httpbin.org/get", params)
		if err != nil {
			b.Fatal(err)
		}
		_ = resp.Close()
	}
}

func BenchmarkRestClient_PostJson(b *testing.B) {
	client := NewRestClient(nil)
	jsonBody := strings.NewReader(`{"test": "benchmark"}`)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		resp, err := client.PostJson("https://httpbin.org/post", jsonBody)
		if err != nil {
			b.Fatal(err)
		}
		_ = resp.Close()
	}
}

func BenchmarkRestClient_buildUrl(b *testing.B) {
	client := NewRestClient(nil)
	params := url.Values{}
	params.Set("page", "1")
	params.Set("limit", "10")
	params.Set("search", "test")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := client.buildUrl("https://example.com/api", params)
		if err != nil {
			b.Fatal(err)
		}
	}
}
