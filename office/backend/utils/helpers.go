package utils

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
)

type jsonResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func ReadJson(w http.ResponseWriter, r *http.Request, data interface{}) error {
	maxBytes := 1048576
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))
	dec := json.NewDecoder(r.Body)
	err := dec.Decode(&data)

	if err != nil {
		return err
	}

	err = dec.Decode(&struct{}{})
	if err != io.EOF {
		return errors.New("body must have a single json value")
	}

	return nil
}

func WriteJson(w http.ResponseWriter, staus int, data interface{}, headers ...http.Header) error {
	out, err := json.Marshal(data)

	if err != nil {
		return err
	}

	if len(headers) > 0 {
		for key, value := range headers[0] {
			w.Header()[key] = value
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(staus)

	_, err = w.Write(out)
	if err != nil {
		return err
	}

	return nil
}

func ResponseToError(w http.ResponseWriter, err error, status ...int) error {
	statusCode := http.StatusBadRequest
	log.Printf("AN ERROR ACCURED %v", err)
	if len(status) > 0 {
		statusCode = status[0]
	}

	var payload jsonResponse
	payload.Error = true
	payload.Message = err.Error()
	payload.Data = nil

	return WriteJson(w, statusCode, payload)
}
