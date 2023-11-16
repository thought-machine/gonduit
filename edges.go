package gonduit

import (
	"github.com/samwestmoreland/gonduit/requests"
	"github.com/samwestmoreland/gonduit/responses"
)

// EdgeSearch performs a call to edge.search endpoint to find object associations.
func (c *Conn) EdgeSearch(req requests.EdgeSearchRequest) (*responses.EdgeSearchResponse, error) {
	var res responses.EdgeSearchResponse

	if err := c.Call("edge.search", &req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
