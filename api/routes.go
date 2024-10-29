package api

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/mail"
	"os"

	"Social-Network-01/api/database"
	"Social-Network-01/api/models"
)

//////////////////////////////////////////////////////////////////////////////////////////
// 						▗▖ ▗▖ ▗▄▄▖▗▄▄▄▖▗▄▄▖  ▗▄▄▖										//
// 						▐▌ ▐▌▐▌   ▐▌   ▐▌ ▐▌▐▌											//
// 						▐▌ ▐▌ ▝▀▚▖▐▛▀▀▘▐▛▀▚▖ ▝▀▚▖										//
// 						▝▚▄▞▘▗▄▄▞▘▐▙▄▄▖▐▌ ▐▌▗▄▄▞▘										//
//////////////////////////////////////////////////////////////////////////////////////////

// Perform the action of registering one user in the database.
//
// `server` is a pointer of the API type (see ./api/api.go). It contains a session reference.
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

// Perform the action of logging one user.
//
// `server` is a pointer of the API type (see ./api/api.go). It contains a session reference.
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

// This method acts as a router for different HTTP methods.
//
// `server` is a pointer of the API type (see ./api/api.go). It contains a session reference.
func (server *API) User(writer http.ResponseWriter, request *http.Request) error {
	ctx, cancel := context.WithTimeout(request.Context(), database.TransactionTimeout)
	defer cancel()

	switch request.Method {
	case http.MethodGet:
		user, err := server.Storage.GetUser(ctx, request.PathValue("userid"))
		if err == sql.ErrNoRows {
			return writeJSON(writer, http.StatusNotFound,
				APIerror{
					http.StatusNotFound,
					"Not found",
					"User not found",
				},
			)
		}
		if err != nil {
			return err
		}

		return writeJSON(writer, http.StatusOK, user)

	case http.MethodDelete:

		sess, err := server.Sessions.GetSession(request)
		if err != nil {
			return err
		}

		if sess.User.Id != request.PathValue("userid") {
			return writeJSON(writer, http.StatusUnauthorized,
				APIerror{
					http.StatusUnauthorized,
					"Unauthorized",
					"You are not authorized to perform this action.",
				})
		}

		err = server.Storage.DeleteUser(ctx, sess.User.Id)
		if err != nil {
			return err
		}

		return writeJSON(writer, http.StatusNoContent, "")
	default:
		return writeJSON(writer, http.StatusMethodNotAllowed,
			APIerror{
				http.StatusMethodNotAllowed,
				"Method Not Allowed",
				"Method not Allowed",
			})
	}
}

func (server *API) GetUser(writer http.ResponseWriter, request *http.Request) error {
	if request.Method != http.MethodGet {
		return writeJSON(writer, http.StatusMethodNotAllowed,
			APIerror{
				http.StatusMethodNotAllowed,
				"Method Not Allowed",
				"Method not Allowed",
			})
	}

	s, err := server.Sessions.GetSession(request)
	if err != nil {
		return writeJSON(writer, http.StatusUnauthorized, "You are unauthorized to access this ressource.")
	}

	return writeJSON(writer, http.StatusOK, s.User)
}

//////////////////////////////////////////////////////////////////////////////////////////
//						▗▄▄▄▖ ▗▄▖ ▗▖   ▗▖    ▗▄▖ ▗▖ ▗▖ ▗▄▄▖								//
//						▐▌   ▐▌ ▐▌▐▌   ▐▌   ▐▌ ▐▌▐▌ ▐▌▐▌								//
//						▐▛▀▀▘▐▌ ▐▌▐▌   ▐▌   ▐▌ ▐▌▐▌ ▐▌ ▝▀▚▖								//
//						▐▌   ▝▚▄▞▘▐▙▄▄▖▐▙▄▄▖▝▚▄▞▘▐▙█▟▌▗▄▄▞▘								//
//////////////////////////////////////////////////////////////////////////////////////////

// Perform the action of following one from another.
//
// `server` is a pointer of the API type (see ./api/api.go). It contains a session reference.
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

