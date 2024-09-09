package api

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"net/mail"

	"Social-Network-01/api/database"
	"Social-Network-01/api/models"
)

func (server *API) Register(writer http.ResponseWriter, request *http.Request) error {
	ctx, cancel := context.WithTimeout(request.Context(), database.TransactionTimeout)
	defer cancel()
	if request.Method != http.MethodPost {
		return writeJSON(writer, http.StatusMethodNotAllowed,
			APIerror{
				http.StatusMethodNotAllowed,
				"Method Not Allowed",
				"Method not allowed: Only POST is supported",
			})
	}

	registerReq := new(models.RegisterRequest)
	err := json.NewDecoder(request.Body).Decode(registerReq)
	if err != nil {
		return writeJSON(writer, http.StatusUnprocessableEntity,
			APIerror{
				http.StatusUnprocessableEntity,
				"Unprocessable Entity",
				"Could not process register request",
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
				http.StatusUnauthorized,
				"Unauthorized",
				"All fields are required",
			})
	}

	if _, err = mail.ParseAddress(registerReq.Email); err != nil {
		return writeJSON(writer, http.StatusBadRequest,
			APIerror{
				http.StatusBadRequest,
				"Bad Request",
				"Invalid Email address",
			})
	}

	ctx, cancel := context.WithTimeout(ctx, database.TransactionTimeout)
	defer cancel()
	user, err := server.Storage.RegisterUser(ctx, registerReq)
	if errors.Is(err, database.ErrConflict) {
		return writeJSON(writer, http.StatusConflict,
			APIerror{
				http.StatusConflict,
				"Conflict",
				"Email address is already taken",
			})
	}
	if err != nil {
		return err
	}

	session := server.Sessions.NewSession(writer, request)
	session.User = user

	return writeJSON(writer, http.StatusCreated, user)
}

func (server *API) Login(writer http.ResponseWriter, request *http.Request) (err error) {
	if request.Method != http.MethodPost {
		return writeJSON(writer, http.StatusMethodNotAllowed,
			APIerror{
				http.StatusMethodNotAllowed,
				"Method Not Allowed",
				"Method not allowed: Only POST is supported",
			})
	}

	loginReq := new(models.LoginRequest)

	if err := json.NewDecoder(request.Body).Decode(loginReq); err != nil {
		return writeJSON(writer, http.StatusUnprocessableEntity,
			APIerror{
				http.StatusUnprocessableEntity,
				"Unprocessable Entity",
				"Could not process login request",
			})
	}

	if loginReq.Email == "" || loginReq.Password == "" {
		return writeJSON(writer, http.StatusBadRequest,
			APIerror{
				http.StatusUnauthorized,
				"Unauthorized",
				"Email and password are required",
			})
	}

	if _, err = mail.ParseAddress(loginReq.Email); err != nil {
		return writeJSON(writer, http.StatusBadRequest,
			APIerror{
				http.StatusBadRequest,
				"Bad Request",
				"Invalid Email address",
			})
	}

	ctx, cancel := context.WithTimeout(request.Context(), database.TransactionTimeout)
	defer cancel()
	user, err := server.Storage.LogUser(ctx, loginReq)
	if err != nil {
		return writeJSON(writer, http.StatusBadRequest,
			APIerror{
				http.StatusBadRequest,
				"Bad Request",
				"Invalid Email or Password",
			})
	}

	session := server.Sessions.NewSession(writer, request)
	session.User = user

	return writeJSON(writer, http.StatusOK, user)
}

func (server *API) User(writer http.ResponseWriter, request *http.Request) error {
	ctx, cancel := context.WithTimeout(request.Context(), database.TransactionTimeout)
	defer cancel()
	switch request.Method {
	case http.MethodGet:
		user, err := server.Storage.GetUser(ctx, request.PathValue("userid"))
		if err != nil {
			if err == sql.ErrNoRows {
				return writeJSON(writer, http.StatusNotFound,
					APIerror{
						http.StatusNotFound,
						"Not found",
						"User not found",
					},
				)
			}
			return err
		}

		if user.Private {
			return writeJSON(writer, http.StatusUnauthorized, APIerror{
				http.StatusUnauthorized,
				"Unauthorized",
				"This account is private",
			})
		}

		return writeJSON(writer, http.StatusOK, user)

	default:
		return writeJSON(writer, http.StatusMethodNotAllowed,
			APIerror{
				http.StatusMethodNotAllowed,
				"Method Not Allowed",
				"Method not Allowed",
			})
	}
}

func (server *API) FollowUser(writer http.ResponseWriter, request *http.Request) error {
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

	err = server.Storage.FollowUser(ctx, request.PathValue("userid"), sess.User.Id)
	if err != nil {
		return err
	}

	return writeJSON(writer, http.StatusCreated, "Created")
}

