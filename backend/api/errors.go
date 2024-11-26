package api

import "net/http"

type APIerror struct {
	Status     int    `json:"status"`
	StatusText string `json:"statusText"`
	Message    string `json:"message"`
}

func (err APIerror) Error() string { return err.StatusText }

func HTTPerror(status int, message ...string) (err APIerror) {
	if message == nil {
		message = append(message, http.StatusText(status))
	}

	return APIerror{
		Status:     status,
		StatusText: http.StatusText(status),
		Message:    message[0],
	}
}
