package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql" // The underscore prefix initializes the module
	"github.com/gorilla/mux"
	hashids "github.com/speps/go-hashids"
)

type Url struct {
	Id           int64  `json:"id,omitempty"`
	Slug         string `json:"slug,omitempty"`
	ShortUrl     string `json:"short_url,omitempty"`
	LongUrl      string `json:"long_url,omitempty"`
	VisitorCount int    `json:"visitor_count,omitempty"`
	CreatedAt    string `json:"created_at,omitempty"`
}

func dbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "adetoyosi"
	dbName := "urlshortener"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}

// Endpoint for redirecting to the original URL
// METHOD : GET
func Root(w http.ResponseWriter, req *http.Request) {
	db := dbConn()
	params := mux.Vars(req)
	var fetchedUrl Url

	// Check if the slug exists in the database
	query := "SELECT long_url,visitor_count FROM links WHERE slug=?"
	err := db.QueryRow(query, params["id"]).Scan(&fetchedUrl.LongUrl, &fetchedUrl.VisitorCount)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("Not found in database."))
		return
	}

	// If the slug exists in the database
	stmt, err := db.Prepare("UPDATE `links` SET `visitor_count` = `visitor_count` + 1 WHERE `slug` = ?")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = stmt.Exec(&fetchedUrl.Slug)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	handleRedirect(w, req, fetchedUrl.LongUrl)
}

func handleRedirect(w http.ResponseWriter, req *http.Request, newUrl string) {
	http.Redirect(w, req, newUrl, http.StatusMovedPermanently)
}

// Generate a slug using hashids library
func generateSlug() string {
	hd := hashids.NewData()
	h, _ := hashids.NewWithData(hd)
	now := time.Now()
	slug, _ := h.Encode([]int{int(now.Unix())})
	return slug
}

// Endpoint for shortening the long URL
// METHOD : POST
func Shorten(w http.ResponseWriter, req *http.Request) {
	db := dbConn()
	var newUrl Url
	_ = json.NewDecoder(req.Body).Decode(&newUrl)

	// Generate a slug
	newUrl.Slug = generateSlug()
	newUrl.ShortUrl = "http://localhost:5050/" + newUrl.Slug
	query := "INSERT INTO links(slug,long_url,short_url,visitor_count,created_at) VALUES (?,?,?,?,?)"
	newUrl.CreatedAt = time.Now().Format(time.RFC3339)
	newUrl.VisitorCount = 0
	result, err := db.Exec(query, newUrl.Slug, newUrl.LongUrl, newUrl.ShortUrl, newUrl.VisitorCount, newUrl.CreatedAt)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}
	newUrl.Id, err = result.LastInsertId()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(newUrl)
	w.WriteHeader(200)

	json.NewEncoder(w).Encode(newUrl)
}

func main() {
	//db := dbConn()
	router := mux.NewRouter()
	router.HandleFunc("/shorten", Shorten).Methods("POST")
	router.HandleFunc("/{id}", Root).Methods("GET")
	log.Fatalln(http.ListenAndServe(":5050", router))
}
