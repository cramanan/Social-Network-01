package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"Social-Network-01/api/database"
	"Social-Network-01/api/websocket"
)

// The API struct inherits from Golang's native http.Server and has built-in:
//   - SQLite3 storage,
//   - Session system
//   - and WebSocket Upgrader.
type API struct {
	http.Server

	//	Storage is the SQLite3 DB. It comes with Forum's fuction such
	//	as create / log users, posts and comments creation.
	Storage *database.SQLite3Store

	//  sessions is the Session Store. It generates, retrieve and end Sessions
	//  using HTTP request's Cookies.
	Sessions *database.SessionStore

	WebSocket websocket.WebSocket
}

// NewAPI initializes and returns a new API server instance with routes and database setup.
func NewAPI(addr string, dbFilePath string) (*API, error) {
	// Create a new API server instance.
	server := new(API)
	server.Server.Addr = addr // Set the address for the API server.

	// Create a new router to handle incoming HTTP requests.
	router := http.NewServeMux()

	// Register routes for user authentication and profile management.
	// TODO: Protect some routes with authentication and authorization logic.
	// Route for user registration.
	router.Handle("/api/register", handleFunc(server.Register))
	// Route for user login.
	router.Handle("/api/login", handleFunc(server.Login))

	// Routes for accessing and managing user information and friend requests.
	// Route to fetch user details by user ID.
	router.Handle("/api/users/{userid}", handleFunc(server.User))
	// Route for user statistics.
	router.Handle("/api/users/{userid}/stats", handleFunc(server.GetUserStats))
	// Route for sending a friend request.
	router.Handle("/api/users/{userid}/send-request", handleFunc(server.SendFollowRequest))
	// Route for accepting a friend request.
	router.Handle("/api/users/{userid}/accept-request", handleFunc(server.AcceptFollowRequest))
	// Route for declining a friend request.
	router.Handle("/api/users/{userid}/decline-request", handleFunc(server.DeclineFollowRequest))
	// Routes related to user chats and friend list management.
	// Route for fetching chats between two users.
	router.Handle("/api/users/{userid}/groups", handleFunc(server.GetUserGroups))
	// Route for fetching a user's friend list.
	router.Handle("/api/users/{userid}/chats", handleFunc(server.GetChatFrom2Userid))

	router.Handle("/api/follow-list", handleFunc(server.GetUserFollowList))

	// Profile-related routes for fetching or updating user profile and posts.
	// Route for fetching user profile.
	router.Handle("/api/profile", handleFunc(server.Profile))
	// Route for fetching posts on a user's profile.
	router.Handle("/api/profile/posts", handleFunc(server.ProfilePosts))
	// Route for fetching followers of a user.
	router.Handle("/api/profile/followers", handleFunc(server.GetProfileFollowers))
	// Route for fetching users followed by a user.
	router.Handle("/api/profile/following", handleFunc(server.GetProfileFollowing))
	router.Handle("/api/profile/groups", handleFunc(server.GetProfileGroups))

	router.Handle("/api/groups", handleFunc(server.Groups))
	router.Handle("/api/groups/{groupid}", handleFunc(server.Group))
	router.Handle("/api/groups/{groupid}/invite", handleFunc(server.InviteUserIntoGroup))
	router.Handle("/api/groups/{groupid}/accept-invite", handleFunc(server.AcceptGroupInvite))
	router.Handle("/api/groups/{groupid}/decline-invite", handleFunc(server.DeclineGroupInvite))
	router.Handle("/api/groups/{groupid}/request", handleFunc(server.RequestGroup))
	router.Handle("/api/groups/{groupid}/members", handleFunc(server.GetGroupMembers))
	router.Handle("/api/groups/{groupid}/posts", handleFunc(server.GetGroupPosts))
	router.Handle("/api/groups/{groupid}/events", handleFunc(server.Events))
	router.Handle("/api/groups/{groupid}/events/{eventid}", handleFunc(server.RegisterUserToEvent))
	router.Handle("/api/groups/{groupid}/chats", handleFunc(server.GetChatFromGroup))
	router.Handle("/api/groups/{groupid}/chatroom", http.HandlerFunc(server.JoinGroupChat))

	router.Handle("/api/posts", handleFunc(server.Posts))
	router.Handle("/api/posts/{postid}", handleFunc(server.GetPostById))
	router.Handle("/api/posts/{postid}/comments", handleFunc(server.Comment))
	router.Handle("/api/posts/{postid}/likes", server.Protected(server.LikePost))

	router.Handle("/api/inbox/group-invites", handleFunc(server.GetGroupInvites))
	router.Handle("/api/inbox/group-requests", handleFunc(server.GetGroupRequests))

	// Route for fetching pending follow requests.
	router.Handle("/api/inbox/follow-requests", handleFunc(server.GetFollowRequests))

	// Initialize the WebSocket for real-time communication.
	server.WebSocket = websocket.NewWebSocket()

	// Route for handling WebSocket connections.
	// Route for WebSocket connections.
	router.Handle("/api/socket", http.HandlerFunc(server.Socket))
	// Route for fetching online users.
	router.Handle("/api/online", handleFunc(server.GetOnlineUsers))

	// Serve image files from the "api/images" directory.
	router.Handle("/api/images/", http.StripPrefix("/api/images/", http.FileServer(http.Dir("api/images"))))

	// Set the router as the handler for the API server.
	server.Server.Handler = router

	// Set up the SQLite database connection and handle any errors.
	storage, err := database.NewSQLite3Store(dbFilePath)
	if err != nil {
		return nil, err // Return an error if the database connection fails.
	}
	server.Storage = storage // Store the database connection in the API server.

	// Initialize the session store for managing user sessions.
	server.Sessions = database.NewSessionStore()

	// Return the fully initialized API server.
	return server, nil
}

