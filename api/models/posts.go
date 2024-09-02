package models

import "time"

type Post struct {
	ID         string    `json:"id"`
	UserID     string    `json:"userID"`
	Username   string    `json:"username"`
	Categories []byte    `json:"categories"`
	Content    string    `json:"content"`
	Created    time.Time `json:"created"`
	Image      *string   `json:"image"`
}

type PostRequest struct {
	Content    string `json:"content"`
	UserID     string
	Username   string
	Categories []string `json:"categories"`
	ImagePath  *string
}
