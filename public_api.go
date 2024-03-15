package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"
)

func listAds(w http.ResponseWriter, r *http.Request) {
	age, _ := strconv.Atoi(r.URL.Query().Get("age"))
	gender := r.URL.Query().Get("gender")
	country := r.URL.Query()["country"]
	platform := r.URL.Query()["platform"]
	offset, _ := strconv.Atoi(r.URL.Query().Get("offset"))
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))

	var filteredAds []Ad
	now := time.Now()

	for _, ad := range ads {
		if ad.StartAt.Before(now) && ad.EndAt.After(now) &&
			(age == 0 || (age >= ad.Conditions.AgeStart && age <= ad.Conditions.AgeEnd)) &&
			(gender == "" || gender == ad.Conditions.Gender) &&
			(len(country) == 0 || contains(ad.Conditions.Country, country)) &&
			(len(platform) == 0 || contains(ad.Conditions.Platform, platform)) {
			filteredAds = append(filteredAds, ad)
		}
	}

	// Pagination
	start := offset - 1
	end := start + limit
	if start < 0 {
		start = 0
	}
	if end > len(filteredAds) {
		end = len(filteredAds)
	}

	response := map[string][]Ad{"items": filteredAds[start:end]}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func contains(s []string, e []string) bool {
	for _, a := range e {
		for _, b := range s {
			if b == a {
				return true
			}
		}
	}
	return false
}
