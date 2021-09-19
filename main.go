package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/handlers"
	"github.com/qwark97/budget_maintainer/model"

	"github.com/gorilla/mux"
)

var (
	HOST         string
	PORT         string
	dbConf       model.DBConf
	ALLOW_ORIGIN = "*"
)

func main() {
	loadConfiguration()
	model.InitDatabase(dbConf)
	db, _ := model.DBConn.DB()
	defer db.Close()
	router := mux.NewRouter()
	wrappedHandler := handlers.LoggingHandler(os.Stdout, router)

	router.HandleFunc("/api/categories/{name:[a-zA-Z0-9]+}", addCategory).Methods("POST")
	router.HandleFunc("/api/categories", fetchCategories).Methods("GET")
	router.HandleFunc("/api/categories/{name:[a-zA-Z0-9]+}", removeCategory).Methods("DELETE")

	router.HandleFunc("/api/operations", addOperation).Methods("POST")
	router.HandleFunc("/api/operations", fetchOperations).Methods("GET")
	router.HandleFunc("/api/operations/{id:[0-9]+}", removeOperation).Methods("DELETE")

	router.HandleFunc("/api/assets/{name:[a-zA-Z0-9]+}", addAsset).Methods("POST")
	router.HandleFunc("/api/assets/{name:[a-zA-Z0-9]+}/increase", increaseAssets).Queries("amount", "{amount:[0-9]+}").Methods("PUT")
	router.HandleFunc("/api/assets/{name:[a-zA-Z0-9]+}/decrease", decreaseAssets).Queries("amount", "{amount:[0-9]+}").Methods("PUT")
	router.HandleFunc("/api/assets/{name:[a-zA-Z0-9]+}", fetchAsset).Methods("GET")
	router.HandleFunc("/api/assets", fetchAssets).Methods("GET")
	router.HandleFunc("/api/assets/{name:[a-zA-Z0-9]+}", removeAsset).Methods("DELETE")

	router.HandleFunc("/api/budget", addBudget).Methods("POST")
	router.HandleFunc("/api/budget/{year:[0-9]+}/{month:[0-9]+}", fetchBudget).Methods("GET")

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
	if err != nil {
		log.Printf("ERROR - %s", err.Error())
	}
}

func loadConfiguration() {
	HOST = os.Getenv("BM_HOST")
	if HOST == "" {
		HOST = "127.0.0.1"
	}

	PORT = os.Getenv("BM_PORT")
	if PORT == "" {
		PORT = "9999"
	}

	dbHost := os.Getenv("BM_DB_HOST")
	if dbHost == "" {
		dbHost = "postgres"
	}
	dbConf.Host = dbHost

	dbPort := os.Getenv("BM_DB_PORT")
	if dbPort == "" {
		dbPort = "5432"
	}
	dbConf.Port = dbPort

	dbUser := os.Getenv("BM_DB_USER")
	if dbUser == "" {
		dbUser = "postgres"
	}
	dbConf.User = dbUser

	dbConf.Pass = os.Getenv("BM_DB_PASS")

	dbName := os.Getenv("BM_DB_NAME")
	if dbName == "" {
		dbName = "postgres"
	}
	dbConf.Name = dbName

	dbTZ := os.Getenv("TZ")
	if dbTZ == "" {
		dbTZ = "Europe/Warsaw"
	}
	dbConf.TZ = dbTZ
}
