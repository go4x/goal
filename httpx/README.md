# HTTPX - Advanced HTTP Client Library for Go

HTTPX is a comprehensive HTTP client library for Go that provides both synchronous and asynchronous HTTP request capabilities. It offers a clean, intuitive API with intelligent method selection, comprehensive error handling, and extensive customization options.

## Features

### 🚀 Core Features
- **Synchronous HTTP Client**: Full-featured REST client with all HTTP methods
- **Asynchronous HTTP Client**: Non-blocking HTTP requests with channel-based results
- **Intelligent Method Selection**: Automatically chooses the appropriate HTTP method based on content type
- **Request Interface**: Flexible request abstraction with fluent API
- **Response Wrapper**: Enhanced response handling with convenience methods
- **Package-level Convenience Methods**: Direct access to HTTP methods via default client

### 🔧 HTTP Methods Support
- **GET**: Standard GET requests with query parameters
- **POST**: POST requests with JSON, form data, or custom body
- **PUT**: PUT requests with JSON, form data, or custom body
- **PATCH**: PATCH requests with JSON, form data, or custom body
- **DELETE**: DELETE requests with optional JSON/form body
- **OPTIONS**: OPTIONS requests for CORS preflight

### 🎯 Advanced Features
- **Context Support**: Full context.Context integration for timeouts and cancellation
- **Request Options**: Functional options pattern for request configuration
- **Batch Operations**: Concurrent request execution with result aggregation
- **Resource Management**: Automatic response body closing and proper resource cleanup
- **Error Handling**: Comprehensive error handling with detailed error information
- **Race Condition Safe**: Tested with Go's race detector

## Installation

```bash
go get github.com/go4x/goal/httpx
```

## Quick Start

### Basic Usage

```go
package main

import (
    "fmt"
    "github.com/go4x/goal/httpx"
)

func main() {
    // Simple GET request
    resp, err := httpx.Get("https://httpbin.org/get")
    if err != nil {
        panic(err)
    }
    defer resp.Close()
    
    fmt.Printf("Status: %d\n", resp.StatusCode)
}
```

### Using RestClient

```go
// Create a new REST client
client := httpx.NewRestClient(nil)

// GET request with query parameters
params := url.Values{"key": {"value"}}
resp, err := client.Get("https://httpbin.org/get", params)
if err != nil {
    panic(err)
}
defer resp.Close()

// POST JSON request
jsonBody := strings.NewReader(`{"name": "John", "age": 30}`)
resp, err = client.PostJson("https://httpbin.org/post", jsonBody)
if err != nil {
    panic(err)
}
defer resp.Close()
```

### Using AsyncClient

```go
// Create an async client
asyncClient := httpx.NewAsyncClient(nil)

// Send asynchronous requests
getResult := <-asyncClient.GetAsync("https://httpbin.org/get", nil)
postResult := <-asyncClient.PostJsonAsync("https://httpbin.org/post", jsonBody)

// Process results
if getResult.Err != nil {
    fmt.Printf("GET failed: %v\n", getResult.Err)
} else {
    defer getResult.Resp.Close()
    fmt.Printf("GET Status: %d\n", getResult.Resp.StatusCode)
}

if postResult.Err != nil {
    fmt.Printf("POST failed: %v\n", postResult.Err)
} else {
    defer postResult.Resp.Close()
    fmt.Printf("POST Status: %d\n", postResult.Resp.StatusCode)
}
```

### Batch Requests

```go
// Create multiple requests
requests := []httpx.Request{
    httpx.MustNewRequest("GET", "https://httpbin.org/get", nil),
    httpx.MustNewRequest("POST", "https://httpbin.org/post", jsonBody, 
        httpx.WithContentType("application/json")),
}

// Execute batch requests
asyncClient := httpx.NewAsyncClient(nil)
results := <-asyncClient.BatchAsync(requests)

// Process all results
for i, result := range results {
    if result.Err != nil {
        fmt.Printf("Request %d failed: %v\n", i, result.Err)
    } else {
        defer result.Resp.Close()
        fmt.Printf("Request %d succeeded: %d\n", i, result.Resp.StatusCode)
    }
}
```

