package middleware

import (
	"../core/code"
	"../core/response"
	"github.com/gin-gonic/gin"
)

func HasAuthorizationHeader() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.Request.Header["Authorization"]
		if auth == nil || auth == "" {
			response.ResponseError(
				c, code.UnAuthorized, "invalid or missing authorization token",
			)
		}
	}
}
