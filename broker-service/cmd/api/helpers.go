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

func (app *Config) readJSON(w http.ResponseWriter, r *http.Request, data interface{}) error {
	maxSize := 1048576 // 1MB
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxSize))

	dec := json.NewDecoder(r.Body)
	error := dec.Decode(&data)
	if error != nil {
		return error
	}

	error = dec.Decode(&struct{}{})
	if error != io.EOF {
		return errors.New("request must have only a single value")
	}
	return nil
}

func (app *Config) writeJSON(w http.ResponseWriter, status int, data interface{}, headers ...http.Header) error {
	out, err := json.Marshal(data)
	if err != nil {
		return err
	}

	if len(headers) > 0 {
		for k, v := range headers[0] {
			w.Header()[k] = v
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	_, err = w.Write(out)
	if err != nil {
		return err
	}
	return nil
}

func (app *Config) errorJSON(w http.ResponseWriter, err error, status ...int) error {
	statusCode := http.StatusInternalServerError

	if len(status) > 0 {
		statusCode = status[0]
	}

	payload := jsonResponse{
		Error:   true,
		Message: err.Error(),
	}

	return app.writeJSON(w, statusCode, payload)
}
