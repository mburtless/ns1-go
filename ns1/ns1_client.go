// Code generated by go-swagger; DO NOT EDIT.

package ns1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"net/url"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	rtclient "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/mburtless/ns1-go-v3/ns1/zones"
)

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "api.nsone.net"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath   string = "/v1/"
	DefaultAuthHeader string = "X-NSONE-Key"
)

// DefaultSchemes are the default schemes found in Meta (info) section of spec file
var DefaultSchemes = []string{"https"}

type Config struct {
	// URL is the base URL of the upstream server
	URL *url.URL
	// Transport is an inner transport for the client
	Transport http.RoundTripper
	// AuthInfo is for authentication
	//AuthInfo  runtime.ClientAuthInfoWriter
	// NS1 api key (value for http request header 'X-NSONE-Key').
	APIKey string
}

// New creates a new ns1 HTTP client.
func New(c Config) (*Ns1, error) {
	var (
		host     = DefaultHost
		basePath = DefaultBasePath
		schemes  = DefaultSchemes
	)

	if c.URL != nil {
		host = c.URL.Host
		basePath = c.URL.Path
		schemes = []string{c.URL.Scheme}
	}

	transport := rtclient.New(host, basePath, schemes)
	if c.Transport != nil {
		transport.Transport = c.Transport
	}

	// Create authInfo
	if c.APIKey == "" {
		return &Ns1{}, errors.Required("APIKey", "Config")
	}
	authInfo := rtclient.APIKeyAuth(DefaultAuthHeader, "header", c.APIKey)

	cli := new(Ns1)
	cli.Transport = transport
	cli.Zones = zones.New(transport, strfmt.Default, authInfo)
	return cli, nil
}

// Ns1 is a client for ns1
type Ns1 struct {
	Zones     *zones.Client
	Transport runtime.ClientTransport
}
