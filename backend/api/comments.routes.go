package api

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"time"

	"Social-Network-01/api/database"
	"Social-Network-01/api/types"
)

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

func (server *API) Comment(writer http.ResponseWriter, request *http.Request) (err error) {
	sess, err := server.Sessions.GetSession(request)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(request.Context(), database.TransactionTimeout)
	defer cancel()

	switch request.Method {
	case http.MethodPost:
		// Multipart form handling
		err = request.ParseMultipartForm(5 * (1 << 20))
		if err != nil {
			return err
		}

		contents, ok := request.MultipartForm.Value["content"]
		if !ok || len(contents) <= 0 {
			return fmt.Errorf("no content")
		}

		postId, ok := request.MultipartForm.Value["postId"]
		if !ok || len(postId) != 1 {
			return fmt.Errorf("invalid number of post id")
		}

		images, ok := request.MultipartForm.File["images"]
		if !ok || len(images) != 1 {
			return fmt.Errorf("invalid number of images")
		}

		image := images[0]

		file, err := image.Open()
		if err != nil {
			return err
		}
		defer file.Close()

		temp, err := os.CreateTemp("api/images", fmt.Sprintf("*-%s", image.Filename))
		if err != nil {
			return err
		}
		defer temp.Close()

		_, err = temp.ReadFrom(file)
		if err != nil {
			return err
		}

		comment := types.Comment{
			UserId:    sess.User.Id,
			PostId:    postId[0],
			Content:   contents[0],
			Image:     fmt.Sprintf("/%s", temp.Name()),
			Timestamp: time.Now(),
		}

		return server.Storage.CreateComment(ctx, comment)

	default:
		return writeJSON(writer, http.StatusMethodNotAllowed, "Method not allowed")
	}
}
