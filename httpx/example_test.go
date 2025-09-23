package httpx

import (
	"fmt"
	"net/url"
	"strings"
)

// ExampleNewRestClient demonstrates creating a new RestClient
func ExampleNewRestClient() {
	client := NewRestClient(nil)
	fmt.Printf("Client created with timeout: %v\n", client.Timeout)
	// Output: Client created with timeout: 1m0s
}

// ExampleRestClient_Get demonstrates making a GET request
func ExampleRestClient_Get() {
	client := NewRestClient(nil)

	// GET request with query parameters
	params := url.Values{}
	params.Set("page", "1")
	params.Set("limit", "10")

	resp, err := client.Get("https://httpbin.org/get", params)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	defer func() { _ = resp.Close() }()

	if resp.IsSuccess() {
		fmt.Printf("Request successful: %d\n", resp.StatusCode)
	}
	// Output: Request successful: 200
}

// ExampleRestClient_GetJson demonstrates making a GET request with JSON body (use with caution)
func ExampleRestClient_GetJson() {
	client := NewRestClient(nil)

	// GET request with JSON body (not recommended - use POST for complex queries instead)
	jsonBody := strings.NewReader(`{"searchCriteria": {"category": "tech", "dateRange": "2024"}}`)

	resp, err := client.GetJson("https://httpbin.org/get", jsonBody)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	defer func() { _ = resp.Close() }()

	if resp.IsSuccess() {
		fmt.Printf("GET JSON request successful: %d\n", resp.StatusCode)
	}
	// Output: GET JSON request successful: 200
}

// ExampleRestClient_GetForm demonstrates making a GET request with form data (use with caution)
func ExampleRestClient_GetForm() {
	client := NewRestClient(nil)

	// GET request with form data (not recommended - use POST for complex queries instead)
	formData := url.Values{}
	formData.Set("search", "complex query")
	formData.Set("filters", "category=tech,date=2024")
	formBody := strings.NewReader(formData.Encode())

	resp, err := client.GetForm("https://httpbin.org/get", formBody)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	defer func() { _ = resp.Close() }()

	if resp.IsSuccess() {
		fmt.Printf("GET form request successful: %d\n", resp.StatusCode)
	}
	// Output: GET form request successful: 200
}

// ExampleRestClient_PostJson demonstrates making a POST request with JSON
func ExampleRestClient_PostJson() {
	client := NewRestClient(nil)

	// POST request with JSON body
	jsonBody := strings.NewReader(`{"name": "John", "age": 30}`)

	resp, err := client.PostJson("https://httpbin.org/post", jsonBody)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	defer func() { _ = resp.Close() }()

	if resp.IsSuccess() {
		fmt.Printf("POST request successful: %d\n", resp.StatusCode)
	}
	// Output: POST request successful: 200
}

// ExampleRestClient_PostForm demonstrates making a POST request with form data
func ExampleRestClient_PostForm() {
	client := NewRestClient(nil)

	// POST request with form data
	formData := url.Values{}
	formData.Set("username", "john_doe")
	formData.Set("password", "secret123")
	formBody := strings.NewReader(formData.Encode())

	resp, err := client.PostForm("https://httpbin.org/post", formBody)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	defer func() { _ = resp.Close() }()

	if resp.IsSuccess() {
		fmt.Printf("Form POST successful: %d\n", resp.StatusCode)
	}
	// Output: Form POST successful: 200
}

// ExampleRestClient_PatchJson demonstrates making a PATCH request with JSON
func ExampleRestClient_PatchJson() {
	client := NewRestClient(nil)

	// PATCH request with JSON body
	jsonBody := strings.NewReader(`{"title": "Updated Title", "completed": true}`)

	resp, err := client.PatchJson("https://httpbin.org/patch", jsonBody)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	defer func() { _ = resp.Close() }()

	if resp.IsSuccess() {
		fmt.Printf("PATCH JSON request successful: %d\n", resp.StatusCode)
	}
	// Output: PATCH JSON request successful: 200
}

// ExampleRestClient_PatchForm demonstrates making a PATCH request with form data
func ExampleRestClient_PatchForm() {
	client := NewRestClient(nil)

	// PATCH request with form data
	formData := url.Values{}
	formData.Set("name", "updated_name")
	formData.Set("status", "active")
	formBody := strings.NewReader(formData.Encode())

	resp, err := client.PatchForm("https://httpbin.org/patch", formBody)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	defer func() { _ = resp.Close() }()

	if resp.IsSuccess() {
		fmt.Printf("PATCH form request successful: %d\n", resp.StatusCode)
	}
	// Output: PATCH form request successful: 200
}

// ExampleRestClient_DeleteJson demonstrates making a DELETE request with JSON body
func ExampleRestClient_DeleteJson() {
	client := NewRestClient(nil)

	// DELETE request with JSON body (e.g., for batch delete or conditional delete)
	jsonBody := strings.NewReader(`{"confirmCode": "ABC123", "reason": "customer_request"}`)

	resp, err := client.DeleteJson("https://httpbin.org/delete", jsonBody)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	defer func() { _ = resp.Close() }()

	if resp.IsSuccess() {
		fmt.Printf("DELETE JSON request successful: %d\n", resp.StatusCode)
	}
	// Output: DELETE JSON request successful: 200
}

// ExampleRestClient_DeleteForm demonstrates making a DELETE request with form data
func ExampleRestClient_DeleteForm() {
	client := NewRestClient(nil)

	// DELETE request with form data
	formData := url.Values{}
	formData.Set("confirmCode", "ABC123")
	formData.Set("reason", "customer_request")
	formBody := strings.NewReader(formData.Encode())

	resp, err := client.DeleteForm("https://httpbin.org/delete", formBody)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	defer func() { _ = resp.Close() }()

	if resp.IsSuccess() {
		fmt.Printf("DELETE form request successful: %d\n", resp.StatusCode)
	}
	// Output: DELETE form request successful: 200
}

// ExampleRestClient_Delete demonstrates making a simple DELETE request
func ExampleRestClient_Delete() {
	client := NewRestClient(nil)

	resp, err := client.Delete("https://httpbin.org/delete")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	defer func() { _ = resp.Close() }()

	if resp.IsSuccess() {
		fmt.Printf("DELETE request successful: %d\n", resp.StatusCode)
	}
	// Output: DELETE request successful: 200
}
