package httpx

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestNewAsyncClient(t *testing.T) {
	// Test with nil client
	client := NewAsyncClient(nil)
	if client == nil {
		t.Fatal("NewAsyncClient should not return nil")
	}
	if client.RestClient != DefaultClient {
		t.Error("NewAsyncClient with nil should use DefaultClient")
	}

	// Test with custom client
	customClient := &http.Client{Timeout: 5 * time.Second}
	asyncClient := NewAsyncClient(customClient)
	if asyncClient == nil {
		t.Fatal("NewAsyncClient should not return nil")
	}
	if asyncClient.RestClient == DefaultClient {
		t.Error("NewAsyncClient should use provided client")
	}
}

func TestAsyncClient_GetAsync(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("Expected GET request, got %s", r.Method)
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("Hello World"))
	}))
	defer server.Close()

	client := NewAsyncClient(nil)

	// Test async GET request
	resultChan := client.GetAsync(server.URL, nil)
	result := <-resultChan

	if result.Err != nil {
		t.Fatalf("GetAsync failed: %v", result.Err)
	}
	if result.Resp == nil {
		t.Fatal("Response should not be nil")
	}
	if !result.Resp.IsSuccess() {
		t.Errorf("Expected successful response, got status: %d", result.Resp.StatusCode)
	}

	body, err := result.Resp.String()
	if err != nil {
		t.Fatalf("Failed to read response body: %v", err)
	}
	if body != "Hello World" {
		t.Errorf("Expected 'Hello World', got: %s", body)
	}

	// Close the response
	_ = result.Resp.Close()
}

func TestAsyncClient_PostJsonAsync(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("Expected POST request, got %s", r.Method)
		}
		if r.Header.Get("Content-Type") != "application/json" {
			t.Errorf("Expected Content-Type: application/json, got: %s", r.Header.Get("Content-Type"))
		}
		w.WriteHeader(http.StatusCreated)
		_, _ = w.Write([]byte(`{"id": 1, "message": "created"}`))
	}))
	defer server.Close()

	client := NewAsyncClient(nil)
	body := strings.NewReader(`{"name": "test"}`)

	// Test async POST JSON request
	resultChan := client.PostJsonAsync(server.URL, body)
	result := <-resultChan

	if result.Err != nil {
		t.Fatalf("PostJsonAsync failed: %v", result.Err)
	}
	if result.Resp == nil {
		t.Fatal("Response should not be nil")
	}
	if result.Resp.StatusCode != http.StatusCreated {
		t.Errorf("Expected status 201, got: %d", result.Resp.StatusCode)
	}

	_ = result.Resp.Close()
}

func mustNewRequest(method, url string, body io.Reader) Request {
	req, err := NewRequest(method, url, body)
	if err != nil {
		panic(err)
	}
	return req
}

func TestAsyncClient_BatchAsync(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("response"))
	}))
	defer server.Close()

	client := NewAsyncClient(nil)

	// Create multiple requests
	requests := []Request{
		mustNewRequest("GET", server.URL, nil),
		mustNewRequest("GET", server.URL, nil),
		mustNewRequest("GET", server.URL, nil),
	}

	// Test batch async requests
	resultChan := client.BatchAsync(requests)
	results := <-resultChan

	if len(results) != len(requests) {
		t.Fatalf("Expected %d results, got %d", len(requests), len(results))
	}

	for i, result := range results {
		if result.Err != nil {
			t.Errorf("Request %d failed: %v", i, result.Err)
		}
		if result.Resp == nil {
			t.Errorf("Response %d should not be nil", i)
		}
		if result.Resp != nil {
			_ = result.Resp.Close()
		}
	}
}

