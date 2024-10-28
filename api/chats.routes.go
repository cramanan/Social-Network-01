package api

import (
	"context"
	"database/sql"
	"log"
	"net/http"

	"Social-Network-01/api/database"
)

func (server *API) Socket(writer http.ResponseWriter, request *http.Request) {
	_, err := server.Sessions.GetSession(request)
	if err != nil {
		writeJSON(writer, http.StatusUnauthorized, "Unauthorized")
		return
	}

	conn, err := server.WSUpgrader.Upgrade(writer, request, nil)
	if err != nil {
		log.Println(err)
		return
	}

	for {
		value := new(any)
		err = conn.ReadJSON(value)
		if err != nil {
			log.Println(err)
			return
		}

		err = conn.WriteJSON(value)
		if err != nil {
			log.Println(err)
			return
		}
	}
}

// Retrieve all chats beetween 2 users from the database.
//
// `server` is a pointer of the API type (see ./api/api.go). It contains a session reference.
func (server *API) GetChatFrom2Userid(writer http.ResponseWriter, request *http.Request) error {
	ctx, cancel := context.WithTimeout(request.Context(), database.TransactionTimeout)
	defer cancel()
	if request.Method == http.MethodGet {
		limit, offset := parseRequestLimitAndOffset(request)
		sessionUser, err := server.Sessions.GetSession(request)
		if err != nil {
			return writeJSON(writer, http.StatusNotFound,
				APIerror{
					http.StatusNotFound,
					"Not found",
					"User does not exist",
				},
			)
		}

		chats, err := server.Storage.GetChats(ctx, request.PathValue("userid"), sessionUser.User.Id, limit, offset)
		if err != nil {
			if err == sql.ErrNoRows {
				return writeJSON(writer, http.StatusNotFound,
					APIerror{
						http.StatusNotFound,
						"Not found",
						"Chat not found",
					},
				)
			}
			return err
		}

		return writeJSON(writer, http.StatusOK, chats)
	}

	return writeJSON(writer, http.StatusMethodNotAllowed,
		APIerror{
			http.StatusMethodNotAllowed,
			"Method Not Allowed",
			"Method not Allowed",
		})
}

// Retrieve all chats beetween 2 users from the database using their userIds.
//
// `server` is a pointer of the API type (see ./api/api.go). It contains a session reference.
func (server *API) GetChatFromGroup(writer http.ResponseWriter, request *http.Request) error {
	ctx, cancel := context.WithTimeout(request.Context(), database.TransactionTimeout)
	defer cancel()

	groupname := request.PathValue("groupname")

	if request.Method != http.MethodGet {
		return writeJSON(writer, http.StatusMethodNotAllowed,
			APIerror{
				http.StatusMethodNotAllowed,
				"Method Not Allowed",
				"Method not Allowed",
			})
	}

	sessionUser, err := server.Sessions.GetSession(request)
	if err != nil {
		return writeJSON(writer, http.StatusNotFound,
			APIerror{
				http.StatusNotFound,
				"Not found",
				"User does not exist",
			},
		)
	}
	limit, offset := parseRequestLimitAndOffset(request)

	chats, err := server.Storage.GetChats(ctx, groupname, sessionUser.User.Id, limit, offset)
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

	return writeJSON(writer, http.StatusOK, chats)
}
