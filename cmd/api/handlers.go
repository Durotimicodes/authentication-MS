package main

import (
	"errors"
	"fmt"
	"net/http"
)

func (app *Config) Authenticate(w http.ResponseWriter, r *http.Request) {
	//request payload
	var requestPayload struct {
		Email    string `json:"email"`
		Password string `json:"password`
	}

	//read json
	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	//if i get pass the above, then validate the user against the database
	users, err := app.Models.User.GetByEmail(requestPayload.Email)
	if err != nil {
		app.errorJSON(w, errors.New("invalid credentials "), http.StatusBadRequest)
		return
	}

	//validate password
	valid, err := users.PasswordMatches(requestPayload.Password)
	if err != nil || !valid {
		app.errorJSON(w, errors.New("invalid credentials"), http.StatusBadRequest)
		return
	}

	//paylaod of response
	payload := JsonResponse{
		Error:   false,
		Message: fmt.Sprintf("Logged in user %s", users.Email),
		Data: users,
	}

	//write the json
	app.writeJSON(w, http.StatusOK, payload)

}

