package ns1

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

// Request is the request to be made
type Request struct {
	Config       Config
	Operation    *Operation
	Body         interface{}
	HTTPRequest  *http.Request
	HTTPResponse *http.Response
	HTTPClient   *http.Client
	Error        error
	Params       interface{}
	Data         interface{}
	//TODO: Add fields for rate limiting here or to client
}

// Operation is the API operation to be made
type Operation struct {
	Name       string
	HTTPMethod string
	HTTPPath   string
	//*Paginator
}

// New returns a new Request pointer for the NS1 API
// operation and parameters
//
// Params is any value of input parameters to be the request payload
// This will be ignored on any GET or DELETE request
// Data is pointer value to an object which the request's response
// payload will be deserialized to.
func New(cfg Config, httpClient *http.Client, operation *Operation, params interface{}, data interface{}) *Request {
	// init fields
	r := &Request{
		Config:     cfg,
		Operation:  operation,
		Body:       nil,
		HTTPClient: httpClient,
		Params:     params,
		Data:       data,
	}

	// TODO make copy of config?
	method := operation.HTTPMethod
	if method == "" {
		method = "GET"
	}
	rel, err := url.Parse(operation.HTTPPath)
	if err != nil {
		r.Error = err
	}
	uri := cfg.endpoint.ResolveReference(rel)

	// Handle body
	r.Body = new(bytes.Buffer)
	if method != "GET" && method != "DELETE" {
		r.SetBufferBody()
	}
	httpReq, err := http.NewRequest(method, uri.String(), r.Body.(*bytes.Buffer))
	if err != nil {
		r.Error = err
	}
	// Handle headers
	httpReq.Header.Add(headerAuth, cfg.apiKey)
	httpReq.Header.Add("User-Agent", cfg.userAgent)
	r.HTTPRequest = httpReq
	return r
}

// SetBufferBody marshals the contents of the Params field, if present, to the Body field.
// If Params field is empty, no action is taken
// If an error is encournterd during marshaling, it is saved to the Error field
func (r *Request) SetBufferBody() {
	if r.Params != nil {
		err := json.NewEncoder(r.Body.(*bytes.Buffer)).Encode(r.Params)
		if err != nil {
			r.Error = err
		}
	}
}

// CheckResponse handles parsing of rest api errors. Returns nil if no error.
func (r *Request) CheckResponse() error {
	if c := r.HTTPResponse.StatusCode; c >= 200 && c <= 299 {
		return nil
	}

	restErr := &RespError{Resp: r.HTTPResponse}

	b, err := ioutil.ReadAll(r.HTTPResponse.Body)
	if err != nil {
		return err
	}
	if len(b) == 0 {
		return restErr
	}

	err = json.Unmarshal(b, restErr)
	if err != nil {
		return err
	}

	return restErr
}

// SendToData wraps Send to send the request.
// The API response is JSON decoded and stored in the Request Data field
func (r *Request) SendToData() error {
	err := r.Send(r.Data)
	if err != nil {
		return err
	}
	return nil
}

// SendToPointer wraps Send to send the request.
// The API response is JSON decoded and stored in the value pointed to by v
// Used when JSON must be unmarshaled to member of Request Data field
func (r *Request) SendToPointer(v interface{}) error {
	err := r.Send(v)
	if err != nil {
		return err
	}
	return nil
}

// Send will send the request returning error if errors are encountered.
func (r *Request) Send(v interface{}) error {
	// TODO: Add context support
	resp, err := r.HTTPClient.Do(r.HTTPRequest)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	r.HTTPResponse = resp
	err = r.CheckResponse()
	if err != nil {
		return err
	}
	// TODO: Parse rate headers
	// validate response
	// Parse to r.data?
	if r.Data != nil {
		// Try to unmarshal body into given type using streaming decoder
		// if err := json.NewDecoder(resp.Body).Decode(&r.Data.(*ListZonesOutput).Zones); err != nil {
		if err := json.NewDecoder(resp.Body).Decode(v); err != nil {
			return err
		}
	}

	return nil
}
