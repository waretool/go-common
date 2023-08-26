package domain

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type JwtServiceSuite struct {
	suite.Suite
}

func TestJwtServiceSuite(t *testing.T) {
	suite.Run(t, new(JwtServiceSuite))
}

func (suite *JwtServiceSuite) TestJwtServiceInterface() {
	suite.Implementsf((*JwtService)(nil), new(dummyJwtService), "JwtService interface is changed")
}

type dummyJwtService struct{}

func (s *dummyJwtService) Generate(consumer Consumer) (string, error) {
	return "some token", nil
}

func (s *dummyJwtService) Valid(tokenString string) (*Claims, bool) {
	return &Claims{}, true
}
