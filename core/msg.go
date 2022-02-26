package core

// MsgFlag list of code and message
var MsgFlag = map[int]string{
	SUCCESS:         "ok",
	SERVER_ERROR:    "fail",
	INVALID_REQUEST: "wrong parameter",
	CREATE_SUCCESS:  "entity created successfully",
}

// GetMsg get message based code
func GetMsg(code int) string {
	msg, ok := MsgFlag[code]
	if ok {
		return msg
	}
	return MsgFlag[SERVER_ERROR]
}
