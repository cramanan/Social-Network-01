package types

import "time"

// Group represents a group with details such as an ID, name, owner, description, image, and timestamp of creation.
type Group struct {
	// Id uniquely identifies the group.
	Id          string    `json:"id"`   
	
	// Name is the name of the group.       
	Name        string    `json:"name"`    

	// Owner is the identifier of the person who owns the group.    
	Owner       string    `json:"owner"`   

	// Description provides a brief explanation about the group.    
	Description string    `json:"description"` 

	// Image is the URL or path to the group's image/logo.
	Image       string    `json:"image"`   

	// Timestamp is the time when the group was created or last modified.    
	Timestamp   time.Time `json:"timestamp"`   
}
