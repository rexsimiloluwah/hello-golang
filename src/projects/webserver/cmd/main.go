package main

import (
	"log"
	"net/http"

	"github.com/rexsimiloluwah/hello-golang/src/projects/webserver/postgres"
	"github.com/rexsimiloluwah/hello-golang/src/projects/webserver/web"
)

func main() {
	store, err := postgres.NewStore("postgres://postgres:secret@localhost/postgres?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	h := web.NewHandler(store)
	http.ListenAndServe(":8050", h)
}
