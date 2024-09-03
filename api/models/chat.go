package models

import "time"

type Chat struct {
	ID          string
	SenderId    string
	RecipientId string
	Content     string
	Timestamp   time.Time
}
