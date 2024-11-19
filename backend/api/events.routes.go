package api

import (
	"Social-Network-01/api/types"
	"context"
	"encoding/json"
	"net/http"
)

func (server *API) RegisterUserToEvent(writer http.ResponseWriter, request *http.Request) (err error) {
	sess, err := server.Sessions.GetSession(request)
	if err != nil {
		return err
	}

	switch request.Method {
	case http.MethodPost:
		return server.Storage.RegisterUserToEvent(context.TODO(), sess.User.Id, request.PathValue("eventid"))
	default:
		return HTTPerror(http.StatusMethodNotAllowed)
	}
}

func (server *API) Events(writer http.ResponseWriter, request *http.Request) (err error) {
	groupId := request.PathValue("groupid")
	switch request.Method {
	case http.MethodPost:
		var event types.Event
		err = json.NewDecoder(request.Body).Decode(&event)
		if err != nil {
			return err
		}

		event.GroupId = groupId

		if !event.Valid() {
			return writeJSON(writer,
				http.StatusBadRequest,
				HTTPerror(http.StatusBadRequest, "Invalid Request"))
		}

		return server.Storage.CreateEvent(context.TODO(), event)

	case http.MethodGet:
		sess, err := server.Sessions.GetSession(request)
		if err != nil {
			return err
		}

		limit, offset := parseRequestLimitAndOffset(request)
		events, err := server.Storage.GetEvents(context.TODO(), sess.User.Id, groupId, limit, offset)
		if err != nil {
			return err
		}
		return writeJSON(writer, http.StatusOK, events)

	default:
		return writeJSON(writer, http.StatusMethodNotAllowed, "Method not allowed")
	}
}
