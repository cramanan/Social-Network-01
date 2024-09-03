package api

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/mail"

	"Social-Network-01/api/database"
	"Social-Network-01/api/models"
)

func (server *API) Register(writer http.ResponseWriter, request *http.Request) error {
	if request.Method != http.MethodPost {
		return writeJSON(writer, http.StatusMethodNotAllowed,
			APIerror{
				Status:  http.StatusMethodNotAllowed,
				Error:   "Method Not Allowed",
				Message: "Method not allowed: Only POST is supported",
			})
	}

	registerReq := new(models.RegisterRequest)
	err := json.NewDecoder(request.Body).Decode(registerReq)
	if err != nil {
		return writeJSON(writer, http.StatusUnprocessableEntity,
			APIerror{
				Status:  http.StatusUnprocessableEntity,
				Error:   "Unprocessable Entity",
				Message: "Could not process register request",
			})
	}

	if registerReq.Nickname == "" ||
		registerReq.Email == "" ||
		registerReq.Password == "" ||
		registerReq.FirstName == "" ||
		registerReq.LastName == "" ||
		registerReq.DateOfBirth == "" {

		return writeJSON(writer, http.StatusUnauthorized,
			APIerror{
				Status:  http.StatusUnauthorized,
				Error:   "Unauthorized",
				Message: "All fields are required",
			})
	}

	if _, err = mail.ParseAddress(registerReq.Email); err != nil {
		return writeJSON(writer, http.StatusBadRequest,
			APIerror{
				Status:  http.StatusBadRequest,
				Error:   "Bad Request",
				Message: "Invalid Email address",
			})
	}

	ctx, cancel := context.WithTimeout(request.Context(), database.TransactionTimeout)
	defer cancel()
	user, err := server.Storage.RegisterUser(ctx, registerReq)
	if errors.Is(err, database.ErrConflict) {
		return writeJSON(writer, http.StatusConflict,
			APIerror{
				Status:  http.StatusConflict,
				Error:   "Conflict",
				Message: "Email address is already taken",
			})
	}
	if err != nil {
		return err
	}

	// session := server.Sessions.NewSession(writer, request)
	// session.User = user
	return writeJSON(writer, http.StatusCreated, user)
}

func (server *API) Login(writer http.ResponseWriter, request *http.Request) error {
	if request.Method != http.MethodPost {
		return writeJSON(writer, http.StatusMethodNotAllowed,
			APIerror{
				Status:  http.StatusMethodNotAllowed,
				Error:   "Method Not Allowed",
				Message: "Method not allowed: Only POST is supported",
			})
	}

	loginReq := new(models.LoginRequest)
	err := json.NewDecoder(request.Body).Decode(loginReq)
	if err != nil {
		return writeJSON(writer, http.StatusUnprocessableEntity,
			APIerror{
				Status:  http.StatusUnprocessableEntity,
				Error:   "Unprocessable Entity",
				Message: "Could not process login request",
			})
	}

	if loginReq.Email == "" ||
		loginReq.Password == "" {
		return writeJSON(writer, http.StatusBadRequest,
			APIerror{
				Status:  http.StatusUnauthorized,
				Error:   "Unauthorized",
				Message: "Email and password are required",
			})
	}

	if _, err = mail.ParseAddress(loginReq.Email); err != nil {
		return writeJSON(writer, http.StatusBadRequest,
			APIerror{
				Status:  http.StatusBadRequest,
				Error:   "Bad Request",
				Message: "Invalid Email address",
			})
	}

	/* => fonctions missing
	ctx, cancel := context.WithTimeout(request.Context(), database.TransactionTimeout)
	defer cancel()
	user, err := server.Storage.LogUser(ctx, loginReq)
	if err != nil {
		return writeJSON(writer, http.StatusBadRequest,
			APIerror{
				Status:  http.StatusBadRequest,
				Error:   "Bad Request",
				Message: "Invalid Email or Password",
			})
	}

	session := server.Sessions.NewSession(writer, request)
	session.User = user

	return writeJSON(writer, http.StatusOK, user)
	*/
	return nil
}

func (server *API) GetAllPostsFromOneUser(writer http.ResponseWriter, request *http.Request) error {
	return nil
}

func (server *API) GetAllPostsFromOneGroup(writer http.ResponseWriter, request *http.Request) error {
	return nil
}

func (server *API) GetAllPostsFromOneUsersFollows(writer http.ResponseWriter, request *http.Request) error {
	return nil
}

func (server *API) GetAllPostsFromOneUsersLikes(writer http.ResponseWriter, request *http.Request) error {
	return nil
}

func (server *API) GetAllCommentsFromOnePost(writer http.ResponseWriter, request *http.Request) error {
	return nil
}

func (server *API) GetUserFromUserid(writer http.ResponseWriter, request *http.Request) error {
	return nil
}

func (server *API) GetAccountFromUserid(writer http.ResponseWriter, request *http.Request) error {
	return nil
}

func (server *API) GetChatFrom2Userid(writer http.ResponseWriter, request *http.Request) error {
	return nil
}
