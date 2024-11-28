package api

import (
	"encoding/json"
	"net/http"

	"Social-Network-01/api/types"
)

	// CreatePost handles the creation of a new post by the user. 
// It processes a multipart form request, extracts post data, and stores the post in the database.
func (server *API) CreatePost(writer http.ResponseWriter, request *http.Request) (err error) {
    // Ensure that the HTTP method is POST, as we are creating a new post
    if request.Method != http.MethodPost {
        return err
    }

    // Retrieve the session of the user making the request
    sess, err := server.Sessions.GetSession(request)
    if err != nil {
		// Return an error if session retrieval fails
        return err 
    }

    // Parse the multipart form data (max 5MB in size)
    err = request.ParseMultipartForm(5 * (1 << 20))
    if err != nil {
		// Return an error if parsing the form fails
        return err 
    }

    post := new(types.Post)

    // Check if the form contains the "data" field with post details
    data, ok := request.MultipartForm.Value["data"]
    if !ok || len(data) < 1 {
        return writeJSON(writer, http.StatusBadRequest, HTTPerror(http.StatusBadRequest, "No content in request"))
    }

    // Unmarshal the "data" field into a Post object
    err = json.Unmarshal([]byte(data[0]), post)
    if err != nil {
		// Return an error if unmarshalling fails
        return err 
    }

    // Associate the post with the user making the request
    post.UserId = sess.User.Id

    // Extract any images attached to the post
    post.Images, err = MultiPartFiles(request)
    if err != nil {
		// Return an error if there is an issue extracting the files
        return err 
    }

    // Save the post in the database
    err = server.Storage.CreatePost(request.Context(), post)
    if err != nil {
		// Return an error if saving the post fails
        return err 
    }

    // Respond with a success message and HTTP status code 201 (Created)
    return writeJSON(writer, http.StatusCreated, "Created")
}

// GetPostById retrieves a post by its unique ID from the database.
func (server *API) GetPostById(writer http.ResponseWriter, request *http.Request) (err error) {
    // Ensure that the HTTP method is GET, as we are retrieving data
    if request.Method != http.MethodGet {
        return writeJSON(writer, http.StatusMethodNotAllowed, HTTPerror(http.StatusMethodNotAllowed))
    }

    // Retrieve the post ID from the request path
    post, err := server.Storage.GetPost(request.Context(), request.PathValue("postid"))
    if err != nil {
		// Return an error if the post is not found or other issues occur
        return err 
    }

    // Respond with the post data
    return writeJSON(writer, http.StatusOK, post)
}

// LikePost allows a user to like a post by its unique ID.
func (server *API) LikePost(writer http.ResponseWriter, request *http.Request) (err error) {
    // Ensure that the HTTP method is POST, as we are performing an action (liking a post)
    if request.Method != http.MethodPost {
        return writeJSON(writer, http.StatusMethodNotAllowed, HTTPerror(http.StatusMethodNotAllowed))
    }

    // Retrieve the session of the user making the request
    sess, err := server.Sessions.GetSession(request)
    if err != nil {
		// Return an error if session retrieval fails
        return err 
    }

    // Like the post by calling the storage layer
    err = server.Storage.LikePost(request.Context(), sess.User.Id, request.PathValue("postid"))
    if err != nil {
		// Return an error if there is an issue with liking the post
        return err 
    }

    // Respond with a success message indicating the action was successful
    return writeJSON(writer, http.StatusOK, "OK")
}

// ProfilePosts retrieves and returns all posts made by the user (profile posts).
func (server *API) ProfilePosts(writer http.ResponseWriter, request *http.Request) (err error) {
    // Retrieve the session of the user making the request
    sess, err := server.Sessions.GetSession(request)
    if err != nil {
		// Return an error if session retrieval fails
        return err 
    }

    // Parse the limit and offset parameters for pagination
    limit, offset := parseRequestLimitAndOffset(request)

    // Retrieve the user's posts from the storage layer, using the provided pagination
    posts, err := server.Storage.GetUserPosts(request.Context(), sess.User.Id, limit, offset)
    if err != nil {
		// Return an error if there is an issue retrieving the posts
        return err 
    }

    // Respond with the retrieved posts data
    return writeJSON(writer, http.StatusOK, posts)
}
