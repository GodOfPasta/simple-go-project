package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func OpenDBReadConnection() (*sql.DB, error) {
	// Hardcode for the testing sake
	// TODO: get creds from env
	connStr := "user=db_read password=admin dbname=postgres sslmode=disable"
	db, err := sql.Open("postgres", connStr)

	return db, err
}

func OpenDBWriteConnection() (*sql.DB, error) {
	// Hardcode for the testing sake
	// TODO: get creds from env
	connStr := "user=db_write password=admin dbname=postgres sslmode=disable"
	db, err := sql.Open("postgres", connStr)

	return db, err
}

func CloseDBConnection(db *sql.DB) {
	db.Close()
}
