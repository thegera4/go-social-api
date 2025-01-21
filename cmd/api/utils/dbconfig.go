package utils

import (
	"os"
	"strconv"
)

// Utility function to load the configuration settings from the environment variables for the "Addr". Returns ":8080" if no value is found.
func LoadAddr() string {
	env := os.Getenv("ADDR")
	if env == "" { return ":8080" }
	return env
}

// Utility function to load the configuration settings from the environment variables for the "DB_ADDR". Returns a default connection string if no value is found.
func LoadDBAddr() string {
	env := os.Getenv("DB_ADDR")
	if env == "" { return "postgresql://admin:adminpassword@localhost:5432/social?sslmode=disable" }
	return env
}

// Utility function to load the configuration settings from the environment variables for the "DB_MAX_OPEN_CONNS". Returns 30 if no value is found.
func LoadMaxOpenConns() int {
	env := os.Getenv("DB_MAX_OPEN_CONNS")
	maxOpenConns, err := strconv.Atoi(env)
	if err != nil || env == "" { return 30 }
	return maxOpenConns
}

// Utility function to load the configuration settings from the environment variables for the "DB_MAX_IDLE_CONNS". Returns 30 if no value is found.
func LoadMaxIdleConns() int {
	env := os.Getenv("DB_MAX_IDLE_CONNS")
	maxIdleConns, err := strconv.Atoi(env)
	if err != nil || env == "" { return 30 }
	return maxIdleConns
}

// Utility function to load the configuration settings from the environment variables for the "DB_MAX_IDLE_TIME". Returns "15m" if no value is found.
func LoadMaxIdleTime() string {
	env := os.Getenv("DB_MAX_IDLE_TIME")
	if env == "" { return "15m" }
	return env
}