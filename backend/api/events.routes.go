package api

import (
	"Social-Network-01/api/types"
	"context"
	"encoding/json"
	"net/http"
)

func (server *API) Events(writer http.ResponseWriter, request *http.Request) (err error) {
	switch request.Method {
	case http.MethodPost:
		var event types.Event
		err = json.NewDecoder(request.Body).Decode(&event)
		if err != nil {
			return err
		}

		event.GroupId = request.PathValue("groupid")

		if !event.Valid() {
			return writeJSON(writer,
				http.StatusBadRequest,
				HTTPerror(http.StatusBadRequest, "Invalid Request"))
		}

		return server.Storage.CreateEvent(context.TODO(), event)

	case http.MethodGet:
		limit, offset := parseRequestLimitAndOffset(request)
		events, err := server.Storage.GetEvents(context.TODO(), request.PathValue("groupid"), limit, offset)
		if err != nil {
			return err
		}
		return writeJSON(writer, http.StatusOK, events)

	default:
		return writeJSON(writer, http.StatusMethodNotAllowed, "Method not allowed")
	}
}
