BEGIN TRANSACTION;
CREATE TABLE IF NOT EXISTS users (
    id BLOB PRIMARY KEY,
    nickname TEXT NOT NULL,
    email TEXT UNIQUE NOT NULL,
    password BLOB NOT NULL,
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    date_of_birth DATE NOT NULL,
    image_path TEXT,
    about_me TEXT,
    timestamp DATE NOT NULL,
    private BOOLEAN NOT NULL
);

CREATE TABLE IF NOT EXISTS posts (
    id BLOB PRIMARY KEY,
    user_id BLOB REFERENCES users(id),
    group_id BLOB REFERENCES groups(id),
    category BLOB NOT NULL,
    timestamp DATE not NULL,
    content TEXT NOT NULL,
    image_path TEXT
);

CREATE TABLE IF NOT EXISTS comments(
    id BLOB PRIMARY KEY,
    user_id BLOB REFERENCES users(id),
    parent_id BLOB REFERENCES posts(id),
    timestamp DATE NOT NULL,
    content TEXT NOT NULL,
    img_path BLOB
);

CREATE TABLE IF NOT EXISTS chats(
    id BLOB PRIMARY KEY,
    sender_id BLOB REFERENCES users(id),
    recipient_id BLOB REFERENCES users(id),
    content TEXT NOT NULL,
    date_of_birth DATE NOT NULL
);

CREATE TABLE IF NOT EXISTS follow_records(
    id BLOB PRIMARY KEY,
    follower_id BLOB REFERENCES users(id),
    user_id BLOB REFERENCES users(id)
);

CREATE TABLE IF NOT EXISTS likes_records(
    user_id BLOB REFERENCES users(id),
    post_id BLOB REFERENCES posts(id)
);

CREATE TABLE IF NOT EXISTS groups(
    id BLOB PRIMARY KEY,
    name TEXT NOT NULL,
    users BLOB NOT NULL
);
COMMIT;