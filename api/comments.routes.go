package api

import (
	"context"
	"database/sql"
	"net/http"

	"Social-Network-01/api/database"
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
