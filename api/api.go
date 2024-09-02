package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"Social-Network-01/api/database"
)

type API struct {
	http.Server
	Storage  *database.SQLite3Store
	Sessions *database.SessionStore
}

type APIerror struct {
	Status  int    `json:"status"`
	Error   string `json:"error"`
	Message string `json:"message"`
}

func NewAPI(addr string) (*API, error) {
	server := new(API)
	server.Server.Addr = addr

	router := http.NewServeMux()

	router.Handle("/api/images/", http.StripPrefix("/api/images/", http.FileServer(http.Dir("api/images"))))
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "dist/index.html")
	})
	router.Handle("/assets/", http.FileServer(http.Dir("dist")))

	server.Server.Handler = router

	storage, err := database.NewSQLite3Store()
	if err != nil {
		return nil, err
	}
	server.Storage = storage

	server.Sessions = database.NewSessionStore()
	return server, nil
}

func parseRequestLimitAndOffset(request *http.Request) (limit, offset *int) {
	params := request.URL.Query()
	if params.Has("limit") {
		val, err := strconv.Atoi(params.Get("limit"))
		if err == nil {
			limit = &val
		}
	}
	if params.Has("limit") {
		val, err := strconv.Atoi(params.Get("limit"))
		if err == nil {
			offset = &val
		}
	}
	return limit, offset
}

func writeJSON(writer http.ResponseWriter, statusCode int, v any) error {
	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(statusCode)
	return json.NewEncoder(writer).Encode(v)
}

type handlerFunc func(http.ResponseWriter, *http.Request) error

func handleFunc(fn handlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := fn(w, r); err != nil {
			log.Println(err)
			writeJSON(w, http.StatusInternalServerError,
				APIerror{
					Status:  http.StatusInternalServerError,
					Error:   "Internal Server Error",
					Message: err.Error(),
				})
		}
	}
}

func (server *API) Protected(fn handlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := server.Sessions.GetSession(r)
		if err != nil {
			log.Println(err)
			writeJSON(w, http.StatusUnauthorized,
				APIerror{
					Status:  http.StatusUnauthorized,
					Error:   "Unauthorized",
					Message: "You are not authorized to access this ressource",
				})
			return
		}
		if err := fn(w, r); err != nil {
			log.Println(err)
			writeJSON(w, http.StatusInternalServerError,
				APIerror{
					Status:  http.StatusInternalServerError,
					Error:   "Internal Server Error",
					Message: err.Error(),
				})
		}
	})
}
