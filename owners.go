package gonduit

import (
	"github.com/thought-machine/gonduit/entities"
	"github.com/thought-machine/gonduit/requests"
	"github.com/thought-machine/gonduit/responses"
)

// OwnersSearch performs a call to owners.search
func (c *Conn) OwnersSearch(req requests.SearchRequest) (*responses.SearchResponse, error) {
	var res responses.SearchResponse
	if err := c.Call("owners.search", &req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
