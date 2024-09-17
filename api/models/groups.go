package models

import "time"

type Group struct {
	Name        string
	Description string
	UsersIds    []string
	Timestamp   time.Time
}
