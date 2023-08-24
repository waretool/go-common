package env

import (
	"github.com/stretchr/testify/suite"
	"go-common/logger"
	"testing"
)

type EnvSuite struct {
	suite.Suite
}

func TestEnvSuite(t *testing.T) {
	suite.Run(t, new(EnvSuite))
}

func (suite *EnvSuite) TestGetEnvMessage() {
	suite.Equal("environment variable '%s' not found, using default value '%v'.", envNotFoundMessage)
	suite.Equal("environment variable '%s' cannot be converted to %s value due to: %s.", envCannotBeConvertedMessage)
}

func (suite *EnvSuite) TestGetEnv() {
	suite.T().Setenv("TEST_ENV", "value-of-test-env")
	environmentVar := GetEnv("TEST_ENV", "")
	suite.Equal("value-of-test-env", environmentVar)
}

func (suite *EnvSuite) TestGetEnvWithDefaultValue() {
	environmentVar := GetEnv("TEST_ENV", "default-value")
	suite.Equal("default-value", environmentVar)
}

func (suite *EnvSuite) TestGetEnvBool() {
	suite.T().Setenv("TEST_ENV", "true")
	environmentVar := GetEnv("TEST_ENV", false)
	suite.Equal(true, environmentVar)
}

func (suite *EnvSuite) TestGetEnvBoolFail() {
	suite.T().Setenv("TEST_ENV", "hello")

	defer func() { logger.GetLogger().ExitFunc = nil }()
	fatal := false
	logger.GetLogger().ExitFunc = func(int) { fatal = true }

	GetEnv("TEST_ENV", false)
	suite.Equal(true, fatal)
}

func (suite *EnvSuite) TestGetEnvInt() {
	suite.T().Setenv("TEST_ENV", "1")
	environmentVar := GetEnv("TEST_ENV", 0)
	suite.Equal(1, environmentVar)
}

func (suite *EnvSuite) TestGetEnvIntFail() {
	suite.T().Setenv("TEST_ENV", "hello")

	defer func() { logger.GetLogger().ExitFunc = nil }()
	fatal := false
	logger.GetLogger().ExitFunc = func(int) { fatal = true }

	GetEnv("TEST_ENV", 0)
	suite.Equal(true, fatal)
}
