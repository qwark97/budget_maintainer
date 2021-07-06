package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

const (
	HOST = "127.0.0.1"
	PORT = "9999"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/api/operations", addOperation).Methods("POST")
	router.HandleFunc("/api/operations", fetchOperations).Methods("GET")
	router.HandleFunc("/api/operations/{id:[0-9]+}", removeOperation).Methods("DELETE")

	router.HandleFunc("/api/assets/{name:[a-zA-Z0-9]+}", addAsset).Methods("POST")
	router.HandleFunc("/api/assets/{name:[a-zA-Z0-9]+}/increase", increaseAssets).Queries("amount", "{amount:[0-9]+}").Methods("PUT")
	router.HandleFunc("/api/assets/{name:[a-zA-Z0-9]+}/decrease", decreaseAssets).Queries("amount", "{amount:[0-9]+}").Methods("PUT")
	router.HandleFunc("/api/assets/{name:[a-zA-Z0-9]+}", fetchAsset).Methods("GET")
	router.HandleFunc("/api/assets", fetchAssets).Methods("GET")
	router.HandleFunc("/api/assets/{name:[a-zA-Z0-9]+}", removeAsset).Methods("DELETE")

	addr := fmt.Sprintf("%s:%s", HOST, PORT)
	server := &http.Server{
		Handler:      router,
		Addr:         addr,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Printf("serve at: %s\n", addr)
	log.Fatal(server.ListenAndServe())
}
