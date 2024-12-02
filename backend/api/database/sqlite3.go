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

// NewSQLite3Store initializes a new SQLite3Store instance by opening a connection
// to the SQLite database file specified by `dbFilePath`.
// Returns a pointer to SQLite3Store or an error if the connection fails.
func NewSQLite3Store(dbFilePath string) (*SQLite3Store, error) {
    db, err := sql.Open("sqlite3", dbFilePath) // Opens the SQLite database file.
    if err != nil {
        return nil, err // Returns an error if the database file cannot be opened.
    }

    return &SQLite3Store{db}, nil // Returns the store instance wrapping the DB connection.
}

// Up applies all available migrations from the specified `migrationDir`
// to upgrade the database schema to the latest version.
// Uses the "github.com/golang-migrate/migrate/v4" library for managing migrations.
func (store *SQLite3Store) Up(migrationDir string) error {
    // Wraps the SQLite DB connection as a migrate-compatible instance.
    instance, err := sqlite3.WithInstance(store.DB, new(sqlite3.Config)) // TODO: refactor
    if err != nil {
        return err // Returns an error if the instance cannot be initialized.
    }

    // Opens the migration source (file directory in this case).
    fSrc, err := new(file.File).Open(migrationDir)
    if err != nil {
        return err // Returns an error if the migration directory cannot be accessed.
    }

    // Creates a new migration instance that links the migration source to the database.
    m, err := migrate.NewWithInstance("file", fSrc, "sqlite3", instance)
    if err != nil {
        return err // Returns an error if the migration instance fails to initialize.
    }

    // Applies all migrations up to the latest version.
    return m.Up()
}

// Down rolls back all applied migrations from the database schema.
// The migrations are loaded from the specified `migrationDir`.
func (store *SQLite3Store) Down(migrationDir string) error {
    // Wraps the SQLite DB connection as a migrate-compatible instance.
    instance, err := sqlite3.WithInstance(store.DB, new(sqlite3.Config))
    if err != nil {
        return err // Returns an error if the instance cannot be initialized.
    }

    // Opens the migration source (file directory in this case).
    fSrc, err := new(file.File).Open(migrationDir)
    if err != nil {
        return err // Returns an error if the migration directory cannot be accessed.
    }

    // Creates a new migration instance that links the migration source to the database.
    m, err := migrate.NewWithInstance("file", fSrc, "sqlite3", instance)
    if err != nil {
        return err // Returns an error if the migration instance fails to initialize.
    }

    // Rolls back all migrations to the base version (essentially a reset).
    return m.Down()
}
