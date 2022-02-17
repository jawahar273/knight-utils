package core

import "net/http"

const (
	// SUCCESS http code 200
	SUCCESS       = http.StatusOK
	CREATESUCCESS = http.StatusCreated

	// Client error
	INVALIDREQUEST = http.StatusBadRequest
	UNAUTHORIZED   = http.StatusUnauthorized
	FORBIDDEN      = http.StatusForbidden
	NOTFOUND       = http.StatusNotFound
	METHODNOTALLOW = http.StatusMethodNotAllowed

	// Server error
	SERVERERROR        = http.StatusInternalServerError
	SERVICEUNAVAILABLE = http.StatusServiceUnavailable
)
