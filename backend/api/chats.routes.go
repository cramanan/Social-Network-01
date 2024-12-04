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
	for _, userConn := range server.WebSocket.Users.Range {
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
		for _, userConn := range server.WebSocket.Users.Range {
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
		return writeJSON(writer, http.StatusMethodNotAllowed, HTTPerror(http.StatusMethodNotAllowed))
	}

	sessionUser, err := server.Sessions.GetSession(request)
	if err != nil {
		return writeJSON(writer, http.StatusNotFound, HTTPerror(http.StatusMethodNotAllowed, "User does not exist"))
	}

	limit, offset := parseRequestLimitAndOffset(request)

	chats, err := server.Storage.GetChats(request.Context(), request.PathValue("userid"), sessionUser.User.Id, limit, offset)
	if err == sql.ErrNoRows {
		return writeJSON(writer, http.StatusNotFound, HTTPerror(http.StatusNotFound, "Chat not found"))
	}
	if err != nil {
		return err
	}

	return writeJSON(writer, http.StatusOK, chats)

}

// GetChatFromGroup handles the HTTP request for fetching chats from a group.
// It retrieves the chats from the group based on the group ID and paginates the result using limit and offset.
func (server *API) GetChatFromGroup(writer http.ResponseWriter, request *http.Request) error {
	// Parse the limit and offset values from the request to handle pagination.
	limit, offset := parseRequestLimitAndOffset(request)

	// Handle the GET request to fetch group chats.
	switch request.Method {
	case http.MethodGet:
		// Fetch the chats from the group using the group ID, limit, and offset.
		chats, err := server.Storage.GetChatsFromGroup(context.TODO(), request.PathValue("groupid"), limit, offset)
		if err != nil {
			return err // Return an error if fetching the chats fails.
		}
		// Write the chats as a JSON response with HTTP status OK.
		return writeJSON(writer, http.StatusOK, chats)

	// Handle unsupported HTTP methods.
	default:
		// Return a Method Not Allowed error if the HTTP method is not GET.
		return HTTPerror(http.StatusMethodNotAllowed)
	}
}

// JoinGroupChat handles the WebSocket connection for joining a group chat.
// It upgrades the HTTP connection to a WebSocket connection and manages the chat session.
func (server *API) JoinGroupChat(writer http.ResponseWriter, request *http.Request) {
	// Get the current session from the request to identify the user.
	sess, err := server.Sessions.GetSession(request)
	if err != nil {
		// Return if there is an error fetching the session (user not logged in or session expired).
		return
	}

	// Upgrade the HTTP request to a WebSocket connection.
	conn, err := server.WebSocket.Upgrade(writer, request, nil)
	if err != nil {
		// Log the error if upgrading to WebSocket fails.
		log.Println(err)
		return
	}

	// Retrieve the group ID from the URL path.
	groupid := request.PathValue("groupid")

	// Check if a chatroom for the group already exists in the WebSocket connections.
	chatroom, ok := server.WebSocket.Chatrooms.Lookup(groupid)
	if !ok {
		// If no chatroom exists for the group, create a new chatroom.
		chatroom = server.WebSocket.Chatrooms.Add(groupid, websocket.NewChatRoom())
		log.Printf("Chatroom for %s created", groupid) // Log the creation of a new chatroom.
	}

	// Add the current user connection to the chatroom.
	chatroom.Add(sess.User.Id, conn)
	// Ensure the user's connection is removed when the function exits.
	defer chatroom.Remove(sess.User.Id)

	// Define a variable to hold the incoming chat message.
	var message types.ServerChat
	var userConn *websocket.MxConn

	// Continuously read messages from the WebSocket connection.
	for {
		// Read the JSON message from the connection.
		err = conn.ReadJSON(&message)
		if err != nil || message.Content == "" {
			// Exit if there's an error reading the message or if the message is empty.
			return
		}

		// Set the sender ID and recipient ID for the message.
		message.SenderId = sess.User.Id
		// Use the group ID as the recipient for storage.
		message.RecipientId = groupid
		// Set the current timestamp for the message.
		message.Timestamp = time.Now()

		// Store the message in the group chat in the database.
		server.Storage.StoreGroupChat(request.Context(), message)

		// Broadcast the message to all users in the chatroom, except the sender.
		for message.RecipientId, userConn = range chatroom.Range {
			if message.RecipientId != message.SenderId {
				// Send the message to each user in the chatroom.
				userConn.WriteJSON(message)
			}
		}

		// Clear the content of the message after broadcasting it.
		message.Content = ""
	}
}
