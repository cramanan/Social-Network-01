CREATE TABLE IF NOT EXISTS users (
    id TEXT PRIMARY KEY,
    nickname TEXT NOT NULL,
    email TEXT UNIQUE NOT NULL,
    password BLOB NOT NULL,
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    date_of_birth DATETIME NOT NULL,
    image_path TEXT NOT NULL DEFAULT "https://commons.wikimedia.org/wiki/File:Default_pfp.svg",
    about_me TEXT,
    private BOOLEAN NOT NULL,
    timestamp DATETIME NOT NULL
);

CREATE TABLE IF NOT EXISTS posts (
    id TEXT PRIMARY KEY,
    user_id TEXT REFERENCES users(id),
    group_id TEXT DEFAULT '00000000' REFERENCES groups(id),
    content TEXT NOT NULL,
    image_path TEXT,
    timestamp DATETIME NOT NULL
);

CREATE TABLE IF NOT EXISTS posts_status (
    post_id TEXT REFERENCES posts(id),
    status_enum INTEGER,
    users_ids BLOB
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
    user_id TEXT REFERENCES users(id),
    follower_id TEXT REFERENCES users(id)
);

CREATE TABLE IF NOT EXISTS likes_records(
    user_id TEXT REFERENCES users(id),
    post_id TEXT REFERENCES posts(id)
);

CREATE TABLE IF NOT EXISTS groups(
    id TEXT NOT NULL PRIMARY KEY,
    name TEXT UNIQUE,
    description TEXT NOT NULL,
    timestamp DATETIME NOT NULL
);

INSERT INTO groups VALUES(
    '00000000', 'Global', 'Global group', date('now')
);