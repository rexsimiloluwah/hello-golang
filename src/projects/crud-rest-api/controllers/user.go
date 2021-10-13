package controllers

import (
	"net/http"

	models "../models"
	u "../utils"
)

func FetchUser(w http.ResponseWriter, req *http.Request) {
	user := req.Context().Value("user").(uint)
	response := models.GetUser(user)
	if !response["status"].(bool) {
		u.BuildResponse(w, 404, response)
		return
	}

	u.BuildResponse(w, 200, response)
}

func FetchUserPosts(w http.ResponseWriter, req *http.Request) {
	// Extract user id from request context
	user := req.Context().Value("user").(uint)
	response := models.GetUserPosts(user)
	if !response["status"].(bool) {
		u.BuildResponse(w, 404, response)
		return
	}

	u.BuildResponse(w, 200, response)
}
