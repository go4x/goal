// Package httpx provides convenient HTTP client functionality with both client-based and package-level methods.
//
// This package offers two ways to make HTTP requests:
//  1. Client-based methods: Create a RestClient instance and use its methods
//  2. Package-level convenience methods: Use the global DefaultClient directly
//
// The package-level methods are convenient for simple use cases where you don't need
// custom client configuration. They use a default client with 60-second timeout.
//
// Example usage:
//
//	// Using package-level methods (convenient for simple cases)
//	resp, err := httpx.Get("https://api.example.com/users", params)
//	resp, err := httpx.PostJson("https://api.example.com/users", jsonBody)
//
//	// Using client-based methods (for custom configuration)
//	client := httpx.NewClient(nil)
//	resp, err := client.Get("https://api.example.com/users", params)
//
// All methods return a *Response object that wraps the standard http.Response
// with additional convenience methods for reading and processing the response.
package httpx

import (
	"io"
	"net/url"
)

// Get sends a GET request to the specified URL with optional query parameters.
// It uses the default HTTP client with 60-second timeout.
//
// Parameters:
//   - url: The target URL for the request
//   - params: Query parameters to append to the URL (can be nil)
//
// Returns:
//   - *Response: A response wrapper with convenience methods
//   - error: Any error that occurred during the request
//
// Example:
//
//	params := url.Values{}
//	params.Set("page", "1")
//	params.Set("limit", "10")
//	resp, err := httpx.Get("https://api.example.com/users", params)
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer resp.Close()
//
//	if resp.IsSuccess() {
//		body, _ := resp.String()
//		fmt.Println(body)
//	}
func Get(url string, params url.Values) (*Response, error) {
	return DefaultClient.Get(url, params)
}

// GetWithBody sends a GET request with a request body.
// Note: This is not recommended by HTTP semantics as GET requests should not have bodies.
// Many proxies, caches, and tools may ignore GET request bodies.
//
// Parameters:
//   - url: The target URL for the request
//   - body: The request body (can be nil)
//
// Returns:
//   - *Response: A response wrapper with convenience methods
//   - error: Any error that occurred during the request
//
// Example:
//
//	body := strings.NewReader(`{"query": "search term"}`)
//	resp, err := httpx.GetWithBody("https://api.example.com/search", body)
//	// Use with caution - consider using POST for requests with bodies
func GetWithBody(url string, body io.Reader) (*Response, error) {
	return DefaultClient.GetWithBody(url, body)
}

// GetJson sends a GET request with a JSON request body.
// Note: This is not recommended by HTTP semantics as GET requests should not have bodies.
// Many proxies, caches, and tools may ignore GET request bodies.
//
// Parameters:
//   - url: The target URL for the request
//   - body: The JSON request body as an io.Reader
//
// Returns:
//   - *Response: A response wrapper with convenience methods
//   - error: Any error that occurred during the request
//
// Example:
//
//	jsonBody := strings.NewReader(`{"searchCriteria": {"category": "tech"}}`)
//	resp, err := httpx.GetJson("https://api.example.com/search", jsonBody)
//	// Use with caution - consider using POST for complex queries
func GetJson(url string, body io.Reader) (*Response, error) {
	return DefaultClient.GetJson(url, body)
}

// GetForm sends a GET request with form-urlencoded request body.
// Note: This is not recommended by HTTP semantics as GET requests should not have bodies.
// Many proxies, caches, and tools may ignore GET request bodies.
//
// Parameters:
//   - url: The target URL for the request
//   - body: The form data as an io.Reader
//
// Returns:
//   - *Response: A response wrapper with convenience methods
//   - error: Any error that occurred during the request
//
// Example:
//
//	formData := url.Values{}
//	formData.Set("search", "query")
//	formBody := strings.NewReader(formData.Encode())
//	resp, err := httpx.GetForm("https://api.example.com/search", formBody)
//	// Use with caution - consider using POST for complex queries
func GetForm(url string, body io.Reader) (*Response, error) {
	return DefaultClient.GetForm(url, body)
}

// Post sends a POST request to the specified URL with the given body.
// It uses the default HTTP client with 60-second timeout.
//
// Parameters:
//   - url: The target URL for the request
//   - body: The request body (can be nil)
//
// Returns:
//   - *Response: A response wrapper with convenience methods
//   - error: Any error that occurred during the request
//
// Example:
//
//	body := strings.NewReader(`{"name": "John", "email": "john@example.com"}`)
//	resp, err := httpx.Post("https://api.example.com/users", body)
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer resp.Close()
//
//	if resp.IsSuccess() {
//		fmt.Println("User created successfully")
//	}
func Post(url string, body io.Reader) (*Response, error) {
	return DefaultClient.Post(url, body)
}

