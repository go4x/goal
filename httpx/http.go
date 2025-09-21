package httpx

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"time"

	"github.com/go4x/goal/errorx"
)

// defaultClient is the default http client with a timeout of 60 seconds
var defaultClient = &http.Client{
	Timeout: time.Second * 60,
}

// H is short for http header, it is a map of string key and string value
type H map[string]string

// builder is a builder for http request
type builder struct {
	client      *http.Client
	url         string
	params      []any
	method      string
	contentType ContentType
	headers     http.Header
	body        io.Reader
}

// R is a wrapper struct for http.Response, providing additional fields for error handling and response body caching.
type R struct {
	*http.Response

	err  error  // Stores any error encountered during the HTTP request or response processing
	body []byte // Caches the response body bytes after reading
	read bool   // Indicates whether the response body has been read
}

// Resp creates a new R with the given http.Response
func Resp(r *http.Response) *R {
	return &R{Response: r}
}

// RespErr creates a new R with the given http.Response and error
func RespErr(r *http.Response, err error) *R {
	return &R{Response: r, err: err}
}

// ok checks if the response is successful(2xx)
func (r *R) ok() bool {
	if !(r.StatusCode >= http.StatusOK && r.StatusCode < http.StatusMultipleChoices) {
		r.wrapErr(fmt.Errorf("%s", r.Status))
		return false
	}
	return true
}

// wrapErr wraps the error with the response error
func (r *R) wrapErr(err error) {
	e := err
	if r.err != nil {
		e = errorx.Wrap(err)
	}
	r.err = e
}

// readAll reads the whole body of the response
func (r *R) readAll() *R {
	if !r.read {
		if r.Body != nil {
			if bs, err := io.ReadAll(r.Body); err != nil {
				r.wrapErr(fmt.Errorf("read body error: %v", err))
				r.body = []byte{}
			} else {
				r.read = true
				r.body = bs
			}
		}
	}
	return r
}

// Clone clones the R
func (r *R) Clone() *R {
	bodyCopy := make([]byte, len(r.readAll().body))
	copy(bodyCopy, r.readAll().body)
	nr := RespErr(r.Response, r.err)
	nr.body = bodyCopy
	nr.read = true // body has been copied, so set read to true
	return nr
}

// Err returns the error of the R
func (r *R) Err() error {
	return r.err
}

// Str returns the body of the R as a string
func (r *R) Str() string {
	if r.ok() {
		return string(r.readAll().body)
	}
	return ""
}

// Bytes returns the body of the R as a byte slice
func (r *R) Bytes() []byte {
	return r.readAll().body
}

// JsonObj unmarshals the body of the R as a json object
func (r *R) JsonObj(v any) any {
	if reflect.TypeOf(v).Kind() != reflect.Pointer {
		panic(fmt.Errorf("param v should be a pointer"))
	}
	if r.ok() {
		if err := json.Unmarshal(r.readAll().body, v); err != nil {
			r.wrapErr(fmt.Errorf("unmarshal json error: %v", err))
		}
	}
	return v
}
