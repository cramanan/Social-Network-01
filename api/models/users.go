package models

import "time"

type User struct {
	Id          string
	Nickname    string    `json:"nickname"`
	Email       string    `json:"email"`
	FirstName   string    `json:"firstName"`
	LastName    string    `json:"lastName"`
	DateOfBirth time.Time `json:"dateOfBirth"`
	ImagePath   *string   `json:"imagePath"`
	AboutMe     *string   `json:"aboutMe"`
	Private     bool      `json:"private"`
	Timestamp   time.Time `json:"timestamp"`
}

type RegisterRequest struct {
	Nickname    string `json:"nickname"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	DateOfBirth string `json:"dateofbirth"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
