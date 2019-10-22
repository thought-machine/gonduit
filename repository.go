package gonduit

import (
	"github.com/thought-machine/gonduit/requests"
	"github.com/thought-machine/gonduit/responses"
)

// RepositoryQuery performs a call to repository.query.
// Deprecated: This method is frozen and will eventually be deprecated. New code should use "diffusion.repository.search" instead.
func (c *Conn) RepositoryQuery(
	req requests.RepositoryQueryRequest,
) (*responses.RepositoryQueryResponse, error) {
	var res responses.RepositoryQueryResponse

	if err := c.Call("repository.query", &req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
