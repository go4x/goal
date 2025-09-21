package httpx

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/go4x/goal/errorx"
)

// Handler is a function that will be called when the request is successful
type Handler func(resp *http.Response)

// ErrHandler is a function that will be called when the request is failed
type ErrHandler func(err error)

// httpBuilder is a builder for http request
type httpBuilder struct {
	builder
	callback   Handler
	errHandler ErrHandler
}

// NewBuilderClient creates a new httpBuilder with a custom http client,
// the difference between NewBuilder and NewBuilderClient is that NewBuilderClient will use the given http client,
// while NewBuilder will use the default http client with a timeout of 60 seconds
func NewBuilderClient(c *http.Client, url string, param ...any) *httpBuilder {
	h := http.Header{}
	b := &httpBuilder{}
	b.client = c
	b.url = url
	b.params = param
	b.headers = h
	return b
}

// NewBuilder creates a new httpBuilder with the default http client
func NewBuilder(url string, param ...any) *httpBuilder {
	return NewBuilderClient(defaultClient, url, param...)
}

// Client sets the http client for the httpBuilder
func (b *httpBuilder) Client(c *http.Client) *httpBuilder {
	b.client = c
	return b
}

// ContentType sets the request content type for the httpBuilder
func (b *httpBuilder) ContentType(contentType ContentType) *httpBuilder {
	b.contentType = contentType
	return b
}

// Header sets a single request header for the httpBuilder
func (b *httpBuilder) Header(key string, v string) *httpBuilder {
	b.headers.Add(key, v)
	return b
}

// Headers sets multiple request headers for the httpBuilder
func (b *httpBuilder) Headers(headers ...H) *httpBuilder {
	for _, h := range headers {
		for k, v := range h {
			b.headers.Add(k, v)
		}
	}
	return b
}

// Body sets the request body for the httpBuilder
func (b *httpBuilder) Body(body io.Reader) *httpBuilder {
	b.body = body
	return b
}

// BodyStr sets the request body for the httpBuilder from a string
func (b *httpBuilder) BodyStr(body string) *httpBuilder {
	b.body = strings.NewReader(body)
	return b
}

func (b *httpBuilder) getClient() *http.Client {
	if b.client == nil {
		return defaultClient
	}
	return b.client
}

// Request sends the request with the given method
func (b *httpBuilder) Request(method string) {
	b.method = method
	req, err := http.NewRequest(b.method, b.url, b.body)
	if err != nil {
		b.errHandler(err)
		return
	}
	if len(b.headers) > 0 {
		req.Header = b.headers
	}
	resp, err := b.getClient().Do(req)
	if err != nil {
		b.errHandler(err)
		return
	}

	if resp.StatusCode != http.StatusOK {
		err = errors.New(resp.Status)
		b.errHandler(err)
	} else {
		b.callback(resp)
	}
}

// Get sends a GET request
func (b *httpBuilder) Get() {
	b.method = http.MethodGet
	var resp *http.Response
	var err error
	if len(b.headers) > 0 {
		req, err := http.NewRequest(b.method, b.url, nil)
		errorx.Throw(err)
		req.Header = b.headers
		resp, err = b.getClient().Do(req)
	} else {
		resp, err = b.getClient().Get(b.url)
	}

	if err != nil {
		b.errHandler(err)
		return
	}

	if resp.StatusCode != http.StatusOK {
		err = errors.New(resp.Status)
		b.errHandler(err)
	} else {
		b.callback(resp)
	}
}

// Post sends a POST request
func (b *httpBuilder) Post() {
	b.method = http.MethodPost
	var resp *http.Response
	var err error
	if len(b.headers) > 0 {
		req, err := http.NewRequest(b.method, b.url, b.body)
		errorx.Throw(err)

		b.mergeHeaders()

		req.Header = b.headers
		resp, err = b.getClient().Do(req)
	} else {
		resp, err = b.getClient().Post(b.url, string(b.contentType), b.body)
	}

	if err != nil {
		b.errHandler(err)
		return
	}

	if resp.StatusCode != http.StatusOK {
		err = errors.New(resp.Status)
		b.errHandler(err)
	} else {
		b.callback(resp)
	}
}

