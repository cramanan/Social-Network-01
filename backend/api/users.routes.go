package api

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/mail"

	"Social-Network-01/api/database"
	"Social-Network-01/api/types"
)

// Perform the action of registering one user in the database.
// `server` is a pointer of the API type (see ./api/api.go). It contains a session reference.
func (server *API) Register(writer http.ResponseWriter, request *http.Request) error {

	if request.Method != http.MethodPost {
		return writeJSON(writer, http.StatusMethodNotAllowed, HTTPerror(http.StatusMethodNotAllowed))
	}

	registerReq := new(types.RegisterRequest)
	err := json.NewDecoder(request.Body).Decode(registerReq)
	if err != nil {
		return writeJSON(writer, http.StatusUnprocessableEntity, HTTPerror(http.StatusUnprocessableEntity))
	}

	if registerReq.Nickname == "" ||
		registerReq.Email == "" ||
		registerReq.Password == "" ||
		registerReq.FirstName == "" ||
		registerReq.LastName == "" ||
		registerReq.DateOfBirth == "" {

		return writeJSON(writer, http.StatusUnauthorized, HTTPerror(http.StatusBadRequest, "All fields are required"))
	}

	if _, err = mail.ParseAddress(registerReq.Email); err != nil {
		return writeJSON(writer, http.StatusBadRequest, HTTPerror(http.StatusBadRequest, "Invalid Email address"))
	}

	user, err := server.Storage.RegisterUser(request.Context(), registerReq)
	if errors.Is(err, database.ErrConflict) {
		return writeJSON(writer, http.StatusConflict, HTTPerror(http.StatusConflict, "Email address is already taken"))
	}
	if err != nil {
		return err
	}

	session := server.Sessions.NewSession(writer, request)
	session.User = user

	return writeJSON(writer, http.StatusCreated, user)
}

// Perform the action of logging one user.
//
// `server` is a pointer of the API type (see ./api/api.go). It contains a session reference.
func (server *API) Login(writer http.ResponseWriter, request *http.Request) (err error) {
	if request.Method != http.MethodPost {
		return writeJSON(writer, http.StatusMethodNotAllowed, HTTPerror(http.StatusMethodNotAllowed))
	}

	loginReq := new(types.LoginRequest)

	if err := json.NewDecoder(request.Body).Decode(loginReq); err != nil {
		return writeJSON(writer, http.StatusUnprocessableEntity, HTTPerror(http.StatusUnprocessableEntity))
	}

	if loginReq.Email == "" || loginReq.Password == "" {
		return writeJSON(writer, http.StatusBadRequest, HTTPerror(http.StatusUnauthorized, "Email and password are required"))
	}

	if _, err = mail.ParseAddress(loginReq.Email); err != nil {
		return writeJSON(writer, http.StatusBadRequest, HTTPerror(http.StatusBadRequest, "Invalid Email address"))
	}

	user, err := server.Storage.LogUser(request.Context(), loginReq)
	if err != nil {
		return writeJSON(writer, http.StatusBadRequest, HTTPerror(http.StatusBadRequest, "Invalid Email or Password"))
	}

	session := server.Sessions.NewSession(writer, request)
	session.User = user

	return writeJSON(writer, http.StatusOK, user)
}

// This method acts as a router for different HTTP methods.
//
// `server` is a pointer of the API type (see ./api/api.go). It contains a session reference.
func (server *API) User(writer http.ResponseWriter, request *http.Request) (err error) {
	switch request.Method {
	case http.MethodGet:
		userId := request.PathValue("userid")
		user, err := server.Storage.GetUser(request.Context(), userId)
		if err == sql.ErrNoRows {
			return writeJSON(writer, http.StatusNotFound, HTTPerror(http.StatusNotFound, "User not found"))
		}
		if err != nil {
			return err
		}

		sess, err := server.Sessions.GetSession(request)
		if err != nil {
			return err
		}

		if !user.IsPrivate || sess.User.Id == userId {
			return writeJSON(writer, http.StatusOK, user)
		}

		follows, err := server.Storage.Follows(request.Context(), userId, sess.User.Id)
		if !follows || err != nil {
			return writeJSON(
				writer,
				http.StatusUnauthorized,
				HTTPerror(http.StatusUnauthorized, "You are not allowed to access this ressource"),
			)
		}

		return writeJSON(writer, http.StatusOK, user)

	default:
		return writeJSON(writer, http.StatusMethodNotAllowed, HTTPerror(http.StatusMethodNotAllowed))
	}
}

