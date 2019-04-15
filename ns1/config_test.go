package ns1

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewConfig(t *testing.T) {
	endpoint, _ := url.Parse(defaultEndpoint)
	expectedConfig := &Config{
		apiKey:    "testkey",
		endpoint:  endpoint,
		userAgent: defaultUserAgent,
		ignoreSSL: defaultIgnoreSSL,
	}
	input := &NewConfigInput{
		APIKey: "testkey",
	}

	c, err := NewConfig(input)
	require.NoError(t, err)
	validateConfig(t, c, expectedConfig)

	// Test creating config with custom endpoint
	customEndpoint := "https://api.foo.com/v1/"
	input.Endpoint = customEndpoint
	expectedConfig.endpoint, _ = url.Parse(customEndpoint)
	c, err = NewConfig(input)
	require.NoError(t, err)
	validateConfig(t, c, expectedConfig)

	// Creating config without apiKey should raise error
	c, err = NewConfig(nil)
	require.Error(t, err)
}

func validateConfig(t *testing.T, config *Config, expected *Config) {
	assert.Equal(t, expected, config)
}
