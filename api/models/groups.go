package models

import "time"

type Group struct {
	Id          string
	Name        string
	UsersIds    []byte
	Content     string
	ImagesPaths []byte
	Timestamp   time.Time
}
