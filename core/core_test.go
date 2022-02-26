package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetMsg(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(GetMsg(SUCCESS), "ok")
	assert.Equal(GetMsg(SERVER_ERROR), "sometime went wrong")
	assert.Equal(GetMsg(INVALID_REQUEST), "invalid request")
	assert.Equal(GetMsg(CREATE_SUCCESS), "entity created successfully")
	assert.Equal(GetMsg(UN_AUTHORIZED), "unauthorized user")
	assert.Equal(GetMsg(METHOD_NOT_ALLOW), "method not allowed")
}
