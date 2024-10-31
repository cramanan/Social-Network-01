package models

import (
	"encoding/json"
	"time"
)

type RawMessage struct {
	Type string          `json:"type"`
	Data json.RawMessage `json:"data"`
}

type Ping struct{}
type RawChat struct {
	RecipientId string `json:"recipientId"`
	Content     string `json:"content"`
}

type Chat struct {
	SenderId    string    `json:"id"`
	RecipientId string    `json:"recipientId"`
	Content     string    `json:"content"`
	Timestamp   time.Time `json:"timestamp"`
}
