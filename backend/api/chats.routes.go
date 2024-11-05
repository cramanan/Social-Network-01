package api

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"sync"
	"time"

	"Social-Network-01/api/database"
	"Social-Network-01/api/models"
	"Social-Network-01/api/websocket"
)

func (server *API) Socket(writer http.ResponseWriter, request *http.Request) {
	// Retrieve session
	sess, err := server.Sessions.GetSession(request)
	if err != nil {
		writeJSON(writer, http.StatusUnauthorized, "Unauthorized")
		return
	}

	// Switch to WebSocket protocol
	conn, err := server.WebSocket.Upgrader.Upgrade(writer, request, nil)
	if err != nil {
		log.Println(err)
		return
	}

	wconn := &websocket.SocketConn{Conn: conn, Mutex: sync.Mutex{}}

	// Range over every online users
	for _, userConn := range server.WebSocket.Users {

		// instantiate a socket message
		ping := models.SocketMessage[models.OnlineUser]{
			Type: "ping",
			Data: models.OnlineUser{User: &models.User{Id: sess.User.Id}, Online: true},
		}

		userConn.WriteJSON(ping) // Write to online conn
	}

	server.WebSocket.Add(sess.User.Id, wconn) // Safely set user as online

	// Program deferred behaviour
	defer func() {
		server.WebSocket.Remove(sess.User.Id) // Safely set user as offline

		// instantiate a socket message
		for _, userConn := range server.WebSocket.Users {
			ping := models.SocketMessage[models.OnlineUser]{
				Type: "ping",
				Data: models.OnlineUser{User: &models.User{Id: sess.User.Id}, Online: false},
			}

			userConn.WriteJSON(ping) // Write to online conn
		}
	}()

	// Keep connection alive
	for {
		// Read user message and unmarshal it into rawChat
		var rawchat models.SocketMessage[models.ClientChat]
		err = conn.ReadJSON(&rawchat)
		if err != nil {
			log.Println(err)
			return
		}

		switch rawchat.Type {
		case "message":

			// Retrieve contacted user
			_, err = server.Storage.GetUser(request.Context(), rawchat.Data.RecipientId)
			if err != nil {
				log.Println(err)
				continue
			}

			// Prepare socket message
			chat := models.SocketMessage[models.ServerChat]{
				Type: "message",
				Data: models.ServerChat{
					SenderId:    sess.User.Id,
					RecipientId: rawchat.Data.RecipientId,
					Content:     rawchat.Data.Content,
					Timestamp:   time.Now(),
				},
			}

			// Store in db
			err = server.Storage.StoreChat(request.Context(), chat.Data)
			if err != nil {
				log.Println(err)
			}

			// Check if the recipient is online
			if recipient, ok := server.WebSocket.Users[chat.Data.RecipientId]; ok {
				recipient.WriteJSON(chat)
			}

			// Send message to connected user
			wconn.WriteJSON(chat)
		}
	}
}

// Retrieve all chats beetween 2 users from the database.
//
// `server` is a pointer of the API type (see ./api/api.go). It contains a session reference.
func (server *API) GetChatFrom2Userid(writer http.ResponseWriter, request *http.Request) error {
	ctx, cancel := context.WithTimeout(request.Context(), database.TransactionTimeout)
	defer cancel()
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

	chats, err := server.Storage.GetChats(ctx, request.PathValue("userid"), sessionUser.User.Id, limit, offset)
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
