package model

import (
	"github.com/gorilla/mux"
	"github.com/waretool/go-common/domain"
	"net/http"
)

type MuxRouter struct {
	*mux.Router
}

func NewMuxRouter() *MuxRouter {
	return &MuxRouter{Router: mux.NewRouter()}
}

func (mr *MuxRouter) UseMiddleware(middleware domain.Middleware) {
	mr.Router.Use(func(next http.Handler) http.Handler {
		handler := middleware(next)
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			handler.ServeHTTP(w, r)
		})
	})
}
