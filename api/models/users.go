package models

import (
	"time"

	"github.com/gofrs/uuid"
)

type User struct {
	Id          uuid.UUID
	Nickname    string
	Email       string
	password    []byte
	FirstName   string
	LastName    string
	DateOfBirth time.Time
	ImagePath   *string
	AboutMe     *string
	Private     bool
	Timestamp   time.Time
}

type RegisterRequest struct {
	Nickname    string `json:"nickname"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	FirstName   string `json:"firstname"`
	LastName    string `json:"lastname"`
	DateOfBirth string `json:"dateofbirth"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
