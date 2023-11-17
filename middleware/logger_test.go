package middleware

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
	"github.com/waretool/go-common/logger"
	"net/http"
	"net/http/httptest"
	"testing"
)

type LoggerSuite struct {
	suite.Suite
}

func (suite *LoggerSuite) SetupTest() {
	gin.SetMode(gin.TestMode)
}

func TestLoggerSuite(t *testing.T) {
	suite.Run(t, new(LoggerSuite))
}

func (suite *LoggerSuite) TestLogger() {
	w := httptest.NewRecorder()
	_, r := gin.CreateTestContext(w)
	r.Use(LogMiddleware([]string{}))

	r.GET("/some-path", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "the request has matched pattern (website)")
	})

	oldOut := logger.GetLogger().Out
	buf := bytes.Buffer{}
	logger.GetLogger().SetOutput(&buf)

	req := httptest.NewRequest(http.MethodGet, "/some-path", nil)
	req.Header = http.Header{
		"Content-Type": {"application/json"},
		"Connection":   {"keep-alive"},
		"Cookie":       {"some-cookie=***"},
	}

	r.ServeHTTP(w, req)

	regex := `"app":"","duration".*,"env":"production","hostname":".*","level":"info","msg":"","request":\{"clientIp":".*","headers":\{"connection":"keep-alive","content-type":"application/json","cookie":"some-co\*\*\*","host":"example.com","user-agent":""\},"method":"GET","path":"/some-path"\},"status":200,"time":".*"`

	suite.Regexp(regex, buf.String())
	// restore log target
	logger.GetLogger().SetOutput(oldOut)
}

func (suite *LoggerSuite) TestLoggerSkipPath() {
	w := httptest.NewRecorder()
	_, r := gin.CreateTestContext(w)
	r.Use(LogMiddleware([]string{"/somePath"}))

	r.GET("/somePath", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "the request has matched pattern (website)")
	})

	oldOut := logger.GetLogger().Out
	buf := bytes.Buffer{}
	logger.GetLogger().SetOutput(&buf)

	req := httptest.NewRequest(http.MethodGet, "/somePath", nil)
	req.Header = http.Header{
		"Content-Type": {"application/json"},
		"Connection":   {"keep-alive"},
		"Cookie":       {"soc-session=***"},
	}

	r.ServeHTTP(w, req)

	suite.Equal("", buf.String())
	// restore log target
	logger.GetLogger().SetOutput(oldOut)
}
