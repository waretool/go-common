package logger

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/suite"
	"io"
	"os"
	"testing"
)

type LoggerSuite struct {
	suite.Suite
}

func (suite *LoggerSuite) SetupTest() {
	logger.SetOutput(io.Discard)
	gin.SetMode(gin.TestMode)
}

func TestLoggerSuite(t *testing.T) {
	suite.Run(t, new(LoggerSuite))
}

func (suite *LoggerSuite) TestLoggerDefaultFormat() {
	setLogFormat()

	logFormatter := logger.Formatter
	suite.Equal(&logrus.TextFormatter{}, logFormatter)
}

func (suite *LoggerSuite) TestLoggerJsonFormat() {
	suite.T().Setenv("ENVIRONMENT", "production")
	setLogFormat()

	logFormatter := logger.Formatter
	suite.Equal(&logrus.JSONFormatter{}, logFormatter)
}

func (suite *LoggerSuite) TestLoggerDefaultLevel() {
	setLogLevel()

	expectedLevel, _ := logrus.ParseLevel("debug")
	suite.Equal(expectedLevel, logger.GetLevel())
}

func (suite *LoggerSuite) TestLoggerDebugLevel() {
	suite.T().Setenv("LOGGER_LEVEL", "debug")
	setLogLevel()

	expectedLevel, _ := logrus.ParseLevel("debug")
	suite.Equal(expectedLevel, logger.GetLevel())
}

func (suite *LoggerSuite) TestLoggerInfoLevel() {
	suite.T().Setenv("LOGGER_LEVEL", "info")
	setLogLevel()

	expectedLevel, _ := logrus.ParseLevel("info")
	suite.Equal(expectedLevel, logger.GetLevel())
}

func (suite *LoggerSuite) TestLoggerWarnLevel() {
	suite.T().Setenv("LOGGER_LEVEL", "warn")
	setLogLevel()

	expectedLevel, _ := logrus.ParseLevel("warn")
	suite.Equal(expectedLevel, logger.GetLevel())
}

func (suite *LoggerSuite) TestLoggerErrorLevel() {
	suite.T().Setenv("LOGGER_LEVEL", "error")
	setLogLevel()

	expectedLevel, _ := logrus.ParseLevel("error")
	suite.Equal(expectedLevel, logger.GetLevel())
}

func (suite *LoggerSuite) TestLoggerInvalidLevel() {
	suite.T().Setenv("LOGGER_LEVEL", "not valid level")
	setLogLevel()

	expectedLevel, _ := logrus.ParseLevel("info")
	suite.Equal(expectedLevel, logger.GetLevel())
}

func (suite *LoggerSuite) TestGetLogger() {
	log := GetLogger()
	suite.Equal(logger, log)
}

func (suite *LoggerSuite) TestFatalf() {
	var buf bytes.Buffer
	logger.SetOutput(&buf)
	defer func() { logger.SetOutput(os.Stderr) }()

	isOsExitCalled := false
	defer func() { logger.ExitFunc = nil }()
	logger.ExitFunc = func(int) { isOsExitCalled = true }

	Fatalf("this is the fatal message")

	suite.Equal(true, isOsExitCalled)
	suite.Contains(buf.String(), "this is the fatal message")
}
