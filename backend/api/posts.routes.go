package api

import (
	"encoding/json"
	"net/http"

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

	err = request.ParseMultipartForm(5 * (1 << 20))
	if err != nil {
		return err
	}

	post := new(types.Post)

	data, ok := request.MultipartForm.Value["data"]
	if !ok || len(data) < 1 {
		return writeJSON(writer, http.StatusBadRequest, HTTPerror(http.StatusBadRequest, "No content in request"))
	}

	err = json.Unmarshal([]byte(data[0]), post)
	if err != nil {
		return err
	}

	post.UserId = sess.User.Id

	post.Images, err = MultiPartFiles(request)
	if err != nil {
		return err
	}

	err = server.Storage.CreatePost(request.Context(), post)
	if err != nil {
		return err
	}

	return writeJSON(writer, http.StatusCreated, "Created")
}

func (server *API) GetPostById(writer http.ResponseWriter, request *http.Request) (err error) {
	// For visibility:
	// _, err = server.Sessions.GetSession(request)
	// if err != nil {
	// 	return err
	// }

	if request.Method != http.MethodGet {
		return writeJSON(writer, http.StatusMethodNotAllowed, HTTPerror(http.StatusMethodNotAllowed))
	}

	post, err := server.Storage.GetPost(request.Context(), request.PathValue("postid"))
	if err != nil {
		return err
	}

	return writeJSON(writer, http.StatusOK, post)
}

func (server *API) LikePost(writer http.ResponseWriter, request *http.Request) (err error) {
	if request.Method != http.MethodPost {
		return writeJSON(writer, http.StatusMethodNotAllowed, HTTPerror(http.StatusMethodNotAllowed))
	}

	sess, err := server.Sessions.GetSession(request)
	if err != nil {
		return err
	}

	err = server.Storage.LikePost(request.Context(), sess.User.Id, request.PathValue("postid"))
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

	posts, err := server.Storage.GetUserPosts(request.Context(), sess.User.Id, limit, offset)
	if err != nil {
		return err
	}
	return writeJSON(writer, http.StatusOK, posts)
}
