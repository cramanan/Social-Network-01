package database_test

import (
	"context"
	"io"
	"log"
	"os"
	"testing"

	. "Social-Network-01/api/database"
	"Social-Network-01/api/models"
)

var store *SQLite3Store

func TestMain(m *testing.M) {
	file, err := os.CreateTemp(".", "sqlite3_test_*.db")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	defer os.Remove(file.Name())

	initFile, err := os.Open("../db/init.sql")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	body, err := io.ReadAll(initFile)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	store, err = NewSQLite3Store(file.Name())
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	_, err = store.Exec(string(body))
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	os.Exit(m.Run())
}

func TestCreateUser(t *testing.T) {
	err := store.Ping()
	if err != nil {
		t.Fatal(err)
	}

	_, err = store.RegisterUser(context.Background(), &models.RegisterRequest{
		Nickname:    "John",
		Email:       "john.doe@mail.com",
		Password:    "secret-password",
		FirstName:   "john",
		LastName:    "Doe",
		DateOfBirth: "2022-22-02",
	})
	if err != nil {
		t.Fatal(err)
	}

}
