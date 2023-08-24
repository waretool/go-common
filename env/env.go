package env

import (
	"go-common/logger"
	"os"
	"strconv"
)

const (
	envNotFoundMessage          = "environment variable '%s' not found, using default value '%v'."
	envCannotBeConvertedMessage = "environment variable '%s' cannot be converted to %s value due to: %s."
)

func GetEnv[T string | bool | int](key string, fallback T) T {
	value, ok := os.LookupEnv(key)
	if !ok {
		logger.Warnf(envNotFoundMessage, key, fallback)
		return fallback
	}

	var ret T
	switch p := any(&ret).(type) {
	case *string:
		*p = value
	case *bool:
		i, err := strconv.ParseBool(value)
		if err != nil {
			logger.Fatalf(envCannotBeConvertedMessage, key, "boolean", err.Error())
		}
		*p = bool(i)
	case *int:
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			logger.Fatalf(envCannotBeConvertedMessage, key, "int", err.Error())
		}
		*p = int(i)
	}
	return ret
}
