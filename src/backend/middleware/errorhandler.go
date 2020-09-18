package middleware

/** The errorhandler middleware provides a unified way to respond to server errors via http */

import (
	"encoding/json"
	"net/http"
)

// HTTPError -- a json encoded error object for http message passing
type HTTPError struct {
	Err     error
	Message string
	Status  int
}

// HandleError handles database errors
func HandleError(w http.ResponseWriter, err *HTTPError) {

	LogError(err.Err)
	handleError(w, HTTPError{
		Message: err.Message,
		Status:  err.Status,
	})
}

// handleError handles errors via http by responding through: [w http.ResponseWriter] with a
// descriptive json encoded error message
func handleError(w http.ResponseWriter, thrownError HTTPError) {
	var statusCode int = 0
	if thrownError.Status == 0 {
		statusCode = 500
	} else {
		statusCode = thrownError.Status
	}

	w.WriteHeader(statusCode)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(thrownError)
}
