package models

import "github.com/gofrs/uuid"

type Comments struct {
	Id       uuid.UUID
	ParentId uuid.UUID
	UserId   uuid.UUID
}