func (b *httpBuilder) mergeHeaders() *httpBuilder {
	if b.contentType != "" {
		b.headers.Add("Content-Type", string(b.contentType))
		// } else {
		// 	if len(b.headers) == 0 {
		// 		return b
		// 	}
		// 	var ct, ok = b.headers["Content-Type"]
		// 	if ok {
		// 		b.contentType = ContentType(ct[0])
		// 	}
	}
	return b
}

// Success sets the handler for the httpBuilder when the request is successful
func (b *httpBuilder) Success(handler Handler) *httpBuilder {
	b.callback = handler
	return b
}

// Failed sets the handler for the httpBuilder when the request is failed
func (b *httpBuilder) Failed(handler ErrHandler) *httpBuilder {
	b.errHandler = handler
	return b
}

// convenient GET methods

// MustGetString sends a GET request and returns the response body as a string.
// if the request is failed, it will panic
func MustGetString(url string, headers ...H) string {
	return getString(url, func(err error) {
		panic(err)
	}, headers...)
}

func getString(url string, errHandler ErrHandler, headers ...H) string {
	var s string
	NewBuilder(url).Success(func(resp *http.Response) {
		bs, err := io.ReadAll(resp.Body)
		if err != nil {
			errHandler(fmt.Errorf("read response data error: %v", err))
			return
		}
		s = string(bs)
	}).Failed(errHandler).Headers(headers...).Get()
	return s
}

// GetString sends a GET request and returns the response body as a string.
// if the request is failed, it will return an error
func GetString(url string, headers ...H) (string, error) {
	var err error
	s := getString(url, func(e error) { err = e }, headers...)
	return s, err
}

// MustGetBytes sends a GET request and returns the response body as a byte slice.
// if the request is failed, it will panic
func MustGetBytes(url string, headers ...H) []byte {
	return getBytes(url, func(err error) {
		panic(fmt.Sprintf("request failed: %v", err))
	}, headers...)
}

func getBytes(url string, errHandler ErrHandler, headers ...H) []byte {
	var ret []byte
	NewBuilder(url).Success(func(resp *http.Response) {
		bytes, err := io.ReadAll(resp.Body)
		if err != nil {
			errHandler(fmt.Errorf("read reponse data error: %v", err))
			return
		}
		ret = bytes
	}).Failed(errHandler).Headers(headers...).Get()
	return ret
}

// GetBytes sends a GET request and returns the response body as a byte slice.
// if the request is failed, it will return an error
func GetBytes(url string, headers ...H) ([]byte, error) {
	var err error
	bs := getBytes(url, func(e error) { err = e }, headers...)
	return bs, err
}

// MustGet sends a GET request with the give handler, if the request is failed, it will panic.
func MustGet(url string, handler Handler, headers ...H) {
	Get(url, handler, func(err error) {
		panic(err)
	}, headers...)
}

// Get is the raw method to send a GET request with the give handler and error handler.
func Get(url string, handler Handler, errHandler ErrHandler, headers ...H) {
	NewBuilder(url).Success(handler).Failed(errHandler).Headers(headers...).Get()
}

// MustGetJson sends a GET request to the specified URL and unmarshals the response body into the provided object t.
// If the request fails or the response cannot be unmarshaled, it will panic.
// T is a generic type representing the type of the object to unmarshal into.
func MustGetJson[T any](url string, t T, headers ...H) T {
	getJson(url, func(err error) {
		panic(fmt.Sprintf("request failed: %v", err))
	}, t, headers...)
	return t
}

func getJson[T any](url string, errHandler ErrHandler, t T, headers ...H) T {
	NewBuilder(url).Success(func(resp *http.Response) {
		bs, err := io.ReadAll(resp.Body)
		if err != nil {
			errHandler(fmt.Errorf("read response data error: %v", err))
			return
		}
		err = json.Unmarshal(bs, t)
		if err != nil {
			errHandler(fmt.Errorf("unmarshal error: %v", err))
			return
		}
	}).Failed(errHandler).Headers(headers...).Get()
	return t
}

// GetJson sends a GET request to the specified URL and unmarshals the response body into the provided object t.
// It returns the object and an error if the request fails or the response cannot be unmarshaled.
// T is a generic type representing the type of the object to unmarshal into.
func GetJson[T any](url string, t T, headers ...H) (T, error) {
	var err error
	obj := getJson(url, func(e error) { err = e }, t, headers...)
	return obj, err
}

