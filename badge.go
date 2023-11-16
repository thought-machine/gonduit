package gonduit

import (
	"github.com/samwestmoreland/gonduit/requests"
	"github.com/samwestmoreland/gonduit/responses"
)

// BadgeSearch performs a call to badge.search
func (c *Conn) BadgeSearch(req requests.SearchRequest) (*responses.SearchResponse, error) {
	var res responses.SearchResponse
	if err := c.Call("badge.search", &req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *Conn) BadgeEdit(req requests.EditRequest) (*responses.EditResponse, error) {
	var res responses.EditResponse
	if err := c.Call("badge.edit", &req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
