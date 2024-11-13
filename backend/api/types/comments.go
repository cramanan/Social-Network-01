package types

import "time"

type Comment struct {
	UserId    string    `json:"userId"`
	PostId    string    `json:"postId"`
	Content   string    `json:"content"`
	Image     string    `json:"images"`
	Timestamp time.Time `json:"timestamp"`
}
