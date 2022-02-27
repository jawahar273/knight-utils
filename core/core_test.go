package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetMsg(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(GetMsg(Success), "ok")
	assert.Equal(GetMsg(ServerError), "sometime went wrong")
	assert.Equal(GetMsg(InvalidRequest), "invalid request")
	assert.Equal(GetMsg(CreateSuccess), "entity created successfully")
	assert.Equal(GetMsg(UnAuthorized), "unauthorized user")
	assert.Equal(GetMsg(MethodNotAllow), "method not allowed")
	assert.Equal(GetMsg(NotFound), "not found")

}
