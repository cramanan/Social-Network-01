package models

import (
	"encoding/json"
	"time"
)

type SocketMessage[T any] struct {
	Type string `json:"type"`
	Data T      `json:"data"`
}

type RawMessage struct {
	Type string          `json:"type"`
	Data json.RawMessage `json:"data"`
}

type ClientChat struct {
	RecipientId string `json:"recipientId"`
	Content     string `json:"content"`
}

type Chat struct {
	SenderId    string    `json:"senderId"`
	RecipientId string    `json:"recipientId"`
	Content     string    `json:"content"`
	Timestamp   time.Time `json:"timestamp"`
}
