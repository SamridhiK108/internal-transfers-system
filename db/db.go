package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB(dataSource string) {
	var err error
	DB, err = sql.Open("postgres", dataSource)
	if err != nil {
		log.Fatal(err)
	}

	if err := DB.Ping(); err != nil {
		log.Fatal(err)
	}

	createTables()
}

func createTables() {
	query := `
    CREATE TABLE IF NOT EXISTS accounts (
        id BIGINT PRIMARY KEY,
        balance NUMERIC(20, 5) NOT NULL
    );
    CREATE TABLE IF NOT EXISTS transactions (
        id SERIAL PRIMARY KEY,
        source_account_id BIGINT,
        destination_account_id BIGINT,
        amount NUMERIC(20, 5),
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );`
	_, err := DB.Exec(query)
	if err != nil {
		log.Fatal("Error creating tables:", err)
	}
}
