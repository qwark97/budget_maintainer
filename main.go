package main

import (
	"fmt"
	"github.com/gorilla/handlers"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

const (
	HOST = "127.0.0.1"
	PORT = "9999"
)

func main() {
	router := mux.NewRouter()
	wrappedHandler := handlers.LoggingHandler(os.Stdout, router)

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
		Handler:      wrappedHandler,
		Addr:         addr,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Printf("serve at: %s\n", addr)
	log.Fatal(server.ListenAndServe())
}

func logSystemErr(err error) {
	log.Printf("ERROR - %s", err.Error())
}
