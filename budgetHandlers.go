package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/qwark97/budget_maintainer/model"
	"net/http"
	"time"
)

func addBudget(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	data, err := model.NewTransitBudget(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		logSystemErr(json.NewEncoder(w).Encode("invalid request body"))
		return
	}

	if data.Year == 0 {
		data.Year = time.Now().Year()
	} else if data.Year < 0 {
		logSystemErr(json.NewEncoder(w).Encode("invalid request body"))
		return
	}
	if data.Month == 0 {
		data.Month = time.Now().Month()
	} else if data.Month < 0 {
		logSystemErr(json.NewEncoder(w).Encode("invalid request body"))
		return
	}

	if err := model.SetBudget(data); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		logSystemErr(json.NewEncoder(w).Encode("cannot save budget position into DB"))
		return
	} else {
		logSystemErr(json.NewEncoder(w).Encode("saved budget position into DB"))
	}
}

func fetchBudget(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	vars := mux.Vars(r)
	year, _ := vars["month"]
	month, _ := vars["month"]
	budget, err := model.LoadBudget(year, month)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		logSystemErr(json.NewEncoder(w).Encode("cannot load budget from DB"))
	} else {
		logSystemErr(json.NewEncoder(w).Encode(budget))
	}
}
