package api

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"Social-Network-01/api/database"
	"Social-Network-01/api/types"
)

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

	req := types.Post{}

	content, ok := request.MultipartForm.Value["content"]
	if !ok || len(content) < 1 {
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

		req.Images[idx] = fmt.Sprintf("/%s", temp.Name())
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

	// _, err = server.Sessions.GetSession(request)
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
