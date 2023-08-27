package utils

import (
	"github.com/waretool/go-common/logger"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

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

func (suite *UtilSuite) TestGetClientIpFromForwardedFrom() {
	request := httptest.NewRequest(http.MethodGet, "/foo", nil)
	request.Header.Set("X-Forwarded-For", "123.123.123.123")

	ip := GetClientIp(request)

	suite.Equal("123.123.123.123", ip)
}

func (suite *UtilSuite) TestGetClientIpFromXRealIp() {
	request := httptest.NewRequest(http.MethodGet, "/foo", nil)
	request.Header.Set("X-Real-IP", "123.123.123.123")

	ip := GetClientIp(request)

	suite.Equal("123.123.123.123", ip)
}

func (suite *UtilSuite) TestGetClientIpFromRemoteAddr() {
	request := httptest.NewRequest(http.MethodGet, "/foo", nil)

	ip := GetClientIp(request)

	expectedIp := request.RemoteAddr
	suite.Equal(expectedIp, ip)
}

func (suite *UtilSuite) TestGetDuration() {
	start := time.Now()
	time.Sleep(1 * time.Second)
	duration := GetDurationInMilliseconds(start)

	logger.Info(duration)
	suite.LessOrEqual(duration, float64(1003))
}