func TestAsyncClient_WithContextAsync(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Simulate slow response
		time.Sleep(100 * time.Millisecond)
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("OK"))
	}))
	defer server.Close()

	client := NewAsyncClient(nil)

	// Test with timeout context
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()

	resultChan := client.WithContextAsync(ctx, "GET", server.URL, nil, nil)
	result := <-resultChan

	// Should timeout
	if result.Err == nil {
		t.Error("Expected timeout error, but got nil")
	}

	// Test with normal context
	ctx2 := context.Background()
	resultChan2 := client.WithContextAsync(ctx2, "GET", server.URL, nil, nil)
	result2 := <-resultChan2

	if result2.Err != nil {
		t.Fatalf("WithContextAsync failed: %v", result2.Err)
	}
	if result2.Resp == nil {
		t.Fatal("Response should not be nil")
	}
	defer func() {
		_ = result2.Resp.Close()
	}()
}

func TestWaitForAll(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("OK"))
	}))
	defer server.Close()

	client := NewAsyncClient(nil)

	// Create multiple async requests
	resultChans := []<-chan AsyncResult{
		client.GetAsync(server.URL, nil),
		client.GetAsync(server.URL, nil),
		client.GetAsync(server.URL, nil),
	}

	// Wait for all results
	results := WaitForAll(resultChans)

	if len(results) != len(resultChans) {
		t.Fatalf("Expected %d results, got %d", len(resultChans), len(results))
	}

	for i, result := range results {
		if result.Err != nil {
			t.Errorf("Result %d failed: %v", i, result.Err)
		}
		if result.Resp == nil {
			t.Errorf("Response %d should not be nil", i)
		}
		if result.Resp != nil {
			_ = result.Resp.Close()
		}
	}
}

func TestWaitForFirst(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Simulate different response times
		time.Sleep(100 * time.Millisecond)
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("OK"))
	}))
	defer server.Close()

	server2 := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Faster response
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("Fast"))
	}))
	defer server2.Close()

	client := NewAsyncClient(nil)

	// Create requests with different response times
	resultChans := []<-chan AsyncResult{
		client.GetAsync(server.URL, nil),  // Slow
		client.GetAsync(server2.URL, nil), // Fast
	}

	// Wait for first result
	result := WaitForFirst(resultChans)

	if result.Err != nil {
		t.Fatalf("WaitForFirst failed: %v", result.Err)
	}
	if result.Resp == nil {
		t.Fatal("Response should not be nil")
	}

	// Should be the fast response
	body, err := result.Resp.String()
	if err != nil {
		t.Fatalf("Failed to read response body: %v", err)
	}
	if body != "Fast" {
		t.Errorf("Expected 'Fast', got: %s", body)
	}

	defer func() {
		_ = result.Resp.Close()
	}()
}

func TestAsyncClient_OptionsAsync(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "OPTIONS" {
			t.Errorf("Expected OPTIONS request, got %s", r.Method)
		}

		// Set CORS headers for OPTIONS request
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("OPTIONS response"))
	}))
	defer server.Close()

	client := NewAsyncClient(nil)

	// Test basic OPTIONS request
	resultChan := client.OptionsAsync(server.URL)
	result := <-resultChan

	if result.Err != nil {
		t.Fatalf("OptionsAsync failed: %v", result.Err)
	}
	if result.Resp == nil {
		t.Fatal("Response should not be nil")
	}
	if !result.Resp.IsSuccess() {
		t.Errorf("Expected successful response, got status: %d", result.Resp.StatusCode)
	}

	// Check CORS headers
	allowOrigin := result.Resp.HeaderValue("Access-Control-Allow-Origin")
	if allowOrigin != "*" {
		t.Errorf("Expected Access-Control-Allow-Origin: *, got: %s", allowOrigin)
	}

	allowMethods := result.Resp.HeaderValue("Access-Control-Allow-Methods")
	if allowMethods != "GET, POST, PUT, DELETE, OPTIONS" {
		t.Errorf("Expected Access-Control-Allow-Methods: GET, POST, PUT, DELETE, OPTIONS, got: %s", allowMethods)
	}

	defer func() {
		_ = result.Resp.Close()
	}()
}

