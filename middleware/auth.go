package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/jawahar273/knight-utils/core"
)

func HasAuthorizationHeader() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.Request.Header["Authorization"]
		if auth == nil || auth == "" {
			response.ResponseError(
				c, core.UnAuthorized, "invalid or missing authorization token",
			)
		}
	}
}
