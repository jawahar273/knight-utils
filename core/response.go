package core

import (
	"net/http"

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

type responseError struct {
	Code  int    `json:"code"`
	Error string `json:"error,omitempty"`
}

type Pagination struct {
	Total      int64 `json:"total"`
	PerPage    int64 `json:"perPage"`
	PageNumber int64 `json:"pageNumber"`
}

func ResponseSuccessList(c *gin.Context, code int, data interface{}, page Pagination) {
	c.JSON(code, responsePagination{
		Code: code,
		Data: data,
		Msg:  GetMsg(code),
		Page: page,
	})
}

// Response function
func ResponseSuccess(c *gin.Context, code int, data interface{}) {
	c.JSON(code, response{
		Code: code,
		Data: data,
		Msg:  GetMsg(code),
	})
}

// ResponseError function
func ResponseError(c *gin.Context, code int, msg string) {
	c.JSON(code, responseError{
		Code:  code,
		Error: msg,
	})
}

func RouterNotFound(c *gin.Context, msg string) {
	c.JSON(http.StatusNotFound, responseError{
		Code:  NOTFOUND,
		Error: msg,
	})
}
