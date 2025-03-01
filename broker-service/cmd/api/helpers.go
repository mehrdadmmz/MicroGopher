package main

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type jsonResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

// readJson reads a JSON object from the request body and stores it in the data parameter passed in
func (app *Config) readJSON(w http.ResponseWriter, r *http.Request, data any) error {
	maxBytes := 1048576 // one megabyte is the maximum size of the request body

	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes)) // limit the size of the request body

	dec := json.NewDecoder(r.Body) // create a new decoder for the request body
	err := dec.Decode(data)        // decode the request body into the data parameter
	if err != nil {
		return err
	}

	err = dec.Decode(&struct{}{})
	if err != io.EOF {
		return errors.New("body must hvae only a single JSON value") // if there is more than one JSON value in the request body, return an error
	}

	return nil
}

// writeJSON writes a JSON object to the response writer with the specified status code and headers
func (app *Config) writeJSON(w http.ResponseWriter, status int, data any, headers ...http.Header) error {
	out, err := json.Marshal(data) // marshal the data parameter into a JSON object
	if err != nil {
		return err
	}

	// if headers are provided, set them in the response writer before writing the JSON object
	if len(headers > 0) {
		for key, value := range headers[0] {
			w.Header()[key] = value
		}
	}

	w.Header().Set("Content-Type", "application/json") // set the content type header to application/json
	w.WriteHeader(status)                              // set the status code of the response writer
	_, err = w.Write(out)
	if err != nil {
		return err
	}

	return nil
}
