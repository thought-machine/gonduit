package gonduit

import (
	"github.com/thought-machine/gonduit/requests"
	"github.com/thought-machine/gonduit/responses"
)

// AlmanacNetworkSearch performs a call to almanac.network.search
func (c *Conn) AlmanacNetworkSearch (req requests.Request) (*responses.SearchResponse, error) {
	var res responses.SearchResponse
	if err := c.Call("almanac.network.search", &req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// AlmanacNetworkEdit performs a call to almanac.network.edit
func (c *Conn) AlmanacNetworkEdit (req requests.EditRequest) (*responses.EditResponse, error) {
	var res responses.EditResponse
	if err := c.Call("almanac.network.edit", &req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// AlmanacDeviceSearch performs a call to almanac.device.search
func (c *Conn) AlmanacDeviceSearch (req requests.Request) (*responses.SearchResponse, error) {
	var res responses.SearchResponse
	if err := c.Call("almanac.device.search", &req, &res); err != nil {
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
func (c *Conn) AlmanacServiceSearch (req requests.Request) (*responses.SearchResponse, error) {
	var res responses.SearchResponse
	if err := c.Call("almanac.service.search", &req, &res); err != nil {
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

// AlmanacBindingSearch performs a call to almanac.binding.search
func (c *Conn) AlmanacBindingSearch (req requests.Request) (*responses.SearchResponse, error) {
	var res responses.SearchResponse
	if err := c.Call("almanac.binding.search", &req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// AlmanacBindingEdit performs a call to almanac.binding.edit
func (c *Conn) AlmanacBindingEdit (req requests.EditRequest) (*responses.EditResponse, error) {
	var res responses.EditResponse
	if err := c.Call("almanac.binding.edit", &req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// AlmanacInterfaceSearch performs a call to almanac.interface.search
func (c *Conn) AlmanacInterfaceSearch (req requests.Request) (*responses.SearchResponse, error) {
	var res responses.SearchResponse
	if err := c.Call("almanac.interface.search", &req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// AlmanacInterfaceEdit performs a call to almanac.interface.edit
func (c *Conn) AlmanacInterfaceEdit (req requests.EditRequest) (*responses.EditResponse, error) {
	var res responses.EditResponse
	if err := c.Call("almanac.interface.edit", &req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

