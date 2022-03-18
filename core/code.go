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
	InvalidRequest = http.StatusBadRequest
	UnAuthorized   = http.StatusUnauthorized
	Forbidden      = http.StatusForbidden
	NotFound       = http.StatusNotFound
	MethodNotAllow = http.StatusMethodNotAllowed
	NotAcceptable  = http.StatusNotAcceptable

	// Server error
	ServerError        = http.StatusInternalServerError
	ServiceUnavailable = http.StatusServiceUnavailable
)

var InValidEntityIdError = fmt.Errorf("invalid entity id")
