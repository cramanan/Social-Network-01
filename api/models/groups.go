package models

import (
	"time"

	"github.com/gofrs/uuid"
)

type Group struct {
	Id          uuid.UUID
	Name        string
	UsersIds    []byte
	Content     string
	ImagesPaths []byte
	Timestamp   time.Time
}
