package database_test

import (
	. "Social-Network-01/api/database"

	"os"
	"testing"
)

func TestSQLite3Store(t *testing.T) {
	// Create a temporary database file
	dbFile, err := os.CreateTemp("", "sqlite3_test.db")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(dbFile.Name())

	// Initialize the SQLite3 store
	storage, err := NewSQLite3Store(dbFile.Name())
	if err != nil {
		t.Fatal(err)
	}

	_, err = storage.Exec("CREATE TABLE test (id INTEGER PRIMARY KEY, name TEXT)")
	if err != nil {
		t.Fatal(err)
	}
}