// PostJson sends a POST request with JSON content type.
// The request body is expected to be JSON data.
//
// Parameters:
//   - url: The target URL for the request
//   - body: The JSON request body as an io.Reader
//
// Returns:
//   - *Response: A response wrapper with convenience methods
//   - error: Any error that occurred during the request
//
// Example:
//
//	jsonBody := strings.NewReader(`{"name": "John", "email": "john@example.com"}`)
//	resp, err := httpx.PostJson("https://api.example.com/users", jsonBody)
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer resp.Close()
//
//	var user map[string]interface{}
//	if resp.IsSuccess() {
//		resp.JSON(&user)
//		fmt.Printf("Created user: %+v\n", user)
//	}
func PostJson(url string, body io.Reader) (*Response, error) {
	return DefaultClient.PostJson(url, body)
}

// PostForm sends a POST request with form-urlencoded content type.
// The request body should be URL-encoded form data.
//
// Parameters:
//   - url: The target URL for the request
//   - body: The form data as an io.Reader (typically from url.Values.Encode())
//
// Returns:
//   - *Response: A response wrapper with convenience methods
//   - error: Any error that occurred during the request
//
// Example:
//
//	formData := url.Values{}
//	formData.Set("name", "John")
//	formData.Set("email", "john@example.com")
//	formBody := strings.NewReader(formData.Encode())
//
//	resp, err := httpx.PostForm("https://api.example.com/users", formBody)
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer resp.Close()
//
//	if resp.IsSuccess() {
//		fmt.Println("User created successfully")
//	}
func PostForm(url string, body io.Reader) (*Response, error) {
	return DefaultClient.PostForm(url, body)
}

// Put sends a PUT request to the specified URL with the given body.
// PUT is typically used for updating or creating resources.
//
// Parameters:
//   - url: The target URL for the request
//   - body: The request body (can be nil)
//
// Returns:
//   - *Response: A response wrapper with convenience methods
//   - error: Any error that occurred during the request
//
// Example:
//
//	body := strings.NewReader(`{"name": "Updated Name", "email": "updated@example.com"}`)
//	resp, err := httpx.Put("https://api.example.com/users/123", body)
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer resp.Close()
//
//	if resp.IsSuccess() {
//		fmt.Println("User updated successfully")
//	}
func Put(url string, body io.Reader) (*Response, error) {
	return DefaultClient.Put(url, body)
}

// PutJson sends a PUT request with JSON content type.
// The request body is expected to be JSON data.
//
// Parameters:
//   - url: The target URL for the request
//   - body: The JSON request body as an io.Reader
//
// Returns:
//   - *Response: A response wrapper with convenience methods
//   - error: Any error that occurred during the request
//
// Example:
//
//	jsonBody := strings.NewReader(`{"name": "Updated Name", "email": "updated@example.com"}`)
//	resp, err := httpx.PutJson("https://api.example.com/users/123", jsonBody)
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer resp.Close()
//
//	var user map[string]interface{}
//	if resp.IsSuccess() {
//		resp.JSON(&user)
//		fmt.Printf("Updated user: %+v\n", user)
//	}
func PutJson(url string, body io.Reader) (*Response, error) {
	return DefaultClient.PutJson(url, body)
}

// PutForm sends a PUT request with form-urlencoded content type.
// The request body should be URL-encoded form data.
//
// Parameters:
//   - url: The target URL for the request
//   - body: The form data as an io.Reader (typically from url.Values.Encode())
//
// Returns:
//   - *Response: A response wrapper with convenience methods
//   - error: Any error that occurred during the request
//
// Example:
//
//	formData := url.Values{}
//	formData.Set("name", "Updated Name")
//	formData.Set("email", "updated@example.com")
//	formBody := strings.NewReader(formData.Encode())
//
//	resp, err := httpx.PutForm("https://api.example.com/users/123", formBody)
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer resp.Close()
//
//	if resp.IsSuccess() {
//		fmt.Println("User updated successfully")
//	}
func PutForm(url string, body io.Reader) (*Response, error) {
	return DefaultClient.PutForm(url, body)
}

