package types

import "time"

// Post represents a post made by a user in a group or on a platform, including the content, associated images, and timestamp.
type Post struct {
	// Id uniquely identifies the post.
	Id string `json:"id"`

	// Username is the name of the user who made the post.
	Username string `json:"username"`

	// UserId is the identifier of the user who made the post.
	UserId string `json:"userId"`

	UserImage string `json:"userImage"`

	// GroupId is the optional ID of the group where the post was made, if applicable.
	GroupId *string `json:"groupId"`

	// Content holds the text content of the post.
	Content string `json:"content"`

	// Images is a list of URLs or paths to images attached to the post.
	Images []string `json:"images"`

	// Timestamp is the time when the post was created.
	Timestamp time.Time `json:"timestamp"`
}

// Comment represents a comment made on a post, including the user's details, content, and timestamp of the comment.
type Comment struct {
	// Username is the name of the user who made the comment.
	Username string `json:"username"`

	// UserImage is the URL or path to the user's profile image.
	UserImage string `json:"userImage"`

	// UserId is the identifier of the user who made the comment.
	UserId string `json:"userId"`

	// PostId identifies the post that this comment is associated with.
	PostId string `json:"postId"`

	// Content holds the text content of the comment.
	Content string `json:"content"`

	// Image is an optional URL or path to an image attached to the comment.
	Image string `json:"image"`

	// Timestamp is the time when the comment was made.
	Timestamp time.Time `json:"timestamp"`
}