//////////////////////////////////////////////////////////////////////////////////////////
// 						▗▄▄▖  ▗▄▖  ▗▄▄▖▗▄▄▄▖▗▄▄▖										//
// 						▐▌ ▐▌▐▌ ▐▌▐▌     █ ▐▌   										//
// 						▐▛▀▘ ▐▌ ▐▌ ▝▀▚▖  █  ▝▀▚▖										//
// 						▐▌   ▝▚▄▞▘▗▄▄▞▘  █ ▗▄▄▞▘										//
//////////////////////////////////////////////////////////////////////////////////////////

func (server *API) CreatePost(writer http.ResponseWriter, request *http.Request) (err error) {
	if request.Method != http.MethodPost {
		return err
	}
	sess, err := server.Sessions.GetSession(request)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(request.Context(), database.TransactionTimeout)
	defer cancel()

	err = request.ParseMultipartForm(5 * (1 << 20))
	if err != nil {
		return err
	}
	req := new(models.PostRequest)

	content, ok := request.MultipartForm.Value["content"]
	if !ok || content == nil {
		return writeJSON(writer, http.StatusBadRequest,
			APIerror{
				http.StatusBadRequest,
				"Bad Request",
				"No content in request",
			})
	}

	req.Content = content[0]
	if req.Content == "" {
		return writeJSON(writer, http.StatusBadRequest,
			APIerror{
				http.StatusBadRequest,
				"Bad Request",
				"Content is empty",
			})
	}

	req.UserId = sess.User.Id

	multipartImages := request.MultipartForm.File["images"]

	req.Images = make([]string, len(multipartImages))

	for idx, fileHeader := range multipartImages {
		file, err := fileHeader.Open()
		if err != nil {
			return err
		}
		defer file.Close()

		temp, err := os.CreateTemp("api/images", fmt.Sprintf("*-%s", fileHeader.Filename))
		if err != nil {
			return err
		}
		defer temp.Close()

		_, err = temp.ReadFrom(file)
		if err != nil {
			return err
		}

		req.Images[idx] = temp.Name()
	}

	err = server.Storage.CreatePost(ctx, req)
	if err != nil {
		return err
	}

	return writeJSON(writer, http.StatusCreated, "Created")
}

func (server *API) Post(writer http.ResponseWriter, request *http.Request) (err error) {
	ctx, cancel := context.WithTimeout(request.Context(), database.TransactionTimeout)
	defer cancel()

	// sess, err := server.Sessions.GetSession(request)
	// if err != nil {
	// 	return err
	// }

	switch request.Method {
	case http.MethodGet:
		post, err := server.Storage.GetPost(ctx, request.PathValue("postid"))
		if err != nil {
			return err
		}

		return writeJSON(writer, http.StatusOK, post)

	default:
		return writeJSON(writer, http.StatusMethodNotAllowed,
			APIerror{
				http.StatusMethodNotAllowed,
				"Method Not Allowed",
				"Method not Allowed",
			})
	}
}

func (server *API) LikePost(writer http.ResponseWriter, request *http.Request) (err error) {
	if request.Method != http.MethodPost {
		return writeJSON(writer, http.StatusMethodNotAllowed,
			APIerror{
				http.StatusMethodNotAllowed,
				"Method Not Allowed",
				"Method not Allowed",
			})
	}

	ctx, cancel := context.WithTimeout(request.Context(), database.TransactionTimeout)
	defer cancel()

	sess, err := server.Sessions.GetSession(request)
	if err != nil {
		return err
	}

	err = server.Storage.LikePost(ctx, sess.User.Id, request.PathValue("postid"))
	if err != nil {
		return err
	}

	return writeJSON(writer, http.StatusOK, "OK")
}

// // Retrieve all posts of one user from the database.
// //
// // `server` is a pointer of the API type (see ./api/api.go). It contains a session reference.
// func (server *API) AllPostsFromOneUser(writer http.ResponseWriter, request *http.Request) error {
// 	ctx, cancel := context.WithTimeout(request.Context(), database.TransactionTimeout)
// 	defer cancel()
// 	if request.Method == http.MethodGet {

