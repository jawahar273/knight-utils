package core

// MsgFlag list of code and message
var MsgFlag = map[int]string{
	Success:            "ok",
	InvalidRequest:     "invalid request",
	CreateSuccess:      "entity created successfully",
	UnModifiedRedirect: "redirect unmodified",
	// 4xx
	UnAuthorized:   "unauthorized user",
	MethodNotAllow: "method not allowed",
	NotFound:       "not found",
	// 5xx
	ServerError: "sometime went wrong",
}

// GetMsg get message based code
func GetMsg(code int) string {
	msg, ok := MsgFlag[code]
	if ok {
		return msg
	}
	return MsgFlag[ServerError]
}