// GetResp sends a GET request to the specified URL with optional headers and returns an *R object
// that wraps the HTTP response and any error encountered during the request.
func GetResp(url string, headers ...H) *R {
	var resp *http.Response
	var err error
	NewBuilder(url).Success(func(r *http.Response) { resp = r }).Failed(func(e error) { err = e }).Headers(headers...).Get()
	return RespErr(resp, err)
}

// convenient POST methods

// MustPostJson sends a POST request with a JSON body to the specified URL and returns the response body as a string.
// If the request fails, it panics with the encountered error.
// url: the target URL to send the POST request to.
// body: the request body, typically a JSON-encoded io.Reader.
// headers: optional headers to include in the request.
func MustPostJson(url string, body io.Reader, headers ...H) string {
	return postJson(url, body, func(err error) {
		panic(err)
	}, headers...)
}

func postJson(url string, body io.Reader, errHandler ErrHandler, headers ...H) string {
	var s string
	NewBuilder(url).ContentType(ContentTypeApplicationJson).Body(body).Success(func(resp *http.Response) {
		bs, err := io.ReadAll(resp.Body)
		if err != nil {
			errHandler(fmt.Errorf("read response data error: %v", err))
			return
		}
		s = string(bs)
	}).Failed(errHandler).Headers(headers...).Post()
	return s
}

// PostJson sends a POST request with a JSON body to the specified URL and returns the response body as a string.
// It returns the string and an error if the request fails or the response cannot be read.
// url: the target URL to send the POST request to.
// body: the request body, typically a JSON-encoded io.Reader.
// headers: optional headers to include in the request.
func PostJson(url string, body io.Reader, headers ...H) (string, error) {
	var err error
	r := postJson(url, body, func(e error) {
		err = e
	}, headers...)
	return r, err
}

// PostJsonr sends a POST request with a JSON body to the specified URL and returns an *R object
// that wraps the HTTP response and any error encountered during the request.
func PostJsonr(url string, body io.Reader, headers ...H) *R {
	var resp *http.Response
	var err error
	NewBuilder(url).ContentType(ContentTypeApplicationJson).Body(body).
		Success(func(r *http.Response) { resp = r }).
		Failed(func(e error) { err = e }).
		Post()
	return RespErr(resp, err)
}

// MustPostForm sends a POST request with a form-urlencoded body to the specified URL and returns the response body as a string.
// If the request fails, it panics with the encountered error.
// url: the target URL to send the POST request to.
// body: the request body, typically a form-urlencoded io.Reader.
// headers: optional headers to include in the request.
func MustPostForm(url string, body io.Reader, headers ...H) string {
	return postForm(url, body, func(err error) {
		panic(err)
	}, headers...)
}

func postForm(url string, body io.Reader, errHandler ErrHandler, headers ...H) string {
	var s string
	NewBuilder(url).ContentType(ContentTypeApplicationFormUrlencoded).Body(body).Success(func(resp *http.Response) {
		bs, err := io.ReadAll(resp.Body)
		if err != nil {
			errHandler(fmt.Errorf("read response data error: %v", err))
			return
		}
		if len(bs) > 0 {
			s = string(bs)
		}
	}).Failed(errHandler).Headers(headers...).Post()
	return s
}

// PostForm sends a POST request with a form-urlencoded body to the specified URL and returns the response body as a string.
// It returns the string and an error if the request fails or the response cannot be read.
// url: the target URL to send the POST request to.
// body: the request body, typically a form-urlencoded io.Reader.
// headers: optional headers to include in the request.
func PostForm(url string, body io.Reader, headers ...H) (string, error) {
	var err error
	s := postForm(url, body, func(e error) {
		err = e
	}, headers...)
	return s, err
}

// PostFormr sends a POST request with a form-urlencoded body to the specified URL and returns an *R object
// that wraps the HTTP response and any error encountered during the request.
func PostFormr(url string, body io.Reader, headers ...H) *R {
	var resp *http.Response
	var err error
	NewBuilder(url).ContentType(ContentTypeApplicationFormUrlencoded).Body(body).
		Success(func(r *http.Response) { resp = r }).
		Failed(func(e error) { err = e }).
		Post()
	return RespErr(resp, err)
}

