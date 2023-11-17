package gonduit

import (
	"github.com/thought-machine/gonduit/requests"
	"github.com/thought-machine/gonduit/responses"
)

// PasteCreate calls the paste.create endpoint.
// Deprecated: This method is frozen and will eventually be deprecated. New code should use "paste.edit" instead.
func (c *Conn) PasteCreate(
	req *requests.PasteCreateRequest,
) (responses.PasteCreateResponse, error) {
	var res responses.PasteCreateResponse

	if err := c.Call("paste.create", &req, &res); err != nil {
		return nil, err
	}

	return res, nil
}

// PasteQuery calls the paste.query endpoint.
// Deprecated: This method is frozen and will eventually be deprecated. New code should use "paste.search" instead.
func (c *Conn) PasteQuery(
	req *requests.PasteQueryRequest,
) (responses.PasteQueryResponse, error) {
	var res responses.PasteQueryResponse

	if err := c.Call("paste.query", &req, &res); err != nil {
		return nil, err
	}

	return res, nil
}

// PasteSearch performs a call to paste.search
func (c *Conn) PasteSearch(req requests.SearchRequest) (*responses.SearchResponse, error) {
	var res responses.SearchResponse
	if err := c.Call("paste.search", &req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
