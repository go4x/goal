package httpx

import (
	"io"
	"net/http"
	"net/url"
	"time"
)

// DefaultClient is a default HTTP client with 60-second timeout
var DefaultClient = &RestClient{Client: &http.Client{Timeout: time.Second * 60}}

// RestClient is a RESTful HTTP client that provides convenient methods for making HTTP requests.
// It wraps the standard http.Client and provides additional functionality for common REST operations.
//
// The RestClient supports all standard HTTP methods (GET, POST, PUT, PATCH, DELETE) with various
// content types (JSON, form data, custom). It also supports request options for setting headers,
// content types, and contexts.
//
// Example usage:
//
//	client := httpx.NewClient(nil)
//
//	// Simple GET request
//	resp, err := client.Get("https://api.example.com/users", params)
//
//	// POST with JSON
//	jsonBody := strings.NewReader(`{"name": "John"}`)
//	resp, err := client.PostJson("https://api.example.com/users", jsonBody)
//
//	// Custom request with options
//	resp, err := client.Post("https://api.example.com/users", body,
//		httpx.WithContentType("application/json"),
//		httpx.WithHeader("Authorization", "Bearer token"))
type RestClient struct {
	*http.Client
}

// NewRestClient creates a new RestClient with the given http.Client.
// If client is nil, a default client with 60-second timeout is used.
func NewRestClient(client *http.Client) *RestClient {
	if client == nil {
		return DefaultClient
	}
	return &RestClient{Client: client}
}

// Get sends a GET request to the specified URL with optional query parameters
func (c *RestClient) Get(url string, params url.Values) (*Response, error) {
	uri, err := c.buildUrl(url, params)
	if err != nil {
		return nil, err
	}
	resp, err := c.Client.Get(uri)
	if err != nil {
		return nil, err
	}
	return NewResponse(resp), nil
}

// GetWithBody sends a GET request with a body (use with caution - not recommended by HTTP semantics)
// Note: Many proxies, caches, and tools may ignore GET request bodies
func (c *RestClient) GetWithBody(url string, body io.Reader, options ...RequestOption) (*Response, error) {
	return c.send(http.MethodGet, url, body, options...)
}

// GetJson sends a GET request with JSON body (use with caution - not recommended by HTTP semantics)
// Note: Many proxies, caches, and tools may ignore GET request bodies
func (c *RestClient) GetJson(url string, body io.Reader, options ...RequestOption) (*Response, error) {
	options = append(options, WithContentType(ContentTypeApplicationJson))
	return c.send(http.MethodGet, url, body, options...)
}

// GetForm sends a GET request with form-urlencoded body (use with caution - not recommended by HTTP semantics)
// Note: Many proxies, caches, and tools may ignore GET request bodies
func (c *RestClient) GetForm(url string, body io.Reader, options ...RequestOption) (*Response, error) {
	options = append(options, WithContentType(ContentTypeApplicationFormUrlencoded))
	return c.send(http.MethodGet, url, body, options...)
}

// buildUrl constructs the full URL by combining base URL and query parameters
func (c *RestClient) buildUrl(s string, params url.Values) (string, error) {
	u, err := url.Parse(s)
	if err != nil {
		return s, err
	}
	if params != nil {
		q := u.Query()
		for k, v := range params {
			for _, value := range v {
				q.Add(k, value)
			}
		}
		u.RawQuery = q.Encode()
	}
	return u.String(), nil
}

// Post sends a POST request with the given body
func (c *RestClient) Post(url string, body io.Reader, options ...RequestOption) (*Response, error) {
	return c.send(http.MethodPost, url, body, options...)
}

// PostJson sends a POST request with JSON content type
func (c *RestClient) PostJson(url string, body io.Reader, options ...RequestOption) (*Response, error) {
	options = append(options, WithContentType(ContentTypeApplicationJson))
	return c.send(http.MethodPost, url, body, options...)
}

// PostForm sends a POST request with form-urlencoded content type
func (c *RestClient) PostForm(url string, body io.Reader, options ...RequestOption) (*Response, error) {
	options = append(options, WithContentType(ContentTypeApplicationFormUrlencoded))
	return c.send(http.MethodPost, url, body, options...)
}

// Put sends a PUT request with the given body
func (c *RestClient) Put(url string, body io.Reader, options ...RequestOption) (*Response, error) {
	return c.send(http.MethodPut, url, body, options...)
}

// PutJson sends a PUT request with JSON content type
func (c *RestClient) PutJson(url string, body io.Reader, options ...RequestOption) (*Response, error) {
	options = append(options, WithContentType(ContentTypeApplicationJson))
	return c.send(http.MethodPut, url, body, options...)
}

// PutForm sends a PUT request with form-urlencoded content type
func (c *RestClient) PutForm(url string, body io.Reader, options ...RequestOption) (*Response, error) {
	options = append(options, WithContentType(ContentTypeApplicationFormUrlencoded))
	return c.send(http.MethodPut, url, body, options...)
}

// Patch sends a PATCH request with the given body
func (c *RestClient) Patch(url string, body io.Reader, options ...RequestOption) (*Response, error) {
	return c.send(http.MethodPatch, url, body, options...)
}

// PatchJson sends a PATCH request with JSON content type
func (c *RestClient) PatchJson(url string, body io.Reader, options ...RequestOption) (*Response, error) {
	options = append(options, WithContentType(ContentTypeApplicationJson))
	return c.send(http.MethodPatch, url, body, options...)
}

// PatchForm sends a PATCH request with form-urlencoded content type
func (c *RestClient) PatchForm(url string, body io.Reader, options ...RequestOption) (*Response, error) {
	options = append(options, WithContentType(ContentTypeApplicationFormUrlencoded))
	return c.send(http.MethodPatch, url, body, options...)
}

// Delete sends a DELETE request
func (c *RestClient) Delete(url string, options ...RequestOption) (*Response, error) {
	return c.send(http.MethodDelete, url, nil, options...)
}

// DeleteJson sends a DELETE request with JSON content type
func (c *RestClient) DeleteJson(url string, body io.Reader, options ...RequestOption) (*Response, error) {
	options = append(options, WithContentType(ContentTypeApplicationJson))
	return c.send(http.MethodDelete, url, body, options...)
}

// DeleteForm sends a DELETE request with form-urlencoded content type
func (c *RestClient) DeleteForm(url string, body io.Reader, options ...RequestOption) (*Response, error) {
	options = append(options, WithContentType(ContentTypeApplicationFormUrlencoded))
	return c.send(http.MethodDelete, url, body, options...)
}

func (c *RestClient) send(method string, url string, body io.Reader, options ...RequestOption) (*Response, error) {
	uri, err := c.buildUrl(url, nil)
	if err != nil {
		return nil, err
	}
	req, err := NewRequest(method, uri, body, options...)
	if err != nil {
		return nil, err
	}
	resp, err := c.Do(req.GetRequest())
	if err != nil {
		return nil, err
	}
	return NewResponse(resp), nil
}

// Options sends an OPTIONS request
func (c *RestClient) Options(url string, options ...RequestOption) (*Response, error) {
	return c.send(http.MethodOptions, url, nil, options...)
}