// parseRequestLimitAndOffset is used to extract the query parameters // with the name: "limit" & "offset".
func parseRequestLimitAndOffset(request *http.Request) (limit, offset int) {
	params := request.URL.Query() // parse the Query
	limit, _ = strconv.Atoi(params.Get("limit"))
	if limit == 0 {
		limit = -1 // set to -1 for SQL Query
	}
	offset, _ = strconv.Atoi(params.Get("offset"))
	return limit, offset
}

// writeJSON writes the JSON encoding of v to the http.ResponseWriter
// and sends it with the provided status code as application/json.
func writeJSON(writer http.ResponseWriter, statusCode int, v any) error {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(statusCode)
	return json.NewEncoder(writer).Encode(v)
}

// api.HandlerFunc has the same signature as an http.HandlerFunc except if any error is returned,
//
//	it will use writeJSON to encode the error.
type handlerFunc func(http.ResponseWriter, *http.Request) error

// api.Handle is the middleware that will change an api.HandlerFunc into a HTTP.HandlerFunc.
func handleFunc(fn handlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// w.Header().Set("Access-Control-Allow-Origin", "*")
		// w.Header().Set("Access-Control-Allow-Methods", "PUT, GET, POST, DELETE, PATCH, OPTIONS")
		// w.Header().Set("Access-Control-Allow-Headers", "*")
		// w.Header().Set("Access-Control-Allow-Credentials", "true")

		// if r.Method == http.MethodOptions {
		// 	w.WriteHeader(http.StatusOK)
		// 	return
		// }

		if err := fn(w, r); err != nil {
			log.Println(err)
			writeJSON(w, http.StatusInternalServerError,
				APIerror{
					http.StatusInternalServerError,
					"Internal Server Error",
					err.Error(),
				})
		}
	}
}

//	The protected method is a middleware that will authenticate the users by finding its session.
//
// if no session is found, it will writeJSON an APIerror with Unauthorized Header.Write
func (server *API) Protected(fn handlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := server.Sessions.GetSession(r)
		if err != nil {
			writeJSON(w, http.StatusUnauthorized, HTTPerror(http.StatusUnauthorized, "You are not authorized to access this ressource"))
			return
		}
		if err := fn(w, r); err != nil {
			log.Println(err)
			writeJSON(w, http.StatusInternalServerError, HTTPerror(http.StatusInternalServerError))
		}
	})
}

// MultiPartFiles processes multipart file uploads from an HTTP request.
// It retrieves files named "images" from the request, stores them as temporary files,
// and returns a list of their paths.
func MultiPartFiles(request *http.Request) (filepaths []string, err error) {
	// Retrieve the multipart files from the request's "images" field.
	multipartImages := request.MultipartForm.File["images"]

	// Initialize the filepaths slice to store the paths of the uploaded files.
	filepaths = make([]string, len(multipartImages))

	// Iterate through each file in the "images" field of the multipart form.
	for idx, fileHeader := range multipartImages {
		// Open the file for reading.
		file, err := fileHeader.Open()
		if err != nil {
			// Return error if file cannot be opened.
			return nil, err
		}
		defer file.Close() // Ensure file is closed when the function exits.

		// Create a temporary file in the "api/images" directory with the original filename.
		temp, err := os.CreateTemp("api/images", fmt.Sprintf("*-%s", fileHeader.Filename))
		if err != nil {
			// Return error if temporary file cannot be created.
			return nil, err
		}
		// Ensure temporary file is closed after use.
		defer temp.Close()

		// Copy the contents of the uploaded file into the temporary file.
		_, err = temp.ReadFrom(file)
		if err != nil {
			// Return error if there is an issue copying the file.
			return nil, err
		}

		// Store the relative path of the temporary file in the filepaths slice.
		filepaths[idx] = fmt.Sprintf("/%s", temp.Name())
	}

	// Return the filepaths of the uploaded images.
	return filepaths, nil
}
