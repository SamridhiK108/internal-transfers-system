package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"internaltransfer/db"
	"internaltransfer/models"
	"net/http"
)

func CreateTransaction(w http.ResponseWriter, r *http.Request) {
	var txn models.Transaction
	if err := json.NewDecoder(r.Body).Decode(&txn); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	tx, err := db.DB.Begin()
	if err != nil {
		http.Error(w, "Failed to start transaction", http.StatusInternalServerError)
		return
	}

	defer tx.Rollback()

	var srcBalance, dstBalance float64
	err = tx.QueryRow("SELECT balance FROM accounts WHERE id = $1 FOR UPDATE", txn.SourceAccountID).Scan(&srcBalance)
	if err == sql.ErrNoRows {
		http.Error(w, "Source account not found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, "Error reading source account", http.StatusInternalServerError)
		return
	}

	err = tx.QueryRow("SELECT balance FROM accounts WHERE id = $1 FOR UPDATE", txn.DestinationAccountID).Scan(&dstBalance)
	if err == sql.ErrNoRows {
		http.Error(w, "Destination account not found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, "Error reading destination account", http.StatusInternalServerError)
		return
	}

	var amount float64
	if _, err := fmt.Sscanf(txn.Amount, "%f", &amount); err != nil || amount <= 0 {
		http.Error(w, "Invalid amount", http.StatusBadRequest)
		return
	}

	if srcBalance < amount {
		http.Error(w, "Insufficient funds", http.StatusBadRequest)
		return
	}

	_, err = tx.Exec("UPDATE accounts SET balance = balance - $1 WHERE id = $2", amount, txn.SourceAccountID)
	if err != nil {
		http.Error(w, "Failed to debit", http.StatusInternalServerError)
		return
	}

	_, err = tx.Exec("UPDATE accounts SET balance = balance + $1 WHERE id = $2", amount, txn.DestinationAccountID)
	if err != nil {
		http.Error(w, "Failed to credit", http.StatusInternalServerError)
		return
	}

	_, err = tx.Exec("INSERT INTO transactions (source_account_id, destination_account_id, amount) VALUES ($1, $2, $3)",
		txn.SourceAccountID, txn.DestinationAccountID, amount)
	if err != nil {
		http.Error(w, "Failed to log transaction", http.StatusInternalServerError)
		return
	}

	if err := tx.Commit(); err != nil {
		http.Error(w, "Transaction failed", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
