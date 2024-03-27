CREATE TABLE stock (
    code TEXT PRIMARY KEY,
    name TEXT NOT NULL,
    current_price REAL NOT NULL,
    created_at TIMESTAMP NOT NULL,
    last_update TIMESTAMP NOT NULL,
    deleted BOOLEAN NOT NULL
);
