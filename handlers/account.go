package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"internaltransfer/db"
	"internaltransfer/models"

	"github.com/gorilla/mux"
)

func CreateAccount(w http.ResponseWriter, r *http.Request) {
	var acc models.Account
	if err := json.NewDecoder(r.Body).Decode(&acc); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	_, err := db.DB.Exec("INSERT INTO accounts (id, balance) VALUES ($1, $2)", acc.ID, acc.Balance)
	if err != nil {
		log.Printf("Error inserting account: %v", err)
		http.Error(w, "Account creation failed", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func GetAccount(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid account ID", http.StatusBadRequest)
		return
	}

	var acc models.Account
	row := db.DB.QueryRow("SELECT id, balance FROM accounts WHERE id = $1", id)
	if err := row.Scan(&acc.ID, &acc.Balance); err != nil {
		http.Error(w, "Account not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(acc)
}