// MustPostBytes sends a POST request with a raw body to the specified URL and returns the response body as a byte slice.
// If the request fails, it panics with the encountered error.
// url: the target URL to send the POST request to.
// ct: the content type of the request body.
// body: the request body, typically a raw io.Reader.
// headers: optional headers to include in the request.
func MustPostBytes(url string, ct ContentType, body io.Reader, headers ...H) []byte {
	return postBytes(url, ct, body, func(err error) {
		panic(fmt.Sprintf("request failed: %v", err))
	}, headers...)
}

func postBytes(url string, ct ContentType, body io.Reader, errHandler ErrHandler, headers ...H) []byte {
	var ret []byte
	NewBuilder(url).ContentType(ct).Body(body).Success(func(resp *http.Response) {
		bytes, err := io.ReadAll(resp.Body)
		if err != nil {
			errHandler(fmt.Errorf("read reponse data error: %v", err))
			return
		}
		ret = bytes
	}).Failed(errHandler).Headers(headers...).Post()
	return ret
}

// PostBytes sends a POST request with a raw body to the specified URL and returns the response body as a byte slice.
// It returns the byte slice and an error if the request fails or the response cannot be read.
// url: the target URL to send the POST request to.
// ct: the content type of the request body.
// body: the request body, typically a raw io.Reader.
// headers: optional headers to include in the request.
func PostBytes(url string, ct ContentType, body io.Reader, headers ...H) ([]byte, error) {
	var err error
	ret := postBytes(url, ct, body, func(e error) { err = e }, headers...)
	return ret, err
}

// MustPost sends a POST request with a raw body to the specified URL and calls the given handler when the request is successful.
// If the request fails, it panics with the encountered error.
// url: the target URL to send the POST request to.
// ct: the content type of the request body.
// body: the request body, typically a raw io.Reader.
// handler: the handler to call when the request is successful.
func MustPost(url string, ct ContentType, body io.Reader, handler Handler, headers ...H) {
	Post(url, ct, body, handler, func(err error) {
		panic(err)
	}, headers...)
}

// Post is the raw post method, it sends a POST request with the specified content type and body to the given URL.
// handler: function to handle the response when the request is successful.
// errHandler: function to handle errors that occur during the request.
// headers: optional HTTP headers to include in the request.
func Post(url string, ct ContentType, body io.Reader, handler Handler, errHandler ErrHandler, headers ...H) {
	NewBuilder(url).ContentType(ct).Body(body).Success(handler).Failed(errHandler).Headers(headers...).Post()
}

// MustPostJsonObj is a generic method that sends a POST request with a JSON body to the specified URL and unmarshals the response body into the provided object t.
// If the request fails or the response cannot be unmarshaled, it will panic.
// T is a generic type representing the type of the object to unmarshal into.
func MustPostJsonObj[T any](url string, body io.Reader, t T, headers ...H) T {
	postJsonObj(url, body, func(err error) {
		panic(fmt.Sprintf("request failed: %v", err))
	}, t, headers...)
	return t
}

func postJsonObj[T any](url string, body io.Reader, errHandler ErrHandler, t T, headers ...H) T {
	NewBuilder(url).ContentType(ContentTypeApplicationJson).Body(body).Success(func(resp *http.Response) {
		bs, err := io.ReadAll(resp.Body)
		if err != nil {
			errHandler(fmt.Errorf("read response data error: %v", err))
			return
		}
		err = json.Unmarshal(bs, t)
		if err != nil {
			errHandler(fmt.Errorf("unmarshal error: %v", err))
		}
	}).Failed(errHandler).Headers(headers...).Post()
	return t
}

// PostJsonObj is a generic method that sends a POST request with a JSON body to the specified URL and unmarshals the response body into the provided object t.
// It returns the object and an error if the request fails or the response cannot be unmarshaled.
// T is a generic type representing the type of the object to unmarshal into.
func PostJsonObj[T any](url string, body io.Reader, t T, headers ...H) (T, error) {
	var err error
	ret := postJsonObj(url, body, func(e error) { err = e }, t, headers...)
	return ret, err
}
