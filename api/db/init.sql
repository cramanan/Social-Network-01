BEGIN TRANSACTION;
CREATE TABLE IF NOT EXISTS users (
    id TEXT PRIMARY KEY,
    nickname TEXT NOT NULL,
    email TEXT UNIQUE NOT NULL,
    password BLOB NOT NULL,
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    date_of_birth DATETIME NOT NULL,
    image_path TEXT,
    about_me TEXT,
    private BOOLEAN NOT NULL,
    timestamp DATETIME NOT NULL
);

CREATE TABLE IF NOT EXISTS posts (
    id TEXT PRIMARY KEY,
    user_id TEXT REFERENCES users(id),
    group_id TEXT DEFAULT '00000000-0000-0000-0000-000000000000' REFERENCES groups(id),
    category BLOB NOT NULL,
    content TEXT NOT NULL,
    image_path TEXT,
    timestamp DATETIME NOT NULL
);

CREATE TABLE IF NOT EXISTS comments(
    id TEXT PRIMARY KEY,
    user_id TEXT REFERENCES users(id),
    parent_id TEXT REFERENCES posts(id),
    content TEXT NOT NULL,
    img_path BLOB,
    timestamp DATETIME NOT NULL
);

CREATE TABLE IF NOT EXISTS chats(
    id TEXT PRIMARY KEY,
    sender_id TEXT REFERENCES users(id),
    recipient_id TEXT REFERENCES users(id),
    content TEXT NOT NULL,
    img_path TEXT,
    timestamp DATETIME NOT NULL
);

CREATE TABLE IF NOT EXISTS follow_records(
    user_id TEXT REFERENCES users(id)
    follower_id TEXT REFERENCES users(id),
);

CREATE TABLE IF NOT EXISTS likes_records(
    user_id TEXT REFERENCES users(id),
    post_id TEXT REFERENCES posts(id)
);

CREATE TABLE IF NOT EXISTS groups(
    id TEXT PRIMARY KEY,
    name TEXT NOT NULL,
    users BLOB NOT NULL
);
COMMIT;