func (server *API) GetFollowersOfUser(writer http.ResponseWriter, request *http.Request) error {
	limit, offset := parseRequestLimitAndOffset(request)

	ctx, cancel := context.WithTimeout(request.Context(), database.TransactionTimeout)
	defer cancel()

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

func (server *API) AllPostsFromOneUser(writer http.ResponseWriter, request *http.Request) error {
	ctx, cancel := context.WithTimeout(request.Context(), database.TransactionTimeout)
	defer cancel()
	if request.Method == http.MethodGet {

		limit, offset := parseRequestLimitAndOffset(request)
		posts, err := server.Storage.GetAllPostsFromOneUser(ctx, request.PathValue("userid"), limit, offset)
		if err != nil {
			if err == sql.ErrNoRows {
				return writeJSON(writer, http.StatusNotFound,
					APIerror{
						http.StatusNotFound,
						"Not found",
						"Posts not found",
					},
				)
			}
			return err
		}
		return writeJSON(writer, http.StatusOK, posts)
	}

	return writeJSON(writer, http.StatusMethodNotAllowed,
		APIerror{
			http.StatusMethodNotAllowed,
			"Method Not Allowed",
			"Method not Allowed",
		})
}

func (server *API) GetAllPostsFromOneGroup(writer http.ResponseWriter, request *http.Request) error {
	ctx, cancel := context.WithTimeout(request.Context(), database.TransactionTimeout)
	defer cancel()
	if request.Method == http.MethodGet {

		limit, offset := parseRequestLimitAndOffset(request)
		posts, err := server.Storage.GetGroupPosts(ctx, request.PathValue("groupid"), limit, offset)
		if err != nil {
			if err == sql.ErrNoRows {
				return writeJSON(writer, http.StatusNotFound,
					APIerror{
						http.StatusNotFound,
						"Not found",
						"Group not found",
					},
				)
			}
			return err
		}

		return writeJSON(writer, http.StatusOK, posts)
	}

	return writeJSON(writer, http.StatusMethodNotAllowed,
		APIerror{
			http.StatusMethodNotAllowed,
			"Method Not Allowed",
			"Method not Allowed",
		})
}

func (server *API) GetAllPostsFromOneUsersFollows(writer http.ResponseWriter, request *http.Request) error {
	ctx, cancel := context.WithTimeout(request.Context(), database.TransactionTimeout)
	defer cancel()
	if request.Method == http.MethodGet {

		limit, offset := parseRequestLimitAndOffset(request)
		posts, err := server.Storage.GetFollowsPosts(ctx, request.PathValue("userid"), limit, offset)
		if err != nil {
			if err == sql.ErrNoRows {
				return writeJSON(writer, http.StatusNotFound,
					APIerror{
						http.StatusNotFound,
						"Not found",
						"User not found",
					},
				)
			}
			return err
		}

		return writeJSON(writer, http.StatusOK, posts)
	}

	return writeJSON(writer, http.StatusMethodNotAllowed,
		APIerror{
			http.StatusMethodNotAllowed,
			"Method Not Allowed",
			"Method not Allowed",
		})
}

func (server *API) GetAllPostsFromOneUsersLikes(writer http.ResponseWriter, request *http.Request) error {
	ctx, cancel := context.WithTimeout(request.Context(), database.TransactionTimeout)
	defer cancel()
	if request.Method == http.MethodGet {

		userError := func(err error) error {
			if err == sql.ErrNoRows {
				return writeJSON(writer, http.StatusNotFound,
					APIerror{
						http.StatusNotFound,
						"Not found",
						"User not found",
					},
				)
			}
			return err
		}
		
		user, err := server.Storage.GetUser(ctx, request.PathValue("userid"))
		if err != nil {
			return userError(err)
		}

		sessionUser, err := server.Sessions.GetSession(request)

		if user.Private && (err != nil || sessionUser.User.Id != user.Id) {
			return writeJSON(writer, http.StatusUnauthorized, APIerror{
				http.StatusUnauthorized,
				"Unauthorized",
				"This account is private",
			})
		}

		limit, offset := parseRequestLimitAndOffset(request)
		posts, err := server.Storage.GetPostsLike(ctx, request.PathValue("userid"), limit, offset)
		if err != nil {
			return userError(err)
		}

		return writeJSON(writer, http.StatusOK, posts)
	}

	return writeJSON(writer, http.StatusMethodNotAllowed,
		APIerror{
			http.StatusMethodNotAllowed,
			"Method Not Allowed",
			"Method not Allowed",
		})
}

func (server *API) GetAllCommentsFromOnePost(writer http.ResponseWriter, request *http.Request) error {
	ctx, cancel := context.WithTimeout(request.Context(), database.TransactionTimeout)
	defer cancel()
	if request.Method == http.MethodGet {

		limit, offset := parseRequestLimitAndOffset(request)
		comments, err := server.Storage.GetComments(ctx, request.PathValue("postid"), limit, offset)
		if err != nil {
			if err == sql.ErrNoRows {
				return writeJSON(writer, http.StatusNotFound,
					APIerror{
						http.StatusNotFound,
						"Not found",
						"Post not found",
					},
				)
			}
			return err
		}

		return writeJSON(writer, http.StatusOK, comments)
	}

	return writeJSON(writer, http.StatusMethodNotAllowed,
		APIerror{
			http.StatusMethodNotAllowed,
			"Method Not Allowed",
			"Method not Allowed",
		})
}

func (server *API) GetChatFrom2Userid(writer http.ResponseWriter, request *http.Request) error {
	ctx, cancel := context.WithTimeout(request.Context(), database.TransactionTimeout)
	defer cancel()
	if request.Method == http.MethodGet {

		limit, offset := parseRequestLimitAndOffset(request)
		sessionUser, err := server.Sessions.GetSession(request)
		if err != nil {
			return writeJSON(writer, http.StatusNotFound,
				APIerror{
					http.StatusNotFound,
					"Not found",
					"User does not exist",
				},
			)
		}

		chats, err := server.Storage.GetChats(ctx, request.PathValue("userid"), sessionUser.User.Id, limit, offset)
		if err != nil {
			if err == sql.ErrNoRows {
				return writeJSON(writer, http.StatusNotFound,
					APIerror{
						http.StatusNotFound,
						"Not found",
						"Chat not found",
					},
				)
			}
			return err
		}

		return writeJSON(writer, http.StatusOK, chats)
	}

	return writeJSON(writer, http.StatusMethodNotAllowed,
		APIerror{
			http.StatusMethodNotAllowed,
			"Method Not Allowed",
			"Method not Allowed",
		})
}
