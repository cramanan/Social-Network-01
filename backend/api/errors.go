package api

import "net/http"

// APIerror represents an error structure that contains status, status text, and a message
// which can be used for consistent API error responses.
type APIerror struct {
	// The HTTP status code for the error.
	Status     int    `json:"status"`     

	// The textual representation of the status code (e.g., "Not Found").
	StatusText string `json:"statusText"` 

	// The custom error message describing the error.
	Message    string `json:"message"`    
}

// Error implements the error interface for the APIerror type.
// This allows APIerror to be used as a standard error type in Go.
func (err APIerror) Error() string {
	// Return the status text (e.g., "Not Found") as the error message.
	return err.StatusText 
}

// HTTPerror is a helper function to create an APIerror instance.
// It takes an HTTP status code and an optional message to generate a structured error response.
func HTTPerror(status int, message ...string) (err APIerror) {
	// If no message is provided, use the default message for the given HTTP status code.
	if message == nil {
		// Default to the status text for the error (e.g., "OK" or "Not Found").
		message = append(message, http.StatusText(status)) 
	}

	// Return an APIerror with the specified status, status text, and custom or default message.
	return APIerror{
		// HTTP status code.
		Status:     status,    

		// HTTP status text (e.g., "Not Found").
		StatusText: http.StatusText(status), 
		
		// The custom or default error message.
		Message:    message[0],                   
	}
}

