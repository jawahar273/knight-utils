package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetMsg(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(GetMsg(SUCCESS), "ok")
	assert.Equal(GetMsg(SERVERERROR), "fail")
	assert.Equal(GetMsg(INVALIDREQUEST), "wrong parameter")

}
