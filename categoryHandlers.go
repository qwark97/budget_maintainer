package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/qwark97/budget_maintainer/model"
	"net/http"
)

func addCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.Header().Add("Access-Control-Allow-Origin", ALLOW_ORIGIN)
	vars := mux.Vars(r)
	categoryName, _ := vars["name"]

	if err := model.SaveCategory(categoryName); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		logSystemErr(json.NewEncoder(w).Encode("cannot save category into DB"))
		return
	} else {
		logSystemErr(json.NewEncoder(w).Encode("saved category into DB"))
	}
}

func removeCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.Header().Add("Access-Control-Allow-Origin", ALLOW_ORIGIN)
	vars := mux.Vars(r)
	categoryName, _ := vars["name"]

	if err := model.DeleteCategory(categoryName); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		logSystemErr(json.NewEncoder(w).Encode("cannot erase category with passed name"))
	}
}

func fetchCategories(w http.ResponseWriter, _ *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.Header().Add("Access-Control-Allow-Origin", ALLOW_ORIGIN)
	categories, err := model.LoadAllCategories()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		logSystemErr(json.NewEncoder(w).Encode("cannot load categories from DB"))
	} else {
		logSystemErr(json.NewEncoder(w).Encode(categories.Elements))
	}
}
