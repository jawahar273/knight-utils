package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
)

func TestGetMsg(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(GetMsg(Success), "success")
	assert.Equal(GetMsg(ServerError), "sometime went wrong")
	assert.Equal(GetMsg(InvalidRequest), "invalid request")
	assert.Equal(GetMsg(CreateSuccess), "success")
	assert.Equal(GetMsg(UnAuthorized), "unauthorized user")
	assert.Equal(GetMsg(MethodNotAllow), "method not allowed")
	assert.Equal(GetMsg(NotFound), "not found")

}

func TestGrpcMsg(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(GrpcCodeToHttpCode(codes.OK), "success")
	assert.Equal(GrpcCodeToHttpCode(codes.Unknown), "sometime went wrong")
}
