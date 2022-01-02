package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/qwark97/budget_maintainer/model"
)

func addAsset(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	assetName, _ := vars["name"]
	err := model.SaveAsset(assetName)
	if err != nil {
		logSystemErr(err)
		w.WriteHeader(http.StatusBadRequest)
		logSystemErr(json.NewEncoder(w).Encode("cannot add such asset"))
	}
}

func increaseAssets(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	assetName, _ := vars["name"]
	amount, _ := strconv.Atoi(r.URL.Query().Get("amount"))
	err := model.AlterAsset(assetName, amount)
	if err != nil {
		logSystemErr(err)
		w.WriteHeader(http.StatusBadRequest)
		logSystemErr(json.NewEncoder(w).Encode("cannot increase such asset"))
	}
}

func decreaseAssets(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	assetName, _ := vars["name"]
	amount, _ := strconv.Atoi(r.URL.Query().Get("amount"))
	err := model.AlterAsset(assetName, -amount)
	if err != nil {
		logSystemErr(err)
		w.WriteHeader(http.StatusBadRequest)
		logSystemErr(json.NewEncoder(w).Encode("cannot decrease such asset"))
	}
}

func fetchAsset(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	assetName, _ := vars["name"]
	if assets, err := model.LoadAsset(assetName); err != nil {
		logSystemErr(err)
		w.WriteHeader(http.StatusBadRequest)
		logSystemErr(json.NewEncoder(w).Encode("cannot fetch such asset"))
	} else {
		logSystemErr(json.NewEncoder(w).Encode(assets))
	}
}

func fetchAssets(w http.ResponseWriter, _ *http.Request) {
	if assets, err := model.LoadAssets(); err != nil {
		logSystemErr(err)
		w.WriteHeader(http.StatusBadRequest)
		logSystemErr(json.NewEncoder(w).Encode("cannot fetch assets"))
	} else {
		logSystemErr(json.NewEncoder(w).Encode(assets))
	}
}

func removeAsset(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	assetName, _ := vars["name"]
	err := model.DeleteAsset(assetName)
	if err != nil {
		logSystemErr(err)
		w.WriteHeader(http.StatusBadRequest)
		logSystemErr(json.NewEncoder(w).Encode("cannot delete such asset"))
	}
}
