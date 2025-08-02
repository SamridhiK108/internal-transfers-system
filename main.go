package main

import (
	"log"
	"net/http"

	"internaltransfer/db"
	"internaltransfer/handlers"

	"github.com/gorilla/mux"
)

func main() {
	db.InitDB("postgres://postgres:5601ds@localhost:5432/transferdb?sslmode=disable")

	router := mux.NewRouter()

	router.HandleFunc("/accounts", handlers.CreateAccount).Methods("POST")
	router.HandleFunc("/accounts/{id}", handlers.GetAccount).Methods("GET")
	router.HandleFunc("/transactions", handlers.CreateTransaction).Methods("POST")

	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
