package httpx

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Response wraps the standard http.Response with additional convenience methods.
type Response struct {
	*http.Response
}

// NewResponse creates a new Response from an http.Response.
func NewResponse(resp *http.Response) *Response {
	return &Response{Response: resp}
}

// IsSuccess returns true if the response status code is in the 2xx range.
func (r *Response) IsSuccess() bool {
	return r.StatusCode >= http.StatusOK && r.StatusCode < http.StatusMultipleChoices
}

// IsClientError returns true if the response status code is in the 4xx range.
func (r *Response) IsClientError() bool {
	return r.StatusCode >= http.StatusBadRequest && r.StatusCode < http.StatusInternalServerError
}

// IsServerError returns true if the response status code is in the 5xx range.
func (r *Response) IsServerError() bool {
	return r.StatusCode >= http.StatusInternalServerError
}

// Status returns the status code and status text.
func (r *Response) Status() (int, string) {
	return r.StatusCode, r.Response.Status
}

// Bytes reads and returns the response body as a byte slice.
func (r *Response) Bytes() ([]byte, error) {
	defer func() { _ = r.Body.Close() }()
	return io.ReadAll(r.Body)
}

// String reads and returns the response body as a string.
func (r *Response) String() (string, error) {
	bytes, err := r.Bytes()
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// JSON unmarshals the response body into the provided interface.
func (r *Response) JSON(v interface{}) error {
	bytes, err := r.Bytes()
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	if err := json.Unmarshal(bytes, v); err != nil {
		return fmt.Errorf("failed to unmarshal JSON: %w", err)
	}

	return nil
}

// Close closes the response body.
func (r *Response) Close() error {
	if r.Body != nil {
		return r.Body.Close()
	}
	return nil
}

// HeaderValue returns the value of the specified header.
func (r *Response) HeaderValue(key string) string {
	return r.Header.Get(key)
}

// ContentType returns the Content-Type header value.
func (r *Response) ContentType() string {
	return r.HeaderValue("Content-Type")
}

// ContentLength returns the Content-Length header value.
func (r *Response) ContentLength() int64 {
	return r.Response.ContentLength
}
