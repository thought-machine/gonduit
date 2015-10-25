package core

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConduitErrorCode(t *testing.T) {
	err := ConduitError{
		code: "404",
		info: "OMG WHY",
	}

	assert.Equal(t, "404", err.Code())
}

func TestConduitErrorInfo(t *testing.T) {
	err := ConduitError{
		code: "404",
		info: "OMG WHY",
	}

	assert.Equal(t, "OMG WHY", err.Info())
}

func TestConduitErrorError(t *testing.T) {
	err := ConduitError{
		code: "404",
		info: "OMG WHY",
	}

	assert.Equal(t, "404: OMG WHY", err.Error())
}

func TestIsConduitError(t *testing.T) {
	err := ConduitError{
		code: "404",
		info: "OMG WHY",
	}
	err2 := fmt.Errorf("OMG")
	var err3 error

	assert.True(t, IsConduitError(error(&err)))
	assert.False(t, IsConduitError(err2))
	assert.False(t, IsConduitError(err3))
}
