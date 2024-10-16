package database

import (
	"database/sql"
	"errors"
	"log"
	"time"

	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/sqlite3"
	"github.com/golang-migrate/migrate/source/file"
	_ "github.com/mattn/go-sqlite3"
)

const TransactionTimeout = 3 * time.Second

// A custom UNIQUE CONSTRAINT error
var ErrConflict = errors.New("Conflict")

// The SQLite3 Store is a sqlite3 DB connection with additionals transaction methods
type SQLite3Store struct{ *sql.DB }

func NewSQLite3Store(dbFilePath string) (*SQLite3Store, error) {
	db, err := sql.Open("sqlite3", dbFilePath)
	if err != nil {
		return nil, err
	}

	instance, err := sqlite3.WithInstance(db, &sqlite3.Config{})
	if err != nil {
		return nil, err
	}

	fSrc, err := (&file.File{}).Open("api/db/migrations")
	if err != nil {
		log.Fatal(err)
	}

	m, err := migrate.NewWithInstance("file", fSrc, "sqlite3", instance)
	if err != nil {
		log.Fatal(err)
	}

	// modify for Down
	if err := m.Up(); err != nil {
		log.Fatal(err)
	}

	return &SQLite3Store{db}, nil
}
