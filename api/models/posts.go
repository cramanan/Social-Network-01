package models

import (
	"time"

	"github.com/gofrs/uuid"
)

type Post struct {
	Id         uuid.UUID
	UserId     uuid.UUID
	GroupId    uuid.UUID
	Categories []byte // JSON
	Content    string
	ImagePath  []byte // JSON
	Timestamp  time.Time
}
