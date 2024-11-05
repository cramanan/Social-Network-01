CREATE TABLE IF NOT EXISTS follow_records(
    user_id TEXT REFERENCES users(id),
    follower_id TEXT REFERENCES users(id)
);

CREATE TABLE IF NOT EXISTS likes_records(
    user_id TEXT REFERENCES users(id),
    post_id TEXT REFERENCES posts(id)
);
