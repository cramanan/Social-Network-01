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

func (server *API) ProfilePosts(writer http.ResponseWriter, request *http.Request) (err error) {
	sess, err := server.Sessions.GetSession(request)
	if err != nil {
		return err
	}

	limit, offset := parseRequestLimitAndOffset(request)

	ctx, cancel := context.WithTimeout(request.Context(), database.TransactionTimeout)
	defer cancel()

	posts, err := server.Storage.GetUserPosts(ctx, sess.User.Id, limit, offset)
	if err != nil {
		return err
	}
	return writeJSON(writer, http.StatusOK, posts)
}
