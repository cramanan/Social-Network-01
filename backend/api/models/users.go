package models

import "time"

type User struct {
	Id          string    `json:"id"`
	Nickname    string    `json:"nickname"`
	Email       string    `json:"email"`
	FirstName   string    `json:"firstName"`
	LastName    string    `json:"lastName"`
	DateOfBirth time.Time `json:"dateOfBirth"`
	ImagePath   string    `json:"image"`
	AboutMe     *string   `json:"aboutMe"`
	Private     bool      `json:"private"`
	Timestamp   time.Time `json:"timestamp"`
}

type UserStats struct {
	Id           string `json:"id"`
	NumFollowers int    `json:"numFollowers"`
	NumFollowing int    `json:"numFollowing"`
	NumPosts     int    `json:"numPosts"`
	NumLikes     int    `json:"numLikes"`
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

type OnlineUser struct {
	*User
	Online bool `json:"online"`
}
