package api

import "net/http"

type APIerror struct {
	Status   int    `json:"status"`
	ErrorMsg string `json:"error"`
	Message  string `json:"message"`
}

func (err APIerror) Error() string { return err.ErrorMsg }

func HTTPerror(status int, message ...string) (err APIerror) {
	if message == nil {
		message = append(message, http.StatusText(status))
	}

	return APIerror{
		Status:   status,
		ErrorMsg: http.StatusText(status),
		Message:  message[0],
	}
}
