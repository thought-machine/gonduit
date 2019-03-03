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

// HarbormasterSendMessage performs a call to harbormaster.sendmessage.
func (c *Conn) HarbormasterSendMessage(req requests.HarbormasterSendMessageRequest) (*responses.HarbormasterSendMessageResponse, error) {
	var res responses.HarbormasterSendMessageResponse
	if err := c.Call("harbormaster.sendmessage", &req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
