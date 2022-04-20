CREATE TABLE tweets(
    id SERIAL PRIMARY KEY, 
    content TEXT NOT NULL,
    likes_count INT DEFAULT 0, 
    retweets_count INT DEFAULT 0
);

CREATE TABLE comments(
    id SERIAL PRIMARY KEY, 
    content TEXT NOT NULL, 
    likes_count INT DEFAULT 0, 
    retweets_count INT DEFAULT 0, 
    tweet_id INT NOT NULL REFERENCES tweets (id) ON DELETE CASCADE
);

CREATE TABLE quotes(
    id SERIAL PRIMARY KEY, 
    content TEXT NOT NULL, 
    likes_count INT DEFAULT 0, 
    retweets_count INT DEFAULT 0, 
    tweet_id INT NOT NULL REFERENCES tweets (id) ON DELETE CASCADE
);