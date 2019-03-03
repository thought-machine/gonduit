package gonduit

import (
	"github.com/thought-machine/gonduit/requests"
	"github.com/thought-machine/gonduit/responses"
)

// HarbormasterCreateArtifact performs a call to harbormaster.createartifact.
func (c *Conn) HarbormasterCreateArtifact(req requests.HarbormasterCreateArtifactRequest) (*responses.HarbormasterCreateArtifactResponse, error) {
	var res responses.HarbormasterCreateArtifactResponse
	if err := c.Call("harbormaster.createartifact", &req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
