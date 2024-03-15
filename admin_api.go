package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func createAd(w http.ResponseWriter, r *http.Request) {
	var newAd Ad
	err := json.NewDecoder(r.Body).Decode(&newAd)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Validate input
	if newAd.Title == "" || newAd.StartAt.IsZero() || newAd.EndAt.IsZero() {
		http.Error(w, "Title, StartAt and EndAt are required fields", http.StatusBadRequest)
		return
	}

	// Add ad to the list
	ads = append(ads, newAd)

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Ad created successfully")
}