## API Reference

### RestClient

#### Constructor
```go
func NewRestClient(client *http.Client) *RestClient
```

#### HTTP Methods
```go
// GET requests
func (rc *RestClient) Get(url string, params url.Values) (*Response, error)
func (rc *RestClient) GetWithBody(url string, body io.Reader) (*Response, error)
func (rc *RestClient) GetJson(url string, body io.Reader) (*Response, error)
func (rc *RestClient) GetForm(url string, body io.Reader) (*Response, error)

// POST requests
func (rc *RestClient) Post(url string, body io.Reader) (*Response, error)
func (rc *RestClient) PostJson(url string, body io.Reader) (*Response, error)
func (rc *RestClient) PostForm(url string, body io.Reader) (*Response, error)

// PUT requests
func (rc *RestClient) Put(url string, body io.Reader) (*Response, error)
func (rc *RestClient) PutJson(url string, body io.Reader) (*Response, error)
func (rc *RestClient) PutForm(url string, body io.Reader) (*Response, error)

// PATCH requests
func (rc *RestClient) Patch(url string, body io.Reader) (*Response, error)
func (rc *RestClient) PatchJson(url string, body io.Reader) (*Response, error)
func (rc *RestClient) PatchForm(url string, body io.Reader) (*Response, error)

// DELETE requests
func (rc *RestClient) Delete(url string) (*Response, error)
func (rc *RestClient) DeleteJson(url string, body io.Reader) (*Response, error)
func (rc *RestClient) DeleteForm(url string, body io.Reader) (*Response, error)

// OPTIONS requests
func (rc *RestClient) Options(url string, options ...RequestOption) (*Response, error)
```

### AsyncClient

#### Constructor
```go
func NewAsyncClient(client *http.Client) *AsyncClient
```

#### Asynchronous Methods
All RestClient methods have corresponding async versions that return `<-chan AsyncResult`:

```go
func (ac *AsyncClient) GetAsync(url string, params url.Values) <-chan AsyncResult
func (ac *AsyncClient) PostJsonAsync(url string, body io.Reader) <-chan AsyncResult
// ... and so on for all HTTP methods
```

#### Batch Operations
```go
func (ac *AsyncClient) BatchAsync(requests []Request) <-chan []AsyncResult
func WaitForAll(resultChans []<-chan AsyncResult) []AsyncResult
func WaitForFirst(resultChans []<-chan AsyncResult) AsyncResult
```

### Request Interface

```go
type Request interface {
    GetURL() *url.URL
    GetMethod() string
    GetBody() (io.ReadCloser, error)
    GetHeader() http.Header
    GetContext() context.Context
    // ... additional getter methods
    
    WithURL(url *url.URL) Request
    WithMethod(method string) Request
    WithBody(body io.ReadCloser) Request
    WithHeader(key string, value string) Request
    WithContext(ctx context.Context) Request
    // ... additional setter methods
}
```

#### Request Creation
```go
func NewRequest(method string, url string, body io.Reader, options ...RequestOption) (Request, error)
func MustNewRequest(method string, url string, body io.Reader, options ...RequestOption) Request
func NewRequestWithContext(ctx context.Context, method string, url string, body io.Reader, options ...RequestOption) (Request, error)
func MustNewRequestWithContext(ctx context.Context, method string, url string, body io.Reader, options ...RequestOption) Request
```

#### Request Options
```go
func WithContentType(contentType string) RequestOption
func WithAuthorization(authorization string) RequestOption
func WithHeader(key string, value string) RequestOption
func WithContext(ctx context.Context) RequestOption
```

### Response Wrapper

```go
type Response struct {
    *http.Response
}

// Status checking
func (r *Response) IsSuccess() bool
func (r *Response) IsClientError() bool
func (r *Response) IsServerError() bool
func (r *Response) Status() int

// Body reading
func (r *Response) Bytes() ([]byte, error)
func (r *Response) String() (string, error)
func (r *Response) JSON(v interface{}) error

// Header utilities
func (r *Response) HeaderValue(key string) string
func (r *Response) ContentType() string
func (r *Response) ContentLength() int64

// Resource management
func (r *Response) Close() error
```

