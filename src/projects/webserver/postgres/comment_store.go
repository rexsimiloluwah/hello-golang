package postgres

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/rexsimiloluwah/hello-golang/src/projects/webserver/entity"
)

func NewCommentStore(db *sqlx.DB) *CommentStore {
	return &CommentStore{
		DB: db,
	}
}

type CommentStore struct {
	*sqlx.DB
}

func (s *CommentStore) GetComment(id int) (entity.Comment, error) {
	var comment entity.Comment
	if err := s.Get(&comment, `SELECT * FROM comments WHERE id = $1`, id); err != nil {
		return entity.Comment{}, fmt.Errorf("error getting comment by id: %v", err)
	}
	return comment, nil
}

func (s *CommentStore) GetCommentsByTweet(tweetId int) ([]entity.Comment, error) {
	var comments []entity.Comment
	if err := s.Select(&comments, `SELECT * FROM comments WHERE tweet_id = &1`, tweetId); err != nil {
		return []entity.Comment{}, fmt.Errorf("error getting comments by tweet: %v", err)
	}
	return comments, nil
}

func (s *CommentStore) CreateComment(comment *entity.Comment) error {
	if err := s.Get(comment, `INSERT INTO comments (content,tweet_id) VALUES ($1, $2) RETURNING *`,
		comment.TweetID,
		comment.Content,
	); err != nil {
		return fmt.Errorf("error creating new comment: %v", err)
	}
	return nil
}

func (s *CommentStore) UpdateComment(comment *entity.Comment) error {
	if err := s.Get(comment, `UPDATE comments SET tweet_id = $1, content = $2 WHERE id = $3 RETURNING *`,
		comment.TweetID,
		comment.Content,
		comment.ID,
	); err != nil {
		return fmt.Errorf("error updating comment: %v", err)
	}
	return nil
}

func (s *CommentStore) DeleteComment(id int) error {
	if _, err := s.Exec(`DELETE FROM comments WHERE id = $1`, id); err != nil {
		return fmt.Errorf("error deleting comment: %v", err)
	}
	return nil
}
