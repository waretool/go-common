package context

import (
	"github.com/gin-gonic/gin"
	"strings"
	"time"
)

func GetClientIP(ctx *gin.Context) string {
	requester := ctx.Request.Header.Get("X-Forwarded-For")
	if len(requester) == 0 {
		requester = ctx.Request.Header.Get("X-Real-IP")
	}
	if len(requester) == 0 {
		requester = ctx.Request.RemoteAddr
	}
	if strings.Contains(requester, ",") {
		requester = strings.Split(requester, ",")[0]
	}

	return requester
}

func GetDurationInMilliseconds(start time.Time) float64 {
	end := time.Now()
	duration := end.Sub(start)
	milliseconds := float64(duration) / float64(time.Millisecond)
	rounded := float64(int(milliseconds*100+.5)) / 100
	return rounded
}
