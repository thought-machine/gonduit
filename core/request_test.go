package core

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"
	"testing"

	"github.com/etcinit/gonduit/requests"
	"github.com/stretchr/testify/assert"
)

func TestMakeRequest(t *testing.T) {
	request := requests.ConduitConnectRequest{
		Client:        "gonduit",
		ClientVersion: "1",
	}
	body, _ := prepareForm(&request)

	req, err := MakeRequest("http://localhost/api/conduit.connect", &request)

	assert.Nil(t, err)
	assert.Equal(t, "http://localhost/api/conduit.connect", req.URL.String())
	assert.Equal(t, "POST", req.Method)

	buffer := new(bytes.Buffer)
	buffer.ReadFrom(req.Body)
	assert.Equal(t, body.Encode(), buffer.String())
}

func TestMakeRequest_withInvalidBody(t *testing.T) {
	_, err := MakeRequest("http://localhost/api/conduit.connect", http.Request{})

	assert.NotNil(t, err)
}

func TestMakeRequest_withInvalidRequest(t *testing.T) {
	request := requests.ConduitConnectRequest{
		Client:        "gonduit",
		ClientVersion: "1",
	}

	_, err := MakeRequest("htcp.\\invaid#&^%", &request)

	assert.NotNil(t, err)
}

func TestSetHeaders(t *testing.T) {
	request := &http.Request{
		Header: http.Header{},
	}

	setHeaders(request)

	assert.NotEqual(t, "", request.Header.Get("Content-Type"))
}

func TestPrepareForm(t *testing.T) {
	request := requests.ConduitConnectRequest{
		Client:        "gonduit",
		ClientVersion: "1",
	}
	marshalled, _ := json.Marshal(request)

	form, err := prepareForm(&request)

	assert.Nil(t, err)
	assert.Equal(t, "true", form.Get("__conduit__"))
	assert.Equal(t, "json", form.Get("output"))
	assert.Equal(t, string(marshalled), form.Get("params"))
}

func TestPrepareForm_withNoBody(t *testing.T) {
	form, err := prepareForm(nil)

	assert.Nil(t, err)
	assert.Equal(t, "", form.Get("__conduit__"))
	assert.Equal(t, "json", form.Get("output"))
	assert.Equal(t, "", form.Get("params"))
}

func TestPrepareForm_withError(t *testing.T) {
	_, err := prepareForm(http.Request{})

	assert.NotNil(t, err)
}

func TestHandleConnectRequest(t *testing.T) {
	request := requests.ConduitConnectRequest{}
	request2 := requests.DifferentialQueryRequest{}

	form := url.Values{}
	form2 := url.Values{}

	handleConnectRequest(&form, &request)
	handleConnectRequest(&form2, &request2)

	assert.Equal(t, "true", form.Get("__conduit__"))
	assert.Equal(t, "", form2.Get("__conduit__"))
}
