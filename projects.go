package gonduit

import (
	"github.com/samwestmoreland/gonduit/requests"
	"github.com/samwestmoreland/gonduit/responses"
)

// ProjectQuery performs a call to project.query.
func (c *Conn) ProjectQuery(
	req requests.ProjectQueryRequest,
) (*responses.ProjectQueryResponse, error) {
	var res responses.ProjectQueryResponse

	if err := c.Call("project.query", &req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// ProjectSearch performs a call to project.search
func (c *Conn) ProjectSearch(req requests.SearchRequest) (*responses.SearchResponse, error) {
	var res responses.SearchResponse
	if err := c.Call("project.search", &req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *Conn) ProjectEdit(req requests.EditRequest) (*responses.EditResponse, error) {
	var res responses.EditResponse
	if err := c.Call("project.edit", &req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *Conn) ProjectColumnSearch(req requests.SearchRequest) (*responses.SearchResponse, error) {
	var res responses.SearchResponse
	if err := c.Call("project.column.search", &req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
