package gonduit

import (
	"github.com/samwestmoreland/gonduit/requests"
	"github.com/samwestmoreland/gonduit/responses"
)

// OwnersSearch performs a call to owners.search
func (c *Conn) OwnersSearch(req requests.SearchRequest) (*responses.SearchResponse, error) {
	var res responses.SearchResponse
	if err := c.Call("owners.search", &req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
