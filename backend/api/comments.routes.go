package api

import (
	"Social-Network-01/api/types"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
)

// Comment handles the HTTP requests related to comments on a post.
// It supports both GET (to fetch comments) and POST (to create a new comment) methods.
func (server *API) Comment(writer http.ResponseWriter, request *http.Request) (err error) {
	// Get the current session from the request to identify the user making the request.
	sess, err := server.Sessions.GetSession(request)
	if err != nil {
		// Return if the session cannot be retrieved (e.g., user not logged in).
		return err 
	}

	// Switch based on the HTTP method of the request (GET or POST).
	switch request.Method {
	// Handle GET request to fetch comments for a specific post.
	case http.MethodGet:
		// Parse the limit and offset from the request to handle pagination.
		limit, offset := parseRequestLimitAndOffset(request)

		// Fetch the comments from the database based on the post ID and pagination parameters.
		comments, err := server.Storage.GetComments(request.Context(), request.PathValue("postid"), limit, offset)
		if err == sql.ErrNoRows {
			// If no comments are found, return a Not Found error.
			return HTTPerror(http.StatusNotFound)
		}
		if err != nil {
			// Return any other errors encountered while fetching comments.
			return err
		}

		// Write the comments as a JSON response with HTTP status OK.
		return writeJSON(writer, http.StatusOK, comments)

	// Handle POST request to create a new comment.
	case http.MethodPost:
		// Parse the multipart form data from the request (up to 5 MB).
		err = request.ParseMultipartForm(5 * (1 << 10))
		if err != nil {
			// Return if there is an error parsing the form data.
			return err 
		}

		// Extract the "data" field from the multipart form data (this should contain the comment details).
		data := request.MultipartForm.Value["data"]
		if len(data) != 1 {
			// If there is no or more than one "data" field, return an error.
			return fmt.Errorf("no data field in multipart form")
		}

		// Create a new Comment struct and unmarshal the JSON data into it.
		comment := new(types.Comment)
		// Set the post ID from the URL path.
		comment.PostId = request.PathValue("postid") 
		err = json.Unmarshal([]byte(data[0]), comment)
		if err != nil {
			// Return if there is an error unmarshaling the JSON data into the Comment struct.
			return err 
		}

		// Handle file uploads from the multipart form (e.g., images attached to the comment).
		files, err := MultiPartFiles(request)
		if err != nil {
			// Return if there is an error processing the files.
			return err 
		}

		// Set the UserId field of the comment to the current user's ID.
		comment.UserId = sess.User.Id

		// If files were uploaded (i.e., images), assign the first file as the comment's image.
		if len(files) >= 1 {
			comment.Image = files[0]
		}

		// Store the new comment in the database.
		return server.Storage.CreateComment(request.Context(), comment)

	// Handle any unsupported HTTP methods (e.g., PUT, DELETE).
	default:
		// Return a Method Not Allowed error for unsupported HTTP methods.
		return writeJSON(writer, http.StatusMethodNotAllowed, "Method not allowed")
	}
}

