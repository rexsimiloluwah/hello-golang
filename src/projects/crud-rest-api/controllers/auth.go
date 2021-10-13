package controllers

import (
	"encoding/json"
	"net/http"

	models "../models"
	u "../utils"
)

func Register(w http.ResponseWriter, req *http.Request) {
	newUser := &models.User{}
	// Decode the request body into a struct
	err := json.NewDecoder(req.Body).Decode(newUser)
	if err != nil {
		u.BuildResponse(w, 400, u.Message(false, "Invalid Request."))
	}

	response := newUser.CreateUser()
	if response["status"].(bool) == false {
		u.BuildResponse(w, 400, response)
		return
	}

	u.BuildResponse(w, 201, response)
}

func Login(w http.ResponseWriter, req *http.Request) {
	user := &models.User{}
	err := json.NewDecoder(req.Body).Decode(user)
	if err != nil {
		u.BuildResponse(w, 400, u.Message(false, "Invalid Request."))
	}

	response := models.Login(user.Email, user.Password)
	if response["status"].(bool) == false {
		u.BuildResponse(w, 400, response)
		return
	}

	u.BuildResponse(w, 201, response)
}
