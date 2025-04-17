package db

import (
	"database/sql"
	"os"

	_ "github.com/lib/pq"
)

func OpenDBReadConnection() (*sql.DB, error) {
	connStr := os.Getenv("DB_READ_URI")
	db, err := sql.Open("postgres", connStr)

	return db, err
}

func OpenDBWriteConnection() (*sql.DB, error) {
	connStr := os.Getenv("DB_WRITE_URI")
	db, err := sql.Open("postgres", connStr)

	return db, err
}

func CloseDBConnection(db *sql.DB) {
	db.Close()
}
