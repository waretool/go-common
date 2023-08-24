package context

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
	"github.com/waretool/go-common/logger"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type ContextSuite struct {
	suite.Suite
}

func (suite *ContextSuite) SetupTest() {
	gin.SetMode(gin.TestMode)
}

func TestFormatterSuite(t *testing.T) {
	suite.Run(t, new(ContextSuite))
}

func (suite *ContextSuite) TestGetClientIpFromForwardedFrom() {
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)

	request := httptest.NewRequest(http.MethodGet, "/foo", nil)
	request.Header.Set("X-Forwarded-For", "123.123.123.123")
	ctx.Request = request

	ip := GetClientIP(ctx)

	suite.Equal("123.123.123.123", ip)
}

func (suite *ContextSuite) TestGetClientIpFromXRealIp() {
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)

	request := httptest.NewRequest(http.MethodGet, "/foo", nil)
	request.Header.Set("X-Real-IP", "123.123.123.123")
	ctx.Request = request

	ip := GetClientIP(ctx)

	suite.Equal("123.123.123.123", ip)
}

func (suite *ContextSuite) TestGetClientIpFromRemoteAddr() {
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = httptest.NewRequest(http.MethodGet, "/foo", nil)

	ip := GetClientIP(ctx)

	expectedIp := ctx.Request.RemoteAddr
	suite.Equal(expectedIp, ip)
}

func (suite *ContextSuite) TestGetDuration() {
	start := time.Now()
	time.Sleep(1 * time.Second)
	duration := GetDurationInMilliseconds(start)

	logger.Info(duration)
	suite.LessOrEqual(duration, float64(1003))
}
