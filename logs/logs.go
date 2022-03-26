package logs

import (
	"time"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var logger *zap.Logger

// SetupZap any custom format can done here
func Setupzap() {
	logger, _ = zap.NewProduction()
}

func SetupDevzap() {
	logger, _ = zap.NewDevelopment()
}

func Info(msg string, fields ...zap.Field) {
	logger.Info(msg, fields...)
}

func Debug(msg string, fields ...zap.Field) {
	logger.Debug(msg, fields...)
}

func Fatal(msg string, fields ...zap.Field) {
	logger.Fatal(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	logger.Error(msg, fields...)
}

func ErrorField(err error) zap.Field {
	return zap.Error(err)
}

func LogMiddleware(router *gin.Engine) {
	// Logs
	// logger, _ := zap.NewProduction()
	defer logger.Sync()
	// Add a ginzap middleware, which:
	//   - Logs all requests, like a combined access and error log.
	//   - Logs to stdout.
	//   - RFC3339 with UTC time format.
	router.Use(ginzap.Ginzap(logger, time.RFC3339, true))
	// Logs all panic to error log
	//   - stack means whether output the stack info.
	router.Use(ginzap.RecoveryWithZap(logger, true))

	// Skip a path
	//r.Use(ginzap.GinzapWithConfig(logger, &ginzap.Config{
	//  TimeFormat: time.RFC3339,
	//  UTC: true,
	//  SkipPaths: []string{"/no_log"},
	//}))

}
