CREATE TABLE IF NOT EXISTS events (
    id TEXT NOT NULL PRIMARY KEY,
    group_id TEXT REFERENCES groups(id),
    title TEXT,
    description TEXT,
    date DATETIME
);

CREATE TABLE IF NOT EXISTS events_records (
    event_id TEXT NOT NULL REFERENCES events(id),
    user_id TEXT NOT NULL REFERENCES users(id),
    UNIQUE(user_id, event_id)
);
