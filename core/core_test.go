package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetMsg(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(GetMsg(success), "ok")
	assert.Equal(GetMsg(serverError), "sometime went wrong")
	assert.Equal(GetMsg(invalidRequest), "invalid request")
	assert.Equal(GetMsg(createSuccess), "entity created successfully")
	assert.Equal(GetMsg(unAuthorized), "unauthorized user")
	assert.Equal(GetMsg(methodNotAllow), "method not allowed")
	assert.Equal(GetMsg(notFound), "not found")

}
