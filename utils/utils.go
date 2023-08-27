package utils

import (
	"github.com/waretool/go-common/env"
	"net/http"
	"strings"
	"time"
)

func IsProduction() bool {
	if env.GetEnv("ENVIRONMENT", "production") == "production" {
		return true
	}
	return false
}

func GetClientIp(r *http.Request) string {
	requester := r.Header.Get("X-Forwarded-For")
	if len(requester) == 0 {
		requester = r.Header.Get("X-Real-IP")
	}
	if len(requester) == 0 {
		requester = r.RemoteAddr
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
