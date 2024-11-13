CREATE TABLE IF NOT EXISTS events (
    id TEXT NOT NULL PRIMARY KEY,
    group_id TEXT REFERENCES groups(id),
    title TEXT,
    description TEXT,
    date DATETIME
);

CREATE TABLE IF NOT EXISTS events_records (
    event_id NOT NULL TEXT REFERENCES events(id),
    user_id NOT NULL TEXT REFERENCES users(id),
    going BOOLEAN NOT NULL
);