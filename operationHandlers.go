package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Operation struct {
	Category string  `json:"category"`
	Amount   float64 `json:"amount"`

	ID        int
	Timestamp int
}
type Operations []Operation

func addOperation(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	var data Operation
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("invalid request body")
		return
	}

	if err := saveOperation(data); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("cannot save operation into DB")
		return
	} else {
		json.NewEncoder(w).Encode("saved operation into DB")
	}
}

func removeOperation(w http.ResponseWriter, r *http.Request) {
	var idToRemove int
	w.Header().Add("Content-Type", "application/json")
	vars := mux.Vars(r)
	// Validation privided by mux
	id, _ := vars["id"]
	idToRemove, _ = strconv.Atoi(id)

	if err := eraseOperation(idToRemove); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("cannot erase operation with passed id")
	}
}

func fetchOperations(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	operations, err := loadAllOperations()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("cannot load operations from DB")
	} else {
		json.NewEncoder(w).Encode(operations)
	}
}
