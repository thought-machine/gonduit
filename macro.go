package gonduit

import (
	"github.com/etcinit/gonduit/requests"
	"github.com/etcinit/gonduit/responses"
)

// MacroCreateMeme performs a call to macro.creatememe.
func (c *Conn) MacroCreateMeme(req requests.MacroCreateMemeRequest) (*responses.MacroCreateMemeResponse, error) {
	var res responses.MacroCreateMemeResponse

	req.Session = c.Session

	if err := c.Call("macro.creatememe", &req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
