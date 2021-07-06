package main

import (
	"encoding/json"
	"github.com/qwark97/budget_maintainer/model"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func addOperation(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	var data model.Operation
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
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

func fetchOperations(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	operations, err := model.LoadAllOperations()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		logSystemErr(json.NewEncoder(w).Encode("cannot load operations from DB"))
	} else {
		logSystemErr(json.NewEncoder(w).Encode(operations))
	}
}
