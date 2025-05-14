package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	initRedis()

	// Register the handler
	http.HandleFunc("/", proxyHandler)

	// Start the server
	fmt.Println("Proxy server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}