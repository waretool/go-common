package utils

import "github.com/waretool/go-common/env"

func IsProduction() bool {
	if env.GetEnv("ENVIRONMENT", "production") == "production" {
		return true
	}
	return false
}
