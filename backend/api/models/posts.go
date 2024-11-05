package models

import "time"

type PostStatus int

const (
	ENUM_PUBLIC PostStatus = iota
	ENUM_PRIVATE
	ENUM_ALMOST_PRIVATE
)

type Post struct {
	Id        string    `json:"id"`
	Username  string    `json:"username"`
	UserId    string    `json:"userId"`
	GroupId   string    `json:"groupId"`
	Content   string    `json:"content"`
	Images    []string  `json:"images"`
	Timestamp time.Time `json:"timestamp"`
}

type PostRequest struct {
	UserId        string
	GroupName     *string    `json:"groupName"`
	Status        PostStatus `json:"status"`
	SelectedUsers []string   `json:"selectedUsers"`
	Content       string     `json:"content"`
	Images        []string
}