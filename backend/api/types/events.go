package types

import (
	"time"
)

const DateFormat = "2006-01-02T15:04"

type Event struct {
	Id          string `json:"id"`
	GroupId     string `json:"groupId"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Date        string `json:"date"`
	Going       bool   `json:"going"`
}

func (e Event) Valid() (isValid bool) {
	isGroupIdValid := e.GroupId != ""
	isTitleValid := e.Title != ""
	isDescriptionValid := e.Description != ""
	_, err := time.Parse(DateFormat, e.Date)
	isTimeValid := err == nil
	return isGroupIdValid && isTitleValid && isDescriptionValid && isTimeValid
}
