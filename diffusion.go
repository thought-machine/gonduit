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
