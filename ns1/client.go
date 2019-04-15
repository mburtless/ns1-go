package ns1

import (
	"crypto/tls"
	"net/http"
)

// Client manages communication with the NS1 Rest API.
type Client struct {
	// Configuration of the client
	config *Config

	// The HTTP client to use when sending requests. Defaults to
	// `http.DefaultClient`.
	httpClient *http.Client
	// From the excellent github-go client.
	//common service // Reuse a single struct instead of allocating one for each service on the heap.

	// Services used for communicating with different components of the NS1 API.
	/*APIKeys       *APIKeysService
	DataFeeds     *DataFeedsService
	DataSources   *DataSourcesService
	Jobs          *JobsService
	Notifications *NotificationsService
	Records       *RecordsService
	Settings      *SettingsService
	Teams         *TeamsService
	Users         *UsersService
	Warnings      *WarningsService*/
	Zones ZonesService
}

// NewClient will return a pointer to a new initialized Client
// TODO: Allow users to pass their own httpclient as param
func NewClient(cfg *Config) *Client {
	c := &Client{
		config: cfg,
	}
	if cfg.ignoreSSL {
		c.httpClient = newInsecureHTTPClient()
	} else {
		c.httpClient = http.DefaultClient
	}

	c.Zones = &ZonesServiceOp{client: c}
	return c
}

// newInsecureHTTPClient returns a http.Client with TLS verification disabled
func newInsecureHTTPClient() *http.Client {
	defaultTransport := http.DefaultTransport.(*http.Transport)
	tr := &http.Transport{
		Proxy:                 defaultTransport.Proxy,
		DialContext:           defaultTransport.DialContext,
		MaxIdleConns:          defaultTransport.MaxIdleConns,
		IdleConnTimeout:       defaultTransport.IdleConnTimeout,
		ExpectContinueTimeout: defaultTransport.ExpectContinueTimeout,
		TLSHandshakeTimeout:   defaultTransport.TLSHandshakeTimeout,
		TLSClientConfig:       &tls.Config{InsecureSkipVerify: true},
	}
	return &http.Client{Transport: tr}
}
