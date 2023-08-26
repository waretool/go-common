// -tags test
package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
	"github.com/waretool/go-common/domain"
	"github.com/waretool/go-common/service"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"
)

type AuthNSuite struct {
	suite.Suite
	adminConsumer domain.Consumer
}

func (suite *AuthNSuite) SetupTest() {
	gin.SetMode(gin.TestMode)
	suite.adminConsumer = &fakeConsumer{
		Uuid:    "ac3c796f-0cdd-4cfd-9e9b-5655fee5f4df",
		Role:    domain.Admin,
		Enabled: true,
		Local:   true,
	}
}

func TestAuthNSuite(t *testing.T) {
	suite.Run(t, new(AuthNSuite))
}

func (suite *AuthNSuite) TestBearerRgx() {
	expected := regexp.MustCompile(`(?i)bearer `)
	suite.EqualValues(expected, bearerRgx)
}

func (suite *AuthNSuite) TestAuthNWithBearer() {
	suite.T().Setenv("JWT_SECRET_KEY", "super secret")
	jwtService := service.NewJwtService()

	validAdminToken, err := jwtService.Generate(suite.adminConsumer)
	if err != nil {
		suite.FailNow("cannot generate admin token to be used in test")
	}
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Add("Authorization", "Bearer "+validAdminToken)
	w := httptest.NewRecorder()

	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		val := r.Context().Value("claims")
		if val == nil {
			suite.FailNow("claims not present in jwt")
		}
		claims, ok := val.(*domain.Claims)
		if !ok {
			suite.FailNow("claims is not domain.Claims")
		}
		suite.EqualValues(suite.adminConsumer.GetRole(), claims.Role)
		suite.EqualValues(suite.adminConsumer.IsEnabled(), claims.Enabled)
		suite.EqualValues(suite.adminConsumer.IsLocal(), claims.Local)
		suite.NotZero(claims.ExpiresAt)
		w.WriteHeader(http.StatusOK)
	})

	handler := AuthN(nextHandler)
	handler.ServeHTTP(w, req)

	suite.Equal(http.StatusOK, w.Code)
}

func (suite *AuthNSuite) TestAuthNWithInvalidBearer() {
	suite.T().Setenv("JWT_SECRET_KEY", "super secret")

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Add("Authorization", "Bearer foobar")
	w := httptest.NewRecorder()

	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		suite.FailNow("this handler must not be reached but the request continued in the chain with invalid jwt")
	})

	handler := AuthN(nextHandler)
	handler.ServeHTTP(w, req)

	suite.Equal(http.StatusUnauthorized, w.Code)
}

func (suite *AuthNSuite) TestAuthNWithoutBearer() {
	suite.T().Setenv("JWT_SECRET_KEY", "super secret")

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()

	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		suite.FailNow("this handler must not be reached but the request continued in the chain without jwt")
	})

	handler := AuthN(nextHandler)
	handler.ServeHTTP(w, req)

	suite.Equal(http.StatusUnauthorized, w.Code)
}
