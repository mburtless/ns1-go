package ns1

import (
	"net/http"
	"net/url"
)

// Request is the request to be made
type Request struct {
	Operation    *Operation
	Body         interface{}
	HTTPRequest  *http.Request
	HTTPResponse *http.Response
	Error        error
	Params       interface{}
	Data         interface{}
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
func New(cfg *Config, operation *Operation, params interface{}, data interface{}) *Request {
	// TODO make copy of config?
	method := operation.HTTPMethod
	if method == "" {
		method = "GET"
	}
	rel, err := url.Parse(operation.HTTPPath)
	uri := cfg.endpoint.ResolveReference(rel)

	httpReq, err := http.NewRequest(method, uri.String(), nil)

	r := &Request{
		Operation:   operation,
		Body:        nil,
		HTTPRequest: httpReq,
		Error:       err,
		Params:      params,
		Data:        data,
	}
	return r
}
