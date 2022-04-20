package entity

type Tweet struct {
	ID            int    `db:"id"`
	Content       string `db:"content"`
	LikesCount    int    `db:"likes_count"`
	RetweetsCount int    `db:"retweets_count"`
}

type Comment struct {
	ID            int    `db:"id"`
	TweetID       int    `db:"tweet_id"`
	Content       string `db:"content"`
	LikesCount    int    `db:"likes_count"`
	RetweetsCount int    `db:"retweets_count"`
}

type Quote struct {
	ID            int    `db:"id"`
	TweetID       int    `db:"tweet_id"`
	Content       string `db:"content"`
	LikesCount    int    `db:"likes_count"`
	RetweetsCount int    `db:"retweets_count"`
}

type TweetStore interface {
	GetTweet(id int) (Tweet, error)
	GetAllTweets() ([]Tweet, error)
	CreateTweet(tweet *Tweet) error
	UpdateTweet(tweet *Tweet) error
	DeleteTweet(id int) error
}

type CommentStore interface {
	GetComment(id int) (Comment, error)
	GetCommentsByTweet(tweetId int) ([]Comment, error)
	CreateComment(comment *Comment) error
	UpdateComment(comment *Comment) error
	DeleteComment(id int) error
}

type QuoteStore interface {
	GetQuote(id int) (Quote, error)
	GetQuotesByTweet(tweetId int) ([]Quote, error)
	CreateQuote(quote *Quote) error
	UpdateQuote(quote *Quote) error
	DeleteQuote(id int) error
}

type Store interface {
	TweetStore
	CommentStore
	QuoteStore
}
