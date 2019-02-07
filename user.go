package gonduit

import (
	"github.com/thought-machine/gonduit/requests"
	"github.com/thought-machine/gonduit/responses"
)

// UserQuery performs a call to user.query.
func (c *Conn) UserQuery(
	req requests.UserQueryRequest,
) (*responses.UserQueryResponse, error) {
	var res responses.UserQueryResponse

	if err := c.Call("user.query", &req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// UserSearch performs a call to user.search
func (c *Conn) UserSearch(req requests.SearchRequest) (*responses.SearchResponse, error) {
	var res responses.SearchResponse
	if err := c.Call("user.search", &req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
