package mongodb

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/rexsimiloluwah/hello-golang/src/projects/golang-mongo-starter/internal/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Connection struct {
	Client *mongo.Client
	ctx    context.Context
}

func NewMongoConnection(cfg *config.Settings) Connection {
	var DB_URI string
	if cfg.Env == "development" {
		DB_URI = fmt.Sprintf("mongodb://%s:%s", cfg.DbHost, cfg.DbPort)
	}

	credentials := options.Credential{
		Username: cfg.DbUser,
		Password: cfg.DbPass,
	}
	fmt.Println(DB_URI)
	clientOpts := options.Client().ApplyURI(DB_URI).SetAuth(credentials)

	// context: cancel the connect operation if it is taking too much time
	ctx, cancel := context.WithTimeout(
		context.Background(),
		10*time.Second,
	)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOpts)

	// ping the database to ensure a valid connection
	err = client.Ping(ctx, readpref.Primary())

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to database.")

	return Connection{
		Client: client,
		ctx:    ctx,
	}
}

// disconnect from the database
func (c Connection) Disconnect() {
	c.Client.Disconnect(c.ctx)
}
