package middleware

import (
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
)

func MatchPattern(queryParam, pattern string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		param := ctx.Param(queryParam)
		if match, err := regexp.MatchString(pattern, param); err != nil || !match {
			ctx.AbortWithStatus(http.StatusNotFound)
			return
		}
		ctx.Next()
	}
}
