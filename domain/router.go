package domain

import "net/http"

type Middleware func(next http.Handler) http.Handler

type Router interface {
	http.Handler
	UseMiddleware(middleware Middleware)
}
