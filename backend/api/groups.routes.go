package api

import (
	"context"
	"database/sql"
	"encoding/json"
	"net/http"

	"Social-Network-01/api/database"
	"Social-Network-01/api/models"
)

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
