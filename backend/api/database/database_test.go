package database_test

import (
	"context"
	"log"
	"os"
	"testing"

	. "Social-Network-01/api/database"
	"Social-Network-01/api/types"
)

var store *SQLite3Store

func TestMain(m *testing.M) {

	file, err := os.CreateTemp("../db", "sqlite3_test_*.db")
	if err != nil {
		log.Fatalln(err)

	}
	defer os.Remove(file.Name())
	store, err = NewSQLite3Store(file.Name())
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	defer store.Close()
	defer os.Exit(m.Run())
}

func TestCreateUser(t *testing.T) {
	err := store.Ping()
	if err != nil {
		t.Fatal(err)
	}

	testCases := []struct {
		desc        string
		registerReq *types.RegisterRequest
		ShouldErr   bool
	}{
		{
			desc: "valid register request",
			registerReq: &types.RegisterRequest{
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
			registerReq: &types.RegisterRequest{
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
			registerReq: &types.RegisterRequest{
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
			registerReq: &types.RegisterRequest{
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
		t.Run(tC.desc, func(t *testing.T) {
			_, err := store.RegisterUser(context.Background(), tC.registerReq)
			if (err != nil) != tC.ShouldErr {
				t.Fatal(err)
			}
		})
	}
}
