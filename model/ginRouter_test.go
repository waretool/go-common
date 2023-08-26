package model

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
	"github.com/waretool/go-common/domain"
	"net/http"
	"testing"
)

type GinRouterSuite struct {
	suite.Suite
}

func TestGinRouterSuite(t *testing.T) {
	suite.Run(t, new(GinRouterSuite))
}

func (suite *GinRouterSuite) TestGinRouterImplementsRouter() {
	suite.Implementsf((*domain.Router)(nil), new(GinRouter), "GinRouter does not implement the domain.Router interface")
}

func (suite *GinRouterSuite) TestNewGinRouter() {
	ginRouter := NewGinRouter()
	suite.Implementsf((*domain.Router)(nil), ginRouter, "GinRouter does not implement the domain.Router interface")
	suite.IsType((*GinRouter)(nil), ginRouter)
	suite.Equal(gin.ReleaseMode, gin.Mode())
}

func (suite *GinRouterSuite) TestUseMiddleware() {
	ginRouter := NewGinRouter()
	ginRouter.UseMiddleware(func(next http.Handler) http.Handler { return nil })
	suite.Equal(2, len(ginRouter.RouterGroup.Handlers))
}
