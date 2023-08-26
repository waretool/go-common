package service

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/suite"
	"github.com/waretool/go-common/domain"
	"github.com/waretool/go-common/logger"
	"os"
	"testing"
	"time"
)

type JwtSuite struct {
	suite.Suite
	consumer domain.Consumer
}

func (suite *JwtSuite) SetupTest() {
	suite.consumer = &fakeConsumer{
		Uuid:    "453e303d-0454-4c0d-8b1f-05e7bb34e679",
		Local:   false,
		Role:    3,
		Enabled: false,
	}
}

func TestJwtSuite(t *testing.T) {
	suite.Run(t, new(JwtSuite))
}

func (suite *JwtSuite) TestNewJwtService() {
	suite.T().Setenv("JWT_SECRET_KEY", "super secret")
	jwtService := NewJwtService()
	suite.Implementsf((*domain.JwtService)(nil), jwtService, "NewJwtService does not return an instance of the JwtService interface")
}

func (suite *JwtSuite) TestNewJwtServiceWithoutEnvJwtSecretKey() {
	log := logger.GetLogger()
	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer func() { log.SetOutput(os.Stderr) }()

	isOsExitCalled := false
	defer func() { log.ExitFunc = nil }()
	log.ExitFunc = func(int) { isOsExitCalled = true }

	NewJwtService() // this method calls Fatalf

	suite.Equal(true, isOsExitCalled)
	suite.Contains(buf.String(), "env variable JWT_SECRET_KEY is missing but is mandatory")
}

func (suite *JwtSuite) TestJwtServiceGenerate() {
	suite.T().Setenv("JWT_SECRET_KEY", "super secret")
	jwtService := NewJwtService()
	tokenString, err := jwtService.Generate(suite.consumer)
	if err != nil {
		suite.T().Fatalf("cannot generate jwt")
	}

	token, _ := jwt.ParseWithClaims(tokenString, &domain.Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("super secret"), nil
	})

	if claims, ok := token.Claims.(*domain.Claims); ok && token.Valid {
		suite.Equal(suite.consumer.GetIdentifier(), claims.Subject)
		suite.Equal(suite.consumer.GetRole(), claims.Role)
		suite.Equal(suite.consumer.IsEnabled(), claims.Enabled)
		suite.Equal(suite.consumer.IsLocal(), claims.Local)
		suite.LessOrEqual(time.Now().Add(900*time.Second).Unix(), claims.ExpiresAt)
	} else {
		suite.T().Fatal("the token does not contain custom claims or is not valid")
	}
}

func (suite *JwtSuite) TestJwtServiceGenerateWithNilConsumer() {
	suite.T().Setenv("JWT_SECRET_KEY", "super secret")
	jwtService := NewJwtService()
	_, err := jwtService.Generate(nil)
	suite.Equal(err, errors.New("unable to generate a jwt from a nil consumer"))
}

func (suite *JwtSuite) TestJwtServiceValid() {
	suite.T().Setenv("JWT_SECRET_KEY", "super secret")
	jwtService := NewJwtService()
	tokenString, err := jwtService.Generate(suite.consumer)
	if err != nil {
		suite.T().Fatal("cannot generate jwt")
	}

	claims, valid := jwtService.Valid(tokenString)
	suite.Equal(true, valid)
	suite.Equal(suite.consumer.GetIdentifier(), claims.Subject)
	suite.Equal(suite.consumer.GetRole(), claims.Role)
	suite.Equal(suite.consumer.IsEnabled(), claims.Enabled)
	suite.Equal(suite.consumer.IsLocal(), claims.Local)
}

func (suite *JwtSuite) TestJwtServiceValidWithInvalidToken() {
	suite.T().Setenv("JWT_SECRET_KEY", "super secret")
	jwtService := NewJwtService()
	tokenString := "this is invalid jwt"

	claims, valid := jwtService.Valid(tokenString)
	suite.Equal(false, valid)
	suite.Nil(claims)
}

func (suite *JwtSuite) TestJwtServiceValidWithInvalidSignMethod() {
	suite.T().Setenv("JWT_SECRET_KEY", "super secret")
	jwtService := NewJwtService()
	tokenString := "eyJhbGciOiJFUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWUsImlhdCI6MTUxNjIzOTAyMn0.tyh-VfuzIxCyGYDlkBA7DfyjrqmSHu6pQ2hoZuFqUSLPNY2N0mpHb3nk5K17HWP_3cYHBw7AhHale5wky6-sVA"

	claims, valid := jwtService.Valid(tokenString)
	suite.Equal(false, valid)
	suite.Nil(claims)
}
