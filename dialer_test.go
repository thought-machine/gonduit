package gonduit

import (
	"net/http/httptest"
	"testing"

	"github.com/etcinit/gonduit/core"
	"github.com/etcinit/gonduit/responses"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestDial(t *testing.T) {
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

	_, err := Dial(ts.URL, &core.ClientOptions{
		APIToken: "some-token",
	})

	assert.Nil(t, err)
}

func TestDial_withInvalid(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.POST("/api/conduit.getcapabilities", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"fake": "fake",
		})
	})

	ts := httptest.NewServer(r)
	defer ts.Close()

	_, err := Dial(ts.URL, &core.ClientOptions{
		APIToken: "some-token",
	})

	assert.NotNil(t, err)
}

func TestAssertSupportedCapabilities(t *testing.T) {
	response := responses.ConduitCapabilitiesResponse{
		Input:  []string{"urlencoded"},
		Output: []string{"json"},
	}

	assert.Nil(t, assertSupportedCapabilities(response, &core.ClientOptions{}))
}

func TestAssertSupportedCapabilities_withMissingInput(t *testing.T) {
	response := responses.ConduitCapabilitiesResponse{
		Input:  []string{"fake"},
		Output: []string{"json"},
	}

	assert.Equal(
		t,
		core.ErrURLEncodedInputUnsupported,
		assertSupportedCapabilities(response, &core.ClientOptions{}),
	)
}

func TestAssertSupportedCapabilities_withMissingOutput(t *testing.T) {
	response := responses.ConduitCapabilitiesResponse{
		Input:  []string{"urlencoded"},
		Output: []string{"fake"},
	}

	assert.Equal(
		t,
		core.ErrJSONOutputUnsupported,
		assertSupportedCapabilities(response, &core.ClientOptions{}),
	)
}

func TestAssertSupportedCapabilities_withNoToken(t *testing.T) {
	response := responses.ConduitCapabilitiesResponse{
		Authentication: []string{"session"},
		Input:          []string{"urlencoded"},
		Output:         []string{"json"},
	}

	assert.Equal(
		t,
		core.ErrTokenAuthUnsupported,
		assertSupportedCapabilities(response, &core.ClientOptions{
			APIToken: "super-secret-token",
		}),
	)
}

func TestAssertSupportedCapabilities_withNoCertificate(t *testing.T) {
	response := responses.ConduitCapabilitiesResponse{
		Authentication: []string{"token"},
		Input:          []string{"urlencoded"},
		Output:         []string{"json"},
	}

	assert.Equal(
		t,
		core.ErrSessionAuthUnsupported,
		assertSupportedCapabilities(response, &core.ClientOptions{
			Cert:     "super-secret-token",
			CertUser: "alice",
		}),
	)
}
