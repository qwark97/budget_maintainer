package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type Assets struct {
	Name   string `json:"name"`
	Amount int    `json:"amount"`
}

func addAsset(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	vars := mux.Vars(r)
	assetName, _ := vars["name"]
	err := saveAsset(assetName)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("cannot add such asset")
	}
}

func increaseAssets(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	vars := mux.Vars(r)
	assetName, _ := vars["name"]
	amount, _ := strconv.Atoi(r.URL.Query().Get("amount"))
	err := alterAsset(assetName, amount)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("cannot increase such asset")
	}
}

func decreaseAssets(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	vars := mux.Vars(r)
	assetName, _ := vars["name"]
	amount, _ := strconv.Atoi(r.URL.Query().Get("amount"))
	err := alterAsset(assetName, -amount)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("cannot decrease such asset")
	}
}

func fetchAsset(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	vars := mux.Vars(r)
	assetName, _ := vars["name"]
	if assets, err := loadAsset(assetName); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("cannot fetch such asset")
	} else {
		json.NewEncoder(w).Encode(assets)
	}
}

func fetchAssets(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	if assets, err := loadAssets(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("cannot fetch assets")
	} else {
		json.NewEncoder(w).Encode(assets)
	}
}

func removeAsset(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	vars := mux.Vars(r)
	assetName, _ := vars["name"]
	err := deleteAsset(assetName)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("cannot delete such asset")
	}
}
