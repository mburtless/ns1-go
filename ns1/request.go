package ns1

import (
	"encoding/json"
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

// NeqRequest returns a new Request pointer for the NS1 API
// operation and parameters
//
// Params is any value of input parameters to be the request payload
// Data is pointer value to an object which the request's response
// payload will be deserialized to.
func New(cfg Config, httpClient *http.Client, operation *Operation, params interface{}, data interface{}) *Request {
	// TODO make copy of config?
	method := operation.HTTPMethod
	if method == "" {
		method = "GET"
	}
	rel, err := url.Parse(operation.HTTPPath)
	uri := cfg.endpoint.ResolveReference(rel)

	// Handle body

	httpReq, err := http.NewRequest(method, uri.String(), nil)

	// Handle headers
	httpReq.Header.Add(headerAuth, cfg.apiKey)
	httpReq.Header.Add("User-Agent", cfg.userAgent)

	// init remaining fields
	r := &Request{
		Config:      cfg,
		Operation:   operation,
		Body:        nil,
		HTTPRequest: httpReq,
		HTTPClient:  httpClient,
		Error:       err,
		Params:      params,
		Data:        data,
	}
	return r
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
