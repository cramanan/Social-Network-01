package api

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"time"

	"Social-Network-01/api/types"
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
	conn, err := server.WebSocket.Upgrade(writer, request, nil)
	if err != nil {
		log.Println(err)
		return
	}

	// Range over every online users
	for _, userConn := range server.WebSocket.Users.Entries() {
		// instantiate a socket message
		ping := types.SocketMessage[types.OnlineUser]{
			Type: "ping",
			Data: types.OnlineUser{User: &types.User{Id: sess.User.Id}, Online: true},
		}

		userConn.WriteJSON(ping) // Write to online conn
	}

	server.WebSocket.Users.Add(sess.User.Id, conn) // Safely set user as online

	// Program deferred behaviour
	defer func() {
		server.WebSocket.Users.Remove(sess.User.Id) // Safely set user as offline

		// instantiate a socket message
		for _, userConn := range server.WebSocket.Users.Entries() {
			ping := types.SocketMessage[types.OnlineUser]{
				Type: "ping",
				Data: types.OnlineUser{User: &types.User{Id: sess.User.Id}, Online: false},
			}

			userConn.WriteJSON(ping) // Write to online conn
		}
	}()

	// Keep connection alive
	for {
		// Read user message and unmarshal it into rawChat
		var rawchat types.SocketMessage[types.ClientChat]
		err = conn.ReadJSON(&rawchat)
		if err != nil {
			return
		}

		switch rawchat.Type {
		case "message":

			// Retrieve contacted user
			_, err = server.Storage.GetUser(request.Context(), rawchat.Data.RecipientId)
			if err != nil {
				continue
			}

			// Prepare socket message
			chat := types.SocketMessage[types.ServerChat]{
				Type: "message",
				Data: types.ServerChat{
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
			if recipient, ok := server.WebSocket.Users.Lookup(chat.Data.RecipientId); ok {
				recipient.WriteJSON(chat)
			}
		}
	}
}

// Retrieve all chats beetween 2 users from the database.
//
// `server` is a pointer of the API type (see ./api/api.go). It contains a session reference.
func (server *API) GetChatFrom2Userid(writer http.ResponseWriter, request *http.Request) error {
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

	chats, err := server.Storage.GetChats(request.Context(), request.PathValue("userid"), sessionUser.User.Id, limit, offset)
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

func (server *API) GetChatFromGroup(writer http.ResponseWriter, request *http.Request) error {
	limit, offset := parseRequestLimitAndOffset(request)

	switch request.Method {
	case http.MethodGet:
		chats, err := server.Storage.GetChatsFromGroup(context.TODO(), request.PathValue("groupid"), limit, offset)
		if err != nil {
			return err
		}
		return writeJSON(writer, http.StatusOK, chats)

	default:
		return HTTPerror(http.StatusMethodNotAllowed)
	}
}

func (server *API) JoinGroupChat(writer http.ResponseWriter, request *http.Request) {
	sess, err := server.Sessions.GetSession(request)
	if err != nil {
		return
	}

	conn, err := server.WebSocket.Upgrade(writer, request, nil)
	if err != nil {
		log.Println(err)
		return
	}

	groupid := request.PathValue("groupid")
	chatroom, ok := server.WebSocket.Chatrooms.Lookup(groupid)
	if !ok {
		chatroom = server.WebSocket.Chatrooms.Add(groupid, websocket.NewChatRoom())
		log.Printf("Chatroom for %s created", groupid)
	}

	chatroom.Add(sess.User.Id, conn)
	defer chatroom.Remove(sess.User.Id)

	var message types.ServerChat
	var userConn *websocket.MxConn

	for {
		err = conn.ReadJSON(&message)
		if err != nil || message.Content == "" {
			return
		}

		message.SenderId = sess.User.Id
		message.RecipientId = groupid // use groupid for storage
		message.Timestamp = time.Now()
		server.Storage.StoreGroupChat(request.Context(), message)

		for message.RecipientId, userConn = range chatroom.Entries() { // then use user id for broadcast
			if message.RecipientId != message.SenderId {
				userConn.WriteJSON(message)
			}
		}
		message.Content = ""
	}
}
