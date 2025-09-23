package httpx

import (
	"fmt"
	"net/url"
	"strings"
)

// ExampleGet demonstrates making a GET request with query parameters
func ExampleGet() {
	// Create query parameters
	params := url.Values{}
	params.Set("page", "1")
	params.Set("limit", "10")
	params.Set("search", "example")

	// Make GET request
	resp, err := Get("https://httpbin.org/get", params)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	defer func() { _ = resp.Close() }()

	if resp.IsSuccess() {
		fmt.Printf("GET request successful: %d\n", resp.StatusCode)
	}
	// Output: GET request successful: 200
}

// ExampleGetJson demonstrates making a GET request with JSON body (use with caution)
func ExampleGetJson() {
	// Create JSON body (not recommended for GET requests)
	jsonBody := strings.NewReader(`{"searchCriteria": {"category": "tech", "dateRange": "2024"}}`)

	// Make GET request with JSON body
	resp, err := GetJson("https://httpbin.org/get", jsonBody)
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

// ExamplePost demonstrates making a POST request
func ExamplePost() {
	// Create request body
	body := strings.NewReader(`{"name": "John Doe", "email": "john@example.com"}`)

	// Make POST request
	resp, err := Post("https://httpbin.org/post", body)
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

// ExamplePostJson demonstrates making a POST request with JSON content type
func ExamplePostJson() {
	// Create JSON body
	jsonBody := strings.NewReader(`{"name": "John Doe", "email": "john@example.com"}`)

	// Make POST request with JSON
	resp, err := PostJson("https://httpbin.org/post", jsonBody)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	defer func() { _ = resp.Close() }()

	if resp.IsSuccess() {
		fmt.Printf("POST JSON request successful: %d\n", resp.StatusCode)
	}
	// Output: POST JSON request successful: 200
}

// ExamplePostForm demonstrates making a POST request with form data
func ExamplePostForm() {
	// Create form data
	formData := url.Values{}
	formData.Set("name", "John Doe")
	formData.Set("email", "john@example.com")
	formBody := strings.NewReader(formData.Encode())

	// Make POST request with form data
	resp, err := PostForm("https://httpbin.org/post", formBody)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	defer func() { _ = resp.Close() }()

	if resp.IsSuccess() {
		fmt.Printf("POST form request successful: %d\n", resp.StatusCode)
	}
	// Output: POST form request successful: 200
}

// ExamplePut demonstrates making a PUT request
func ExamplePut() {
	// Create request body for updating a resource
	body := strings.NewReader(`{"name": "Updated Name", "email": "updated@example.com"}`)

	// Make PUT request
	resp, err := Put("https://httpbin.org/put", body)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	defer func() { _ = resp.Close() }()

	if resp.IsSuccess() {
		fmt.Printf("PUT request successful: %d\n", resp.StatusCode)
	}
	// Output: PUT request successful: 200
}

// ExamplePutJson demonstrates making a PUT request with JSON content type
func ExamplePutJson() {
	// Create JSON body for updating a resource
	jsonBody := strings.NewReader(`{"name": "Updated Name", "email": "updated@example.com"}`)

	// Make PUT request with JSON
	resp, err := PutJson("https://httpbin.org/put", jsonBody)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	defer func() { _ = resp.Close() }()

	if resp.IsSuccess() {
		fmt.Printf("PUT JSON request successful: %d\n", resp.StatusCode)
	}
	// Output: PUT JSON request successful: 200
}

// ExamplePatch demonstrates making a PATCH request for partial updates
func ExamplePatch() {
	// Create request body for partial update
	body := strings.NewReader(`{"email": "newemail@example.com"}`)

	// Make PATCH request
	resp, err := Patch("https://httpbin.org/patch", body)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	defer func() { _ = resp.Close() }()

	if resp.IsSuccess() {
		fmt.Printf("PATCH request successful: %d\n", resp.StatusCode)
	}
	// Output: PATCH request successful: 200
}

// ExamplePatchJson demonstrates making a PATCH request with JSON content type
func ExamplePatchJson() {
	// Create JSON body for partial update
	jsonBody := strings.NewReader(`{"email": "newemail@example.com"}`)

	// Make PATCH request with JSON
	resp, err := PatchJson("https://httpbin.org/patch", jsonBody)
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

// ExamplePatchForm demonstrates making a PATCH request with form data
func ExamplePatchForm() {
	// Create form data for partial update
	formData := url.Values{}
	formData.Set("email", "newemail@example.com")
	formBody := strings.NewReader(formData.Encode())

	// Make PATCH request with form data
	resp, err := PatchForm("https://httpbin.org/patch", formBody)
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

// ExampleDelete demonstrates making a DELETE request
func ExampleDelete() {
	// Make DELETE request
	resp, err := Delete("https://httpbin.org/delete")
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

// ExampleDeleteJson demonstrates making a DELETE request with JSON body
func ExampleDeleteJson() {
	// Create JSON body for conditional delete
	jsonBody := strings.NewReader(`{"confirmCode": "ABC123", "reason": "user_request"}`)

	// Make DELETE request with JSON body
	resp, err := DeleteJson("https://httpbin.org/delete", jsonBody)
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

// ExampleDeleteForm demonstrates making a DELETE request with form data
func ExampleDeleteForm() {
	// Create form data for conditional delete
	formData := url.Values{}
	formData.Set("confirmCode", "ABC123")
	formData.Set("reason", "user_request")
	formBody := strings.NewReader(formData.Encode())

	// Make DELETE request with form data
	resp, err := DeleteForm("https://httpbin.org/delete", formBody)
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

// ExampleGetWithBody demonstrates making a GET request with body (use with caution)
func ExampleGetWithBody() {
	// Create request body (not recommended for GET requests)
	body := strings.NewReader(`{"query": "search term"}`)

	// Make GET request with body
	resp, err := GetWithBody("https://httpbin.org/get", body)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	defer func() { _ = resp.Close() }()

	if resp.IsSuccess() {
		fmt.Printf("GET with body request successful: %d\n", resp.StatusCode)
	}
	// Output: GET with body request successful: 200
}

// ExampleGetForm demonstrates making a GET request with form data (use with caution)
func ExampleGetForm() {
	// Create form data (not recommended for GET requests)
	formData := url.Values{}
	formData.Set("search", "complex query")
	formData.Set("filters", "category=tech,date=2024")
	formBody := strings.NewReader(formData.Encode())

	// Make GET request with form data
	resp, err := GetForm("https://httpbin.org/get", formBody)
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

// ExamplePutForm demonstrates making a PUT request with form data
func ExamplePutForm() {
	// Create form data for updating a resource
	formData := url.Values{}
	formData.Set("name", "Updated Name")
	formData.Set("email", "updated@example.com")
	formBody := strings.NewReader(formData.Encode())

	// Make PUT request with form data
	resp, err := PutForm("https://httpbin.org/put", formBody)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	defer func() { _ = resp.Close() }()

	if resp.IsSuccess() {
		fmt.Printf("PUT form request successful: %d\n", resp.StatusCode)
	}
	// Output: PUT form request successful: 200
}
