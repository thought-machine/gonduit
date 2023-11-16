package gonduit

import (
	"github.com/samwestmoreland/gonduit/requests"
	"github.com/samwestmoreland/gonduit/responses"
)

// UserQuery performs a call to user.query.
// Deprecated: This method is frozen and will eventually be deprecated. New code should use "user.search" instead.
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

// UserEdit performs a call to user.edit
func (c *Conn) UserEdit(req requests.EditRequest) (*responses.EditResponse, error) {
	var res responses.EditResponse
	if err := c.Call("user.edit", &req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
