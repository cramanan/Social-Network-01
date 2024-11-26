package database

import (
	"database/sql"
	"errors"

	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/sqlite3"
	"github.com/golang-migrate/migrate/source/file"
	_ "github.com/mattn/go-sqlite3"
)

// A custom UNIQUE CONSTRAINT error
var ErrConflict = errors.New("Conflict")

// The SQLite3 Store is a sqlite3 DB connection with additionals transaction methods
type SQLite3Store struct{ *sql.DB }

func NewSQLite3Store(dbFilePath string) (*SQLite3Store, error) {
	db, err := sql.Open("sqlite3", dbFilePath)
	if err != nil {
		return nil, err
	}

	return &SQLite3Store{db}, nil
}

func (store *SQLite3Store) Up(migrationDir string) error {
	instance, err := sqlite3.WithInstance(store.DB, new(sqlite3.Config)) // TODO: refactor
	if err != nil {
		return err
	}

	fSrc, err := new(file.File).Open(migrationDir)
	if err != nil {
		return err
	}

	m, err := migrate.NewWithInstance("file", fSrc, "sqlite3", instance)
	if err != nil {
		return err
	}

	return m.Up()
}

func (store *SQLite3Store) Down(migrationDir string) error {
	instance, err := sqlite3.WithInstance(store.DB, new(sqlite3.Config))
	if err != nil {
		return err
	}

	fSrc, err := new(file.File).Open(migrationDir)
	if err != nil {
		return err
	}

	m, err := migrate.NewWithInstance("file", fSrc, "sqlite3", instance)
	if err != nil {
		return err
	}

	return m.Down()
}
