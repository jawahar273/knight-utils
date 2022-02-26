package core

import "net/http"

const (
	// SUCCESS http code 200
	SUCCESS        = http.StatusOK
	CREATE_SUCCESS = http.StatusCreated

	// Client error
	INVALID_REQUEST  = http.StatusBadRequest
	UNAUTHORIZED     = http.StatusUnauthorized
	FORBIDDEN        = http.StatusForbidden
	NOTFOUND         = http.StatusNotFound
	METHOD_NOT_ALLOW = http.StatusMethodNotAllowed

	// Server error
	SERVER_ERROR        = http.StatusInternalServerError
	SERVICE_UNAVAILABLE = http.StatusServiceUnavailable
)
