CREATE TABLE sessions (
    id SERIAL PRIMARY KEY,
    user_id INT UNIUE,
    token_hash TEXT UNIQUE NOT NULL
);