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
