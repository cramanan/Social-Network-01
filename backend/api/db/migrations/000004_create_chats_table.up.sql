CREATE TABLE IF NOT EXISTS chats(
    sender_id TEXT REFERENCES users(id),
    recipient_id TEXT REFERENCES users(id),
    content TEXT NOT NULL,
    timestamp DATETIME NOT NULL
);

CREATE TABLE IF NOT EXISTS group_chats(
    sender_id TEXT REFERENCES users(id),
    group_id TEXT REFERENCES groups(id),
    content TEXT NOT NULL,
    timestamp DATETIME NOT NULL
);
