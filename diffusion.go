package gonduit

import (
	"github.com/thought-machine/gonduit/requests"
	"github.com/thought-machine/gonduit/responses"
)

// DiffusionQueryCommits performs a call to diffusion.querycommits.
func (c *Conn) DiffusionQueryCommits(
	req requests.DiffusionQueryCommitsRequest,
) (*responses.DiffusionQueryCommitsResponse, error) {
	var res responses.DiffusionQueryCommitsResponse

	if err := c.Call("diffusion.querycommits", &req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// DiffusionCommitSearch performs a call to diffusion.commit.search
func (c *Conn) DiffusionCommitSearch(req requests.SearchRequest) (*responses.SearchResponse, error) {
	var res responses.SearchResponse
	if err := c.Call("diffusion.commit.search", &req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// DiffusionRepositorySearch performs a call to diffusion.repository.search
func (c *Conn) DiffusionRepositorySearch(req requests.SearchRequest) (*responses.SearchResponse, error) {
	var res responses.SearchResponse
	if err := c.Call("diffusion.repository.search", &req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
