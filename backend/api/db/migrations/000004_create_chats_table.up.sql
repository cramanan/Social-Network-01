CREATE TABLE IF NOT EXISTS chats(
    sender_id TEXT REFERENCES users(id),
    recipient_id TEXT REFERENCES users(id),
    content TEXT NOT NULL,
    timestamp DATETIME NOT NULL
);