package models

import "time"

type Post struct {
	Id         string
	UserId     string
	GroupId    string
	Categories []byte // JSON
	Content    string
	ImagePath  []byte // JSON
	Timestamp  time.Time
}
