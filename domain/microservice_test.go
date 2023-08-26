package domain

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type MicroserviceSuite struct {
	suite.Suite
}

func TestMicroserviceSuite(t *testing.T) {
	suite.Run(t, new(MicroserviceSuite))
}

func (suite *MicroserviceSuite) TestMicroserviceInterface() {
	suite.Implementsf((*Microservice)(nil), new(fakeMicroservice), "Microservice interface is changed")
}
