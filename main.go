package main

import (
	"log"
	"net/http"
)

func main() {
	// Admin API routes
	http.HandleFunc("/api/v1/ad", createAd)

	// Public API routes
	http.HandleFunc("/api/v1/ad", listAds)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
