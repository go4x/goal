package httpx

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// AsyncResult represents the result of an asynchronous HTTP request.
type AsyncResult struct {
	Resp *Response
	Err  error
}

// AsyncClient provides asynchronous HTTP request capabilities.
// It wraps RestClient and provides methods that return channels for non-blocking operations.
type AsyncClient struct {
	*RestClient
}

// NewAsyncClient creates a new AsyncClient with the given http.Client.
// If client is nil, it uses the default client.
func NewAsyncClient(client *http.Client) *AsyncClient {
	if client == nil {
		return &AsyncClient{RestClient: DefaultClient}
	}
	return &AsyncClient{RestClient: NewRestClient(client)}
}

// GetAsync sends an asynchronous GET request and returns a channel that will receive the result.
func (ac *AsyncClient) GetAsync(url string, params url.Values) <-chan AsyncResult {
	resultChan := make(chan AsyncResult, 1)

	go func() {
		resp, err := ac.Get(url, params)
		resultChan <- AsyncResult{Resp: resp, Err: err}
	}()

	return resultChan
}

// GetWithBodyAsync sends an asynchronous GET request with body and returns a channel that will receive the result.
func (ac *AsyncClient) GetWithBodyAsync(url string, body io.Reader) <-chan AsyncResult {
	resultChan := make(chan AsyncResult, 1)

	go func() {
		resp, err := ac.GetWithBody(url, body)
		resultChan <- AsyncResult{Resp: resp, Err: err}
	}()

	return resultChan
}

// GetJsonAsync sends an asynchronous GET request with JSON body and returns a channel that will receive the result.
func (ac *AsyncClient) GetJsonAsync(url string, body io.Reader) <-chan AsyncResult {
	resultChan := make(chan AsyncResult, 1)

	go func() {
		resp, err := ac.GetJson(url, body)
		resultChan <- AsyncResult{Resp: resp, Err: err}
	}()

	return resultChan
}

// GetFormAsync sends an asynchronous GET request with form data and returns a channel that will receive the result.
func (ac *AsyncClient) GetFormAsync(url string, body io.Reader) <-chan AsyncResult {
	resultChan := make(chan AsyncResult, 1)

	go func() {
		resp, err := ac.GetForm(url, body)
		resultChan <- AsyncResult{Resp: resp, Err: err}
	}()

	return resultChan
}

// PostAsync sends an asynchronous POST request and returns a channel that will receive the result.
func (ac *AsyncClient) PostAsync(url string, body io.Reader) <-chan AsyncResult {
	resultChan := make(chan AsyncResult, 1)

	go func() {
		resp, err := ac.Post(url, body)
		resultChan <- AsyncResult{Resp: resp, Err: err}
	}()

	return resultChan
}

// PostJsonAsync sends an asynchronous POST request with JSON content type and returns a channel that will receive the result.
func (ac *AsyncClient) PostJsonAsync(url string, body io.Reader) <-chan AsyncResult {
	resultChan := make(chan AsyncResult, 1)

	go func() {
		resp, err := ac.PostJson(url, body)
		resultChan <- AsyncResult{Resp: resp, Err: err}
	}()

	return resultChan
}

// PostFormAsync sends an asynchronous POST request with form data and returns a channel that will receive the result.
func (ac *AsyncClient) PostFormAsync(url string, body io.Reader) <-chan AsyncResult {
	resultChan := make(chan AsyncResult, 1)

	go func() {
		resp, err := ac.PostForm(url, body)
		resultChan <- AsyncResult{Resp: resp, Err: err}
	}()

	return resultChan
}

// PutAsync sends an asynchronous PUT request and returns a channel that will receive the result.
func (ac *AsyncClient) PutAsync(url string, body io.Reader) <-chan AsyncResult {
	resultChan := make(chan AsyncResult, 1)

	go func() {
		resp, err := ac.Put(url, body)
		resultChan <- AsyncResult{Resp: resp, Err: err}
	}()

	return resultChan
}

// PutJsonAsync sends an asynchronous PUT request with JSON content type and returns a channel that will receive the result.
func (ac *AsyncClient) PutJsonAsync(url string, body io.Reader) <-chan AsyncResult {
	resultChan := make(chan AsyncResult, 1)

	go func() {
		resp, err := ac.PutJson(url, body)
		resultChan <- AsyncResult{Resp: resp, Err: err}
	}()

	return resultChan
}

