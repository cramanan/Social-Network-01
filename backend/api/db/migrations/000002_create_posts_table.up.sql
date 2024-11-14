CREATE TABLE IF NOT EXISTS posts (
    id TEXT PRIMARY KEY,
    user_id TEXT REFERENCES users(id),
    group_id TEXT DEFAULT '00000000' REFERENCES groups(id),
    content TEXT NOT NULL,
    timestamp DATETIME NOT NULL
);

CREATE TABLE IF NOT EXISTS posts_images (
    post_id TEXT REFERENCES posts(id),
    path TEXT UNIQUE NOT NULL
);

CREATE TABLE IF NOT EXISTS likes_records(
    user_id TEXT REFERENCES users(id),
    post_id TEXT REFERENCES posts(id)
);

