package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"Social-Network-01/api/database"
	"Social-Network-01/api/types"
)

// Group retrieves a group from the database using its ID (groupid) from the request path.
// It checks if the request method is GET and returns the group details or an error.
func (server *API) Group(writer http.ResponseWriter, request *http.Request) error {
	// Check if the request method is GET, if not return a 405 Method Not Allowed error.
	if request.Method != http.MethodGet {
		return writeJSON(writer, http.StatusMethodNotAllowed, HTTPerror(http.StatusMethodNotAllowed))
	}

	// Retrieve the group ID from the request path.
	groupid := request.PathValue("groupid")

	// Fetch the group from the storage (database).
	group, err := server.Storage.GetGroup(request.Context(), groupid)
	if err == sql.ErrNoRows {
		// If no group is found, return a 404 Not Found error.
		return writeJSON(writer, http.StatusNotFound, HTTPerror(http.StatusNotFound))
	}
	if err != nil {
		// Return other errors encountered while fetching the group.
		return err
	}

	// Return the group details as a JSON response with a 200 OK status.
	return writeJSON(writer, http.StatusOK, group)
}

// GetGroupPosts retrieves the posts of a group. It supports pagination with limit and offset.
func (server *API) GetGroupPosts(writer http.ResponseWriter, request *http.Request) (err error) {
	// Check if the request method is GET, if not return a 405 Method Not Allowed error.
	if request.Method != http.MethodGet {
		return writeJSON(writer, http.StatusMethodNotAllowed, HTTPerror(http.StatusMethodNotAllowed))
	}

	// Parse the limit and offset for pagination from the request.
	limit, offset := parseRequestLimitAndOffset(request)

	// Fetch the posts of the group from the storage.
	posts, err := server.Storage.GetGroupPosts(request.Context(), request.PathValue("groupid"), limit, offset)
	if err != nil {
		// Return any error encountered while fetching the posts.
		return err
	}

	// Return the posts as a JSON response with a 200 OK status.
	return writeJSON(writer, http.StatusOK, posts)
}

// Groups handles both creating a new group (POST method) and fetching existing groups (GET method).
func (server *API) Groups(writer http.ResponseWriter, request *http.Request) (err error) {
	// Retrieve the session for the current user.
	sess, err := server.Sessions.GetSession(request)
	if err != nil {
		// Return error if session retrieval fails.
		return err 
	}

	switch request.Method {
	case http.MethodPost:
		// Handle group creation. Parse the multipart form for group data and file uploads.
		err = request.ParseMultipartForm(5 * (1 << 20)) // Max 5MB for form data.
		if err != nil {
			// Return error if parsing the multipart form fails.
			return err 
		}

		// Retrieve the group data from the multipart form.
		data := request.MultipartForm.Value["data"]
		if len(data) != 1 {
			return fmt.Errorf("no data field in multipart form") // Return error if no group data is found.
		}

		// Deserialize the group data into the Group struct.
		newGroup := new(types.Group)
		err = json.Unmarshal([]byte(data[0]), &newGroup)
		if err != nil {
			// Return error if deserialization fails.
			return err 
		}

		// Retrieve any uploaded files (e.g., group image).
		files, err := MultiPartFiles(request)
		if err != nil {
			// Return error if file retrieval fails.
			return err 
		}
		// If no image is uploaded, use a default image.
		if len(files) != 1 {
			files = append(files, "https://upload.wikimedia.org/wikipedia/commons/2/2c/Default_pfp.svg")
		}

		// Set the group image and the current user as the owner.
		newGroup.Image = files[0]
		newGroup.Owner = sess.User.Id

		// Ensure required fields are not empty.
		if newGroup.Name == "" || newGroup.Description == "" {
			return writeJSON(writer, http.StatusBadRequest, HTTPerror(http.StatusBadRequest, "All fields are required"))
		}

		// Create the new group in the database.
		err = server.Storage.NewGroup(request.Context(), newGroup)
		if err == database.ErrConflict {
			// Return error if a group with the same name already exists.
			return writeJSON(writer, http.StatusConflict, HTTPerror(http.StatusConflict, "This group already exists"))
		}
		if err != nil {
			// Return any other error encountered during group creation.
			return err 
		}

		// Return a success response with HTTP Status Created (201).
		return writeJSON(writer, http.StatusCreated, "Created")

	case http.MethodGet:
		// Handle fetching existing groups. Parse the limit and offset for pagination.
		limit, offset := parseRequestLimitAndOffset(request)

		// Fetch the groups from the storage.
		groups, err := server.Storage.GetGroups(request.Context(), limit, offset)
		if err != nil {
			// Return error if fetching groups fails.
			return err 
		}
		// Return the list of groups as a JSON response with HTTP Status OK.
		return writeJSON(writer, http.StatusOK, groups)

	default:
		// Return a 405 Method Not Allowed error for unsupported HTTP methods.
		return writeJSON(writer, http.StatusMethodNotAllowed, HTTPerror(http.StatusMethodNotAllowed))
	}
}

// InviteUserIntoGroup allows a user to invite another user to join a group.
func (server *API) InviteUserIntoGroup(writer http.ResponseWriter, request *http.Request) (err error) {
	// Check if the request method is POST, if not return a 405 Method Not Allowed error.
	if request.Method != http.MethodPost {
		return writeJSON(writer, http.StatusMethodNotAllowed, HTTPerror(http.StatusMethodNotAllowed))
	}

	// Retrieve the session for the current user.
	sess, err := server.Sessions.GetSession(request)
	if err != nil {
		// Return error if session retrieval fails.
		return err 
	}

	// Parse the incoming JSON payload for the group ID and user ID.
	payload := new(struct {
		GroupId string `json:"groupId"`
		UserId  string `json:"userId"`
	})

	// Check if the user has permission to invite another user to the group.
	allowed, err := server.Storage.AllowGroupInvite(request.Context(), sess.User.Id, payload.UserId, payload.GroupId)
	if err != nil {
		// Return error if checking permission fails.
		return err 
	}

	if !allowed {
		// Return a 401 Unauthorized error if the user is not allowed to invite the other user.
		return writeJSON(writer, http.StatusUnauthorized, HTTPerror(http.StatusUnauthorized))
	}

	// Add the user to the group (invited user).
	return server.Storage.UserJoinGroup(request.Context(), payload.UserId, payload.GroupId, false)
}

// RequestGroup allows a user to request to join a group.
func (server *API) RequestGroup(writer http.ResponseWriter, request *http.Request) (err error) {
	// Retrieve the session for the current user.
	sess, err := server.Sessions.GetSession(request)
	if err != nil {
		// Return error if session retrieval fails.
		return err 
	}

	// Retrieve the group ID from the request path.
	groupid := request.PathValue("groupid")

	// Check if the user is allowed to request to join the group.
	allowed, err := server.Storage.AllowGroupRequest(request.Context(), groupid, sess.User.Id)
	if err != nil {
		// Return error if checking permission fails.
		return err 
	}

	if !allowed {
		// Return a 401 Unauthorized error if the user is not allowed to request to join the group.
		return writeJSON(writer, http.StatusUnauthorized, HTTPerror(http.StatusUnauthorized))
	}

	// Add the user to the group as a request (pending approval).
	err = server.Storage.UserJoinGroup(request.Context(), sess.User.Id, groupid, true)
	if err != nil {
		// Return error if joining the group fails.
		return err 
	}

	// Return a success response with HTTP Status OK.
	return writeJSON(writer, http.StatusOK, "OK")
}
