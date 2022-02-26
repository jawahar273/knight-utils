package core

import (
	"github.com/gin-gonic/gin"
)

// Response struct
type response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type responsePagination struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
	Page Pagination  `json:"page"`
}

// ErrorField are use to display the field of error and reason under message attribute
type ErrorFieldAndMessage struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type responseError struct {
	Code   int                    `json:"code"`
	Errors []ErrorFieldAndMessage `json:"errors,omitempty"`
	Msg    string                 `json:"message"`
	Title  string                 `json:"title"`
}

// Pagination use as meta data on returning list of entitys
type Pagination struct {
	Total      int64 `json:"total"`
	PerPage    int64 `json:"perPage"`
	PageNumber int64 `json:"pageNumber"`
}

// ResponseSuccessList return the list of entity with pagination metadata
func ResponseSuccessList(c *gin.Context, code int, data interface{}, page Pagination) {
	c.JSON(code, responsePagination{
		Code: code,
		Data: data,
		Msg:  GetMsg(code),
		Page: page,
	})
}

// ResponseSuccess return single entity
func ResponseSuccess(c *gin.Context, code int, data interface{}) {
	c.JSON(code, response{
		Code: code,
		Data: data,
		Msg:  GetMsg(code),
	})
}

// ResponseError return error with structure
// example:
// {
//	"code": 400,
//	"title": "Invalid request",
//	"errors": [{
//		"field": "user_name",
//		"message": "user name is required",
//	}],
//	"message": ""
// }
func ResponseError(c *gin.Context, code int, msg string, errors []ErrorFieldAndMessage) {
	c.JSON(code, responseError{
		Code:   code,
		Title:  GetMsg(code),
		Errors: errors,
		Msg:    msg,
	})
}

// func ResponseNotFound(c *gin.Context, msg string) {
// 	c.JSON(http.StatusNotFound, responseError{
// 		Code:  NOT_FOUND,
// 		E: GetMsg(msg),
// 	})
// }
