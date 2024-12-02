package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"Social-Network-01/api/database"
	"Social-Network-01/api/types"
)

// Retrieve the group from the database using its id.
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
	sess, err := server.Sessions.GetSession(request)
	if err != nil {
		return err
	}

	if request.Method != http.MethodGet {
		return writeJSON(writer, http.StatusMethodNotAllowed, HTTPerror(http.StatusMethodNotAllowed))
	}

	groupid := request.PathValue("groupid")
	inGroup, err := server.Storage.UserInGroup(request.Context(), groupid, sess.User.Id)
	if err != nil {
		return err
	}

	if !inGroup {
		return writeJSON(writer, http.StatusUnauthorized, HTTPerror(http.StatusUnauthorized))
	}

	limit, offset := parseRequestLimitAndOffset(request)
	posts, err := server.Storage.GetGroupPosts(request.Context(), &groupid, limit, offset)
	if err != nil {
		return err
	}

	return writeJSON(writer, http.StatusOK, posts)
}

func (server *API) Groups(writer http.ResponseWriter, request *http.Request) (err error) {
	sess, err := server.Sessions.GetSession(request)
	if err != nil {
		return err
	}

	switch request.Method {
	case http.MethodPost:
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
		newGroup.Owner = sess.User.Id

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
	if request.Method != http.MethodPost {
		return writeJSON(writer, http.StatusMethodNotAllowed, HTTPerror(http.StatusMethodNotAllowed))
	}

	sess, err := server.Sessions.GetSession(request)
	if err != nil {
		return err
	}

	payload := new(struct {
		GroupId string `json:"groupId"`
		UserId  string `json:"userId"`
	})

	err = json.NewDecoder(request.Body).Decode(&payload)
	if err != nil {
		return err
	}

	if payload.GroupId == "" || payload.UserId == "" {
		return writeJSON(writer, http.StatusBadRequest, HTTPerror(http.StatusBadRequest))
	}

	hostInGroup, err := server.Storage.UserInGroup(request.Context(), payload.GroupId, sess.User.Id)
	if err != nil {
		return err
	}

	if !hostInGroup {
		return writeJSON(writer, http.StatusUnauthorized, HTTPerror(http.StatusUnauthorized))
	}

	guestInGroup, err := server.Storage.UserInGroup(request.Context(), payload.GroupId, payload.UserId)
	if err != nil {
		return err
	}
	if guestInGroup {
		return writeJSON(writer, http.StatusConflict, HTTPerror(http.StatusConflict))
	}

	return server.Storage.UserJoinGroup(request.Context(), payload.UserId, payload.GroupId, false)
}

func (server *API) RequestGroup(writer http.ResponseWriter, request *http.Request) (err error) {
	sess, err := server.Sessions.GetSession(request)
	if err != nil {
		return err
	}

	groupid := request.PathValue("groupid")

	err = server.Storage.UserJoinGroup(request.Context(), sess.User.Id, groupid, true)
	if err != nil {
		return err
	}

	return writeJSON(writer, http.StatusOK, "OK")
}

func (server *API) GetGroupInvites(writer http.ResponseWriter, request *http.Request) error {
	sess, err := server.Sessions.GetSession(request)
	if err != nil {
		return err
	}

	invites, err := server.Storage.GetGroupInvites(request.Context(), sess.User.Id)
	if err != nil {
		return err
	}

	return writeJSON(writer, http.StatusOK, invites)
}

func (server *API) AcceptGroupInvite(writer http.ResponseWriter, request *http.Request) error {
	sess, err := server.Sessions.GetSession(request)
	if err != nil {
		return err
	}

	groupid := request.PathValue("groupid")
	group, err := server.Storage.GetGroup(request.Context(), groupid)
	if err != nil {
		return err
	}
	if group.Owner != sess.User.Id {
		return writeJSON(writer, http.StatusUnauthorized, HTTPerror(http.StatusUnauthorized))
	}

	err = server.Storage.AcceptGroupInvite(request.Context(), sess.User.Id, groupid)
	if err != nil {
		return err
	}

	return writeJSON(writer, http.StatusOK, "OK")
}

func (server *API) DeclineGroupInvite(writer http.ResponseWriter, request *http.Request) error {
	sess, err := server.Sessions.GetSession(request)
	if err != nil {
		return err
	}

	groupid := request.PathValue("groupid")
	group, err := server.Storage.GetGroup(request.Context(), groupid)
	if err != nil {
		return err
	}
	if group.Owner != sess.User.Id {
		return writeJSON(writer, http.StatusUnauthorized, HTTPerror(http.StatusUnauthorized))
	}

	err = server.Storage.DeclineGroupInvite(request.Context(), sess.User.Id, groupid)
	if err != nil {
		return err
	}

	return writeJSON(writer, http.StatusOK, "OK")
}

func (server *API) GetGroupRequests(writer http.ResponseWriter, request *http.Request) error {
	sess, err := server.Sessions.GetSession(request)
	if err != nil {
		return err
	}

	invites, err := server.Storage.GetGroupRequests(request.Context(), sess.User.Id)
	if err != nil {
		return err
	}

	return writeJSON(writer, http.StatusOK, invites)
}
