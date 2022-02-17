package core

// MsgFlag list of code and message
var MsgFlag = map[int]string{
	SUCCESS:        "ok",
	SERVERERROR:    "fail",
	INVALIDREQUEST: "wrong parameter",
	CREATESUCCESS:  "entity created successfully",
}

// GetMsg get message based code
func GetMsg(code int) string {
	msg, ok := MsgFlag[code]
	if ok {
		return msg
	}
	return MsgFlag[SERVERERROR]
}