### Package-level Convenience Methods

```go
// All HTTP methods are available at package level
func Get(url string, params url.Values) (*Response, error)
func PostJson(url string, body io.Reader) (*Response, error)
func PutForm(url string, body io.Reader) (*Response, error)
// ... and so on for all HTTP methods
```

## Examples

### Context with Timeout

```go
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()

// Using RestClient
resp, err := client.GetWithContext(ctx, "https://httpbin.org/delay/2")
if err != nil {
    panic(err)
}
defer resp.Close()

// Using AsyncClient
result := <-asyncClient.WithContextAsync(ctx, "GET", "https://httpbin.org/delay/2", nil, nil)
if result.Err != nil {
    fmt.Printf("Request failed: %v\n", result.Err)
}
```

### Custom Headers

```go
// Using Request Options
req, err := httpx.NewRequest("POST", "https://api.example.com/data", body,
    httpx.WithContentType("application/json"),
    httpx.WithAuthorization("Bearer token123"),
    httpx.WithHeader("X-Custom-Header", "value"),
)
if err != nil {
    panic(err)
}

// Using RestClient with Options
resp, err := client.PostJson("https://api.example.com/data", body,
    httpx.WithAuthorization("Bearer token123"),
)
```

### Error Handling

```go
resp, err := httpx.Get("https://httpbin.org/status/404")
if err != nil {
    // Network or request creation error
    fmt.Printf("Request failed: %v\n", err)
    return
}
defer resp.Close()

if !resp.IsSuccess() {
    // HTTP error response
    body, _ := resp.String()
    fmt.Printf("HTTP Error %d: %s\n", resp.StatusCode, body)
    return
}

// Success
fmt.Printf("Success: %d\n", resp.StatusCode)
```

### JSON Handling

```go
// Sending JSON
type User struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
}

user := User{Name: "John", Age: 30}
jsonData, _ := json.Marshal(user)
resp, err := httpx.PostJson("https://api.example.com/users", strings.NewReader(string(jsonData)))

// Receiving JSON
type ResponseData struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
}

var responseData ResponseData
err = resp.JSON(&responseData)
if err != nil {
    panic(err)
}
```

## Best Practices

### 1. Always Close Responses
```go
resp, err := httpx.Get("https://example.com")
if err != nil {
    panic(err)
}
defer resp.Close() // Always close the response
```

### 2. Use Context for Timeouts
```go
ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
defer cancel()

resp, err := httpx.GetWithContext(ctx, "https://example.com")
```

### 3. Handle Errors Appropriately
```go
resp, err := httpx.Get("https://example.com")
if err != nil {
    // Handle network errors
    return fmt.Errorf("request failed: %w", err)
}
defer resp.Close()

if !resp.IsSuccess() {
    // Handle HTTP errors
    return fmt.Errorf("HTTP error: %d", resp.StatusCode)
}
```

### 4. Use Batch Operations for Multiple Requests
```go
requests := []httpx.Request{
    httpx.MustNewRequest("GET", "https://api1.example.com", nil),
    httpx.MustNewRequest("GET", "https://api2.example.com", nil),
}

asyncClient := httpx.NewAsyncClient(nil)
results := <-asyncClient.BatchAsync(requests)

for _, result := range results {
    if result.Err != nil {
        // Handle individual request errors
        continue
    }
    defer result.Resp.Close()
    // Process successful responses
}
```

## Performance

HTTPX is designed for high performance with minimal overhead:

- **Zero-copy operations** where possible
- **Efficient resource management** with automatic cleanup
- **Concurrent request execution** with proper synchronization
- **Minimal memory allocations** in hot paths

### Benchmarks

```bash
go test -bench=. ./httpx
```

## Testing

Run the test suite:

```bash
# Run all tests
go test ./httpx

# Run tests with coverage
go test -cover ./httpx

# Run tests with race detection
go test -race ./httpx

# Run benchmarks
go test -bench=. ./httpx
```

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the Apache License 2.0 - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- Built on top of Go's standard `net/http` package
- Inspired by modern HTTP client libraries
- Thanks to the Go community for feedback and contributions