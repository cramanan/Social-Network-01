package models

import "time"

type PostStatus int

const (
	ENUM_PUBLIC PostStatus = iota
	ENUM_PRIVATE
	ENUM_ALMOST_PRIVATE
)

type Post struct {
	Id        string
	UserId    string
	GroupName string
	Content   string
	Images    []string
	Timestamp time.Time
}

type PostRequest struct {
	UserId        string
	GroupName     string     `json:"groupName"`
	Status        PostStatus `json:"status"`
	SelectedUsers []string   `json:"selectedUsers"`
	Content       string     `json:"content"`
	Images        []string
}
