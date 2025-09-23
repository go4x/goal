package httpx

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"regexp"
	"strings"
	"time"
)

// normalizePort replaces dynamic port numbers with a placeholder for consistent output
func normalizePort(input string) string {
	// Replace dynamic port with placeholder
	re := regexp.MustCompile(`:\d+`)
	return re.ReplaceAllString(input, ":PORT")
}

func ExampleNewAsyncClient() {
	// Create an async client with default HTTP client
	client := NewAsyncClient(nil)

	// Create a test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("Hello Async!"))
	}))
	defer server.Close()

	// Make an async GET request
	resultChan := client.GetAsync(server.URL, nil)

	// Do other work while request is in progress
	fmt.Println("Request sent, doing other work...")

	// Wait for the result
	result := <-resultChan

	if result.Err != nil {
		fmt.Printf("Error: %v\n", result.Err)
		return
	}

	defer func() {
		_ = result.Resp.Close()
	}()
	body, _ := result.Resp.String()
	fmt.Printf("Response: %s\n", body)

	// Output:
	// Request sent, doing other work...
	// Response: Hello Async!
}

func ExampleAsyncClient_PostJsonAsync() {
	// Create an async client
	client := NewAsyncClient(nil)

	// Create a test server that expects JSON
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Content-Type") != "application/json" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusCreated)
		_, _ = w.Write([]byte(`{"id": 1, "status": "created"}`))
	}))
	defer server.Close()

	// Prepare JSON data
	jsonData := `{"name": "test", "value": 42}`
	body := strings.NewReader(jsonData)

	// Make an async POST JSON request
	resultChan := client.PostJsonAsync(server.URL, body)

	// Wait for the result
	result := <-resultChan

	if result.Err != nil {
		fmt.Printf("Error: %v\n", result.Err)
		return
	}

	defer func() {
		_ = result.Resp.Close()
	}()
	fmt.Printf("Status: %d\n", result.Resp.StatusCode)

	// Output:
	// Status: 201
}

func ExampleAsyncClient_BatchAsync() {
	// Create an async client
	client := NewAsyncClient(nil)

	// Create a test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("Batch response"))
	}))
	defer server.Close()

	// Create multiple requests with intelligent method selection
	getReq := MustNewRequest("GET", server.URL, nil)
	postReq := MustNewRequest("POST", server.URL, strings.NewReader(`{"name": "test"}`), WithContentType("application/json"))
	putReq := MustNewRequest("PUT", server.URL, strings.NewReader("name=test&value=123"), WithContentType("application/x-www-form-urlencoded"))

	requests := []Request{getReq, postReq, putReq}

	// Send all requests concurrently
	resultChan := client.BatchAsync(requests)

	// Wait for all results
	results := <-resultChan

	fmt.Printf("Received %d responses\n", len(results))

	// Process results
	for i, result := range results {
		if result.Err != nil {
			fmt.Printf("Request %d failed: %v\n", i, result.Err)
			continue
		}
		defer func() {
			_ = result.Resp.Close()
		}()
		fmt.Printf("Request %d: Status %d\n", i, result.Resp.StatusCode)
	}

	// Output:
	// Received 3 responses
	// Request 0: Status 200
	// Request 1: Status 200
	// Request 2: Status 200
}

func ExampleAsyncClient_WithContextAsync() {
	// Create an async client
	client := NewAsyncClient(nil)

	// Create a slow test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(200 * time.Millisecond) // Simulate slow response
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("Slow response"))
	}))
	defer server.Close()

	// Create a context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	// Make an async request with context
	resultChan := client.WithContextAsync(ctx, "GET", server.URL, nil, nil)

	// Wait for the result
	result := <-resultChan

	// show normalized error message (replace dynamic port with placeholder)
	if result.Err != nil {
		normalizedErr := normalizePort(result.Err.Error())
		fmt.Printf("Request failed (expected due to timeout): %s\n", normalizedErr)
	} else {
		defer func() {
			_ = result.Resp.Close()
		}()
		body, _ := result.Resp.String()
		fmt.Printf("Response: %s\n", body)
	}

	// Output:
	// Request failed (expected due to timeout): Get "http://127.0.0.1:PORT": context deadline exceeded
}

func ExampleWaitForAll() {
	// Create an async client
	client := NewAsyncClient(nil)

	// Create test servers
	server1 := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(50 * time.Millisecond)
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("Server 1"))
	}))
	defer server1.Close()

	server2 := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(100 * time.Millisecond)
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("Server 2"))
	}))
	defer server2.Close()

	// Start multiple async requests
	resultChans := []<-chan AsyncResult{
		client.GetAsync(server1.URL, nil),
		client.GetAsync(server2.URL, nil),
	}

	// Wait for all results
	results := WaitForAll(resultChans)

	fmt.Printf("All %d requests completed\n", len(results))

	// Process results
	for i, result := range results {
		if result.Err != nil {
			fmt.Printf("Request %d failed: %v\n", i, result.Err)
			continue
		}
		defer func() {
			_ = result.Resp.Close()
		}()
		body, _ := result.Resp.String()
		fmt.Printf("Request %d: %s\n", i, body)
	}

	// Output:
	// All 2 requests completed
	// Request 0: Server 1
	// Request 1: Server 2
}

