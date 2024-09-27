package models

import "time"

type Status int

const (
	ENUM_PUBLIC Status = iota
	ENUM_PRIVATE
	ENUM_ALMOST_PRIVATE
)

type Post struct {
	Id        string
	UserId    string
	Status    Status
	GroupId   string
	Content   string
	ImagePath []string
	Timestamp time.Time
}

type PostRequest struct {
	UserId     string
	GroupId    string
	Categories []string
	Content    string
	Images     []string
}
