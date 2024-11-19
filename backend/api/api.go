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

	gorilla "github.com/gorilla/websocket"
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

func NewAPI(addr string, dbFilePath string) (*API, error) {
	server := new(API)
	server.Server.Addr = addr

	router := http.NewServeMux()

	router.Handle("/api/register", handleFunc(server.Register))
	router.Handle("/api/login", handleFunc(server.Login))

	router.Handle("/api/user/{userid}", handleFunc(server.User))
	router.Handle("/api/user/{userid}/stats", handleFunc(server.GetUserStats))
	router.Handle("/api/user/{userid}/send-request", handleFunc(server.SendFriendRequest))
	router.Handle("/api/user/{userid}/accept-request", handleFunc(server.AcceptFriendRequest))
	router.Handle("/api/user/{userid}/chats", handleFunc(server.GetChatFrom2Userid))
	router.Handle("/api/friend-list", handleFunc(server.GetUserFriendList))
	router.Handle("/api/friend-requests", handleFunc(server.GetFriendRequests))
	// router.Handle("/api/user/{userid}/followers", handleFunc(server.GetFollowersOfUser))
	// router.Handle("/api/user/{userid}/posts", handleFunc(server.AllPostsFromOneUser))
	router.Handle("/api/profile", handleFunc(server.Profile))
	router.Handle("/api/profile/posts", handleFunc(server.ProfilePosts))
	router.Handle("/api/profile/followers", handleFunc(server.GetProfileFollowers))
	router.Handle("/api/profile/following", handleFunc(server.GetProfileFollowing))

	router.Handle("/api/groups", handleFunc(server.GetGroups))
	router.Handle("/api/group/{groupname}", handleFunc(server.Group))
	router.Handle("/api/group/{groupid}/posts", handleFunc(server.GetGroupPosts))
	router.Handle("/api/group/{groupid}/events", handleFunc(server.Events))
	router.Handle("/api/group/{groupid}/events/{eventid}", handleFunc(server.RegisterUserToEvent))

	router.Handle("/api/create/group", handleFunc(server.CreateGroup))

	// router.Handle("/api/group/{groupname}/chats", handleFunc(server.GetChatFromGroup))

	router.Handle("/api/post", handleFunc(server.CreatePost))
	router.Handle("/api/post/{postid}", handleFunc(server.Post))
	router.Handle("/api/post/{postid}/comment", handleFunc(server.Comment))
	router.Handle("/api/post/{postid}/like", server.Protected(server.LikePost))
	// router.Handle("/api/post/{postid}/comments", handleFunc(server.GetAllCommentsFromOnePost))

	// router.Handle("/api/posts/follows/{userid}", handleFunc(server.GetAllPostsFromOneUsersFollows))
	// router.Handle("/api/posts/likes/{userid}", handleFunc(server.GetAllPostsFromOneUsersLikes))
	// router.Handle("/api/chats/{userid}", handleFunc(server.GetChatFrom2Userid))

	server.WebSocket.Upgrader = gorilla.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true // TODO: Check origin
		},
	}

	server.WebSocket.Users = make(map[string]*websocket.SocketConn)
	router.HandleFunc("/api/socket", server.Socket)
	router.Handle("/api/online", handleFunc(server.GetOnlineUsers))

	router.Handle("/api/images/", http.StripPrefix("/api/images/", http.FileServer(http.Dir("api/images"))))

	server.Server.Handler = router

	storage, err := database.NewSQLite3Store(dbFilePath)
	if err != nil {
		return nil, err
	}
	server.Storage = storage

	server.Sessions = database.NewSessionStore()
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
			log.Println(err)
			writeJSON(w, http.StatusUnauthorized,
				APIerror{
					http.StatusUnauthorized,
					"Unauthorized",
					"You are not authorized to access this ressource",
				})
			return
		}
		if err := fn(w, r); err != nil {
			log.Println(err)
			writeJSON(w, http.StatusInternalServerError,
				APIerror{
					http.StatusInternalServerError,
					"Internal Server Error",
					err.Error(),
				})
		}
	})
}

func MultiPartFiles(request *http.Request) (filepaths []string, err error) {
	multipartImages := request.MultipartForm.File["images"]

	filepaths = make([]string, len(multipartImages))

	for idx, fileHeader := range multipartImages {
		file, err := fileHeader.Open()
		if err != nil {
			return nil, err
		}
		defer file.Close()

		temp, err := os.CreateTemp("api/images", fmt.Sprintf("*-%s", fileHeader.Filename))
		if err != nil {
			return nil, err
		}
		defer temp.Close()

		_, err = temp.ReadFrom(file)
		if err != nil {
			return nil, err
		}

		filepaths[idx] = fmt.Sprintf("/%s", temp.Name())
	}
	return filepaths, nil
}
