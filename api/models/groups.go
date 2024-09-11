package models

import "time"

type Group struct {
	Id          string
	Name        string
	Description string
	UsersIds    []byte
	Timestamp   time.Time
}
