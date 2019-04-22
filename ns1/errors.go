package ns1

import (
	"errors"
	"fmt"
	"net/http"
)

var (
	ErrAPIKeyReq = errors.New("apiKey required")
)

// Response wraps stdlib http response.
/*type Response struct {
	*http.Response
}*/

// Error contains all http responses outside the 2xx range.
type RespError struct {
	// HTTP response that caused this error
	Resp *http.Response

	// Error message
	Message string
}

// Satisfy std lib error interface.
func (re *RespError) Error() string {
	return fmt.Sprintf("%v %v: %d %v", re.Resp.Request.Method, re.Resp.Request.URL, re.Resp.StatusCode, re.Message)
}
