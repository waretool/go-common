package domain

import (
	"github.com/stretchr/testify/suite"
	"github.com/waretool/go-common/domain"
	"testing"
)

type MuxRouterSuite struct {
	suite.Suite
}

func TestMuxRouterSuite(t *testing.T) {
	suite.Run(t, new(MuxRouterSuite))
}

func (suite *MuxRouterSuite) TestMuxRouterImplementsRouter() {
	suite.Implementsf((*domain.Router)(nil), new(MuxRouter), "MuxRouter does not implement the domain.Router interface")
}

func (suite *MuxRouterSuite) TestNewMuxRouter() {
	muxRouter := NewMuxRouter()
	suite.Implementsf((*domain.Router)(nil), muxRouter, "MuxRouter does not implement the domain.Router interface")
	suite.IsType((*MuxRouter)(nil), muxRouter)
}

func (suite *MuxRouterSuite) TestUseMiddleware() {
	// todo implement test
}
