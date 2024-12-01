package util

import "os"

func GetEnvOr(key, defaultVaue string) string {
	if envValue := os.Getenv(key); envValue != "" {
		return envValue
	}

	return defaultVaue
}
