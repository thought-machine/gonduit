package gonduit

import (
	"github.com/thought-machine/gonduit/requests"
	"github.com/thought-machine/gonduit/responses"
)

// AlmanacDeviceSearch performs a call to almanac.device.search
func (c *Conn) AlmanacDeviceSearch () (*responses.SearchResponse, error) {
	var res responses.SearchResponse
	if err := c.Call("almanac.device.search", &requests.Request{}, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// AlmanacDeviceEdit performs a call to almanac.device.edit
func (c *Conn) AlmanacDeviceEdit (req requests.EditRequest) (*responses.EditResponse, error) {
	var res responses.EditResponse
	if err := c.Call("almanac.device.edit", &req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// AlmanacServiceSearch performs a call to almanac.service.search
func (c *Conn) AlmanacServiceSearch () (*responses.SearchResponse, error) {
	var res responses.SearchResponse
	if err := c.Call("almanac.service.search", &requests.Request{}, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// AlmanacServiceEdit performs a call to almanac.service.edit
func (c *Conn) AlmanacServiceEdit (req requests.EditRequest) (*responses.EditResponse, error) {
	var res responses.EditResponse
	if err := c.Call("almanac.service.edit", &req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

