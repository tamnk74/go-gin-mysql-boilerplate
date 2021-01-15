package config

import (
	"os"
)

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}

var PORT = getEnv("PORT", "8000")
var DB_USERNAME = getEnv("DB_USERNAME", "root")
var DB_PASSWORD = getEnv("DB_PASSWORD", "")
var DB_HOST = getEnv("DB_HOST", "127.0.0.1")
var DB_PORT = getEnv("DB_PORT", "3306")
var DB_NAME = getEnv("DB_NAME", "test")
