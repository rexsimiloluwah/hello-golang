package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	models "../models"
	u "../utils"
	"github.com/gorilla/mux"
)

func CreatePost(w http.ResponseWriter, req *http.Request) {
	// Extract the user ID from the parsed auth token
	user := req.Context().Value("user").(uint)
	newPost := &models.Post{}
	var response map[string]interface{}
	err := json.NewDecoder(req.Body).Decode(newPost)
	if err != nil {
		response = u.Message(false, "Invalid Request.")
		u.BuildResponse(w, 400, response)
		return
	}
	newPost.UserId = user
	response = newPost.CreatePost()
	if !response["status"].(bool) {
		u.BuildResponse(w, 400, response)
		return
	}
	u.BuildResponse(w, 201, response)
}

func FetchAllPosts(w http.ResponseWriter, req *http.Request) {
	response := models.GetPosts()
	if !response["status"].(bool) {
		u.BuildResponse(w, 404, response)
		return
	}

	u.BuildResponse(w, 200, response)
}

func FetchPostById(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	id, _ := strconv.ParseUint(params["id"], 10, 32)
	response := models.GetPostById(uint(id))
	if !response["status"].(bool) {
		u.BuildResponse(w, 404, response)
		return
	}

	u.BuildResponse(w, 200, response)
}

func UpdatePost(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	id, _ := strconv.ParseUint(params["id"], 10, 32)
	var fields, response map[string]interface{}
	err := json.NewDecoder(req.Body).Decode(&fields)
	if err != nil {
		u.BuildResponse(w, 400, u.Message(false, "Invalid Request."))
		return
	}

	response = models.UpdatePost(uint(id), fields)
	if !response["status"].(bool) {
		u.BuildResponse(w, 400, response)
		return
	}

	u.BuildResponse(w, 200, response)

}
func DeletePost(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	id, _ := strconv.ParseUint(params["id"], 10, 32)
	var response map[string]interface{}
	response = models.DeletePost(uint(id))
	if !response["status"].(bool) {
		u.BuildResponse(w, 400, response)
		return
	}

	u.BuildResponse(w, 200, response)
}
