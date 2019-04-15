package ns1

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewInsecureHTTPClient(t *testing.T) {
	client := newInsecureHTTPClient()

	tr := client.Transport.(*http.Transport)
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
	client := NewClient(config)
	assert.Equal(t, testAPIKey, client.config.apiKey)

	// Test creating client with SSL disabled
	input.IgnoreSSL = true
	config, _ = NewConfig(input)
	client = NewClient(config)
	tr := client.httpClient.Transport.(*http.Transport)
	assert.True(t, tr.TLSClientConfig.InsecureSkipVerify)
}
