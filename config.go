package main

import (
	"os"
)

// this retrieves the api key from environment variables (future improvement)
func GetAPIKey() string {
	return os.Getenv("PROXY_API_KEY")
}