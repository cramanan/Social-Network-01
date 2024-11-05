CREATE TABLE IF NOT EXISTS follow_records(
    user_id TEXT NOT NULL REFERENCES users(id),
    follower_id TEXT NOT NULL REFERENCES users(id),
    accepted BOOLEAN
);

CREATE TABLE IF NOT EXISTS likes_records(
    user_id TEXT REFERENCES users(id),
    post_id TEXT REFERENCES posts(id)
);