// PutFormAsync sends an asynchronous PUT request with form data and returns a channel that will receive the result.
func (ac *AsyncClient) PutFormAsync(url string, body io.Reader) <-chan AsyncResult {
	resultChan := make(chan AsyncResult, 1)

	go func() {
		resp, err := ac.PutForm(url, body)
		resultChan <- AsyncResult{Resp: resp, Err: err}
	}()

	return resultChan
}

// PatchAsync sends an asynchronous PATCH request and returns a channel that will receive the result.
func (ac *AsyncClient) PatchAsync(url string, body io.Reader) <-chan AsyncResult {
	resultChan := make(chan AsyncResult, 1)

	go func() {
		resp, err := ac.Patch(url, body)
		resultChan <- AsyncResult{Resp: resp, Err: err}
	}()

	return resultChan
}

// PatchJsonAsync sends an asynchronous PATCH request with JSON content type and returns a channel that will receive the result.
func (ac *AsyncClient) PatchJsonAsync(url string, body io.Reader) <-chan AsyncResult {
	resultChan := make(chan AsyncResult, 1)

	go func() {
		resp, err := ac.PatchJson(url, body)
		resultChan <- AsyncResult{Resp: resp, Err: err}
	}()

	return resultChan
}

// PatchFormAsync sends an asynchronous PATCH request with form data and returns a channel that will receive the result.
func (ac *AsyncClient) PatchFormAsync(url string, body io.Reader) <-chan AsyncResult {
	resultChan := make(chan AsyncResult, 1)

	go func() {
		resp, err := ac.PatchForm(url, body)
		resultChan <- AsyncResult{Resp: resp, Err: err}
	}()

	return resultChan
}

// DeleteAsync sends an asynchronous DELETE request and returns a channel that will receive the result.
func (ac *AsyncClient) DeleteAsync(url string) <-chan AsyncResult {
	resultChan := make(chan AsyncResult, 1)

	go func() {
		resp, err := ac.Delete(url)
		resultChan <- AsyncResult{Resp: resp, Err: err}
	}()

	return resultChan
}

// DeleteJsonAsync sends an asynchronous DELETE request with JSON body and returns a channel that will receive the result.
func (ac *AsyncClient) DeleteJsonAsync(url string, body io.Reader) <-chan AsyncResult {
	resultChan := make(chan AsyncResult, 1)

	go func() {
		resp, err := ac.DeleteJson(url, body)
		resultChan <- AsyncResult{Resp: resp, Err: err}
	}()

	return resultChan
}

// DeleteFormAsync sends an asynchronous DELETE request with form data and returns a channel that will receive the result.
func (ac *AsyncClient) DeleteFormAsync(url string, body io.Reader) <-chan AsyncResult {
	resultChan := make(chan AsyncResult, 1)

	go func() {
		resp, err := ac.DeleteForm(url, body)
		resultChan <- AsyncResult{Resp: resp, Err: err}
	}()

	return resultChan
}

// OptionsAsync sends an asynchronous OPTIONS request and returns a channel that will receive the result.
func (ac *AsyncClient) OptionsAsync(url string, options ...RequestOption) <-chan AsyncResult {
	resultChan := make(chan AsyncResult, 1)

	go func() {
		resp, err := ac.Options(url, options...)
		resultChan <- AsyncResult{Resp: resp, Err: err}
	}()

	return resultChan
}

// BatchAsync sends multiple asynchronous requests and returns a channel that will receive all results.
// The results will be in the same order as the requests.
func (ac *AsyncClient) BatchAsync(requests []Request) <-chan []AsyncResult {
	resultChan := make(chan []AsyncResult, 1)

	go func() {
		results := make([]AsyncResult, len(requests))

		// Use goroutines for each request
		type resultWithIndex struct {
			index  int
			result AsyncResult
		}

		innerResultChan := make(chan resultWithIndex, len(requests))

		// Start all requests concurrently
		for i, req := range requests {
			go func(i int, req Request) {
				resp, err := ac.executeAsyncRequest(req)

				innerResultChan <- resultWithIndex{
					index:  i,
					result: AsyncResult{Resp: resp, Err: err},
				}
			}(i, req)
		}

		// Collect all results
		for i := 0; i < len(requests); i++ {
			result := <-innerResultChan
			results[result.index] = result.result
		}

		resultChan <- results
	}()

	return resultChan
}

