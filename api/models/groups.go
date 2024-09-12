package models

import "time"

type Group struct {
	Id          string
	Name        string
	Description string
	UsersIds    []string
	Timestamp   time.Time
}
