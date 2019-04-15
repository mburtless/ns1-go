package ns1

import (
	"net/url"
)

const (
	clientVersion    = "3.0.0"
	defaultEndpoint  = "https://api.nsone.net/v1/"
	defaultUserAgent = "go-ns1/" + clientVersion
	defaultIgnoreSSL = false

	headerAuth          = "X-NSONE-Key"
	headerRateLimit     = "X-Ratelimit-Limit"
	headerRateRemaining = "X-Ratelimit-Remaining"
	headerRatePeriod    = "X-Ratelimit-Period"
	ns1APIKeyEnvVar     = "NS1_APIKEY"
	ns1EndpointEnvVar   = "NS1_ENDPOINT"
)

// Config provides configuration for client
// TODO: don't export to prevent use without constructor?
type Config struct {
	// NS1 api key (value for http request header 'X-NSONE-Key').
	apiKey string

	// NS1 rest endpoint
	endpoint *url.URL

	// NS1 go rest user agent (value for http request header 'User-Agent').
	userAgent string

	// Whether to ignore ssl validation
	ignoreSSL bool

	// TODO: Add retryer or RateLimitFunc field
}

// NewConfigInput provides parameters for NewConfig
type NewConfigInput struct {
	// NS1 api key (value for http request header 'X-NSONE-Key').
	APIKey string
	// NS1 rest endpoint, overrides default if given.
	Endpoint string
	// NS1 go rest user agent (value for http request header 'User-Agent').
	UserAgent string
	// Whether to ignore ssl validation
	IgnoreSSL bool
}

// NewConfig returns pointer to a new Config
func NewConfig(input *NewConfigInput) (*Config, error) {
	if input == nil {
		input = &NewConfigInput{}
	}

	if input.APIKey == "" {
		return nil, ErrAPIKeyReq
	}

	c := &Config{
		userAgent: defaultUserAgent,
		apiKey:    input.APIKey,
		ignoreSSL: defaultIgnoreSSL,
	}

	var err error
	if input.Endpoint == "" {
		c.endpoint, err = url.Parse(defaultEndpoint)
	} else {
		c.endpoint, err = url.Parse(input.Endpoint)
	}
	if err != nil {
		return nil, err
	}

	if input.UserAgent != "" {
		c.userAgent = input.UserAgent
	}

	if input.IgnoreSSL {
		c.ignoreSSL = input.IgnoreSSL
	}

	return c, nil
}
