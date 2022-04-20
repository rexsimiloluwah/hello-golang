package postgres

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func NewStore(dbUri string) (*Store, error) {
	db, err := sqlx.Connect("postgres", dbUri)
	if err != nil {
		return nil, fmt.Errorf("error opening database: %v", err)
	}

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("error pinging/connecting to database: %v", err)
	}

	return &Store{
		TweetStore:   NewTweetStore(db),
		CommentStore: NewCommentStore(db),
		QuoteStore:   NewQuoteStore(db),
	}, nil
}

type Store struct {
	*TweetStore
	*CommentStore
	*QuoteStore
}
