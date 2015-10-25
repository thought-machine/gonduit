package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMakeHTTPClient(t *testing.T) {
	client := makeHTTPClient(&ClientOptions{
		InsecureSkipVerify: true,
	})

	client2 := makeHTTPClient(&ClientOptions{
		InsecureSkipVerify: false,
	})

	assert.NotNil(t, client)
	assert.NotNil(t, client2)
}
