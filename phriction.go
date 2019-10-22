package gonduit

import (
	"github.com/thought-machine/gonduit/requests"
	"github.com/thought-machine/gonduit/responses"
)

// PhrictionInfo performs a call to phriction.info
// Deprecated: This method is frozen and will eventually be deprecated. New code should use "phriction.document.search" instead.
func (c *Conn) PhrictionInfo(
	req requests.PhrictionInfoRequest,
) (*responses.PhrictionInfoResponse, error) {
	var res responses.PhrictionInfoResponse

	if err := c.Call("phriction.info", &req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// PhrictionContentSearch performs a call to phriction.content.search
func (c *Conn) PhrictionContentSearch(req requests.SearchRequest) (*responses.SearchResponse, error) {
	var res responses.SearchResponse
	if err := c.Call("phriction.content.search", &req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// PhrictionDocumentSearch performs a call to phriction.document.search
func (c *Conn) PhrictionDocumentSearch(req requests.SearchRequest) (*responses.SearchResponse, error) {
	var res responses.SearchResponse
	if err := c.Call("phriction.document.search", &req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
