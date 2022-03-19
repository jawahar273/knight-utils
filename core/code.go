package core

import (
	"fmt"
	"net/http"
)

const (
	// SUCCESS http code 200
	Success       = http.StatusOK
	CreateSuccess = http.StatusCreated

	// Redirect
	UnModifiedRedirect = http.StatusNotModified

	// Client error
	InvalidRequest      = http.StatusBadRequest
	UnAuthorized        = http.StatusUnauthorized
	Forbidden           = http.StatusForbidden
	NotFound            = http.StatusNotFound
	MethodNotAllow      = http.StatusMethodNotAllowed
	NotAcceptable       = http.StatusNotAcceptable
	ToManyRequest       = http.StatusTooManyRequests
	RequestTimeout      = http.StatusRequestTimeout
	Conflict            = http.StatusConflict
	ClientClosedRequest = 499

	// Server error
	ServerError        = http.StatusInternalServerError
	Unimplemented      = http.StatusNotImplemented
	ServiceUnavailable = http.StatusServiceUnavailable
	GatewayTimeOut     = http.StatusGatewayTimeout
)

var ErrInValidEntityId = fmt.Errorf("invalid entity id")
