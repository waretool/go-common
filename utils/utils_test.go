package utils

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
)

type UtilSuite struct {
	suite.Suite
}

func (suite *UtilSuite) SetupTest() {
	gin.SetMode(gin.TestMode)
}

func TestUtilSuite(t *testing.T) {
	suite.Run(t, new(UtilSuite))
}

func (suite *UtilSuite) TestIsProduction() {
	var tests = []struct {
		environment string
		expected    bool
	}{
		{
			"production",
			true,
		},
		{
			"development",
			false,
		},
	}

	for _, test := range tests {
		suite.T().Setenv("ENVIRONMENT", test.environment)
		result := IsProduction()
		suite.Equal(test.expected, result)
	}

}
