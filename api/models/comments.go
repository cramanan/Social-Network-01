package models

import "time"

type Comments struct {
	Id        string
	UserId    string
	ParentId  string
	Content   string
	ImgPath   []string
	TimeStamp time.Time
}
