package api

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"time"

	"Social-Network-01/api/database"
	"Social-Network-01/api/models"
)

func (server *API) Socket(writer http.ResponseWriter, request *http.Request) {
	// Retrieve session
	sess, err := server.Sessions.GetSession(request)
	if err != nil {
		writeJSON(writer, http.StatusUnauthorized, "Unauthorized")
		return
	}

	// Switch to WebSocket protocol
	conn, err := server.WSUpgrader.Upgrade(writer, request, nil)
	if err != nil {
		log.Println(err)
		return
	}

	// Range over every online users
	for _, userConn := range server.users {

		// instantiate a socket message
		ping := models.SocketMessage[models.OnlineUser]{
			Type: "ping",
			Data: models.OnlineUser{User: &models.User{Id: sess.User.Id}, Online: true},
		}

		userConn.WriteJSON(ping) // Write to online conn
	}

	server.Lock()                     //
	server.users[sess.User.Id] = conn // Safely set user as online
	server.Unlock()                   //

	// Program deferred behaviour
	defer func() {
		server.Lock()                      //
		delete(server.users, sess.User.Id) // Safely set user as offline
		server.Unlock()                    //

		// instantiate a socket message
		for _, userConn := range server.users {
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
		if recipient, ok := server.users[chat.Data.RecipientId]; ok {
			recipient.WriteJSON(chat)
		}

		// Send message to connected user
		conn.WriteJSON(chat)
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
