package main

import (
	"github.com/rexsimiloluwah/hello-golang/src/projects/golang-mongo-starter/internal/api"
	"github.com/rexsimiloluwah/hello-golang/src/projects/golang-mongo-starter/internal/config"
	"github.com/rexsimiloluwah/hello-golang/src/projects/golang-mongo-starter/pkg/mongodb"
)

func main() {
	cfg := config.New()
	db := mongodb.NewMongoConnection(cfg)
	defer db.Disconnect()
	api := api.New(cfg, db.Client)
	api.StartServer()
}