// 		limit, offset := parseRequestLimitAndOffset(request)
// 		posts, err := server.Storage.GetAllPostsFromOneUser(ctx, request.PathValue("userid"), limit, offset)
// 		if err != nil {
// 			if err == sql.ErrNoRows {
// 				return writeJSON(writer, http.StatusNotFound,
// 					APIerror{
// 						http.StatusNotFound,
// 						"Not found",
// 						"Posts not found",
// 					},
// 				)
// 			}
// 			return err
// 		}
// 		return writeJSON(writer, http.StatusOK, posts)
// 	}

// 	return writeJSON(writer, http.StatusMethodNotAllowed,
// 		APIerror{
// 			http.StatusMethodNotAllowed,
// 			"Method Not Allowed",
// 			"Method not Allowed",
// 		})
// }

// // Retrieve all posts of one group from the database.
// //
// // `server` is a pointer of the API type (see ./api/api.go). It contains a session reference.
// func (server *API) GetAllPostsFromOneGroup(writer http.ResponseWriter, request *http.Request) error {
// 	ctx, cancel := context.WithTimeout(request.Context(), database.TransactionTimeout)
// 	defer cancel()
// 	if request.Method == http.MethodGet {

// 		limit, offset := parseRequestLimitAndOffset(request)
// 		posts, err := server.Storage.GetGroupPosts(ctx, request.PathValue("groupid"), limit, offset)
// 		if err != nil {
// 			if err == sql.ErrNoRows {
// 				return writeJSON(writer, http.StatusNotFound,
// 					APIerror{
// 						http.StatusNotFound,
// 						"Not found",
// 						"Group not found",
// 					},
// 				)
// 			}
// 			return err
// 		}

// 		return writeJSON(writer, http.StatusOK, posts)
// 	}

// 	return writeJSON(writer, http.StatusMethodNotAllowed,
// 		APIerror{
// 			http.StatusMethodNotAllowed,
// 			"Method Not Allowed",
// 			"Method not Allowed",
// 		})
// }

// // Retrieve all posts of a user's follows from the database.
// //
// // `server` is a pointer of the API type (see ./api/api.go). It contains a session reference.
// func (server *API) GetAllPostsFromOneUsersFollows(writer http.ResponseWriter, request *http.Request) error {
// 	ctx, cancel := context.WithTimeout(request.Context(), database.TransactionTimeout)
// 	defer cancel()
// 	if request.Method == http.MethodGet {

// 		limit, offset := parseRequestLimitAndOffset(request)
// 		posts, err := server.Storage.GetFollowsPosts(ctx, request.PathValue("userid"), limit, offset)
// 		if err != nil {
// 			if err == sql.ErrNoRows {
// 				return writeJSON(writer, http.StatusNotFound,
// 					APIerror{
// 						http.StatusNotFound,
// 						"Not found",
// 						"User not found",
// 					},
// 				)
// 			}
// 			return err
// 		}

// 		return writeJSON(writer, http.StatusOK, posts)
// 	}

// 	return writeJSON(writer, http.StatusMethodNotAllowed,
// 		APIerror{
// 			http.StatusMethodNotAllowed,
// 			"Method Not Allowed",
// 			"Method not Allowed",
// 		})
// }

// // Retrieve all posts of ones likes from the database.
// //
// // `server` is a pointer of the API type (see ./api/api.go). It contains a session reference.
// func (server *API) GetAllPostsFromOneUsersLikes(writer http.ResponseWriter, request *http.Request) error {
// 	ctx, cancel := context.WithTimeout(request.Context(), database.TransactionTimeout)
// 	defer cancel()
// 	if request.Method == http.MethodGet {

// 		userError := func(err error) error {
// 			if err == sql.ErrNoRows {
// 				return writeJSON(writer, http.StatusNotFound,
// 					APIerror{
// 						http.StatusNotFound,
// 						"Not found",
// 						"User not found",
// 					},
// 				)
// 			}
// 			return err
// 		}

// 		user, err := server.Storage.GetUser(ctx, request.PathValue("userid"))
// 		if err != nil {
// 			return userError(err)
// 		}

// 		sessionUser, err := server.Sessions.GetSession(request)

// 		if user.Private && (err != nil || sessionUser.User.Id != user.Id) {
// 			return writeJSON(writer, http.StatusUnauthorized, APIerror{
// 				http.StatusUnauthorized,
// 				"Unauthorized",
// 				"This account is private",
// 			})
// 		}

