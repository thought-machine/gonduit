package gonduit

import (
	"github.com/etcinit/gonduit/requests"
	"github.com/etcinit/gonduit/responses"
)

// RepositoryQuery performs a call to repository.query.
func (c *Conn) RepositoryQuery(req requests.RepositoryQueryRequest) (*responses.RepositoryQueryResponse, error) {
	var res responses.RepositoryQueryResponse

	req.Session = c.Session

	if err := c.Call("repository.query", &req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
