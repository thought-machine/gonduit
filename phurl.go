package gonduit

import (
	"github.com/thought-machine/gonduit/requests"
	"github.com/thought-machine/gonduit/responses"
)

// PhurlsSearch performs a call to phurls.search
func (c *Conn) PhurlsSearch(req requests.SearchRequest) (*responses.SearchResponse, error) {
	var res responses.SearchResponse
	if err := c.Call("phurls.search", &req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
