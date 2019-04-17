package ns1

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	mux    *http.ServeMux
	ctx    = context.TODO()
	client *Client
	server *httptest.Server
)

func setup() {
	// create server
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)
	//url, _ := url.Parse(server.URL)

	//create config
	apiKey := "testkey"
	ci := &NewConfigInput{
		APIKey:   apiKey,
		Endpoint: fmt.Sprintf("%s/v1/", server.URL),
	}
	config, _ := NewConfig(ci)

	client = NewClient(config)
}

func teardown() {
	server.Close()
}

func testMethod(t *testing.T, r *http.Request, expected string) {
	if expected != r.Method {
		t.Errorf("Request method = %v, expected %v", r.Method, expected)
	}
}

func TestNewInsecureHTTPClient(t *testing.T) {
	c := newInsecureHTTPClient()

	tr := c.Transport.(*http.Transport)
	assert.True(t, tr.TLSClientConfig.InsecureSkipVerify)

	/*defaultTransport := http.DefaultTransport.(*http.Transport)
	tr.TLSClientConfig = nil
	assert.Equal(t, defaultTransport, tr)*/
}

func TestNewClient(t *testing.T) {
	testAPIKey := "testkey"
	input := &NewConfigInput{
		APIKey: testAPIKey,
	}
	config, _ := NewConfig(input)
	c := NewClient(config)
	assert.Equal(t, testAPIKey, c.config.apiKey)

	// Test creating client with SSL disabled
	input.IgnoreSSL = true
	config, _ = NewConfig(input)
	c = NewClient(config)
	tr := c.httpClient.Transport.(*http.Transport)
	assert.True(t, tr.TLSClientConfig.InsecureSkipVerify)
}
