CREATE TABLE IF NOT EXISTS follow_records(
    user_id TEXT NOT NULL REFERENCES users(id),
    follower_id TEXT NOT NULL REFERENCES users(id),
    accepted BOOLEAN,

    UNIQUE (user_id, follower_id)
);