package core

// MsgFlag list of code and message
var MsgFlag = map[int]string{
	success:        "ok",
	invalidRequest: "invalid request",
	createSuccess:  "entity created successfully",
	// 4xx
	unAuthorized:   "unauthorized user",
	methodNotAllow: "method not allowed",
	notFound:       "not found",
	// 5xx
	serverError: "sometime went wrong",
}

// GetMsg get message based code
func GetMsg(code int) string {
	msg, ok := MsgFlag[code]
	if ok {
		return msg
	}
	return MsgFlag[serverError]
}
