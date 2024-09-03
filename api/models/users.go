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
