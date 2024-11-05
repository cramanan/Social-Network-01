package api

import (
	"context"
	"fmt"
	"net/http"

	"Social-Network-01/api/database"
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

	err = server.Storage.SendFriendRequest(ctx, request.PathValue("userid"), sess.User.Id)
	if err != nil {
		return err
	}

	return writeJSON(writer, http.StatusCreated, "Created")
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
func (server *API) GetFollowersOfUser(writer http.ResponseWriter, request *http.Request) error {
	ctx, cancel := context.WithTimeout(request.Context(), database.TransactionTimeout)
	defer cancel()

	if request.Method != http.MethodGet {
		return writeJSON(writer, http.StatusMethodNotAllowed, APIerror{
			http.StatusMethodNotAllowed,
			"Method Not Allowed",
			"Only GET is allowed",
		})
	}

	limit, offset := parseRequestLimitAndOffset(request)

	user, err := server.Storage.GetUser(ctx, request.PathValue("userid"))
	if err != nil {
		return err
	}

	if user.Private {
		return writeJSON(writer, http.StatusUnauthorized, APIerror{
			http.StatusUnauthorized,
			"Unauthorized",
			"This account is private",
		})
	}

	users, err := server.Storage.GetFollowersOfUser(ctx, request.PathValue("userid"), limit, offset)
	if err != nil {
		return err
	}

	return writeJSON(writer, http.StatusOK, users)
}
