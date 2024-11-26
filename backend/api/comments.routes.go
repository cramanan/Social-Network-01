package api

import (
	"Social-Network-01/api/types"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
)

func (server *API) Comment(writer http.ResponseWriter, request *http.Request) (err error) {
	sess, err := server.Sessions.GetSession(request)
	if err != nil {
		return err
	}

	switch request.Method {
	case http.MethodGet:
		limit, offset := parseRequestLimitAndOffset(request)
		comments, err := server.Storage.GetComments(request.Context(), request.PathValue("postid"), limit, offset)
		if err == sql.ErrNoRows {
			return HTTPerror(http.StatusNotFound)
		}
		if err != nil {
			return err
		}

		return writeJSON(writer, http.StatusOK, comments)

	case http.MethodPost:
		err = request.ParseMultipartForm(5 * (1 << 10))
		if err != nil {
			return err
		}

		data := request.MultipartForm.Value["data"]
		if len(data) != 1 {
			return fmt.Errorf("no data field in multipart form")
		}

		comment := new(types.Comment)
		comment.PostId = request.PathValue("postid")
		err = json.Unmarshal([]byte(data[0]), comment)
		if err != nil {
			return err
		}

		files, err := MultiPartFiles(request)
		if err != nil {
			return err
		}

		comment.UserId = sess.User.Id
		if len(files) >= 1 {
			comment.Image = files[0]
		}

		return server.Storage.CreateComment(request.Context(), comment)

	default:
		return writeJSON(writer, http.StatusMethodNotAllowed, "Method not allowed")
	}
}
