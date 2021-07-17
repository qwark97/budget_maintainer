package main

import (
	"encoding/json"
	"github.com/qwark97/budget_maintainer/model"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

func addOperation(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	data, err := model.NewOperation(r.Body)
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

	if err := model.SaveOperation(data); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		logSystemErr(json.NewEncoder(w).Encode("cannot save operation into DB"))
		return
	} else {
		logSystemErr(json.NewEncoder(w).Encode("saved operation into DB"))
	}
}

func removeOperation(w http.ResponseWriter, r *http.Request) {
	var idToRemove int
	w.Header().Add("Content-Type", "application/json")
	vars := mux.Vars(r)
	id, _ := vars["id"]
	idToRemove, _ = strconv.Atoi(id)

	if err := model.DeleteOperation(idToRemove); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		logSystemErr(json.NewEncoder(w).Encode("cannot erase operation with passed id"))
	}
}

func fetchOperations(w http.ResponseWriter, _ *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	operations, err := model.LoadAllOperations()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		logSystemErr(json.NewEncoder(w).Encode("cannot load operations from DB"))
	} else {
		logSystemErr(json.NewEncoder(w).Encode(operations))
	}
}
