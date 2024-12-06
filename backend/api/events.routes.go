package api

import (
	"Social-Network-01/api/types"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// RegisterUserToEvent allows a user to register for an event.
// The function checks the HTTP method and registers the user to the event based on the event ID in the request.
func (server *API) RegisterUserToEvent(writer http.ResponseWriter, request *http.Request) (err error) {
	// Retrieve the session of the user making the request.
	sess, err := server.Sessions.GetSession(request)
	if err != nil {
		// If session retrieval fails, return the error.
		return err
	}

	// Switch based on the HTTP method for the request.
	switch request.Method {
	case http.MethodPost:
		// Call the storage layer to register the user for the event.
		// The event ID is obtained from the URL path (request.PathValue("eventid")).
		return server.Storage.RegisterUserToEvent(context.TODO(), sess.User.Id, request.PathValue("eventid"))
	default:
		// If the method is not POST, return an HTTP error (Method Not Allowed).
		return HTTPerror(http.StatusMethodNotAllowed)
	}
}

// Events handles the creation and retrieval of events in a group.
// It supports POST (to create an event) and GET (to retrieve events).
func (server *API) Events(writer http.ResponseWriter, request *http.Request) (err error) {
	sess, err := server.Sessions.GetSession(request)
	if err != nil {
		return err
	}

	// Retrieve the group ID from the URL path.
	groupId := request.PathValue("groupid")

	allowed := server.Storage.UserInGroup(request.Context(), groupId, sess.User.Id)
	if !allowed {
		return writeJSON(writer, http.StatusUnauthorized, HTTPerror(http.StatusUnauthorized))
	}

	// Switch based on the HTTP method for the request.
	switch request.Method {
	case http.MethodPost:
		// If the method is POST, decode the event data from the request body into an Event object.
		var event types.Event
		err = json.NewDecoder(request.Body).Decode(&event)
		if err != nil {
			// If decoding fails, return the error.
			return err
		}

		// Set the group ID for the event.
		event.GroupId = groupId

		// Validate the event. If it's invalid, return a bad request error.
		if !event.Valid() {
			return writeJSON(writer,
				http.StatusBadRequest,
				HTTPerror(http.StatusBadRequest, "Invalid Request"))
		}

		// Store the event in the database.
		var created *types.Event
		created, err = server.Storage.CreateEvent(context.TODO(), event)
		if err != nil {
			return err
		}

		if event.Going {
			server.Storage.RegisterUserToEvent(request.Context(), sess.User.Id, created.Id)
		}

		var members []types.User
		members, err = server.Storage.GetGroupMembers(request.Context(), groupId, -1, 0)
		if err != nil {
			return err
		}

		for _, member := range members {
			if conn, ok := server.WebSocket.Users.Lookup(member.Id); ok {
				conn.WriteJSON(types.SocketMessage[string]{
					Type: "event",
					Data: fmt.Sprintf("%s created a new event: %s", sess.User.Nickname, event.Title),
				})
			}
		}

		return

	case http.MethodGet:
		// If the method is GET, retrieve the session for the current user.
		sess, err := server.Sessions.GetSession(request)
		if err != nil {
			// If session retrieval fails, return the error.
			return err
		}

		// Parse the limit and offset parameters from the request.
		limit, offset := parseRequestLimitAndOffset(request)

		// Retrieve the events for the user in the specified group with pagination (limit and offset).
		events, err := server.Storage.GetEvents(context.TODO(), sess.User.Id, groupId, limit, offset)
		if err != nil {
			// If event retrieval fails, return the error.
			return err
		}

		// Write the events to the response in JSON format with HTTP status OK.
		return writeJSON(writer, http.StatusOK, events)

	default:
		// If the method is not POST or GET, return an HTTP error (Method Not Allowed).
		return writeJSON(writer, http.StatusMethodNotAllowed, "Method not allowed")
	}
}
