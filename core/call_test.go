package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetEndpointURI(t *testing.T) {
	assert.Equal(
		t,
		"phabricator.gonduit.wow/api/conduit.connect",
		GetEndpointURI("phabricator.gonduit.wow/", "conduit.connect"),
	)
}
