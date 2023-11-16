package gonduit

import (
	"github.com/samwestmoreland/gonduit/requests"
	"github.com/samwestmoreland/gonduit/responses"
)

// RemarkupProcess performs a call to remarkup.process
func (c *Conn) RemarkupProcess(
	req requests.RemarkupProcessRequest,
) (*responses.RemarkupProcessResponse, error) {
	var res responses.RemarkupProcessResponse

	if err := c.Call("remarkup.process", &req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