func ExampleWaitForFirst() {
	// Create an async client
	client := NewAsyncClient(nil)

	// Create test servers with different response times
	server1 := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(100 * time.Millisecond) // Slow
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("Slow response"))
	}))
	defer server1.Close()

	server2 := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Fast response
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("Fast response"))
	}))
	defer server2.Close()

	// Start multiple async requests
	resultChans := []<-chan AsyncResult{
		client.GetAsync(server1.URL, nil), // Slow
		client.GetAsync(server2.URL, nil), // Fast
	}

	// Wait for the first result
	result := WaitForFirst(resultChans)

	if result.Err != nil {
		fmt.Printf("Request failed: %v\n", result.Err)
		return
	}

	defer func() {
		_ = result.Resp.Close()
	}()
	body, _ := result.Resp.String()
	fmt.Printf("First response: %s\n", body)

	// Output:
	// First response: Fast response
}

func ExampleAsyncClient_errorHandling() {
	// Create an async client
	client := NewAsyncClient(nil)

	// Try to make a request to an invalid URL
	resultChan := client.GetAsync("invalid-url", nil)

	// Wait for the result
	result := <-resultChan

	if result.Err != nil {
		fmt.Printf("Expected error: %v\n", result.Err)
	} else {
		defer func() {
			_ = result.Resp.Close()
		}()
		fmt.Printf("Unexpected success: %d\n", result.Resp.StatusCode)
	}

	// Output:
	// Expected error: Get "invalid-url": unsupported protocol scheme ""
}

func ExampleAsyncClient_mixedMethods() {
	// Create an async client
	client := NewAsyncClient(nil)

	// Create a test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(fmt.Sprintf("Method: %s", r.Method)))
	}))
	defer server.Close()

	// Start requests with different HTTP methods
	resultChans := []<-chan AsyncResult{
		client.GetAsync(server.URL, nil),
		client.PostAsync(server.URL, strings.NewReader("data")),
		client.PutAsync(server.URL, strings.NewReader("data")),
		client.DeleteAsync(server.URL),
	}

	// Wait for all results
	results := WaitForAll(resultChans)

	// Process results
	for i, result := range results {
		if result.Err != nil {
			fmt.Printf("Request %d failed: %v\n", i, result.Err)
			continue
		}
		defer func() {
			_ = result.Resp.Close()
		}()
		body, _ := result.Resp.String()
		fmt.Printf("Request %d: %s\n", i, body)
	}

	// Output:
	// Request 0: Method: GET
	// Request 1: Method: POST
	// Request 2: Method: PUT
	// Request 3: Method: DELETE
}

func ExampleAsyncClient_intelligentMethodSelection() {
	// Create an async client
	client := NewAsyncClient(nil)

	// Create a test server that shows the received Content-Type
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(fmt.Sprintf("Method: %s, Content-Type: %s", r.Method, r.Header.Get("Content-Type"))))
	}))
	defer server.Close()

	// Create requests with different Content-Types - the client will automatically
	// choose the appropriate method (PostJson, PostForm, etc.)
	requests := []Request{
		// POST with JSON - will automatically use PostJson
		MustNewRequest("POST", server.URL, strings.NewReader(`{"name": "test"}`), WithContentType("application/json")),

		// POST with Form - will automatically use PostForm
		MustNewRequest("POST", server.URL, strings.NewReader("name=test&value=123"), WithContentType("application/x-www-form-urlencoded")),

		// PUT with JSON - will automatically use PutJson
		MustNewRequest("PUT", server.URL, strings.NewReader(`{"id": 1}`), WithContentType("application/json")),

		// DELETE with JSON - will automatically use DeleteJson
		MustNewRequest("DELETE", server.URL, strings.NewReader(`{"id": 1}`), WithContentType("application/json")),

		// OPTIONS request - will automatically use Options
		MustNewRequest("OPTIONS", server.URL, nil),
	}

	// Send all requests concurrently
	resultChan := client.BatchAsync(requests)
	results := <-resultChan

	// Process results
	for i, result := range results {
		if result.Err != nil {
			fmt.Printf("Request %d failed: %v\n", i, result.Err)
			continue
		}
		defer func() {
			_ = result.Resp.Close()
		}()
		body, _ := result.Resp.String()
		fmt.Printf("Request %d: %s\n", i, body)
	}

	// Output:
	// Request 0: Method: POST, Content-Type: application/json
	// Request 1: Method: POST, Content-Type: application/x-www-form-urlencoded
	// Request 2: Method: PUT, Content-Type: application/json
	// Request 3: Method: DELETE, Content-Type: application/json
	// Request 4: Method: OPTIONS, Content-Type:
}

func ExampleAsyncClient_OptionsAsync() {
	// Create an async client
	client := NewAsyncClient(nil)

	// Create a test server that handles OPTIONS requests
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "OPTIONS" {
			// Set CORS headers
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte("CORS preflight response"))
			return
		}
		w.WriteHeader(http.StatusMethodNotAllowed)
	}))
	defer server.Close()

	// Make an async OPTIONS request
	resultChan := client.OptionsAsync(server.URL)
	result := <-resultChan

	if result.Err != nil {
		fmt.Printf("Error: %v\n", result.Err)
		return
	}

	defer func() {
		_ = result.Resp.Close()
	}()

	// Check CORS headers
	allowOrigin := result.Resp.HeaderValue("Access-Control-Allow-Origin")
	allowMethods := result.Resp.HeaderValue("Access-Control-Allow-Methods")

	fmt.Printf("Status: %d\n", result.Resp.StatusCode)
	fmt.Printf("Allow-Origin: %s\n", allowOrigin)
	fmt.Printf("Allow-Methods: %s\n", allowMethods)

	// Output:
	// Status: 200
	// Allow-Origin: *
	// Allow-Methods: GET, POST, PUT, DELETE
}
