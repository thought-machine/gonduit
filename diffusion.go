package gonduit

import (
	"github.com/etcinit/gonduit/entities"
	"github.com/etcinit/gonduit/requests"
	"github.com/etcinit/gonduit/responses"
)

// DiffusionQueryCommits performs a call to diffusion.querycommits.
func (c *Conn) DiffusionQueryCommits(req requests.DiffusionQueryCommitsRequest) (*responses.DiffusionQueryCommitsResponse, error) {
	var res responses.DiffusionQueryCommitsResponse

	req.Session = c.Session

	if err := c.Call("diffusion.querycommits", &req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// RepositoryQuery performs a call to repository.query.
func (c *Conn) RepositoryQuery(req requests.RepositoryQueryRequest) (*[]entities.DiffusionRepository, error) {
	var res []entities.DiffusionRepository

	req.Session = c.Session

	if err := c.Call("repository.query", &req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
