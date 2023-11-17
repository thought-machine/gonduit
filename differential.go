package gonduit

import (
	"fmt"
	"github.com/samwestmoreland/gonduit/requests"
	"github.com/samwestmoreland/gonduit/responses"
)

// DifferentialQuery performs a call to differential.query.
// Deprecated: This method is frozen and will eventually be deprecated. New code should use "differential.revision.search" instead.
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
// Deprecated: This method is frozen and will eventually be deprecated. New code should use "differential.diff.search" instead.
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
	fmt.Printf("req: %+v\n", req)
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

// DifferentialGetRawDiff performs a call to differential.getrawdiff
func (c *Conn) DifferentialGetRawDiff(req requests.DifferentialGetRawDiffRequest) (*responses.DifferentialGetRawDiffResponse, error) {
	var res responses.DifferentialGetRawDiffResponse
	fmt.Printf("req: %+v\n", req)
	if err := c.Call("differential.getrawdiff", &req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
