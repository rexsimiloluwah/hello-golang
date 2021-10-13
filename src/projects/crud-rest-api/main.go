// Main app entry point
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"./controllers"
	"./middlewares"
	"./models"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func HelloWorld(w http.ResponseWriter, req *http.Request) {
	user := &models.User{}
	fmt.Println(user)
	fmt.Println("Hello world")
	w.WriteHeader(200)
	w.Write([]byte("Hello World."))
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err.Error())
	}

	var PORT string
	PORT = os.Getenv("PORT")
	if PORT == "" {
		PORT = "5050" // Default port
	}

	fmt.Println("PORT:- ", PORT)
	// Initialize Mux router
	router := mux.NewRouter()
	router.HandleFunc("/", HelloWorld)
	router.HandleFunc("/api/v1/auth/register", controllers.Register).Methods("POST")
	router.HandleFunc("/api/v1/auth/login", controllers.Login).Methods("POST")
	router.HandleFunc("/api/v1/post", controllers.CreatePost).Methods("POST")
	router.HandleFunc("/api/v1/posts", controllers.FetchAllPosts).Methods("GET")
	router.HandleFunc("/api/v1/posts/{id}", controllers.FetchPostById).Methods("GET")
	router.HandleFunc("/api/v1/posts/{id}", controllers.UpdatePost).Methods("PUT")
	router.HandleFunc("/api/v1/posts/{id}", controllers.DeletePost).Methods("DELETE")
	router.HandleFunc("/api/v1/user", controllers.FetchUser).Methods("GET")
	router.HandleFunc("/api/v1/user/posts", controllers.FetchUserPosts).Methods("GET")

	router.Use(middlewares.AuthRequired)
	log.Fatal(http.ListenAndServe(":"+PORT, router))
}