func TestAsyncClient_AllMethods(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("OK"))
	}))
	defer server.Close()

	client := NewAsyncClient(nil)
	body := strings.NewReader("test body")

	testCases := []struct {
		name   string
		method func() <-chan AsyncResult
	}{
		{"GetAsync", func() <-chan AsyncResult { return client.GetAsync(server.URL, nil) }},
		{"GetWithBodyAsync", func() <-chan AsyncResult { return client.GetWithBodyAsync(server.URL, body) }},
		{"GetJsonAsync", func() <-chan AsyncResult { return client.GetJsonAsync(server.URL, body) }},
		{"GetFormAsync", func() <-chan AsyncResult { return client.GetFormAsync(server.URL, body) }},
		{"PostAsync", func() <-chan AsyncResult { return client.PostAsync(server.URL, body) }},
		{"PostJsonAsync", func() <-chan AsyncResult { return client.PostJsonAsync(server.URL, body) }},
		{"PostFormAsync", func() <-chan AsyncResult { return client.PostFormAsync(server.URL, body) }},
		{"PutAsync", func() <-chan AsyncResult { return client.PutAsync(server.URL, body) }},
		{"PutJsonAsync", func() <-chan AsyncResult { return client.PutJsonAsync(server.URL, body) }},
		{"PutFormAsync", func() <-chan AsyncResult { return client.PutFormAsync(server.URL, body) }},
		{"PatchAsync", func() <-chan AsyncResult { return client.PatchAsync(server.URL, body) }},
		{"PatchJsonAsync", func() <-chan AsyncResult { return client.PatchJsonAsync(server.URL, body) }},
		{"PatchFormAsync", func() <-chan AsyncResult { return client.PatchFormAsync(server.URL, body) }},
		{"DeleteAsync", func() <-chan AsyncResult { return client.DeleteAsync(server.URL) }},
		{"DeleteJsonAsync", func() <-chan AsyncResult { return client.DeleteJsonAsync(server.URL, body) }},
		{"DeleteFormAsync", func() <-chan AsyncResult { return client.DeleteFormAsync(server.URL, body) }},
		{"OptionsAsync", func() <-chan AsyncResult { return client.OptionsAsync(server.URL) }},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			resultChan := tc.method()
			result := <-resultChan

			if result.Err != nil {
				t.Errorf("%s failed: %v", tc.name, result.Err)
			}
			if result.Resp == nil {
				t.Errorf("%s response should not be nil", tc.name)
			}
			if result.Resp != nil {
				_ = result.Resp.Close()
			}
		})
	}
}

func TestAsyncRequest_WithContentType(t *testing.T) {
	req := mustNewRequest("POST", "https://example.com", nil)
	reqPtr := req.WithHeader("Content-Type", "application/json")

	if reqPtr.GetHeader().Get("Content-Type") != "application/json" {
		t.Errorf("Expected ContentType 'application/json', got: %s", reqPtr.GetHeader().Get("Content-Type"))
	}
}

func TestAsyncRequest_WithHeader(t *testing.T) {
	req := mustNewRequest("POST", "https://example.com", nil)
	reqPtr := req.WithHeader("Authorization", "Bearer token")

	if reqPtr.GetHeader().Get("Authorization") != "Bearer token" {
		t.Errorf("Expected Authorization header 'Bearer token', got: %s", reqPtr.GetHeader().Get("Authorization"))
	}
}

