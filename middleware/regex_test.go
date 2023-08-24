package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
)

type RegexSuite struct {
	suite.Suite
}

func (suite *RegexSuite) SetupTest() {
	gin.SetMode(gin.TestMode)
}

func TestRegexSuite(t *testing.T) {
	suite.Run(t, new(RegexSuite))
}

func (suite *RegexSuite) TestRegexMatchOk() {
	w := httptest.NewRecorder()
	_, r := gin.CreateTestContext(w)

	r.GET("/somePath/:type", MatchPattern("type", "(website)"), func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "the request has matched pattern (website)")
	})

	req := httptest.NewRequest(http.MethodGet, "/somePath/website", nil)

	r.ServeHTTP(w, req)

	suite.Equal(http.StatusOK, w.Code)
	suite.Equal("the request has matched pattern (website)", w.Body.String())
}

func (suite *RegexSuite) TestRegexMatchKo() {
	w := httptest.NewRecorder()
	_, r := gin.CreateTestContext(w)

	r.GET("/somePath/:type", MatchPattern("type", "(website)"), func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "this code is unreachable because unknown-value does not match the regex (website)")
	})

	req := httptest.NewRequest(http.MethodGet, "/somePath/unknown-value", nil)

	r.ServeHTTP(w, req)

	suite.Equal(http.StatusNotFound, w.Code)
	suite.Empty(w.Body.String())
}
