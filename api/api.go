package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"Social-Network-01/api/database"
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
}

type APIerror struct {
	Status   int    `json:"status"`
	ErrorMsg string `json:"error"`
	Message  string `json:"message"`
}

func (err APIerror) Error() string {
	return err.ErrorMsg
}

func NewAPI(addr string, dbFilePath string) (*API, error) {
	server := new(API)
	server.Server.Addr = addr

	router := http.NewServeMux()

	router.HandleFunc("/api/register", handleFunc(server.Register))
	router.HandleFunc("/api/login", handleFunc(server.Login))

	router.HandleFunc("/api/user/{userid}", handleFunc(server.User))
	router.HandleFunc("/api/user/{userid}/follow", handleFunc(server.FollowUser))
	router.HandleFunc("/api/user/{userid}/followers", handleFunc(server.GetFollowersOfUser))
	router.HandleFunc("/api/user/{userid}/posts", handleFunc(server.AllPostsFromOneUser))
	router.HandleFunc("/api/group/{groupid}/posts", handleFunc(server.GetAllPostsFromOneGroup))
	router.HandleFunc("/api/group/{groupid}/chats", handleFunc(server.GetChatFromGroup))
	router.HandleFunc("/api/group/{groupid}", handleFunc(server.Group))

	// router.HandleFunc("/api/posts/follows/{userid}", handleFunc(server.GetAllPostsFromOneUsersFollows))
	// router.HandleFunc("/api/posts/likes/{userid}", handleFunc(server.GetAllPostsFromOneUsersLikes))
	// router.HandleFunc("/api/post/{postid}/comments", handleFunc(server.GetAllCommentsFromOnePost))
	// router.HandleFunc("/api/chats/{userid}", handleFunc(server.GetChatFrom2Userid))

	router.Handle("/images/", http.FileServer(http.Dir("api/images")))

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
	writer.Header().Add("Content-Type", "application/json")
	writer.Header().Add("Access-Control-Allow-Origin", "*")
	writer.WriteHeader(statusCode)
	return json.NewEncoder(writer).Encode(v)
}

// api.HandlerFunc has the same signature as an http.HandlerFunc except if any error is returned,
//
//	it will use writeJSON to encode the error.
type handlerFunc func(http.ResponseWriter, *http.Request) error

// api.HandleFunc is the middleware that will change an api.HandlerFunc into a HTTP.HandlerFunc.
func handleFunc(fn handlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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
