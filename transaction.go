package gonduit

import (
	"github.com/samwestmoreland/gonduit/requests"
	"github.com/samwestmoreland/gonduit/responses"
)

//TransactionSearch performs a call to transaction.search.
func (c *Conn) TransactionSearch(req requests.TransactionSearchRequest) (*responses.TransactionSearchResponse, error) {
	var res responses.TransactionSearchResponse
	if err := c.Call("transaction.search", &req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
