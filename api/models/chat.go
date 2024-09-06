package models

import "time"

type Chat struct {
	ID          string
	SenderId    string
	RecipientId string
	Content     string
	ImgPath     string
	Timestamp   time.Time
}
