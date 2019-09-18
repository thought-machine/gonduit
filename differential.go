package gonduit

import (
	"github.com/thought-machine/gonduit/requests"
	"github.com/thought-machine/gonduit/responses"
)

// DifferentialQuery performs a call to differential.query.
func (c *Conn) DifferentialQuery(
	req requests.DifferentialQueryRequest,
) (*responses.DifferentialQueryResponse, error) {
	var res responses.DifferentialQueryResponse

	if err := c.Call("differential.query", &req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// DifferentialQueryDiffs performs a call to differential.querydiffs.
func (c *Conn) DifferentialQueryDiffs(req requests.DifferentialQueryDiffsRequest) (responses.DifferentialQueryDiffsResponse, error) {
	var res responses.DifferentialQueryDiffsResponse
	if err := c.Call("differential.querydiffs", &req, &res); err != nil {
		return nil, err
	}
	return res, nil
}

// DifferentialDiffSearch performs a call to differential.diff.search
func (c *Conn) DifferentialDiffSearch(req requests.SearchRequest) (*responses.SearchResponse, error) {
	var res responses.SearchResponse
	if err := c.Call("differential.diff.search", &req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// DifferentialRevisionSearch performs a call to differential.revision.search
func (c *Conn) DifferentialRevisionSearch(req requests.SearchRequest) (*responses.SearchResponse, error) {
	var res responses.SearchResponse
	if err := c.Call("differential.revision.search", &req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// DifferentialRevisionEdit performs a call to differential.revision.edit
func (c *Conn) DifferentialRevisionEdit(req requests.EditRequest) (*responses.EditResponse, error) {
	var res responses.EditResponse
	if err := c.Call("differential.revision.edit", &req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
