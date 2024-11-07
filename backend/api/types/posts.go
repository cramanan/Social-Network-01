package types

import "time"

type Post struct {
	Id        string    `json:"id"`
	Username  string    `json:"username"`
	UserId    string    `json:"userId"`
	GroupId   string    `json:"groupId"`
	Content   string    `json:"content"`
	Images    []string  `json:"images"`
	Timestamp time.Time `json:"timestamp"`
}
