package gonduit

import (
	"testing"

	"github.com/thought-machine/gonduit/core"
	"github.com/thought-machine/gonduit/test/server"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestConduitQuery(t *testing.T) {
	s := server.New()
	defer s.Close()
	s.RegisterCapabilities()

	s.RegisterMethod("conduit.query", 200, gin.H{
		"result": gin.H{
			"phid.query": gin.H{
				"description": "Retrieve information about arbitrary PHIDs.",
				"params": gin.H{
					"phids": "required list<phid>",
				},
				"return": "nonempty dict<string, wild>",
			},
		},
	})

	c, err := Dial(s.GetURL(), &core.ClientOptions{
		APIToken: "some-token",
	})
	assert.Nil(t, err)

	r, err := c.ConduitQuery()
	assert.Nil(t, err)

	assert.Equal(
		t,
		"nonempty dict<string, wild>",
		(*r)["phid.query"].Return,
	)
}
