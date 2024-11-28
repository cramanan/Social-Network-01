package types

import (
	"time"
)

const DateFormat = "2006-01-02T15:04"

// Event represents an event with details such as an ID, group ID, title, description, date, and attendance status.
type Event struct {
	// Id uniquely identifies the event.
	Id          string `json:"id"`          

	// GroupId identifies the group associated with the event.
	GroupId     string `json:"groupId"`     

	// Title is the name of the event.
	Title       string `json:"title"`      

	// Description provides a detailed explanation of the event.
	Description string `json:"description"` 

	// Date represents the event's date in string format.
	Date        string `json:"date"`       
	 
	// Going indicates whether the user is attending the event.
	Going       bool   `json:"going"`       
}

// Valid checks whether the event has valid fields. It returns true if all the fields are valid according to the following rules:
func (e Event) Valid() (isValid bool) {
	// Check if GroupId is not empty.
	isGroupIdValid := e.GroupId != ""

	// Check if Title is not empty.
	isTitleValid := e.Title != ""

	// Check if Description is not empty.
	isDescriptionValid := e.Description != ""

	// Check if Date is in a valid format using time.Parse.
	_, err := time.Parse(DateFormat, e.Date)
	isTimeValid := err == nil

	// Return true if all checks pass.
	return isGroupIdValid && isTitleValid && isDescriptionValid && isTimeValid
}

