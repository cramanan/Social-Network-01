package models

import "time"

type Post struct {
	Id         string
	UserId     string
	GroupId    string
	Categories []string
	Content    string
	ImagePath  []string
	Timestamp  time.Time
}

type PostRequest struct {
	UserId     string
	GroupId    string
	Categories []string
	Content    string
	Images     []string
}
