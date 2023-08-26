package domain

import (
	"github.com/stretchr/testify/suite"
	"net/http"
	"testing"
)

type RouterSuite struct {
	suite.Suite
}

func TestRouterSuite(t *testing.T) {
	suite.Run(t, new(RouterSuite))
}

func (suite *RouterSuite) TestRouterInterface() {
	suite.Implementsf((*Router)(nil), new(dummyRouter), "Router interface is changed")
}

type dummyRouter struct{}

func (r *dummyRouter) ServeHTTP(http.ResponseWriter, *http.Request) {}

func (r *dummyRouter) UseMiddleware(middleware Middleware) {}
