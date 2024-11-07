package api

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"Social-Network-01/api/database"
	"Social-Network-01/api/types"
)

// Create a new group in the database.
//
// `server` is a pointer of the API type (see ./api/api.go). It contains a session reference.
func (server *API) CreateGroup(writer http.ResponseWriter, request *http.Request) (err error) {
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

	newGroup := new(types.Group)
	err = request.ParseMultipartForm(5 * (1 << 20))
	if err != nil {
		return err
	}

	data := request.MultipartForm.Value["data"]
	if len(data) != 1 {
		return fmt.Errorf("no data field in multipart form")
	}

	err = json.Unmarshal([]byte(data[0]), newGroup)
	if err != nil {
		return err
	}

	files, err := MultiPartFiles(request)
	if len(files) != 1 {
		files = append(files, "https://upload.wikimedia.org/wikipedia/commons/2/2c/Default_pfp.svg")
	}

	newGroup.Image = files[0]

	if newGroup.Name == "" || newGroup.Description == "" {
		return writeJSON(writer, http.StatusBadRequest,
			APIerror{
				http.StatusBadRequest,
				"Bad Request",
				"All fields are required",
			})
	}

	group, err := server.Storage.NewGroup(ctx, newGroup)
	if err == database.ErrConflict {
		return writeJSON(writer, http.StatusConflict, APIerror{
			http.StatusConflict,
			"Conflict",
			"This group already exists",
		})
	}
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
