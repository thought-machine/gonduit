package gonduit

import "github.com/etcinit/gonduit/responses"

// ConduitQuery performs a call to conduit.query.
func (c *Conn) ConduitQuery() (*responses.ConduitQueryResponse, error) {
	var res responses.ConduitQueryResponse

	if err := c.Call("conduit.query", nil, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
