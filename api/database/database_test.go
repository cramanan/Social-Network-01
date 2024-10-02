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
		log.Fatalln(err)
	}
	defer os.Remove(file.Name())

	initFile, err := os.Open("../db/init.sql")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	body, err := io.ReadAll(initFile)
	if err != nil {
		log.Fatalln(err)
	}

	store, err = NewSQLite3Store(file.Name())
	if err != nil {
		log.Fatalln(err)

	}
	defer store.Close()

	_, err = store.Exec(string(body))
	if err != nil {
		log.Fatalln(err)
	}

	os.Remove(file.Name())
	os.Exit(m.Run())
}

func TestCreateUser(t *testing.T) {
	err := store.Ping()
	if err != nil {
		t.Fatal(err)
	}

	testCases := []struct {
		desc        string
		registerReq *models.RegisterRequest
		ShouldErr   bool
	}{
		{
			desc: "valid register request",
			registerReq: &models.RegisterRequest{
				Nickname:    "nickname",
				Email:       "email@example.com",
				Password:    "password",
				FirstName:   "first",
				LastName:    "last",
				DateOfBirth: "2000-01-01",
			},
			ShouldErr: false,
		},
		{
			desc: "missing nickname",
			registerReq: &models.RegisterRequest{
				Email:       "email@example.com",
				Password:    "password",
				FirstName:   "first",
				LastName:    "last",
				DateOfBirth: "2000-01-01",
			},
			ShouldErr: true,
		},
		{
			desc: "weak password",
			registerReq: &models.RegisterRequest{
				Nickname:    "nickname",
				Email:       "email@example.com",
				Password:    "weak",
				FirstName:   "first",
				LastName:    "last",
				DateOfBirth: "2000-01-01",
			},
			ShouldErr: true,
		},
		{
			desc: "invalid date of birth",
			registerReq: &models.RegisterRequest{
				Nickname:    "nickname",
				Email:       "email@example.com",
				Password:    "password",
				FirstName:   "first",
				LastName:    "last",
				DateOfBirth: "invalid date",
			},
			ShouldErr: true,
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(tc *testing.T) {
			_, err := store.RegisterUser(context.TODO(), tC.registerReq)
			if (err != nil) != tC.ShouldErr {
				tc.Fatal(err)
			}
		})
	}
}