// 		limit, offset := parseRequestLimitAndOffset(request)
// 		posts, err := server.Storage.GetPostsLike(ctx, request.PathValue("userid"), limit, offset)
// 		if err != nil {
// 			return userError(err)
// 		}

// 		return writeJSON(writer, http.StatusOK, posts)
// 	}

// 	return writeJSON(writer, http.StatusMethodNotAllowed,
// 		APIerror{
// 			http.StatusMethodNotAllowed,
// 			"Method Not Allowed",
// 			"Method not Allowed",
// 		})
// }

//////////////////////////////////////////////////////////////////////////////////////////
// 						 ▗▄▄▖ ▗▄▖ ▗▖  ▗▖▗▖  ▗▖▗▄▄▄▖▗▖  ▗▖▗▄▄▄▖▗▄▄▖						//
// 						▐▌   ▐▌ ▐▌▐▛▚▞▜▌▐▛▚▞▜▌▐▌   ▐▛▚▖▐▌  █ ▐▌   						//
// 						▐▌   ▐▌ ▐▌▐▌  ▐▌▐▌  ▐▌▐▛▀▀▘▐▌ ▝▜▌  █  ▝▀▚▖						//
// 						▝▚▄▄▖▝▚▄▞▘▐▌  ▐▌▐▌  ▐▌▐▙▄▄▖▐▌  ▐▌  █ ▗▄▄▞▘						//
//////////////////////////////////////////////////////////////////////////////////////////

// Retrieve all comments of one post from the database.
//
// `server` is a pointer of the API type (see ./api/api.go). It contains a session reference.
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

//////////////////////////////////////////////////////////////////////////////////////////
// 						 ▗▄▄▖▗▖ ▗▖ ▗▄▖▗▄▄▄▖▗▄▄▖											//
// 						▐▌   ▐▌ ▐▌▐▌ ▐▌ █ ▐▌   											//
// 						▐▌   ▐▛▀▜▌▐▛▀▜▌ █  ▝▀▚▖											//
// 						▝▚▄▄▖▐▌ ▐▌▐▌ ▐▌ █ ▗▄▄▞▘											//
//////////////////////////////////////////////////////////////////////////////////////////

func (server *API) Socket(writer http.ResponseWriter, request *http.Request) {
	_, err := server.Sessions.GetSession(request)
	if err != nil {
		writeJSON(writer, http.StatusUnauthorized, "Unauthorized")
		return
	}

	conn, err := server.WSUpgrader.Upgrade(writer, request, nil)
	if err != nil {
		log.Println(err)
		return
	}

	for {
		value := new(any)
		err = conn.ReadJSON(value)
		if err != nil {
			log.Println(err)
			return
		}

		err = conn.WriteJSON(value)
		if err != nil {
			log.Println(err)
			return
		}
	}
}

// Retrieve all chats beetween 2 users from the database.
//
// `server` is a pointer of the API type (see ./api/api.go). It contains a session reference.
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

// Retrieve all chats beetween 2 users from the database using their userIds.
//
// `server` is a pointer of the API type (see ./api/api.go). It contains a session reference.
func (server *API) GetChatFromGroup(writer http.ResponseWriter, request *http.Request) error {
	ctx, cancel := context.WithTimeout(request.Context(), database.TransactionTimeout)
	defer cancel()

	groupname := request.PathValue("groupname")

	if request.Method != http.MethodGet {
		return writeJSON(writer, http.StatusMethodNotAllowed,
			APIerror{
				http.StatusMethodNotAllowed,
				"Method Not Allowed",
				"Method not Allowed",
			})
	}

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
	limit, offset := parseRequestLimitAndOffset(request)

	chats, err := server.Storage.GetChats(ctx, groupname, sessionUser.User.Id, limit, offset)
	if err == sql.ErrNoRows {
		return writeJSON(writer, http.StatusNotFound,
			APIerror{
				http.StatusNotFound,
				"Not found",
				"Chat not found",
			},
		)
	}
	if err != nil {
		return err
	}

	return writeJSON(writer, http.StatusOK, chats)
}

