package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	contextUtils "go-common/context"
	"go-common/logger"
	"slices"
	"time"
)

func LogMiddleware(skipPaths []string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()

		// Process Request
		ctx.Next()

		if slices.Contains(skipPaths, ctx.Request.RequestURI) {
			return
		} else {
			duration := contextUtils.GetDurationInMilliseconds(start)
			cookie := ctx.GetHeader("Cookie")
			if cookieLen := len(cookie); cookieLen > 0 {
				cookie = cookie[:cookieLen/2] + "***"
			}
			logger.GetLogger().WithFields(logrus.Fields{
				"request": map[string]interface{}{
					"clientIp": contextUtils.GetClientIP(ctx),
					"method":   ctx.Request.Method,
					"path":     ctx.Request.RequestURI,
					"headers": map[string]interface{}{
						"host":         ctx.Request.Host,
						"connection":   ctx.GetHeader("Connection"),
						"content-type": ctx.GetHeader("Content-Type"),
						"cookie":       cookie,
						"user-agent":   ctx.GetHeader("User-Agent"),
					},
				},
				"status":   ctx.Writer.Status(),
				"duration": duration,
			}).Info()
		}
	}
}