// executeAsyncRequest intelligently executes an async request based on method and content type.
func (ac *AsyncClient) executeAsyncRequest(req Request) (*Response, error) {
	// Determine content type
	contentType := req.GetHeader().Get("Content-Type")

	body, err := req.GetBody()
	if err != nil {
		return nil, err
	}

	// Execute request based on method and content type
	switch req.GetMethod() {
	case http.MethodGet:
		if body != nil {
			// GET with body
			switch contentType {
			case ContentTypeApplicationJson:
				return ac.GetJson(req.GetURL().String(), body)
			case ContentTypeApplicationFormUrlencoded:
				return ac.GetForm(req.GetURL().String(), body)
			default:
				return ac.GetWithBody(req.GetURL().String(), body)
			}
		}
		return ac.Get(req.GetURL().String(), req.GetRequest().URL.Query())

	case http.MethodPost:
		if body == nil {
			return ac.Post(req.GetURL().String(), nil)
		}
		switch contentType {
		case ContentTypeApplicationJson:
			return ac.PostJson(req.GetURL().String(), body)
		case ContentTypeApplicationFormUrlencoded:
			return ac.PostForm(req.GetURL().String(), body)
		default:
			return ac.Post(req.GetURL().String(), body)
		}

	case http.MethodPut:
		if body == nil {
			return ac.Put(req.GetURL().String(), nil)
		}
		switch contentType {
		case ContentTypeApplicationJson:
			return ac.PutJson(req.GetURL().String(), body)
		case ContentTypeApplicationFormUrlencoded:
			return ac.PutForm(req.GetURL().String(), body)
		default:
			return ac.Put(req.GetURL().String(), body)
		}

	case http.MethodPatch:
		if body == nil {
			return ac.Patch(req.GetURL().String(), nil)
		}
		switch contentType {
		case ContentTypeApplicationJson:
			return ac.PatchJson(req.GetURL().String(), body)
		case ContentTypeApplicationFormUrlencoded:
			return ac.PatchForm(req.GetURL().String(), body)
		default:
			return ac.Patch(req.GetURL().String(), body)
		}

	case http.MethodDelete:
		if body == nil {
			return ac.Delete(req.GetURL().String())
		}
		switch contentType {
		case ContentTypeApplicationJson:
			return ac.DeleteJson(req.GetURL().String(), body)
		case ContentTypeApplicationFormUrlencoded:
			return ac.DeleteForm(req.GetURL().String(), body)
		default:
			// For DELETE with body, we need to use a custom approach
			return ac.deleteWithBody(req.GetURL().String(), body)
		}

	case http.MethodOptions:
		return ac.Options(req.GetURL().String())

	default:
		return nil, fmt.Errorf("unsupported HTTP method: %s", req.GetMethod())
	}
}

// deleteWithBody sends a DELETE request with body using the underlying HTTP client.
func (ac *AsyncClient) deleteWithBody(url string, body io.Reader) (*Response, error) {
	req, err := http.NewRequest(http.MethodDelete, url, body)
	if err != nil {
		return nil, err
	}

	resp, err := ac.Do(req)
	if err != nil {
		return nil, err
	}

	return NewResponse(resp), nil
}

// WithContextAsync sends an asynchronous request with context support.
// This allows for timeout and cancellation control.
func (ac *AsyncClient) WithContextAsync(ctx context.Context, method, url string, params url.Values, body io.Reader) <-chan AsyncResult {
	resultChan := make(chan AsyncResult, 1)

	go func() {
		// Create a request with context
		req, err := NewRequestWithContext(ctx, method, url, body)
		if err != nil {
			resultChan <- AsyncResult{Resp: nil, Err: err}
			return
		}

		// Send the request
		resp, err := ac.Do(req.GetRequest())
		if err != nil {
			resultChan <- AsyncResult{Resp: nil, Err: err}
			return
		}

		response := NewResponse(resp)
		resultChan <- AsyncResult{Resp: response, Err: nil}
	}()

	return resultChan
}

// WaitForAll waits for multiple async requests to complete and returns all results.
// This is a convenience function that collects results from multiple channels.
func WaitForAll(resultChans []<-chan AsyncResult) []AsyncResult {
	results := make([]AsyncResult, len(resultChans))

	for i, resultChan := range resultChans {
		results[i] = <-resultChan
	}

	return results
}

// WaitForFirst waits for the first async request to complete and returns its result.
// Other requests will continue in the background.
func WaitForFirst(resultChans []<-chan AsyncResult) AsyncResult {
	// Create a channel to receive the first result
	firstResult := make(chan AsyncResult, 1)

	// Start goroutines for each channel
	for _, resultChan := range resultChans {
		go func(ch <-chan AsyncResult) {
			result := <-ch
			select {
			case firstResult <- result:
			default:
				// Channel already has a result, ignore
			}
		}(resultChan)
	}

	return <-firstResult
}
