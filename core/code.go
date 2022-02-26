package core

import "net/http"

const (
	// SUCCESS http code 200
	SUCCESS        = http.StatusOK
	CREATE_SUCCESS = http.StatusCreated

	// Client error
	INVALID_REQUEST  = http.StatusBadRequest
	UN_AUTHORIZED    = http.StatusUnauthorized
	FORBIDDEN        = http.StatusForbidden
	NOT_FOUND        = http.StatusNotFound
	METHOD_NOT_ALLOW = http.StatusMethodNotAllowed

	// Server error
	SERVER_ERROR        = http.StatusInternalServerError
	SERVICE_UNAVAILABLE = http.StatusServiceUnavailable
)