// Patch sends a PATCH request to the specified URL with the given body.
// PATCH is typically used for partial updates of resources.
//
// Parameters:
//   - url: The target URL for the request
//   - body: The request body (can be nil)
//
// Returns:
//   - *Response: A response wrapper with convenience methods
//   - error: Any error that occurred during the request
//
// Example:
//
//	body := strings.NewReader(`{"email": "newemail@example.com"}`)
//	resp, err := httpx.Patch("https://api.example.com/users/123", body)
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer resp.Close()
//
//	if resp.IsSuccess() {
//		fmt.Println("User patched successfully")
//	}
func Patch(url string, body io.Reader) (*Response, error) {
	return DefaultClient.Patch(url, body)
}

// PatchJson sends a PATCH request with JSON content type.
// The request body is expected to be JSON data with partial updates.
//
// Parameters:
//   - url: The target URL for the request
//   - body: The JSON request body as an io.Reader
//
// Returns:
//   - *Response: A response wrapper with convenience methods
//   - error: Any error that occurred during the request
//
// Example:
//
//	jsonBody := strings.NewReader(`{"email": "newemail@example.com"}`)
//	resp, err := httpx.PatchJson("https://api.example.com/users/123", jsonBody)
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer resp.Close()
//
//	var user map[string]interface{}
//	if resp.IsSuccess() {
//		resp.JSON(&user)
//		fmt.Printf("Patched user: %+v\n", user)
//	}
func PatchJson(url string, body io.Reader) (*Response, error) {
	return DefaultClient.PatchJson(url, body)
}

// PatchForm sends a PATCH request with form-urlencoded content type.
// The request body should be URL-encoded form data with partial updates.
//
// Parameters:
//   - url: The target URL for the request
//   - body: The form data as an io.Reader (typically from url.Values.Encode())
//
// Returns:
//   - *Response: A response wrapper with convenience methods
//   - error: Any error that occurred during the request
//
// Example:
//
//	formData := url.Values{}
//	formData.Set("email", "newemail@example.com")
//	formBody := strings.NewReader(formData.Encode())
//
//	resp, err := httpx.PatchForm("https://api.example.com/users/123", formBody)
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer resp.Close()
//
//	if resp.IsSuccess() {
//		fmt.Println("User patched successfully")
//	}
func PatchForm(url string, body io.Reader) (*Response, error) {
	return DefaultClient.PatchForm(url, body)
}

// Delete sends a DELETE request to the specified URL.
// DELETE is typically used for removing resources.
//
// Parameters:
//   - url: The target URL for the request
//
// Returns:
//   - *Response: A response wrapper with convenience methods
//   - error: Any error that occurred during the request
//
// Example:
//
//	resp, err := httpx.Delete("https://api.example.com/users/123")
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer resp.Close()
//
//	if resp.IsSuccess() {
//		fmt.Println("User deleted successfully")
//	}
func Delete(url string) (*Response, error) {
	return DefaultClient.Delete(url)
}

// DeleteJson sends a DELETE request with JSON request body.
// This is useful for conditional deletes or batch delete operations.
//
// Parameters:
//   - url: The target URL for the request
//   - body: The JSON request body as an io.Reader
//
// Returns:
//   - *Response: A response wrapper with convenience methods
//   - error: Any error that occurred during the request
//
// Example:
//
//	jsonBody := strings.NewReader(`{"confirmCode": "ABC123", "reason": "user_request"}`)
//	resp, err := httpx.DeleteJson("https://api.example.com/users/123", jsonBody)
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer resp.Close()
//
//	if resp.IsSuccess() {
//		fmt.Println("User deleted with confirmation")
//	}
func DeleteJson(url string, body io.Reader) (*Response, error) {
	return DefaultClient.DeleteJson(url, body)
}

// DeleteForm sends a DELETE request with form-urlencoded request body.
// This is useful for conditional deletes or batch delete operations.
//
// Parameters:
//   - url: The target URL for the request
//   - body: The form data as an io.Reader (typically from url.Values.Encode())
//
// Returns:
//   - *Response: A response wrapper with convenience methods
//   - error: Any error that occurred during the request
//
// Example:
//
//	formData := url.Values{}
//	formData.Set("confirmCode", "ABC123")
//	formData.Set("reason", "user_request")
//	formBody := strings.NewReader(formData.Encode())
//
//	resp, err := httpx.DeleteForm("https://api.example.com/users/123", formBody)
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer resp.Close()
//
//	if resp.IsSuccess() {
//		fmt.Println("User deleted with confirmation")
//	}
func DeleteForm(url string, body io.Reader) (*Response, error) {
	return DefaultClient.DeleteForm(url, body)
}
