CREATE TABLE IF NOT EXISTS groups(
    id TEXT NOT NULL PRIMARY KEY,
    name TEXT UNIQUE,
    owner TEXT NOT NULL REFERENCES users(id),
    description TEXT NOT NULL,
    image TEXT NOT NULL DEFAULT 'https://upload.wikimedia.org/wikipedia/commons/2/2c/Default_pfp.svg',
    timestamp DATETIME NOT NULL
);

CREATE TABLE IF NOT EXISTS groups_record (
    group_id TEXT NOT NULL REFERENCES groups(id),
    user_id TEXT NOT NULL REFERENCES users(id),
    is_request BOOLEAN NOT NULL,
    accepted BOOLEAN NOT NULL DEFAULT FALSE,

    UNIQUE(group_id, user_id, accepted)
);