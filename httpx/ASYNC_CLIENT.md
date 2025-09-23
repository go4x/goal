# AsyncClient - Asynchronous HTTP Client Guide

The `AsyncClient` provides non-blocking HTTP request capabilities using Go's channels and goroutines. It's designed for scenarios where you need to perform multiple HTTP requests concurrently or handle requests without blocking the main execution flow.

## Table of Contents

- [Overview](#overview)
- [Basic Usage](#basic-usage)
- [Asynchronous Methods](#asynchronous-methods)
- [Batch Operations](#batch-operations)
- [Context Support](#context-support)
- [Result Aggregation](#result-aggregation)
- [Error Handling](#error-handling)
- [Best Practices](#best-practices)
- [Performance Considerations](#performance-considerations)
- [Examples](#examples)

## Overview

The `AsyncClient` wraps the `RestClient` and provides asynchronous versions of all HTTP methods. Instead of blocking until a request completes, these methods return channels that will receive the result when the request finishes.

### Key Features

- **Non-blocking Operations**: All HTTP methods return channels for asynchronous execution
- **Intelligent Method Selection**: Automatically chooses the correct HTTP method based on content type
- **Batch Processing**: Execute multiple requests concurrently with result aggregation
- **Context Support**: Full integration with Go's context for timeouts and cancellation
- **Resource Management**: Automatic cleanup of response resources

## Basic Usage

### Creating an AsyncClient

```go
package main

import (
    "fmt"
    "github.com/go4x/goal/httpx"
)

func main() {
    // Create an async client with default HTTP client
    client := httpx.NewAsyncClient(nil)
    
    // Or create with a custom HTTP client
    customClient := &http.Client{Timeout: 30 * time.Second}
    asyncClient := httpx.NewAsyncClient(customClient)
}
```

### Simple Asynchronous Request

```go
// Send an asynchronous GET request
resultChan := client.GetAsync("https://httpbin.org/get", nil)

// Do other work while the request is in progress
fmt.Println("Request sent, doing other tasks...")

// Wait for the result
result := <-resultChan
if result.Err != nil {
    fmt.Printf("Request failed: %v\n", result.Err)
    return
}
defer result.Resp.Close()

fmt.Printf("Response status: %d\n", result.Resp.StatusCode)
```

## Asynchronous Methods

All HTTP methods have asynchronous counterparts that return `<-chan AsyncResult`:

### GET Methods

```go
// Simple GET request
result := <-client.GetAsync("https://api.example.com/users", nil)

// GET with query parameters
params := url.Values{"page": {"1"}, "limit": {"10"}}
result := <-client.GetAsync("https://api.example.com/users", params)

// GET with body (uncommon but supported)
body := strings.NewReader(`{"query": "search term"}`)
result := <-client.GetWithBodyAsync("https://api.example.com/search", body)

// GET with JSON body
result := <-client.GetJsonAsync("https://api.example.com/search", body)

// GET with form data
formData := strings.NewReader("name=value&email=user@example.com")
result := <-client.GetFormAsync("https://api.example.com/search", formData)
```

### POST Methods

```go
// POST with JSON
jsonBody := strings.NewReader(`{"name": "John", "email": "john@example.com"}`)
result := <-client.PostJsonAsync("https://api.example.com/users", jsonBody)

// POST with form data
formData := strings.NewReader("name=John&email=john@example.com")
result := <-client.PostFormAsync("https://api.example.com/users", formData)

// POST with custom body
result := <-client.PostAsync("https://api.example.com/upload", fileReader)
```

### PUT, PATCH, DELETE Methods

```go
// PUT operations
result := <-client.PutAsync("https://api.example.com/users/1", body)
result := <-client.PutJsonAsync("https://api.example.com/users/1", jsonBody)
result := <-client.PutFormAsync("https://api.example.com/users/1", formData)

// PATCH operations
result := <-client.PatchAsync("https://api.example.com/users/1", body)
result := <-client.PatchJsonAsync("https://api.example.com/users/1", jsonBody)
result := <-client.PatchFormAsync("https://api.example.com/users/1", formData)

// DELETE operations
result := <-client.DeleteAsync("https://api.example.com/users/1")
result := <-client.DeleteJsonAsync("https://api.example.com/users/1", jsonBody)
result := <-client.DeleteFormAsync("https://api.example.com/users/1", formData)

// OPTIONS request
result := <-client.OptionsAsync("https://api.example.com/users")
```

## Batch Operations

### BatchAsync - Concurrent Request Execution

The `BatchAsync` method allows you to execute multiple requests concurrently and collect all results:

```go
// Create multiple requests
requests := []httpx.Request{
    httpx.MustNewRequest("GET", "https://api1.example.com/data", nil),
    httpx.MustNewRequest("POST", "https://api2.example.com/process", jsonBody,
        httpx.WithContentType("application/json")),
    httpx.MustNewRequest("PUT", "https://api3.example.com/update", formData,
        httpx.WithContentType("application/x-www-form-urlencoded")),
}

// Execute all requests concurrently
asyncClient := httpx.NewAsyncClient(nil)
resultsChan := asyncClient.BatchAsync(requests)

// Wait for all results
results := <-resultsChan

// Process results
for i, result := range results {
    if result.Err != nil {
        fmt.Printf("Request %d failed: %v\n", i, result.Err)
        continue
    }
    defer result.Resp.Close()
    
    fmt.Printf("Request %d succeeded with status: %d\n", i, result.Resp.StatusCode)
}
```

### Intelligent Method Selection

The `BatchAsync` method uses intelligent method selection based on the HTTP method and Content-Type header:

```go
requests := []httpx.Request{
    // This will automatically use GetAsync
    httpx.MustNewRequest("GET", "https://api.example.com/users", nil),
    
    // This will automatically use PostJsonAsync
    httpx.MustNewRequest("POST", "https://api.example.com/users", jsonBody,
        httpx.WithContentType("application/json")),
    
    // This will automatically use PutFormAsync
    httpx.MustNewRequest("PUT", "https://api.example.com/users/1", formData,
        httpx.WithContentType("application/x-www-form-urlencoded")),
}

results := <-asyncClient.BatchAsync(requests)
```

## Context Support

### WithContextAsync

Use `WithContextAsync` for requests that need timeout or cancellation control:

```go
// Create a context with timeout
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()

// Send request with context
resultChan := asyncClient.WithContextAsync(ctx, "GET", "https://httpbin.org/delay/3", nil, nil)

// Wait for result
result := <-resultChan
if result.Err != nil {
    if errors.Is(result.Err, context.DeadlineExceeded) {
        fmt.Println("Request timed out")
    } else {
        fmt.Printf("Request failed: %v\n", result.Err)
    }
    return
}
defer result.Resp.Close()

fmt.Printf("Request completed: %d\n", result.Resp.StatusCode)
```

### Context with Batch Operations

```go
ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
defer cancel()

// Create requests with context
requests := []httpx.Request{
    httpx.MustNewRequestWithContext(ctx, "GET", "https://api1.example.com", nil),
    httpx.MustNewRequestWithContext(ctx, "GET", "https://api2.example.com", nil),
    httpx.MustNewRequestWithContext(ctx, "GET", "https://api3.example.com", nil),
}

results := <-asyncClient.BatchAsync(requests)
// All requests will be cancelled if context times out
```

## Result Aggregation

### WaitForAll

Collect results from multiple independent async operations:

```go
// Start multiple independent requests
var resultChans []<-chan httpx.AsyncResult

resultChans = append(resultChans, client.GetAsync("https://api1.example.com", nil))
resultChans = append(resultChans, client.PostJsonAsync("https://api2.example.com", jsonBody))
resultChans = append(resultChans, client.PutAsync("https://api3.example.com", body))

// Wait for all to complete
results := httpx.WaitForAll(resultChans)

for i, result := range results {
    if result.Err != nil {
        fmt.Printf("Request %d failed: %v\n", i, result.Err)
        continue
    }
    defer result.Resp.Close()
    fmt.Printf("Request %d succeeded\n", i)
}
```

### WaitForFirst

Get the first completed result from multiple operations:

```go
// Start multiple requests with different delays
var resultChans []<-chan httpx.AsyncResult

resultChans = append(resultChans, client.GetAsync("https://slow-api.example.com", nil))
resultChans = append(resultChans, client.GetAsync("https://fast-api.example.com", nil))
resultChans = append(resultChans, client.GetAsync("https://medium-api.example.com", nil))

// Get the first result (likely from fast-api)
firstResult := httpx.WaitForFirst(resultChans)

if firstResult.Err != nil {
    fmt.Printf("First request failed: %v\n", firstResult.Err)
} else {
    defer firstResult.Resp.Close()
    fmt.Printf("First request completed: %d\n", firstResult.Resp.StatusCode)
}

// Note: Other requests continue in the background
```

## Error Handling

### Basic Error Handling

```go
result := <-client.GetAsync("https://api.example.com/data", nil)

if result.Err != nil {
    // Handle network or request creation errors
    fmt.Printf("Request failed: %v\n", result.Err)
    return
}
defer result.Resp.Close()

if !result.Resp.IsSuccess() {
    // Handle HTTP error responses
    body, _ := result.Resp.String()
    fmt.Printf("HTTP error %d: %s\n", result.Resp.StatusCode, body)
    return
}

// Success
fmt.Printf("Success: %d\n", result.Resp.StatusCode)
```

### Error Handling with Batch Operations

```go
requests := []httpx.Request{
    httpx.MustNewRequest("GET", "https://api1.example.com", nil),
    httpx.MustNewRequest("GET", "https://api2.example.com", nil),
    httpx.MustNewRequest("GET", "https://api3.example.com", nil),
}

results := <-asyncClient.BatchAsync(requests)

successCount := 0
errorCount := 0

for i, result := range results {
    if result.Err != nil {
        errorCount++
        fmt.Printf("Request %d failed: %v\n", i, result.Err)
        continue
    }
    defer result.Resp.Close()
    
    if result.Resp.IsSuccess() {
        successCount++
        fmt.Printf("Request %d succeeded: %d\n", i, result.Resp.StatusCode)
    } else {
        errorCount++
        fmt.Printf("Request %d returned error: %d\n", i, result.Resp.StatusCode)
    }
}

fmt.Printf("Summary: %d successful, %d failed\n", successCount, errorCount)
```

## Best Practices

### 1. Always Close Response Bodies

```go
result := <-client.GetAsync("https://api.example.com", nil)
if result.Err != nil {
    return result.Err
}
defer result.Resp.Close() // Always close the response
```

### 2. Use Context for Timeouts

```go
ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
defer cancel()

result := <-client.WithContextAsync(ctx, "GET", "https://api.example.com", nil, nil)
```

### 3. Handle Errors Appropriately

```go
result := <-client.GetAsync("https://api.example.com", nil)
if result.Err != nil {
    // Log the error and decide whether to retry
    log.Printf("Request failed: %v", result.Err)
    return fmt.Errorf("request failed: %w", result.Err)
}
defer result.Resp.Close()

if !result.Resp.IsSuccess() {
    return fmt.Errorf("HTTP error: %d", result.Resp.StatusCode)
}
```

### 4. Use Batch Operations for Multiple Related Requests

```go
// Good: Use batch operations for related requests
requests := []httpx.Request{
    httpx.MustNewRequest("GET", "https://api.example.com/users", nil),
    httpx.MustNewRequest("GET", "https://api.example.com/posts", nil),
    httpx.MustNewRequest("GET", "https://api.example.com/comments", nil),
}

results := <-asyncClient.BatchAsync(requests)

// Bad: Sequential async requests
result1 := <-client.GetAsync("https://api.example.com/users", nil)
result2 := <-client.GetAsync("https://api.example.com/posts", nil)
result3 := <-client.GetAsync("https://api.example.com/comments", nil)
```

### 5. Limit Concurrent Requests

```go
// Use a semaphore to limit concurrent requests
semaphore := make(chan struct{}, 10) // Limit to 10 concurrent requests

var wg sync.WaitGroup
for _, url := range urls {
    wg.Add(1)
    go func(u string) {
        defer wg.Done()
        
        semaphore <- struct{}{} // Acquire
        defer func() { <-semaphore }() // Release
        
        result := <-client.GetAsync(u, nil)
        // Process result...
    }(url)
}
wg.Wait()
```

## Performance Considerations

### 1. Goroutine Management

AsyncClient creates one goroutine per request. For high-volume scenarios, consider:

- Using `BatchAsync` to reduce goroutine overhead
- Implementing connection pooling with custom `http.Client`
- Using request rate limiting

### 2. Memory Usage

Each async request creates:
- One goroutine
- One channel
- One `AsyncResult` struct

For thousands of concurrent requests, monitor memory usage.

### 3. Connection Pooling

```go
// Configure HTTP client for better performance
customClient := &http.Client{
    Transport: &http.Transport{
        MaxIdleConns:        100,
        MaxIdleConnsPerHost: 10,
        IdleConnTimeout:     90 * time.Second,
    },
    Timeout: 30 * time.Second,
}

asyncClient := httpx.NewAsyncClient(customClient)
```

## Examples

### Example 1: API Data Aggregation

```go
func aggregateUserData(userID string) (*UserData, error) {
    client := httpx.NewAsyncClient(nil)
    
    // Start all requests concurrently
    userChan := client.GetAsync(fmt.Sprintf("https://api.example.com/users/%s", userID), nil)
    postsChan := client.GetAsync(fmt.Sprintf("https://api.example.com/users/%s/posts", userID), nil)
    followersChan := client.GetAsync(fmt.Sprintf("https://api.example.com/users/%s/followers", userID), nil)
    
    // Wait for all results
    results := httpx.WaitForAll([]<-chan httpx.AsyncResult{
        userChan, postsChan, followersChan,
    })
    
    var userData UserData
    
    // Process user data
    if results[0].Err != nil {
        return nil, fmt.Errorf("failed to fetch user: %w", results[0].Err)
    }
    defer results[0].Resp.Close()
    
    err := results[0].Resp.JSON(&userData.User)
    if err != nil {
        return nil, fmt.Errorf("failed to parse user data: %w", err)
    }
    
    // Process posts data
    if results[1].Err != nil {
        return nil, fmt.Errorf("failed to fetch posts: %w", results[1].Err)
    }
    defer results[1].Resp.Close()
    
    err = results[1].Resp.JSON(&userData.Posts)
    if err != nil {
        return nil, fmt.Errorf("failed to parse posts data: %w", err)
    }
    
    // Process followers data
    if results[2].Err != nil {
        return nil, fmt.Errorf("failed to fetch followers: %w", results[2].Err)
    }
    defer results[2].Resp.Close()
    
    err = results[2].Resp.JSON(&userData.Followers)
    if err != nil {
        return nil, fmt.Errorf("failed to parse followers data: %w", err)
    }
    
    return &userData, nil
}
```

### Example 2: Health Check Service

```go
func checkServiceHealth(services []string) map[string]bool {
    client := httpx.NewAsyncClient(&http.Client{Timeout: 5 * time.Second})
    
    // Create health check requests
    var requests []httpx.Request
    for _, service := range services {
        requests = append(requests, 
            httpx.MustNewRequest("GET", fmt.Sprintf("https://%s/health", service), nil))
    }
    
    // Execute all health checks concurrently
    results := <-client.BatchAsync(requests)
    
    healthStatus := make(map[string]bool)
    for i, result := range results {
        service := services[i]
        healthStatus[service] = result.Err == nil && result.Resp.IsSuccess()
        if result.Resp != nil {
            result.Resp.Close()
        }
    }
    
    return healthStatus
}
```

### Example 3: Rate-Limited Requests

```go
func fetchWithRateLimit(urls []string, maxConcurrent int) []string {
    client := httpx.NewAsyncClient(nil)
    semaphore := make(chan struct{}, maxConcurrent)
    
    var wg sync.WaitGroup
    results := make([]string, len(urls))
    
    for i, url := range urls {
        wg.Add(1)
        go func(index int, u string) {
            defer wg.Done()
            
            semaphore <- struct{}{} // Acquire semaphore
            defer func() { <-semaphore }() // Release semaphore
            
            result := <-client.GetAsync(u, nil)
            if result.Err == nil && result.Resp.IsSuccess() {
                body, _ := result.Resp.String()
                results[index] = body
            }
            if result.Resp != nil {
                result.Resp.Close()
            }
        }(i, url)
    }
    
    wg.Wait()
    return results
}
```

This guide covers the essential aspects of using the AsyncClient. For more advanced usage patterns and integration examples, refer to the main [README.md](README.md) documentation.
