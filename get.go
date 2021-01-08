package env

import (
	"os"
	"strconv"
)

// GetString func returns environment variable value as a string value,
// If variable doesn't exist or is not set, returns fallback value
func GetString(key string, fallback string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return fallback
	}
	return value
}

// GetBool func returns environment variable value as a boolean value,
// If variable doesn't exist or is not set, returns fallback value
func GetBool(key string, fallback bool) bool {
	value, exists := os.LookupEnv(key)
	if !exists {
		return fallback
	}
	if value == "true" || value == "1" {
		return true
	}
	return false
}

// GetInt func returns environment variable value as a integer value,
// If variable doesn't exist or is not set, returns fallback value
func GetInt(key string, fallback int) int {
	value, exists := os.LookupEnv(key)
	if !exists {
		return fallback
	}
	res, err := strconv.ParseInt(value, 10, 32)
	if err != nil {
		return fallback
	}
	return int(res)
}

// GetFloat func returns environment variable value as a float value,
// If variable doesn't exist or is not set, returns fallback value
func GetFloat(key string, fallback float64) float64 {
	value, exists := os.LookupEnv(key)
	if !exists {
		return fallback
	}
	res, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return fallback
	}
	return res
}
