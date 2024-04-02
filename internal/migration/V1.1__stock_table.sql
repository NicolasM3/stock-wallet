CREATE TABLE stock (
    id SERIAL PRIMARY KEY,
    code TEXT NOT NULL,
    name TEXT NOT NULL,
    current_price REAL NOT NULL,
    created_at TIMESTAMP NOT NULL,
    last_update TIMESTAMP NOT NULL,
    deleted BOOLEAN NOT NULL
);
