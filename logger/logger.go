package logger

import (
	"github.com/sirupsen/logrus"
	"os"
	"strings"
)

var logger = logrus.New()

func init() {
	setLogFormat()
	setLogLevel()
}

func setLogFormat() {
	hostname, _ := os.Hostname()

	// do not use package env to avoid import cycle
	format, ok := os.LookupEnv("LOGGER_FORMAT")
	if !ok {
		format = "json"
	}

	// do not use package env to avoid import cycle
	environment, ok := os.LookupEnv("ENVIRONMENT")
	if !ok {
		environment = "production"
	}

	f := customFormatter{
		additionalFields: map[string]string{
			"app":      os.Getenv("APP_NAME"),
			"env":      environment,
			"hostname": hostname,
		},
	}

	switch format {
	case "json":
		f.formatter = &logrus.JSONFormatter{}
	default:
		f.formatter = &logrus.TextFormatter{FullTimestamp: true}
	}
	logger.SetFormatter(f)
}

func setLogLevel() {
	// do not use package env to avoid import cycle
	level, ok := os.LookupEnv("LOGGER_LEVEL")
	if !ok {
		level = "debug"
	}
	switch strings.ToLower(level) {
	case "debug":
		logger.SetLevel(logrus.DebugLevel)
	case "info":
		logger.SetLevel(logrus.InfoLevel)
	case "warn":
		logger.SetLevel(logrus.WarnLevel)
	case "error":
		logger.SetLevel(logrus.ErrorLevel)
	default:
		logger.SetLevel(logrus.InfoLevel)
		logger.Warnf("invalid log level supplied: '%s'", level)
	}
}

func GetLogger() *logrus.Logger {
	return logger
}

func Fatalf(format string, args ...interface{}) {
	logger.Fatalf(format, args...)
}

func Fatal(args ...interface{}) {
	logger.Fatal(args...)
}

func Errorf(format string, args ...interface{}) {
	logger.Errorf(format, args...)
}

func Error(args ...interface{}) {
	logger.Error(args...)
}

func Warnf(format string, args ...interface{}) {
	logger.Warnf(format, args...)
}

func Warn(args ...interface{}) {
	logger.Warn(args...)
}

func Infof(format string, args ...interface{}) {
	logger.Infof(format, args...)
}

func Info(args ...interface{}) {
	logger.Info(args...)
}

func Debugf(format string, args ...interface{}) {
	logger.Debugf(format, args...)
}

func Debug(args ...interface{}) {
	logger.Debug(args...)
}
