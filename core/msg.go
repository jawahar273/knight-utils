package core

import (
	"google.golang.org/grpc/codes"
)

// msgFlag list of code and message
var msgFlag = map[int]string{
	Success:       "success",
	CreateSuccess: "success",

	// Redirect
	UnModifiedRedirect: "redirect unmodified",

	// 4xx
	InvalidRequest:      "invalid request",
	UnAuthorized:        "unauthorized user",
	Forbidden:           "access forbidden",
	NotFound:            "not found",
	MethodNotAllow:      "method not allowed",
	NotAcceptable:       "given request not accepted by the server",
	ToManyRequest:       "too many request",
	RequestTimeout:      "request timeout",
	Conflict:            "conflict in given request",
	ClientClosedRequest: "rquest cancelled the request",

	// 5xx
	ServerError:        "sometime went wrong",
	Unimplemented:      "service is not implemented",
	ServiceUnavailable: "service is unavailable",
	GatewayTimeOut:     "gateway timeout",
}

// GetMsg get message based code
func GetMsg(code int) string {
	msg, ok := msgFlag[code]
	if ok {
		return msg
	}
	return msgFlag[ServerError]
}

// https://gist.github.com/hamakn/708b9802ca845eb59f3975dbb3ae2a01
var grpcHttpMsg = map[codes.Code]int{

	codes.OK: Success,
	// CreateSuccess,

	// Redirect
	// UnModifiedRedirect

	// 4xx
	codes.InvalidArgument:    InvalidRequest,
	codes.FailedPrecondition: InvalidRequest,
	codes.OutOfRange:         InvalidRequest,
	codes.Unauthenticated:    UnAuthorized,
	codes.PermissionDenied:   Forbidden,
	codes.NotFound:           NotFound,
	codes.AlreadyExists:      Conflict,
	codes.Aborted:            Conflict,
	// MethodNotAllow,
	// NotAcceptable
	codes.ResourceExhausted: ToManyRequest,
	// RequestTimeout,
	codes.Canceled: ClientClosedRequest,

	// 5xx
	codes.Unknown:          ServerError,
	codes.Internal:         ServerError,
	codes.DataLoss:         ServerError,
	codes.Unimplemented:    Unimplemented,
	codes.Unavailable:      ServiceUnavailable,
	codes.DeadlineExceeded: GatewayTimeOut,
}

func GrpcCodeToHttpCode(code codes.Code) string {
	httpStatusCode, ok := grpcHttpMsg[code]
	if ok {
		return GetMsg(httpStatusCode)
	}

	return GetMsg(grpcHttpMsg[codes.Unknown])
}
