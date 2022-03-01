package core

import "net/http"

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

	// Server error
	ServerError        = http.StatusInternalServerError
	ServiceUnavailable = http.StatusServiceUnavailable
)
