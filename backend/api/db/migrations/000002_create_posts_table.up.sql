CREATE TABLE IF NOT EXISTS posts (
    id TEXT PRIMARY KEY,
    user_id TEXT NOT NULL REFERENCES users(id),
    group_id TEXT REFERENCES groups(id),
    content TEXT NOT NULL,
    privacy_level TEXT NOT NULL CHECK (privacy_level IN ('public', 'private', 'almost_private')),
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


CREATE TABLE IF NOT EXISTS post_visibility(
    user_id TEXT REFERENCES users(id),
    post_id TEXT REFERENCES posts(id)
);

