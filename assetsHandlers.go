package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type Assets struct {
	Amount int `json:"amount"`
}

func increaseAssets(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	amount, _ := strconv.Atoi(r.URL.Query().Get("amount"))
	err := alterAssets(amount)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("cannot increase assets")
	}
}
func fetchAssets(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	if assets, err := loadAssets(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("cannot decrease assets")
	} else {
		json.NewEncoder(w).Encode(assets)
	}
}

func decreaseAssets(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	amount, _ := strconv.Atoi(r.URL.Query().Get("amount"))
	err := alterAssets(-amount)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("cannot decrease assets")
	}
}
