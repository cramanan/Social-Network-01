CREATE TABLE IF NOT EXISTS comments(
    user_id TEXT REFERENCES users(id),
    post_id TEXT REFERENCES posts(id),
    content TEXT NOT NULL,
    image_path TEXT NOT NULL,
    timestamp DATETIME NOT NULL
);
