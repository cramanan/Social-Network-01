package types

import "time"

// User represents a user's profile information, including personal details and privacy settings.
type User struct {
	// Id uniquely identifies the user.
	Id          string    `json:"id"`

	// Nickname is the user's chosen display name.
	Nickname    string    `json:"nickname"`

	// Email is the user's email address.
	Email       string    `json:"email"`

	// FirstName is the user's given name.
	FirstName   string    `json:"firstName"`

	// LastName is the user's family name.
	LastName    string    `json:"lastName"`

	// DateOfBirth is the user's birth date.
	DateOfBirth time.Time `json:"dateOfBirth"`

	// ImagePath is the URL or path to the user's profile image.
	ImagePath   string    `json:"image"`

	// AboutMe is an optional description of the user. It can be null if not provided.
	AboutMe     *string   `json:"aboutMe"`

	// IsPrivate indicates whether the user's profile is private.
	IsPrivate   bool      `json:"isPrivate"`

	// Timestamp is the time when the user profile was created or last updated.
	Timestamp   time.Time `json:"timestamp"`
}

// UserStats holds statistics related to a user's activity, such as followers, following, posts, and likes.
type UserStats struct {
	// Id uniquely identifies the user.
	Id           string `json:"id"`

	// NumFollowers represents how many people follow this user.
	NumFollowers int    `json:"numFollowers"`

	// NumFollowing represents how many people this user is following.
	NumFollowing int    `json:"numFollowing"`

	// NumPosts represents how many posts this user has made.
	NumPosts     int    `json:"numPosts"`

	// NumLikes represents how many likes this user has received.
	NumLikes     int    `json:"numLikes"`
}

// RegisterRequest represents the data required for a new user to register in the system.
type RegisterRequest struct {
	// Nickname is the user's chosen display name during registration.
	Nickname    string `json:"nickname"`

	// Email is the user's email address for registration.
	Email       string `json:"email"`

	// Password is the user's password chosen during registration.
	Password    string `json:"password"`

	// FirstName is the user's given name during registration.
	FirstName   string `json:"firstName"`

	// LastName is the user's family name during registration.
	LastName    string `json:"lastName"`

	// DateOfBirth is the user's birth date in string format during registration.
	DateOfBirth string `json:"dateofbirth"`
}

// LoginRequest represents the credentials required for a user to log in.
type LoginRequest struct {
	// Email is the user's email address for login.
	Email    string `json:"email"`

	// Password is the user's password for login.
	Password string `json:"password"`
}

// OnlineUser represents a user with an additional status to indicate if they are currently online.
type OnlineUser struct {
	// Embedding the User struct to inherit all fields from the User struct.
	*User    

	// Online indicates whether the user is currently online.
	Online bool `json:"online"`
}
