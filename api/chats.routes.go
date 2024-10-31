package api

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"reflect"

	"Social-Network-01/api/database"
	"Social-Network-01/api/models"
)

func GenericUnmarshal[T any](raw json.RawMessage) (value T, err error) {
	err = json.Unmarshal(raw, &value)
	if err != nil {
		log.Println(err, reflect.TypeFor[T]())
		return
	}
	return
}

func (server *API) Socket(writer http.ResponseWriter, request *http.Request) {
	sess, err := server.Sessions.GetSession(request)
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
		var raw models.RawMessage
		err = conn.ReadJSON(&raw)
		if err != nil {
			log.Println(err)
			return
		}

		switch raw.Type {
		case "message":
			rawchat, err := GenericUnmarshal[models.RawChat](raw.Data)
			if err != nil {
				log.Println(err)
				break
			}

			_, err = server.Storage.GetUser(request.Context(), rawchat.RecipientId)
			if err != nil {
				log.Println(err)
				break
			}

			chat := models.Chat{
				SenderId:    sess.User.Id,
				RecipientId: rawchat.RecipientId,
				Content:     rawchat.Content,
			}

			log.Println(chat)

			err = server.Storage.StoreChat(request.Context(), chat)
			if err != nil {
				log.Println(err)
			}

		case "ping":
			err = conn.WriteJSON("PONG")
			if err != nil {
				log.Println(err)
			}
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
