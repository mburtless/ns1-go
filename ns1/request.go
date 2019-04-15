package ns1

// Request is the request to be made
type Request struct {
	Operation *Operation
	Body      interface{}
}

// Operation is the API operation to be made
type Operation struct {
	Name       string
	HTTPMethod string
	HTTPPath   string
	//*Paginator
}
