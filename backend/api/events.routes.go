package api

import (
	"Social-Network-01/api/types"
	"context"
	"encoding/json"
	"log"
	"net/http"
)

func (server *API) GetGroupEvents(writer http.ResponseWriter, request *http.Request) (err error) {
	limit, offset := parseRequestLimitAndOffset(request)
	events, err := server.Storage.GetEvents(context.TODO(), request.PathValue("groupid"), limit, offset)
	if err != nil {
		return err
	}
	log.Println(len(events))
	return writeJSON(writer, http.StatusOK, events)
}

func (server *API) CreateEvent(writer http.ResponseWriter, request *http.Request) (err error) {
	if request.Method != http.MethodPost {
		return writeJSON(writer, http.StatusMethodNotAllowed, "Method not allowed")
	}

	var event types.Event
	err = json.NewDecoder(request.Body).Decode(&event)
	if err != nil {
		return err
	}

	if !event.Valid() {
		return writeJSON(writer,
			http.StatusBadRequest,
			HTTPerror(http.StatusBadRequest, "Invalid Request"))
	}

	return server.Storage.CreateEvent(context.TODO(), event)
}
