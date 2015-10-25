package core

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strings"

	"github.com/etcinit/gonduit/requests"
)

// MakeRequest creates a new requests to the conduit API.
func MakeRequest(endpointURL string, params interface{}) (*http.Request, error) {
	// First, we begin by building the request content, which will be encoded as
	// a urlencoded form.
	form, err := prepareForm(params)
	if err != nil {
		return nil, err
	}

	// Next, we begin building the HTTP request.
	req, err := http.NewRequest(
		"POST",
		endpointURL,
		strings.NewReader(form.Encode()),
	)
	if err != nil {
		return nil, err
	}

	// Finally, we set some headers.
	setHeaders(req)

	return req, nil
}

func setHeaders(req *http.Request) {
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
}

func prepareForm(params interface{}) (url.Values, error) {
	form := url.Values{}
	form.Add("output", "json")

	if params != nil {
		b, err := json.Marshal(params)
		if err != nil {
			return nil, err
		}

		form.Add("params", string(b))

		handleConnectRequest(&form, params)
	}

	return form, nil
}

func handleConnectRequest(form *url.Values, params interface{}) {
	_, isConduitConnect := params.(*requests.ConduitConnectRequest)
	if isConduitConnect {
		form.Add("__conduit__", "true")
	}
}
