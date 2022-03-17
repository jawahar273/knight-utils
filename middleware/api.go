package middleware

import (
	"github.com/gin-contrib/secure"
	"github.com/gin-gonic/gin"
)

// import (
// 	"github.com/gin-gonic/gin"

// 	"strconv"

// 	"github.com/prometheus/client_golang/prometheus"
// )

// var TotalRequest = prometheus.NewCounterVec(
// 	prometheus.CounterOpts{
// 		Name: "http_request_total",
// 		Help: "Number of get request",
// 	},
// 	[]string{"status"},
// )

// func ApiMerticMiddleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		status := strconv.Itoa(c.Writer.Status())
// 		TotalRequest.WithLabelValues(status).Inc()
// 	}
// }

func DefaultSecureGinMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		secure.New(secure.DefaultConfig())
		c.Header("X-Content-Type-Options", "nosniff")
		c.Header("X-Frame-Options", "deny")
		c.Header("Content-Security-Policy", "default-src 'none'")
		c.Header("X-Powered-By", "")
		c.Header("content-type", "application/json")
	}
}
