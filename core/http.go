package core

import (
	"crypto/tls"
	"net/http"
	"time"
)

// ClientOptions are options that can be set on the HTTP client.
type ClientOptions struct {
	APIToken string

	Cert       string
	CertUser   string
	SessionKey string

	InsecureSkipVerify bool

	RoundTripper http.RoundTripper
	Timeout      time.Duration
}

// makeHttpClient creates a new HTTP client for making API requests.
func makeHTTPClient(options *ClientOptions) *http.Client {
	// Default to a 30 second timeout
	timeout := 30 * time.Second
	if options.Timeout != 0 {
		timeout = options.Timeout
	}

	var transport http.RoundTripper

	transport = &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: options.InsecureSkipVerify,
		},
	}
	if options.RoundTripper != nil {
		transport = options.RoundTripper
	}

	return &http.Client{
		Transport: transport,
		Timeout: timeout,
	}
}
