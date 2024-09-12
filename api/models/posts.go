package models

import "time"

type Post struct {
	Id        string
	UserId    string
	GroupId   string
	Content   string
	ImagePath []string
	Timestamp time.Time
}
