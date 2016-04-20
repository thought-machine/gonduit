package core

import (
	"net/http/httptest"
	"testing"

	"github.com/etcinit/gonduit/responses"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetEndpointURI(t *testing.T) {
	assert.Equal(
		t,
		"phabricator.gonduit.wow/api/conduit.connect",
		GetEndpointURI("phabricator.gonduit.wow/", "conduit.connect"),
	)
}

func TestPerformCall(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.POST("/api/conduit.getcapabilities", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"result": map[string][]string{
				"authentication": []string{"token", "session"},
				"signatures":     []string{"consign"},
				"input":          []string{"json", "urlencoded"},
				"output":         []string{"json"},
			},
		})
	})

	ts := httptest.NewServer(r)
	defer ts.Close()

	result := map[string]interface{}{}

	err := PerformCall(
		ts.URL+"/api/conduit.getcapabilities",
		map[string]interface{}{},
		&result,
		&ClientOptions{},
	)

	assert.Nil(t, err)
}

func TestPerformCall_withEmptyArray(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.POST("/api/phid.lookup", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"result": []string{},
		})
	})

	ts := httptest.NewServer(r)
	defer ts.Close()

	var result responses.PHIDLookupResponse
	ptr := &result

	err := PerformCall(
		ts.URL+"/api/phid.lookup",
		map[string]interface{}{},
		&ptr,
		&ClientOptions{},
	)

	assert.Nil(t, err)
}

func TestPerformCall_withErrorCode(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.POST("/api/conduit.getcapabilities", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"result": map[string][]string{
				"authentication": []string{"token", "session"},
				"signatures":     []string{"consign"},
				"input":          []string{"json", "urlencoded"},
				"output":         []string{"json"},
			},
			"error_code": "ERR-CONDUIT-CORE",
			"error_info": "Something bad happened",
		})
	})

	ts := httptest.NewServer(r)
	defer ts.Close()

	result := map[string]interface{}{}

	err := PerformCall(
		ts.URL+"/api/conduit.getcapabilities",
		map[string]interface{}{},
		&result,
		&ClientOptions{},
	)

	assert.Equal(t, &ConduitError{
		code: "ERR-CONDUIT-CORE",
		info: "Something bad happened",
	}, err)
}

func TestPerformCall_withMissingResults(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.POST("/api/conduit.getcapabilities", func(c *gin.Context) {
		c.JSON(200, gin.H{})
	})

	ts := httptest.NewServer(r)
	defer ts.Close()

	result := map[string]interface{}{}

	err := PerformCall(
		ts.URL+"/api/conduit.getcapabilities",
		map[string]interface{}{},
		&result,
		&ClientOptions{},
	)

	assert.Equal(t, ErrMissingResults, err)
}
