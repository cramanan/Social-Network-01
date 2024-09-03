package models

import (
	"time"

	"github.com/gofrs/uuid"
)

type Chat struct {
	ID          uuid.UUID
	SenderId    uuid.UUID
	RecipientId uuid.UUID
	Content     string
	Timestamp   time.Time
}
