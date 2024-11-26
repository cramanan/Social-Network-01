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

type Comment struct {
	Username  string    `json:"username"`
	UserImage string    `json:"userImage"`
	UserId    string    `json:"userId"`
	PostId    string    `json:"postId"`
	Content   string    `json:"content"`
	Image     string    `json:"image"`
	Timestamp time.Time `json:"timestamp"`
}