func (server *API) Profile(writer http.ResponseWriter, request *http.Request) (err error) {
	sess, err := server.Sessions.GetSession(request)
	if err != nil {
		return err
	}

	switch request.Method {
	case http.MethodGet:
		s, err := server.Sessions.GetSession(request)
		if err != nil {
			return writeJSON(writer, http.StatusUnauthorized, HTTPerror(http.StatusUnauthorized, "You are unauthorized to access this ressource."))
		}

		return writeJSON(writer, http.StatusOK, s.User)

	case http.MethodPatch:
		err = request.ParseMultipartForm(5 * (1 << 20))
		if err != nil {
			return err
		}

		user := types.User{}

		data := request.MultipartForm.Value["data"]
		if len(data) != 1 {
			return fmt.Errorf("invalid number of datas")
		}

		err = json.Unmarshal([]byte(data[0]), &user)
		if err != nil {
			return err
		}

		imgs, err := MultiPartFiles(request)
		if err != nil {
			return err
		}

		//TODO: DELETE old profile picture

		if len(imgs) > 0 {
			user.ImagePath = imgs[0]
		} else {
			user.ImagePath = ""
		}

		modified, err := server.Storage.UpdateUser(request.Context(), sess.User.Id, user)
		if err != nil {
			return err
		}
		sess.User = *modified

		return writeJSON(writer, http.StatusOK, modified)

	case http.MethodDelete:
		sess, err := server.Sessions.GetSession(request)
		if err != nil {
			return err
		}

		if sess.User.Id != request.PathValue("userid") {
			return writeJSON(writer, http.StatusUnauthorized, HTTPerror(http.StatusUnauthorized, "You are not authorized to perform this action."))
		}

		err = server.Storage.DeleteUser(request.Context(), sess.User.Id)
		if err != nil {
			return err
		}

		return writeJSON(writer, http.StatusNoContent, "")

	default:
		return writeJSON(writer, http.StatusMethodNotAllowed, HTTPerror(http.StatusMethodNotAllowed))
	}
}

func (server *API) GetUserStats(writer http.ResponseWriter, request *http.Request) error {
	stats, err := server.Storage.GetUserStats(request.Context(), request.PathValue("userid"))
	if err != nil {
		return err
	}
	return writeJSON(writer, http.StatusOK, stats)
}

func (server *API) GetOnlineUsers(writer http.ResponseWriter, request *http.Request) error {
	sess, err := server.Sessions.GetSession(request)
	if err != nil {
		return err
	}

	limit, offset := parseRequestLimitAndOffset(request)

	users, err := server.Storage.GetMessagedUsers(request.Context(), sess.User.Id, limit, offset)
	if err != nil {
		return err
	}

	onlineUsers := make([]types.OnlineUser, len(users))

	for idx, user := range users {
		onlineUsers[idx] = types.OnlineUser{User: user}
		_, onlineUsers[idx].Online = server.WebSocket.Users.Lookup(user.Id)
	}

	return writeJSON(writer, http.StatusOK, onlineUsers)
}

func (server *API) GetUserFriendList(writer http.ResponseWriter, request *http.Request) (err error) {
	sess, err := server.Sessions.GetSession(request)
	if err != nil {
		return err
	}

	limit, offset := parseRequestLimitAndOffset(request)

	users, err := server.Storage.GetUserFriendList(context.TODO(), sess.User.Id, limit, offset)
	if err != nil {
		return err
	}

	onlineUsers := make([]types.OnlineUser, len(users))

	for idx, user := range users {
		onlineUsers[idx] = types.OnlineUser{User: user}
		_, onlineUsers[idx].Online = server.WebSocket.Users.Lookup(user.Id)
	}

	return writeJSON(writer, http.StatusOK, onlineUsers)
}
