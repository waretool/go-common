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

	hostname, _ := os.Hostname()
	expected := customFormatter{
		additionalFields: map[string]string{
			"app":      os.Getenv("APP_NAME"),
			"env":      "production",
			"hostname": hostname,
		},
		formatter: &logrus.JSONFormatter{},
	}

	logFormatter := logger.Formatter
	suite.Equal(expected, logFormatter)
}

func (suite *LoggerSuite) TestLoggerTextFormat() {
	suite.T().Setenv("LOGGER_FORMAT", "text")
	setLogFormat()

	hostname, _ := os.Hostname()
	expected := customFormatter{
		additionalFields: map[string]string{
			"app":      os.Getenv("APP_NAME"),
			"env":      "production",
			"hostname": hostname,
		},
		formatter: &logrus.TextFormatter{FullTimestamp: true},
	}

	logFormatter := logger.Formatter
	suite.Equal(expected, logFormatter)
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

func (suite *LoggerSuite) TestFatal() {
	var buf bytes.Buffer
	logger.SetOutput(&buf)
	defer func() { logger.SetOutput(os.Stderr) }()

	isOsExitCalled := false
	defer func() { logger.ExitFunc = nil }()
	logger.ExitFunc = func(int) { isOsExitCalled = true }

	Fatal("this is the fatal message")

	suite.Equal(true, isOsExitCalled)
	suite.Contains(buf.String(), "this is the fatal message")
}

func (suite *LoggerSuite) TestErrorf() {
	var buf bytes.Buffer
	logger.SetOutput(&buf)
	defer func() { logger.SetOutput(os.Stderr) }()

	Errorf("this is the error message")

	suite.Contains(buf.String(), "this is the error message")
}

func (suite *LoggerSuite) TestError() {
	var buf bytes.Buffer
	logger.SetOutput(&buf)
	defer func() { logger.SetOutput(os.Stderr) }()

	Error("this is the error message")

	suite.Contains(buf.String(), "this is the error message")
}

func (suite *LoggerSuite) TestWarnf() {
	var buf bytes.Buffer
	logger.SetOutput(&buf)
	defer func() { logger.SetOutput(os.Stderr) }()

	Warnf("this is the warn message")

	suite.Contains(buf.String(), "this is the warn message")
}

func (suite *LoggerSuite) TestWarn() {
	var buf bytes.Buffer
	logger.SetOutput(&buf)
	defer func() { logger.SetOutput(os.Stderr) }()

	Warn("this is the warn message")

	suite.Contains(buf.String(), "this is the warn message")
}

func (suite *LoggerSuite) TestInfof() {
	var buf bytes.Buffer
	logger.SetOutput(&buf)
	defer func() { logger.SetOutput(os.Stderr) }()

	Infof("this is the info message")

	suite.Contains(buf.String(), "this is the info message")
}

func (suite *LoggerSuite) TestInfo() {
	var buf bytes.Buffer
	logger.SetOutput(&buf)
	defer func() { logger.SetOutput(os.Stderr) }()

	Info("this is the info message")

	suite.Contains(buf.String(), "this is the info message")
}

func (suite *LoggerSuite) TestDebugf() {
	var buf bytes.Buffer
	logger.SetOutput(&buf)
	defer func() { logger.SetOutput(os.Stderr) }()

	Debugf("this is the debug message")

	suite.Contains(buf.String(), "this is the debug message")
}

func (suite *LoggerSuite) TestDebug() {
	var buf bytes.Buffer
	logger.SetOutput(&buf)
	defer func() { logger.SetOutput(os.Stderr) }()

	Debug("this is the debug message")

	suite.Contains(buf.String(), "this is the debug message")
}
