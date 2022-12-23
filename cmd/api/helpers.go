package main

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

// json format// type declarator
type JsonResponse struct {
	Error   bool        `json:"error"`
	Message string      `json:"message"`
	Date    interface{} `json:"data"`
}

// function readJSON
func (app *Config) readJSON(w http.ResponseWriter, r *http.Request, data interface{}) error {

	//specify the max bytes to read
	maxBytes := 1048576

	//declare a request body
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes)) // prevent clients from sending larger file and wasting server resources
	//decode the json
	dec := json.NewDecoder(r.Body)
	//decode it into data
	err := dec.Decode(data)
	if err != nil {
		return err
	}

	err = dec.Decode(&struct{}{})
	if err != io.EOF {
		return errors.New("body must have only a single JSON value")
	}

	return nil

}

// function write json
func (app *Config) writeJSON(w http.ResponseWriter, status int, data interface{}, headers ...http.Header) error {
	//marshal returns a slice of bytes
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
	w.WriteHeader(status)
	_, err = w.Write(out)
	if err != nil {
		return err
	}

	return nil

}

//function error json
func (app *Config) errorJSON(w http.ResponseWriter,err error, status ...int) error {
	statuCode := http.StatusBadRequest

	if len(status) > 0 {
		statuCode = status[0]
	}

	var payload JsonResponse
	payload.Error = true
	payload.Message = err.Error()

	return app.writeJSON(w,statuCode, payload)
}
