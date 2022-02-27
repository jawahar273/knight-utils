package core

import "net/http"

const (
	// SUCCESS http code 200
	success       = http.StatusOK
	createSuccess = http.StatusCreated

	// Client error
	invalidRequest = http.StatusBadRequest
	unAuthorized   = http.StatusUnauthorized
	forbidden      = http.StatusForbidden
	notFound       = http.StatusNotFound
	methodNotAllow = http.StatusMethodNotAllowed

	// Server error
	serverError        = http.StatusInternalServerError
	serviceUnavailable = http.StatusServiceUnavailable
)