func TestAsyncClient_executeAsyncRequest_IntelligentMethodSelection(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprintf("Method: %s, Content-Type: %s", r.Method, r.Header.Get("Content-Type"))))
	}))
	defer server.Close()

	client := NewAsyncClient(nil)

	testCases := []struct {
		name        string
		req         Request
		expectedCT  string
		description string
	}{
		{
			name:        "POST with JSON Content-Type",
			req:         mustNewRequest("POST", server.URL, strings.NewReader(`{"name": "test"}`)).WithHeader("Content-Type", "application/json"),
			expectedCT:  "application/json",
			description: "Should automatically use PostJson",
		},
		{
			name:        "POST with Form Content-Type",
			req:         mustNewRequest("POST", server.URL, strings.NewReader("name=test&value=123")).WithHeader("Content-Type", "application/x-www-form-urlencoded"),
			expectedCT:  "application/x-www-form-urlencoded",
			description: "Should automatically use PostForm",
		},
		{
			name:        "POST with Header Content-Type",
			req:         mustNewRequest("POST", server.URL, strings.NewReader(`{"name": "test"}`)).WithHeader("Content-Type", "application/json"),
			expectedCT:  "application/json",
			description: "Should detect Content-Type from headers",
		},
		{
			name:        "PUT with JSON Content-Type",
			req:         mustNewRequest("PUT", server.URL, strings.NewReader(`{"name": "test"}`)).WithHeader("Content-Type", "application/json"),
			expectedCT:  "application/json",
			description: "Should automatically use PutJson",
		},
		{
			name:        "PATCH with Form Content-Type",
			req:         mustNewRequest("PATCH", server.URL, strings.NewReader("name=test")).WithHeader("Content-Type", "application/x-www-form-urlencoded"),
			expectedCT:  "application/x-www-form-urlencoded",
			description: "Should automatically use PatchForm",
		},
		{
			name:        "DELETE with JSON Content-Type",
			req:         mustNewRequest("DELETE", server.URL, strings.NewReader(`{"id": 1}`)).WithHeader("Content-Type", "application/json"),
			expectedCT:  "application/json",
			description: "Should automatically use DeleteJson",
		},
		{
			name:        "GET with JSON body",
			req:         mustNewRequest("GET", server.URL, strings.NewReader(`{"query": "test"}`)).WithHeader("Content-Type", "application/json"),
			expectedCT:  "application/json",
			description: "Should automatically use GetJson",
		},
		{
			name:        "OPTIONS request",
			req:         mustNewRequest("OPTIONS", server.URL, nil),
			expectedCT:  "",
			description: "Should automatically use Options",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			resp, err := client.executeAsyncRequest(tc.req)
			if err != nil {
				t.Fatalf("%s failed: %v", tc.description, err)
			}
			defer func() {
				_ = resp.Close()
			}()

			if !resp.IsSuccess() {
				t.Errorf("%s: Expected successful response, got status: %d", tc.description, resp.StatusCode)
			}

			body, err := resp.String()
			if err != nil {
				t.Fatalf("Failed to read response body: %v", err)
			}

			if !strings.Contains(body, tc.expectedCT) {
				t.Errorf("%s: Expected response to contain Content-Type '%s', got: %s", tc.description, tc.expectedCT, body)
			}
		})
	}
}

func TestAsyncClient_deleteWithBody(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "DELETE" {
			t.Errorf("Expected DELETE request, got %s", r.Method)
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("DELETE with body response"))
	}))
	defer server.Close()

	client := NewAsyncClient(nil)
	body := strings.NewReader(`{"id": 1}`)

	// Test deleteWithBody directly
	resp, err := client.deleteWithBody(server.URL, body)
	if err != nil {
		t.Fatalf("deleteWithBody failed: %v", err)
	}
	defer func() {
		_ = resp.Close()
	}()

	if !resp.IsSuccess() {
		t.Errorf("Expected successful response, got status: %d", resp.StatusCode)
	}

	responseBody, _ := resp.String()
	if responseBody != "DELETE with body response" {
		t.Errorf("Expected 'DELETE with body response', got: %s", responseBody)
	}
}

// Benchmark tests
func BenchmarkAsyncClient_GetAsync(b *testing.B) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("OK"))
	}))
	defer server.Close()

	client := NewAsyncClient(nil)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		resultChan := client.GetAsync(server.URL, nil)
		result := <-resultChan
		if result.Resp != nil {
			_ = result.Resp.Close()
		}
	}
}

func BenchmarkAsyncClient_BatchAsync(b *testing.B) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("OK"))
	}))
	defer server.Close()

	client := NewAsyncClient(nil)
	requests := []Request{
		mustNewRequest("GET", server.URL, nil),
		mustNewRequest("GET", server.URL, nil),
		mustNewRequest("GET", server.URL, nil),
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		resultChan := client.BatchAsync(requests)
		results := <-resultChan
		for _, result := range results {
			if result.Resp != nil {
				_ = result.Resp.Close()
			}
		}
	}
}
