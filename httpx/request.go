package httpx

import (
	"context"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
)

// RequestOption is a function type that configures a Request.
type RequestOption func(Request)

// Request defines the interface for HTTP request operations.
// It provides methods to access and modify HTTP request properties.
type Request interface {
	GetURL() *url.URL
	GetMethod() string
	GetBody() (io.ReadCloser, error)
	GetHeader() http.Header
	GetContext() context.Context
	GetForm() url.Values
	GetQuery() url.Values
	GetProto() string
	GetProtoMajor() int
	GetProtoMinor() int
	GetMultipartForm() (*multipart.Form, error)
	GetRequest() *http.Request
	WithURL(url *url.URL) Request
	WithMethod(method string) Request
	WithBody(body io.ReadCloser) Request
	WithHeader(key string, value string) Request
	WithContext(ctx context.Context) Request
	ParseForm() error
}

// request is the default implementation of the Request interface.
// It wraps the standard http.Request to provide additional functionality.
type request struct {
	*http.Request
}

// GetURL returns the request URL.
func (req *request) GetURL() *url.URL {
	return req.URL
}

// GetMethod returns the HTTP method of the request.
func (req *request) GetMethod() string {
	return req.Method
}

// GetBody returns the request body as an io.ReadCloser.
func (req *request) GetBody() (io.ReadCloser, error) {
	return req.Body, nil
}

// GetHeader returns the request headers.
func (req *request) GetHeader() http.Header {
	return req.Header
}

// GetContext returns the request context.
func (req *request) GetContext() context.Context {
	return req.Context()
}

// GetForm returns the parsed form data.
func (req *request) GetForm() url.Values {
	return req.Form
}

// GetQuery returns the query parameters.
func (req *request) GetQuery() url.Values {
	return req.URL.Query()
}

// GetProto returns the HTTP protocol version.
func (req *request) GetProto() string {
	return req.Proto
}

// GetProtoMajor returns the HTTP protocol major version.
func (req *request) GetProtoMajor() int {
	return req.ProtoMajor
}

// GetProtoMinor returns the HTTP protocol minor version.
func (req *request) GetProtoMinor() int {
	return req.ProtoMinor
}

// GetMultipartForm returns the parsed multipart form data.
func (req *request) GetMultipartForm() (*multipart.Form, error) {
	return req.MultipartForm, nil
}

// GetRequest returns the underlying http.Request.
func (req *request) GetRequest() *http.Request {
	return req.Request
}

// WithURL sets the request URL and returns the modified request.
func (req *request) WithURL(url *url.URL) Request {
	req.URL = url
	return req
}

// WithMethod sets the HTTP method and returns the modified request.
func (req *request) WithMethod(method string) Request {
	req.Method = method
	return req
}

// WithBody sets the request body and returns the modified request.
func (req *request) WithBody(body io.ReadCloser) Request {
	req.Body = body
	return req
}

// WithHeader sets a header key-value pair and returns the modified request.
func (req *request) WithHeader(key string, value string) Request {
	req.Header.Set(key, value)
	return req
}

// WithContext sets the request context and returns the modified request.
func (req *request) WithContext(ctx context.Context) Request {
	req.Request = req.Request.WithContext(ctx)
	return req
}

// ParseForm parses the request body as form data.
func (req *request) ParseForm() error {
	return req.Request.ParseForm()
}

// NewRequest creates a new Request with the given method, URL, and body.
// It applies the provided options to configure the request.
func NewRequest(method string, url string, body io.Reader, options ...RequestOption) (Request, error) {
	hreq, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	req := &request{Request: hreq}
	for _, option := range options {
		option(req)
	}
	return req, nil
}

// MustNewRequest creates a new Request and panics if there's an error.
// This is useful for testing or when you're certain the request creation will succeed.
func MustNewRequest(method string, url string, body io.Reader, options ...RequestOption) Request {
	req, err := NewRequest(method, url, body, options...)
	if err != nil {
		panic(err)
	}
	return req
}

// NewRequestWithContext creates a new Request with the given context, method, URL, and body.
// It applies the provided options to configure the request.
func NewRequestWithContext(ctx context.Context, method string, url string, body io.Reader, options ...RequestOption) (Request, error) {
	hreq, err := http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		return nil, err
	}
	req := &request{Request: hreq}
	for _, option := range options {
		option(req)
	}
	return req, nil
}

// MustNewRequestWithContext creates a new Request with context and panics if there's an error.
// This is useful for testing or when you're certain the request creation will succeed.
func MustNewRequestWithContext(ctx context.Context, method string, url string, body io.Reader, options ...RequestOption) Request {
	req, err := NewRequestWithContext(ctx, method, url, body, options...)
	if err != nil {
		panic(err)
	}
	return req
}

// WithContentType returns a RequestOption that sets the Content-Type header.
func WithContentType(contentType string) RequestOption {
	return func(req Request) {
		req.WithHeader("Content-Type", contentType)
	}
}

// WithAuthorization returns a RequestOption that sets the Authorization header.
func WithAuthorization(authorization string) RequestOption {
	return func(req Request) {
		req.WithHeader("Authorization", authorization)
	}
}

// WithHeader returns a RequestOption that sets a header key-value pair.
func WithHeader(key string, value string) RequestOption {
	return func(req Request) {
		req.WithHeader(key, value)
	}
}

// WithContext returns a RequestOption that sets the request context.
func WithContext(ctx context.Context) RequestOption {
	return func(req Request) {
		req.WithContext(ctx)
	}
}
