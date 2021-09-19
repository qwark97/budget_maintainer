package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/qwark97/budget_maintainer/model"
	"net/http"
	"strconv"
	"time"
)

func addBudget(w http.ResponseWriter, r *http.Request) {
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
	vars := mux.Vars(r)
	year, _ := vars["year"]
	month, _ := vars["month"]

	yearToFetch, _ := strconv.Atoi(year)
	if yearToFetch < 2000 || yearToFetch > 2100 {
		logSystemErr(json.NewEncoder(w).Encode("invalid request body"))
		return
	}
	monthToFetch, _ := strconv.Atoi(month)
	if monthToFetch < 1 || monthToFetch > 12 {
		logSystemErr(json.NewEncoder(w).Encode("invalid request body"))
		return
	}

	budget, err := model.LoadBudget(yearToFetch, monthToFetch)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		logSystemErr(json.NewEncoder(w).Encode("cannot load budget from DB"))
	} else {
		logSystemErr(json.NewEncoder(w).Encode(budget))
	}
}
