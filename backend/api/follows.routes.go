package api

import (
	"context"
	"fmt"
	"net/http"

	"Social-Network-01/api/types"
)

// GetFriendRequests retrieves all the friend requests for the currently logged-in user.
// It checks the session, fetches the list of users who sent friend requests, and returns them.
func (server *API) GetFriendRequests(writer http.ResponseWriter, request *http.Request) (err error) {
	// Retrieve the session for the current user.
	sess, err := server.Sessions.GetSession(request)
	if err != nil {
		// Return error if session retrieval fails.
		return err 
	}

	// Retrieve the list of friend requests from the storage for the current user.
	users, err := server.Storage.GetFriendRequests(request.Context(), sess.User.Id)
	if err != nil {
		// Return error if fetching friend requests fails.
		return err 
	}

	// Return the friend requests as a JSON response with HTTP Status OK.
	return writeJSON(writer, http.StatusOK, users)
}

// SendFriendRequest handles sending or canceling a friend request between users.
// It checks if the user is trying to follow themselves, sends a friend request or unfollows based on the current state.
func (server *API) SendFriendRequest(writer http.ResponseWriter, request *http.Request) error {
	// Only allow POST method for this route. If not POST, return Method Not Allowed.
	if request.Method != http.MethodPost {
		return writeJSON(writer, http.StatusMethodNotAllowed, HTTPerror(http.StatusMethodNotAllowed))
	}

	// Retrieve the session of the current user.
	sess, err := server.Sessions.GetSession(request)
	if err != nil {
		// Return error if session retrieval fails.
		return err 
	}

	// Prevent the user from following themselves.
	if sess.User.Id == request.PathValue("userid") {
		return fmt.Errorf("cannot follow yourself")
	}

	// Retrieve the user ID of the person being followed or unfollowed.
	userId := request.PathValue("userid")

	// Check if the current user already follows the target user.
	follows, err := server.Storage.Follows(request.Context(), userId, sess.User.Id)
	if err != nil {
		// Return error if checking follow status fails.
		return err 
	}

	// Define the method to call based on whether the user is following or not.
	var methodToUse func(context.Context, string, string) error

	if !follows {
		// If not following, send a friend request.
		methodToUse = server.Storage.SendFriendRequest
		// Notify the user through WebSocket about the incoming friend request.
		if conn, ok := server.WebSocket.Users.Lookup(userId); ok {
			conn.WriteJSON(types.SocketMessage[string]{
				Type: "friend-request",
				Data: fmt.Sprintf("%s has sent you a friend request.", sess.User.Nickname),
			})
		}
	} else {
		// If already following, unfollow the user.
		methodToUse = server.Storage.UnfollowUser
	}

	// Execute the appropriate method (either send a friend request or unfollow).
	err = methodToUse(request.Context(), userId, sess.User.Id)
	if err != nil {
		// Return error if the operation fails.
		return err 
	}

	// Return a success response with "OK".
	return writeJSON(writer, http.StatusOK, "OK")
}

// AcceptFriendRequest allows the currently logged-in user to accept a friend request from another user.
func (server *API) AcceptFriendRequest(writer http.ResponseWriter, request *http.Request) (err error) {
	// Retrieve the session for the current user.
	sess, err := server.Sessions.GetSession(request)
	if err != nil {
		// Return error if session retrieval fails.
		return err 
	}

	// Retrieve the ID of the user whose friend request is being accepted.
	followerId := request.PathValue("userid")

	// Call the storage method to accept the friend request.
	err = server.Storage.AcceptFriendRequest(request.Context(), sess.User.Id, followerId)
	if err != nil {
		// Return error if accepting the friend request fails.
		return err 
	}

	// Return a success response with HTTP Status OK.
	return writeJSON(writer, http.StatusOK, http.StatusOK)
}

// DeclineFriendRequest allows the user to decline a friend request.
func (server *API) DeclineFriendRequest(writer http.ResponseWriter, request *http.Request) (err error) {
	// Retrieve the session for the current user.
	sess, err := server.Sessions.GetSession(request)
	if err != nil {
		// Return error if session retrieval fails.
		return err 
	}

	// Retrieve the ID of the user whose friend request is being declined.
	followerId := request.PathValue("userid")

	// Call the storage method to unfollow the user (decline the request).
	err = server.Storage.UnfollowUser(request.Context(), sess.User.Id, followerId)
	if err != nil {
		// Return error if declining the friend request fails.
		return err 
	}

	// Return a success response with HTTP Status OK.
	return writeJSON(writer, http.StatusOK, http.StatusOK)
}

// GetProfileFollowers retrieves the list of followers for the currently logged-in user, with pagination support.
func (server *API) GetProfileFollowers(writer http.ResponseWriter, request *http.Request) error {
	// Only allow GET method for this route. If not GET, return Method Not Allowed.
	if request.Method != http.MethodGet {
		return writeJSON(writer, http.StatusMethodNotAllowed, HTTPerror(http.StatusMethodNotAllowed))
	}

	// Retrieve the session for the current user.
	sess, err := server.Sessions.GetSession(request)
	if err != nil {
		// Return error if session retrieval fails.
		return err 
	}

	// Parse the limit and offset parameters from the request for pagination.
	limit, offset := parseRequestLimitAndOffset(request)

	// Retrieve the list of followers from the storage.
	users, err := server.Storage.GetProfileFollowers(request.Context(), sess.User.Id, limit, offset)
	if err != nil {
		// Return error if fetching followers fails.
		return err 
	}

	// Return the list of followers as a JSON response with HTTP Status OK.
	return writeJSON(writer, http.StatusOK, users)
}

// GetProfileFollowing retrieves the list of users the currently logged-in user is following, with pagination support.
func (server *API) GetProfileFollowing(writer http.ResponseWriter, request *http.Request) error {
	// Only allow GET method for this route. If not GET, return Method Not Allowed.
	if request.Method != http.MethodGet {
		return writeJSON(writer, http.StatusMethodNotAllowed, HTTPerror(http.StatusMethodNotAllowed))
	}

	// Retrieve the session for the current user.
	sess, err := server.Sessions.GetSession(request)
	if err != nil {
		// Return error if session retrieval fails.
		return err 
	}

	// Parse the limit and offset parameters from the request for pagination.
	limit, offset := parseRequestLimitAndOffset(request)

	// Retrieve the list of users the current user is following from the storage.
	users, err := server.Storage.GetProfileFollowing(request.Context(), sess.User.Id, limit, offset)
	if err != nil {
		// Return error if fetching following list fails.
		return err 
	}

	// Return the list of users the current user is following as a JSON response with HTTP Status OK.
	return writeJSON(writer, http.StatusOK, users)
}
