package core

// MsgFlag list of code and message
var MsgFlag = map[int]string{
	SUCCESS:          "ok",
	SERVER_ERROR:     "sometime went wrong",
	INVALID_REQUEST:  "invalid request",
	CREATE_SUCCESS:   "entity created successfully",
	UN_AUTHORIZED:    "unauthorized user",
	METHOD_NOT_ALLOW: "method not allowed",
}

// GetMsg get message based code
func GetMsg(code int) string {
	msg, ok := MsgFlag[code]
	if ok {
		return msg
	}
	return MsgFlag[SERVER_ERROR]
}
