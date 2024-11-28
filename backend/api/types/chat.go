package types

import (
	"encoding/json"
	"time"
)

// SocketMessage is a generic type that represents a message sent over a socket connection, with a specific type and data payload.
type SocketMessage[T any] struct {
	// Type indicates the type of the message.
	Type string `json:"type"` 

	// Data contains the message's actual content, which can vary based on the type.
	Data T      `json:"data"` 
}

// RawMessage represents a raw socket message, where the data is stored as a raw JSON message (without specific struct decoding).
type RawMessage struct {
	// Type indicates the type of the raw message.
	Type string          `json:"type"` 

	// Data holds the raw JSON content that can be parsed later.
	Data json.RawMessage `json:"data"` 
}

// ClientChat represents the structure of a chat message sent from the client, including the recipient and message content.
type ClientChat struct {
	// RecipientId identifies the recipient of the chat message.
	RecipientId string `json:"recipientId"` 

	// Content holds the actual text of the chat message.
	Content     string `json:"content"`     
}

// ServerChat represents a chat message sent from the server, containing sender, recipient, message content, and timestamp.
type ServerChat struct {
	// SenderId identifies the sender of the chat message.
	SenderId    string    `json:"senderId"`    

	// RecipientId identifies the recipient of the chat message.
	RecipientId string    `json:"recipientId"` 

	// Content holds the actual text of the chat message.
	Content     string    `json:"content"`     
	
	// Timestamp marks the time when the message was sent.
	Timestamp   time.Time `json:"timestamp"`   
}

