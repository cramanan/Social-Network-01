package types

import "time"

type Group struct {
	Id          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Image       string    `json:"image"`
	Timestamp   time.Time `json:"timestamp"`
}
