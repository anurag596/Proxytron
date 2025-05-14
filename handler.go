package main

import (
	"fmt"
	"net/http"
)

// HTTP handler function
func proxyHandler(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Query().Get("url")
	if url == "" {
		http.Error(w, "URL parameter missing", http.StatusBadRequest)
		return
	}

	// Fetch data from cache or origin
	response, err := fetchAndCache(url)
	if err != nil {
		http.Error(w, "Error fetching URL", http.StatusInternalServerError)
		return
	}

	// Return the response to the client
	fmt.Fprint(w, response)
}
