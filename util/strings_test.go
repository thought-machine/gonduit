package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContainsString(t *testing.T) {
	assert.True(t, ContainsString(
		[]string{"wow", "omg", "omg", "yes"},
		"omg",
	))

	assert.False(t, ContainsString(
		[]string{"wow", "omg", "omg", "yes"},
		"omgomg",
	))
}
