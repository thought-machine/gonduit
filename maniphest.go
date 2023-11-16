package gonduit

import (
	"github.com/samwestmoreland/gonduit/entities"
	"github.com/samwestmoreland/gonduit/requests"
	"github.com/samwestmoreland/gonduit/responses"
)

// ManiphestQuery performs a call to maniphest.query.
// Deprecated: This method is frozen and will eventually be deprecated. New code should use "maniphest.search" instead.
func (c *Conn) ManiphestQuery(
	req requests.ManiphestQueryRequest,
) (*responses.ManiphestQueryResponse, error) {
	var res responses.ManiphestQueryResponse

	if err := c.Call("maniphest.query", &req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// ManiphestCreateTask performs a call to maniphest.createtask.
// Deprecated: This method is frozen and will eventually be deprecated. New code should use "maniphest.edit" instead.
func (c *Conn) ManiphestCreateTask(
	req requests.ManiphestCreateTaskRequest,
) (*entities.ManiphestTask, error) {
	var res entities.ManiphestTask

	if err := c.Call("maniphest.createtask", &req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// ManiphestGetTaskTransactions performs a call to maniphest.gettasktransactions
// Deprecated: This method is frozen and will eventually be deprecated. New code should use "transaction.search" instead.
func (c *Conn) ManiphestGetTaskTransactions(
	req requests.ManiphestGetTaskTransactions,
) (*responses.ManiphestGetTaskTransactionsResponse, error) {
	var res responses.ManiphestGetTaskTransactionsResponse

	if err := c.Call("maniphest.gettasktransactions", &req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// ManiphestSearch performs a call to maniphest.search
func (c *Conn) ManiphestSearch(req requests.SearchRequest) (*responses.SearchResponse, error) {
	var res responses.SearchResponse
	if err := c.Call("maniphest.search", &req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// ManiphestStatusSearch performs a call to maniphest.status.search
func (c *Conn) ManiphestStatusSearch(req requests.SearchRequest) (*responses.SearchResponse, error) {
	var res responses.SearchResponse
	if err := c.Call("maniphest.status.search", &req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// ManiphestEdit performs a call to maniphest.edit
func (c *Conn) ManiphestEdit(req requests.EditRequest) (*responses.EditResponse, error) {
	var res responses.EditResponse
	if err := c.Call("maniphest.edit", &req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
