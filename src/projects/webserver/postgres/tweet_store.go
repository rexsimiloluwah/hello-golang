package postgres

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/rexsimiloluwah/hello-golang/src/projects/webserver/entity"
)

func NewTweetStore(db *sqlx.DB) *TweetStore {
	return &TweetStore{
		DB: db,
	}
}

type TweetStore struct {
	*sqlx.DB
}

func (s *TweetStore) GetTweet(id int) (entity.Tweet, error) {
	var tweet entity.Tweet
	if err := s.Get(&tweet, `SELECT * FROM tweets WHERE id=$1`, id); err != nil {
		return entity.Tweet{}, fmt.Errorf("error getting tweet by id: %v", err)
	}
	return tweet, nil
}

func (s *TweetStore) GetAllTweets() ([]entity.Tweet, error) {
	var tweets []entity.Tweet
	if err := s.Select(&tweets, `SELECT * FROM tweets`); err != nil {
		return []entity.Tweet{}, fmt.Errorf("error getting all tweets: %v", err)
	}
	return tweets, nil
}

func (s *TweetStore) CreateTweet(tweet *entity.Tweet) error {
	err := s.Get(tweet, `INSERT INTO tweets (content) VALUES ($1) RETURNING *`, tweet.Content)
	if err != nil {
		return fmt.Errorf("error creating new tweet: %v", err)
	}
	return nil
}

func (s *TweetStore) UpdateTweet(tweet *entity.Tweet) error {
	err := s.Get(tweet, `UPDATE tweets SET content = $1 WHERE id = $2 RETURNING *`, tweet.Content, tweet.ID)
	if err != nil {
		return fmt.Errorf("error updating tweet with id=%d : %v", tweet.ID, err)
	}
	return nil
}

func (s *TweetStore) DeleteTweet(id int) error {
	if _, err := s.Exec(`DELETE FROM tweets WHERE id = $1`, id); err != nil {
		return fmt.Errorf("error deleting tweet with id=%d : %v", id, err)
	}
	return nil
}
