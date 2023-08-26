package middleware

import (
	"context"
	"github.com/waretool/go-common/service"
	"net/http"
	"regexp"
)

var bearerRgx = regexp.MustCompile(`(?i)bearer `)

func AuthN(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header, ok := r.Header["Authorization"]
		if !ok || len(header) != 1 {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		jwt := bearerRgx.ReplaceAllString(header[0], "")
		if claims, valid := service.NewJwtService().Valid(jwt); !valid || !claims.Enabled {
			w.WriteHeader(http.StatusUnauthorized)
			return
		} else {
			ctx := context.WithValue(r.Context(), "claims", claims)
			h.ServeHTTP(w, r.WithContext(ctx))
		}
	})
}
