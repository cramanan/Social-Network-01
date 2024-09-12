package database_test

import (
	. "Social-Network-01/api/database"

	"testing"

	"os"
)

func TestSQLite3Store(t *testing.T) {
	// Create a temporary database file
	dbFile, err := os.CreateTemp("", "sqlite3_test_")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(dbFile.Name())

	// Initialize the SQLite3 store
	storage, err := NewSQLite3Store(dbFile.Name())
	if err != nil {
		t.Fatal(err)
	}

	// Test database operations
	t.Run("CreateTable", func(t *testing.T) {
		_, err := storage.Exec("CREATE TABLE test (id INTEGER PRIMARY KEY, name TEXT)")
		if err != nil {
			t.Fatal(err)
		}
	})
}
