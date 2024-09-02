package models

import (
	"time"
)

type Comment struct {
	ID       string    `json:"id"`
	UserID   string    `json:"userid"`
	PostID   string    `json:"postid"`
	Username string    `json:"username"`
	Content  string    `json:"content"`
	Created  time.Time `json:"created"`
}

type CommentRequest struct {
	UserID   string
	PostID   string
	Username string
	Content  string `json:"content"`
}
