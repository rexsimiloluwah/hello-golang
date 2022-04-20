package postgres

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/rexsimiloluwah/hello-golang/src/projects/webserver/entity"
)

func NewQuoteStore(db *sqlx.DB) *QuoteStore {
	return &QuoteStore{
		DB: db,
	}
}

type QuoteStore struct {
	*sqlx.DB
}

func (s *QuoteStore) GetQuote(id int) (entity.Quote, error) {
	var quote entity.Quote
	if err := s.Get(&quote, `SELECT * FROM quotes WHERE id = $1`, id); err != nil {
		return entity.Quote{}, fmt.Errorf("error getting quote by id: %v", err)
	}
	return quote, nil
}

func (s *QuoteStore) GetQuotesByTweet(tweetId int) ([]entity.Quote, error) {
	var quotes []entity.Quote
	if err := s.Select(&quotes, `SELECT * FROM quotes WHERE tweet_id = &1`, tweetId); err != nil {
		return []entity.Quote{}, fmt.Errorf("error getting quotes by tweet: %v", err)
	}
	return quotes, nil
}

func (s *QuoteStore) CreateQuote(quote *entity.Quote) error {
	if err := s.Get(quote, `INSERT INTO quotes (content, tweet_id) VALUES ($1,$2) RETURNING *`,
		quote.TweetID,
		quote.Content,
	); err != nil {
		return fmt.Errorf("error creating new quote: %v", err)
	}
	return nil
}

func (s *QuoteStore) UpdateQuote(quote *entity.Quote) error {
	if err := s.Get(quote, `UPDATE quotes SET tweet_id = $1, content = $2 WHERE id = $3 RETURNING *`,
		quote.TweetID,
		quote.Content,
		quote.ID,
	); err != nil {
		return fmt.Errorf("error updating quote: %v", err)
	}
	return nil
}

func (s *QuoteStore) DeleteQuote(id int) error {
	if _, err := s.Exec(`DELETE FROM quotes WHERE id = $1`, id); err != nil {
		return fmt.Errorf("error deleting quote: %v", err)
	}
	return nil
}
