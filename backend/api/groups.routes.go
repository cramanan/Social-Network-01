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

// Retrieve the group from the database using its name.
//
// `server` is a pointer of the API type (see ./api/api.go). It contains a session reference.
func (server *API) Group(writer http.ResponseWriter, request *http.Request) error {
	if request.Method != http.MethodGet {
		return writeJSON(writer, http.StatusMethodNotAllowed, HTTPerror(http.StatusMethodNotAllowed))
	}

	groupid := request.PathValue("groupid")

	group, err := server.Storage.GetGroup(request.Context(), groupid)
	if err == sql.ErrNoRows {
		return writeJSON(writer, http.StatusNotFound, HTTPerror(http.StatusNotFound))
	}
	if err != nil {
		return err
	}

	return writeJSON(writer, http.StatusOK, group)
}

func (server *API) GetGroupPosts(writer http.ResponseWriter, request *http.Request) (err error) {
	if request.Method != http.MethodGet {
		return writeJSON(writer, http.StatusMethodNotAllowed, HTTPerror(http.StatusMethodNotAllowed))
	}

	limit, offset := parseRequestLimitAndOffset(request)

	posts, err := server.Storage.GetGroupPosts(request.Context(), request.PathValue("groupid"), limit, offset)
	if err != nil {
		return err
	}

	return writeJSON(writer, http.StatusOK, posts)
}

func (server *API) Groups(writer http.ResponseWriter, request *http.Request) (err error) {
	switch request.Method {
	case http.MethodPost:
		if request.Method != http.MethodPost {
			return writeJSON(writer, http.StatusMethodNotAllowed, HTTPerror(http.StatusMethodNotAllowed))
		}

		err = request.ParseMultipartForm(5 * (1 << 20))
		if err != nil {
			return err
		}

		data := request.MultipartForm.Value["data"]
		if len(data) != 1 {
			return fmt.Errorf("no data field in multipart form")
		}

		newGroup := new(types.Group)
		err = json.Unmarshal([]byte(data[0]), &newGroup)
		if err != nil {
			return err
		}

		files, err := MultiPartFiles(request)
		if err != nil {
			return err
		}
		if len(files) != 1 {
			files = append(files, "https://upload.wikimedia.org/wikipedia/commons/2/2c/Default_pfp.svg")
		}

		newGroup.Image = files[0]

		if newGroup.Name == "" || newGroup.Description == "" {
			return writeJSON(writer, http.StatusBadRequest, HTTPerror(http.StatusBadRequest, "All fields are required"))
		}

		err = server.Storage.NewGroup(request.Context(), newGroup)
		if err == database.ErrConflict {
			return writeJSON(writer, http.StatusConflict, HTTPerror(http.StatusConflict, "This group already exists"))
		}
		if err != nil {
			return err
		}

		return writeJSON(writer, http.StatusCreated, "Created")

	case http.MethodGet:
		limit, offset := parseRequestLimitAndOffset(request)
		groups, err := server.Storage.GetGroups(request.Context(), limit, offset)
		if err != nil {
			return err
		}
		return writeJSON(writer, http.StatusOK, groups)

	default:
		return writeJSON(writer, http.StatusMethodNotAllowed, HTTPerror(http.StatusMethodNotAllowed))
	}
}

func (server *API) InviteUserIntoGroup(writer http.ResponseWriter, request *http.Request) (err error) {
	sess, err := server.Sessions.GetSession(request)
	if err != nil {
		return err
	}

	payload := new(struct {
		GroupId string `json:"groupId"`
		UserId  string `json:"userId"`
	})

	allowed, err := server.Storage.AllowGroupInvite(context.TODO(), sess.User.Id, payload.UserId, payload.GroupId)
	if err != nil {
		return err
	}

	if !allowed {
		return writeJSON(writer, http.StatusMethodNotAllowed, HTTPerror(http.StatusMethodNotAllowed))
	}

	return server.Storage.InviteUserIntoGroup(context.TODO(), payload.UserId, payload.GroupId)
}
