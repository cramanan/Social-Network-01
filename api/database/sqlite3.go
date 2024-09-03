package database

import (
	"database/sql"
	"errors"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

const TransactionTimeout = 3 * time.Second

// A custom UNIQUE CONSTRAINT error
var ErrConflict = errors.New("Conflict")

// The SQLite3 Store is a sqlite3 DB connection with additionals transaction methods
type SQLite3Store struct{ *sql.DB }

func NewSQLite3Store() (*SQLite3Store, error) {
	db, err := sql.Open("sqlite3", "db/db.sqlite3")
	if err != nil {
		return nil, err
	}
	return &SQLite3Store{db}, nil
}