//////////////////////////////////////////////////////////////////////////////////////////
// 						 ▗▄▄▖▗▄▄▖  ▗▄▖ ▗▖ ▗▖▗▄▄▖  ▗▄▄▖									//
// 						▐▌   ▐▌ ▐▌▐▌ ▐▌▐▌ ▐▌▐▌ ▐▌▐▌										//
// 						▐▌▝▜▌▐▛▀▚▖▐▌ ▐▌▐▌ ▐▌▐▛▀▘  ▝▀▚▖									//
// 						▝▚▄▞▘▐▌ ▐▌▝▚▄▞▘▝▚▄▞▘▐▌   ▗▄▄▞▘									//
//////////////////////////////////////////////////////////////////////////////////////////

// Create a new group in the database.
//
// `server` is a pointer of the API type (see ./api/api.go). It contains a session reference.
func (server *API) CreateGroup(writer http.ResponseWriter, request *http.Request) error {
	ctx, cancel := context.WithTimeout(request.Context(), database.TransactionTimeout)
	defer cancel()

	if request.Method != http.MethodPost {
		return writeJSON(writer, http.StatusMethodNotAllowed,
			APIerror{
				http.StatusMethodNotAllowed,
				"Method Not Allowed",
				"Method not Allowed",
			})
	}

	newGroup := new(models.Group)
	err := json.NewDecoder(request.Body).Decode(newGroup)
	if err != nil {
		return writeJSON(writer, http.StatusUnprocessableEntity,
			APIerror{
				http.StatusUnprocessableEntity,
				"Unprocessable Entity",
				"Could not process register request",
			})
	}

	if newGroup.Name == "" ||
		newGroup.Description == "" {
		return writeJSON(writer, http.StatusBadRequest,
			APIerror{
				http.StatusBadRequest,
				"Bad Request",
				"All fields are required",
			})
	}

	group, err := server.Storage.NewGroup(ctx, newGroup)
	if err != nil {
		return err
	}

	return writeJSON(writer, http.StatusOK, group)
}

// Retrieve the group from the database using its name.
//
// `server` is a pointer of the API type (see ./api/api.go). It contains a session reference.
func (server *API) Group(writer http.ResponseWriter, request *http.Request) error {
	if request.Method != http.MethodGet {
		return writeJSON(writer, http.StatusMethodNotAllowed,
			APIerror{
				http.StatusMethodNotAllowed,
				"Method Not Allowed",
				"Method not Allowed",
			})
	}

	ctx, cancel := context.WithTimeout(request.Context(), database.TransactionTimeout)
	defer cancel()

	groupname := request.PathValue("groupname")

	group, err := server.Storage.GetGroup(ctx, groupname)
	if err == sql.ErrNoRows {
		return writeJSON(writer, http.StatusNotFound,
			APIerror{
				http.StatusNotFound,
				"Not found",
				"Chat not found",
			},
		)
	}
	if err != nil {
		return err
	}

	return writeJSON(writer, http.StatusOK, group)
}

func (server *API) GetGroupPosts(writer http.ResponseWriter, request *http.Request) (err error) {
	if request.Method != http.MethodGet {
		return writeJSON(writer, http.StatusMethodNotAllowed,
			APIerror{
				http.StatusMethodNotAllowed,
				"Method Not Allowed",
				"Method not Allowed",
			})
	}

	ctx, cancel := context.WithTimeout(request.Context(), database.TransactionTimeout)
	defer cancel()

	limit, offset := parseRequestLimitAndOffset(request)

	posts, err := server.Storage.GetGroupPosts(ctx, request.PathValue("groupid"), limit, offset)
	if err != nil {
		return err
	}

	return writeJSON(writer, http.StatusOK, posts)
}

func (server *API) GetGroups(writer http.ResponseWriter, request *http.Request) (err error) {
	ctx, cancel := context.WithTimeout(request.Context(), database.TransactionTimeout)
	defer cancel()

	limit, offset := parseRequestLimitAndOffset(request)

	groups, err := server.Storage.GetGroups(ctx, limit, offset)
	if err != nil {
		return err
	}

	return writeJSON(writer, http.StatusOK, groups)
}
