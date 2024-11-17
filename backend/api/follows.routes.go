package api

import (
	"context"
	"fmt"
	"net/http"

	"Social-Network-01/api/database"
	"Social-Network-01/api/types"
)

// Perform the action of following one from another.
//
// `server` is a pointer of the API type (see ./api/api.go). It contains a session reference.
func (server *API) SendFriendRequest(writer http.ResponseWriter, request *http.Request) error {
	ctx, cancel := context.WithTimeout(request.Context(), database.TransactionTimeout)
	defer cancel()
	if request.Method != http.MethodPost {
		return writeJSON(writer, http.StatusMethodNotAllowed, APIerror{
			http.StatusMethodNotAllowed,
			"Method Not Allowed",
			"Only POST is allowed",
		})
	}

	sess, err := server.Sessions.GetSession(request)
	if err != nil {
		return err
	}

	if sess.User.Id == request.PathValue("userid") {
		return fmt.Errorf("cannot follow yourself")
	}

	userId := request.PathValue("userid")

	follows, err := server.Storage.Follows(ctx, userId, sess.User.Id)
	if err != nil {
		return err
	}
	var methodToUse func(context.Context, string, string) error

	if !follows {
		methodToUse = server.Storage.SendFriendRequest
		if conn, ok := server.WebSocket.Users[userId]; ok {
			conn.WriteJSON(types.SocketMessage[string]{
				Type: "friend-request",
				Data: fmt.Sprintf("%s has sent you a friend request.", sess.User.Nickname),
			})
		}

	} else {
		methodToUse = server.Storage.UnfollowUser
	}

	err = methodToUse(ctx, userId, sess.User.Id)
	if err != nil {
		return err
	}

	return writeJSON(writer, http.StatusOK, "OK")
}

func (server *API) AcceptFriendRequest(writer http.ResponseWriter, request *http.Request) (err error) {
	sess, err := server.Sessions.GetSession(request)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(request.Context(), database.TransactionTimeout)
	defer cancel()

	followerId := request.PathValue("userid")

	err = server.Storage.AcceptFriendRequest(ctx, sess.User.Id, followerId)
	if err != nil {
		return err
	}

	return writeJSON(writer, http.StatusOK, http.StatusOK)
}

// Retrieve all follower of a user from the database.
//
// `server` is a pointer of the API type (see ./api/api.go). It contains a session reference.
func (server *API) GetProfileFollowers(writer http.ResponseWriter, request *http.Request) error {
	ctx, cancel := context.WithTimeout(request.Context(), database.TransactionTimeout)
	defer cancel()

	if request.Method != http.MethodGet {
		return writeJSON(writer, http.StatusMethodNotAllowed, APIerror{
			http.StatusMethodNotAllowed,
			"Method Not Allowed",
			"Only GET is allowed",
		})
	}

	sess, err := server.Sessions.GetSession(request)
	if err != nil {
		return err
	}

	limit, offset := parseRequestLimitAndOffset(request)

	users, err := server.Storage.GetProfileFollowers(ctx, sess.User.Id, limit, offset)
	if err != nil {
		return err
	}

	return writeJSON(writer, http.StatusOK, users)
}

func (server *API) GetProfileFollowing(writer http.ResponseWriter, request *http.Request) error {
	ctx, cancel := context.WithTimeout(request.Context(), database.TransactionTimeout)
	defer cancel()

	if request.Method != http.MethodGet {
		return writeJSON(writer, http.StatusMethodNotAllowed, APIerror{
			http.StatusMethodNotAllowed,
			"Method Not Allowed",
			"Only GET is allowed",
		})
	}

	sess, err := server.Sessions.GetSession(request)
	if err != nil {
		return err
	}

	limit, offset := parseRequestLimitAndOffset(request)

	users, err := server.Storage.GetProfileFollowing(ctx, sess.User.Id, limit, offset)
	if err != nil {
		return err
	}

	return writeJSON(writer, http.StatusOK, users)
}
