package domain

import (
	"github.com/gin-gonic/gin"
	adapter "github.com/gwatts/gin-adapter"
	"github.com/waretool/go-common/domain"
	"github.com/waretool/go-common/utils"
)

type GinRouter struct {
	*gin.Engine
}

func NewGinRouter() *GinRouter {
	if utils.IsProduction() {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()
	r.Use(gin.Recovery())
	return &GinRouter{Engine: r}
}

func (gr *GinRouter) UseMiddleware(middleware domain.Middleware) {
	gr.Use(adapter.Wrap(middleware))
}
