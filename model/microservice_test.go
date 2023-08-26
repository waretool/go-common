package model

import (
	"github.com/stretchr/testify/suite"
	"github.com/waretool/go-common/domain"
	"testing"
)

type MicroserviceSuite struct {
	suite.Suite
}

func TestMicroserviceSuite(t *testing.T) {
	suite.Run(t, new(MicroserviceSuite))
}

func (suite *MicroserviceSuite) TestNewMicroservice() {
	mux := NewMicroservice("foo", NewGinRouter())
	gin := NewMicroservice("bar", NewMuxRouter())
	suite.Implements((*domain.Microservice)(nil), mux)
	suite.Implements((*domain.Microservice)(nil), gin)
}
