package middleware

import (
	"github.com/gin-gonic/gin"

	"strconv"

	"github.com/prometheus/client_golang/prometheus"
)

var TotalRequest = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "http_request_total",
		Help: "Number of get request",
	},
	[]string{"status"},
)

func ApiMerticMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		status := strconv.Itoa(c.Writer.Status())
		TotalRequest.WithLabelValues(status).Inc()
	}
}
