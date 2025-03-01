package main

import (
	"net/http"
)

// Broker is a handler that will be used to handle requests to the broker service
// because it is a handler, it requires a response writer and a pointer to a request
func (app *Config) Broker(w http.ResponseWriter, r *http.Request) {
	payload := jsonResponse{
		Error:   false,
		Message: "Hit the broker",
	}

	_ = app.writeJSON(w, http.StatusOK, payload)
}
